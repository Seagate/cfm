// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/common/datastore"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

const ID_PREFIX_BLADE_DFLT string = "blade"

type Blade struct {
	Id          string
	Uri         string
	Status      common.ConnectionStatus
	Socket      SocketDetails
	ApplianceId string
	Memory      map[string]*BladeMemory
	Ports       map[string]*CxlBladePort
	Resources   map[string]*BladeResource

	// Cached data
	NumResourceChannels int32 // Number of channels (dimms) detected on blade across all BladeResources
	ResourceSizeMib     int32

	// Backend access data
	backendOps        backend.BackendOperations
	creds             *openapi.Credentials // Used during resync
	lastSyncTimeStamp time.Time
}

type RequestNewBlade struct {
	BladeId     string
	Ip          string
	Port        uint16
	ApplianceId string
	Status      common.ConnectionStatus
	BackendOps  backend.BackendOperations
	Creds       *openapi.Credentials
}

func NewBlade(ctx context.Context, r *RequestNewBlade) (*Blade, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewBlade: ", "request", r)

	b := Blade{
		Id:          r.BladeId,
		Uri:         GetCfmUriBladeId(r.ApplianceId, r.BladeId),
		Socket:      *NewSocketDetails(r.Ip, r.Port),
		ApplianceId: r.ApplianceId,
		Status:      r.Status,
		Ports:       make(map[string]*CxlBladePort),
		Resources:   make(map[string]*BladeResource),
		Memory:      make(map[string]*BladeMemory),
		backendOps:  r.BackendOps,
		creds:       r.Creds,
	}

	err := b.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("init blade [%s] failure: %w", b.Id, err)
		logger.Error(newErr, "failure: new blade")
		return nil, newErr
	}

	b.SetSync(ctx)

	logger.V(2).Info("success: new blade", "bladeId", b.Id, "applianceId", b.ApplianceId)

	return &b, nil
}

func (b *Blade) SetSync(ctx context.Context) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> SetSync: ", "bladeId", b.Id)
	b.lastSyncTimeStamp = time.Now()
}

func (b *Blade) CheckSync(ctx context.Context) bool {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> CheckSync: ", "bladeId", b.Id)

	if time.Since(b.lastSyncTimeStamp).Seconds() > common.SyncCheckTimeoutSeconds {
		b.SetSync(ctx) // renew the timestamp
		return true
	}
	return false
}

type RequestAssignMemory struct {
	MemoryId  string
	PortId    string
	Operation string
}

func (b *Blade) AssignMemory(ctx context.Context, r *RequestAssignMemory) (*openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> AssignMemory: ", "memoryId", r.MemoryId, "portId", r.PortId, "Operation", r.Operation, "bladeId", b.Id, "applianceId", b.ApplianceId)

	// var status string

	if r.Operation == "assign" {

		request := backend.AssignMemoryRequest{
			PortId:   r.PortId,
			MemoryId: r.MemoryId,
		}

		settings := backend.ConfigurationSettings{}
		_, err := b.backendOps.AssignMemory(ctx, &settings, &request)
		if err != nil {
			newErr := fmt.Errorf("assign memory (backend) failure: request [%v]: %w", r, err)
			logger.Error(newErr, "failure: assign memory")
			return nil, &common.RequestError{StatusCode: common.StatusAssignMemoryFailure, Err: newErr}
		}
	} else {
		requestUnassign := backend.UnassignMemoryRequest{
			PortId:   r.PortId,
			MemoryId: r.MemoryId,
		}

		settings := backend.ConfigurationSettings{}
		_, err := b.backendOps.UnassignMemory(ctx, &settings, &requestUnassign)
		if err != nil {
			newErr := fmt.Errorf("unassign memory (backend) failure: request [%v]: %w", r, err)
			logger.Error(newErr, "failure: unassign memory")
			return nil, &common.RequestError{StatusCode: common.StatusUnassignMemoryFailure, Err: newErr}
		}
	}

	//Invalidate all related object cache's
	port, _ := b.GetPortById(ctx, r.PortId)
	port.InvalidateCache()
	memory, _ := b.GetMemoryById(ctx, r.MemoryId)
	memory.InvalidateCache()
	// Invalidate the cooresponding resource caches
	var resourcesToUpdate []string
	if len(memory.resourceIds) != 0 {
		resourcesToUpdate = memory.resourceIds
	} else {
		resourcesToUpdate = b.GetAllResourceIds(ctx)
	}
	for _, resourceId := range resourcesToUpdate {
		resource, err := b.GetResourceById(ctx, resourceId)
		if err != nil {
			continue
		}
		resource.InvalidateCache()
	}

	memoryRegion := openapi.MemoryRegion{
		Id:                  r.MemoryId,
		Status:              "", //Unused		//status
		Type:                openapi.MEMORYTYPE_MEMORY_TYPE_REGION,
		SizeMiB:             -1, //Not Implemented
		Bandwidth:           -1, //Not Implemented
		Latency:             -1, //Not Implemented
		MemoryApplianceId:   b.ApplianceId,
		MemoryBladeId:       b.Id,
		MemoryAppliancePort: r.PortId,
		MappedHostId:        NOT_IMPLEMENTED,
		MappedHostPort:      NOT_IMPLEMENTED,
	}

	logger.V(2).Info("success: assign memory", "memoryId", memoryRegion.Id, "portId", memoryRegion.MemoryAppliancePort, "bladeId", memoryRegion.MemoryBladeId, "applianceId", memoryRegion.MemoryApplianceId)

	return &memoryRegion, nil
}

