/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PortV1110Reset - This action resets this port.
type PortV1110Reset struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertPortV1110ResetRequired checks if the required fields are not zero-ed
func AssertPortV1110ResetRequired(obj PortV1110Reset) error {
	return nil
}

// AssertPortV1110ResetConstraints checks if the values respects the defined constraints
func AssertPortV1110ResetConstraints(obj PortV1110Reset) error {
	return nil
}
