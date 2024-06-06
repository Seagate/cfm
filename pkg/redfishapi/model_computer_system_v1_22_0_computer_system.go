/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"time"
)

// ComputerSystemV1220ComputerSystem - The ComputerSystem schema represents a computer or system instance and the software-visible resources, or items within the data plane, such as memory, CPU, and other devices that it can access.  Details of those resources or subsystems are also linked through this resource.
type ComputerSystemV1220ComputerSystem struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ComputerSystemV1220Actions `json:"Actions,omitempty"`

	// The user-definable tag that can track this computer system for inventory or other client purposes.
	AssetTag *string `json:"AssetTag,omitempty"`

	Bios OdataV4IdRef `json:"Bios,omitempty"`

	// The version of the system BIOS or primary system firmware.
	BiosVersion *string `json:"BiosVersion,omitempty"`

	Boot ComputerSystemV1220Boot `json:"Boot,omitempty"`

	BootProgress ComputerSystemV1220BootProgress `json:"BootProgress,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	Composition ComputerSystemV1220Composition `json:"Composition,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	EthernetInterfaces OdataV4IdRef `json:"EthernetInterfaces,omitempty"`

	FabricAdapters OdataV4IdRef `json:"FabricAdapters,omitempty"`

	GraphicalConsole ComputerSystemV1220HostGraphicalConsole `json:"GraphicalConsole,omitempty"`

	GraphicsControllers OdataV4IdRef `json:"GraphicsControllers,omitempty"`

	// The DNS host name, without any domain information.
	HostName *string `json:"HostName,omitempty"`

	HostWatchdogTimer ComputerSystemV1220WatchdogTimer `json:"HostWatchdogTimer,omitempty"`

	HostedServices ComputerSystemV1220HostedServices `json:"HostedServices,omitempty"`

	// The hosting roles that this computer system supports.
	HostingRoles []ComputerSystemV1220HostingRole `json:"HostingRoles,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	IdlePowerSaver ComputerSystemV1220IdlePowerSaver `json:"IdlePowerSaver,omitempty"`

	IndicatorLED ComputerSystemV1220IndicatorLed `json:"IndicatorLED,omitempty"`

	KeyManagement ComputerSystemV1220KeyManagement `json:"KeyManagement,omitempty"`

	// The date and time when the system was last reset or rebooted.
	LastResetTime time.Time `json:"LastResetTime,omitempty"`

	Links ComputerSystemV1220Links `json:"Links,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	LogServices OdataV4IdRef `json:"LogServices,omitempty"`

	// The manufacturer or OEM of this system.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// An indication of whether the system is in manufacturing mode.  Manufacturing mode is a special boot mode, not normally available to end users, that modifies features and settings for use while the system is being manufactured and tested.
	ManufacturingMode *bool `json:"ManufacturingMode,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	// Deprecated
	Measurements []SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`

	Memory OdataV4IdRef `json:"Memory,omitempty"`

	MemoryDomains OdataV4IdRef `json:"MemoryDomains,omitempty"`

	MemorySummary ComputerSystemV1220MemorySummary `json:"MemorySummary,omitempty"`

	// The product name for this system, without the manufacturer name.
	Model *string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	NetworkInterfaces OdataV4IdRef `json:"NetworkInterfaces,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	OperatingSystem OdataV4IdRef `json:"OperatingSystem,omitempty"`

	// The link to a collection of PCIe devices that this computer system uses.
	PCIeDevices []OdataV4IdRef `json:"PCIeDevices,omitempty"`

	// The number of items in a collection.
	PCIeDevicesodataCount int64 `json:"PCIeDevices@odata.count,omitempty"`

	// The link to a collection of PCIe functions that this computer system uses.
	PCIeFunctions []OdataV4IdRef `json:"PCIeFunctions,omitempty"`

	// The number of items in a collection.
	PCIeFunctionsodataCount int64 `json:"PCIeFunctions@odata.count,omitempty"`

	// The part number for this system.
	PartNumber *string `json:"PartNumber,omitempty"`

	// The number of seconds to delay power on after a `Reset` action requesting `PowerCycle`.  Zero seconds indicates no delay.
	PowerCycleDelaySeconds *float32 `json:"PowerCycleDelaySeconds,omitempty"`

	PowerMode ComputerSystemV1220PowerMode `json:"PowerMode,omitempty"`

	// The number of seconds to delay power off during a reset.  Zero seconds indicates no delay to power off.
	PowerOffDelaySeconds *float32 `json:"PowerOffDelaySeconds,omitempty"`

	// The number of seconds to delay power on after a power cycle or during a reset.  Zero seconds indicates no delay to power up.
	PowerOnDelaySeconds *float32 `json:"PowerOnDelaySeconds,omitempty"`

	PowerRestorePolicy ComputerSystemV1220PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`

	PowerState ResourcePowerState `json:"PowerState,omitempty"`

	ProcessorSummary ComputerSystemV1220ProcessorSummary `json:"ProcessorSummary,omitempty"`

	Processors OdataV4IdRef `json:"Processors,omitempty"`

	// The link to a collection of redundancy entities.  Each entity specifies a kind and level of redundancy and a collection, or redundancy set, of other computer systems that provide the specified redundancy to this computer system.
	Redundancy []RedundancyRedundancy `json:"Redundancy,omitempty"`

	// The number of items in a collection.
	RedundancyodataCount int64 `json:"Redundancy@odata.count,omitempty"`

	// The manufacturer SKU for this system.
	SKU *string `json:"SKU,omitempty"`

	SecureBoot OdataV4IdRef `json:"SecureBoot,omitempty"`

	SerialConsole ComputerSystemV1220HostSerialConsole `json:"SerialConsole,omitempty"`

	// The serial number for this system.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	SimpleStorage OdataV4IdRef `json:"SimpleStorage,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	Storage OdataV4IdRef `json:"Storage,omitempty"`

	// The sub-model for this system.
	SubModel *string `json:"SubModel,omitempty"`

	SystemType ComputerSystemV1220SystemType `json:"SystemType,omitempty"`

	// An array of trusted modules in the system.
	// Deprecated
	TrustedModules []ComputerSystemV1220TrustedModules `json:"TrustedModules,omitempty"`

	USBControllers OdataV4IdRef `json:"USBControllers,omitempty"`

	UUID string `json:"UUID,omitempty"`

	VirtualMedia OdataV4IdRef `json:"VirtualMedia,omitempty"`

	VirtualMediaConfig ComputerSystemV1220VirtualMediaConfig `json:"VirtualMediaConfig,omitempty"`
}

