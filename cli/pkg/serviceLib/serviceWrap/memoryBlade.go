/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package serviceWrap

import (
	"context"
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
		newErr := fmt.Errorf("failure: blade %s memory: %w", operation, err)
		msg := fmt.Sprintf("%T: Execute FAILURE", bladeRequest)
		klog.ErrorS(newErr, msg, "response", response, "bladeRequest", bladeRequest)
		return nil, newErr
	}

	klog.V(3).InfoS("BladesAssignMemory success", "operation", operation, "memoryId", region.GetId(), "portId", region.GetMemoryAppliancePort(), "size", region.GetSizeMiB(), "applId", region.GetMemoryApplianceId(), "bladeId", region.GetMemoryBladeId())

	return region, nil
}
