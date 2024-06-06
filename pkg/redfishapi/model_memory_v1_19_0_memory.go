/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryV1190Memory - The Memory schema represents a memory device, such as a DIMM, and its configuration.  It also describes the location, such as a slot, socket, or bay, where a unit can be installed, by populating a resource instance with an absent state if a unit is not present.
type MemoryV1190Memory struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions MemoryV1190Actions `json:"Actions,omitempty"`

	// The boundary that memory regions are allocated on, measured in mebibytes (MiB).
	AllocationAlignmentMiB *int64 `json:"AllocationAlignmentMiB,omitempty"`

	// The size of the smallest unit of allocation for a memory region in mebibytes (MiB).
	AllocationIncrementMiB *int64 `json:"AllocationIncrementMiB,omitempty"`

	// Speeds supported by this memory device.
	AllowedSpeedsMHz []int64 `json:"AllowedSpeedsMHz,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	BaseModuleType MemoryV1190BaseModuleType `json:"BaseModuleType,omitempty"`

	// The bus width, in bits.
	BusWidthBits *int64 `json:"BusWidthBits,omitempty"`

	CXL MemoryV1190Cxl `json:"CXL,omitempty"`

	// Total size of the cache portion memory in MiB.
	CacheSizeMiB *int64 `json:"CacheSizeMiB,omitempty"`

	// Memory capacity in mebibytes (MiB).
	CapacityMiB *int64 `json:"CapacityMiB,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	// An indication of whether the configuration of this memory device is locked and cannot be altered.
	ConfigurationLocked *bool `json:"ConfigurationLocked,omitempty"`

	// Data width in bits.
	DataWidthBits *int64 `json:"DataWidthBits,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Device ID.
	// Deprecated
	DeviceID *string `json:"DeviceID,omitempty"`

	// Location of the memory device in the platform.
	// Deprecated
	DeviceLocator *string `json:"DeviceLocator,omitempty"`

	// An indication of whether this memory is enabled.
	Enabled bool `json:"Enabled,omitempty"`

	EnvironmentMetrics OdataV4IdRef `json:"EnvironmentMetrics,omitempty"`

	ErrorCorrection MemoryV1190ErrorCorrection `json:"ErrorCorrection,omitempty"`

	// Version of API supported by the firmware.
	FirmwareApiVersion *string `json:"FirmwareApiVersion,omitempty"`

	// Revision of firmware on the memory controller.
	FirmwareRevision *string `json:"FirmwareRevision,omitempty"`

	// Function classes by the memory device.
	// Deprecated
	FunctionClasses []string `json:"FunctionClasses,omitempty"`

	HealthData MemoryV1190HealthData `json:"HealthData,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether rank spare is enabled for this memory device.
	IsRankSpareEnabled *bool `json:"IsRankSpareEnabled,omitempty"`

	// An indication of whether a spare device is enabled for this memory device.
	IsSpareDeviceEnabled *bool `json:"IsSpareDeviceEnabled,omitempty"`

	Links MemoryV1190Links `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	Log OdataV4IdRef `json:"Log,omitempty"`

	// Total size of the logical memory in MiB.
	LogicalSizeMiB *int64 `json:"LogicalSizeMiB,omitempty"`

	// The memory device manufacturer.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// Set of maximum power budgets supported by the memory device in milliwatt units.
	MaxTDPMilliWatts []int64 `json:"MaxTDPMilliWatts,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	// Deprecated
	Measurements []SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`

	MemoryDeviceType MemoryV1190MemoryDeviceType `json:"MemoryDeviceType,omitempty"`

	MemoryLocation MemoryV1190MemoryLocation `json:"MemoryLocation,omitempty"`

	// Media of this memory device.
	MemoryMedia []MemoryV1190MemoryMedia `json:"MemoryMedia,omitempty"`

	// The manufacturer ID of the memory subsystem controller of this memory device.
	MemorySubsystemControllerManufacturerID *string `json:"MemorySubsystemControllerManufacturerID,omitempty"`

	// The product ID of the memory subsystem controller of this memory device.
	MemorySubsystemControllerProductID *string `json:"MemorySubsystemControllerProductID,omitempty"`

	MemoryType MemoryV1190MemoryType `json:"MemoryType,omitempty"`

	Metrics OdataV4IdRef `json:"Metrics,omitempty"`

	// The product model number of this device.
	Model *string `json:"Model,omitempty"`

	// The manufacturer ID of this memory device.
	ModuleManufacturerID *string `json:"ModuleManufacturerID,omitempty"`

	// The product ID of this memory device.
	ModuleProductID *string `json:"ModuleProductID,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The total non-volatile memory capacity in mebibytes (MiB).
	NonVolatileSizeLimitMiB int64 `json:"NonVolatileSizeLimitMiB,omitempty"`

	// Total size of the non-volatile portion memory in MiB.
	NonVolatileSizeMiB *int64 `json:"NonVolatileSizeMiB,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Memory modes supported by the memory device.
	OperatingMemoryModes []MemoryV1190OperatingMemoryModes `json:"OperatingMemoryModes,omitempty"`

	// Operating speed of the memory device in MHz or MT/s as appropriate.
	OperatingSpeedMhz *int64 `json:"OperatingSpeedMhz,omitempty"`

	OperatingSpeedRangeMHz ControlControlRangeExcerpt `json:"OperatingSpeedRangeMHz,omitempty"`

	// The product part number of this device.
	PartNumber *string `json:"PartNumber,omitempty"`

	// Total number of persistent regions this memory device can support.
	PersistentRegionNumberLimit *int64 `json:"PersistentRegionNumberLimit,omitempty"`

	// Total size of persistent regions in mebibytes (MiB).
	PersistentRegionSizeLimitMiB *int64 `json:"PersistentRegionSizeLimitMiB,omitempty"`

	// Maximum size of a single persistent region in mebibytes (MiB).
	PersistentRegionSizeMaxMiB *int64 `json:"PersistentRegionSizeMaxMiB,omitempty"`

	// The maximum number of media error records this device can track in its poison list.
	PoisonListMaxMediaErrorRecords int64 `json:"PoisonListMaxMediaErrorRecords,omitempty"`

	PowerManagementPolicy MemoryV1190PowerManagementPolicy `json:"PowerManagementPolicy,omitempty"`

	// Number of ranks available in the memory device.
	RankCount *int64 `json:"RankCount,omitempty"`

	// Memory regions information within the memory device.
	Regions []MemoryV1190RegionSet `json:"Regions,omitempty"`

	SecurityCapabilities MemoryV1190SecurityCapabilities `json:"SecurityCapabilities,omitempty"`

	SecurityState MemoryV1190SecurityStates `json:"SecurityState,omitempty"`

	SecurityStates MemoryV1190SecurityStateInfo `json:"SecurityStates,omitempty"`

	// The product serial number of this device.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// Number of unused spare devices available in the memory device.
	SpareDeviceCount *int64 `json:"SpareDeviceCount,omitempty"`

	// The spare part number of the memory.
	SparePartNumber *string `json:"SparePartNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// Subsystem device ID.
	// Deprecated
	SubsystemDeviceID *string `json:"SubsystemDeviceID,omitempty"`

	// SubSystem vendor ID.
	// Deprecated
	SubsystemVendorID *string `json:"SubsystemVendorID,omitempty"`

	// Vendor ID.
	// Deprecated
	VendorID *string `json:"VendorID,omitempty"`

	// Total number of volatile regions this memory device can support.
	VolatileRegionNumberLimit *int64 `json:"VolatileRegionNumberLimit,omitempty"`

	// Total size of volatile regions in mebibytes (MiB).
	VolatileRegionSizeLimitMiB *int64 `json:"VolatileRegionSizeLimitMiB,omitempty"`

	// Maximum size of a single volatile region in mebibytes (MiB).
	VolatileRegionSizeMaxMiB *int64 `json:"VolatileRegionSizeMaxMiB,omitempty"`

	// The total volatile memory capacity in mebibytes (MiB).
	VolatileSizeLimitMiB int64 `json:"VolatileSizeLimitMiB,omitempty"`

	// Total size of the volatile portion memory in MiB.
	VolatileSizeMiB *int64 `json:"VolatileSizeMiB,omitempty"`
}