type RequestComposeMemory struct {
	PortId  string
	SizeMib int32
	Qos     openapi.Qos
}

// ComposeMemory: Create a new memory region of the requested size and, if a port is requested, assign it to a port.
// In ComposeMemory, there are two steps in backend, one is AllocateMemory and the other one is AssignMemory
// AllocateMemory creates memory region (memorychunk), AssignMemory connects the memory region to an port
func (b *Blade) ComposeMemory(ctx context.Context, r *RequestComposeMemory) (*openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ComposeMemory: ", "portId", r.PortId, "SizeMiB", r.SizeMib, "Qos", r.Qos, "bladeId", b.Id, "applianceId", b.ApplianceId)

	// Update all resource details
	for _, resource := range b.Resources {
		_, err := resource.GetDetails(ctx)
		if err != nil {
			newErr := fmt.Errorf("get details failure on blade [%s] resource [%s]: %w", b.Id, resource.Id, err)
			logger.Error(newErr, "failure: compose memory")
			return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
		}
	}
	resourceIds := b.findResourcesByQoS(r.SizeMib, r.Qos)
	if resourceIds == nil {
		newErr := fmt.Errorf("unable to find resources by qos during compose: appliance [%s] blade [%s] request [%v]", b.ApplianceId, b.Id, r)
		logger.Error(newErr, "failure: compose memory")
		return nil, &common.RequestError{StatusCode: common.StatusBladeGetMemoryResourceBlocksFailure, Err: newErr}
	}

	memoryRegion, err := b.ComposeMemoryByResource(ctx, r.PortId, resourceIds)
	if err != nil {
		if memoryRegion != nil {
			newErr := fmt.Errorf("compose memory by resource allocation success but port assignment failure: %w", err)
			logger.Error(newErr, "partial success: compose memory")
			logger.V(2).Info("partial success: compose memory: ", "memoryId", memoryRegion.Id, "SizeMiB", memoryRegion.SizeMiB, "bladeId", memoryRegion.MemoryBladeId, "applianceId", memoryRegion.MemoryApplianceId)
			return memoryRegion, &common.RequestError{StatusCode: common.StatusComposePartialSuccess, Err: newErr}
		} else {
			newErr := fmt.Errorf("compose memory by resource failure during compose memory: %w", err)
			logger.Error(newErr, "failure: compose memory")
			return nil, &common.RequestError{StatusCode: common.StatusComposeMemoryByResourceFailure, Err: newErr}
		}
	}

	logger.V(2).Info("success: compose memory", "memoryId", memoryRegion.Id, "portId", memoryRegion.MemoryAppliancePort, "SizeMiB", memoryRegion.SizeMiB, "bladeId", memoryRegion.MemoryBladeId, "applianceId", memoryRegion.MemoryApplianceId)

	return memoryRegion, nil
}

