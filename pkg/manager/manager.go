// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

/////////////
// Globals //
/////////////

const (
	SUCCESS         string = "success"
	FAILURE         string = "failure"
	UNKNOWN         string = "unknown"
	NOT_IMPLEMENTED string = "not implemented"
	NOT_APPLICABLE  string = "n\\a"
	NOT_SUPPORTED   string = "not supported"
)

var deviceCache *DevicesCache

// Perform one-time initialization steps
func init() {
	deviceCache = NewDevicesCache()
}

/////////////////////////
// Appliance Functions //
/////////////////////////

func AddAppliance(ctx context.Context, c *openapi.Credentials) (*Appliance, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> AddAppliance: ")

	// Create a new cfm-service Appliance object
	appliance, err := NewAppliance(ctx, c.CustomId)
	if err != nil || appliance == nil {
		newErr := fmt.Errorf("new appliance creation failure: %w", err)
		logger.Error(newErr, "failure: add appliance")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	// Cache it
	err = deviceCache.AddAppliance(appliance)
	if err != nil {
		newErr := fmt.Errorf("add appliance [%s] failure: %w", appliance.Id, err)
		logger.Error(newErr, "failure: add appliance")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	logger.V(2).Info("success: add appliance", "applianceId", appliance.Id)

	return appliance, nil
}

func DeleteApplianceById(ctx context.Context, applianceId string) (*Appliance, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteApplianceById: ", "applianceId", applianceId)

	// query cache
	appliance, ok := deviceCache.GetApplianceByIdOk(applianceId)
	if !ok {
		logger.V(2).Info("appliance not found during delete: do nothing", "applianceId", applianceId)
		return &Appliance{Id: applianceId}, nil
	}

	err := appliance.DeleteAllBlades(ctx) // cache and hardware interactions here
	if err != nil {
		newErr := fmt.Errorf("failed to delete all appliance [%s] blades: %w", appliance.Id, err)
		logger.Error(newErr, "failure: delete appliance by id")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	// delete appliance from cache
	a := deviceCache.DeleteApplianceById(appliance.Id)
	if a == nil {
		newErr := fmt.Errorf("appliance [%s] cache delete failed", appliance.Id)
		logger.Error(newErr, "failure: delete appliance by id")
		return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
	}

	logger.V(2).Info("success: delete appliance by id", "applianceId", a.Id)

	return a, nil
}

func GetAllApplianceIds() []string {
	return deviceCache.GetAllApplianceIds()
}

func GetApplianceById(ctx context.Context, applianceId string) (*Appliance, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetApplianceById: ", "applianceId", applianceId)

	appliance, err := deviceCache.GetApplianceById(applianceId)
	if err != nil {
		logger.Error(err, "failure: get appliance by id")
		return nil, &common.RequestError{StatusCode: common.StatusApplianceIdDoesNotExist, Err: err}
	}

	logger.V(2).Info("success: get appliance by id", "applianceId", applianceId)

	return appliance, nil
}

func GetAppliances(ctx context.Context) map[string]*Appliance {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetAppliances: ")

	appliances := deviceCache.GetAppliances()

	logger.V(2).Info("success: get appliances", "count", len(appliances))

	return appliances
}

////////////////////
// Host Functions //
////////////////////

func AddHost(ctx context.Context, c *openapi.Credentials) (*Host, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> AddHost: ", "cred", c)

	backendName := ""
	if c.Backend != "" {
		backendName = c.Backend
	} else {
		backendName = common.GetContextString(ctx, common.KeyBackend)
	}
	logger.V(2).Info("found host backend", "backend", backendName)

	ops, err := backend.NewBackendInterface(backendName, nil)
	if err != nil || ops == nil {
		newErr := fmt.Errorf("failed to initialize backend interface [%s]: %w", backendName, err)
		logger.Error(newErr, "failure: add host")
		return nil, &common.RequestError{StatusCode: common.StatusBackendInterfaceFailure, Err: newErr}
	}

	// Apply default value for optional Protocol field in the request
	if c.Protocol == "" {
		c.Protocol = "https"
	}

	req := backend.CreateSessionRequest{
		Ip:       c.IpAddress,
		Port:     c.Port,
		Username: c.Username,
		Password: c.Password,
		Insecure: c.Insecure,
		Protocol: c.Protocol,
	}

	settings := backend.ConfigurationSettings{}

	// Create a new session
	response, err := ops.CreateSession(ctx, &settings, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("create session failure at [%s:%d] using interface [%s]: %w", c.IpAddress, c.Port, backendName, err)
		logger.Error(newErr, "failure: add host")
		return nil, &common.RequestError{StatusCode: common.StatusHostCreateSessionFailure, Err: newErr}
	}

	hostId := c.CustomId
	if hostId == "" {
		// Generate default id using last N digits of the session id combined with the default prefix
		hostId = fmt.Sprintf("%s-%s", ID_PREFIX_HOST_DFLT, response.SessionId[(len(response.SessionId)-common.NumUuidCharsForId):])
	}

	// Check for duplicate ID.
	_, exists := deviceCache.GetHostByIdOk(hostId)
	if exists {
		req := backend.DeleteSessionRequest{}
		response, err := ops.DeleteSession(ctx, &settings, &req)
		if err != nil || response == nil {
			newErr := fmt.Errorf("failed to delete session [%s:%d] after failed duplicate hostId [%s] check: %w", c.IpAddress, c.Port, hostId, err)
			logger.Error(newErr, "failure: add host (during duplicate hostId check)")
			return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
		}

		newErr := fmt.Errorf("invalid id: hostId [%s] already exists in cfm-service", hostId)
		logger.Error(newErr, "failure: add host")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	// Create a new cfm-service Host object
	r := RequestNewHost{
		HostId:     hostId,
		Ip:         c.IpAddress,
		Port:       uint16(c.Port),
		BackendOps: ops,
	}

	host, err := NewHost(ctx, &r)
	if err != nil || host == nil {
		req := backend.DeleteSessionRequest{}
		response, err := ops.DeleteSession(ctx, &settings, &req)
		if err != nil || response == nil {
			newErr := fmt.Errorf("failed to delete session [%s:%d] after failed host [%s] object creation: %w", c.IpAddress, c.Port, hostId, err)
			logger.Error(newErr, "failure: add host")
			return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
		}

		newErr := fmt.Errorf("new host object creation failure: %w", err)
		logger.Error(newErr, "failure: add host")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	// Cache it
	deviceCache.AddHost(host)

	logger.V(2).Info("success: add host", "hostId", host.Id)

	return host, nil
}

func DeleteHostById(ctx context.Context, hostId string) (*Host, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteHostById: ", "hostId", hostId)

	// query cache
	host, ok := deviceCache.GetHostByIdOk(hostId)
	if !ok {
		logger.V(2).Info("host not found during delete: do nothing", "hostId", hostId)
		return &Host{Id: hostId}, nil
	}

	ops := host.backendOps

	// delete the session
	settings := backend.ConfigurationSettings{}
	req := backend.DeleteSessionRequest{}

	response, err := ops.DeleteSession(ctx, &settings, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("failed to delete host [%s] backend [%s] session [%s]: %w", host.Id, ops.GetBackendInfo(ctx).BackendName, host.Socket.String(), err)
		logger.Error(newErr, "failure: delete host by id")
		return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
	}

	// delete host from cache
	h := deviceCache.DeleteHostById(host.Id)
	if h == nil {
		newErr := fmt.Errorf("host [%s] cache delete failed", host.Id)
		logger.Error(newErr, "failure: delete host by id")
		return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
	}

	logger.V(2).Info("success: delete host by id", "hostId", h.Id)

	return h, nil
}

func GetAllHostIds() []string {
	return deviceCache.GetAllHostIds()
}

// GetHostById - Returns the manager.Host object containing the matching hostId.
func GetHostById(ctx context.Context, hostId string) (*Host, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetHostById: ", "hostId", hostId)

	host, err := deviceCache.GetHostById(hostId)
	if err != nil {
		logger.Error(err, "failure: get host by id")
		return nil, &common.RequestError{StatusCode: common.StatusHostIdDoesNotExist, Err: err}
	}

	logger.V(2).Info("success: get host by id", "hostId", hostId)

	return host, nil
}

func GetHosts(ctx context.Context) map[string]*Host {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetHosts: ")

	hosts := deviceCache.GetHosts()

	logger.V(2).Info("success: get hosts", "count", len(hosts))

	return hosts
}
