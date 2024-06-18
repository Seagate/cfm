/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ActionInfoV142ParameterTypes string

// List of ActionInfoV142ParameterTypes
const (
	ACTIONINFOV142PARAMETERTYPES_BOOLEAN      ActionInfoV142ParameterTypes = "Boolean"
	ACTIONINFOV142PARAMETERTYPES_NUMBER       ActionInfoV142ParameterTypes = "Number"
	ACTIONINFOV142PARAMETERTYPES_NUMBER_ARRAY ActionInfoV142ParameterTypes = "NumberArray"
	ACTIONINFOV142PARAMETERTYPES_STRING       ActionInfoV142ParameterTypes = "String"
	ACTIONINFOV142PARAMETERTYPES_STRING_ARRAY ActionInfoV142ParameterTypes = "StringArray"
	ACTIONINFOV142PARAMETERTYPES_OBJECT       ActionInfoV142ParameterTypes = "Object"
	ACTIONINFOV142PARAMETERTYPES_OBJECT_ARRAY ActionInfoV142ParameterTypes = "ObjectArray"
)

// AssertActionInfoV142ParameterTypesRequired checks if the required fields are not zero-ed
func AssertActionInfoV142ParameterTypesRequired(obj ActionInfoV142ParameterTypes) error {
	return nil
}

// AssertActionInfoV142ParameterTypesConstraints checks if the values respects the defined constraints
func AssertActionInfoV142ParameterTypesConstraints(obj ActionInfoV142ParameterTypes) error {
	return nil
}