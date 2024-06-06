/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryV1190DisableMasterPassphrase - Disables the master passphrase for the given region.
type MemoryV1190DisableMasterPassphrase struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertMemoryV1190DisableMasterPassphraseRequired checks if the required fields are not zero-ed
func AssertMemoryV1190DisableMasterPassphraseRequired(obj MemoryV1190DisableMasterPassphrase) error {
	return nil
}

// AssertMemoryV1190DisableMasterPassphraseConstraints checks if the values respects the defined constraints
func AssertMemoryV1190DisableMasterPassphraseConstraints(obj MemoryV1190DisableMasterPassphrase) error {
	return nil
}