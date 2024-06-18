/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PcieFunctionV151Links - The links to other resources that are related to this resource.
type PcieFunctionV151Links struct {
	CXLLogicalDevice OdataV4IdRef `json:"CXLLogicalDevice,omitempty"`

	// An array of links to the drives that this PCIe function produces.
	Drives []OdataV4IdRef `json:"Drives,omitempty"`

	// The number of items in a collection.
	DrivesodataCount int64 `json:"Drives@odata.count,omitempty"`

	// An array of links to the Ethernet interfaces that this PCIe function produces.
	EthernetInterfaces []OdataV4IdRef `json:"EthernetInterfaces,omitempty"`

	// The number of items in a collection.
	EthernetInterfacesodataCount int64 `json:"EthernetInterfaces@odata.count,omitempty"`

	// An array of links to the memory domains that the PCIe function produces.
	MemoryDomains []OdataV4IdRef `json:"MemoryDomains,omitempty"`

	// The number of items in a collection.
	MemoryDomainsodataCount int64 `json:"MemoryDomains@odata.count,omitempty"`

	// An array of links to the network device functions that the PCIe function produces.
	NetworkDeviceFunctions []OdataV4IdRef `json:"NetworkDeviceFunctions,omitempty"`

	// The number of items in a collection.
	NetworkDeviceFunctionsodataCount int64 `json:"NetworkDeviceFunctions@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeDevice OdataV4IdRef `json:"PCIeDevice,omitempty"`

	Processor OdataV4IdRef `json:"Processor,omitempty"`

	// An array of links to the storage controllers that this PCIe function produces.
	StorageControllers []StorageStorageController `json:"StorageControllers,omitempty"`

	// The number of items in a collection.
	StorageControllersodataCount int64 `json:"StorageControllers@odata.count,omitempty"`
}

// AssertPcieFunctionV151LinksRequired checks if the required fields are not zero-ed
func AssertPcieFunctionV151LinksRequired(obj PcieFunctionV151Links) error {
	if err := AssertOdataV4IdRefRequired(obj.CXLLogicalDevice); err != nil {
		return err
	}
	for _, el := range obj.Drives {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EthernetInterfaces {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MemoryDomains {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NetworkDeviceFunctions {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.PCIeDevice); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Processor); err != nil {
		return err
	}
	for _, el := range obj.StorageControllers {
		if err := AssertStorageStorageControllerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPcieFunctionV151LinksConstraints checks if the values respects the defined constraints
func AssertPcieFunctionV151LinksConstraints(obj PcieFunctionV151Links) error {
	return nil
}