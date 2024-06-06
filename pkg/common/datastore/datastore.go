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

// DeleteHost: Delete an host from the data store
func (c *DataStore) DeleteHost(hostId string) {
	delete(c.SavedHosts, hostId)
}

// Init: initialize the data store using command line args, ENV, or a file
func (c *DataStore) InitDataStore(ctx context.Context, args []string) error {

	DStore().Store()

	return nil
}

type ApplianceDataStore struct {
	Credentials *openapi.Credentials       `json:"credentials"`
	SavedBlades map[string]*BladeDataStore `json:"saved-blades"`
}

func NewApplianceDataStore(creds *openapi.Credentials) *ApplianceDataStore {
	return &ApplianceDataStore{
		Credentials: creds,
		SavedBlades: make(map[string]*BladeDataStore),
	}
}

func (c *ApplianceDataStore) AddBlade(creds *openapi.Credentials) {
	c.SavedBlades[creds.CustomId] = NewBladeDataStore(creds)
}

func (c *ApplianceDataStore) DeleteBlade(bladeId string) {
	delete(c.SavedBlades, bladeId)
}

type BladeDataStore struct {
	Credentials *openapi.Credentials `json:"credentials"`
}

func NewBladeDataStore(creds *openapi.Credentials) *BladeDataStore {
	return &BladeDataStore{
		Credentials: creds,
	}
}

type HostDataStore struct {
	Credentials *openapi.Credentials `json:"credentials"`
}

func NewHostDataStore(creds *openapi.Credentials) *HostDataStore {
	return &HostDataStore{
		Credentials: creds,
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
		appliance.Credentials.CustomId = applianceId
		_, err = s.AppliancesPost(ctx, *appliance.Credentials)
		if err != nil {
			logger.V(2).Info("cfm-service: appliance restore failure", "applianceId", applianceId)
			appliancesToDelete = append(appliancesToDelete, applianceId)
			continue
		}

		bladesToDelete := make(map[string]string)
		for bladeId, blade := range appliance.SavedBlades {
			blade.Credentials.CustomId = bladeId
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
		host.Credentials.CustomId = hostId
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
