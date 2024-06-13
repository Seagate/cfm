/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// SoftwareInventoryMeasurementBlock - The DSP0274-defined measurement block information.
type SoftwareInventoryMeasurementBlock struct {

	// The hexadecimal string representation of the numeric value of the DSP0274-defined Measurement field of the measurement block.
	Measurement *string `json:"Measurement,omitempty"`

	// The DSP0274-defined Index field of the measurement block.
	MeasurementIndex *int64 `json:"MeasurementIndex,omitempty"`

	// The DSP0274-defined MeasurementSize field of the measurement block.
	MeasurementSize *int64 `json:"MeasurementSize,omitempty"`

	// The DSP0274-defined MeasurementSpecification field of the measurement block.
	MeasurementSpecification *int64 `json:"MeasurementSpecification,omitempty"`
}

// AssertSoftwareInventoryMeasurementBlockRequired checks if the required fields are not zero-ed
func AssertSoftwareInventoryMeasurementBlockRequired(obj SoftwareInventoryMeasurementBlock) error {
	return nil
}

// AssertSoftwareInventoryMeasurementBlockConstraints checks if the values respects the defined constraints
func AssertSoftwareInventoryMeasurementBlockConstraints(obj SoftwareInventoryMeasurementBlock) error {
	return nil
}
