// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

func AddAppliance(client *service.APIClient, creds *service.Credentials) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesPost(context.Background())
	request = request.Credentials(*creds)
	appliance, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: AddAppliance", "applianceId", appliance.GetId())

	return appliance, nil
}

func DeleteApplianceById(client *service.APIClient, applId string) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesDeleteById(context.Background(), applId)
	appliance, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: DeleteApplianceById", "applianceId", appliance.GetId())

	return appliance, nil
}

func GetAllAppliances(client *service.APIClient) (*[]*service.Appliance, error) {
	var appliances []*service.Appliance

	//Get existing appliances
	requestGetAppls := client.DefaultAPI.AppliancesGet(context.Background())
	collection, response, err := requestGetAppls.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestGetAppls, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestGetAppls, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(4).InfoS("success: AppliancesGet", "count", collection.GetMemberCount())

	for _, member := range collection.GetMembers() {
		id := ReadLastItemFromUri(member.GetUri())
		requestGetApplById := client.DefaultAPI.AppliancesGetById(context.Background(), id)
		appliance, response, err := requestGetApplById.Execute()
		if err != nil {
			// Decode the JSON response into a struct
			var status service.StatusMessage
			if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
				newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestGetApplById, err)
				klog.V(4).Info(newErr)
				return nil, newErr
			}

			newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
				requestGetApplById, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		klog.V(4).InfoS("success: AppliancesGetById", "applianceId", appliance.GetId())

		appliances = append(appliances, appliance)
	}

	klog.V(3).InfoS("success: GetAllAppliances", "count", len(appliances))

	return &appliances, nil
}

func ResyncApplianceById(client *service.APIClient, applianceId string) (*service.Appliance, error) {
	request := client.DefaultAPI.AppliancesResyncById(context.Background(), applianceId)
	appliance, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: ResyncApplianceById", "applianceId", appliance.GetId())

	return appliance, nil
}
