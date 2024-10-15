// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

//////////////////////////////////////////////////////////////////
///////////////////////// BladeMemory ////////////////////////////
//////////////////////////////////////////////////////////////////

type BladeMemory struct {
	Id          string
	Uri         string
	BladeId     string
	ApplianceId string
	// Status      string

	// Cached data
	cacheUpdated bool // Don't know what cache update mechanism looks like yet.  Start with this.
	details      openapi.MemoryRegion
	resourceIds  []string

	// Backend access data
	backendOps backend.BackendOperations
}

func NewBladeMemoryById(ctx context.Context, applianceId, bladeId, memoryId string, ops backend.BackendOperations) *BladeMemory {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewBladeMemoryById: ", "memoryId", memoryId, "bladeId", bladeId, "applianceId", applianceId, "backend", ops.GetBackendInfo(ctx).BackendName)

	m := BladeMemory{
		Id:          memoryId,
		Uri:         GetCfmUriBladeMemoryId(applianceId, bladeId, memoryId),
		BladeId:     bladeId,
		ApplianceId: applianceId,

		backendOps: ops,
	}

	logger.V(2).Info("success: new blade memory", "memoryId", m.Id, "bladeId", m.BladeId, "applianceId", m.ApplianceId)

	return &m
}

func (m *BladeMemory) GetDetails(ctx context.Context) (openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetDetails: ", "memoryId", m.Id, "bladeId", m.BladeId, "applianceId", m.ApplianceId)

	if !m.cacheUpdated {

		req := backend.GetMemoryByIdRequest{
			MemoryId: m.Id,
		}

		response, err := m.backendOps.GetMemoryById(ctx, &backend.ConfigurationSettings{}, &req)
		if err != nil || response == nil {
			if response != nil && response.Status == "Not Found" {
				newErr := fmt.Errorf("memory [%s] doesn't exist: appliance [%s] blade [%s]: %w", m.Id, m.ApplianceId, m.BladeId, err)
				logger.Error(newErr, "failure: get details")
				return openapi.MemoryRegion{}, &common.RequestError{StatusCode: common.StatusMemoryIdDoesNotExist, Err: newErr}
			}
			newErr := fmt.Errorf("failed to get memory by id (backend): appliance [%s] blade [%s] memory [%s]: %w", m.ApplianceId, m.BladeId, m.Id, err)
			logger.Error(newErr, "failure: get details")
			return openapi.MemoryRegion{}, &common.RequestError{StatusCode: common.StatusBladeGetMemoryByIdFailure, Err: newErr}
		}

		m.details = openapi.MemoryRegion{
			Id:                  response.MemoryRegion.MemoryId,
			Status:              "", //Unused
			Type:                openapi.MemoryType(response.MemoryRegion.Type),
			SizeMiB:             response.MemoryRegion.SizeMiB,
			Bandwidth:           -1, // Not implemented
			Latency:             -1, // Not implemented
			MemoryApplianceId:   m.ApplianceId,
			MemoryBladeId:       m.BladeId,
			MemoryAppliancePort: response.MemoryRegion.PortId,
			MappedHostId:        NOT_IMPLEMENTED,
			MappedHostPort:      NOT_IMPLEMENTED,
		}

		if response.MemoryRegion.Bandwidth != 0 {
			m.details.Bandwidth = response.MemoryRegion.Bandwidth
		}

		if response.MemoryRegion.Latency != 0 {
			m.details.Latency = response.MemoryRegion.Latency
		}

		m.ValidateCache()
	}

	logger.V(2).Info("success: get details", "memoryId", m.Id, "bladeId", m.BladeId, "applianceId", m.ApplianceId)

	return m.details, nil
}

func (m *BladeMemory) InvalidateCache() {
	m.cacheUpdated = false
}

func (m *BladeMemory) ValidateCache() {
	m.cacheUpdated = true
}

func (m *BladeMemory) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "memoryId", m.Id, "bladeId", m.BladeId, "applianceId", m.ApplianceId)

	m.InvalidateCache()

	_, err := m.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("memory [%s] init failed on blade [%s]: %w", m.Id, m.BladeId, err)
		logger.Error(newErr, "failure: init blade memory")
		return &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	logger.V(2).Info("success: init blade memory", "memoryId", m.Id, "bladeId", m.BladeId, "applianceId", m.ApplianceId)

	return nil
}

//////////////////////////////////////////////////////////////////
///////////////////////// HostMemory /////////////////////////////
//////////////////////////////////////////////////////////////////

