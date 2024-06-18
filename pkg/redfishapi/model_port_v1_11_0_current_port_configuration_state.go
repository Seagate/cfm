/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110CurrentPortConfigurationState string

// List of PortV1110CurrentPortConfigurationState
const (
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_DISABLED           PortV1110CurrentPortConfigurationState = "Disabled"
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_BIND_IN_PROGRESS   PortV1110CurrentPortConfigurationState = "BindInProgress"
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_UNBIND_IN_PROGRESS PortV1110CurrentPortConfigurationState = "UnbindInProgress"
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_DSP                PortV1110CurrentPortConfigurationState = "DSP"
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_USP                PortV1110CurrentPortConfigurationState = "USP"
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_RESERVED           PortV1110CurrentPortConfigurationState = "Reserved"
	PORTV1110CURRENTPORTCONFIGURATIONSTATE_FABRIC_LINK        PortV1110CurrentPortConfigurationState = "FabricLink"
)

// AssertPortV1110CurrentPortConfigurationStateRequired checks if the required fields are not zero-ed
func AssertPortV1110CurrentPortConfigurationStateRequired(obj PortV1110CurrentPortConfigurationState) error {
	return nil
}

// AssertPortV1110CurrentPortConfigurationStateConstraints checks if the values respects the defined constraints
func AssertPortV1110CurrentPortConfigurationStateConstraints(obj PortV1110CurrentPortConfigurationState) error {
	return nil
}
