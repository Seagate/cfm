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
///////////////////////// HostMemoryDevice ///////////////////////
//////////////////////////////////////////////////////////////////

const (
	HOST_MEMORY_DEVICE_ID_PREFIX string = "memdev"
)

type HostMemoryDevice struct {
	Id               string
	PhysicalDeviceId string
	LogicalDeviceId  string
	Uri              string
	HostId           string
	// Status           string

	// Cached data
	cacheUpdated bool // Don't know what cache update mechanism looks like yet.  Start with this.
	Details      openapi.MemoryDeviceInformation
	CxlSn        string

	// Backend access data
	backendOps backend.BackendOperations
}

func NewHostMemoryDeviceById(ctx context.Context, hostId, physicalDeviceId, logicalDeviceId string, ops backend.BackendOperations) (*HostMemoryDevice, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewHostMemoryDeviceById: ", "physicalDeviceId", physicalDeviceId, "logicalDeviceId", logicalDeviceId, "hostId", hostId, "backend", ops.GetBackendInfo(ctx).BackendName)

	memdevId := GenerateFrontendHostMemdevId(physicalDeviceId, logicalDeviceId)

	d := HostMemoryDevice{
		Id:               memdevId,
		PhysicalDeviceId: physicalDeviceId,
		LogicalDeviceId:  logicalDeviceId,
		Uri:              GetCfmUriHostMemoryDeviceId(hostId, memdevId),
		HostId:           hostId,

		backendOps: ops,
	}

	logger.V(2).Info("success: new host memory device", "memdevId", d.Id, "hostId", d.HostId)

	return &d, nil
}

func (d *HostMemoryDevice) GetDetails(ctx context.Context) (openapi.MemoryDeviceInformation, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetDetails: ", "memdevId", d.Id, "hostId", d.HostId)

	if !d.cacheUpdated {

		req := backend.GetMemoryDeviceDetailsRequest{
			PhysicalDeviceId: d.PhysicalDeviceId,
			LogicalDeviceId:  d.LogicalDeviceId,
		}

		response, err := d.backendOps.GetMemoryDeviceDetails(ctx, &backend.ConfigurationSettings{}, &req)
		if err != nil {
			if err.Error() == "Not Found" {
				newErr := fmt.Errorf("memory device not found (backend): host [%s] memory device [%s]: %w", d.HostId, d.Id, err)
				logger.Error(newErr, "failure: get host memory device details")
				return openapi.MemoryDeviceInformation{}, &common.RequestError{StatusCode: common.StatusMemoryDeviceIdDoesNotExist, Err: newErr}
			}

			newErr := fmt.Errorf("failed to get details (backend): host [%s] memory device [%s]: %w", d.HostId, d.Id, err)
			logger.Error(newErr, "failure: get host memory device details")
			return openapi.MemoryDeviceInformation{}, &common.RequestError{StatusCode: common.StatusGetMemoryDevicesDetailsFailure, Err: newErr}
		}

		d.Details.Id = d.Id
		d.Details.DeviceType = response.DeviceType
		d.Details.MemorySizeMiB = response.MemorySizeMiB

		d.Details.StatusState = response.Status

		if response.LinkStatus != nil {
			d.Details.LinkStatus.CurrentSpeedGTps = response.LinkStatus.CurrentSpeedGTps
			d.Details.LinkStatus.CurrentWidth = response.LinkStatus.CurrentWidth
			d.Details.LinkStatus.MaxSpeedGTps = response.LinkStatus.MaxSpeedGTps
			d.Details.LinkStatus.MaxWidth = response.LinkStatus.MaxWidth
		}

		d.CxlSn = response.SerialNumber

		d.ValidateCache()
	}

	return d.Details, nil
}

func (d *HostMemoryDevice) InvalidateCache() {
	d.cacheUpdated = false
}

func (d *HostMemoryDevice) ValidateCache() {
	// d.cacheUpdated = true
	d.cacheUpdated = false // Temporarily disable host cache usage
}

func (d *HostMemoryDevice) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "memdevId", d.Id, "host", d.HostId)

	d.InvalidateCache()

	_, err := d.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("memory-device [%s] init failed on host [%s]: %w", d.Id, d.HostId, err)
		logger.Error(newErr, "failure: init host memory-device")
		return newErr
	}

	logger.V(2).Info("success: init host memory-device", "memdevId", d.Id, "host", d.HostId)

	return nil
}

//////////////////////////////////////////////////////////////////
//////////////////////////// Helpers /////////////////////////////
//////////////////////////////////////////////////////////////////

func GenerateFrontendHostMemdevId(backendPortId, logicalDeviceId string) string {
	// Current host memory device id string format: memdevXX-XX.X
	return fmt.Sprintf("%s%s.%s", HOST_MEMORY_DEVICE_ID_PREFIX, backendPortId, logicalDeviceId)
}