type HostMemory struct {
	Id     string
	Uri    string
	HostId string
	// Status string

	// Cached data
	cacheUpdated bool // Don't know what cache update mechanism looks like yet.  Start with this.
	details      openapi.MemoryRegion

	// Backend access data
	backendOps backend.BackendOperations
}

func NewHostMemoryById(ctx context.Context, hostId, memoryId string, ops backend.BackendOperations) (*HostMemory, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewHostMemoryById: ", "memoryId", memoryId, "hostId", hostId, "backend", ops.GetBackendInfo(ctx).BackendName)

	m := HostMemory{
		Id:         memoryId,
		Uri:        GetCfmUriHostMemoryId(hostId, memoryId),
		HostId:     hostId,
		backendOps: ops,
	}

	logger.V(2).Info("success: new host memory", "memoryId", m.Id, "hostId", m.HostId)

	return &m, nil
}

func (m *HostMemory) GetDetails(ctx context.Context) (openapi.MemoryRegion, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetDetails: ", "memoryId", m.Id, "hostId", m.HostId)

	if !m.cacheUpdated {

		req := backend.GetMemoryByIdRequest{
			MemoryId: m.Id,
		}

		response, err := m.backendOps.GetMemoryById(ctx, &backend.ConfigurationSettings{}, &req)
		if err != nil || response == nil {
			if response != nil && response.Status == "Not Found" {
				newErr := fmt.Errorf("memory [%s] doesn't exist on host [%s]: %w", m.Id, m.HostId, err)
				logger.Error(newErr, "failure: get details")
				return openapi.MemoryRegion{}, &common.RequestError{StatusCode: common.StatusMemoryIdDoesNotExist, Err: newErr}
			}

			newErr := fmt.Errorf("failed to get memory by id (backend): host [%s] memory [%s]: %w", m.HostId, m.Id, err)
			logger.Error(newErr, "failure: get details")
			return openapi.MemoryRegion{}, &common.RequestError{StatusCode: common.StatusHostGetMemoryByIdFailure, Err: newErr}
		}

		m.details = openapi.MemoryRegion{
			Id:                  response.MemoryRegion.MemoryId,
			Status:              "", //Unused
			Type:                openapi.MemoryType(response.MemoryRegion.Type),
			SizeMiB:             response.MemoryRegion.SizeMiB,
			MemoryApplianceId:   NOT_IMPLEMENTED,
			MemoryBladeId:       NOT_IMPLEMENTED,
			MemoryAppliancePort: NOT_IMPLEMENTED,
			MappedHostId:        m.HostId,
			MappedHostPort:      NOT_IMPLEMENTED,
		}

		if response.MemoryRegion.Bandwidth != 0 {
			m.details.Bandwidth = response.MemoryRegion.Bandwidth
		}

		if response.MemoryRegion.Latency != 0 {
			m.details.Latency = response.MemoryRegion.Latency
		}

		m.ValidateCache()
	}

	logger.V(2).Info("success: get details", "memoryId", m.Id, "hostId", m.HostId)

	return m.details, nil
}

func (m *HostMemory) GetTotals(ctx context.Context) (*ResponseHostMemoryTotals, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetTotals: ", "memoryId", m.Id, "hostId", m.HostId)

	var local, remote int32

	details, err := m.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("failed to get details for host [%s] memory [%s]: %w", m.HostId, m.Id, err)
		logger.Error(newErr, "failure: get totals")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	if string(details.Type) == string(openapi.MEMORYTYPE_MEMORY_TYPE_LOCAL) {
		local = details.SizeMiB
	} else if string(details.Type) == string(openapi.MEMORYTYPE_MEMORY_TYPE_CXL) {
		remote = details.SizeMiB
	}

	totals := ResponseHostMemoryTotals{
		LocalMemoryMib:  local,
		RemoteMemoryMib: remote,
	}

	logger.V(4).Info("success: get totals", "memoryId", m.Id, "hostId", m.HostId, "local(MiB)", local, "remote(MiB)", remote)

	return &totals, nil
}

func (m *HostMemory) InvalidateCache() {
	m.cacheUpdated = false
}

func (m *HostMemory) ValidateCache() {
	m.cacheUpdated = true
}

func (m *HostMemory) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "memoryId", m.Id, "host", m.HostId)

	m.InvalidateCache()

	_, err := m.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("memory [%s] init failed on host [%s]: %w", m.Id, m.HostId, err)
		logger.Error(newErr, "failure: init host memory")
		return newErr
	}

	logger.V(2).Info("success: init host memory", "memoryId", m.Id, "host", m.HostId)

	return nil
}
