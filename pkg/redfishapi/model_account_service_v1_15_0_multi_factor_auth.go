/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// AccountServiceV1150MultiFactorAuth - Multi-factor authentication settings.
type AccountServiceV1150MultiFactorAuth struct {
	ClientCertificate AccountServiceV1150ClientCertificate `json:"ClientCertificate,omitempty"`

	GoogleAuthenticator AccountServiceV1150GoogleAuthenticator `json:"GoogleAuthenticator,omitempty"`

	MicrosoftAuthenticator AccountServiceV1150MicrosoftAuthenticator `json:"MicrosoftAuthenticator,omitempty"`

	OneTimePasscode AccountServiceV1150OneTimePasscode `json:"OneTimePasscode,omitempty"`

	SecurID AccountServiceV1150SecurId `json:"SecurID,omitempty"`
}

// AssertAccountServiceV1150MultiFactorAuthRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150MultiFactorAuthRequired(obj AccountServiceV1150MultiFactorAuth) error {
	if err := AssertAccountServiceV1150ClientCertificateRequired(obj.ClientCertificate); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150GoogleAuthenticatorRequired(obj.GoogleAuthenticator); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150MicrosoftAuthenticatorRequired(obj.MicrosoftAuthenticator); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150OneTimePasscodeRequired(obj.OneTimePasscode); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150SecurIdRequired(obj.SecurID); err != nil {
		return err
	}
	return nil
}

// AssertAccountServiceV1150MultiFactorAuthConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150MultiFactorAuthConstraints(obj AccountServiceV1150MultiFactorAuth) error {
	return nil
}