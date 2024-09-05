// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package datastore

import (
	"context"
	"fmt"

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
	Status      ConnectionStatus
}

// UpdateAppliance: Update an appliance's data
func (c *DataStore) UpdateAppliance(r *ApplianceUpdateRequest) error {
	appliance, exists := c.SavedAppliances[r.ApplianceId]
	if !exists {
		return fmt.Errorf("appliance [%s] not found in data store appliance update", r.ApplianceId)
	}

	appliance.ConnectionStatus = r.Status

	return nil
}

type BladeUpdateRequest struct {
	ApplianceId string
	BladeId     string
	Status      ConnectionStatus
}

// UpdateBlade: Update an appliance's data
func (c *DataStore) UpdateBlade(r *BladeUpdateRequest) error {
	appliance, exists := c.SavedAppliances[r.ApplianceId]
	if !exists {
		return fmt.Errorf("appliance [%s] not found in data store blade update", r.ApplianceId)
	}

	appliance.UpdateBlade(r)

	return nil
}

type HostUpdateRequest struct {
	HostId string
	Status ConnectionStatus
}

// UpdateHost: Update a host's data
func (c *DataStore) UpdateHost(r *HostUpdateRequest) error {
	host, exists := c.SavedHosts[r.HostId]
	if !exists {
		return fmt.Errorf("host [%s] not found in data store host update", r.HostId)
	}

	host.ConnectionStatus = r.Status

	return nil
}

type ApplianceDataStore struct {
	Credentials      *openapi.Credentials       `json:"credentials"`
	SavedBlades      map[string]*BladeDataStore `json:"saved-blades"`
	ConnectionStatus ConnectionStatus
}

func NewApplianceDataStore(creds *openapi.Credentials) *ApplianceDataStore {
	return &ApplianceDataStore{
		Credentials: creds,
		SavedBlades: make(map[string]*BladeDataStore),
	}
}

func (a *ApplianceDataStore) AddBlade(creds *openapi.Credentials) {
	a.SavedBlades[creds.CustomId] = NewBladeDataStore(creds)
}

func (a *ApplianceDataStore) DeleteBlade(bladeId string) {
	delete(a.SavedBlades, bladeId)
}

func (a *ApplianceDataStore) UpdateBlade(r *BladeUpdateRequest) error {
	blade, exists := a.SavedBlades[r.BladeId]
	if !exists {
		return fmt.Errorf("blade [%s] not found in data store blade update", r.BladeId)
	}

	blade.ConnectionStatus = r.Status

	return nil
}

type ConnectionStatus string

const (
	Active   ConnectionStatus = "active"
	Inactive ConnectionStatus = "inactive"
)

type BladeDataStore struct {
	Credentials      *openapi.Credentials `json:"credentials"`
	ConnectionStatus ConnectionStatus
}

func NewBladeDataStore(creds *openapi.Credentials) *BladeDataStore {
	return &BladeDataStore{
		Credentials:      creds,
		ConnectionStatus: Active,
	}
}

type HostDataStore struct {
	Credentials      *openapi.Credentials `json:"credentials"`
	ConnectionStatus ConnectionStatus
}

func NewHostDataStore(creds *openapi.Credentials) *HostDataStore {
	return &HostDataStore{
		Credentials:      creds,
		ConnectionStatus: Active,
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

	logger.V(2).Info("cfm-service: restoring saved hosts")
	var hostsToDelete []string
	for hostId, host := range c.SavedHosts {
		_, err = s.HostsPost(ctx, *host.Credentials)
		if err != nil {
			logger.V(2).Info("cfm-service: host restore failure", "hostId", hostId)
			hostsToDelete = append(hostsToDelete, hostId)
			continue
		}
	}

	for _, hostId := range hostsToDelete {
		delete(c.SavedHosts, hostId)
	}
}
