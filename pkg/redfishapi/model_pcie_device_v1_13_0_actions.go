/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PcieDeviceV1130Actions - The available actions for this resource.
type PcieDeviceV1130Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertPcieDeviceV1130ActionsRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130ActionsRequired(obj PcieDeviceV1130Actions) error {
	return nil
}

// AssertPcieDeviceV1130ActionsConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130ActionsConstraints(obj PcieDeviceV1130Actions) error {
	return nil
}
