/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// IpAddressesIpv6Address - This type describes an IPv6 address.
type IpAddressesIpv6Address struct {

	// The IPv6 address.
	Address *string `json:"Address,omitempty"`

	AddressOrigin *IpAddressesV115Ipv6AddressOrigin `json:"AddressOrigin,omitempty"`

	AddressState *IpAddressesV115AddressState `json:"AddressState,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PrefixLength int64 `json:"PrefixLength,omitempty"`
}

// AssertIpAddressesIpv6AddressRequired checks if the required fields are not zero-ed
func AssertIpAddressesIpv6AddressRequired(obj IpAddressesIpv6Address) error {
	return nil
}

// AssertIpAddressesIpv6AddressConstraints checks if the values respects the defined constraints
func AssertIpAddressesIpv6AddressConstraints(obj IpAddressesIpv6Address) error {
	if obj.PrefixLength < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PrefixLength > 128 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
