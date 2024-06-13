/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type MemoryV1190MemoryType string

// List of MemoryV1190MemoryType
const (
	MEMORYV1190MEMORYTYPE_DRAM         MemoryV1190MemoryType = "DRAM"
	MEMORYV1190MEMORYTYPE_NVDIMM_N     MemoryV1190MemoryType = "NVDIMM_N"
	MEMORYV1190MEMORYTYPE_NVDIMM_F     MemoryV1190MemoryType = "NVDIMM_F"
	MEMORYV1190MEMORYTYPE_NVDIMM_P     MemoryV1190MemoryType = "NVDIMM_P"
	MEMORYV1190MEMORYTYPE_INTEL_OPTANE MemoryV1190MemoryType = "IntelOptane"
)

// AssertMemoryV1190MemoryTypeRequired checks if the required fields are not zero-ed
func AssertMemoryV1190MemoryTypeRequired(obj MemoryV1190MemoryType) error {
	return nil
}

// AssertMemoryV1190MemoryTypeConstraints checks if the values respects the defined constraints
func AssertMemoryV1190MemoryTypeConstraints(obj MemoryV1190MemoryType) error {
	return nil
}
