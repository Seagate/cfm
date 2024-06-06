/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceLocation - The location of a resource.
type ResourceLocation struct {

	// The altitude of the resource in meters.
	AltitudeMeters *float32 `json:"AltitudeMeters,omitempty"`

	// An array of contact information.
	Contacts []ResourceV1180ContactInfo `json:"Contacts,omitempty"`

	// The location of the resource.
	// Deprecated
	Info *string `json:"Info,omitempty"`

	// The format of the Info property.
	// Deprecated
	InfoFormat *string `json:"InfoFormat,omitempty"`

	// The latitude of the resource.
	Latitude *float32 `json:"Latitude,omitempty"`

	// The longitude of the resource in degree units.
	Longitude *float32 `json:"Longitude,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PartLocation ResourceV1180PartLocation `json:"PartLocation,omitempty"`

	// Human-readable string to enable differentiation between PartLocation values for parts in the same enclosure, which might include hierarchical information of containing PartLocation values for the part.
	PartLocationContext *string `json:"PartLocationContext,omitempty"`

	PhysicalAddress ResourceV1180PhysicalAddress `json:"PhysicalAddress,omitempty"`

	Placement ResourceV1180Placement `json:"Placement,omitempty"`

	// Deprecated
	PostalAddress ResourceV1180PostalAddress `json:"PostalAddress,omitempty"`
}

// AssertResourceLocationRequired checks if the required fields are not zero-ed
func AssertResourceLocationRequired(obj ResourceLocation) error {
	for _, el := range obj.Contacts {
		if err := AssertResourceV1180ContactInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertResourceV1180PartLocationRequired(obj.PartLocation); err != nil {
		return err
	}
	if err := AssertResourceV1180PhysicalAddressRequired(obj.PhysicalAddress); err != nil {
		return err
	}
	if err := AssertResourceV1180PlacementRequired(obj.Placement); err != nil {
		return err
	}
	if err := AssertResourceV1180PostalAddressRequired(obj.PostalAddress); err != nil {
		return err
	}
	return nil
}

// AssertResourceLocationConstraints checks if the values respects the defined constraints
func AssertResourceLocationConstraints(obj ResourceLocation) error {
	return nil
}