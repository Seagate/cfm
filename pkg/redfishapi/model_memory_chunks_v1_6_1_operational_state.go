/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type MemoryChunksV161OperationalState string

// List of MemoryChunksV161OperationalState
const (
	MEMORYCHUNKSV161OPERATIONALSTATE_ONLINE  MemoryChunksV161OperationalState = "Online"
	MEMORYCHUNKSV161OPERATIONALSTATE_OFFLINE MemoryChunksV161OperationalState = "Offline"
)

// AssertMemoryChunksV161OperationalStateRequired checks if the required fields are not zero-ed
func AssertMemoryChunksV161OperationalStateRequired(obj MemoryChunksV161OperationalState) error {
	return nil
}

// AssertMemoryChunksV161OperationalStateConstraints checks if the values respects the defined constraints
func AssertMemoryChunksV161OperationalStateConstraints(obj MemoryChunksV161OperationalState) error {
	return nil
}
