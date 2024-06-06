/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PortV1110Sfp - A small form-factor pluggable (SFP) device attached to a port.
type PortV1110Sfp struct {
	FiberConnectionType PortV1110FiberConnectionType `json:"FiberConnectionType,omitempty"`

	// The manufacturer of this SFP.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	MediumType PortV1110MediumType `json:"MediumType,omitempty"`

	// The part number for this SFP.
	PartNumber *string `json:"PartNumber,omitempty"`

	// The serial number for this SFP.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The types of SFP devices that can be attached to this port.
	SupportedSFPTypes []PortV1110SfpType `json:"SupportedSFPTypes,omitempty"`

	Type PortV1110SfpType `json:"Type,omitempty"`
}

// AssertPortV1110SfpRequired checks if the required fields are not zero-ed
func AssertPortV1110SfpRequired(obj PortV1110Sfp) error {
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertPortV1110SfpConstraints checks if the values respects the defined constraints
func AssertPortV1110SfpConstraints(obj PortV1110Sfp) error {
	return nil
}