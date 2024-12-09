// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package flags

import "fmt"

// Family of getter functions to allow retrieval of consistent strings, for each option, when building the cobra.Command Example field string across multiple commands.

func GetOptionExampleLhGroupServiceTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleLhServiceIp(),
		GetOptionExampleLhServicePort(),
		GetOptionExampleLhServiceInsecure(),
		GetOptionExampleLhServiceProtocol())
}

func GetOptionExampleLhServiceIp() string {
	return fmt.Sprintf("--%s %s", SERVICE_NET_IP, SERVICE_NET_IP_DFLT)
}

func GetOptionExampleLhServicePort() string {
	return fmt.Sprintf("--%s %d", SERVICE_NET_PORT, SERVICE_NET_PORT_DFLT)
}

func GetOptionExampleLhServiceInsecure() string {
	return fmt.Sprintf("--%s", SERVICE_INSECURE)
}

func GetOptionExampleLhServiceProtocol() string {
	return fmt.Sprintf("--%s %s", SERVICE_PROTOCOL, SERVICE_PROTOCOL_DFLT)
}

func GetOptionExampleLhApplianceId() string {
	return fmt.Sprintf("--%s applianceId", APPLIANCE_ID)
}

func GetOptionExampleLhBladeId() string {
	return fmt.Sprintf("--%s bladeId", BLADE_ID)
}

func GetOptionExampleLhHostId() string {
	return fmt.Sprintf("--%s hostId", HOST_ID)
}

func GetOptionExampleLhNewId() string {
	return fmt.Sprintf("--%s newId", NEW_ID)
}

func GetOptionExampleLhApplianceUsername() string {
	return fmt.Sprintf("--%s %s", APPLIANCE_USERNAME, APPLIANCE_USERNAME_DFLT)
}

func GetOptionExampleLhAppliancePassword() string {
	return fmt.Sprintf("--%s %s", APPLIANCE_PASSWORD, APPLIANCE_PASSWORD_DFLT)
}

func GetOptionExampleLhGroupApplianceTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleLhApplianceIp(),
		GetOptionExampleLhAppliancePort(),
		GetOptionExampleLhApplianceInsecure(),
		GetOptionExampleLhApplianceProtocol())
}

func GetOptionExampleLhApplianceIp() string {
	return fmt.Sprintf("--%s %s", APPLIANCE_NET_IP, APPLIANCE_NET_IP_DFLT)
}

func GetOptionExampleLhAppliancePort() string {
	return fmt.Sprintf("--%s %d", APPLIANCE_NET_PORT, APPLIANCE_NET_PORT_DFLT)
}

func GetOptionExampleLhApplianceInsecure() string {
	return fmt.Sprintf("--%s", APPLIANCE_INSECURE)
}

func GetOptionExampleLhApplianceProtocol() string {
	return fmt.Sprintf("--%s %s", APPLIANCE_PROTOCOL, APPLIANCE_PROTOCOL_DFLT)
}

func GetOptionExampleLhBladeUsername() string {
	return fmt.Sprintf("--%s %s", BLADE_USERNAME, BLADE_USERNAME_DFLT)
}

func GetOptionExampleLhBladePassword() string {
	return fmt.Sprintf("--%s %s", BLADE_PASSWORD, BLADE_PASSWORD_DFLT)
}

func GetOptionExampleLhGroupBladeTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleLhBladeIp(),
		GetOptionExampleLhBladePort(),
		GetOptionExampleLhBladeInsecure(),
		GetOptionExampleLhBladeProtocol())
}

func GetOptionExampleLhBladeIp() string {
	return fmt.Sprintf("--%s %s", BLADE_NET_IP, BLADE_NET_IP_DFLT)
}

func GetOptionExampleLhBladePort() string {
	return fmt.Sprintf("--%s %d", BLADE_NET_PORT, BLADE_NET_PORT_DFLT)
}

func GetOptionExampleLhBladeInsecure() string {
	return fmt.Sprintf("--%s", BLADE_INSECURE)
}

func GetOptionExampleLhBladeProtocol() string {
	return fmt.Sprintf("--%s %s", BLADE_PROTOCOL, BLADE_PROTOCOL_DFLT)
}

