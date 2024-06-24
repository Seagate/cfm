/*
 * Composer and Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type StatusMessageStatus struct {

	// A unique status code value
	Code int32 `json:"code"`

	// A description of the status code
	Message string `json:"message"`
}

// AssertStatusMessageStatusRequired checks if the required fields are not zero-ed
func AssertStatusMessageStatusRequired(obj StatusMessageStatus) error {
	elements := map[string]interface{}{
		"code":    obj.Code,
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertStatusMessageStatusConstraints checks if the values respects the defined constraints
func AssertStatusMessageStatusConstraints(obj StatusMessageStatus) error {
	return nil
}
