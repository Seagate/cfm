/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ConnectionV131MemoryChunkInfo - The combination of permissions and memory chunk information.
type ConnectionV131MemoryChunkInfo struct {

	// Supported I/O access capabilities.
	AccessCapabilities []ConnectionV131AccessCapability `json:"AccessCapabilities,omitempty"`

	AccessState ConnectionV131AccessState `json:"AccessState,omitempty"`

	MemoryChunk OdataV4IdRef `json:"MemoryChunk,omitempty"`
}

// AssertConnectionV131MemoryChunkInfoRequired checks if the required fields are not zero-ed
func AssertConnectionV131MemoryChunkInfoRequired(obj ConnectionV131MemoryChunkInfo) error {
	if err := AssertOdataV4IdRefRequired(obj.MemoryChunk); err != nil {
		return err
	}
	return nil
}

// AssertConnectionV131MemoryChunkInfoConstraints checks if the values respects the defined constraints
func AssertConnectionV131MemoryChunkInfoConstraints(obj ConnectionV131MemoryChunkInfo) error {
	return nil
}