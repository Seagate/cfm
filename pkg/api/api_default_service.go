// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
/*
Composable Fabric Manager Service OpenAPI

This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.

API version: 1.6.1
Generated by: OpenAPI Generator (https://openapi-generator.tech)
*/

package api

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"cfm/pkg/common"
	"cfm/pkg/manager"
	"cfm/pkg/openapi"
)

const (
	MAX_COUNT_APPLIANCES = 32
	MAX_COUNT_BLADES     = 8
	MAX_COUNT_HOSTS      = 32
)

// CfmApiService is a service that implements the logic for the cfm-service.
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type CfmApiService struct {
	Version string
}

// NewCfmApiService creates a default api service
func NewCfmApiService(version string) openapi.DefaultAPIServicer {
	return &CfmApiService{Version: version}
}

// AppliancesDeleteById -
func (cfm *CfmApiService) AppliancesDeleteById(ctx context.Context, applianceId string) (openapi.ImplResponse, error) {
	var a openapi.Appliance

	appliance, err := manager.DeleteApplianceById(ctx, applianceId)
	if err != nil {
		a = openapi.Appliance{Id: applianceId}

	} else {

		a = openapi.Appliance{
			Id:        appliance.Id,
			IpAddress: "", // Unused (May need for POC4)
			Port:      0,  // Unused (May need for POC4)
			Status:    "", // Unused
			Blades: openapi.MemberItem{
				Uri: manager.GetCfmUriBlades(appliance.Id),
			},
			TotalMemoryAvailableMiB: -1, // Not Implemented
			TotalMemoryAllocatedMiB: -1, // Not Implemented
		}
	}

	return openapi.Response(http.StatusOK, a), nil
}

