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

type ComposeMemoryRequest struct {

	// The CXL port name on the Memory Appliance
	Port string `json:"Port,omitempty"`

	// A mebibyte equals 2**20 or 1,048,576 bytes.
	MemorySizeMiB int32 `json:"memorySizeMiB"`

	QoS Qos `json:"QoS"`
}

// AssertComposeMemoryRequestRequired checks if the required fields are not zero-ed
func AssertComposeMemoryRequestRequired(obj ComposeMemoryRequest) error {
	elements := map[string]interface{}{
		"memorySizeMiB": obj.MemorySizeMiB,
		"QoS":           obj.QoS,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertComposeMemoryRequestConstraints checks if the values respects the defined constraints
func AssertComposeMemoryRequestConstraints(obj ComposeMemoryRequest) error {
	if obj.MemorySizeMiB < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}