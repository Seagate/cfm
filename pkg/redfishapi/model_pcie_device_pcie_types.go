/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PcieDevicePcieTypes string

// List of PcieDevicePcieTypes
const (
	PCIEDEVICEPCIETYPES_GEN1 PcieDevicePcieTypes = "Gen1"
	PCIEDEVICEPCIETYPES_GEN2 PcieDevicePcieTypes = "Gen2"
	PCIEDEVICEPCIETYPES_GEN3 PcieDevicePcieTypes = "Gen3"
	PCIEDEVICEPCIETYPES_GEN4 PcieDevicePcieTypes = "Gen4"
	PCIEDEVICEPCIETYPES_GEN5 PcieDevicePcieTypes = "Gen5"
)

// AssertPcieDevicePcieTypesRequired checks if the required fields are not zero-ed
func AssertPcieDevicePcieTypesRequired(obj PcieDevicePcieTypes) error {
	return nil
}

// AssertPcieDevicePcieTypesConstraints checks if the values respects the defined constraints
func AssertPcieDevicePcieTypesConstraints(obj PcieDevicePcieTypes) error {
	return nil
}