// AppliancesGet -
func (cfm *CfmApiService) AppliancesGet(ctx context.Context) (openapi.ImplResponse, error) {
	// order returned uris by appliance id
	applianceIds := manager.GetAllApplianceIds()
	sort.Strings(applianceIds)

	appliances := manager.GetAppliances(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(appliances)),
	}
	for _, id := range applianceIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: appliances[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// AppliancesGetById -
func (cfm *CfmApiService) AppliancesGetById(ctx context.Context, applianceId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	totals, err := appliance.GetResourceTotals(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	a := openapi.Appliance{
		Id:        appliance.Id,
		IpAddress: "", // Unused
		Port:      0,  // Unused
		Status:    "", // Unused
		Blades: openapi.MemberItem{
			Uri: manager.GetCfmUriBlades(appliance.Id),
		},
		TotalMemoryAvailableMiB: totals.TotalMemoryAvailableMiB,
		TotalMemoryAllocatedMiB: totals.TotalMemoryAllocatedMiB,
	}

	return openapi.Response(http.StatusOK, a), nil
}

// AppliancesPost -
func (cfm *CfmApiService) AppliancesPost(ctx context.Context, credentials openapi.Credentials) (openapi.ImplResponse, error) {
	appliances := manager.GetAppliances(ctx)
	if len(appliances) >= MAX_COUNT_APPLIANCES {
		err := common.RequestError{
			StatusCode: common.StatusAppliancesExceedMaximum,
			Err:        fmt.Errorf("cfm-service at maximum appliance capacity (%d)", MAX_COUNT_APPLIANCES),
		}
		return formatErrorResp(ctx, &err)
	}

	appliance, err := manager.AddAppliance(ctx, &credentials)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	a := openapi.Appliance{
		Id:        appliance.Id,
		IpAddress: "", // Unused
		Port:      0,  // Unused
		Status:    "", // Unused
		Blades: openapi.MemberItem{
			Uri: manager.GetCfmUriBlades(appliance.Id),
		},
		TotalMemoryAvailableMiB: 0, // Currently, new appliances have no initial knowledge of their blades
		TotalMemoryAllocatedMiB: 0, // Currently, new appliances have no initial knowledge of their blades
	}

	return openapi.Response(http.StatusCreated, a), nil
}

// AppliancesRenameById -
func (cfm *CfmApiService) AppliancesRenameById(ctx context.Context, applianceId string, newApplianceId string) (openapi.ImplResponse, error) {
	// Make sure the newApplianceId doesn't exist
	_, exist := manager.GetApplianceById(ctx, newApplianceId)
	if exist == nil {
		err := common.RequestError{
			StatusCode: common.StatusApplianceIdDuplicate,
			Err:        fmt.Errorf("the new name (%s) already exists", newApplianceId),
		}
		return formatErrorResp(ctx, &err)
	}

	// Make sure the appliance exists
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	//Rename the appliance with the new id
	newAppliance, err := manager.RenameAppliance(ctx, appliance, newApplianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, newAppliance), nil
}

// AppliancesResync -
func (cfm *CfmApiService) AppliancesResyncById(ctx context.Context, applianceId string) (openapi.ImplResponse, error) {
	appliance, err := manager.ResyncApplianceById(ctx, applianceId)

	if err != nil {
		if appliance != nil {
			return openapi.Response(http.StatusPartialContent, appliance), nil
		} else {
			return formatErrorResp(ctx, err.(*common.RequestError))
		}
	}

	totals, err := appliance.GetResourceTotals(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	a := openapi.Appliance{
		Id:        appliance.Id,
		IpAddress: "", // Unused
		Port:      0,  // Unused
		Status:    "", // Unused
		Blades: openapi.MemberItem{
			Uri: manager.GetCfmUriBlades(appliance.Id),
		},
		TotalMemoryAvailableMiB: totals.TotalMemoryAvailableMiB,
		TotalMemoryAllocatedMiB: totals.TotalMemoryAllocatedMiB,
	}

	return openapi.Response(http.StatusOK, a), nil
}

// BladesAssignMemoryById - Assign\Unassign the specified memory region (iedntified via memoryId) to\from the specified portId.  Assign\Unassigned is set in request's Operation parameter.
func (cfm *CfmApiService) BladesAssignMemoryById(ctx context.Context, applianceId string, bladeId string, memoryId string, assignMemoryRequest openapi.AssignMemoryRequest) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	r := manager.RequestAssignMemory{
		MemoryId:  memoryId,
		PortId:    assignMemoryRequest.Port,
		Operation: assignMemoryRequest.Operation,
	}

	memory, err := blade.AssignMemory(ctx, &r)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, memory), nil
}

// BladesComposeMemory - Using available blade resources, provision a new blade memory region and assign it to the designated port [Note: provision + assign = compose]. If composeMemoryRequest.Port == nil, provision only.
func (cfm *CfmApiService) BladesComposeMemory(ctx context.Context, applianceId string, bladeId string, composeMemoryRequest openapi.ComposeMemoryRequest) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	r := manager.RequestComposeMemory{
		PortId:  composeMemoryRequest.Port,
		SizeMib: composeMemoryRequest.MemorySizeMiB,
		Qos:     composeMemoryRequest.QoS,
	}

	memory, err := blade.ComposeMemory(ctx, &r)
	if err != nil {
		if memory != nil {
			return openapi.Response(http.StatusPartialContent, memory), nil
		} else {
			return formatErrorResp(ctx, err.(*common.RequestError))
		}
	}

	return openapi.Response(http.StatusCreated, memory), nil
}

// BladesComposeMemoryByResource -
func (cfm *CfmApiService) BladesComposeMemoryByResource(ctx context.Context, applianceId string, bladeId string, composeMemoryByResourceRequest openapi.ComposeMemoryByResourceRequest) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	memory, err := blade.ComposeMemoryByResource(ctx, composeMemoryByResourceRequest.Port, composeMemoryByResourceRequest.MemoryResources)
	if err != nil {
		if memory != nil {
			return openapi.Response(http.StatusPartialContent, memory), nil
		} else {
			return formatErrorResp(ctx, err.(*common.RequestError))
		}
	}

	return openapi.Response(http.StatusCreated, memory), nil
}

