/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ChassisV1250IntrusionSensor string

// List of ChassisV1250IntrusionSensor
const (
	CHASSISV1250INTRUSIONSENSOR_NORMAL             ChassisV1250IntrusionSensor = "Normal"
	CHASSISV1250INTRUSIONSENSOR_HARDWARE_INTRUSION ChassisV1250IntrusionSensor = "HardwareIntrusion"
	CHASSISV1250INTRUSIONSENSOR_TAMPERING_DETECTED ChassisV1250IntrusionSensor = "TamperingDetected"
)

// AssertChassisV1250IntrusionSensorRequired checks if the required fields are not zero-ed
func AssertChassisV1250IntrusionSensorRequired(obj ChassisV1250IntrusionSensor) error {
	return nil
}

// AssertChassisV1250IntrusionSensorConstraints checks if the values respects the defined constraints
func AssertChassisV1250IntrusionSensorConstraints(obj ChassisV1250IntrusionSensor) error {
	return nil
}
