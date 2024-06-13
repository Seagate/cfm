/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryDomainV150MemoryDomain - The MemoryDomain schema describes a memory domain and its configuration.  Memory domains indicate to the client which memory, or DIMMs, can be grouped together in memory chunks to represent addressable memory.
type MemoryDomainV150MemoryDomain struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions MemoryDomainV150Actions `json:"Actions,omitempty"`

	// An indication of whether this memory domain supports the provisioning of blocks of memory.
	AllowsBlockProvisioning *bool `json:"AllowsBlockProvisioning,omitempty"`

	// An indication of whether this memory domain supports the creation of memory chunks.
	AllowsMemoryChunkCreation *bool `json:"AllowsMemoryChunkCreation,omitempty"`

	// An indication of whether this memory domain supports the creation of memory chunks with mirroring enabled.
	AllowsMirroring *bool `json:"AllowsMirroring,omitempty"`

	// An indication of whether this memory domain supports the creation of memory chunks with sparing enabled.
	AllowsSparing *bool `json:"AllowsSparing,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The interleave sets for the memory chunk.
	InterleavableMemorySets []MemoryDomainV150MemorySet `json:"InterleavableMemorySets,omitempty"`

	Links MemoryDomainV150Links `json:"Links,omitempty"`

	// The incremental size, from the minimum size, allowed for a memory chunk within this domain in mebibytes (MiB).
	MemoryChunkIncrementMiB *int64 `json:"MemoryChunkIncrementMiB,omitempty"`

	MemoryChunks OdataV4IdRef `json:"MemoryChunks,omitempty"`

	// The total size of the memory domain in mebibytes (MiB).
	MemorySizeMiB *int64 `json:"MemorySizeMiB,omitempty"`

	// The minimum size allowed for a memory chunk within this domain in mebibytes (MiB).
	MinMemoryChunkSizeMiB *int64 `json:"MinMemoryChunkSizeMiB,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`
}

// AssertMemoryDomainV150MemoryDomainRequired checks if the required fields are not zero-ed
func AssertMemoryDomainV150MemoryDomainRequired(obj MemoryDomainV150MemoryDomain) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertMemoryDomainV150ActionsRequired(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.InterleavableMemorySets {
		if err := AssertMemoryDomainV150MemorySetRequired(el); err != nil {
			return err
		}
	}
	if err := AssertMemoryDomainV150LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.MemoryChunks); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertMemoryDomainV150MemoryDomainConstraints checks if the values respects the defined constraints
func AssertMemoryDomainV150MemoryDomainConstraints(obj MemoryDomainV150MemoryDomain) error {
	return nil
}
