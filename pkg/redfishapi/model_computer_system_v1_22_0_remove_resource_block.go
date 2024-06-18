/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ComputerSystemV1220RemoveResourceBlock - This action removes a resource block from a system.
type ComputerSystemV1220RemoveResourceBlock struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertComputerSystemV1220RemoveResourceBlockRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220RemoveResourceBlockRequired(obj ComputerSystemV1220RemoveResourceBlock) error {
	return nil
}

// AssertComputerSystemV1220RemoveResourceBlockConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220RemoveResourceBlockConstraints(obj ComputerSystemV1220RemoveResourceBlock) error {
	return nil
}
