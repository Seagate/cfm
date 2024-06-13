/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// ServiceRootV1161DeepOperations - The information about deep operations that the service supports.
type ServiceRootV1161DeepOperations struct {

	// An indication of whether the service supports the deep PATCH operation.
	DeepPATCH bool `json:"DeepPATCH,omitempty"`

	// An indication of whether the service supports the deep POST operation.
	DeepPOST bool `json:"DeepPOST,omitempty"`

	// The maximum levels of resources allowed in deep operations.
	MaxLevels int64 `json:"MaxLevels,omitempty"`
}

// AssertServiceRootV1161DeepOperationsRequired checks if the required fields are not zero-ed
func AssertServiceRootV1161DeepOperationsRequired(obj ServiceRootV1161DeepOperations) error {
	return nil
}

// AssertServiceRootV1161DeepOperationsConstraints checks if the values respects the defined constraints
func AssertServiceRootV1161DeepOperationsConstraints(obj ServiceRootV1161DeepOperations) error {
	if obj.MaxLevels < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
