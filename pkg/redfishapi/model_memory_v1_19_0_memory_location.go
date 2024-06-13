/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryV1190MemoryLocation - Memory connection information to sockets and memory controllers.
type MemoryV1190MemoryLocation struct {

	// The channel number to which the memory device is connected.
	Channel *int64 `json:"Channel,omitempty"`

	// The memory controller number to which the memory device is connected.
	MemoryController *int64 `json:"MemoryController,omitempty"`

	// The slot number to which the memory device is connected.
	Slot *int64 `json:"Slot,omitempty"`

	// The socket number to which the memory device is connected.
	Socket *int64 `json:"Socket,omitempty"`
}

// AssertMemoryV1190MemoryLocationRequired checks if the required fields are not zero-ed
func AssertMemoryV1190MemoryLocationRequired(obj MemoryV1190MemoryLocation) error {
	return nil
}

// AssertMemoryV1190MemoryLocationConstraints checks if the values respects the defined constraints
func AssertMemoryV1190MemoryLocationConstraints(obj MemoryV1190MemoryLocation) error {
	return nil
}
