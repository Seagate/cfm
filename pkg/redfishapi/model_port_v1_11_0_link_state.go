/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110LinkState string

// List of PortV1110LinkState
const (
	PORTV1110LINKSTATE_ENABLED  PortV1110LinkState = "Enabled"
	PORTV1110LINKSTATE_DISABLED PortV1110LinkState = "Disabled"
)

// AssertPortV1110LinkStateRequired checks if the required fields are not zero-ed
func AssertPortV1110LinkStateRequired(obj PortV1110LinkState) error {
	return nil
}

// AssertPortV1110LinkStateConstraints checks if the values respects the defined constraints
func AssertPortV1110LinkStateConstraints(obj PortV1110LinkState) error {
	return nil
}
