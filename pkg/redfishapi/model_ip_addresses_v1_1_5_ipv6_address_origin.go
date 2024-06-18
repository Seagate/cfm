/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type IpAddressesV115Ipv6AddressOrigin string

// List of IpAddressesV115Ipv6AddressOrigin
const (
	IPADDRESSESV115IPV6ADDRESSORIGIN_STATIC     IpAddressesV115Ipv6AddressOrigin = "Static"
	IPADDRESSESV115IPV6ADDRESSORIGIN_DHCPV6     IpAddressesV115Ipv6AddressOrigin = "DHCPv6"
	IPADDRESSESV115IPV6ADDRESSORIGIN_LINK_LOCAL IpAddressesV115Ipv6AddressOrigin = "LinkLocal"
	IPADDRESSESV115IPV6ADDRESSORIGIN_SLAAC      IpAddressesV115Ipv6AddressOrigin = "SLAAC"
)

// AssertIpAddressesV115Ipv6AddressOriginRequired checks if the required fields are not zero-ed
func AssertIpAddressesV115Ipv6AddressOriginRequired(obj IpAddressesV115Ipv6AddressOrigin) error {
	return nil
}

// AssertIpAddressesV115Ipv6AddressOriginConstraints checks if the values respects the defined constraints
func AssertIpAddressesV115Ipv6AddressOriginConstraints(obj IpAddressesV115Ipv6AddressOrigin) error {
	return nil
}