// AssertMemoryV1190MemoryRequired checks if the required fields are not zero-ed
func AssertMemoryV1190MemoryRequired(obj MemoryV1190Memory) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertMemoryV1190ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertMemoryV1190CxlRequired(obj.CXL); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Certificates); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.EnvironmentMetrics); err != nil {
		return err
	}
	if err := AssertMemoryV1190HealthDataRequired(obj.HealthData); err != nil {
		return err
	}
	if err := AssertMemoryV1190LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Log); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockRequired(el); err != nil {
			return err
		}
	}
	if err := AssertMemoryV1190MemoryLocationRequired(obj.MemoryLocation); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Metrics); err != nil {
		return err
	}
	if err := AssertControlControlRangeExcerptRequired(obj.OperatingSpeedRangeMHz); err != nil {
		return err
	}
	if err := AssertMemoryV1190PowerManagementPolicyRequired(obj.PowerManagementPolicy); err != nil {
		return err
	}
	for _, el := range obj.Regions {
		if err := AssertMemoryV1190RegionSetRequired(el); err != nil {
			return err
		}
	}
	if err := AssertMemoryV1190SecurityCapabilitiesRequired(obj.SecurityCapabilities); err != nil {
		return err
	}
	if err := AssertMemoryV1190SecurityStateInfoRequired(obj.SecurityStates); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertMemoryV1190MemoryConstraints checks if the values respects the defined constraints
func AssertMemoryV1190MemoryConstraints(obj MemoryV1190Memory) error {
	return nil
}
