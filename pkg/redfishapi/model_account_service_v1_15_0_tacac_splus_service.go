/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// AccountServiceV1150TacacSplusService - Various settings to parse a TACACS+ service.
type AccountServiceV1150TacacSplusService struct {

	// The TACACS+ service authorization argument.
	AuthorizationService string `json:"AuthorizationService,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Indicates the allowed TACACS+ password exchange protocols.
	PasswordExchangeProtocols []AccountServiceV1150TacacSplusPasswordExchangeProtocol `json:"PasswordExchangeProtocols,omitempty"`

	// Indicates the name of the TACACS+ argument name in an authorization request.
	PrivilegeLevelArgument *string `json:"PrivilegeLevelArgument,omitempty"`
}

// AssertAccountServiceV1150TacacSplusServiceRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150TacacSplusServiceRequired(obj AccountServiceV1150TacacSplusService) error {
	return nil
}

// AssertAccountServiceV1150TacacSplusServiceConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150TacacSplusServiceConstraints(obj AccountServiceV1150TacacSplusService) error {
	return nil
}