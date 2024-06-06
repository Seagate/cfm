/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceStatus - The status and health of a resource and its children.
type ResourceStatus struct {

	// Conditions in this resource that require attention.
	Conditions []ResourceCondition `json:"Conditions,omitempty"`

	Health ResourceHealth `json:"Health,omitempty"`

	HealthRollup ResourceHealth `json:"HealthRollup,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	State ResourceState `json:"State,omitempty"`
}

// AssertResourceStatusRequired checks if the required fields are not zero-ed
func AssertResourceStatusRequired(obj ResourceStatus) error {
	for _, el := range obj.Conditions {
		if err := AssertResourceConditionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertResourceStatusConstraints checks if the values respects the defined constraints
func AssertResourceStatusConstraints(obj ResourceStatus) error {
	return nil
}