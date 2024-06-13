/*
 * Composable Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Qos int32

// List of Qos
const (
	QOS__1 Qos = 1
	QOS__2 Qos = 2
	QOS__4 Qos = 4
	QOS__8 Qos = 8
)

// AssertQosRequired checks if the required fields are not zero-ed
func AssertQosRequired(obj Qos) error {
	return nil
}

// AssertQosConstraints checks if the values respects the defined constraints
func AssertQosConstraints(obj Qos) error {
	return nil
}
