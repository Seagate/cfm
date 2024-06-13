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

// ComputerSystemV1220MemorySummary - The memory of the system in general detail.
type ComputerSystemV1220MemorySummary struct {
	MemoryMirroring ComputerSystemV1220MemoryMirroring `json:"MemoryMirroring,omitempty"`

	Metrics OdataV4IdRef `json:"Metrics,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The total configured operating system-accessible memory (RAM), measured in GiB.
	TotalSystemMemoryGiB *float32 `json:"TotalSystemMemoryGiB,omitempty"`

	// The total configured, system-accessible persistent memory, measured in GiB.
	TotalSystemPersistentMemoryGiB *float32 `json:"TotalSystemPersistentMemoryGiB,omitempty"`
}

// AssertComputerSystemV1220MemorySummaryRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220MemorySummaryRequired(obj ComputerSystemV1220MemorySummary) error {
	if err := AssertOdataV4IdRefRequired(obj.Metrics); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1220MemorySummaryConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220MemorySummaryConstraints(obj ComputerSystemV1220MemorySummary) error {
	if obj.TotalSystemMemoryGiB != nil && *obj.TotalSystemMemoryGiB < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TotalSystemPersistentMemoryGiB != nil && *obj.TotalSystemPersistentMemoryGiB < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
