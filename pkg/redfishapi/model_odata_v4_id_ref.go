/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// OdataV4IdRef - A reference to a resource.
type OdataV4IdRef struct {

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id,omitempty"`
}

// AssertOdataV4IdRefRequired checks if the required fields are not zero-ed
func AssertOdataV4IdRefRequired(obj OdataV4IdRef) error {
	return nil
}

// AssertOdataV4IdRefConstraints checks if the values respects the defined constraints
func AssertOdataV4IdRefConstraints(obj OdataV4IdRef) error {
	return nil
}