// ComposeMemoryByResource: Create a new memory region from the requested resource block ids and, if a port is requested, assign it to a port.
// In ComposeMemoryByResource, there are two steps in backend, one is AllocateMemory and the other one is AssignMemory
// AllocateMemory creates memory region (memorychunk), AssignMemory connects the memory region to a port
func (b *Blade) ComposeMemoryByResource(ctx context.Context, portId string, resourceIds []string) (*openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ComposeMemoryByResource: ", "portId", portId, "resourceIds", resourceIds, "bladeId", b.Id, "applianceId", b.ApplianceId)

	// Execute provision
	requestAlloc := backend.AllocateMemoryByResourceRequest{
		MemoryResoureIds: resourceIds,
	}

	settings := backend.ConfigurationSettings{}
	responseAlloc, errAlloc := b.backendOps.AllocateMemoryByResource(ctx, &settings, &requestAlloc)
	if errAlloc != nil {
		newErr := fmt.Errorf("allocate memory (backed) failure during compose by resource: appliance [%s] blade [%s] resourceIds [%s]: %w", b.ApplianceId, b.Id, resourceIds, errAlloc)
		logger.Error(newErr, "failure: compose memory by resource")
		return nil, &common.RequestError{StatusCode: common.StatusComposeMemoryByResourceFailure, Err: newErr}
	}

	// Add the new memory region to cfm-service
	memory, err := b.addMemoryById(ctx, responseAlloc.MemoryId)
	if err != nil {
		newErr := fmt.Errorf("add memory by id failed [%s] during compose by resource: appliance [%s] blade [%s] resourceIds [%s]: %w", responseAlloc.MemoryId, b.ApplianceId, b.Id, resourceIds, err)
		logger.Error(newErr, "failure: compose memory by resource")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	// Save there resource ids.  Used when memory is freed
	memory.resourceIds = append(memory.resourceIds, resourceIds...)

	// query memory region to obtain capacity
	memoryRegion, err := memory.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("get details failure for memory [%s] during compose by resource: appliance [%s] blade [%s] resourceIds [%s]: %w", responseAlloc.MemoryId, b.ApplianceId, b.Id, resourceIds, err)
		logger.Error(newErr, "failure: compose memory by resource")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	//Invalidate all related object cache's
	for _, resourceId := range resourceIds {
		resource, _ := b.GetResourceById(ctx, resourceId)
		resource.InvalidateCache()
	}

	// If no port given, allocate only
	if portId == "" {
		logger.V(2).Info("success: compose memory by resource", "memoryId", memoryRegion.Id, "portId", "no port", "resourceIds", resourceIds, "bladeId", memoryRegion.MemoryBladeId, "applianceId", memoryRegion.MemoryApplianceId)
		return &memoryRegion, nil
	}

	requestAssign := RequestAssignMemory{
		MemoryId:  memoryRegion.Id,
		PortId:    portId,
		Operation: "assign",
	}

	memoryAssign, errAssign := b.AssignMemory(ctx, &requestAssign)
	if errAssign != nil {
		newErr := fmt.Errorf("assign memory failure during compose memory by resource: requestAssign [%v]: %w", requestAssign, errAssign)
		logger.Error(newErr, "failure: compose memory by resource")
		return &memoryRegion, &common.RequestError{StatusCode: common.StatusComposePartialSuccess, Err: newErr}
	}

	// Save assigned port
	memoryRegion.MemoryAppliancePort = memoryAssign.MemoryAppliancePort
	memoryRegion.Status = "" //Unused
	memoryRegion.MappedHostId = NOT_IMPLEMENTED
	memoryRegion.MappedHostPort = NOT_IMPLEMENTED

	logger.V(2).Info("success: compose memory by resource", "memoryId", memoryRegion.Id, "portId", memoryRegion.MemoryAppliancePort, "resourceIds", resourceIds, "bladeId", memoryRegion.MemoryBladeId, "applianceId", memoryRegion.MemoryApplianceId)

	return &memoryRegion, nil
}

func (b *Blade) FreeMemoryById(ctx context.Context, memoryId string) (*openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> FreeMemory: ", "memoryId", memoryId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	settings := backend.ConfigurationSettings{}

	// Check if the memory region exists
	getMemoryRequest := backend.GetMemoryByIdRequest{
		MemoryId: memoryId,
	}

	getMemoryResponse, getMemoryErr := b.backendOps.GetMemoryById(ctx, &settings, &getMemoryRequest)
	if getMemoryErr != nil {
		newErr := fmt.Errorf("get memory by id (backend) failure on appliance [%s] blade [%s] memory [%s]: %w", b.ApplianceId, b.Id, memoryId, getMemoryErr)
		logger.Error(newErr, "failure: free memory by id")
		return nil, &common.RequestError{StatusCode: common.StatusBladeGetMemoryByIdFailure, Err: newErr}
	}

	// If present, remove the assigned port before deallocating the memory region
	if getMemoryResponse.MemoryRegion.PortId != "" {
		unassignRequest := backend.UnassignMemoryRequest{
			MemoryId: memoryId,
		}

		_, unassignErr := b.backendOps.UnassignMemory(ctx, &settings, &unassignRequest)
		if unassignErr != nil {
			newErr := fmt.Errorf("unassign memory (backend) failure on appliance [%s] blade [%s] memory [%s]: %w", b.ApplianceId, b.Id, memoryId, unassignErr)
			logger.Error(newErr, "failure: free memory by id")
			return nil, &common.RequestError{StatusCode: common.StatusUnassignMemoryFailure, Err: newErr}
		}

		//Invalidate all related object cache's
		port, _ := b.GetPortById(ctx, getMemoryResponse.MemoryRegion.PortId)
		port.InvalidateCache()
	}

	req := backend.FreeMemoryRequest{
		MemoryId: memoryId,
	}

	response, err := b.backendOps.FreeMemoryById(ctx, &settings, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("free memory by id (backend) failure on appliance [%s] blade [%s] memory [%s]: %w", b.ApplianceId, b.Id, memoryId, err)
		logger.Error(newErr, "failure: free memory by id")
		return nil, &common.RequestError{StatusCode: common.StatusBladeFreeMemoryFailure, Err: newErr}
	}

	// Invalidate the cooresponding resource caches
	var resourcesToUpdate []string
	memory, _ := b.GetMemoryById(ctx, memoryId)
	if len(memory.resourceIds) != 0 {
		resourcesToUpdate = memory.resourceIds
	} else {
		resourcesToUpdate = b.GetAllResourceIds(ctx)
	}
	for _, resourceId := range resourcesToUpdate {
		resource, err := b.GetResourceById(ctx, resourceId)
		if err != nil {
			continue
		}
		resource.InvalidateCache()
	}

	memoryRegion := openapi.MemoryRegion{
		Id:                  getMemoryResponse.MemoryRegion.MemoryId,
		Status:              "", //Unused
		Type:                openapi.MemoryType(getMemoryResponse.MemoryRegion.Type),
		SizeMiB:             getMemoryResponse.MemoryRegion.SizeMiB,
		Bandwidth:           getMemoryResponse.MemoryRegion.Bandwidth,
		Latency:             getMemoryResponse.MemoryRegion.Latency,
		MemoryApplianceId:   b.ApplianceId,
		MemoryBladeId:       b.Id,
		MemoryAppliancePort: getMemoryResponse.MemoryRegion.PortId,
		MappedHostId:        NOT_IMPLEMENTED,
		MappedHostPort:      NOT_IMPLEMENTED,
	}

	// delete memory object from blade
	delete(b.Memory, memoryRegion.Id)

	logger.V(2).Info("success: free memory", "memoryId", memoryRegion.Id, "bladeId", memoryRegion.MemoryBladeId, "applianceId", memoryRegion.MemoryApplianceId)

	return &memoryRegion, nil
}

func (b *Blade) GetAllMemoryIds(ctx context.Context) []string {
	var ids []string

	for id := range b.GetMemory(ctx) {
		ids = append(ids, id)
	}

	return ids
}

func (b *Blade) GetAllPortIds(ctx context.Context) []string {
	var ids []string

	for id := range b.GetPorts(ctx) {
		ids = append(ids, id)
	}

	return ids
}

func (b *Blade) GetAllResourceIds(ctx context.Context) []string {
	var ids []string

	for id := range b.GetResources(ctx) {
		ids = append(ids, id)
	}

	return ids
}

func (b *Blade) GetMemoryById(ctx context.Context, memoryId string) (*BladeMemory, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryById: ", "memoryId", memoryId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return nil, nil
	}

	memory, ok := b.Memory[memoryId]
	if !ok {
		newErr := fmt.Errorf("memory [%s] not found on appliance [%s] blade [%s] ", memoryId, b.ApplianceId, b.Id)
		logger.Error(newErr, "failure: get memory by id")
		return nil, &common.RequestError{StatusCode: common.StatusMemoryIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get memory by id", "memoryId", memoryId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return memory, nil
}

func (b *Blade) GetMemory(ctx context.Context) map[string]*BladeMemory {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemory: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return make(map[string]*BladeMemory)
	}

	memory := b.Memory

	logger.V(2).Info("success: get memory", "count", len(memory), "bladeId", b.Id, "applianceId", b.ApplianceId)

	return memory
}

// GetMemoryBackend - Returns slice of hardware memory id's
func (b *Blade) GetMemoryBackend(ctx context.Context) ([]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryBackend: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return make([]string, 0), nil
	}

	req := backend.GetMemoryRequest{}
	response, err := b.backendOps.GetMemory(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("get memory (backend) [%s] failure on blade [%s]: %w", b.backendOps.GetBackendInfo(ctx).BackendName, b.Id, err)
		logger.Error(newErr, "failure: get memory(backend)")
		return nil, &common.RequestError{StatusCode: common.StatusBladeGetMemoryFailure, Err: newErr}
	}

	logger.V(2).Info("success: get memory(backend)", "memoryIds", response.MemoryIds, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return response.MemoryIds, nil
}

func (b *Blade) GetNetIp() string {
	return b.Socket.IpAddress
}

func (b *Blade) GetNetPort() uint16 {
	return b.Socket.Port
}

func (b *Blade) GetPortById(ctx context.Context, portId string) (*CxlBladePort, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetPortById: ", "portId", portId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return nil, nil
	}

	port, ok := b.Ports[portId]
	if !ok {
		newErr := fmt.Errorf("port [%s] not found on appliance [%s] blade [%s] ", portId, b.ApplianceId, b.Id)
		logger.Error(newErr, "failure: get port by id")
		return nil, &common.RequestError{StatusCode: common.StatusPortIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get port by id(blade) (cache)", "portId", portId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return port, nil
}

func (b *Blade) GetPorts(ctx context.Context) map[string]*CxlBladePort {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetPorts: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return make(map[string]*CxlBladePort)
	}

	ports := b.Ports

	logger.V(2).Info("success: get ports(blade) (cache)", "count", len(ports), "bladeId", b.Id, "applianceId", b.ApplianceId)

	return ports
}

// GetPortsBackend - Returns slice of hardware port id's
func (b *Blade) GetPortsBackend(ctx context.Context) ([]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetPortsBackend: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return make([]string, 0), nil
	}

	req := backend.GetPortsRequest{}
	response, err := b.backendOps.GetPorts(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("get ports (backend) [%s] failure on blade [%s]: %w", b.backendOps.GetBackendInfo(ctx).BackendName, b.Id, err)
		logger.Error(newErr, "failure: get ports(backend)")
		return nil, &common.RequestError{StatusCode: common.StatusBladeGetPortsFailure, Err: newErr}
	}

	logger.V(2).Info("success: get ports(backend)", "portIds", response.PortIds, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return response.PortIds, nil
}

func (b *Blade) GetResourceById(ctx context.Context, resourceId string) (*BladeResource, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetResourceById: ", "resourceId", resourceId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return nil, nil
	}

	resource, ok := b.Resources[resourceId]
	if !ok {
		newErr := fmt.Errorf("resource [%s] not found on appliance [%s] blade [%s] ", resourceId, b.ApplianceId, b.Id)
		logger.Error(newErr, "failure: get resource by id")
		return nil, &common.RequestError{StatusCode: common.StatusResourceIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get resource by id", "resourceId", resourceId, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return resource, nil
}

func (b *Blade) GetResources(ctx context.Context) map[string]*BladeResource {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetResources: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return make(map[string]*BladeResource)
	}

	resources := b.Resources

	logger.V(2).Info("success: get resources(cache)", "count", len(resources), "bladeId", b.Id, "applianceId", b.ApplianceId)

	return resources
}

// GetResourcesBackend - Returns slice of hardware resource id's
func (b *Blade) GetResourcesBackend(ctx context.Context) ([]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetResourcesBackend: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	if !b.IsOnline(ctx) {
		// If blade offline, not an error.  Just no information to return.
		return make([]string, 0), nil
	}

	req := backend.MemoryResourceBlocksRequest{}
	response, err := b.backendOps.GetMemoryResourceBlocks(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("get resources (backend) [%s] failure on blade [%s]: %w", b.backendOps.GetBackendInfo(ctx).BackendName, b.Id, err)
		logger.Error(newErr, "failure: get resources(backend)")
		return nil, &common.RequestError{StatusCode: common.StatusBladeGetMemoryResourceBlocksFailure, Err: newErr}
	}

	logger.V(2).Info("success: get resources(backend)", "resourceIds", response.MemoryResources, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return response.MemoryResources, nil
}

func (b *Blade) GetResourceTotals(ctx context.Context) (*ResponseResourceTotals, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetResourceTotals: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	var totalAvail, totalAlloc int32

	response := ResponseResourceTotals{
		TotalMemoryAvailableMiB: 0,
		TotalMemoryAllocatedMiB: 0,
	}

	for _, resource := range b.GetResources(ctx) {
		totals, err := resource.GetTotals(ctx)
		if err != nil || totals == nil {
			newErr := fmt.Errorf("failed to get resource totals: appliance [%s] blade [%s] resource [%s]: %w", b.ApplianceId, b.Id, resource.Id, err)
			logger.Error(newErr, "failure: get resource totals: blade")
			return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
		}

		totalAvail += totals.TotalMemoryAvailableMiB
		totalAlloc += totals.TotalMemoryAllocatedMiB
	}

	response.TotalMemoryAvailableMiB = totalAvail
	response.TotalMemoryAllocatedMiB = totalAlloc

	logger.V(2).Info("success: get resource totals", "totals", response, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return &response, nil
}

func (b *Blade) InvalidateCache() {
	for _, m := range b.Memory {
		m.InvalidateCache()
	}

	for _, p := range b.Ports {
		p.InvalidateCache()
	}

	for _, r := range b.Resources {
		r.InvalidateCache()
	}
}

func (b *Blade) IsOnline(ctx context.Context) bool {
	return b.Status == common.ONLINE
}

// UpdateConnectionStatusBackend - Query the blade for backend root and sesssion status and then update the manager's blade status accordingly.
func (b *Blade) UpdateConnectionStatusBackend(ctx context.Context) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> UpdateConnectionStatusBackend: ", "bladeId", b.Id)

	status := b.backendOps.GetBackendStatus(ctx)
	if status.FoundRootService {
		if status.FoundSession {
			b.Status = common.ONLINE
		} else {
			b.Status = common.FOUND
		}
	} else {
		b.Status = common.OFFLINE
	}

	// Update datastore status
	applianceDatum, _ := datastore.DStore().GetDataStore().GetApplianceDatumById(b.ApplianceId)
	bladeDatum, _ := applianceDatum.GetBladeDatumById(ctx, b.Id)
	bladeDatum.SetConnectionStatus(&b.Status)
	datastore.DStore().Store()

	logger.V(2).Info("success: update blade status(backend)", "status", b.Status, "bladeId", b.Id)
}

/////////////////////////////////////
//////// Private Functions //////////
/////////////////////////////////////

func (b *Blade) addMemoryById(ctx context.Context, memoryId string) (*BladeMemory, error) {
	logger := klog.FromContext(ctx)

	_, exists := b.Memory[memoryId]
	if exists {
		newErr := fmt.Errorf("appliance [%s] blade [%s] already contains memory with id [%s]", b.ApplianceId, b.Id, memoryId)
		return nil, &common.RequestError{StatusCode: common.StatusMemoryIdDuplicate, Err: newErr}
	}

	// Create the new object
	memory := NewBladeMemoryById(ctx, b.ApplianceId, b.Id, memoryId, b.backendOps)

	// Initialize new object
	err := memory.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("memory [%s] object init failed for appliance [%s] blade [%s]: %w", memory.Id, b.ApplianceId, b.Id, err)
		logger.Error(newErr, "failure: add memory by id")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	// Add the new object
	b.Memory[memory.Id] = memory

	return memory, nil
}

func (b *Blade) addPortById(ctx context.Context, portId string) (*CxlBladePort, error) {
	logger := klog.FromContext(ctx)

	_, exists := b.Ports[portId]
	if exists {
		newErr := fmt.Errorf("appliance [%s] blade [%s] already contains port with id [%s]", b.ApplianceId, b.Id, portId)
		return nil, &common.RequestError{StatusCode: common.StatusPortIdDuplicate, Err: newErr}
	}

	// Create the new object
	port := NewCxlBladePortById(ctx, b.ApplianceId, b.Id, portId, b.backendOps)

	// Initialize new object
	err := port.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("port [%s] object init failed for appliance [%s] blade [%s]: %w", port.Id, b.ApplianceId, b.Id, err)
		logger.Error(newErr, "failure: add port by id")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	// Add the new object
	b.Ports[port.Id] = port

	return port, nil
}

func (b *Blade) addResourceById(ctx context.Context, resourceId string) (*BladeResource, error) {
	logger := klog.FromContext(ctx)

	_, exists := b.Resources[resourceId]
	if exists {
		return nil, fmt.Errorf("appliance [%s] blade [%s] already contains resource with id [%s]", b.ApplianceId, b.Id, resourceId)
	}

	// Create the new object
	resource, err := NewBladeResourceById(ctx, b.ApplianceId, b.Id, resourceId, b.backendOps)
	if err != nil {
		newErr := fmt.Errorf("resource [%s] object creation failed for appliance [%s] blade [%s]: %w", resourceId, b.ApplianceId, b.Id, err)
		logger.Error(newErr, "failure: add resource by id")
		return nil, newErr
	}

	// Initialize new object
	err = resource.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("resource [%s] object init failed for appliance [%s] blade [%s]: %w", resource.Id, b.ApplianceId, b.Id, err)
		logger.Error(newErr, "failure: add resource by id")
		return nil, newErr
	}

	// Add the new object
	b.Resources[resource.Id] = resource

	logger.V(2).Info("success: add resource by id", "resourceId", resource.Id, "bladeId", b.Id, "applianceId", b.ApplianceId)

	return resource, nil
}

// countMaxConsecutiveTrue: Count the longest consecutive length of the true elements from the input boolean array
func (b *Blade) countMaxConsecutiveTrue(array []bool) int {
	maxLen := 0
	currentLen := 0
	for _, element := range array {
		if element {
			currentLen++
		} else {
			currentLen = 0
		}
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

// countTrue: Count the total of the true elements from the input boolean array
func (b *Blade) countTrue(array []bool) int {
	sum := 0
	for _, element := range array {
		if element {
			sum++
		}
	}

	return sum
}

/*  Selecting the memory resource blocks for Size and QoS request, using findResourcesByQoS()

User inputs the desired memory size and Quality of Service ( QoS ) for the desired memory chunk. QoS
represents the bandwidth of the memory chunk could provide. In memory appliance design, the
QoS is directly associated with how many memory channels are interleaved under the memory chunk.
Each memory channel is split into memory resource blocks, each of the same size and
ordered consecutively in the physical memory space indicated by the resource id
The requirements when creating a memory chunks are:
	* QoS indicates the memory channels used for the memory chunk
	* The number of memory resource blocks used from each channel are equal
	* The memory resource blocks used within the same channel are consecutive

The algorith:
	1. Each resource is identified with its channel number, position in the channel and the state ( free or NOT free )
	2. Calculate how many resources are required from a channel
	3. For each channel, find an available consecutive resource id list for the requested length,
	   the number of remaining unused resource on the channel,
	   the largest number of remaining consecutive unused resources on the channel
		3.1 Check consecutive resource size for requested resources length
		3.2 Iterate each potential resource start position
			3.2.1 Count the longest remain consecutive resource counts
			3.2.2 Update the candidate resource id list for the largest remaining consecutive resource counts
	4. Sort the channel candidate list by largest remaining consecutive resources length. Reordering will allow us to
		4.1 Check if QoS can be satisfied
		4.2 Pick the channels to provide the largest numberof remaining consecutive unused resources for future composition.
	6. Return the combined the resource id list if the size and QoS requirements were met
*/

type candidateData struct {
	resourceIds                []string
	unusedResources            int
	unusedResourcesConsecutive int
}

/*
FindResourceBlockByQoS: Returns the resource ids to meet the composition requirement
*/
func (b *Blade) findResourcesByQoS(sizeMib int32, qos openapi.Qos) []string {
	var resourceIds []string

	// Validate input
	targetQos := int32(qos) // r.QoS represents the minimum number of channels requested to create the memory region
	if targetQos > b.NumResourceChannels || targetQos == 0 {
		targetQos = b.NumResourceChannels // Can only use the max number of channels available
	}

	resourcesNeededPerChannel := int(math.Ceil((float64(sizeMib) / float64(b.ResourceSizeMib)) / float64(targetQos)))
	candidates := make([]candidateData, b.NumResourceChannels)
	for i := int32(0); i < b.NumResourceChannels; i++ {
		candidates[i] = *b.findResourceCandidatesByChannel(i, resourcesNeededPerChannel)
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].unusedResourcesConsecutive >= candidates[j].unusedResourcesConsecutive
	})

	// QoS defines the # of memory resource "channels" the requested memory must span (ie - the bandwidth)
	// The "candidate" array contains available memory resources, for a given QoS, for each channel.
	// (Note that sometimes a channel can have NO options)
	// So, if you have a QoS=2 and 4 memory resources channels, there are 2 possible QoS "groups' available in the candidate array to evaluate here.
	qosGroupIndexStepSize := targetQos
	for qosGroupStartIndex := int32(0); qosGroupStartIndex < b.NumResourceChannels; qosGroupStartIndex = qosGroupStartIndex + qosGroupIndexStepSize {
		start := qosGroupStartIndex
		end := start + qosGroupIndexStepSize

		noResourcesAvailable := false
		for i := start; i < end; i++ {
			if len(candidates[i].resourceIds) == 0 {
				noResourcesAvailable = true
				break
			}
		}

		if noResourcesAvailable {
			continue
		}

		for i := start; i < end; i++ {
			resourceIds = append(resourceIds, candidates[i].resourceIds[:]...)
		}

		break
	}

	return resourceIds
}

/*
findResourceBlockFromChannel: Find a list of consecutive resources to meet the required minimum
resource count (for the channel) while still leaving the largest possible number of consecutive unused
resources available.

	Returns:
		the list of the resource ids that meet selection criteria -- empty list if none selected
		the channel's total number of remaining unused resources -- 0 if none selected
		the channel's largest number of remaining consecutive unused resources -- 0 if none selected
*/
func (b *Blade) findResourceCandidatesByChannel(channelId int32, minResourceCnt int) *candidateData {
	var selectedResourceIds []string

	channelCompositionStatus := b.getCompositionStatusForChannel(channelId)

	unusedResources := b.countTrue(channelCompositionStatus)
	unusedResourcesConsecutive := b.countMaxConsecutiveTrue(channelCompositionStatus)
	if unusedResourcesConsecutive >= minResourceCnt {
		optimalUnusedConsecutive := 0
		optimalStartIdx := -1
		for startIdx := 0; startIdx < (len(channelCompositionStatus) - minResourceCnt + 1); startIdx++ {
			invalidPsn := false
			tmpStatus := make([]bool, len(channelCompositionStatus))
			copy(tmpStatus, channelCompositionStatus)
			for i := 0; i < minResourceCnt; i++ {
				if tmpStatus[startIdx+i] {
					tmpStatus[startIdx+i] = false
				} else {
					invalidPsn = true
					break
				}
			}
			if invalidPsn {
				continue
			}

			tmpUnusedConsecutive := b.countMaxConsecutiveTrue(tmpStatus)
			if tmpUnusedConsecutive >= optimalUnusedConsecutive {
				optimalUnusedConsecutive = tmpUnusedConsecutive
				optimalStartIdx = startIdx
			}
		}

		if optimalStartIdx != -1 {
			for _, resource := range b.Resources {
				if resource.GetChannelId() == channelId {
					if resource.GetChannelResourceIndex() >= int32(optimalStartIdx) && resource.GetChannelResourceIndex() < int32(optimalStartIdx+minResourceCnt) {
						selectedResourceIds = append(selectedResourceIds, resource.Id)
					}
				}
			}
			unusedResources -= minResourceCnt
			unusedResourcesConsecutive = optimalUnusedConsecutive
		}
	} else {
		unusedResources = 0
		unusedResourcesConsecutive = 0
	}

	result := candidateData{
		resourceIds:                selectedResourceIds,
		unusedResources:            unusedResources,
		unusedResourcesConsecutive: unusedResourcesConsecutive,
	}

	return &result
}

/*
getCompositionStatusForChannel: Convert the composition status for all resources on this channel into
a boolean array.  Index the array by resource position index. "true" indicates the resource is free for compose.
*/
func (b *Blade) getCompositionStatusForChannel(channelId int32) []bool {
	ChannelCnt := 0
	for _, resource := range b.Resources {
		if resource.GetChannelId() == channelId {
			ChannelCnt++
		}
	}

	compositionStatus := make([]bool, ChannelCnt)
	for _, resource := range b.Resources {
		if resource.GetChannelId() == channelId {
			if resource.GetCompositionState() == backend.RESOURCE_STATE_UNUSED {
				compositionStatus[resource.GetChannelResourceIndex()] = true
			} else {
				compositionStatus[resource.GetChannelResourceIndex()] = false
			}
		}
	}

	return compositionStatus
}

func (b *Blade) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	err := b.initPorts(ctx)
	if err != nil {
		newErr := fmt.Errorf("blade [%s] ports init failed for appliance [%s]: %w", b.Id, b.ApplianceId, err)
		logger.Error(newErr, "failure: init blade")
		return newErr
	}

	err = b.initResources(ctx)
	if err != nil {
		newErr := fmt.Errorf("blade [%s] resources init failed for appliance [%s]: %w", b.Id, b.ApplianceId, err)
		logger.Error(newErr, "failure: init blade")
		return newErr
	}

	err = b.initMemory(ctx)
	if err != nil {
		newErr := fmt.Errorf("blade [%s] memory init failed for appliance [%s]: %w", b.Id, b.ApplianceId, err)
		logger.Error(newErr, "failure: init blade")
		return newErr
	}

	logger.V(2).Info("success: init", "bladeId", b.Id, "applianceId", b.ApplianceId)

	return nil
}

