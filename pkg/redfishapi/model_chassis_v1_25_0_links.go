/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ChassisV1250Links - The links to other resources that are related to this resource.
type ChassisV1250Links struct {

	// An array of links to the cables connected to this chassis.
	Cables []OdataV4IdRef `json:"Cables,omitempty"`

	// The number of items in a collection.
	CablesodataCount int64 `json:"Cables@odata.count,omitempty"`

	// An array of links to the computer systems that this chassis directly and wholly contains.
	ComputerSystems []OdataV4IdRef `json:"ComputerSystems,omitempty"`

	// The number of items in a collection.
	ComputerSystemsodataCount int64 `json:"ComputerSystems@odata.count,omitempty"`

	// An array of links to cooling loops connected to this chassis.
	ConnectedCoolingLoops []OdataV4IdRef `json:"ConnectedCoolingLoops,omitempty"`

	// The number of items in a collection.
	ConnectedCoolingLoopsodataCount int64 `json:"ConnectedCoolingLoops@odata.count,omitempty"`

	ContainedBy OdataV4IdRef `json:"ContainedBy,omitempty"`

	// An array of links to any other chassis that this chassis has in it.
	Contains []OdataV4IdRef `json:"Contains,omitempty"`

	// The number of items in a collection.
	ContainsodataCount int64 `json:"Contains@odata.count,omitempty"`

	// An array of links to resources or objects that cool this chassis.  Normally, the link is for either a chassis or a specific set of fans.
	// Deprecated
	CooledBy []OdataV4IdRef `json:"CooledBy,omitempty"`

	// The number of items in a collection.
	CooledByodataCount int64 `json:"CooledBy@odata.count,omitempty"`

	// An array of links to cooling unit functionality contained in this chassis.
	CoolingUnits []OdataV4IdRef `json:"CoolingUnits,omitempty"`

	// The number of items in a collection.
	CoolingUnitsodataCount int64 `json:"CoolingUnits@odata.count,omitempty"`

	// An array of links to the drives located in this chassis.
	Drives []OdataV4IdRef `json:"Drives,omitempty"`

	// The number of items in a collection.
	DrivesodataCount int64 `json:"Drives@odata.count,omitempty"`

	Facility OdataV4IdRef `json:"Facility,omitempty"`

	// An array of links to the fans that cool this chassis.
	Fans []OdataV4IdRef `json:"Fans,omitempty"`

	// The number of items in a collection.
	FansodataCount int64 `json:"Fans@odata.count,omitempty"`

	// An array of links to the managers responsible for managing this chassis.
	ManagedBy []OdataV4IdRef `json:"ManagedBy,omitempty"`

	// The number of items in a collection.
	ManagedByodataCount int64 `json:"ManagedBy@odata.count,omitempty"`

	// An array of links to the managers located in this chassis.
	ManagersInChassis []OdataV4IdRef `json:"ManagersInChassis,omitempty"`

	// The number of items in a collection.
	ManagersInChassisodataCount int64 `json:"ManagersInChassis@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe devices located in this chassis.
	// Deprecated
	PCIeDevices []OdataV4IdRef `json:"PCIeDevices,omitempty"`

	// The number of items in a collection.
	PCIeDevicesodataCount int64 `json:"PCIeDevices@odata.count,omitempty"`

	PowerDistribution OdataV4IdRef `json:"PowerDistribution,omitempty"`

	// An array of links to the outlets that provide power to this chassis.
	PowerOutlets []OdataV4IdRef `json:"PowerOutlets,omitempty"`

	// The number of items in a collection.
	PowerOutletsodataCount int64 `json:"PowerOutlets@odata.count,omitempty"`

	// An array of links to the power supplies that provide power to this chassis.
	PowerSupplies []OdataV4IdRef `json:"PowerSupplies,omitempty"`

	// The number of items in a collection.
	PowerSuppliesodataCount int64 `json:"PowerSupplies@odata.count,omitempty"`

	// An array of links to resources or objects that power this chassis.  Normally, the link is for either a chassis or a specific set of power supplies.
	// Deprecated
	PoweredBy []OdataV4IdRef `json:"PoweredBy,omitempty"`

	// The number of items in a collection.
	PoweredByodataCount int64 `json:"PoweredBy@odata.count,omitempty"`

	// An array of links to the processors located in this chassis.
	Processors []OdataV4IdRef `json:"Processors,omitempty"`

	// The number of items in a collection.
	ProcessorsodataCount int64 `json:"Processors@odata.count,omitempty"`

	// An array of links to the resource blocks located in this chassis.
	ResourceBlocks []OdataV4IdRef `json:"ResourceBlocks,omitempty"`

	// The number of items in a collection.
	ResourceBlocksodataCount int64 `json:"ResourceBlocks@odata.count,omitempty"`

	// An array of links to the storage subsystems connected to or inside this chassis.
	Storage []OdataV4IdRef `json:"Storage,omitempty"`

	// The number of items in a collection.
	StorageodataCount int64 `json:"Storage@odata.count,omitempty"`

	// An array of links to the switches located in this chassis.
	Switches []OdataV4IdRef `json:"Switches,omitempty"`

	// The number of items in a collection.
	SwitchesodataCount int64 `json:"Switches@odata.count,omitempty"`
}

// AssertChassisV1250LinksRequired checks if the required fields are not zero-ed
func AssertChassisV1250LinksRequired(obj ChassisV1250Links) error {
	for _, el := range obj.Cables {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ComputerSystems {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnectedCoolingLoops {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.ContainedBy); err != nil {
		return err
	}
	for _, el := range obj.Contains {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CooledBy {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CoolingUnits {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Drives {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.Facility); err != nil {
		return err
	}
	for _, el := range obj.Fans {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ManagedBy {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ManagersInChassis {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PCIeDevices {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.PowerDistribution); err != nil {
		return err
	}
	for _, el := range obj.PowerOutlets {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PowerSupplies {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PoweredBy {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Processors {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ResourceBlocks {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Storage {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Switches {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertChassisV1250LinksConstraints checks if the values respects the defined constraints
func AssertChassisV1250LinksConstraints(obj ChassisV1250Links) error {
	return nil
}