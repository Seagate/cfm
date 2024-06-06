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

// PortV1110LldpTransmit - Link Layer Data Protocol (LLDP) data being transmitted on this link.
type PortV1110LldpTransmit struct {

	// Link Layer Data Protocol (LLDP) chassis ID.
	ChassisId *string `json:"ChassisId,omitempty"`

	ChassisIdSubtype PortV1110Ieee802IdSubtype `json:"ChassisIdSubtype,omitempty"`

	// The IPv4 management address to be transmitted from this endpoint.
	ManagementAddressIPv4 *string `json:"ManagementAddressIPv4,omitempty"`

	// The IPv6 management address to be transmitted from this endpoint.
	ManagementAddressIPv6 *string `json:"ManagementAddressIPv6,omitempty"`

	// The management MAC address to be transmitted from this endpoint.
	ManagementAddressMAC *string `json:"ManagementAddressMAC,omitempty"`

	// The management VLAN ID to be transmitted from this endpoint.
	ManagementVlanId *int64 `json:"ManagementVlanId,omitempty"`

	// A colon-delimited string of hexadecimal octets identifying a port to be transmitted from this endpoint.
	PortId *string `json:"PortId,omitempty"`

	PortIdSubtype PortV1110Ieee802IdSubtype `json:"PortIdSubtype,omitempty"`

	// The system capabilities to be transmitted from this endpoint.
	SystemCapabilities []PortV1110LldpSystemCapabilities `json:"SystemCapabilities,omitempty"`

	// The system description to be transmitted from this endpoint.
	SystemDescription *string `json:"SystemDescription,omitempty"`

	// The system name to be transmitted from this endpoint.
	SystemName *string `json:"SystemName,omitempty"`
}

// AssertPortV1110LldpTransmitRequired checks if the required fields are not zero-ed
func AssertPortV1110LldpTransmitRequired(obj PortV1110LldpTransmit) error {
	return nil
}

// AssertPortV1110LldpTransmitConstraints checks if the values respects the defined constraints
func AssertPortV1110LldpTransmitConstraints(obj PortV1110LldpTransmit) error {
	if obj.ManagementVlanId != nil && *obj.ManagementVlanId < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ManagementVlanId != nil && *obj.ManagementVlanId > 4095 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
