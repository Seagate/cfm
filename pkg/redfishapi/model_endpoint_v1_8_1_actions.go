/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// EndpointV181Actions - The available actions for this resource.
type EndpointV181Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertEndpointV181ActionsRequired checks if the required fields are not zero-ed
func AssertEndpointV181ActionsRequired(obj EndpointV181Actions) error {
	return nil
}

// AssertEndpointV181ActionsConstraints checks if the values respects the defined constraints
func AssertEndpointV181ActionsConstraints(obj EndpointV181Actions) error {
	return nil
}
