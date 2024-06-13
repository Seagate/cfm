/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryV1190SetPassphrase - Set passphrase for the given regions.
type MemoryV1190SetPassphrase struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertMemoryV1190SetPassphraseRequired checks if the required fields are not zero-ed
func AssertMemoryV1190SetPassphraseRequired(obj MemoryV1190SetPassphrase) error {
	return nil
}

// AssertMemoryV1190SetPassphraseConstraints checks if the values respects the defined constraints
func AssertMemoryV1190SetPassphraseConstraints(obj MemoryV1190SetPassphrase) error {
	return nil
}
