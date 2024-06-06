/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PortV1110Links - The links to other resources that are related to this resource.
type PortV1110Links struct {

	// An array of links to the endpoints at the other end of the link.
	AssociatedEndpoints []OdataV4IdRef `json:"AssociatedEndpoints,omitempty"`

	// The number of items in a collection.
	AssociatedEndpointsodataCount int64 `json:"AssociatedEndpoints@odata.count,omitempty"`

	// An array of links to the cables connected to this port.
	Cables []OdataV4IdRef `json:"Cables,omitempty"`

	// The number of items in a collection.
	CablesodataCount int64 `json:"Cables@odata.count,omitempty"`

	// An array of links to the remote device ports at the other end of the link.
	ConnectedPorts []OdataV4IdRef `json:"ConnectedPorts,omitempty"`

	// The number of items in a collection.
	ConnectedPortsodataCount int64 `json:"ConnectedPorts@odata.count,omitempty"`

	// An array of links to the switch ports at the other end of the link.
	ConnectedSwitchPorts []OdataV4IdRef `json:"ConnectedSwitchPorts,omitempty"`

	// The number of items in a collection.
	ConnectedSwitchPortsodataCount int64 `json:"ConnectedSwitchPorts@odata.count,omitempty"`

	// An array of links to the switches at the other end of the link.
	ConnectedSwitches []OdataV4IdRef `json:"ConnectedSwitches,omitempty"`

	// The number of items in a collection.
	ConnectedSwitchesodataCount int64 `json:"ConnectedSwitches@odata.count,omitempty"`

	// The links to the Ethernet interfaces this port provides.
	EthernetInterfaces []OdataV4IdRef `json:"EthernetInterfaces,omitempty"`

	// The number of items in a collection.
	EthernetInterfacesodataCount int64 `json:"EthernetInterfaces@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertPortV1110LinksRequired checks if the required fields are not zero-ed
func AssertPortV1110LinksRequired(obj PortV1110Links) error {
	for _, el := range obj.AssociatedEndpoints {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Cables {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnectedPorts {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnectedSwitchPorts {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnectedSwitches {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EthernetInterfaces {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPortV1110LinksConstraints checks if the values respects the defined constraints
func AssertPortV1110LinksConstraints(obj PortV1110Links) error {
	return nil
}