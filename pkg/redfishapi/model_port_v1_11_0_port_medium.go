/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110PortMedium string

// List of PortV1110PortMedium
const (
	PORTV1110PORTMEDIUM_ELECTRICAL PortV1110PortMedium = "Electrical"
	PORTV1110PORTMEDIUM_OPTICAL    PortV1110PortMedium = "Optical"
)

// AssertPortV1110PortMediumRequired checks if the required fields are not zero-ed
func AssertPortV1110PortMediumRequired(obj PortV1110PortMedium) error {
	return nil
}

// AssertPortV1110PortMediumConstraints checks if the values respects the defined constraints
func AssertPortV1110PortMediumConstraints(obj PortV1110PortMedium) error {
	return nil
}
