/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ComputerSystemV1220InterfaceTypeSelection : The enumerations of InterfaceTypeSelection specify the method for switching the TrustedModule InterfaceType, for instance between TPM1_2 and TPM2_0, if supported.
type ComputerSystemV1220InterfaceTypeSelection string

// List of ComputerSystemV1220InterfaceTypeSelection
const (
	COMPUTERSYSTEMV1220INTERFACETYPESELECTION_NONE            ComputerSystemV1220InterfaceTypeSelection = "None"
	COMPUTERSYSTEMV1220INTERFACETYPESELECTION_FIRMWARE_UPDATE ComputerSystemV1220InterfaceTypeSelection = "FirmwareUpdate"
	COMPUTERSYSTEMV1220INTERFACETYPESELECTION_BIOS_SETTING    ComputerSystemV1220InterfaceTypeSelection = "BiosSetting"
	COMPUTERSYSTEMV1220INTERFACETYPESELECTION_OEM_METHOD      ComputerSystemV1220InterfaceTypeSelection = "OemMethod"
)

// AssertComputerSystemV1220InterfaceTypeSelectionRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220InterfaceTypeSelectionRequired(obj ComputerSystemV1220InterfaceTypeSelection) error {
	return nil
}

// AssertComputerSystemV1220InterfaceTypeSelectionConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220InterfaceTypeSelectionConstraints(obj ComputerSystemV1220InterfaceTypeSelection) error {
	return nil
}
