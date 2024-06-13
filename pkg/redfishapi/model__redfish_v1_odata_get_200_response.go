/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// RedfishV1OdataGet200Response - The OData service document from a Redfish service.
type RedfishV1OdataGet200Response struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context"`

	// The list of services provided by the Redfish service.
	Value []RedfishV1OdataGet200ResponseValueInner `json:"value"`
}

// AssertRedfishV1OdataGet200ResponseRequired checks if the required fields are not zero-ed
func AssertRedfishV1OdataGet200ResponseRequired(obj RedfishV1OdataGet200Response) error {
	elements := map[string]interface{}{
		"@odata.context": obj.OdataContext,
		"value":          obj.Value,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Value {
		if err := AssertRedfishV1OdataGet200ResponseValueInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRedfishV1OdataGet200ResponseConstraints checks if the values respects the defined constraints
func AssertRedfishV1OdataGet200ResponseConstraints(obj RedfishV1OdataGet200Response) error {
	return nil
}
