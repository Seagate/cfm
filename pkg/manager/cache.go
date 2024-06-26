// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
/*
The primary struct (DeviceCache) serves 2 purposes.  It is:
1.) The primary location where all information about recognized external devices is saved(cached), and
2.) The primary control mechanism for cfm-service to interact with those devices when executing cfm-service requests.

DeviceCache contains a complete intertface to access any of the primary and secondary objects within the cache.
"Primary objects" are defined as: Appliance, Host
"Secondary objects" are defined as: Blade, Port, Resource, Memory, MemoryDevices

Primary objects can do Get, Add and Delete operations
Secondary objects can only do Get operations.  Add\Delete is handled directly within the primary object.

The main reason for this detailed interface into the cache is cfm-service's need to communicate with multiple devices
at the same time.  Specifically, when cfm-service is communicating with 1 device (say, a Blade), there are times when
cfm-service needs to know what Host that Blade is physically connected to (via a CXL cable).  Since a given Blade
object has no knowledge of Host Socket info, it must be able to search the DeviceCache using information that the
Blade can obtain (like a global CXL ID) to identify which Host it's connect to.  This interface gives devices the
ability to easily create there own searches through the cache.
*/

package manager

import (
	"fmt"
)

type DevicesCache struct {
	appliances map[string]*Appliance
	hosts      map[string]*Host
}

func NewDevicesCache() *DevicesCache {
	c := DevicesCache{
		appliances: make(map[string]*Appliance),
		hosts:      make(map[string]*Host),
	}

	return &c
}

/////////////////////////
// Appliance Functions //
/////////////////////////

func (c *DevicesCache) AddAppliance(appliance *Appliance) error {
	_, ok := c.appliances[appliance.Id]
	if ok {
		return fmt.Errorf("cache already contains appliance with id [%s]", appliance.Id)
	}

	c.appliances[appliance.Id] = appliance

	return nil
}

func (c *DevicesCache) DeleteApplianceById(applianceId string) *Appliance {
	appliance, ok := c.GetApplianceByIdOk(applianceId)
	if !ok {
		return nil
	}

	delete(c.appliances, appliance.Id)

	return appliance

}

func (c *DevicesCache) GetAllApplianceIds() []string {
	var ids []string

	for id := range c.appliances {
		ids = append(ids, id)
	}

	return ids
}

func (c *DevicesCache) GetApplianceById(applianceId string) (*Appliance, error) {
	appliance, ok := c.GetApplianceByIdOk(applianceId)
	if !ok {
		return nil, fmt.Errorf("appliance [%s] doesn't exist", applianceId)
	}

	return appliance, nil
}

func (c *DevicesCache) GetApplianceByIdOk(applianceId string) (*Appliance, bool) {
	appliance, ok := c.appliances[applianceId]

	return appliance, ok
}

// GetAppliances
func (c *DevicesCache) GetAppliances() map[string]*Appliance {
	return c.appliances
}

/////////////////////
// Blade Functions //
/////////////////////

func (c *DevicesCache) GetBladeById(applianceId, bladeId string) (*Blade, error) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, fmt.Errorf("appliance [%s] blade [%s] doesn't exist", applianceId, bladeId)
	}

	return blade, nil
}

func (c *DevicesCache) GetBladeByIdOk(applianceId, bladeId string) (*Blade, bool) {
	appliance, ok := c.GetApplianceByIdOk(applianceId)
	if !ok {
		return nil, ok
	}

	blade, ok := appliance.Blades[bladeId]

	return blade, ok
}

func (c *DevicesCache) GetBladesOk(applianceId string) (map[string]*Blade, bool) {
	appliance, ok := c.GetApplianceByIdOk(applianceId)
	if !ok {
		return nil, ok
	}

	return appliance.Blades, true
}

//There are no "add" or "delete" BladeMemory functions.  Add\Delete is handled within each Blade object.

func (c *DevicesCache) GetBladeMemoryById(applianceId, bladeId, memoryId string) (*BladeMemory, error) {
	memory, ok := c.GetBladeMemoryByIdOk(applianceId, bladeId, memoryId)
	if !ok {
		return nil, fmt.Errorf("appliance [%s] blade [%s] memory [%s] doesn't exist", applianceId, bladeId, memoryId)
	}

	return memory, nil
}

func (c *DevicesCache) GetBladeMemoryByIdOk(applianceId, bladeId, memoryId string) (*BladeMemory, bool) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, ok
	}

	memory, ok := blade.Memory[memoryId]

	return memory, ok
}

func (c *DevicesCache) GetBladeMemoryOk(applianceId, bladeId string) (map[string]*BladeMemory, bool) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, ok
	}

	return blade.Memory, true
}

//There are no "add" or "delete" BladePort functions.  Add\Delete is handled within each Blade object.

func (c *DevicesCache) GetBladePortById(applianceId, bladeId, portId string) (*CxlBladePort, error) {
	port, ok := c.GetBladePortByIdOk(applianceId, bladeId, portId)
	if !ok {
		return nil, fmt.Errorf("appliance [%s] blade [%s] port [%s] doesn't exist", applianceId, bladeId, portId)
	}

	return port, nil
}