func GetOptionExampleLhHostUsername() string {
	return fmt.Sprintf("--%s %s", HOST_USERNAME, HOST_USERNAME_DFLT)
}

func GetOptionExampleLhHostPassword() string {
	return fmt.Sprintf("--%s %s", HOST_PASSWORD, HOST_PASSWORD_DFLT)
}

func GetOptionExampleLhGroupHostTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleLhHostIp(),
		GetOptionExampleLhHostPort(),
		GetOptionExampleLhHostInsecure(),
		GetOptionExampleLhHostProtocol())
}

func GetOptionExampleLhHostIp() string {
	return fmt.Sprintf("--%s %s", HOST_NET_IP, HOST_NET_IP_DFLT)
}

func GetOptionExampleLhHostPort() string {
	return fmt.Sprintf("--%s %d", HOST_NET_PORT, HOST_NET_PORT_DFLT)
}

func GetOptionExampleLhHostInsecure() string {
	return fmt.Sprintf("--%s", HOST_INSECURE)
}

func GetOptionExampleLhHostProtocol() string {
	return fmt.Sprintf("--%s %s", HOST_PROTOCOL, HOST_PROTOCOL_DFLT)
}

func GetOptionExampleLhMemoryId() string {
	return fmt.Sprintf("--%s memoryId", MEMORY_ID)
}

func GetOptionExampleLhMemoryQos() string {
	return fmt.Sprintf("--%s %d", MEMORY_QOS, MEMORY_QOS_DFLT)
}

func GetOptionExampleLhMemoryDeviceId() string {
	return fmt.Sprintf("--%s memoryDeviceId", MEMORY_DEVICE_ID)
}

func GetOptionExampleLhPortId() string {
	return fmt.Sprintf("--%s portId", PORT_ID)
}

func GetOptionExampleLhResourceId() string {
	return fmt.Sprintf("--%s resourceId", RESOURCE_ID)
}

func GetOptionExampleLhResourceSize() string {
	return fmt.Sprintf("--%s %s", RESOURCE_SIZE, SIZE_DFLT)
}

func GetOptionExampleShGroupServiceTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleShServiceIp(),
		GetOptionExampleShServicePort(),
		GetOptionExampleShServiceInsecure(),
		GetOptionExampleShServiceProtocol())
}

func GetOptionExampleShServiceIp() string {
	return fmt.Sprintf("-%s %s", SERVICE_NET_IP_SH, SERVICE_NET_IP_DFLT)
}

func GetOptionExampleShServicePort() string {
	return fmt.Sprintf("-%s %d", SERVICE_NET_PORT_SH, SERVICE_NET_PORT_DFLT)
}

func GetOptionExampleShServiceInsecure() string {
	return fmt.Sprintf("-%s", SERVICE_INSECURE_SH)
}

func GetOptionExampleShServiceProtocol() string {
	return fmt.Sprintf("-%s %s", SERVICE_PROTOCOL_SH, SERVICE_PROTOCOL_DFLT)
}

func GetOptionExampleShApplianceId() string {
	return fmt.Sprintf("-%s applianceId", APPLIANCE_ID_SH)
}

func GetOptionExampleShBladeId() string {
	return fmt.Sprintf("-%s bladeId", BLADE_ID_SH)
}

func GetOptionExampleShHostId() string {
	return fmt.Sprintf("-%s hostId", HOST_ID_SH)
}

func GetOptionExampleShNewId() string {
	return fmt.Sprintf("-%s newId", NEW_ID_SH)
}

func GetOptionExampleShApplianceUsername() string {
	return fmt.Sprintf("-%s %s", APPLIANCE_USERNAME_SH, APPLIANCE_USERNAME_DFLT)
}

func GetOptionExampleShAppliancePassword() string {
	return fmt.Sprintf("-%s %s", APPLIANCE_PASSWORD_SH, APPLIANCE_PASSWORD_DFLT)
}

func GetOptionExampleShGroupApplianceTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleShApplianceIp(),
		GetOptionExampleShAppliancePort(),
		GetOptionExampleShApplianceInsecure(),
		GetOptionExampleShApplianceProtocol())
}

