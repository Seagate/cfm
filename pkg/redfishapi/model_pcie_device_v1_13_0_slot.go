/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// PcieDeviceV1130Slot - The PCIe slot associated with a PCIe device.
type PcieDeviceV1130Slot struct {

	// An indication of whether this PCIe slot supports hotplug.
	HotPluggable *bool `json:"HotPluggable,omitempty"`

	LaneSplitting PcieDeviceV1130LaneSplittingType `json:"LaneSplitting,omitempty"`

	// The number of PCIe lanes supported by this slot.
	Lanes *int64 `json:"Lanes,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	PCIeType PcieDevicePcieTypes `json:"PCIeType,omitempty"`

	SlotType PcieDeviceV1130SlotType `json:"SlotType,omitempty"`
}

// AssertPcieDeviceV1130SlotRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130SlotRequired(obj PcieDeviceV1130Slot) error {
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	return nil
}

// AssertPcieDeviceV1130SlotConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130SlotConstraints(obj PcieDeviceV1130Slot) error {
	if obj.Lanes != nil && *obj.Lanes > 32 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
