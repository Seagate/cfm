/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ComputerSystemV1220BootSourceOverrideEnabled string

// List of ComputerSystemV1220BootSourceOverrideEnabled
const (
	COMPUTERSYSTEMV1220BOOTSOURCEOVERRIDEENABLED_DISABLED   ComputerSystemV1220BootSourceOverrideEnabled = "Disabled"
	COMPUTERSYSTEMV1220BOOTSOURCEOVERRIDEENABLED_ONCE       ComputerSystemV1220BootSourceOverrideEnabled = "Once"
	COMPUTERSYSTEMV1220BOOTSOURCEOVERRIDEENABLED_CONTINUOUS ComputerSystemV1220BootSourceOverrideEnabled = "Continuous"
)

// AssertComputerSystemV1220BootSourceOverrideEnabledRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220BootSourceOverrideEnabledRequired(obj ComputerSystemV1220BootSourceOverrideEnabled) error {
	return nil
}

// AssertComputerSystemV1220BootSourceOverrideEnabledConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220BootSourceOverrideEnabledConstraints(obj ComputerSystemV1220BootSourceOverrideEnabled) error {
	return nil
}
