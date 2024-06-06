/*
Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

*/

package serviceWrap

import (
	"context"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

func AddAppliance(client *service.APIClient, creds *service.Credentials) (*service.Appliance, error) {
	newReqAddAppl := client.DefaultAPI.AppliancesPost(context.Background())
	newReqAddAppl = newReqAddAppl.Credentials(*creds)
	addedAppl, response, err := newReqAddAppl.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", newReqAddAppl)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: appliances post: %s", err)
	}

	klog.V(3).InfoS("AppliancesPost success", "ID", addedAppl.GetId())

	return addedAppl, nil
}

func DeleteApplianceById(client *service.APIClient, applId string) (*service.Appliance, error) {
	newReqDelApplById := client.DefaultAPI.AppliancesDeleteById(context.Background(), applId)
	deletedAppl, response, err := newReqDelApplById.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", newReqDelApplById)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: delete appliance by id failure")
	}

	klog.V(3).InfoS("AppliancesDeleteById success", "ID", deletedAppl.GetId())

	return deletedAppl, nil
}

func GetAllAppliances(client *service.APIClient) (*[]*service.Appliance, error) {
	var appliances []*service.Appliance

	//Get existing appliances
	newReqGetAppls := client.DefaultAPI.AppliancesGet(context.Background())
	collection, response, err := newReqGetAppls.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", newReqGetAppls)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: appliances get: %s", err)
	}

	klog.V(3).InfoS("AppliancesGet success", "count", collection.GetMemberCount())

	for _, member := range collection.GetMembers() {
		id := ReadLastItemFromUri(member.GetUri())
		newReqGetApplById := client.DefaultAPI.AppliancesGetById(context.Background(), id)
		appliance, response, err := newReqGetApplById.Execute()
		if err != nil {
			msg := fmt.Sprintf("%T: Execute FAILURE", newReqGetApplById)
			klog.ErrorS(err, msg, "response", response)
			return nil, fmt.Errorf("failure: appliances get by id: %s", err)
		}

		klog.V(3).InfoS("AppliancesGetById success", "ID", appliance.GetId())

		appliances = append(appliances, appliance)
	}

	klog.V(4).InfoS("Discovered appliances", "count", len(appliances))

	return &appliances, nil
}
