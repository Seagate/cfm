/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// RedundancyV141Actions - The available actions for this resource.
type RedundancyV141Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertRedundancyV141ActionsRequired checks if the required fields are not zero-ed
func AssertRedundancyV141ActionsRequired(obj RedundancyV141Actions) error {
	return nil
}

// AssertRedundancyV141ActionsConstraints checks if the values respects the defined constraints
func AssertRedundancyV141ActionsConstraints(obj RedundancyV141Actions) error {
	return nil
}