/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ActionInfoV142Parameters - The information about a parameter included in a Redfish action for this resource.
type ActionInfoV142Parameters struct {

	// The allowable numeric values or duration values, inclusive ranges of values, and incremental step values for this parameter as applied to this action target.
	AllowableNumbers []*string `json:"AllowableNumbers,omitempty"`

	// The allowable pattern for this parameter as applied to this action target.
	AllowablePattern *string `json:"AllowablePattern,omitempty"`

	// Descriptions of allowable values for this parameter.
	AllowableValueDescriptions []*string `json:"AllowableValueDescriptions,omitempty"`

	// The allowable values for this parameter as applied to this action target.
	AllowableValues []*string `json:"AllowableValues,omitempty"`

	// The maximum number of array elements allowed for this parameter.
	ArraySizeMaximum *int64 `json:"ArraySizeMaximum,omitempty"`

	// The minimum number of array elements required for this parameter.
	ArraySizeMinimum *int64 `json:"ArraySizeMinimum,omitempty"`

	DataType *ActionInfoV142ParameterTypes `json:"DataType,omitempty"`

	// The maximum supported value for this parameter.
	MaximumValue *float32 `json:"MaximumValue,omitempty"`

	// The minimum supported value for this parameter.
	MinimumValue *float32 `json:"MinimumValue,omitempty"`

	// The name of the parameter for this action.
	Name string `json:"Name"`

	// The data type of an object-based parameter.
	ObjectDataType *string `json:"ObjectDataType,omitempty"`

	// An indication of whether the parameter is required to complete this action.
	Required bool `json:"Required,omitempty"`
}

// AssertActionInfoV142ParametersRequired checks if the required fields are not zero-ed
func AssertActionInfoV142ParametersRequired(obj ActionInfoV142Parameters) error {
	elements := map[string]interface{}{
		"Name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertActionInfoV142ParametersConstraints checks if the values respects the defined constraints
func AssertActionInfoV142ParametersConstraints(obj ActionInfoV142Parameters) error {
	return nil
}