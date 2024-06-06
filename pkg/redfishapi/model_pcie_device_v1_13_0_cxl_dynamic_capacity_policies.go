/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PcieDeviceV1130CxlDynamicCapacityPolicies string

// List of PcieDeviceV1130CxlDynamicCapacityPolicies
const (
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_FREE         PcieDeviceV1130CxlDynamicCapacityPolicies = "Free"
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_CONTIGUOUS   PcieDeviceV1130CxlDynamicCapacityPolicies = "Contiguous"
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_PRESCRIPTIVE PcieDeviceV1130CxlDynamicCapacityPolicies = "Prescriptive"
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_TAG_BASED    PcieDeviceV1130CxlDynamicCapacityPolicies = "TagBased"
)

// AssertPcieDeviceV1130CxlDynamicCapacityPoliciesRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130CxlDynamicCapacityPoliciesRequired(obj PcieDeviceV1130CxlDynamicCapacityPolicies) error {
	return nil
}

// AssertPcieDeviceV1130CxlDynamicCapacityPoliciesConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130CxlDynamicCapacityPoliciesConstraints(obj PcieDeviceV1130CxlDynamicCapacityPolicies) error {
	return nil
}