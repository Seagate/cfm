// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/common/datastore"
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
	appliance, err := NewAppliance(ctx, c)
	if err != nil || appliance == nil {
		newErr := fmt.Errorf("new appliance creation failure: %w", err)
		logger.Error(newErr, "failure: add appliance")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	// Cache it
	err = deviceCache.AddAppliance(appliance)
	if err != nil {
		newErr := fmt.Errorf("add appliance [%s] failure: %w", appliance.Id, err)
		logger.Error(newErr, "failure: add appliance")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
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
		newErr := fmt.Errorf("failed to get appliance [%s]", applianceId)
		logger.Error(newErr, "failure: delete appliance by id")
		return nil, &common.RequestError{StatusCode: common.StatusApplianceIdDoesNotExist, Err: newErr}
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
		return nil, &common.RequestError{StatusCode: common.StatusApplianceDeleteSessionFailure, Err: newErr}
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
		newErr := fmt.Errorf("appliance [%s] doesn't exist", applianceId)
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
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

func ResyncApplianceById(ctx context.Context, applianceId string) (*Appliance, *[]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ResyncApplianceById: ", "applianceId", applianceId)

	var failedBladeIds []string

	appliance, err := deviceCache.GetApplianceById(applianceId)
	if err != nil {
		newErr := fmt.Errorf("get appliance by id [%s] failure: %w", appliance.Id, err)
		logger.Error(newErr, "failure: resync appliance by id")
		return nil, &failedBladeIds, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	bladeIds := appliance.GetAllBladeIds()
	for _, id := range bladeIds {
		_, err := appliance.ResyncBladeById(ctx, id)
		if err != nil {
			newErr := fmt.Errorf("resync blade by id [%s] failure: appliance [%s]: %w", id, appliance.Id, err)
			logger.Error(newErr, "failure: resync appliance by id: handle and continue")

			failedBladeIds = append(failedBladeIds, id)
		}
	}

	if len(failedBladeIds) == 0 {
		logger.V(2).Info("success: resync appliance", "applianceId", applianceId, "bladeIds", bladeIds)
		return appliance, &failedBladeIds, nil
	} else if len(failedBladeIds) < len(bladeIds) {
		newErr := fmt.Errorf("resync appliance by id [%s]: some failure(s): blade(s) [%s]: %w", appliance.Id, failedBladeIds, err)
		logger.Error(newErr, "partial success: resync appliance by id")
		return appliance, &failedBladeIds, &common.RequestError{StatusCode: common.StatusApplianceResyncPartialSuccess, Err: newErr}
	} else {
		newErr := fmt.Errorf("resync appliance by id [%s] failure: %w", appliance.Id, err)
		logger.Error(newErr, "failure: resync appliance by id")
		return nil, &failedBladeIds, &common.RequestError{StatusCode: common.StatusApplianceResyncFailure, Err: newErr}
	}
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
	if hostId == "" { // Order CustomeId > HostSN > UUID
		hostId = response.ChassisSN
		if hostId == "" {
			// Generate default id using last N digits of the session id combined with the default prefix
			hostId = fmt.Sprintf("%s-%s", ID_PREFIX_HOST_DFLT, response.SessionId[(len(response.SessionId)-common.NumUuidCharsForId):])
		}
		c.CustomId = hostId
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
		return nil, &common.RequestError{StatusCode: common.StatusHostIdDuplicate, Err: newErr}
	}

	// Create a new cfm-service Host object
	r := RequestNewHost{
		HostId:     hostId,
		Ip:         c.IpAddress,
		Port:       uint16(c.Port),
		Status:     common.ONLINE,
		BackendOps: ops,
		Creds:      c,
	}

	host, err := NewHost(ctx, &r)
	if err != nil || host == nil {
		req := backend.DeleteSessionRequest{}
		response, deleErr := ops.DeleteSession(ctx, &settings, &req)
		if deleErr != nil || response == nil {
			newErr := fmt.Errorf("failed to delete session [%s:%d] after failed host [%s] object creation: %w", c.IpAddress, c.Port, hostId, err)
			logger.Error(newErr, "failure: add host")
			return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
		}

		newErr := fmt.Errorf("new host object creation failure: %w", err)
		logger.Error(newErr, "failure: add host")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	// Add host to device cache
	deviceCache.AddHost(host)

	// Add host to datastore
	datastore.DStore().GetDataStore().AddHost(c)
	datastore.DStore().Store()

	logger.V(2).Info("success: add host", "hostId", host.Id)

	return host, nil
}

func DeleteHostById(ctx context.Context, hostId string) (*Host, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteHostById: ", "hostId", hostId)

	// query cache
	host, ok := deviceCache.GetHostByIdOk(hostId)
	if !ok {
		newErr := fmt.Errorf("failed to get host [%s]", hostId)
		logger.Error(newErr, "failure: delete host by id")
		return nil, &common.RequestError{StatusCode: common.StatusHostIdDoesNotExist, Err: newErr}
	}

	ops := host.backendOps

	// delete the session
	settings := backend.ConfigurationSettings{}
	req := backend.DeleteSessionRequest{}

	response, err := ops.DeleteSession(ctx, &settings, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("failed to delete host [%s] backend [%s] session [%s]: %w", host.Id, ops.GetBackendInfo(ctx).BackendName, host.Socket.String(), err)
		logger.Error(newErr, "failure: delete host by id")

		// Currently, backend ALWAYS deletes the host session from the backend map.
		// Delete host from manager cache and datastore as well
		logger.V(2).Info("force host deletion after backend delete session failure", "hostId", host.Id)
		deviceCache.DeleteHostById(host.Id)
		datastore.DStore().GetDataStore().DeleteHost(host.Id)
		datastore.DStore().Store()

		return nil, &common.RequestError{StatusCode: common.StatusHostDeleteSessionFailure, Err: newErr}
	}

	// delete host from manager cache
	h := deviceCache.DeleteHostById(host.Id)

	// delete host from datastore
	datastore.DStore().GetDataStore().DeleteHost(host.Id)
	datastore.DStore().Store()

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
		newErr := fmt.Errorf("failure: get host by id [%s]: %w", hostId, err)
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	host.UpdateConnectionStatusBackend(ctx)

	logger.V(2).Info("success: get host by id", "status", host.Status, "hostId", host.Id)

	return host, nil
}

func GetHosts(ctx context.Context) map[string]*Host {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetHosts: ")

	hosts := deviceCache.GetHosts()

	logger.V(2).Info("success: get hosts", "count", len(hosts))

	return hosts
}

func ResyncHostById(ctx context.Context, hostId string) (*Host, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ResyncHostById: ", "hostId", hostId)

	// query cache
	host, ok := deviceCache.GetHostByIdOk(hostId)
	if !ok || host == nil {
		newErr := fmt.Errorf("failed to get host [%s]", hostId)
		logger.Error(newErr, "failure: resync host by id")
		return nil, &common.RequestError{StatusCode: common.StatusHostIdDoesNotExist, Err: newErr}
	}

	host.UpdateConnectionStatusBackend(ctx)

	logger.V(2).Info("success: resync host", "hostId", host.Id)

	return host, nil
}
