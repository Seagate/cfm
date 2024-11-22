// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

func AddAppliance(client *service.APIClient, creds *service.Credentials) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesPost(context.Background())
	request = request.Credentials(*creds)
	appliance, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: AddAppliance", "applianceId", appliance.GetId())

	return appliance, nil
}

func DeleteApplianceById(client *service.APIClient, applId string) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesDeleteById(context.Background(), applId)
	appliance, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: DeleteApplianceById", "applianceId", appliance.GetId())

	return appliance, nil
}

func GetAllAppliances(client *service.APIClient) (*[]*service.Appliance, error) {
	var appliances []*service.Appliance

	//Get existing appliances
	request := client.DefaultAPI.AppliancesGet(context.Background())
	collection, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(4).InfoS("success: AppliancesGet", "count", collection.GetMemberCount())

	for _, member := range collection.GetMembers() {
		id := ReadLastItemFromUri(member.GetUri())
		request2 := client.DefaultAPI.AppliancesGetById(context.Background(), id)
		appliance, response2, err := request2.Execute()
		if response2 != nil {
			defer response2.Body.Close() // Required by http lib implementation.
		}
		if err != nil {
			newErr := handleServiceError(response2, err)
			return nil, fmt.Errorf("execute failure(%T): %w", request2, newErr)
		}

		klog.V(4).InfoS("success: AppliancesGetById", "applianceId", appliance.GetId())

		appliances = append(appliances, appliance)
	}

	klog.V(3).InfoS("success: GetAllAppliances", "count", len(appliances))

	return &appliances, nil
}

func RenameApplianceById(client *service.APIClient, applianceId string, newApplianceId string) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesUpdateById(context.Background(), applianceId)
	request = request.NewApplianceId(newApplianceId)
	appliance, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: RenameApplianceById", "request", request)

	return appliance, nil
}

func ResyncApplianceById(client *service.APIClient, applianceId string) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesResyncById(context.Background(), applianceId)
	appliance, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: ResyncApplianceById", "applianceId", appliance.GetId())

	return appliance, nil
}
