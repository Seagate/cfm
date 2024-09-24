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
	ApplianceData map[string]*ApplianceDatum `json:"appliance-data"`
	HostData      map[string]*HostDatum      `json:"host-data"`
}

func NewDataStore() *DataStore {
	return &DataStore{
		ApplianceData: make(map[string]*ApplianceDatum),
		HostData:      make(map[string]*HostDatum),
	}
}

// AddApplianceDatum: Add a new appliance datum to the data store
func (c *DataStore) AddApplianceDatum(creds *openapi.Credentials) {
	c.ApplianceData[creds.CustomId] = NewApplianceDatum(creds)
}

// AddHostDatum: Add a new host datum to the data store
func (c *DataStore) AddHostDatum(creds *openapi.Credentials) {
	c.HostData[creds.CustomId] = NewHostDatum(creds)
}

// DeleteApplianceDatumById: Delete an appliance from the data store
func (c *DataStore) DeleteApplianceDatumById(applianceId string) {
	delete(c.ApplianceData, applianceId)
}

// DeleteHostDatumById: Delete a host datum from the data store
func (c *DataStore) DeleteHostDatumById(hostId string) {
	delete(c.HostData, hostId)
}

// GetApplianceDatumById: Retrieve an appliance datum from the data store
func (c *DataStore) GetApplianceDatumById(applianceId string) (*ApplianceDatum, error) {
	datum, exists := c.ApplianceData[applianceId]
	if !exists {
		return nil, fmt.Errorf("appliance datum [%s] not found in data store", applianceId)
	}

	return datum, nil
}

// GetHostDatumById: Retrieve a host datum from the data store
func (c *DataStore) GetHostDatumById(hostId string) (*HostDatum, error) {
	datum, exists := c.HostData[hostId]
	if !exists {
		return nil, fmt.Errorf("host datum [%s] not found in data store", hostId)
	}

	return datum, nil
}

// Init: initialize the data store using command line args, ENV, or a file
func (c *DataStore) InitDataStore(ctx context.Context, args []string) error {

	DStore().Store()

	return nil
}

type ApplianceDatum struct {
	Credentials      *openapi.Credentials    `json:"credentials"`
	BladeData        map[string]*BladeDatum  `json:"blade-data"`
	ConnectionStatus common.ConnectionStatus `json:"connection-status"`
}

func NewApplianceDatum(creds *openapi.Credentials) *ApplianceDatum {
	return &ApplianceDatum{
		Credentials:      creds,
		BladeData:        make(map[string]*BladeDatum),
		ConnectionStatus: common.NOT_APPLICABLE, // Will use for single-BMC appliance
	}
}

func (a *ApplianceDatum) AddBladeDatum(creds *openapi.Credentials) {
	a.BladeData[creds.CustomId] = NewBladeDatum(creds)
}

func (a *ApplianceDatum) DeleteBladeDatumById(bladeId string) {
	delete(a.BladeData, bladeId)
}

func (a *ApplianceDatum) GetBladeDatumById(ctx context.Context, bladeId string) (*BladeDatum, error) {
	logger := klog.FromContext(ctx)

	blade, exists := a.BladeData[bladeId]
	if !exists {
		err := fmt.Errorf("blade datum [%s] not found in appliance data [%s] in data store", bladeId, a.Credentials.CustomId)
		logger.Error(err, "failure: update blade")
		return nil, err
	}

	return blade, nil
}

func (a *ApplianceDatum) SetConnectionStatus(status common.ConnectionStatus) {
	a.ConnectionStatus = status
}

type BladeDatum struct {
	Credentials      *openapi.Credentials    `json:"credentials"`
	ConnectionStatus common.ConnectionStatus `json:"connection-status"`
}

func NewBladeDatum(creds *openapi.Credentials) *BladeDatum {
	return &BladeDatum{
		Credentials:      creds,
		ConnectionStatus: common.ONLINE,
	}
}

func (b *BladeDatum) SetConnectionStatus(status *common.ConnectionStatus) {
	b.ConnectionStatus = *status
}

type HostDatum struct {
	Credentials      *openapi.Credentials    `json:"credentials"`
	ConnectionStatus common.ConnectionStatus `json:"connection-status"`
}

func NewHostDatum(creds *openapi.Credentials) *HostDatum {
	return &HostDatum{
		Credentials:      creds,
		ConnectionStatus: common.ONLINE,
	}
}

func (h *HostDatum) SetConnectionStatus(status *common.ConnectionStatus) {
	h.ConnectionStatus = *status
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
	for applianceId, applianceDatum := range c.ApplianceData {
		_, err = s.AppliancesPost(ctx, *applianceDatum.Credentials)
		if err != nil {
			logger.V(2).Info("cfm-service: appliance restore failure", "applianceId", applianceId)
			appliancesToDelete = append(appliancesToDelete, applianceId)
			continue
		}

		bladesToDelete := make(map[string]string)
		for bladeId, bladeDatum := range applianceDatum.BladeData {
			_, err = s.BladesPost(ctx, applianceId, *bladeDatum.Credentials)
			if err != nil {
				logger.V(2).Info("cfm-service: blade restore failure", "bladeId", bladeId, "applianceId", applianceId)
				bladesToDelete[applianceId] = bladeId
			}
		}

		for applianceId, bladeId := range bladesToDelete {
			delete(c.ApplianceData[applianceId].BladeData, bladeId)
		}
	}

	for _, applianceId := range appliancesToDelete {
		delete(c.ApplianceData, applianceId)
	}

	logger.V(2).Info("cfm-service: restoring saved hosts")
	for hostId, hostDatum := range c.HostData {
		_, err = s.HostsPost(ctx, *hostDatum.Credentials)
		if err != nil {
			logger.V(2).Info("cfm-service: host datum restore failure", "hostId", hostId)
			continue
		}
	}
}
