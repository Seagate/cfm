/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110PortConnectionType string

// List of PortV1110PortConnectionType
const (
	PORTV1110PORTCONNECTIONTYPE_NOT_CONNECTED   PortV1110PortConnectionType = "NotConnected"
	PORTV1110PORTCONNECTIONTYPE_N_PORT          PortV1110PortConnectionType = "NPort"
	PORTV1110PORTCONNECTIONTYPE_POINT_TO_POINT  PortV1110PortConnectionType = "PointToPoint"
	PORTV1110PORTCONNECTIONTYPE_PRIVATE_LOOP    PortV1110PortConnectionType = "PrivateLoop"
	PORTV1110PORTCONNECTIONTYPE_PUBLIC_LOOP     PortV1110PortConnectionType = "PublicLoop"
	PORTV1110PORTCONNECTIONTYPE_GENERIC         PortV1110PortConnectionType = "Generic"
	PORTV1110PORTCONNECTIONTYPE_EXTENDER_FABRIC PortV1110PortConnectionType = "ExtenderFabric"
	PORTV1110PORTCONNECTIONTYPE_F_PORT          PortV1110PortConnectionType = "FPort"
	PORTV1110PORTCONNECTIONTYPE_E_PORT          PortV1110PortConnectionType = "EPort"
	PORTV1110PORTCONNECTIONTYPE_TE_PORT         PortV1110PortConnectionType = "TEPort"
	PORTV1110PORTCONNECTIONTYPE_NP_PORT         PortV1110PortConnectionType = "NPPort"
	PORTV1110PORTCONNECTIONTYPE_G_PORT          PortV1110PortConnectionType = "GPort"
	PORTV1110PORTCONNECTIONTYPE_NL_PORT         PortV1110PortConnectionType = "NLPort"
	PORTV1110PORTCONNECTIONTYPE_FL_PORT         PortV1110PortConnectionType = "FLPort"
	PORTV1110PORTCONNECTIONTYPE_EX_PORT         PortV1110PortConnectionType = "EXPort"
	PORTV1110PORTCONNECTIONTYPE_U_PORT          PortV1110PortConnectionType = "UPort"
	PORTV1110PORTCONNECTIONTYPE_D_PORT          PortV1110PortConnectionType = "DPort"
)

// AssertPortV1110PortConnectionTypeRequired checks if the required fields are not zero-ed
func AssertPortV1110PortConnectionTypeRequired(obj PortV1110PortConnectionType) error {
	return nil
}

// AssertPortV1110PortConnectionTypeConstraints checks if the values respects the defined constraints
func AssertPortV1110PortConnectionTypeConstraints(obj PortV1110PortConnectionType) error {
	return nil
}
