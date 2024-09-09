// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/openapi"

	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

const ID_PREFIX_APPLIANCE_DFLT string = "memory-appliance"

type Appliance struct {
	Id     string
	Uri    string
	Blades map[string]*Blade
	// BackendOps backend.BackendOperations	// For POC4?
}

// NewAppliance - Creates a new Appliance object.
func NewAppliance(ctx context.Context, c *openapi.Credentials) (*Appliance, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewAppliance: ")

	applianceId := c.CustomId
	if applianceId == "" {
		// Generate uuid here and combine the last N digits with the prefix to be the appliance default id
		uuid := uuid.New().String()
		applianceId = fmt.Sprintf("%s-%s", ID_PREFIX_APPLIANCE_DFLT, uuid[(len(uuid)-common.NumUuidCharsForId):])
		c.CustomId = applianceId
	}

	// Check for duplicate ID
	_, exists := deviceCache.GetApplianceByIdOk(applianceId)
	if exists {
		newErr := fmt.Errorf("invalid id: applianceId [%s] already exists in cfm-service", applianceId)
		return nil, &common.RequestError{StatusCode: common.StatusApplianceIdDuplicate, Err: newErr}
	}

	a := Appliance{
		Id:     applianceId,
		Uri:    GetCfmUriApplianceId(applianceId),
		Blades: make(map[string]*Blade),
	}

	logger.V(2).Info("success: new appliance", "applianceId", applianceId)

	return &a, nil
}

