/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ComputerSystemV1220WatchdogWarningActions : The enumerations of WatchdogWarningActions specify the choice of action to take when the host watchdog timer is close (typically 3-10 seconds) to reaching its timeout value.
type ComputerSystemV1220WatchdogWarningActions string

// List of ComputerSystemV1220WatchdogWarningActions
const (
	COMPUTERSYSTEMV1220WATCHDOGWARNINGACTIONS_NONE                 ComputerSystemV1220WatchdogWarningActions = "None"
	COMPUTERSYSTEMV1220WATCHDOGWARNINGACTIONS_DIAGNOSTIC_INTERRUPT ComputerSystemV1220WatchdogWarningActions = "DiagnosticInterrupt"
	COMPUTERSYSTEMV1220WATCHDOGWARNINGACTIONS_SMI                  ComputerSystemV1220WatchdogWarningActions = "SMI"
	COMPUTERSYSTEMV1220WATCHDOGWARNINGACTIONS_MESSAGING_INTERRUPT  ComputerSystemV1220WatchdogWarningActions = "MessagingInterrupt"
	COMPUTERSYSTEMV1220WATCHDOGWARNINGACTIONS_SCI                  ComputerSystemV1220WatchdogWarningActions = "SCI"
	COMPUTERSYSTEMV1220WATCHDOGWARNINGACTIONS_OEM                  ComputerSystemV1220WatchdogWarningActions = "OEM"
)

// AssertComputerSystemV1220WatchdogWarningActionsRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220WatchdogWarningActionsRequired(obj ComputerSystemV1220WatchdogWarningActions) error {
	return nil
}

// AssertComputerSystemV1220WatchdogWarningActionsConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220WatchdogWarningActionsConstraints(obj ComputerSystemV1220WatchdogWarningActions) error {
	return nil
}
