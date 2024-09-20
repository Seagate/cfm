// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"
	"strings"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/common/datastore"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

const ID_PREFIX_HOST_DFLT string = "host"

type Host struct {
	Id            string
	Uri           string
	Status        common.ConnectionStatus
	Socket        SocketDetails
	Ports         map[string]*CxlHostPort
	MemoryDevices map[string]*HostMemoryDevice
	Memory        map[string]*HostMemory

	// Backend access data
	backendOps backend.BackendOperations
	creds      *openapi.Credentials // Used during resync
}

var HostMemoryDomain = map[string]openapi.MemoryType{
	"CXL":  openapi.MEMORYTYPE_MEMORY_TYPE_CXL,
	"DIMM": openapi.MEMORYTYPE_MEMORY_TYPE_LOCAL,
}

type RequestNewHost struct {
	HostId     string
	Ip         string
	Port       uint16
	Status     common.ConnectionStatus
	BackendOps backend.BackendOperations
	Creds      *openapi.Credentials
}

func NewHost(ctx context.Context, r *RequestNewHost) (*Host, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewHost: ", "request", r)

	h := Host{
		Id:            r.HostId,
		Uri:           GetCfmUriHostId(r.HostId),
		Socket:        *NewSocketDetails(r.Ip, r.Port),
		Ports:         make(map[string]*CxlHostPort),
		Status:        r.Status,
		MemoryDevices: make(map[string]*HostMemoryDevice),
		Memory:        make(map[string]*HostMemory),
		backendOps:    r.BackendOps,
		creds:         r.Creds,
	}

	err := h.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("init host [%s] failure: %w", h.Id, err)
		logger.Error(newErr, "failure: new host")
		return nil, newErr
	}

	logger.V(2).Info("success: new host", "hostId", h.Id)

	return &h, nil
}

