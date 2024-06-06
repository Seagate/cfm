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

// ComputerSystemV1220IdlePowerSaver - The idle power saver settings of a computer system.
type ComputerSystemV1220IdlePowerSaver struct {

	// An indication of whether idle power saver is enabled.
	Enabled bool `json:"Enabled,omitempty"`

	// The duration in seconds the computer system is below the EnterUtilizationPercent value before the idle power save is activated.
	EnterDwellTimeSeconds *int64 `json:"EnterDwellTimeSeconds,omitempty"`

	// The percentage of utilization when the computer system enters idle power save.  If the computer system's utilization goes below this value, it enters idle power save.
	EnterUtilizationPercent *float32 `json:"EnterUtilizationPercent,omitempty"`

	// The duration in seconds the computer system is above the ExitUtilizationPercent value before the idle power save is stopped.
	ExitDwellTimeSeconds *int64 `json:"ExitDwellTimeSeconds,omitempty"`

	// The percentage of utilization when the computer system exits idle power save.  If the computer system's utilization goes above this value, it exits idle power save.
	ExitUtilizationPercent *float32 `json:"ExitUtilizationPercent,omitempty"`
}

// AssertComputerSystemV1220IdlePowerSaverRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220IdlePowerSaverRequired(obj ComputerSystemV1220IdlePowerSaver) error {
	return nil
}

// AssertComputerSystemV1220IdlePowerSaverConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220IdlePowerSaverConstraints(obj ComputerSystemV1220IdlePowerSaver) error {
	if obj.EnterDwellTimeSeconds != nil && *obj.EnterDwellTimeSeconds < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EnterUtilizationPercent != nil && *obj.EnterUtilizationPercent < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ExitDwellTimeSeconds != nil && *obj.ExitDwellTimeSeconds < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ExitUtilizationPercent != nil && *obj.ExitUtilizationPercent < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
