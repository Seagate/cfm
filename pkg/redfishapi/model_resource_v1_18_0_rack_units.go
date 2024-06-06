/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceV1180RackUnits : The type of rack unit in use.
type ResourceV1180RackUnits string

// List of ResourceV1180RackUnits
const (
	RESOURCEV1180RACKUNITS_OPEN_U  ResourceV1180RackUnits = "OpenU"
	RESOURCEV1180RACKUNITS_EIA_310 ResourceV1180RackUnits = "EIA_310"
)

// AssertResourceV1180RackUnitsRequired checks if the required fields are not zero-ed
func AssertResourceV1180RackUnitsRequired(obj ResourceV1180RackUnits) error {
	return nil
}

// AssertResourceV1180RackUnitsConstraints checks if the values respects the defined constraints
func AssertResourceV1180RackUnitsConstraints(obj ResourceV1180RackUnits) error {
	return nil
}