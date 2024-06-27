// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import "fmt"

func GetCfmUriVersion() string {
	return "/cfm/v1"
}

func GetCfmUriAppliances() string {
	return "/cfm/v1/appliances"
}

func GetCfmUriApplianceId(applianceId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s", applianceId)
}

func GetCfmUriBlades(applianceId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades", applianceId)
}

func GetCfmUriBladeId(applianceId, bladeId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s", applianceId, bladeId)
}

func GetCfmUriBladeResources(applianceId, bladeId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s/resources", applianceId, bladeId)
}

func GetCfmUriBladeResourceId(applianceId, bladeId, resourceId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s/resources/%s", applianceId, bladeId, resourceId)
}

func GetCfmUriBladePorts(applianceId, bladeId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s/ports", applianceId, bladeId)
}

func GetCfmUriBladePortId(applianceId, bladeId, portId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s/ports/%s", applianceId, bladeId, portId)
}

func GetCfmUriBladeMemory(applianceId, bladeId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s/memory", applianceId, bladeId)
}

func GetCfmUriBladeMemoryId(applianceId, bladeId, memoryId string) string {
	return fmt.Sprintf("/cfm/v1/appliances/%s/blades/%s/memory/%s", applianceId, bladeId, memoryId)
}

func GetCfmUriHosts() string {
	return "/cfm/v1/hosts"
}

func GetCfmUriHostId(hostId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s", hostId)
}

func GetCfmUriHostMemory(hostId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s/memory", hostId)
}

func GetCfmUriHostMemoryId(hostId, memoryId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s/memory/%s", hostId, memoryId)
}

func GetCfmUriHostMemoryDevices(hostId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s/memory-devices", hostId)
}

func GetCfmUriHostMemoryDeviceId(hostId, memDevId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s/memory-devices/%s", hostId, memDevId)
}

func GetCfmUriHostPorts(hostId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s/ports", hostId)
}

func GetCfmUriHostPortId(hostId, portId string) string {
	return fmt.Sprintf("/cfm/v1/hosts/%s/ports/%s", hostId, portId)
}
