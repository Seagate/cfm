/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ResourceHealth string

// List of ResourceHealth
const (
	RESOURCEHEALTH_OK       ResourceHealth = "OK"
	RESOURCEHEALTH_WARNING  ResourceHealth = "Warning"
	RESOURCEHEALTH_CRITICAL ResourceHealth = "Critical"
)

// AssertResourceHealthRequired checks if the required fields are not zero-ed
func AssertResourceHealthRequired(obj ResourceHealth) error {
	return nil
}

// AssertResourceHealthConstraints checks if the values respects the defined constraints
func AssertResourceHealthConstraints(obj ResourceHealth) error {
	return nil
}