// BladesDeleteById - As long as the appliance id is valid, guarenteed blade deletion from service
func (cfm *CfmApiService) BladesDeleteById(ctx context.Context, applianceId string, bladeId string) (openapi.ImplResponse, error) {
	var b openapi.Blade

	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.DeleteBladeById(ctx, bladeId)
	if err != nil {
		b = openapi.Blade{Id: bladeId}

	} else {

		b = openapi.Blade{
			Id:        blade.Id,
			IpAddress: blade.GetNetIp(),
			Port:      int32(blade.GetNetPort()),
			Status:    string(blade.Status),
			Ports: openapi.MemberItem{
				Uri: manager.GetCfmUriBladePorts(appliance.Id, blade.Id),
			},
			Resources: openapi.MemberItem{
				Uri: manager.GetCfmUriBladeResources(appliance.Id, blade.Id),
			},
			Memory: openapi.MemberItem{
				Uri: manager.GetCfmUriBladeMemory(appliance.Id, blade.Id),
			},
			TotalMemoryAvailableMiB: -1, // Not Implemented
			TotalMemoryAllocatedMiB: -1, // Not Implemented
		}
	}

	return openapi.Response(http.StatusOK, b), nil
}

// BladesFreeMemoryById -
func (cfm *CfmApiService) BladesFreeMemoryById(ctx context.Context, applianceId string, bladeId string, memoryId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	memory, err := blade.FreeMemoryById(ctx, memoryId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, memory), nil
}

