/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ChassisV1250Doors - The doors or access panels of the chassis.
type ChassisV1250Doors struct {
	Front ChassisV1250Door `json:"Front,omitempty"`

	Rear ChassisV1250Door `json:"Rear,omitempty"`
}

// AssertChassisV1250DoorsRequired checks if the required fields are not zero-ed
func AssertChassisV1250DoorsRequired(obj ChassisV1250Doors) error {
	if err := AssertChassisV1250DoorRequired(obj.Front); err != nil {
		return err
	}
	if err := AssertChassisV1250DoorRequired(obj.Rear); err != nil {
		return err
	}
	return nil
}

// AssertChassisV1250DoorsConstraints checks if the values respects the defined constraints
func AssertChassisV1250DoorsConstraints(obj ChassisV1250Doors) error {
	return nil
}
