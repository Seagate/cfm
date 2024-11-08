// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/common/datastore"
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
		Status:      common.ONLINE,
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

	// Add blade to datastore
	applianceDatum, _ := datastore.DStore().GetDataStore().GetApplianceDatumById(a.Id)
	applianceDatum.AddBladeDatum(c)
	datastore.DStore().Store()

	logger.V(2).Info("success: add blade", "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

// ReplaceBladeById: Replace a pre-existing cached blade object with a new one.
// This function is used when a new new backend session is required for the blade.
func (a *Appliance) ReplaceBladeById(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ReplaceBladeById: ", "applianceId", a.Id, "bladeId", bladeId)

	// query for blade
	blade, ok := a.Blades[bladeId]
	if !ok {
		newErr := fmt.Errorf("appliance [%s] blade [%s] not found during replace by id", bladeId, a.Id)
		logger.Error(newErr, "failure: replace blade by id")

		return nil, &common.RequestError{StatusCode: common.StatusBladeIdDoesNotExist, Err: newErr}
	}

	creds := blade.creds
	ops := blade.backendOps

	req := backend.CreateSessionRequest{
		Ip:       creds.IpAddress,
		Port:     creds.Port,
		Username: creds.Username,
		Password: creds.Password,
		Insecure: creds.Insecure,
		Protocol: creds.Protocol,
	}

	settings := backend.ConfigurationSettings{}

	// Create a new session
	response, err := ops.CreateSession(ctx, &settings, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("create session failure at [%s:%d] using interface [%s]: %w", creds.IpAddress, creds.Port, ops.GetBackendInfo(ctx).BackendName, err)
		logger.Error(newErr, "failure: replace blade by id")
		return nil, &common.RequestError{StatusCode: common.StatusBladeCreateSessionFailure, Err: newErr}
	}

	// Create the new Blade
	r := RequestNewBlade{
		BladeId:     bladeId,
		ApplianceId: a.Id,
		Ip:          creds.IpAddress,
		Status:      common.ONLINE,
		Port:        uint16(creds.Port),
		BackendOps:  ops,
		Creds:       creds,
	}

	replacementBlade, err := NewBlade(ctx, &r)
	if err != nil || replacementBlade == nil {
		req := backend.DeleteSessionRequest{}
		response, deleErr := ops.DeleteSession(ctx, &settings, &req)
		if deleErr != nil || response == nil {
			newErr := fmt.Errorf("failed to delete session [%s:%d] after failed blade [%s] object creation: %w", creds.IpAddress, creds.Port, bladeId, err)
			logger.Error(newErr, "failure: replace blade by id")
			return nil, &common.RequestError{StatusCode: common.StatusBladeDeleteSessionFailure, Err: newErr}
		}

		newErr := fmt.Errorf("appliance [%s] new blade object creation failure: %w", a.Id, err)
		logger.Error(newErr, "failure: replace blade by id")
		return nil, &common.RequestError{StatusCode: common.StatusManagerInitializationFailure, Err: newErr}
	}

	// Replace blade in appliance
	a.Blades[blade.Id] = replacementBlade

	// Replace blade in datastore
	applianceDatum, _ := datastore.DStore().GetDataStore().GetApplianceDatumById(a.Id)
	applianceDatum.DeleteBladeDatumById(blade.Id)
	applianceDatum.AddBladeDatum(creds)
	datastore.DStore().Store()

	logger.V(2).Info("success: replace blade by id", "bladeId", replacementBlade.Id, "applianceId", a.Id)

	return replacementBlade, nil
}

func (a *Appliance) DeleteAllBlades(ctx context.Context) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteAllBlades: ", "applianceId", a.Id)

	for id := range a.Blades {
		a.DeleteBladeById(ctx, id) // ignore any errors
	}

	logger.V(2).Info("success: delete all blades", "applianceId", a.Id)
}

