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

// ServiceRootV1161Expand - The information about the use of `$expand` in the service.
type ServiceRootV1161Expand struct {

	// An indication of whether the service supports the asterisk (`*`) option of the `$expand` query parameter.
	ExpandAll bool `json:"ExpandAll,omitempty"`

	// An indication of whether the service supports the `$levels` option of the `$expand` query parameter.
	Levels bool `json:"Levels,omitempty"`

	// An indication of whether this service supports the tilde (`~`) option of the `$expand` query parameter.
	Links bool `json:"Links,omitempty"`

	// The maximum `$levels` option value in the `$expand` query parameter.
	MaxLevels int64 `json:"MaxLevels,omitempty"`

	// An indication of whether the service supports the period (`.`) option of the `$expand` query parameter.
	NoLinks bool `json:"NoLinks,omitempty"`
}

// AssertServiceRootV1161ExpandRequired checks if the required fields are not zero-ed
func AssertServiceRootV1161ExpandRequired(obj ServiceRootV1161Expand) error {
	return nil
}

// AssertServiceRootV1161ExpandConstraints checks if the values respects the defined constraints
func AssertServiceRootV1161ExpandConstraints(obj ServiceRootV1161Expand) error {
	if obj.MaxLevels < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
