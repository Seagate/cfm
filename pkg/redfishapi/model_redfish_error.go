/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// RedfishError - The error payload from a Redfish service.
type RedfishError struct {
	Error RedfishErrorError `json:"error"`
}

// AssertRedfishErrorRequired checks if the required fields are not zero-ed
func AssertRedfishErrorRequired(obj RedfishError) error {
	elements := map[string]interface{}{
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRedfishErrorErrorRequired(obj.Error); err != nil {
		return err
	}
	return nil
}

// AssertRedfishErrorConstraints checks if the values respects the defined constraints
func AssertRedfishErrorConstraints(obj RedfishError) error {
	return nil
}
