/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// PortV1110LldpReceive - Link Layer Data Protocol (LLDP) data received from the remote partner across this link.
type PortV1110LldpReceive struct {

	// Link Layer Data Protocol (LLDP) chassis ID received from the remote partner across this link.
	ChassisId *string `json:"ChassisId,omitempty"`

	ChassisIdSubtype PortV1110Ieee802IdSubtype `json:"ChassisIdSubtype,omitempty"`

	// The IPv4 management address received from the remote partner across this link.
	ManagementAddressIPv4 *string `json:"ManagementAddressIPv4,omitempty"`

	// The IPv6 management address received from the remote partner across this link.
	ManagementAddressIPv6 *string `json:"ManagementAddressIPv6,omitempty"`

	// The management MAC address received from the remote partner across this link.
	ManagementAddressMAC *string `json:"ManagementAddressMAC,omitempty"`

	// The management VLAN ID received from the remote partner across this link.
	ManagementVlanId *int64 `json:"ManagementVlanId,omitempty"`

	// A colon-delimited string of hexadecimal octets identifying a port.
	PortId *string `json:"PortId,omitempty"`

	PortIdSubtype PortV1110Ieee802IdSubtype `json:"PortIdSubtype,omitempty"`

	// The system capabilities received from the remote partner across this link.
	SystemCapabilities []PortV1110LldpSystemCapabilities `json:"SystemCapabilities,omitempty"`

	// The system description received from the remote partner across this link.
	SystemDescription *string `json:"SystemDescription,omitempty"`

	// The system name received from the remote partner across this link.
	SystemName *string `json:"SystemName,omitempty"`
}

// AssertPortV1110LldpReceiveRequired checks if the required fields are not zero-ed
func AssertPortV1110LldpReceiveRequired(obj PortV1110LldpReceive) error {
	return nil
}

// AssertPortV1110LldpReceiveConstraints checks if the values respects the defined constraints
func AssertPortV1110LldpReceiveConstraints(obj PortV1110LldpReceive) error {
	if obj.ManagementVlanId != nil && *obj.ManagementVlanId < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ManagementVlanId != nil && *obj.ManagementVlanId > 4095 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
