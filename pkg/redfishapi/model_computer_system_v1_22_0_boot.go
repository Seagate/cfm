/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// ComputerSystemV1220Boot - The boot information for this resource.
type ComputerSystemV1220Boot struct {

	// Ordered array of boot source aliases representing the persistent boot order associated with this computer system.
	AliasBootOrder []ComputerSystemBootSource `json:"AliasBootOrder,omitempty"`

	// The number of attempts the system will automatically retry booting.
	AutomaticRetryAttempts *int64 `json:"AutomaticRetryAttempts,omitempty"`

	AutomaticRetryConfig ComputerSystemV1220AutomaticRetryConfig `json:"AutomaticRetryConfig,omitempty"`

	// The BootOptionReference of the Boot Option to perform a one-time boot from when BootSourceOverrideTarget is `UefiBootNext`.
	BootNext *string `json:"BootNext,omitempty"`

	BootOptions OdataV4IdRef `json:"BootOptions,omitempty"`

	// An array of BootOptionReference strings that represent the persistent boot order for with this computer system.  Changes to the boot order typically require a system reset before they take effect.  It is likely that a client finds the `@Redfish.Settings` term in this resource, and if it is found, the client makes requests to change boot order settings by modifying the resource identified by the `@Redfish.Settings` term.
	BootOrder []*string `json:"BootOrder,omitempty"`

	BootOrderPropertySelection ComputerSystemV1220BootOrderTypes `json:"BootOrderPropertySelection,omitempty"`

	BootSourceOverrideEnabled ComputerSystemV1220BootSourceOverrideEnabled `json:"BootSourceOverrideEnabled,omitempty"`

	BootSourceOverrideMode ComputerSystemV1220BootSourceOverrideMode `json:"BootSourceOverrideMode,omitempty"`

	BootSourceOverrideTarget ComputerSystemBootSource `json:"BootSourceOverrideTarget,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	// The URI to boot from when BootSourceOverrideTarget is set to `UefiHttp`.
	HttpBootUri *string `json:"HttpBootUri,omitempty"`

	// The number of remaining automatic retry boots.
	RemainingAutomaticRetryAttempts *int64 `json:"RemainingAutomaticRetryAttempts,omitempty"`

	StopBootOnFault ComputerSystemV1220StopBootOnFault `json:"StopBootOnFault,omitempty"`

	TrustedModuleRequiredToBoot ComputerSystemV1220TrustedModuleRequiredToBoot `json:"TrustedModuleRequiredToBoot,omitempty"`

	// The UEFI device path of the device from which to boot when BootSourceOverrideTarget is `UefiTarget`.
	UefiTargetBootSourceOverride *string `json:"UefiTargetBootSourceOverride,omitempty"`
}

// AssertComputerSystemV1220BootRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220BootRequired(obj ComputerSystemV1220Boot) error {
	if err := AssertOdataV4IdRefRequired(obj.BootOptions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Certificates); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1220BootConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220BootConstraints(obj ComputerSystemV1220Boot) error {
	if obj.AutomaticRetryAttempts != nil && *obj.AutomaticRetryAttempts < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RemainingAutomaticRetryAttempts != nil && *obj.RemainingAutomaticRetryAttempts < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}