/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ServiceRootV1161Links - The links to other resources that are related to this resource.
type ServiceRootV1161Links struct {
	ManagerProvidingService OdataV4IdRef `json:"ManagerProvidingService,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	Sessions OdataV4IdRef `json:"Sessions"`
}

// AssertServiceRootV1161LinksRequired checks if the required fields are not zero-ed
func AssertServiceRootV1161LinksRequired(obj ServiceRootV1161Links) error {
	elements := map[string]interface{}{
		"Sessions": obj.Sessions,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertOdataV4IdRefRequired(obj.ManagerProvidingService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Sessions); err != nil {
		return err
	}
	return nil
}

// AssertServiceRootV1161LinksConstraints checks if the values respects the defined constraints
func AssertServiceRootV1161LinksConstraints(obj ServiceRootV1161Links) error {
	return nil
}