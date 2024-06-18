/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ChassisV1250IndicatorLed string

// List of ChassisV1250IndicatorLed
const (
	CHASSISV1250INDICATORLED_UNKNOWN  ChassisV1250IndicatorLed = "Unknown"
	CHASSISV1250INDICATORLED_LIT      ChassisV1250IndicatorLed = "Lit"
	CHASSISV1250INDICATORLED_BLINKING ChassisV1250IndicatorLed = "Blinking"
	CHASSISV1250INDICATORLED_OFF      ChassisV1250IndicatorLed = "Off"
)

// AssertChassisV1250IndicatorLedRequired checks if the required fields are not zero-ed
func AssertChassisV1250IndicatorLedRequired(obj ChassisV1250IndicatorLed) error {
	return nil
}

// AssertChassisV1250IndicatorLedConstraints checks if the values respects the defined constraints
func AssertChassisV1250IndicatorLedConstraints(obj ChassisV1250IndicatorLed) error {
	return nil
}
