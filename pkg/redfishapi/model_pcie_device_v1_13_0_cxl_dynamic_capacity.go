/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PcieDeviceV1130CxlDynamicCapacity - The CXL dynamic capacity device (DCD) information for a CXL device.
type PcieDeviceV1130CxlDynamicCapacity struct {

	// The set of selection policies supported by the CXL device when dynamic capacity is added.
	AddCapacityPoliciesSupported []PcieDeviceV1130CxlDynamicCapacityPolicies `json:"AddCapacityPoliciesSupported,omitempty"`

	// The maximum number of dynamic capacity memory regions available per host from this CXL device.
	MaxDynamicCapacityRegions *int64 `json:"MaxDynamicCapacityRegions,omitempty"`

	// The maximum number of hosts supported by this CXL device.
	MaxHosts *int64 `json:"MaxHosts,omitempty"`

	// The set of memory block sizes supported by memory regions in this CXL device.
	MemoryBlockSizesSupported []PcieDeviceV1130CxlRegionBlockSizes `json:"MemoryBlockSizesSupported,omitempty"`

	// The set of removal policies supported by the CXL device when dynamic capacity is released.
	ReleaseCapacityPoliciesSupported []PcieDeviceV1130CxlDynamicCapacityPolicies `json:"ReleaseCapacityPoliciesSupported,omitempty"`

	// An indication of whether the sanitization on capacity release is configurable for the memory regions in this CXL device.
	SanitizationOnReleaseSupport []PcieDeviceV1130CxlRegionSanitization `json:"SanitizationOnReleaseSupport,omitempty"`

	// The total memory media capacity of the CXL device available for dynamic assignment in mebibytes (MiB).
	TotalDynamicCapacityMiB *int64 `json:"TotalDynamicCapacityMiB,omitempty"`
}

// AssertPcieDeviceV1130CxlDynamicCapacityRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130CxlDynamicCapacityRequired(obj PcieDeviceV1130CxlDynamicCapacity) error {
	for _, el := range obj.MemoryBlockSizesSupported {
		if err := AssertPcieDeviceV1130CxlRegionBlockSizesRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SanitizationOnReleaseSupport {
		if err := AssertPcieDeviceV1130CxlRegionSanitizationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPcieDeviceV1130CxlDynamicCapacityConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130CxlDynamicCapacityConstraints(obj PcieDeviceV1130CxlDynamicCapacity) error {
	return nil
}
