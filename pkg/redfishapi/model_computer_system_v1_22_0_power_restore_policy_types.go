/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ComputerSystemV1220PowerRestorePolicyTypes : The enumerations of PowerRestorePolicyTypes specify the choice of power state for the system when power is applied.
type ComputerSystemV1220PowerRestorePolicyTypes string

// List of ComputerSystemV1220PowerRestorePolicyTypes
const (
	COMPUTERSYSTEMV1220POWERRESTOREPOLICYTYPES_ALWAYS_ON  ComputerSystemV1220PowerRestorePolicyTypes = "AlwaysOn"
	COMPUTERSYSTEMV1220POWERRESTOREPOLICYTYPES_ALWAYS_OFF ComputerSystemV1220PowerRestorePolicyTypes = "AlwaysOff"
	COMPUTERSYSTEMV1220POWERRESTOREPOLICYTYPES_LAST_STATE ComputerSystemV1220PowerRestorePolicyTypes = "LastState"
)

// AssertComputerSystemV1220PowerRestorePolicyTypesRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220PowerRestorePolicyTypesRequired(obj ComputerSystemV1220PowerRestorePolicyTypes) error {
	return nil
}

// AssertComputerSystemV1220PowerRestorePolicyTypesConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220PowerRestorePolicyTypesConstraints(obj ComputerSystemV1220PowerRestorePolicyTypes) error {
	return nil
}