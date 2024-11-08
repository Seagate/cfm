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

// BladeResource represents a single memory resource block on the memory appliance blade.
type BladeResource struct {
	Id          string
	Uri         string
	BladeId     string
	ApplianceId string
	// Status      string

	// Cached data
	cacheUpdated bool // Don't know what cache update mechanism looks like yet.  Start with this.
	details      openapi.MemoryResourceBlock

	// Backend access data
	backendOps backend.BackendOperations
}

func NewBladeResourceById(ctx context.Context, applianceId, bladeId, resourceId string, ops backend.BackendOperations) (*BladeResource, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewBladeResourceById: ", "resourceId", resourceId, "bladeId", bladeId, "applianceId", applianceId, "backend", ops.GetBackendInfo(ctx).BackendName)

	r := BladeResource{
		Id:          resourceId,
		Uri:         GetCfmUriBladeResourceId(applianceId, bladeId, resourceId),
		BladeId:     bladeId,
		ApplianceId: applianceId,

		backendOps: ops,
	}

	logger.V(2).Info("success: new blade resource", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	return &r, nil
}

func (r *BladeResource) GetCompositionState() string {
	return r.details.CompositionStatus.CompositionState
}

func (r *BladeResource) GetChannelId() int32 {
	return r.details.ChannelId
}

func (r *BladeResource) GetChannelResourceIndex() int32 {
	return r.details.ChannelResourceIndex
}

func (r *BladeResource) GetDetails(ctx context.Context) (openapi.MemoryResourceBlock, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetDetails: ", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	if !r.cacheUpdated {

		req := backend.MemoryResourceBlockByIdRequest{
			ResourceId: r.Id,
		}

		response, err := r.backendOps.GetMemoryResourceBlockById(ctx, &backend.ConfigurationSettings{}, &req)
		if err != nil || response == nil {
			newErr := fmt.Errorf("failed to get resource by id (backend): appliance [%s] blade [%s] resource [%s]: %w", r.ApplianceId, r.BladeId, r.Id, err)
			logger.Error(newErr, "failure: get details")
			return openapi.MemoryResourceBlock{}, &common.RequestError{StatusCode: common.StatusBladeGetMemoryResourceBlockDetailsFailure, Err: newErr}
		}

		r.details = openapi.MemoryResourceBlock{
			Id:                   response.MemoryResourceBlock.Id,
			CapacityMiB:          response.MemoryResourceBlock.CapacityMiB,
			ChannelId:            response.MemoryResourceBlock.ChannelId,
			ChannelResourceIndex: response.MemoryResourceBlock.ChannelResourceIdx,
			DataWidthBits:        response.MemoryResourceBlock.DataWidthBits,
			MemoryType:           response.MemoryResourceBlock.MemoryType,
			MemoryDeviceType:     response.MemoryResourceBlock.MemoryDeviceType,
			Manufacturer:         response.MemoryResourceBlock.Manufacturer,
			OperatingSpeedMhz:    response.MemoryResourceBlock.OperatingSpeedMhz,
			PartNumber:           response.MemoryResourceBlock.PartNumber,
			SerialNumber:         response.MemoryResourceBlock.SerialNumber,
			RankCount:            response.MemoryResourceBlock.RankCount,
		}

		r.details.CompositionStatus = openapi.MemoryResourceBlockCompositionStatus{
			CompositionState:     response.MemoryResourceBlock.CompositionStatus.CompositionState.String(),
			MaxCompositions:      response.MemoryResourceBlock.CompositionStatus.MaxCompositions,
			NumberOfCompositions: response.MemoryResourceBlock.CompositionStatus.NumberOfCompositions,
		}

		r.ValidateCache()
	}

	logger.V(2).Info("success: get details", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	return r.details, nil
}

func (r *BladeResource) GetCapacityMib() int32 {
	return r.details.CapacityMiB
}

func (r *BladeResource) GetTotals(ctx context.Context) (*ResponseResourceTotals, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetTotals: ", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	var allocated, available int32

	details, err := r.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("failed to get details for appliance [%s] blade [%s] resource [%s]: %w", r.ApplianceId, r.BladeId, r.Id, err)
		logger.Error(newErr, "failure: get resource totals: resource")
		return nil, &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	if details.CompositionStatus.CompositionState == backend.RESOURCE_STATE_COMPOSED ||
		details.CompositionStatus.CompositionState == backend.RESOURCE_STATE_RESERVED {
		allocated = details.CapacityMiB
		logger.V(2).Info("GetResourceTotals", "state", details.CompositionStatus.CompositionState, "allocated", allocated)
	} else if details.CompositionStatus.CompositionState == backend.RESOURCE_STATE_UNUSED {
		available = details.CapacityMiB
		logger.V(2).Info("GetResourceTotals", "state", details.CompositionStatus.CompositionState, "available", available)
	} else {
		logger.V(2).Info("GetResourceTotals", "state", details.CompositionStatus.CompositionState, "unaccounted", details.CapacityMiB)
	}

	totals := ResponseResourceTotals{
		TotalMemoryAvailableMiB: available,
		TotalMemoryAllocatedMiB: allocated,
	}

	logger.V(2).Info("success: get resource totals", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	return &totals, nil
}

func (r *BladeResource) InvalidateCache() {
	r.cacheUpdated = false
}

// UpdateDetails - Update object with new backend information
func (r *BladeResource) UpdateDetails(status *backend.MemoryResourceBlockCompositionStatus) {
	r.details.CompositionStatus.CompositionState = status.CompositionState.String()
}

func (r *BladeResource) ValidateCache() {
	r.cacheUpdated = true
}

func (r *BladeResource) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	r.InvalidateCache()

	_, err := r.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("resource [%s] init failed on blade [%s]: %w", r.Id, r.BladeId, err)
		logger.Error(newErr, "failure: init blade resource")
		return &common.RequestError{StatusCode: err.(*common.RequestError).StatusCode, Err: newErr}
	}

	logger.V(2).Info("success: init blade resource", "resourceId", r.Id, "bladeId", r.BladeId, "applianceId", r.ApplianceId)

	return nil
}
