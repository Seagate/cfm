/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PcieDeviceV1130CxlDeviceType string

// List of PcieDeviceV1130CxlDeviceType
const (
	PCIEDEVICEV1130CXLDEVICETYPE_TYPE1 PcieDeviceV1130CxlDeviceType = "Type1"
	PCIEDEVICEV1130CXLDEVICETYPE_TYPE2 PcieDeviceV1130CxlDeviceType = "Type2"
	PCIEDEVICEV1130CXLDEVICETYPE_TYPE3 PcieDeviceV1130CxlDeviceType = "Type3"
)

// AssertPcieDeviceV1130CxlDeviceTypeRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130CxlDeviceTypeRequired(obj PcieDeviceV1130CxlDeviceType) error {
	return nil
}

// AssertPcieDeviceV1130CxlDeviceTypeConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130CxlDeviceTypeConstraints(obj PcieDeviceV1130CxlDeviceType) error {
	return nil
}