func (b *Blade) initMemory(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> initMemory: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	memoryIds, err := b.GetMemoryBackend(ctx)
	if err != nil {
		newErr := fmt.Errorf("blade [%s] init failed during get memory: %w", b.Id, err)
		logger.Error(newErr, "failure: init memory: blade")
		return newErr
	}

	for _, memoryId := range memoryIds {
		_, err := b.addMemoryById(ctx, memoryId)
		if err != nil {
			newErr := fmt.Errorf("add memory by id failed for blade [%s]: %w", b.Id, err)
			logger.Error(newErr, "failure: init memory: blade")
			return newErr
		}
	}

	logger.V(2).Info("success: init memory", "bladeId", b.Id, "applianceId", b.ApplianceId)

	return nil
}

func (b *Blade) initPorts(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> initPorts: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	portIds, err := b.GetPortsBackend(ctx)
	if err != nil || portIds == nil {
		newErr := fmt.Errorf("blade [%s] init failed during get ports: %w", b.Id, err)
		logger.Error(newErr, "failure: init ports: blade")
		return newErr
	}

	for _, portId := range portIds {
		_, err := b.addPortById(ctx, portId)
		if err != nil {
			newErr := fmt.Errorf("add port by id failed for blade [%s]: %w", b.Id, err)
			logger.Error(newErr, "failure: init port: blade")
			return newErr
		}
	}

	logger.V(2).Info("success: init ports", "bladeId", b.Id, "applianceId", b.ApplianceId)

	return nil
}

