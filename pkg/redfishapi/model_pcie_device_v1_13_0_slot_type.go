/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PcieDeviceV1130SlotType string

// List of PcieDeviceV1130SlotType
const (
	PCIEDEVICEV1130SLOTTYPE_FULL_LENGTH PcieDeviceV1130SlotType = "FullLength"
	PCIEDEVICEV1130SLOTTYPE_HALF_LENGTH PcieDeviceV1130SlotType = "HalfLength"
	PCIEDEVICEV1130SLOTTYPE_LOW_PROFILE PcieDeviceV1130SlotType = "LowProfile"
	PCIEDEVICEV1130SLOTTYPE_MINI        PcieDeviceV1130SlotType = "Mini"
	PCIEDEVICEV1130SLOTTYPE_M2          PcieDeviceV1130SlotType = "M2"
	PCIEDEVICEV1130SLOTTYPE_OEM         PcieDeviceV1130SlotType = "OEM"
	PCIEDEVICEV1130SLOTTYPE_OCP3_SMALL  PcieDeviceV1130SlotType = "OCP3Small"
	PCIEDEVICEV1130SLOTTYPE_OCP3_LARGE  PcieDeviceV1130SlotType = "OCP3Large"
	PCIEDEVICEV1130SLOTTYPE_U2          PcieDeviceV1130SlotType = "U2"
)

// AssertPcieDeviceV1130SlotTypeRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130SlotTypeRequired(obj PcieDeviceV1130SlotType) error {
	return nil
}

// AssertPcieDeviceV1130SlotTypeConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130SlotTypeConstraints(obj PcieDeviceV1130SlotType) error {
	return nil
}
