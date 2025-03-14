package services

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"

	"cfm/pkg/common"
	"cfm/pkg/common/datastore"
	"cfm/pkg/openapi"
)

// DiscoverDevices function to call the DiscoverDevices API
func DiscoverDevices(ctx context.Context, apiService openapi.DefaultAPIServicer, deviceType string) (openapi.ImplResponse, error) {
	resp, _ := apiService.DiscoverDevices(ctx, deviceType)
	if resp.Code >= 300 {
		err := fmt.Errorf("error discovering devices of type %s: %+v", deviceType, resp)
		log.Print(err)
		return resp, err
	} else {
		log.Printf("Discovered devices of type %s: %+v", deviceType, resp)
		return resp, nil
	}
}

func AddDiscoveredDevices(ctx context.Context, apiService openapi.DefaultAPIServicer, blades openapi.ImplResponse, hosts openapi.ImplResponse) {
	// Verify the existence of the default appliance; if it doesn't exist, add it
	datastore.DStore().Restore()
	data := datastore.DStore().GetDataStore()
	_, err := datastore.DStore().GetDataStore().GetApplianceDatumById(common.DefaultApplianceCredentials.CustomId)
	if err != nil {
		datastore.DStore().GetDataStore().AddApplianceDatum(common.DefaultApplianceCredentials)
		datastore.DStore().Store()
	}

	// Add blades
	bladeBodyBytes, ok := blades.Body.([]*openapi.DiscoveredDevice)
	if !ok {
		log.Fatalf("Response body is not []byte")
	}
	applianceDatum, _ := datastore.DStore().GetDataStore().GetApplianceDatumById(common.DefaultApplianceCredentials.CustomId)
	for _, bladeDevice := range bladeBodyBytes {
		_, exist := data.GetBladeDatumByIp(bladeDevice.Address)
		if !exist {
			newCredentials := *common.DefaultBladeCredentials
			newCredentials.IpAddress = bladeDevice.Address

			// Remove the .local suffix (e.g. blade device name: granite00.local) from the device name by splitting it with . and assign it to the customId
			deviceName := strings.SplitN(bladeDevice.Name, ".", 2)[0]
			newCredentials.CustomId = deviceName

			applianceDatum.AddBladeDatum(&newCredentials)
			datastore.DStore().Store()
		}
	}

	// Add cxl-hosts
	hostBodyBytes, ok := hosts.Body.([]*openapi.DiscoveredDevice)
	if !ok {
		log.Fatalf("Response body is not []byte")
	}
	for _, hostDevice := range hostBodyBytes {
		_, exist := data.GetHostDatumByIp(hostDevice.Address)
		if !exist {
			newCredentials := *common.DefaultHostCredentials
			newCredentials.IpAddress = hostDevice.Address

			// Remove the .local suffix (e.g. host device name: host00.local) from the device name by splitting it with . and assign it to the customId
			deviceName := strings.SplitN(hostDevice.Name, ".", 2)[0]
			newCredentials.CustomId = deviceName

			datastore.DStore().GetDataStore().AddHostDatum(&newCredentials)
			datastore.DStore().Store()
		}

	}
}
