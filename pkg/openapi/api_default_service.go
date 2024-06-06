/*
 * Composer and Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() DefaultAPIServicer {
	return &DefaultAPIService{}
}

// AppliancesDeleteById -
func (s *DefaultAPIService) AppliancesDeleteById(ctx context.Context, applianceId string) (ImplResponse, error) {
	// TODO - update AppliancesDeleteById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Appliance{}) or use other options such as http.Ok ...
	// return Response(200, Appliance{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AppliancesDeleteById method not implemented")
}

// AppliancesGet -
func (s *DefaultAPIService) AppliancesGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update AppliancesGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AppliancesGet method not implemented")
}

// AppliancesGetById -
func (s *DefaultAPIService) AppliancesGetById(ctx context.Context, applianceId string) (ImplResponse, error) {
	// TODO - update AppliancesGetById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Appliance{}) or use other options such as http.Ok ...
	// return Response(200, Appliance{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AppliancesGetById method not implemented")
}

// AppliancesPost -
func (s *DefaultAPIService) AppliancesPost(ctx context.Context, credentials Credentials) (ImplResponse, error) {
	// TODO - update AppliancesPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Appliance{}) or use other options such as http.Ok ...
	// return Response(201, Appliance{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AppliancesPost method not implemented")
}

// BladesAssignMemoryById -
func (s *DefaultAPIService) BladesAssignMemoryById(ctx context.Context, applianceId string, bladeId string, memoryId string, assignMemoryRequest AssignMemoryRequest) (ImplResponse, error) {
	// TODO - update BladesAssignMemoryById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(200, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesAssignMemoryById method not implemented")
}

// BladesComposeMemory -
func (s *DefaultAPIService) BladesComposeMemory(ctx context.Context, applianceId string, bladeId string, composeMemoryRequest ComposeMemoryRequest) (ImplResponse, error) {
	// TODO - update BladesComposeMemory with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(201, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(409, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(409, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesComposeMemory method not implemented")
}

// BladesComposeMemoryByResource -
func (s *DefaultAPIService) BladesComposeMemoryByResource(ctx context.Context, applianceId string, bladeId string, composeMemoryByResourceRequest ComposeMemoryByResourceRequest) (ImplResponse, error) {
	// TODO - update BladesComposeMemoryByResource with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(201, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(409, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(409, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesComposeMemoryByResource method not implemented")
}

// BladesDeleteById -
func (s *DefaultAPIService) BladesDeleteById(ctx context.Context, applianceId string, bladeId string) (ImplResponse, error) {
	// TODO - update BladesDeleteById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Blade{}) or use other options such as http.Ok ...
	// return Response(200, Blade{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesDeleteById method not implemented")
}

// BladesFreeMemoryById -
func (s *DefaultAPIService) BladesFreeMemoryById(ctx context.Context, applianceId string, bladeId string, memoryId string) (ImplResponse, error) {
	// TODO - update BladesFreeMemoryById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(200, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesFreeMemoryById method not implemented")
}

// BladesGet -
func (s *DefaultAPIService) BladesGet(ctx context.Context, applianceId string) (ImplResponse, error) {
	// TODO - update BladesGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGet method not implemented")
}

// BladesGetById -
func (s *DefaultAPIService) BladesGetById(ctx context.Context, applianceId string, bladeId string) (ImplResponse, error) {
	// TODO - update BladesGetById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Blade{}) or use other options such as http.Ok ...
	// return Response(200, Blade{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetById method not implemented")
}

// BladesGetMemory -
func (s *DefaultAPIService) BladesGetMemory(ctx context.Context, applianceId string, bladeId string) (ImplResponse, error) {
	// TODO - update BladesGetMemory with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetMemory method not implemented")
}

// BladesGetMemoryById -
func (s *DefaultAPIService) BladesGetMemoryById(ctx context.Context, applianceId string, bladeId string, memoryId string) (ImplResponse, error) {
	// TODO - update BladesGetMemoryById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(200, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetMemoryById method not implemented")
}

// BladesGetPortById -
func (s *DefaultAPIService) BladesGetPortById(ctx context.Context, applianceId string, bladeId string, portId string) (ImplResponse, error) {
	// TODO - update BladesGetPortById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PortInformation{}) or use other options such as http.Ok ...
	// return Response(200, PortInformation{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetPortById method not implemented")
}

// BladesGetPorts -
func (s *DefaultAPIService) BladesGetPorts(ctx context.Context, applianceId string, bladeId string) (ImplResponse, error) {
	// TODO - update BladesGetPorts with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetPorts method not implemented")
}

// BladesGetResourceById -
func (s *DefaultAPIService) BladesGetResourceById(ctx context.Context, applianceId string, bladeId string, resourceId string) (ImplResponse, error) {
	// TODO - update BladesGetResourceById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryResourceBlock{}) or use other options such as http.Ok ...
	// return Response(200, MemoryResourceBlock{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetResourceById method not implemented")
}

// BladesGetResources -
func (s *DefaultAPIService) BladesGetResources(ctx context.Context, applianceId string, bladeId string) (ImplResponse, error) {
	// TODO - update BladesGetResources with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesGetResources method not implemented")
}

// BladesPost -
func (s *DefaultAPIService) BladesPost(ctx context.Context, applianceId string, credentials Credentials) (ImplResponse, error) {
	// TODO - update BladesPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Blade{}) or use other options such as http.Ok ...
	// return Response(201, Blade{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("BladesPost method not implemented")
}

// CfmGet -
func (s *DefaultAPIService) CfmGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update CfmGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	// return Response(200, string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CfmGet method not implemented")
}

// CfmV1Get -
func (s *DefaultAPIService) CfmV1Get(ctx context.Context) (ImplResponse, error) {
	// TODO - update CfmV1Get with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ServiceInformation{}) or use other options such as http.Ok ...
	// return Response(200, ServiceInformation{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CfmV1Get method not implemented")
}

// HostGetMemory -
func (s *DefaultAPIService) HostGetMemory(ctx context.Context, hostId string) (ImplResponse, error) {
	// TODO - update HostGetMemory with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostGetMemory method not implemented")
}

// HostsComposeMemory -
func (s *DefaultAPIService) HostsComposeMemory(ctx context.Context, hostId string, composeMemoryRequest ComposeMemoryRequest) (ImplResponse, error) {
	// TODO - update HostsComposeMemory with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(201, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(409, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(409, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsComposeMemory method not implemented")
}

// HostsDeleteById -
func (s *DefaultAPIService) HostsDeleteById(ctx context.Context, hostId string) (ImplResponse, error) {
	// TODO - update HostsDeleteById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Host{}) or use other options such as http.Ok ...
	// return Response(200, Host{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsDeleteById method not implemented")
}

// HostsFreeMemoryById -
func (s *DefaultAPIService) HostsFreeMemoryById(ctx context.Context, hostId string, memoryId string) (ImplResponse, error) {
	// TODO - update HostsFreeMemoryById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(200, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsFreeMemoryById method not implemented")
}

// HostsGet - Get CXL Host information.
func (s *DefaultAPIService) HostsGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update HostsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGet method not implemented")
}

// HostsGetById - Get information for a single CXL Host.
func (s *DefaultAPIService) HostsGetById(ctx context.Context, hostId string) (ImplResponse, error) {
	// TODO - update HostsGetById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Host{}) or use other options such as http.Ok ...
	// return Response(200, Host{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGetById method not implemented")
}

// HostsGetMemoryById -
func (s *DefaultAPIService) HostsGetMemoryById(ctx context.Context, hostId string, memoryId string) (ImplResponse, error) {
	// TODO - update HostsGetMemoryById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryRegion{}) or use other options such as http.Ok ...
	// return Response(200, MemoryRegion{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGetMemoryById method not implemented")
}

// HostsGetMemoryDeviceById -
func (s *DefaultAPIService) HostsGetMemoryDeviceById(ctx context.Context, hostId string, memoryDeviceId string) (ImplResponse, error) {
	// TODO - update HostsGetMemoryDeviceById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryDeviceInformation{}) or use other options such as http.Ok ...
	// return Response(200, MemoryDeviceInformation{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGetMemoryDeviceById method not implemented")
}

// HostsGetMemoryDevices -
func (s *DefaultAPIService) HostsGetMemoryDevices(ctx context.Context, hostId string) (ImplResponse, error) {
	// TODO - update HostsGetMemoryDevices with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGetMemoryDevices method not implemented")
}

// HostsGetPortById -
func (s *DefaultAPIService) HostsGetPortById(ctx context.Context, hostId string, portId string) (ImplResponse, error) {
	// TODO - update HostsGetPortById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PortInformation{}) or use other options such as http.Ok ...
	// return Response(200, PortInformation{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGetPortById method not implemented")
}

// HostsGetPorts -
func (s *DefaultAPIService) HostsGetPorts(ctx context.Context, hostId string) (ImplResponse, error) {
	// TODO - update HostsGetPorts with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Collection{}) or use other options such as http.Ok ...
	// return Response(200, Collection{}), nil

	// TODO: Uncomment the next line to return response Response(404, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(404, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsGetPorts method not implemented")
}

// HostsPost - Add a CXL host to be managed by CFM.
func (s *DefaultAPIService) HostsPost(ctx context.Context, credentials Credentials) (ImplResponse, error) {
	// TODO - update HostsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Host{}) or use other options such as http.Ok ...
	// return Response(201, Host{}), nil

	// TODO: Uncomment the next line to return response Response(400, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(400, StatusMessage{}), nil

	// TODO: Uncomment the next line to return response Response(500, StatusMessage{}) or use other options such as http.Ok ...
	// return Response(500, StatusMessage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HostsPost method not implemented")
}

// RootGet -
func (s *DefaultAPIService) RootGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update RootGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	// return Response(200, string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RootGet method not implemented")
}