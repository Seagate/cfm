/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ResourceV1190Orientation : The orientations for the ordering of the part location ordinal value.
type ResourceV1190Orientation string

// List of ResourceV1190Orientation
const (
	RESOURCEV1190ORIENTATION_FRONT_TO_BACK ResourceV1190Orientation = "FrontToBack"
	RESOURCEV1190ORIENTATION_BACK_TO_FRONT ResourceV1190Orientation = "BackToFront"
	RESOURCEV1190ORIENTATION_TOP_TO_BOTTOM ResourceV1190Orientation = "TopToBottom"
	RESOURCEV1190ORIENTATION_BOTTOM_TO_TOP ResourceV1190Orientation = "BottomToTop"
	RESOURCEV1190ORIENTATION_LEFT_TO_RIGHT ResourceV1190Orientation = "LeftToRight"
	RESOURCEV1190ORIENTATION_RIGHT_TO_LEFT ResourceV1190Orientation = "RightToLeft"
)

// AssertResourceV1190OrientationRequired checks if the required fields are not zero-ed
func AssertResourceV1190OrientationRequired(obj ResourceV1190Orientation) error {
	return nil
}

// AssertResourceV1190OrientationConstraints checks if the values respects the defined constraints
func AssertResourceV1190OrientationConstraints(obj ResourceV1190Orientation) error {
	return nil
}
