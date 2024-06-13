/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110LldpSystemCapabilities string

// List of PortV1110LldpSystemCapabilities
const (
	PORTV1110LLDPSYSTEMCAPABILITIES_NONE                PortV1110LldpSystemCapabilities = "None"
	PORTV1110LLDPSYSTEMCAPABILITIES_BRIDGE              PortV1110LldpSystemCapabilities = "Bridge"
	PORTV1110LLDPSYSTEMCAPABILITIES_DOCSIS_CABLE_DEVICE PortV1110LldpSystemCapabilities = "DOCSISCableDevice"
	PORTV1110LLDPSYSTEMCAPABILITIES_OTHER               PortV1110LldpSystemCapabilities = "Other"
	PORTV1110LLDPSYSTEMCAPABILITIES_REPEATER            PortV1110LldpSystemCapabilities = "Repeater"
	PORTV1110LLDPSYSTEMCAPABILITIES_ROUTER              PortV1110LldpSystemCapabilities = "Router"
	PORTV1110LLDPSYSTEMCAPABILITIES_STATION             PortV1110LldpSystemCapabilities = "Station"
	PORTV1110LLDPSYSTEMCAPABILITIES_TELEPHONE           PortV1110LldpSystemCapabilities = "Telephone"
	PORTV1110LLDPSYSTEMCAPABILITIES_WLAN_ACCESS_POINT   PortV1110LldpSystemCapabilities = "WLANAccessPoint"
)

// AssertPortV1110LldpSystemCapabilitiesRequired checks if the required fields are not zero-ed
func AssertPortV1110LldpSystemCapabilitiesRequired(obj PortV1110LldpSystemCapabilities) error {
	return nil
}

// AssertPortV1110LldpSystemCapabilitiesConstraints checks if the values respects the defined constraints
func AssertPortV1110LldpSystemCapabilitiesConstraints(obj PortV1110LldpSystemCapabilities) error {
	return nil
}
