/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110ConnectedDeviceMode string

// List of PortV1110ConnectedDeviceMode
const (
	PORTV1110CONNECTEDDEVICEMODE_DISCONNECTED                    PortV1110ConnectedDeviceMode = "Disconnected"
	PORTV1110CONNECTEDDEVICEMODE_RCD                             PortV1110ConnectedDeviceMode = "RCD"
	PORTV1110CONNECTEDDEVICEMODE_CXL68_B_FLIT_AND_VH             PortV1110ConnectedDeviceMode = "CXL68BFlitAndVH"
	PORTV1110CONNECTEDDEVICEMODE_STANDARD256_B_FLIT              PortV1110ConnectedDeviceMode = "Standard256BFlit"
	PORTV1110CONNECTEDDEVICEMODE_CXL_LATENCY_OPTIMIZED256_B_FLIT PortV1110ConnectedDeviceMode = "CXLLatencyOptimized256BFlit"
	PORTV1110CONNECTEDDEVICEMODE_PBR                             PortV1110ConnectedDeviceMode = "PBR"
)

// AssertPortV1110ConnectedDeviceModeRequired checks if the required fields are not zero-ed
func AssertPortV1110ConnectedDeviceModeRequired(obj PortV1110ConnectedDeviceMode) error {
	return nil
}

// AssertPortV1110ConnectedDeviceModeConstraints checks if the values respects the defined constraints
func AssertPortV1110ConnectedDeviceModeConstraints(obj PortV1110ConnectedDeviceMode) error {
	return nil
}
