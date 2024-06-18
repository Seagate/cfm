/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceStatus - The status and health of a resource and its children.
type ResourceStatus struct {

	// Conditions in this resource that require attention.
	Conditions []ResourceStatusConditionsInner `json:"Conditions,omitempty"`

	Health *ResourceHealth `json:"Health,omitempty"`

	HealthRollup *ResourceHealth `json:"HealthRollup,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	State *ResourceState `json:"State,omitempty"`
}

// AssertResourceStatusRequired checks if the required fields are not zero-ed
func AssertResourceStatusRequired(obj ResourceStatus) error {
	for _, el := range obj.Conditions {
		if err := AssertResourceStatusConditionsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertResourceStatusConstraints checks if the values respects the defined constraints
func AssertResourceStatusConstraints(obj ResourceStatus) error {
	return nil
}
