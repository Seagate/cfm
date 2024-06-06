/*
 * Composer and Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

type MemoryRegion struct {

	// The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters.
	Id string `json:"id"`

	// A response string
	Status string `json:"status,omitempty"`

	Type MemoryType `json:"type"`

	// A mebibyte equals 2**20 or 1,048,576 bytes.
	SizeMiB int32 `json:"sizeMiB"`

	// Memory bandwidth in the unit of GigaBytes per second
	Bandwidth int32 `json:"bandwidth,omitempty"`

	// Memory latency in the unit of nanosecond
	Latency int32 `json:"latency,omitempty"`

	// The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters.
	MemoryApplianceId string `json:"memoryApplianceId,omitempty"`

	// The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters.
	MemoryBladeId string `json:"memoryBladeId,omitempty"`

	// The CXL port name on the Memory Appliance
	MemoryAppliancePort string `json:"memoryAppliancePort,omitempty"`

	// The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters.
	MappedHostId string `json:"mappedHostId,omitempty"`

	// The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters.
	MappedHostPort string `json:"mappedHostPort,omitempty"`
}

// AssertMemoryRegionRequired checks if the required fields are not zero-ed
func AssertMemoryRegionRequired(obj MemoryRegion) error {
	elements := map[string]interface{}{
		"id":      obj.Id,
		"type":    obj.Type,
		"sizeMiB": obj.SizeMiB,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertMemoryRegionConstraints checks if the values respects the defined constraints
func AssertMemoryRegionConstraints(obj MemoryRegion) error {
	if obj.SizeMiB < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Bandwidth < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Latency < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}