/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type MemoryV1190ErrorCorrection string

// List of MemoryV1190ErrorCorrection
const (
	MEMORYV1190ERRORCORRECTION_NO_ECC         MemoryV1190ErrorCorrection = "NoECC"
	MEMORYV1190ERRORCORRECTION_SINGLE_BIT_ECC MemoryV1190ErrorCorrection = "SingleBitECC"
	MEMORYV1190ERRORCORRECTION_MULTI_BIT_ECC  MemoryV1190ErrorCorrection = "MultiBitECC"
	MEMORYV1190ERRORCORRECTION_ADDRESS_PARITY MemoryV1190ErrorCorrection = "AddressParity"
)

// AssertMemoryV1190ErrorCorrectionRequired checks if the required fields are not zero-ed
func AssertMemoryV1190ErrorCorrectionRequired(obj MemoryV1190ErrorCorrection) error {
	return nil
}

// AssertMemoryV1190ErrorCorrectionConstraints checks if the values respects the defined constraints
func AssertMemoryV1190ErrorCorrectionConstraints(obj MemoryV1190ErrorCorrection) error {
	return nil
}
