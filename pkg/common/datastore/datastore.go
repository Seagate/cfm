// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package datastore

import (
	"context"
	"fmt"

	"cfm/pkg/common"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

const (
	DefaultDataStoreFile = "cfmdatastore.json" // Default JSON datastore file
)

type DataStore struct {
	SavedAppliances map[string]*ApplianceDataStore `json:"saved-appliances"`
	SavedHosts      map[string]*HostDataStore      `json:"saved-hosts"`
}

func NewDataStore() *DataStore {
	return &DataStore{
		SavedAppliances: make(map[string]*ApplianceDataStore),
		SavedHosts:      make(map[string]*HostDataStore),
	}
}

// AddAppliance: Add a new appliance to the data store
func (c *DataStore) AddAppliance(creds *openapi.Credentials) {
	c.SavedAppliances[creds.CustomId] = NewApplianceDataStore(creds)
}

// AddBlade: Add a new blade to the data store
func (c *DataStore) AddBlade(creds *openapi.Credentials, applianceId string) error {
	appliance, exists := c.SavedAppliances[applianceId]
	if !exists {
		return fmt.Errorf("appliance [%s] not found in data store add", applianceId)
	}

	appliance.AddBlade(creds)

	return nil
}

// AddHost: Add a new host to the data store
func (c *DataStore) AddHost(creds *openapi.Credentials) {
	c.SavedHosts[creds.CustomId] = NewHostDataStore(creds)
}

// ContainsAppliance: Checks to see if the supplied IDs represent a valid datastore Appliance
func (c *DataStore) ContainsAppliance(applianceId string) bool {
	_, exists := c.SavedAppliances[applianceId]

	return exists
}

// ContainsBlade: Checks to see if the supplied IDs represent a valid datastore Blade
func (c *DataStore) ContainsBlade(bladeId, applianceId string) bool {
	appliance, exists := c.SavedAppliances[applianceId]
	if !exists {
		return false
	}

	_, exists = appliance.SavedBlades[bladeId]

	return exists
}

// ContainsHost: Checks to see if the supplied IDs represent a valid datastore Host
func (c *DataStore) ContainsHost(hostId string) bool {
	_, exists := c.SavedHosts[hostId]

	return exists
}

// DeleteAppliance: Delete an appliance from the data store
func (c *DataStore) DeleteAppliance(applianceId string) {
	delete(c.SavedAppliances, applianceId)
}

// DeleteBlade: Delete a blade from an appliance's data store
func (c *DataStore) DeleteBlade(bladeId, applianceId string) error {
	appliance, exists := c.SavedAppliances[applianceId]
	if !exists {
		return fmt.Errorf("appliance [%s] not found in data store delete", applianceId)
	}

	appliance.DeleteBlade(bladeId)

	return nil
}

// DeleteHost: Delete a host from the data store
func (c *DataStore) DeleteHost(hostId string) {
	delete(c.SavedHosts, hostId)
}

// Init: initialize the data store using command line args, ENV, or a file
func (c *DataStore) InitDataStore(ctx context.Context, args []string) error {

	DStore().Store()

	return nil
}

type ApplianceUpdateRequest struct {
	ApplianceId string
	Status      common.ConnectionStatus
}

// UpdateAppliance: Update an appliance's data
func (c *DataStore) UpdateAppliance(ctx context.Context, r *ApplianceUpdateRequest) error {
	logger := klog.FromContext(ctx)

	appliance, exists := c.SavedAppliances[r.ApplianceId]
	if !exists {
		err := fmt.Errorf("appliance [%s] not found in data store during appliance update", r.ApplianceId)
		logger.Error(err, "failure: update appliance")
		return err
	}

	if r.Status != "" {
		appliance.ConnectionStatus = r.Status
	}

	return nil
}

type BladeUpdateRequest struct {
	ApplianceId string
	BladeId     string
	Status      common.ConnectionStatus
}

// UpdateBlade: Update an appliance's data
func (c *DataStore) UpdateBlade(ctx context.Context, r *BladeUpdateRequest) error {
	logger := klog.FromContext(ctx)

	appliance, exists := c.SavedAppliances[r.ApplianceId]
	if !exists {
		err := fmt.Errorf("appliance [%s] not found in data store during blade update", r.ApplianceId)
		logger.Error(err, "failure: update blade")
		return err
	}

	return appliance.UpdateBlade(ctx, r)
}

type UpdateHostStatusRequest struct {
	HostId string
	Status common.ConnectionStatus
}

// UpdateHost: Update a host's status
func (c *DataStore) UpdateHostStatus(ctx context.Context, r *UpdateHostStatusRequest) error {
	logger := klog.FromContext(ctx)

	host, exists := c.SavedHosts[r.HostId]
	if !exists {
		err := fmt.Errorf("host [%s] not found in data store during host status update", r.HostId)
		logger.Error(err, "failure: update host status")
		return err
	}

	host.ConnectionStatus = r.Status

	return nil
}

type ApplianceDataStore struct {
	Credentials      *openapi.Credentials       `json:"credentials"`
	SavedBlades      map[string]*BladeDataStore `json:"saved-blades"`
	ConnectionStatus common.ConnectionStatus    `json:"connection-status"`
}

func NewApplianceDataStore(creds *openapi.Credentials) *ApplianceDataStore {
	return &ApplianceDataStore{
		Credentials:      creds,
		SavedBlades:      make(map[string]*BladeDataStore),
		ConnectionStatus: common.NOT_APPLICABLE,
	}
}

func (a *ApplianceDataStore) AddBlade(creds *openapi.Credentials) {
	a.SavedBlades[creds.CustomId] = NewBladeDataStore(creds)
}

func (a *ApplianceDataStore) DeleteBlade(bladeId string) {
	delete(a.SavedBlades, bladeId)
}

func (a *ApplianceDataStore) UpdateBlade(ctx context.Context, r *BladeUpdateRequest) error {
	logger := klog.FromContext(ctx)

	blade, exists := a.SavedBlades[r.BladeId]
	if !exists {
		err := fmt.Errorf("blade [%s] not found in data store during blade update", r.BladeId)
		logger.Error(err, "failure: update blade")
		return err
	}

	if r.Status != "" {
		blade.ConnectionStatus = r.Status
	}

	return nil
}

type BladeDataStore struct {
	Credentials      *openapi.Credentials    `json:"credentials"`
	ConnectionStatus common.ConnectionStatus `json:"connection-status"`
}

func NewBladeDataStore(creds *openapi.Credentials) *BladeDataStore {
	return &BladeDataStore{
		Credentials:      creds,
		ConnectionStatus: common.ONLINE,
	}
}

type HostDataStore struct {
	Credentials      *openapi.Credentials    `json:"credentials"`
	ConnectionStatus common.ConnectionStatus `json:"connection-status"`
}

func NewHostDataStore(creds *openapi.Credentials) *HostDataStore {
	return &HostDataStore{
		Credentials:      creds,
		ConnectionStatus: common.ONLINE,
	}
}

////////////////////////////////////////
///////////// Helpers //////////////////
////////////////////////////////////////

// ReloadDataStore: Loads the saved data store information back into cfm-service
func ReloadDataStore(ctx context.Context, s openapi.DefaultAPIServicer, c *DataStore) {
	var err error

	logger := klog.FromContext(ctx)

	logger.V(2).Info("cfm-service: restoring saved appliances")
	var appliancesToDelete []string
	for applianceId, appliance := range c.SavedAppliances {
		_, err = s.AppliancesPost(ctx, *appliance.Credentials)
		if err != nil {
			logger.V(2).Info("cfm-service: appliance restore failure", "applianceId", applianceId)
			appliancesToDelete = append(appliancesToDelete, applianceId)
			continue
		}

		bladesToDelete := make(map[string]string)
		for bladeId, blade := range appliance.SavedBlades {
			_, err = s.BladesPost(ctx, applianceId, *blade.Credentials)
			if err != nil {
				logger.V(2).Info("cfm-service: blade restore failure", "bladeId", bladeId, "applianceId", applianceId)
				bladesToDelete[applianceId] = bladeId
			}
		}

		for applianceId, bladeId := range bladesToDelete {
			delete(c.SavedAppliances[applianceId].SavedBlades, bladeId)
		}
	}

	for _, applianceId := range appliancesToDelete {
		delete(c.SavedAppliances, applianceId)
	}

	logger.V(2).Info("cfm-service: reloading saved hosts")
	for hostId, host := range c.SavedHosts {
		_, err = s.HostsPost(ctx, *host.Credentials)
		if err != nil {
			logger.V(2).Info("cfm-service: host reload failure", "hostId", hostId)
			continue
		}
	}
}
