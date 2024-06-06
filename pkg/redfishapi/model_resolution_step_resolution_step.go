/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// ResolutionStepResolutionStep - This type describes a recommended step of the service-defined resolution.
type ResolutionStepResolutionStep struct {

	// The parameters of the action URI for a resolution step.
	ActionParameters []ActionInfoParameters `json:"ActionParameters,omitempty"`

	// The action URI for a resolution step.
	ActionURI string `json:"ActionURI,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The priority in the set of resolution steps.
	Priority *int64 `json:"Priority,omitempty"`

	ResolutionType ResolutionStepV100ResolutionType `json:"ResolutionType"`

	// The number of retries for a resolution step.
	RetryCount *int64 `json:"RetryCount,omitempty"`

	// The interval between retries for a resolution step.
	RetryIntervalSeconds *int64 `json:"RetryIntervalSeconds,omitempty"`

	// The target URI of the component for a resolution step.
	TargetComponentURI *string `json:"TargetComponentURI,omitempty"`
}

// AssertResolutionStepResolutionStepRequired checks if the required fields are not zero-ed
func AssertResolutionStepResolutionStepRequired(obj ResolutionStepResolutionStep) error {
	elements := map[string]interface{}{
		"ResolutionType": obj.ResolutionType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ActionParameters {
		if err := AssertActionInfoParametersRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertResolutionStepResolutionStepConstraints checks if the values respects the defined constraints
func AssertResolutionStepResolutionStepConstraints(obj ResolutionStepResolutionStep) error {
	if obj.Priority != nil && *obj.Priority < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RetryCount != nil && *obj.RetryCount < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RetryIntervalSeconds != nil && *obj.RetryIntervalSeconds < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
