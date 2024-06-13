/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryV1190SecurityStateInfo - The security states of a memory device.
type MemoryV1190SecurityStateInfo struct {

	// An indication of whether an incorrect master passphrase attempt count has been reached.
	MasterPassphraseAttemptCountReached *bool `json:"MasterPassphraseAttemptCountReached,omitempty"`

	// An indication of whether an incorrect user passphrase attempt count has been reached.
	UserPassphraseAttemptCountReached *bool `json:"UserPassphraseAttemptCountReached,omitempty"`
}

// AssertMemoryV1190SecurityStateInfoRequired checks if the required fields are not zero-ed
func AssertMemoryV1190SecurityStateInfoRequired(obj MemoryV1190SecurityStateInfo) error {
	return nil
}

// AssertMemoryV1190SecurityStateInfoConstraints checks if the values respects the defined constraints
func AssertMemoryV1190SecurityStateInfoConstraints(obj MemoryV1190SecurityStateInfo) error {
	return nil
}
