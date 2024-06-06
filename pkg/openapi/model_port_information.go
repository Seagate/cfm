/*
 * Composer and Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PortInformation struct {

	// The id of this resource
	Id string `json:"id"`

	// The global CXL identifier
	GCxlId string `json:"gCxlId,omitempty"`

	// A full path to the resource with id as the last component
	LinkedPortUri string `json:"linkedPortUri,omitempty"`

	// The protocol being sent over this port
	PortProtocol string `json:"portProtocol,omitempty"`

	// The physical connection medium for this port
	PortMedium string `json:"portMedium,omitempty"`

	// The current speed of this port
	CurrentSpeedGbps int32 `json:"currentSpeedGbps,omitempty"`

	// The health of the resource
	StatusHealth string `json:"statusHealth"`

	// The state of the resource
	StatusState string `json:"statusState"`

	// The number of lanes, phys, or other physical transport links that this port contains
	Width int32 `json:"width,omitempty"`

	// Status of the link, such as LinkUp or LinkDown
	LinkStatus string `json:"linkStatus,omitempty"`
}

// AssertPortInformationRequired checks if the required fields are not zero-ed
func AssertPortInformationRequired(obj PortInformation) error {
	elements := map[string]interface{}{
		"id":           obj.Id,
		"statusHealth": obj.StatusHealth,
		"statusState":  obj.StatusState,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPortInformationConstraints checks if the values respects the defined constraints
func AssertPortInformationConstraints(obj PortInformation) error {
	return nil
}