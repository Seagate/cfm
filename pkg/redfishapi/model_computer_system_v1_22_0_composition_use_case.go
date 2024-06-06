/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ComputerSystemV1220CompositionUseCase string

// List of ComputerSystemV1220CompositionUseCase
const (
	COMPUTERSYSTEMV1220COMPOSITIONUSECASE_RESOURCE_BLOCK_CAPABLE ComputerSystemV1220CompositionUseCase = "ResourceBlockCapable"
	COMPUTERSYSTEMV1220COMPOSITIONUSECASE_EXPANDABLE_SYSTEM      ComputerSystemV1220CompositionUseCase = "ExpandableSystem"
)

// AssertComputerSystemV1220CompositionUseCaseRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220CompositionUseCaseRequired(obj ComputerSystemV1220CompositionUseCase) error {
	return nil
}

// AssertComputerSystemV1220CompositionUseCaseConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220CompositionUseCaseConstraints(obj ComputerSystemV1220CompositionUseCase) error {
	return nil
}
