/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ComputerSystemV1220HostedServices - The services that might be running or installed on the system.
type ComputerSystemV1220HostedServices struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	StorageServices OdataV4IdRef `json:"StorageServices,omitempty"`
}

// AssertComputerSystemV1220HostedServicesRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220HostedServicesRequired(obj ComputerSystemV1220HostedServices) error {
	if err := AssertOdataV4IdRefRequired(obj.StorageServices); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1220HostedServicesConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220HostedServicesConstraints(obj ComputerSystemV1220HostedServices) error {
	return nil
}
