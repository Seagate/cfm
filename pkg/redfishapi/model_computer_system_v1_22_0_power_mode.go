/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ComputerSystemV1220PowerMode string

// List of ComputerSystemV1220PowerMode
const (
	COMPUTERSYSTEMV1220POWERMODE_MAXIMUM_PERFORMANCE          ComputerSystemV1220PowerMode = "MaximumPerformance"
	COMPUTERSYSTEMV1220POWERMODE_BALANCED_PERFORMANCE         ComputerSystemV1220PowerMode = "BalancedPerformance"
	COMPUTERSYSTEMV1220POWERMODE_POWER_SAVING                 ComputerSystemV1220PowerMode = "PowerSaving"
	COMPUTERSYSTEMV1220POWERMODE_STATIC                       ComputerSystemV1220PowerMode = "Static"
	COMPUTERSYSTEMV1220POWERMODE_OS_CONTROLLED                ComputerSystemV1220PowerMode = "OSControlled"
	COMPUTERSYSTEMV1220POWERMODE_OEM                          ComputerSystemV1220PowerMode = "OEM"
	COMPUTERSYSTEMV1220POWERMODE_EFFICIENCY_FAVOR_POWER       ComputerSystemV1220PowerMode = "EfficiencyFavorPower"
	COMPUTERSYSTEMV1220POWERMODE_EFFICIENCY_FAVOR_PERFORMANCE ComputerSystemV1220PowerMode = "EfficiencyFavorPerformance"
)

// AssertComputerSystemV1220PowerModeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220PowerModeRequired(obj ComputerSystemV1220PowerMode) error {
	return nil
}

// AssertComputerSystemV1220PowerModeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220PowerModeConstraints(obj ComputerSystemV1220PowerMode) error {
	return nil
}