func (b *Blade) initResources(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> initResources: ", "bladeId", b.Id, "applianceId", b.ApplianceId)

	resourceIds, err := b.GetResourcesBackend(ctx)
	if err != nil || resourceIds == nil {
		newErr := fmt.Errorf("blade [%s] init failed during get resources: %w", b.Id, err)
		logger.Error(newErr, "failure: init resources: blade")
		return newErr
	}

	for _, resourceId := range resourceIds {
		_, err := b.addResourceById(ctx, resourceId)
		if err != nil {
			newErr := fmt.Errorf("add resource by id failed for blade [%s]: %w", b.Id, err)
			logger.Error(newErr, "failure: init resource: blade")
			return newErr
		}
	}

	// Scan all resources to determine the max number of channels(ie:dimms) available
	channels := make(map[int32]bool)
	for _, r := range b.Resources {
		if _, ok := channels[r.GetChannelId()]; !ok {
			channels[r.GetChannelId()] = true
		}
		if b.ResourceSizeMib == 0 {
			b.ResourceSizeMib = r.GetCapacityMib()
		}
	}
	b.NumResourceChannels = int32(len(channels))

	logger.V(2).Info("success: init resources", "bladeId", b.Id, "applianceId", b.ApplianceId)

	return nil
}

//////////////////////////////////////////////////////////////////
//////////////////////////// Helpers /////////////////////////////
//////////////////////////////////////////////////////////////////

// GetCxlSnFromGCxlId - convert cached cxl global id to a cxl port serial number.
// Example: "cb-7b-6a-39-22-df-e1-00:0000" to "0xcb7b6a3922dfe100"
func GetCxlSnFromGCxlId(gCxlId string) string {
	var cxlSn string

	elements := strings.Split(strings.Split(gCxlId, ":")[0], "-")
	if len(elements) > 0 {
		cxlSn = fmt.Sprintf("0x%s", strings.Join(elements, ""))
	}

	return cxlSn
}

// GetGCxlIdFromCxlSn - convert cached cxl port serial number to a cxl global id
// Example: "0xcb7b6a3922dfe100" to "cb-7b-6a-39-22-df-e1-00:0000"
func GetGCxlIdFromCxlSn(cxlSn string) string {
	var gCxlId string

	trimmed := strings.TrimPrefix(cxlSn, "0x")
	for i, chr := range strings.Split(trimmed, "") {
		gCxlId += chr
		if i != 0 && i%2 == 1 {
			gCxlId += "-"
		}
	}
	gCxlId = strings.TrimRight(gCxlId, "-")
	gCxlId += ":0000"
	return gCxlId
}
