/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ChassisV1250IntrusionSensorReArm string

// List of ChassisV1250IntrusionSensorReArm
const (
	CHASSISV1250INTRUSIONSENSORREARM_MANUAL    ChassisV1250IntrusionSensorReArm = "Manual"
	CHASSISV1250INTRUSIONSENSORREARM_AUTOMATIC ChassisV1250IntrusionSensorReArm = "Automatic"
)

// AssertChassisV1250IntrusionSensorReArmRequired checks if the required fields are not zero-ed
func AssertChassisV1250IntrusionSensorReArmRequired(obj ChassisV1250IntrusionSensorReArm) error {
	return nil
}

// AssertChassisV1250IntrusionSensorReArmConstraints checks if the values respects the defined constraints
func AssertChassisV1250IntrusionSensorReArmConstraints(obj ChassisV1250IntrusionSensorReArm) error {
	return nil
}
