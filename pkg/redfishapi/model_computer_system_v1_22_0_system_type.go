/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ComputerSystemV1220SystemType string

// List of ComputerSystemV1220SystemType
const (
	COMPUTERSYSTEMV1220SYSTEMTYPE_PHYSICAL               ComputerSystemV1220SystemType = "Physical"
	COMPUTERSYSTEMV1220SYSTEMTYPE_VIRTUAL                ComputerSystemV1220SystemType = "Virtual"
	COMPUTERSYSTEMV1220SYSTEMTYPE_OS                     ComputerSystemV1220SystemType = "OS"
	COMPUTERSYSTEMV1220SYSTEMTYPE_PHYSICALLY_PARTITIONED ComputerSystemV1220SystemType = "PhysicallyPartitioned"
	COMPUTERSYSTEMV1220SYSTEMTYPE_VIRTUALLY_PARTITIONED  ComputerSystemV1220SystemType = "VirtuallyPartitioned"
	COMPUTERSYSTEMV1220SYSTEMTYPE_COMPOSED               ComputerSystemV1220SystemType = "Composed"
	COMPUTERSYSTEMV1220SYSTEMTYPE_DPU                    ComputerSystemV1220SystemType = "DPU"
)

// AssertComputerSystemV1220SystemTypeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220SystemTypeRequired(obj ComputerSystemV1220SystemType) error {
	return nil
}

// AssertComputerSystemV1220SystemTypeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220SystemTypeConstraints(obj ComputerSystemV1220SystemType) error {
	return nil
}
