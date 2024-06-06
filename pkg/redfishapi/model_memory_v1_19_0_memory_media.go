/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type MemoryV1190MemoryMedia string

// List of MemoryV1190MemoryMedia
const (
	MEMORYV1190MEMORYMEDIA_DRAM            MemoryV1190MemoryMedia = "DRAM"
	MEMORYV1190MEMORYMEDIA_NAND            MemoryV1190MemoryMedia = "NAND"
	MEMORYV1190MEMORYMEDIA_INTEL3_DX_POINT MemoryV1190MemoryMedia = "Intel3DXPoint"
	MEMORYV1190MEMORYMEDIA_PROPRIETARY     MemoryV1190MemoryMedia = "Proprietary"
)

// AssertMemoryV1190MemoryMediaRequired checks if the required fields are not zero-ed
func AssertMemoryV1190MemoryMediaRequired(obj MemoryV1190MemoryMedia) error {
	return nil
}

// AssertMemoryV1190MemoryMediaConstraints checks if the values respects the defined constraints
func AssertMemoryV1190MemoryMediaConstraints(obj MemoryV1190MemoryMedia) error {
	return nil
}