// AssertComputerSystemV1220ComputerSystemRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220ComputerSystemRequired(obj ComputerSystemV1220ComputerSystem) error {
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

	if err := AssertComputerSystemV1220ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Bios); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220BootRequired(obj.Boot); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220BootProgressRequired(obj.BootProgress); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Certificates); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220CompositionRequired(obj.Composition); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.EthernetInterfaces); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.FabricAdapters); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220HostGraphicalConsoleRequired(obj.GraphicalConsole); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.GraphicsControllers); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220WatchdogTimerRequired(obj.HostWatchdogTimer); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220HostedServicesRequired(obj.HostedServices); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220IdlePowerSaverRequired(obj.IdlePowerSaver); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220KeyManagementRequired(obj.KeyManagement); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.LogServices); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.Memory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.MemoryDomains); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220MemorySummaryRequired(obj.MemorySummary); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.NetworkInterfaces); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.OperatingSystem); err != nil {
		return err
	}
	for _, el := range obj.PCIeDevices {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PCIeFunctions {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertComputerSystemV1220ProcessorSummaryRequired(obj.ProcessorSummary); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Processors); err != nil {
		return err
	}
	for _, el := range obj.Redundancy {
		if err := AssertRedundancyRedundancyRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.SecureBoot); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220HostSerialConsoleRequired(obj.SerialConsole); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.SimpleStorage); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Storage); err != nil {
		return err
	}
	for _, el := range obj.TrustedModules {
		if err := AssertComputerSystemV1220TrustedModulesRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.USBControllers); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.VirtualMedia); err != nil {
		return err
	}
	if err := AssertComputerSystemV1220VirtualMediaConfigRequired(obj.VirtualMediaConfig); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1220ComputerSystemConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220ComputerSystemConstraints(obj ComputerSystemV1220ComputerSystem) error {
	return nil
}