// BladesGet -
func (cfm *CfmApiService) BladesGet(ctx context.Context, applianceId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by blade id
	bladeIds := appliance.GetAllBladeIds()
	sort.Strings(bladeIds)

	blades := appliance.GetBlades(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(blades)),
	}
	for _, id := range bladeIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: blades[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// BladesGetById -
func (cfm *CfmApiService) BladesGetById(ctx context.Context, applianceId string, bladeId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	totals, err := blade.GetResourceTotals(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	b := openapi.Blade{
		Id:        blade.Id,
		IpAddress: blade.GetNetIp(),
		Port:      int32(blade.GetNetPort()),
		Status:    string(blade.Status),
		Ports: openapi.MemberItem{
			Uri: manager.GetCfmUriBladePorts(appliance.Id, blade.Id),
		},
		Resources: openapi.MemberItem{
			Uri: manager.GetCfmUriBladeResources(appliance.Id, blade.Id),
		},
		Memory: openapi.MemberItem{
			Uri: manager.GetCfmUriBladeMemory(appliance.Id, blade.Id),
		},
		TotalMemoryAvailableMiB: totals.TotalMemoryAvailableMiB,
		TotalMemoryAllocatedMiB: totals.TotalMemoryAllocatedMiB,
	}

	return openapi.Response(http.StatusOK, b), nil
}

// BladesGetMemory -
func (cfm *CfmApiService) BladesGetMemory(ctx context.Context, applianceId string, bladeId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by memory id
	memoryIds := blade.GetAllMemoryIds(ctx)
	sort.Strings(memoryIds)

	memory := blade.GetMemory(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(memory)),
	}
	for _, id := range memoryIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: memory[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// BladesGetMemoryById -
func (cfm *CfmApiService) BladesGetMemoryById(ctx context.Context, applianceId string, bladeId string, memoryId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	memory, err := blade.GetMemoryById(ctx, memoryId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	if memory != nil {
		details, err := memory.GetDetails(ctx)
		if err != nil {
			return formatErrorResp(ctx, err.(*common.RequestError))
		}

		return openapi.Response(http.StatusOK, details), nil

	} else {
		return openapi.Response(http.StatusOK, openapi.MemoryRegion{}), nil
	}
}

// BladesGetPortById -
func (cfm *CfmApiService) BladesGetPortById(ctx context.Context, applianceId string, bladeId string, portId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	port, err := blade.GetPortById(ctx, portId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	if port != nil {
		details, err := port.GetDetails(ctx)
		if err != nil {
			return formatErrorResp(ctx, err.(*common.RequestError))
		}

		return openapi.Response(http.StatusOK, details), nil

	} else {
		return openapi.Response(http.StatusOK, openapi.PortInformation{}), nil
	}
}

// BladesGetPorts -
func (cfm *CfmApiService) BladesGetPorts(ctx context.Context, applianceId string, bladeId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by port id
	portIds := blade.GetAllPortIds(ctx)
	sort.Strings(portIds)

	ports := blade.GetPorts(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(ports)),
	}
	for _, id := range portIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: ports[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// BladesRenameById -
func (cfm *CfmApiService) BladesRenameById(ctx context.Context, applianceId string, bladeId string, newBladeId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// Make sure the bladeId exists
	// Get the blade information from the manager level and is used for renaming
	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// Make sure the newBladeId doesn't exist
	existBladeIds := appliance.GetAllBladeIds()
	for _, id := range existBladeIds {
		if newBladeId == id {
			err := common.RequestError{
				StatusCode: common.StatusBladeIdDuplicate,
				Err:        fmt.Errorf("the new name (%s) already exists", newBladeId),
			}
			return formatErrorResp(ctx, &err)
		}
	}

	//Rename the appliance with the new id
	newBlade, err := manager.RenameBlade(ctx, blade, newBladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, newBlade), nil
}

// BladesGetResourceById -
func (cfm *CfmApiService) BladesGetResourceById(ctx context.Context, applianceId string, bladeId string, resourceId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	resource, err := blade.GetResourceById(ctx, resourceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	if resource != nil {
		details, err := resource.GetDetails(ctx)
		if err != nil {
			return formatErrorResp(ctx, err.(*common.RequestError))
		}

		return openapi.Response(http.StatusOK, details), nil

	} else {
		return openapi.Response(http.StatusOK, openapi.MemoryResourceBlock{}), nil
	}
}

// BladesGetResources -
func (cfm *CfmApiService) BladesGetResources(ctx context.Context, applianceId string, bladeId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.GetBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by resourse id
	resourceIds := blade.GetAllResourceIds(ctx)
	sort.Strings(resourceIds)

	resources := blade.GetResources(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(resources)),
	}
	for _, id := range resourceIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: resources[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// BladesPost -
func (cfm *CfmApiService) BladesPost(ctx context.Context, applianceId string, credentials openapi.Credentials) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	if len(appliance.Blades) == MAX_COUNT_BLADES {
		err := common.RequestError{
			StatusCode: common.StatusBladesExceedMaximum,
			Err:        fmt.Errorf("cfm-service at maximum blade capacity (%d) for this appliance (%s)", MAX_COUNT_BLADES, applianceId),
		}
		return formatErrorResp(ctx, &err)
	}

	blade, err := appliance.AddBlade(ctx, &credentials)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	totals, err := blade.GetResourceTotals(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	b := openapi.Blade{
		Id:        blade.Id,
		IpAddress: blade.GetNetIp(),
		Port:      int32(blade.GetNetPort()),
		Status:    string(blade.Status),
		Ports: openapi.MemberItem{
			Uri: manager.GetCfmUriBladePorts(appliance.Id, blade.Id),
		},
		Resources: openapi.MemberItem{
			Uri: manager.GetCfmUriBladeResources(appliance.Id, blade.Id),
		},
		Memory: openapi.MemberItem{
			Uri: manager.GetCfmUriBladeMemory(appliance.Id, blade.Id),
		},
		TotalMemoryAvailableMiB: totals.TotalMemoryAvailableMiB,
		TotalMemoryAllocatedMiB: totals.TotalMemoryAllocatedMiB,
	}

	return openapi.Response(http.StatusCreated, b), nil
}

// BladesResync -
func (cfm *CfmApiService) BladesResyncById(ctx context.Context, applianceId string, bladeId string) (openapi.ImplResponse, error) {
	appliance, err := manager.GetApplianceById(ctx, applianceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	blade, err := appliance.ResyncBladeById(ctx, bladeId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	totals, err := blade.GetResourceTotals(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	b := openapi.Blade{
		Id:        blade.Id,
		IpAddress: blade.GetNetIp(),
		Port:      int32(blade.GetNetPort()),
		Status:    string(blade.Status),
		Ports: openapi.MemberItem{
			Uri: manager.GetCfmUriBladePorts(appliance.Id, blade.Id),
		},
		Resources: openapi.MemberItem{
			Uri: manager.GetCfmUriBladeResources(appliance.Id, blade.Id),
		},
		Memory: openapi.MemberItem{
			Uri: manager.GetCfmUriBladeMemory(appliance.Id, blade.Id),
		},
		TotalMemoryAvailableMiB: totals.TotalMemoryAvailableMiB,
		TotalMemoryAllocatedMiB: totals.TotalMemoryAllocatedMiB,
	}

	return openapi.Response(http.StatusOK, b), nil
}

// HostGetMemory -
func (cfm *CfmApiService) HostGetMemory(ctx context.Context, hostId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by memory id
	memoryIds := host.GetAllMemoryIds(ctx)
	sort.Strings(memoryIds)

	memory := host.GetMemory(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(memory)),
	}
	for _, id := range memoryIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: memory[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// HostsComposeMemory -
func (cfm *CfmApiService) HostsComposeMemory(ctx context.Context, hostId string, composeMemoryRequest openapi.ComposeMemoryRequest) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	r := manager.RequestComposeMemory{
		PortId:  composeMemoryRequest.Port,
		SizeMib: composeMemoryRequest.MemorySizeMiB,
		Qos:     composeMemoryRequest.QoS,
	}

	memory, err := host.ComposeMemory(ctx, &r)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusCreated, memory), nil
}

// HostsDeleteById - Guarenteed host deletion from service.
func (cfm *CfmApiService) HostsDeleteById(ctx context.Context, hostId string) (openapi.ImplResponse, error) {
	var h openapi.Host

	host, err := manager.DeleteHostById(ctx, hostId)
	if err != nil && host == nil {
		h = openapi.Host{Id: hostId}

	} else {

		h = openapi.Host{
			Id:        host.Id,
			IpAddress: host.GetNetIp(),
			Port:      int32(host.GetNetPort()),
			Status:    string(host.Status),
			Ports: openapi.MemberItem{
				Uri: manager.GetCfmUriHostPorts(host.Id),
			},
			Memory: openapi.MemberItem{
				Uri: manager.GetCfmUriHostMemory(host.Id),
			},
			MemoryDevices: openapi.MemberItem{
				Uri: manager.GetCfmUriHostMemoryDevices(host.Id),
			},
			LocalMemoryMiB:  -1, // Not implemented
			RemoteMemoryMiB: -1, // Not implemented
		}
	}

	return openapi.Response(http.StatusOK, h), nil
}

// HostsFreeMemoryById -
func (cfm *CfmApiService) HostsFreeMemoryById(ctx context.Context, hostId string, memoryId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	memory, err := host.FreeMemoryById(ctx, memoryId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, memory), nil
}

// HostsGet - Get CXL Host information.
func (cfm *CfmApiService) HostsGet(ctx context.Context) (openapi.ImplResponse, error) {
	// order returned uris by host id
	hostIds := manager.GetAllHostIds()
	sort.Strings(hostIds)

	hosts := manager.GetHosts(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(hosts)),
	}
	for _, id := range hostIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: hosts[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// HostsGetById - Get information for a single CXL Host.
func (cfm *CfmApiService) HostsGetById(ctx context.Context, hostId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil || host == nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	totals, err := host.GetMemoryTotals(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	h := openapi.Host{
		Id:        host.Id,
		IpAddress: host.GetNetIp(),
		Port:      int32(host.GetNetPort()),
		Status:    string(host.Status),
		Ports: openapi.MemberItem{
			Uri: manager.GetCfmUriHostPorts(host.Id),
		},
		MemoryDevices: openapi.MemberItem{
			Uri: manager.GetCfmUriHostMemoryDevices(host.Id),
		},
		Memory: openapi.MemberItem{
			Uri: manager.GetCfmUriHostMemory(host.Id),
		},
		LocalMemoryMiB:  totals.LocalMemoryMib,
		RemoteMemoryMiB: totals.RemoteMemoryMib,
	}

	return openapi.Response(http.StatusOK, h), nil
}

// HostsGetMemoryById -
func (cfm *CfmApiService) HostsGetMemoryById(ctx context.Context, hostId string, memoryId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	memory, err := host.GetMemoryById(ctx, memoryId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	details, err := memory.GetDetails(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, details), nil
}

// HostsGetMemoryDeviceById -
func (cfm *CfmApiService) HostsGetMemoryDeviceById(ctx context.Context, hostId string, memorydeviceId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	memdev, err := host.GetMemoryDeviceById(ctx, memorydeviceId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	details, err := memdev.GetDetails(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, details), nil
}

// HostsGetMemoryDevices -
func (cfm *CfmApiService) HostsGetMemoryDevices(ctx context.Context, hostId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by memory device id
	memdevIds := host.GetAllMemoryDeviceIds(ctx)
	sort.Strings(memdevIds)

	memdevs := host.GetMemoryDevices(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(memdevs)),
	}
	for _, id := range memdevIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: memdevs[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// HostsGetPortById -
func (cfm *CfmApiService) HostsGetPortById(ctx context.Context, hostId string, portId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	port, err := host.GetPortById(ctx, portId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	details, err := port.GetDetails(ctx)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, details), nil
}

// HostsGetPorts -
func (cfm *CfmApiService) HostsGetPorts(ctx context.Context, hostId string) (openapi.ImplResponse, error) {
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// order returned uris by port id
	portIds := host.GetAllPortIds(ctx)
	sort.Strings(portIds)

	ports := host.GetPorts(ctx)

	response := openapi.Collection{
		MemberCount: int32(len(ports)),
	}
	for _, id := range portIds {
		response.Members = append(response.Members, openapi.MemberItem{
			Uri: ports[id].Uri,
		})
	}

	return openapi.Response(http.StatusOK, response), nil
}

// HostsPost - Add a CXL host to be managed by CFM.
func (cfm *CfmApiService) HostsPost(ctx context.Context, credentials openapi.Credentials) (openapi.ImplResponse, error) {
	hosts := manager.GetHosts(ctx)
	if len(hosts) >= MAX_COUNT_HOSTS {
		err := common.RequestError{
			StatusCode: common.StatusHostsExceedMaximum,
			Err:        fmt.Errorf("cfm-service at maximum host capacity (%d)", MAX_COUNT_HOSTS),
		}
		return formatErrorResp(ctx, &err)
	}

	host, err := manager.AddHost(ctx, &credentials)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	h := openapi.Host{
		Id:        host.Id,
		IpAddress: host.GetNetIp(),
		Port:      int32(host.GetNetPort()),
		Status:    string(host.Status),
		Ports: openapi.MemberItem{
			Uri: manager.GetCfmUriHostPorts(host.Id),
		},
		Memory: openapi.MemberItem{
			Uri: manager.GetCfmUriHostMemory(host.Id),
		},
		MemoryDevices: openapi.MemberItem{
			Uri: manager.GetCfmUriHostMemoryDevices(host.Id),
		},
		LocalMemoryMiB:  -1, // Not implemented
		RemoteMemoryMiB: -1, // Not implemented
	}

	return openapi.Response(http.StatusCreated, h), nil
}

// HostsRenameById -
func (cfm *CfmApiService) HostsRenameById(ctx context.Context, hostId string, newHostId string) (openapi.ImplResponse, error) {
	// Make sure the hostId exists
	// Get the host information from the manager level and is used for renaming
	host, err := manager.GetHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	// Make sure the newHostId doesn't exist
	_, exist := manager.GetHostById(ctx, newHostId)
	if exist == nil {
		err := common.RequestError{
			StatusCode: common.StatusHostIdDuplicate,
			Err:        fmt.Errorf("the new name (%s) already exists", newHostId),
		}
		return formatErrorResp(ctx, &err)
	}

	//Rename the cxl host with the new id
	newHost, err := manager.RenameHost(ctx, host, newHostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	return openapi.Response(http.StatusOK, newHost), nil
}

// HostsResync -
func (cfm *CfmApiService) HostsResyncById(ctx context.Context, hostId string) (openapi.ImplResponse, error) {
	host, err := manager.ResyncHostById(ctx, hostId)
	if err != nil {
		return formatErrorResp(ctx, err.(*common.RequestError))
	}

	h := openapi.Host{
		Id:        host.Id,
		IpAddress: host.GetNetIp(),
		Port:      int32(host.GetNetPort()),
		Status:    string(host.Status),
		Ports: openapi.MemberItem{
			Uri: manager.GetCfmUriHostPorts(host.Id),
		},
		Memory: openapi.MemberItem{
			Uri: manager.GetCfmUriHostMemory(host.Id),
		},
		MemoryDevices: openapi.MemberItem{
			Uri: manager.GetCfmUriHostMemoryDevices(host.Id),
		},
		LocalMemoryMiB:  -1, // Not implemented
		RemoteMemoryMiB: -1, // Not implemented
	}

	return openapi.Response(http.StatusOK, h), nil
}

func formatErrorResp(ctx context.Context, re *common.RequestError) (openapi.ImplResponse, error) {
	// Use the Go language type assertion to convert the returned enhanced error into the RequestError type
	status := openapi.StatusMessage{
		Uri:     common.GetContextString(ctx, common.KeyUri),
		Details: re.Error(),
		Status: openapi.StatusMessageStatus{
			Code:    int32(re.StatusCode),
			Message: re.StatusCode.String(),
		},
	}
	return openapi.Response(re.StatusCode.HttpStatusCode(), status), nil
}

type CfmVersionType struct {
	Version string `json:"v1"`
}

// CfmGet -
func (cfm *CfmApiService) CfmGet(ctx context.Context) (openapi.ImplResponse, error) {
	response := CfmVersionType{
		Version: "/cfm/v1",
	}

	return openapi.Response(http.StatusOK, response), nil
}

// CfmV1Get -
func (cfm *CfmApiService) CfmV1Get(ctx context.Context) (openapi.ImplResponse, error) {
	response := openapi.ServiceInformation{
		Version: cfm.Version,
	}

	response.Resources = append(response.Resources, openapi.ServiceResource{
		Uri:         manager.GetCfmUriAppliances(),
		Methods:     "GET, POST",
		Description: "Manage retrieval and addition of memory appliances under cfm-service management",
	})

	response.Resources = append(response.Resources, openapi.ServiceResource{
		Uri:         manager.GetCfmUriHosts(),
		Methods:     "GET, POST",
		Description: "Manage retrieval and addition of CXL hosts under cfm-service management",
	})

	return openapi.Response(http.StatusOK, response), nil
}

// RootGet -
func (cfm *CfmApiService) RootGet(ctx context.Context) (openapi.ImplResponse, error) {
	response := openapi.StatusMessage{
		Uri:     "/",
		Details: fmt.Sprintf("Composable Fabric Manager (CFM) Service API. Use 'http get /cfm' to see supported versions."),
		Status: openapi.StatusMessageStatus{
			Code:    int32(common.StatusOK),
			Message: common.StatusOK.String(),
		},
	}
	return openapi.Response(http.StatusOK, response), nil
}
