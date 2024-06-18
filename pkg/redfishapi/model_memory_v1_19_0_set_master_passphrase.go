/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryV1190SetMasterPassphrase - Sets the master passphrase for the given region.
type MemoryV1190SetMasterPassphrase struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertMemoryV1190SetMasterPassphraseRequired checks if the required fields are not zero-ed
func AssertMemoryV1190SetMasterPassphraseRequired(obj MemoryV1190SetMasterPassphrase) error {
	return nil
}

// AssertMemoryV1190SetMasterPassphraseConstraints checks if the values respects the defined constraints
func AssertMemoryV1190SetMasterPassphraseConstraints(obj MemoryV1190SetMasterPassphrase) error {
	return nil
}
