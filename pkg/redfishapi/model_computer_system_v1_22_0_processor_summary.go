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

// ComputerSystemV1220ProcessorSummary - The central processors of the system in general detail.
type ComputerSystemV1220ProcessorSummary struct {

	// The number of processor cores in the system.
	CoreCount *int64 `json:"CoreCount,omitempty"`

	// The number of physical processors in the system.
	Count *int64 `json:"Count,omitempty"`

	// The number of logical processors in the system.
	LogicalProcessorCount *int64 `json:"LogicalProcessorCount,omitempty"`

	Metrics OdataV4IdRef `json:"Metrics,omitempty"`

	// The processor model for the primary or majority of processors in this system.
	Model *string `json:"Model,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// An indication of whether threading is enabled on all processors in this system.
	ThreadingEnabled bool `json:"ThreadingEnabled,omitempty"`
}

// AssertComputerSystemV1220ProcessorSummaryRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220ProcessorSummaryRequired(obj ComputerSystemV1220ProcessorSummary) error {
	if err := AssertOdataV4IdRefRequired(obj.Metrics); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1220ProcessorSummaryConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220ProcessorSummaryConstraints(obj ComputerSystemV1220ProcessorSummary) error {
	if obj.CoreCount != nil && *obj.CoreCount < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Count != nil && *obj.Count < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.LogicalProcessorCount != nil && *obj.LogicalProcessorCount < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
