/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type SessionV171SessionTypes string

// List of SessionV171SessionTypes
const (
	SESSIONV171SESSIONTYPES_HOST_CONSOLE        SessionV171SessionTypes = "HostConsole"
	SESSIONV171SESSIONTYPES_MANAGER_CONSOLE     SessionV171SessionTypes = "ManagerConsole"
	SESSIONV171SESSIONTYPES_IPMI                SessionV171SessionTypes = "IPMI"
	SESSIONV171SESSIONTYPES_KVMIP               SessionV171SessionTypes = "KVMIP"
	SESSIONV171SESSIONTYPES_OEM                 SessionV171SessionTypes = "OEM"
	SESSIONV171SESSIONTYPES_REDFISH             SessionV171SessionTypes = "Redfish"
	SESSIONV171SESSIONTYPES_VIRTUAL_MEDIA       SessionV171SessionTypes = "VirtualMedia"
	SESSIONV171SESSIONTYPES_WEB_UI              SessionV171SessionTypes = "WebUI"
	SESSIONV171SESSIONTYPES_OUTBOUND_CONNECTION SessionV171SessionTypes = "OutboundConnection"
)

// AssertSessionV171SessionTypesRequired checks if the required fields are not zero-ed
func AssertSessionV171SessionTypesRequired(obj SessionV171SessionTypes) error {
	return nil
}

// AssertSessionV171SessionTypesConstraints checks if the values respects the defined constraints
func AssertSessionV171SessionTypesConstraints(obj SessionV171SessionTypes) error {
	return nil
}
