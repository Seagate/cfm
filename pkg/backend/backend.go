// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package backend

import (
	"context"
	"errors"
)

type BackendOperations interface {
	CreateSession(context.Context, *ConfigurationSettings, *CreateSessionRequest) (*CreateSessionResponse, error)
	DeleteSession(context.Context, *ConfigurationSettings, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	GetMemoryResourceBlocks(context.Context, *ConfigurationSettings, *MemoryResourceBlocksRequest) (*MemoryResourceBlocksResponse, error)
	GetMemoryResourceBlockStatuses(ctx context.Context, settings *ConfigurationSettings, req *MemoryResourceBlockStatusesRequest) (*MemoryResourceBlockStatusesResponse, error)
	GetMemoryResourceBlockById(context.Context, *ConfigurationSettings, *MemoryResourceBlockByIdRequest) (*MemoryResourceBlockByIdResponse, error)
	GetPorts(context.Context, *ConfigurationSettings, *GetPortsRequest) (*GetPortsResponse, error)
	GetHostPortPcieDevices(ctx context.Context, settings *ConfigurationSettings, req *GetPortsRequest) (*GetPortsResponse, error)
	GetPortDetails(context.Context, *ConfigurationSettings, *GetPortDetailsRequest) (*GetPortDetailsResponse, error)
	GetHostPortSnById(ctx context.Context, settings *ConfigurationSettings, req *GetHostPortSnByIdRequest) (*GetHostPortSnByIdResponse, error)
	GetMemoryDevices(context.Context, *ConfigurationSettings, *GetMemoryDevicesRequest) (*GetMemoryDevicesResponse, error)
	GetMemoryDeviceDetails(context.Context, *ConfigurationSettings, *GetMemoryDeviceDetailsRequest) (*GetMemoryDeviceDetailsResponse, error)
	GetMemory(context.Context, *ConfigurationSettings, *GetMemoryRequest) (*GetMemoryResponse, error)
	AllocateMemory(context.Context, *ConfigurationSettings, *AllocateMemoryRequest) (*AllocateMemoryResponse, error)
	AllocateMemoryByResource(context.Context, *ConfigurationSettings, *AllocateMemoryByResourceRequest) (*AllocateMemoryByResourceResponse, error)
	FreeMemoryById(context.Context, *ConfigurationSettings, *FreeMemoryRequest) (*FreeMemoryResponse, error)
	AssignMemory(context.Context, *ConfigurationSettings, *AssignMemoryRequest) (*AssignMemoryResponse, error)
	UnassignMemory(context.Context, *ConfigurationSettings, *UnassignMemoryRequest) (*UnassignMemoryResponse, error)
	GetMemoryById(context.Context, *ConfigurationSettings, *GetMemoryByIdRequest) (*GetMemoryByIdResponse, error)
	GetBackendInfo(context.Context) *GetBackendInfoResponse
	GetBackendStatus(context.Context) *GetBackendStatusResponse
}

type commonService struct {
	version string
	session interface{}
}

type httpfishService struct {
	service commonService
	be      BackendOperations
}

// Supported interfaces
const (
	HttpfishServiceName string = "httpfish"
)

// NewBackendInterface : To return specific implementation of backend service interface
func NewBackendInterface(service string, parameters map[string]string) (BackendOperations, error) {
	localService, err := buildCommonService(parameters)
	if err == nil {
		if service == HttpfishServiceName {
			return &httpfishService{service: localService}, nil
		}
		return nil, errors.New("Invalid service: " + service)
	}
	return nil, err
}

// buildCommonService: Build a common service and initialize its version
func buildCommonService(config map[string]string) (commonService, error) {
	service := commonService{}
	if config != nil {
		service.version = config["version"]
	}
	return service, nil
}
