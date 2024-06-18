/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type RedundancyV142RedundancyMode string

// List of RedundancyV142RedundancyMode
const (
	REDUNDANCYV142REDUNDANCYMODE_FAILOVER      RedundancyV142RedundancyMode = "Failover"
	REDUNDANCYV142REDUNDANCYMODE_NM            RedundancyV142RedundancyMode = "N+m"
	REDUNDANCYV142REDUNDANCYMODE_SHARING       RedundancyV142RedundancyMode = "Sharing"
	REDUNDANCYV142REDUNDANCYMODE_SPARING       RedundancyV142RedundancyMode = "Sparing"
	REDUNDANCYV142REDUNDANCYMODE_NOT_REDUNDANT RedundancyV142RedundancyMode = "NotRedundant"
)

// AssertRedundancyV142RedundancyModeRequired checks if the required fields are not zero-ed
func AssertRedundancyV142RedundancyModeRequired(obj RedundancyV142RedundancyMode) error {
	return nil
}

// AssertRedundancyV142RedundancyModeConstraints checks if the values respects the defined constraints
func AssertRedundancyV142RedundancyModeConstraints(obj RedundancyV142RedundancyMode) error {
	return nil
}
