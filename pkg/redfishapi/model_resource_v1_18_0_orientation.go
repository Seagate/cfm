/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceV1180Orientation : The orientations for the ordering of the part location ordinal value.
type ResourceV1180Orientation string

// List of ResourceV1180Orientation
const (
	RESOURCEV1180ORIENTATION_FRONT_TO_BACK ResourceV1180Orientation = "FrontToBack"
	RESOURCEV1180ORIENTATION_BACK_TO_FRONT ResourceV1180Orientation = "BackToFront"
	RESOURCEV1180ORIENTATION_TOP_TO_BOTTOM ResourceV1180Orientation = "TopToBottom"
	RESOURCEV1180ORIENTATION_BOTTOM_TO_TOP ResourceV1180Orientation = "BottomToTop"
	RESOURCEV1180ORIENTATION_LEFT_TO_RIGHT ResourceV1180Orientation = "LeftToRight"
	RESOURCEV1180ORIENTATION_RIGHT_TO_LEFT ResourceV1180Orientation = "RightToLeft"
)

// AssertResourceV1180OrientationRequired checks if the required fields are not zero-ed
func AssertResourceV1180OrientationRequired(obj ResourceV1180Orientation) error {
	return nil
}

// AssertResourceV1180OrientationConstraints checks if the values respects the defined constraints
func AssertResourceV1180OrientationConstraints(obj ResourceV1180Orientation) error {
	return nil
}