func (c *DevicesCache) GetBladePortByIdOk(applianceId, bladeId, portId string) (*CxlBladePort, bool) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, ok
	}

	port, ok := blade.Ports[portId]

	return port, ok
}

func (c *DevicesCache) GetBladePortsOk(applianceId, bladeId string) (map[string]*CxlBladePort, bool) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, ok
	}

	return blade.Ports, true
}

//There are no "add" or "delete" BladeResource functions.  Add\Delete is handled within each Blade object.

func (c *DevicesCache) GetBladeResourceById(applianceId, bladeId, resourceId string) (*BladeResource, error) {
	resource, ok := c.GetBladeResourceByIdOk(applianceId, bladeId, resourceId)
	if !ok {
		return nil, fmt.Errorf("appliance [%s] blade [%s] resource [%s] doesn't exist", applianceId, bladeId, resourceId)
	}

	return resource, nil
}

func (c *DevicesCache) GetBladeResourceByIdOk(applianceId, bladeId, resourceId string) (*BladeResource, bool) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, ok
	}

	resource, ok := blade.Resources[resourceId]

	return resource, ok
}

func (c *DevicesCache) GetBladeResourcesOk(applianceId, bladeId string) (map[string]*BladeResource, bool) {
	blade, ok := c.GetBladeByIdOk(applianceId, bladeId)
	if !ok {
		return nil, ok
	}

	return blade.Resources, true
}

////////////////////
// Host Functions //
////////////////////

func (c *DevicesCache) AddHost(host *Host) error {
	_, ok := c.hosts[host.Id]
	if ok {
		return fmt.Errorf("cache already contains host with id [%s]", host.Id)
	}

	c.hosts[host.Id] = host

	return nil
}

func (c *DevicesCache) DeleteHostById(hostId string) *Host {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil
	}

	delete(c.hosts, host.Id)

	return host
}

func (c *DevicesCache) GetAllHostIds() []string {
	var ids []string

	// CACHE: Get
	for id := range c.hosts {
		ids = append(ids, id)
	}

	return ids
}

func (c *DevicesCache) GetHostById(hostId string) (*Host, error) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, fmt.Errorf("host [%s] doesn't exist", hostId)
	}

	return host, nil
}

func (c *DevicesCache) GetHostByIdOk(hostId string) (*Host, bool) {
	host, ok := c.hosts[hostId]

	return host, ok
}

// GetHosts - returns a slice of all active host URIs
func (c *DevicesCache) GetHosts() map[string]*Host {
	return c.hosts
}

//There are no "add" or "delete" HostMempry functions.  Add\Delete is handled within each Host object.

func (c *DevicesCache) GetHostMemoryById(hostId, memoryId string) (*HostMemory, error) {
	memory, ok := c.GetHostMemoryByIdOk(hostId, memoryId)
	if !ok {
		return nil, fmt.Errorf("host [%s] memory [%s] doesn't exist", hostId, memoryId)
	}

	return memory, nil
}

func (c *DevicesCache) GetHostMemoryByIdOk(hostId, memoryId string) (*HostMemory, bool) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, ok
	}

	memory, ok := host.Memory[memoryId]

	return memory, ok
}

func (c *DevicesCache) GetHostMemoryOk(hostId string) (map[string]*HostMemory, bool) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, ok
	}

	return host.Memory, true
}

//There are no "add" or "delete" HostMemoryDevice functions.  Add\Delete is handled within each Host object.

func (c *DevicesCache) GetHostMemoryDeviceById(hostId, memdevId string) (*HostMemoryDevice, error) {
	port, ok := c.GetHostMemoryDeviceByIdOk(hostId, memdevId)
	if !ok {
		return nil, fmt.Errorf("host [%s] memory device [%s] doesn't exist", hostId, memdevId)
	}

	return port, nil
}

func (c *DevicesCache) GetHostMemoryDeviceByIdOk(hostId, memdevId string) (*HostMemoryDevice, bool) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, ok
	}

	memdev, ok := host.MemoryDevices[memdevId]

	return memdev, ok
}

func (c *DevicesCache) GetHostMemoryDevicesOk(hostId string) (map[string]*HostMemoryDevice, bool) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, ok
	}

	return host.MemoryDevices, true
}

//There are no "add" or "delete" HostPort functions.  Add\Delete is handled within each Host object.

func (c *DevicesCache) GetHostPortById(hostId, portId string) (*CxlHostPort, error) {
	port, ok := c.GetHostPortByIdOk(hostId, portId)
	if !ok {
		return nil, fmt.Errorf("host [%s] port [%s] doesn't exist", hostId, portId)
	}

	return port, nil
}

func (c *DevicesCache) GetHostPortByIdOk(hostId, portId string) (*CxlHostPort, bool) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, ok
	}

	port, ok := host.Ports[portId]

	return port, ok
}

func (c *DevicesCache) GetHostPortsOk(hostId string) (map[string]*CxlHostPort, bool) {
	host, ok := c.GetHostByIdOk(hostId)
	if !ok {
		return nil, ok
	}

	return host.Ports, true
}