// DeleteBladeById: Delete the blade from: backend, deviceCache and datastore
// Function is designed to always delete the corresponding bladeId from ALL these locations, regardless of error.
func (a *Appliance) DeleteBladeById(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteBladeById: ", "bladeId", bladeId, "applianceId", a.Id)

	// Currently, backend ALWAYS deletes the blade session from the backend map.  Do the same in this (manager) layer too
	defer a.DeleteBladeByIdManager(ctx, bladeId) //Ensure this ALWAYS runs

	blade, err := a.DeleteBladeByIdBackend(ctx, bladeId)
	if err != nil || blade == nil {
		logger.V(2).Info("success: delete blade by id after backend session failure", "bladeId", blade.Id, "applianceId", a.Id)
		return blade, err
	}

	logger.V(2).Info("success: delete blade by id", "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

// DeleteBladeByIdBackend: Delete the blade from backend only
func (a *Appliance) DeleteBladeByIdBackend(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteBladeBackendById: ", "bladeId", bladeId, "applianceId", a.Id)

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
		logger.Error(newErr, "failure: delete blade by id (backend)")

		return blade, &common.RequestError{StatusCode: common.StatusBladeDeleteSessionFailure, Err: newErr} // Still return the blade for recovery
	}

	logger.V(2).Info("success: delete blade by id (backend)", "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

// DeleteBladeByIdManager: Delete the blade from manager layer only (appliance blade map and datastore)
func (a *Appliance) DeleteBladeByIdManager(ctx context.Context, bladeId string) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> DeleteBladeByIdManager: ", "bladeId", bladeId, "applianceId", a.Id)

	// delete blade from manager cache
	delete(a.Blades, bladeId)

	// delete blade from datastore
	applianceDatum, _ := datastore.DStore().GetDataStore().GetApplianceDatumById(a.Id)
	applianceDatum.DeleteBladeDatumById(bladeId)
	datastore.DStore().Store()
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
	var err error

	blade, ok := a.Blades[bladeId]
	if !ok {
		newErr := fmt.Errorf("appliance [%s] blade [%s] doesn't exist", a.Id, bladeId)
		logger.Error(newErr, "failure: get blade by id")
		return nil, &common.RequestError{StatusCode: common.StatusBladeIdDoesNotExist, Err: newErr}
	}

	// Check for resync
	if blade.CheckSync(ctx) {
		logger.V(4).Info("initiating auto-resync check", "bladeId", bladeId, "applianceId", a.Id)
		blade.UpdateConnectionStatusBackend(ctx)
		if blade.Status == common.FOUND { // good power, bad session
			blade, err = a.ResyncBladeById(ctx, bladeId)
			if err != nil {
				newErr := fmt.Errorf("failed to resync blade by id [%s]: %w", bladeId, err)
				logger.Error(newErr, "failure: get blade by id")
				return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
			}

			logger.V(2).Info("success: auto resync blade", "bladeId", bladeId)
		} else {
			blade.SetSync(ctx)
		}
	}

	logger.V(2).Info("success: get blade by id", "status", blade.Status, "bladeId", blade.Id, "applianceId", a.Id)

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

func (a *Appliance) AddBladeBack(ctx context.Context, c *openapi.Credentials) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> Add Blade Back: ", "bladeId", c.CustomId, "applianceId", a.Id)

	// add blade back
	blade, err := a.AddBlade(ctx, c)
	if err != nil {
		newErr := fmt.Errorf("failed to add blade [%s] back", c.CustomId)
		logger.Error(newErr, "failure: add blade back")
		return nil, &common.RequestError{StatusCode: common.StatusBladeCreateSessionFailure, Err: newErr}
	}

	blade.UpdateConnectionStatusBackend(ctx)

	logger.V(2).Info("success: add blade back", "status", blade.Status, "bladeId", blade.Id, "applianceId", a.Id)

	return blade, nil
}

func (a *Appliance) ResyncBladeById(ctx context.Context, bladeId string) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ResyncBladeById: ", "bladeId", bladeId, "applianceId", a.Id)

	// query device cache
	blade, ok := deviceCache.GetBladeByIdOk(a.Id, bladeId)
	if !ok || blade == nil {
		newErr := fmt.Errorf("failed to get blade [%s]", bladeId)
		logger.Error(newErr, "failure: resync blade by id")
		return nil, &common.RequestError{StatusCode: common.StatusBladeIdDoesNotExist, Err: newErr}
	}

	blade, err := a.DeleteBladeByIdBackend(ctx, bladeId)
	if err != nil {
		logger.Error(err, "resync blade by id: ignoring delete blade by id beackend failure")
	}

	blade.UpdateConnectionStatusBackend(ctx) // update status here in case of failure during replacement

	blade, err = a.ReplaceBladeById(ctx, blade.Id)
	if err != nil {
		newErr := fmt.Errorf("failed to replace blade by id: appliance [%s] blade [%s]: %w", a.Id, bladeId, err)
		logger.Error(newErr, "failure: resync blade by id")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	logger.V(2).Info("success: resync blade by id", "status", blade.Status, "bladeId", bladeId, "applianceId", a.Id)

	return blade, nil
}

/////////////////////////////////////
//////// Private Functions //////////
/////////////////////////////////////
