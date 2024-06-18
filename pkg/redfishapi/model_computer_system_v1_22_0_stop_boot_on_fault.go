/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ComputerSystemV1220StopBootOnFault string

// List of ComputerSystemV1220StopBootOnFault
const (
	COMPUTERSYSTEMV1220STOPBOOTONFAULT_NEVER     ComputerSystemV1220StopBootOnFault = "Never"
	COMPUTERSYSTEMV1220STOPBOOTONFAULT_ANY_FAULT ComputerSystemV1220StopBootOnFault = "AnyFault"
)

// AssertComputerSystemV1220StopBootOnFaultRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220StopBootOnFaultRequired(obj ComputerSystemV1220StopBootOnFault) error {
	return nil
}

// AssertComputerSystemV1220StopBootOnFaultConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220StopBootOnFaultConstraints(obj ComputerSystemV1220StopBootOnFault) error {
	return nil
}
