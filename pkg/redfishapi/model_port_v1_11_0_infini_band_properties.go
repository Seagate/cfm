/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PortV1110InfiniBandProperties - InfiniBand-specific properties for a port.
type PortV1110InfiniBandProperties struct {

	// An array of configured node GUIDs that are associated with this network port, including the programmed address of the lowest-numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedNodeGUIDs []*string `json:"AssociatedNodeGUIDs,omitempty"`

	// An array of configured port GUIDs that are associated with this network port, including the programmed address of the lowest-numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedPortGUIDs []*string `json:"AssociatedPortGUIDs,omitempty"`

	// An array of configured system GUIDs that are associated with this network port, including the programmed address of the lowest-numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedSystemGUIDs []*string `json:"AssociatedSystemGUIDs,omitempty"`
}

// AssertPortV1110InfiniBandPropertiesRequired checks if the required fields are not zero-ed
func AssertPortV1110InfiniBandPropertiesRequired(obj PortV1110InfiniBandProperties) error {
	return nil
}

// AssertPortV1110InfiniBandPropertiesConstraints checks if the values respects the defined constraints
func AssertPortV1110InfiniBandPropertiesConstraints(obj PortV1110InfiniBandProperties) error {
	return nil
}