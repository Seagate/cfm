/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// EndpointV181ConnectedEntity - Represents a remote resource that is connected to the network accessible to this endpoint.
type EndpointV181ConnectedEntity struct {
	EntityLink OdataV4IdRef `json:"EntityLink,omitempty"`

	EntityPciId EndpointV181PciId `json:"EntityPciId,omitempty"`

	EntityRole EndpointV181EntityRole `json:"EntityRole,omitempty"`

	EntityType EndpointV181EntityType `json:"EntityType,omitempty"`

	GenZ EndpointV181GenZ `json:"GenZ,omitempty"`

	// Identifiers for the remote entity.
	Identifiers []ResourceIdentifier `json:"Identifiers,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Class Code, Subclass, and Programming Interface code of this PCIe function.
	// Deprecated
	PciClassCode *string `json:"PciClassCode,omitempty"`

	// The PCI ID of the connected entity.
	// Deprecated
	PciFunctionNumber *int64 `json:"PciFunctionNumber,omitempty"`
}

// AssertEndpointV181ConnectedEntityRequired checks if the required fields are not zero-ed
func AssertEndpointV181ConnectedEntityRequired(obj EndpointV181ConnectedEntity) error {
	if err := AssertOdataV4IdRefRequired(obj.EntityLink); err != nil {
		return err
	}
	if err := AssertEndpointV181PciIdRequired(obj.EntityPciId); err != nil {
		return err
	}
	if err := AssertEndpointV181GenZRequired(obj.GenZ); err != nil {
		return err
	}
	for _, el := range obj.Identifiers {
		if err := AssertResourceIdentifierRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEndpointV181ConnectedEntityConstraints checks if the values respects the defined constraints
func AssertEndpointV181ConnectedEntityConstraints(obj EndpointV181ConnectedEntity) error {
	return nil
}
