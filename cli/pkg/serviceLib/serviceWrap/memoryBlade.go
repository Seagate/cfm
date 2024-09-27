// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

// Send compose request via cfm-service client
func BladesAssignMemory(client *service.APIClient, applianceId, bladeId, memoryId, portId, operation string) (*service.MemoryRegion, error) {

	// create new AssignMemoryRequest
	assignRequest := service.NewAssignMemoryRequest(portId, operation)

	// create new ApiBladesAssignMemoryByIdRequest
	bladeRequest := client.DefaultAPI.BladesAssignMemoryById(context.Background(), applianceId, bladeId, memoryId)

	// add AssignMemoryRequest to ApiBladesAssignMemoryByIdRequest
	bladeRequest = bladeRequest.AssignMemoryRequest(*assignRequest)

	// Execute ApiBladesAssignMemoryByIdRequest
	region, response, err := bladeRequest.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", bladeRequest, err)
			klog.ErrorS(newErr, "failure: BladesAssignMemory")

			return nil, fmt.Errorf("failure: BladesAssignMemory: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			bladeRequest, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: BladesAssignMemory")

		return nil, fmt.Errorf("failure: BladesAssignMemory: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("sucess: BladesAssignMemory", "operation", operation, "memoryId", region.GetId(), "portId", region.GetMemoryAppliancePort(), "size", region.GetSizeMiB(), "applId", region.GetMemoryApplianceId(), "bladeId", region.GetMemoryBladeId())

	return region, nil
}
