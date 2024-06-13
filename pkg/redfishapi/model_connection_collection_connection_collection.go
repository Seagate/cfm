/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ConnectionCollectionConnectionCollection - The collection of Connection resource instances.
type ConnectionCollectionConnectionCollection struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []OdataV4IdRef `json:"Members"`

	// The number of items in a collection.
	MembersodataCount int64 `json:"Members@odata.count"`

	// The URI to the resource containing the next set of partial members.
	MembersodataNextLink string `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertConnectionCollectionConnectionCollectionRequired checks if the required fields are not zero-ed
func AssertConnectionCollectionConnectionCollectionRequired(obj ConnectionCollectionConnectionCollection) error {
	elements := map[string]interface{}{
		"@odata.id":           obj.OdataId,
		"@odata.type":         obj.OdataType,
		"Members":             obj.Members,
		"Members@odata.count": obj.MembersodataCount,
		"Name":                obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Members {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertConnectionCollectionConnectionCollectionConstraints checks if the values respects the defined constraints
func AssertConnectionCollectionConnectionCollectionConstraints(obj ConnectionCollectionConnectionCollection) error {
	return nil
}
