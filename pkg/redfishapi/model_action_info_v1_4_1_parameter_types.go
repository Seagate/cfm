/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ActionInfoV141ParameterTypes string

// List of ActionInfoV141ParameterTypes
const (
	ACTIONINFOV141PARAMETERTYPES_BOOLEAN      ActionInfoV141ParameterTypes = "Boolean"
	ACTIONINFOV141PARAMETERTYPES_NUMBER       ActionInfoV141ParameterTypes = "Number"
	ACTIONINFOV141PARAMETERTYPES_NUMBER_ARRAY ActionInfoV141ParameterTypes = "NumberArray"
	ACTIONINFOV141PARAMETERTYPES_STRING       ActionInfoV141ParameterTypes = "String"
	ACTIONINFOV141PARAMETERTYPES_STRING_ARRAY ActionInfoV141ParameterTypes = "StringArray"
	ACTIONINFOV141PARAMETERTYPES_OBJECT       ActionInfoV141ParameterTypes = "Object"
	ACTIONINFOV141PARAMETERTYPES_OBJECT_ARRAY ActionInfoV141ParameterTypes = "ObjectArray"
)

// AssertActionInfoV141ParameterTypesRequired checks if the required fields are not zero-ed
func AssertActionInfoV141ParameterTypesRequired(obj ActionInfoV141ParameterTypes) error {
	return nil
}

// AssertActionInfoV141ParameterTypesConstraints checks if the values respects the defined constraints
func AssertActionInfoV141ParameterTypesConstraints(obj ActionInfoV141ParameterTypes) error {
	return nil
}