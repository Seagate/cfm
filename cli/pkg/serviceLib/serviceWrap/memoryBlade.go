// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

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
	request := client.DefaultAPI.BladesAssignMemoryById(context.Background(), applianceId, bladeId, memoryId)

	// add AssignMemoryRequest to ApiBladesAssignMemoryByIdRequest
	request = request.AssignMemoryRequest(*assignRequest)

	// Execute ApiBladesAssignMemoryByIdRequest
	region, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("sucess: BladesAssignMemory", "operation", operation, "memoryId", region.GetId(), "portId", region.GetMemoryAppliancePort(), "size", region.GetSizeMiB(), "applId", region.GetMemoryApplianceId(), "bladeId", region.GetMemoryBladeId())

	return region, nil
}
