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

// ComputerSystemV1220HostSerialConsole - The information about the serial console services that this system provides.
type ComputerSystemV1220HostSerialConsole struct {
	IPMI ComputerSystemV1220SerialConsoleProtocol `json:"IPMI,omitempty"`

	// The maximum number of service sessions, regardless of protocol, that this system can support.
	MaxConcurrentSessions int64 `json:"MaxConcurrentSessions,omitempty"`

	SSH ComputerSystemV1220SerialConsoleProtocol `json:"SSH,omitempty"`

	Telnet ComputerSystemV1220SerialConsoleProtocol `json:"Telnet,omitempty"`
}

// AssertComputerSystemV1220HostSerialConsoleRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220HostSerialConsoleRequired(obj ComputerSystemV1220HostSerialConsole) error {
	if err := AssertComputerSystemV1220SerialConsoleProtocolRequired(obj.IPMI); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220SerialConsoleProtocolRequired(obj.SSH); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220SerialConsoleProtocolRequired(obj.Telnet); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1220HostSerialConsoleConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220HostSerialConsoleConstraints(obj ComputerSystemV1220HostSerialConsole) error {
	if obj.MaxConcurrentSessions < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
