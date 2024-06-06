/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ChassisV1250ChassisType string

// List of ChassisV1250ChassisType
const (
	CHASSISV1250CHASSISTYPE_RACK              ChassisV1250ChassisType = "Rack"
	CHASSISV1250CHASSISTYPE_BLADE             ChassisV1250ChassisType = "Blade"
	CHASSISV1250CHASSISTYPE_ENCLOSURE         ChassisV1250ChassisType = "Enclosure"
	CHASSISV1250CHASSISTYPE_STAND_ALONE       ChassisV1250ChassisType = "StandAlone"
	CHASSISV1250CHASSISTYPE_RACK_MOUNT        ChassisV1250ChassisType = "RackMount"
	CHASSISV1250CHASSISTYPE_CARD              ChassisV1250ChassisType = "Card"
	CHASSISV1250CHASSISTYPE_CARTRIDGE         ChassisV1250ChassisType = "Cartridge"
	CHASSISV1250CHASSISTYPE_ROW               ChassisV1250ChassisType = "Row"
	CHASSISV1250CHASSISTYPE_POD               ChassisV1250ChassisType = "Pod"
	CHASSISV1250CHASSISTYPE_EXPANSION         ChassisV1250ChassisType = "Expansion"
	CHASSISV1250CHASSISTYPE_SIDECAR           ChassisV1250ChassisType = "Sidecar"
	CHASSISV1250CHASSISTYPE_ZONE              ChassisV1250ChassisType = "Zone"
	CHASSISV1250CHASSISTYPE_SLED              ChassisV1250ChassisType = "Sled"
	CHASSISV1250CHASSISTYPE_SHELF             ChassisV1250ChassisType = "Shelf"
	CHASSISV1250CHASSISTYPE_DRAWER            ChassisV1250ChassisType = "Drawer"
	CHASSISV1250CHASSISTYPE_MODULE            ChassisV1250ChassisType = "Module"
	CHASSISV1250CHASSISTYPE_COMPONENT         ChassisV1250ChassisType = "Component"
	CHASSISV1250CHASSISTYPE_IP_BASED_DRIVE    ChassisV1250ChassisType = "IPBasedDrive"
	CHASSISV1250CHASSISTYPE_RACK_GROUP        ChassisV1250ChassisType = "RackGroup"
	CHASSISV1250CHASSISTYPE_STORAGE_ENCLOSURE ChassisV1250ChassisType = "StorageEnclosure"
	CHASSISV1250CHASSISTYPE_IMMERSION_TANK    ChassisV1250ChassisType = "ImmersionTank"
	CHASSISV1250CHASSISTYPE_HEAT_EXCHANGER    ChassisV1250ChassisType = "HeatExchanger"
	CHASSISV1250CHASSISTYPE_POWER_STRIP       ChassisV1250ChassisType = "PowerStrip"
	CHASSISV1250CHASSISTYPE_OTHER             ChassisV1250ChassisType = "Other"
)

// AssertChassisV1250ChassisTypeRequired checks if the required fields are not zero-ed
func AssertChassisV1250ChassisTypeRequired(obj ChassisV1250ChassisType) error {
	return nil
}

// AssertChassisV1250ChassisTypeConstraints checks if the values respects the defined constraints
func AssertChassisV1250ChassisTypeConstraints(obj ChassisV1250ChassisType) error {
	return nil
}
