/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceV1190Location - The location of a resource.
type ResourceV1190Location struct {

	// The altitude of the resource in meters.
	AltitudeMeters *float32 `json:"AltitudeMeters,omitempty"`

	// An array of contact information.
	Contacts []ResourceLocationContactsInner `json:"Contacts,omitempty"`

	// The location of the resource.
	// Deprecated
	Info *string `json:"Info,omitempty"`

	// The format of the `Info` property.
	// Deprecated
	InfoFormat *string `json:"InfoFormat,omitempty"`

	// The latitude of the resource.
	Latitude *float32 `json:"Latitude,omitempty"`

	// The longitude of the resource in degree units.
	Longitude *float32 `json:"Longitude,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PartLocation ResourceV1190PartLocation `json:"PartLocation,omitempty"`

	// Human-readable string to enable differentiation between `PartLocation` values for parts in the same enclosure, which might include hierarchical information of containing `PartLocation` values for the part.
	PartLocationContext *string `json:"PartLocationContext,omitempty"`

	PhysicalAddress ResourceV1190PhysicalAddress `json:"PhysicalAddress,omitempty"`

	Placement ResourceV1190Placement `json:"Placement,omitempty"`

	// Deprecated
	PostalAddress ResourceV1190PostalAddress `json:"PostalAddress,omitempty"`
}

// AssertResourceV1190LocationRequired checks if the required fields are not zero-ed
func AssertResourceV1190LocationRequired(obj ResourceV1190Location) error {
	for _, el := range obj.Contacts {
		if err := AssertResourceLocationContactsInnerRequired(el); err != nil {
			return err
		}
	}
	if err := AssertResourceV1190PartLocationRequired(obj.PartLocation); err != nil {
		return err
	}
	if err := AssertResourceV1190PhysicalAddressRequired(obj.PhysicalAddress); err != nil {
		return err
	}
	if err := AssertResourceV1190PlacementRequired(obj.Placement); err != nil {
		return err
	}
	if err := AssertResourceV1190PostalAddressRequired(obj.PostalAddress); err != nil {
		return err
	}
	return nil
}

// AssertResourceV1190LocationConstraints checks if the values respects the defined constraints
func AssertResourceV1190LocationConstraints(obj ResourceV1190Location) error {
	return nil
}