// AddBlade: Open a new session with a blade, create the new Blade object and then cache it
func (a *Appliance) AddBlade(ctx context.Context, c *openapi.Credentials) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> AddBlade: ", "applianceId", a.Id, "cred", c)

	backendName := ""
	if c.Backend != "" {
		backendName = c.Backend
	} else {
		backendName = common.GetContextString(ctx, common.KeyBackend)
	}
	logger.V(2).Info("found blade backend", "backend", backendName)

	ops, err := backend.NewBackendInterface(backendName, nil)
	if err != nil || ops == nil {
		newErr := fmt.Errorf("failed to initialize backend interface [%s]: %w", backendName, err)
		logger.Error(newErr, "failure: add blade")
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
		logger.Error(newErr, "failure: add blade")
		return nil, &common.RequestError{StatusCode: common.StatusBladeCreateSessionFailure, Err: newErr}
	}

	bladeId := c.CustomId
	if bladeId == "" { // Order CustomeId > BladeSN > UUID
		bladeId = response.ChassisSN
		if bladeId == "" {
			// Generate default id using last N digits of the session id combined with the default prefix
			bladeId = fmt.Sprintf("%s-%s", ID_PREFIX_BLADE_DFLT, response.SessionId[(len(response.SessionId)-common.NumUuidCharsForId):])
		}
		c.CustomId = bladeId
	}

	// Check for duplicate ID
	_, exists := a.Blades[bladeId]
	if exists {
		req := backend.DeleteSessionRequest{}
		response, err := ops.DeleteSession(ctx, &settings, &req)
		if err != nil || response == nil {
			newErr := fmt.Errorf("failed to delete session [%s:%d] after failed duplicate bladeId [%s] check: %w", c.IpAddress, c.Port, bladeId, err)
			logger.Error(newErr, "failure: add blade")
			return nil, &common.RequestError{StatusCode: common.StatusBladeDeleteSessionFailure, Err: newErr}
		}

		newErr := fmt.Errorf("invalid id: bladeId [%s] already exists on appliance [%s] ", bladeId, a.Id)
		logger.Error(newErr, "failure: add blade")
		return nil, &common.RequestError{StatusCode: common.StatusBladeIdDuplicate, Err: newErr}
	}

	// Create the new Blade
	r := RequestNewBlade{
		BladeId:     bladeId,
		ApplianceId: a.Id,
		Ip:          c.IpAddress,
		Port:        uint16(c.Port),
		BackendOps:  ops,
		Creds:       c,
	}

	blade, err := NewBlade(ctx, &r)
	if err != nil || blade == nil {
		req := backend.DeleteSessionRequest{}
		response, deleErr := ops.DeleteSession(ctx, &settings, &req)
		if deleErr != nil || response == nil {
			newErr := fmt.Errorf("failed to delete session [%s:%d] after failed blade [%s] object creation: %w", c.IpAddress, c.Port, bladeId, err)
			logger.Error(newErr, "failure: add blade")
			return nil, &common.RequestError{StatusCode: common.StatusBladeDeleteSessionFailure, Err: newErr}
		}

		newErr := fmt.Errorf("appliance [%s] new blade object creation failure: %w", a.Id, err)
		logger.Error(newErr, "failure: add blade")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	// Add blade to appliance
	a.Blades[blade.Id] = blade

	logger.V(2).Info("success: add blade", "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

func (a *Appliance) DeleteAllBlades(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteAllBlades: ", "applianceId", a.Id)

	for id := range a.Blades {
		_, err := a.DeleteBladeById(ctx, id)
		if err != nil {
			return err
		}
	}

	logger.V(2).Info("success: delete all blades", "applianceId", a.Id)

	return nil
}

// DeleteBladeById: Delete the blade backend session and the local blade cache
func (a *Appliance) DeleteBladeById(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteBladeById: ", "bladeId", bladeId, "applianceId", a.Id)

	// query for blade
	blade, ok := a.Blades[bladeId]
	if !ok {
		logger.V(2).Info("blade not found during delete:", "bladeId", bladeId, "applianceId", a.Id)
		newErr := fmt.Errorf("blade [%s] not found during delete", bladeId)
		return nil, &common.RequestError{StatusCode: common.StatusBladeIdDoesNotExist, Err: newErr}
	}

	// get blade backend
	ops := blade.backendOps

	// delete the blade session
	settings := backend.ConfigurationSettings{}
	req := backend.DeleteSessionRequest{}

	response, err := ops.DeleteSession(ctx, &settings, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("failed to delete blade [%s] backend [%s] session [%s]: %w", blade.Id, ops.GetBackendInfo(ctx).BackendName, blade.Socket.String(), err)
		logger.Error(newErr, "failure: delete blade by id")

		// Currently, backend ALWAYS deletes the blade session from the backend map.  For now, need to delete blade from appliance map as well.
		logger.V(2).Info("force blade deletion after backend session failure", "bladeId", blade.Id, "applianceId", a.Id)
		delete(a.Blades, blade.Id)

		return nil, &common.RequestError{StatusCode: common.StatusBladeDeleteSessionFailure, Err: newErr}
	}

	// delete blade
	delete(a.Blades, blade.Id)

	logger.V(2).Info("success: delete blade by id", "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

func (a *Appliance) GetAllBladeIds() []string {
	var ids []string

	for id := range a.Blades {
		ids = append(ids, id)
	}

	return ids
}

func (a *Appliance) GetBladeById(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetBladeById: ", "bladeId", bladeId, "applianceId", a.Id)

	blade, ok := a.Blades[bladeId]
	if !ok {
		newErr := fmt.Errorf("appliance [%s] blade [%s] doesn't exist", a.Id, bladeId)
		logger.Error(newErr, "failure: get blade by id")
		return nil, &common.RequestError{StatusCode: common.StatusBladeIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get blade by id", "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

func (a *Appliance) GetBlades(ctx context.Context) map[string]*Blade {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetBlades: ")

	blades := a.Blades

	logger.V(2).Info("success: get blades", "count", len(blades))

	return blades
}

type ResponseResourceTotals struct {
	TotalMemoryAvailableMiB int32
	TotalMemoryAllocatedMiB int32
}

func (a *Appliance) GetResourceTotals(ctx context.Context) (*ResponseResourceTotals, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetResourceTotals: ", "applianceId", a.Id)

	var totalAvail, totalAlloc int32

	for _, blade := range a.Blades {
		totals, err := blade.GetResourceTotals(ctx)
		if err != nil {
			newErr := fmt.Errorf("failed to get resource totals: appliance [%s] blade [%s]: %w", a.Id, blade.Id, err)
			logger.Error(newErr, "failure: get resource totals: appliance")
			return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
		}

		totalAvail += totals.TotalMemoryAvailableMiB
		totalAlloc += totals.TotalMemoryAllocatedMiB
	}

	response := ResponseResourceTotals{
		TotalMemoryAvailableMiB: totalAvail,
		TotalMemoryAllocatedMiB: totalAlloc,
	}

	logger.V(2).Info("success: get resource totals", "applianceId", a.Id)

	return &response, nil
}

func (a *Appliance) InvalidateCache() {
	for _, b := range a.Blades {
		b.InvalidateCache()
	}
}

func (a *Appliance) ResyncBladeById(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ResyncBladeById: ", "bladeId", bladeId, "applianceId", a.Id)

	blade, err := a.DeleteBladeById(ctx, bladeId)
	if err != nil {
		newErr := fmt.Errorf("failed to resync blade(delete): appliance [%s] blade [%s]: %w", a.Id, bladeId, err)
		logger.Error(newErr, "failure: resync blade: ignoring")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	blade, err = a.AddBlade(ctx, blade.creds)
	if err != nil {
		newErr := fmt.Errorf("failed to resync blade(add): appliance [%s] blade [%s]: %w", a.Id, bladeId, err)
		logger.Error(newErr, "failure: resync blade")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	logger.V(2).Info("success: resync blade", "bladeId", bladeId, "applianceId", a.Id)

	return blade, nil
}
