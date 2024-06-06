/*
 * Composer and Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ServiceResource struct {

	// A full path to the resource with id as the last component
	Uri string `json:"uri"`

	// The service(s) available for the specified URI
	Methods string `json:"methods"`

	// The description of service(s) offered by the URI
	Description string `json:"description"`
}

// AssertServiceResourceRequired checks if the required fields are not zero-ed
func AssertServiceResourceRequired(obj ServiceResource) error {
	elements := map[string]interface{}{
		"uri":         obj.Uri,
		"methods":     obj.Methods,
		"description": obj.Description,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertServiceResourceConstraints checks if the values respects the defined constraints
func AssertServiceResourceConstraints(obj ServiceResource) error {
	return nil
}