func (h *Host) ComposeMemory(ctx context.Context, r *RequestComposeMemory) (*openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> ComposeMemory(Host): ", "request", r, "hostId", h.Id)

	blade, bladePort, err := h.getBladeAndPortByHostPortId(ctx, r.PortId) // Input portId is a hostPortId
	if err != nil {
		newErr := fmt.Errorf("connected blade port not found: host [%s] request [%v]", h.Id, r)
		logger.Error(newErr, "failure: compose memory(host)")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	req := RequestComposeMemory{
		PortId:  bladePort.Id,
		SizeMib: r.SizeMib,
		Qos:     r.Qos,
	}

	memory, err := blade.ComposeMemory(ctx, &req)
	if err != nil {
		newErr := fmt.Errorf("blade [%s] port [%s] compose memory failure: host [%s] request [%v]", blade.Id, bladePort.Id, h.Id, r)
		logger.Error(newErr, "failure: compose memory(host)")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	memory.MappedHostId = h.Id
	memory.MappedHostPort = r.PortId
	// ??? memory.MappedHostMemoryId = hostMemoryId  // Unknown.  Still need to power cycle cxl-host

	logger.V(2).Info("success: compose memory(host)", "memoryId", memory.Id, "hostId", h.Id, "hostPortId", r.PortId, "bladeId", blade.Id, "bladePortId", bladePort.Id)

	return memory, nil
}

func (h *Host) FreeMemoryById(ctx context.Context, hostMemoryId string) (*openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> FreeMemoryById(Host): ", "hostMemoryId", hostMemoryId, "hostId", h.Id)

	hostMemory, err := h.GetMemoryById(ctx, hostMemoryId)
	if err != nil {
		return nil, &common.RequestError{StatusCode: common.StatusHostFreeMemoryFailure, Err: err}
	}

	hostMemoryDetails, err := hostMemory.GetDetails(ctx)
	if err != nil {
		return nil, &common.RequestError{StatusCode: common.StatusHostFreeMemoryFailure, Err: err}
	}

	hostPortId := hostMemoryDetails.MappedHostPort

	blade, bladePort, err := h.getBladeAndPortByHostPortId(ctx, hostPortId)
	if err != nil {
		newErr := fmt.Errorf("couldn't find connected port on blade [%s] using host [%s] memory [%s] port [%s] ", blade.Id, h.Id, hostMemoryId, hostPortId)
		logger.Error(newErr, "failure: free memory by id(host)")
		return nil, &common.RequestError{StatusCode: common.StatusHostFreeMemoryFailure, Err: newErr}
	}

	var bladeMemoryId string
	for id, mem := range blade.GetMemory(ctx) {
		bladeMemoryDetails, err := mem.GetDetails(ctx)
		if err != nil {
			newErr := fmt.Errorf("couldn't retrieve memory objects on blade [%s] using host [%s] memory [%s] port [%s] ", blade.Id, h.Id, hostMemoryId, hostPortId)
			logger.Error(newErr, "failure: free memory by id(host)")
			return nil, &common.RequestError{StatusCode: common.StatusHostFreeMemoryFailure, Err: newErr}
		}
		if bladeMemoryDetails.MemoryAppliancePort == bladePort.Id {
			bladeMemoryId = id
			break
		}
	}
	if bladeMemoryId == "" {
		newErr := fmt.Errorf("couldn't find memory on blade [%s] using host [%s] memory [%s] port [%s] ", blade.Id, h.Id, hostMemoryId, hostPortId)
		logger.Error(newErr, "failure: free memory by id(host)")
		return nil, &common.RequestError{StatusCode: common.StatusHostFreeMemoryFailure, Err: newErr}
	}

	memory, err := blade.FreeMemoryById(ctx, bladeMemoryId)
	if err != nil {
		newErr := fmt.Errorf("blade [%s] memoryId [%s] free memory failure: host [%s] memoryId [%s]", blade.Id, bladeMemoryId, h.Id, hostMemoryId)
		logger.Error(newErr, "failure: free memory by id(host)")
		return nil, &common.RequestError{StatusCode: common.StatusHostFreeMemoryFailure, Err: newErr}
	}

	memory.MappedHostId = h.Id
	memory.MappedHostPort = hostPortId

	logger.V(2).Info("success: free memory by id(host)", "memoryId", memory.Id, "hostId", h.Id, "hostPortId", hostPortId, "bladeId", blade.Id, "bladePortId", bladePort.Id)

	return memory, nil
}

func (h *Host) GetAllMemoryIds() []string {
	var ids []string

	for id := range h.Memory {
		ids = append(ids, id)
	}

	return ids
}

func (h *Host) GetAllMemoryDeviceIds() []string {
	var ids []string

	for id := range h.MemoryDevices {
		ids = append(ids, id)
	}

	return ids
}

func (h *Host) GetAllPortIds() []string {
	var ids []string

	for id := range h.Ports {
		ids = append(ids, id)
	}

	return ids
}

func (h *Host) GetMemoryById(ctx context.Context, memoryId string) (*HostMemory, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryById: ", "memoryId", memoryId, "hostId", h.Id)

	memory, ok := h.Memory[memoryId]
	if !ok {
		newErr := fmt.Errorf("memory [%s] not found on host [%s]", memoryId, h.Id)
		logger.Error(newErr, "failure: get memory by id")
		return nil, &common.RequestError{StatusCode: common.StatusMemoryIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get memory by id", "memoryId", memoryId, "hostId", h.Id)

	return memory, nil
}

func (h *Host) GetMemory(ctx context.Context) map[string]*HostMemory {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemory: ", "hostId", h.Id)

	memory := h.Memory

	logger.V(2).Info("success: get memory", "count", len(memory), "hostId", h.Id)

	return memory
}

// GetMemoryBackend - Returns slice of memory id's
func (h *Host) GetMemoryBackend(ctx context.Context) ([]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryBackend: ", "hostId", h.Id)

	req := backend.GetMemoryRequest{}
	response, err := h.backendOps.GetMemory(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("get memory (backend) [%s] failure on host [%s]: %w", h.backendOps.GetBackendInfo(ctx).BackendName, h.Id, err)
		logger.Error(newErr, "failure: get memory(host) (backend)")
		return nil, &common.RequestError{StatusCode: common.StatusHostGetMemoryFailure, Err: newErr}
	}

	return response.MemoryIds, nil
}

func (h *Host) GetMemoryDeviceById(ctx context.Context, memdevId string) (*HostMemoryDevice, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryDeviceById: ", "memdevId", memdevId, "hostId", h.Id)

	memdev, ok := h.MemoryDevices[memdevId]
	if !ok {
		newErr := fmt.Errorf("memory device [%s] not found on host [%s]", memdevId, h.Id)
		logger.Error(newErr, "failure: get memory device by id")
		return nil, &common.RequestError{StatusCode: common.StatusMemoryDeviceIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get memory device by id", "memdevId", memdevId, "hostId", h.Id)

	return memdev, nil
}

func (h *Host) GetMemoryDevices(ctx context.Context) map[string]*HostMemoryDevice {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryDevices: ", "hostId", h.Id)

	memdevs := h.MemoryDevices

	logger.V(2).Info("success: get memory devices", "count", len(memdevs), "hostId", h.Id)

	return memdevs
}

func (h *Host) GetRedfishMemoryDomains(ctx context.Context) []string {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetRedfishMemoryDomains: ", "hostId", h.Id)

	memdomains := []string{}

	for memdom := range HostMemoryDomain {
		memdomains = append(memdomains, memdom)
	}

	logger.V(2).Info("success: get memory devices", "count", len(memdomains), "hostId", h.Id)

	return memdomains
}

func (h *Host) GetMemoryDomainAllMemoryIds(ctx context.Context, domain string) ([]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryDomainAllMemoryIds: ", "hostId", h.Id)

	var ids []string

	typeString, exist := HostMemoryDomain[domain]
	if exist {
		for id, mem := range h.Memory {
			details, err := mem.GetDetails(ctx)
			if err != nil {
				return nil, err
			}
			if details.Type == typeString {
				ids = append(ids, id)
			}
		}
		logger.V(2).Info("success: get memory devices", "count", len(ids), "hostId", h.Id)

		return ids, nil
	} else {
		newErr := fmt.Errorf("memory domain [%s] not found on host [%s]", domain, h.Id)
		logger.Error(newErr, "failure: get memory device by domain id")
		return nil, &common.RequestError{StatusCode: common.StatusMemoryDeviceIdDoesNotExist, Err: newErr}

	}
}

// GetMemoryDevicesBackend - Returns map of host physical device id (key) to 1 or more logical device ids (value)
func (h *Host) GetMemoryDevicesBackend(ctx context.Context) (map[string][]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryDevicesBackend: ", "hostId", h.Id)

	req := backend.GetMemoryDevicesRequest{}
	response, err := h.backendOps.GetMemoryDevices(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("get memory devices (backend) [%s] failure on host [%s]: %w", h.backendOps.GetBackendInfo(ctx).BackendName, h.Id, err)
		logger.Error(newErr, "failure: get memory devices(backend)")
		return nil, &common.RequestError{StatusCode: common.StatusGetMemoryDevicesFailure, Err: newErr}
	}

	logger.V(2).Info("success: get memory devices(backend)", "DeviceIdMap", response.DeviceIdMap)

	return response.DeviceIdMap, nil
}

func (h *Host) GetMemoryTotals(ctx context.Context) (*ResponseHostMemoryTotals, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetMemoryTotals: ", "hostId", h.Id)

	var local, remote int32

	response := ResponseHostMemoryTotals{
		LocalMemoryMib:  0,
		RemoteMemoryMib: 0,
	}

	if h.Status == common.OFFLINE {
		return &response, nil
	}

	for _, memory := range h.Memory {
		totals, err := memory.GetTotals(ctx)
		if err != nil || totals == nil {
			newErr := fmt.Errorf("failed to get memory totals: host [%s] memory [%s]: %w", h.Id, memory.Id, err)
			logger.Error(newErr, "failure: get memory totals: host")
			return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
		}

		local += totals.LocalMemoryMib
		remote += totals.RemoteMemoryMib
	}

	response.LocalMemoryMib = local
	response.RemoteMemoryMib = remote

	logger.V(2).Info("success: get memory totals", "hostId", h.Id)

	return &response, nil
}

func (h *Host) GetPortById(ctx context.Context, portId string) (*CxlHostPort, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetPortById: ", "portId", portId, "hostId", h.Id)

	port, ok := h.Ports[portId]
	if !ok {
		newErr := fmt.Errorf("port [%s] not found on host [%s]", portId, h.Id)
		logger.Error(newErr, "failure: get port by id")
		return nil, &common.RequestError{StatusCode: common.StatusPortIdDoesNotExist, Err: newErr}
	}

	logger.V(2).Info("success: get port by id", "portId", portId, "hostId", h.Id)

	return port, nil
}

func (h *Host) GetPorts(ctx context.Context) map[string]*CxlHostPort {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetPorts: ", "hostId", h.Id)

	ports := h.Ports

	logger.V(2).Info("success: get ports", "count", len(ports), "hostId", h.Id)

	return ports
}

// GetPortsBackend - Returns slice of backend port id's
func (h *Host) GetPortsBackend(ctx context.Context) ([]string, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetPortsBackend: ", "hostId", h.Id)

	req := backend.GetPortsRequest{}
	response, err := h.backendOps.GetHostPortPcieDevices(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		newErr := fmt.Errorf("get ports (backend) [%s] failure on host [%s]: %w", h.backendOps.GetBackendInfo(ctx).BackendName, h.Id, err)
		logger.Error(newErr, "failure: get ports(backend)")
		return nil, &common.RequestError{StatusCode: common.StatusHostGetPortsFailure, Err: newErr}
	}

	logger.V(2).Info("success: get ports(backend)", "portIds", response.PortIds)

	return response.PortIds, nil
}

func (h *Host) GetNetIp() string {
	return h.Socket.IpAddress
}

func (h *Host) GetNetPort() uint16 {
	return h.Socket.Port
}

func (h *Host) InvalidateCache() {
	for _, m := range h.Memory {
		m.InvalidateCache()
	}

	for _, p := range h.Ports {
		p.InvalidateCache()
	}

	for _, d := range h.MemoryDevices {
		d.InvalidateCache()
	}
}

// UpdateConnectionStatusBackend - Query the host root service to verify continued connection and update the object status accordingly.
func (h *Host) UpdateConnectionStatusBackend(ctx context.Context) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetConnectionsStatusBackend: ", "hostId", h.Id)

	req := backend.GetRootServiceRequest{}
	response, err := h.backendOps.GetRootService(ctx, &backend.ConfigurationSettings{}, &req)
	if err != nil || response == nil {
		h.Status = common.OFFLINE
	} else {
		h.Status = common.ONLINE
	}

	// Update datastore status
	r := datastore.UpdateHostStatusRequest{
		HostId: h.Id,
		Status: common.ConnectionStatus(h.Status),
	}
	datastore.DStore().GetDataStore().UpdateHostStatus(ctx, &r)
	datastore.DStore().Store()

	logger.V(2).Info("update host status(backend)", "status", h.Status, "hostId", h.Id)
}

type ResponseHostMemoryTotals struct {
	LocalMemoryMib  int32
	RemoteMemoryMib int32
}

/////////////////////////////////////
//////// Private Functions //////////
/////////////////////////////////////

func (h *Host) addMemoryById(ctx context.Context, memoryId string) (*HostMemory, error) {
	logger := klog.FromContext(ctx)

	_, exists := h.Memory[memoryId]
	if exists {
		return nil, fmt.Errorf("host [%s] already contains memory with id [%s]", h.Id, memoryId)
	}

	// Create the new object
	memory, err := NewHostMemoryById(ctx, h.Id, memoryId, h.backendOps)
	if err != nil {
		return nil, fmt.Errorf("memory object creation failed for host [%s]: %w", h.Id, err)
	}

	// Initialize new object
	err = memory.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("memory [%s] object init failed for host [%s]: %w", memory.Id, h.Id, err)
		logger.Error(newErr, "failure: add memory by id")
		return nil, newErr
	}

	// Add the new object
	h.Memory[memory.Id] = memory

	return memory, nil
}

func (h *Host) addMemoryDeviceById(ctx context.Context, physicalDeviceId, logicalDeviceId string) (*HostMemoryDevice, error) {
	logger := klog.FromContext(ctx)

	// Create the new object
	memdev, err := NewHostMemoryDeviceById(ctx, h.Id, physicalDeviceId, logicalDeviceId, h.backendOps)
	if err != nil {
		return nil, fmt.Errorf("memory device object creation failed for host [%s]: %w", h.Id, err)
	}

	// Initialize new object
	err = memdev.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("memory-device [%s] object init failed for host [%s]: %w", memdev.Id, h.Id, err)
		logger.Error(newErr, "failure: add memory-device by id")
		return nil, newErr
	}

	_, exists := h.MemoryDevices[memdev.Id]
	if exists {
		return nil, fmt.Errorf("host [%s] already contains memory device with id [%s]", h.Id, memdev.Id)
	}

	// Add the new object
	h.MemoryDevices[memdev.Id] = memdev

	return memdev, nil
}

func (h *Host) addPortById(ctx context.Context, backendPortId string) (*CxlHostPort, error) {
	logger := klog.FromContext(ctx)

	_, exists := h.Ports[backendPortId]
	if exists {
		return nil, fmt.Errorf("host [%s] already contains port with id [%s]", h.Id, backendPortId)
	}

	// Create the new object
	port, err := NewCxlHostPortById(ctx, h.Id, backendPortId, h.backendOps)
	if err != nil {
		return nil, fmt.Errorf("port object creation failed for host [%s]: %w", h.Id, err)
	}

	// Initialize new object
	err = port.init(ctx)
	if err != nil {
		newErr := fmt.Errorf("port [%s] object init failed for host [%s]: %w", port.Id, h.Id, err)
		logger.Error(newErr, "failure: add port by id")
		return nil, newErr
	}

	// Add the new object
	h.Ports[port.Id] = port

	return port, nil
}

// getBladeAndPortByHostPortId - If connected, retrieves the Blade and BladePort cfm-service objects for the given HostPort id
func (h *Host) getBladeAndPortByHostPortId(ctx context.Context, hostPortId string) (*Blade, *CxlBladePort, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> findBladeAndPortByHostPortId: ", "hostPortId", hostPortId, "hostId", h.Id)

	hostPort, err := h.GetPortById(ctx, hostPortId)
	if err != nil {
		return nil, nil, err
	}

	details, err := hostPort.GetDetails(ctx)
	if err != nil {
		return nil, nil, err
	}

	elements := strings.Split(details.LinkedPortUri, "/")
	l := len(elements)
	if l <= 6 {
		return nil, nil, fmt.Errorf("invalid linked port uri [%s]", details.LinkedPortUri)
	}

	applianceId := elements[l-5]
	bladeId := elements[l-3]
	bladePortId := elements[l-1]

	blade, err := deviceCache.GetBladeById(applianceId, bladeId)
	if err != nil {
		return nil, nil, err
	}

	bladePort, err := blade.GetPortById(ctx, bladePortId)
	if err != nil {
		return nil, nil, err
	}

	return blade, bladePort, nil

}

func (h *Host) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "hostId", h.Id)

	err := h.initPorts(ctx)
	if err != nil {
		newErr := fmt.Errorf("host [%s] ports init failed: %w", h.Id, err)
		logger.Error(newErr, "failure: init host")
		return newErr
	}

	err = h.initMemoryDevices(ctx)
	if err != nil {
		newErr := fmt.Errorf("host [%s] memory devices init failed: %w", h.Id, err)
		logger.Error(newErr, "failure: init host")
		return newErr
	}

	err = h.initMemory(ctx)
	if err != nil {
		newErr := fmt.Errorf("host [%s] memory init failed: %w", h.Id, err)
		logger.Error(newErr, "failure: init host")
		return newErr
	}

	logger.V(2).Info("success: init", "hostId", h.Id)

	return nil
}

func (h *Host) initMemory(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> initMemory: ", "hostId", h.Id)

	memoryIds, err := h.GetMemoryBackend(ctx)
	if err != nil || memoryIds == nil {
		newErr := fmt.Errorf("host [%s] init failed during get memory: %w", h.Id, err)
		logger.Error(newErr, "failure: init memory: host")
		return newErr
	}

	for _, memoryId := range memoryIds {
		_, err := h.addMemoryById(ctx, memoryId)
		if err != nil {
			newErr := fmt.Errorf("add memory by id failed for host [%s]: %w", h.Id, err)
			logger.Error(newErr, "failure: init memory: host")
			return newErr
		}
	}

	logger.V(2).Info("success: init memory", "hostId", h.Id)

	return nil
}

func (h *Host) initMemoryDevices(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> initMemoryDevices: ", "hostId", h.Id)

	deviceIdMap, err := h.GetMemoryDevicesBackend(ctx)
	if err != nil {
		newErr := fmt.Errorf("host [%s] init failed during get memory devices: %w", h.Id, err)
		logger.Error(newErr, "failure: init memory devices: host")
		return newErr
	}

	for physicalDeviceId, logicalDeviceIds := range deviceIdMap {
		for _, id := range logicalDeviceIds {
			_, err := h.addMemoryDeviceById(ctx, physicalDeviceId, id)
			if err != nil {
				newErr := fmt.Errorf("add memory device by id failed for host [%s] physDev [%s] logicalDev [%s]: %w", h.Id, physicalDeviceId, id, err)
				logger.Error(newErr, "failure: init memory device: host")
				return newErr
			}
		}
	}

	logger.V(2).Info("success: init memory devices", "hostId", h.Id)

	return nil
}

func (h *Host) initPorts(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> initPorts: ", "hostId", h.Id)

	backendPortIds, err := h.GetPortsBackend(ctx)
	if err != nil || backendPortIds == nil {
		newErr := fmt.Errorf("host [%s] init failed during get ports: %w", h.Id, err)
		logger.Error(newErr, "failure: init ports: host")
		return newErr
	}

	for _, portId := range backendPortIds {
		_, err := h.addPortById(ctx, portId)
		if err != nil {
			newErr := fmt.Errorf("add port by id failed for host [%s]: %w", h.Id, err)
			logger.Error(newErr, "failure: init port: host")
			return newErr
		}
	}

	logger.V(2).Info("success: init ports", "hostId", h.Id)

	return nil
}
