/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MessageV121Message - The message that the Redfish service returns.
type MessageV121Message struct {

	// The human-readable message.
	Message string `json:"Message,omitempty"`

	// An array of message arguments that are substituted for the arguments in the message when looked up in the message registry.
	MessageArgs []string `json:"MessageArgs,omitempty"`

	// The identifier for the message.
	MessageId string `json:"MessageId"`

	MessageSeverity ResourceHealth `json:"MessageSeverity,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A set of properties described by the message.
	RelatedProperties []string `json:"RelatedProperties,omitempty"`

	// Used to provide suggestions on how to resolve the situation that caused the message.
	Resolution string `json:"Resolution,omitempty"`

	// The list of recommended steps to resolve the situation that caused the message.
	ResolutionSteps []ResolutionStepResolutionStep `json:"ResolutionSteps,omitempty"`

	// The severity of the message.
	// Deprecated
	Severity string `json:"Severity,omitempty"`
}

// AssertMessageV121MessageRequired checks if the required fields are not zero-ed
func AssertMessageV121MessageRequired(obj MessageV121Message) error {
	elements := map[string]interface{}{
		"MessageId": obj.MessageId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ResolutionSteps {
		if err := AssertResolutionStepResolutionStepRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMessageV121MessageConstraints checks if the values respects the defined constraints
func AssertMessageV121MessageConstraints(obj MessageV121Message) error {
	return nil
}