func GetOptionExampleShApplianceIp() string {
	return fmt.Sprintf("-%s %s", APPLIANCE_NET_IP_SH, APPLIANCE_NET_IP_DFLT)
}

func GetOptionExampleShAppliancePort() string {
	return fmt.Sprintf("-%s %d", APPLIANCE_NET_PORT_SH, APPLIANCE_NET_PORT_DFLT)
}

func GetOptionExampleShApplianceInsecure() string {
	return fmt.Sprintf("-%s", APPLIANCE_INSECURE_SH)
}

func GetOptionExampleShApplianceProtocol() string {
	return fmt.Sprintf("-%s %s", APPLIANCE_PROTOCOL_SH, APPLIANCE_PROTOCOL_DFLT)
}

func GetOptionExampleShBladeUsername() string {
	return fmt.Sprintf("-%s %s", BLADE_USERNAME_SH, BLADE_USERNAME_DFLT)
}

func GetOptionExampleShBladePassword() string {
	return fmt.Sprintf("-%s %s", BLADE_PASSWORD_SH, BLADE_PASSWORD_DFLT)
}

func GetOptionExampleShGroupBladeTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleShBladeIp(),
		GetOptionExampleShBladePort(),
		GetOptionExampleShBladeInsecure(),
		GetOptionExampleShBladeProtocol())
}

func GetOptionExampleShBladeIp() string {
	return fmt.Sprintf("-%s %s", BLADE_NET_IP_SH, BLADE_NET_IP_DFLT)
}

func GetOptionExampleShBladePort() string {
	return fmt.Sprintf("-%s %d", BLADE_NET_PORT_SH, BLADE_NET_PORT_DFLT)
}

func GetOptionExampleShBladeInsecure() string {
	return fmt.Sprintf("-%s", BLADE_INSECURE_SH)
}

func GetOptionExampleShBladeProtocol() string {
	return fmt.Sprintf("-%s %s", BLADE_PROTOCOL_SH, BLADE_PROTOCOL_DFLT)
}

func GetOptionExampleShHostUsername() string {
	return fmt.Sprintf("-%s %s", HOST_USERNAME_SH, HOST_USERNAME_DFLT)
}

func GetOptionExampleShHostPassword() string {
	return fmt.Sprintf("-%s %s", HOST_PASSWORD_SH, HOST_PASSWORD_DFLT)
}

func GetOptionExampleShGroupHostTcp() string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionExampleShHostIp(),
		GetOptionExampleShHostPort(),
		GetOptionExampleShHostInsecure(),
		GetOptionExampleShHostProtocol())
}

func GetOptionExampleShHostIp() string {
	return fmt.Sprintf("-%s %s", HOST_NET_IP_SH, HOST_NET_IP_DFLT)
}

func GetOptionExampleShHostPort() string {
	return fmt.Sprintf("-%s %d", HOST_NET_PORT_SH, HOST_NET_PORT_DFLT)
}

func GetOptionExampleShHostInsecure() string {
	return fmt.Sprintf("-%s", HOST_INSECURE_SH)
}

func GetOptionExampleShHostProtocol() string {
	return fmt.Sprintf("-%s %s", HOST_PROTOCOL_SH, HOST_PROTOCOL_DFLT)
}

func GetOptionExampleShMemoryId() string {
	return fmt.Sprintf("-%s memoryId", MEMORY_ID_SH)
}

func GetOptionExampleShMemoryQos() string {
	return fmt.Sprintf("-%s %d", MEMORY_QOS_SH, MEMORY_QOS_DFLT)
}

func GetOptionExampleShMemoryDeviceId() string {
	return fmt.Sprintf("-%s memoryDeviceId", MEMORY_DEVICE_ID_SH)
}

func GetOptionExampleShPortId() string {
	return fmt.Sprintf("-%s portId", PORT_ID_SH)
}

func GetOptionExampleShResourceId() string {
	return fmt.Sprintf("-%s resourceId", RESOURCE_ID_SH)
}

func GetOptionExampleShResourceSize() string {
	return fmt.Sprintf("-%s %s", RESOURCE_SIZE_SH, SIZE_DFLT)
}
