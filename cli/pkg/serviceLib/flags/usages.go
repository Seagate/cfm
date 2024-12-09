// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package flags

import "fmt"

// Family of getter functions to allow retrieval of consistent strings, for each option, when building the cobra.Command Usage field string across multiple commands.
//
// Usage syntax is defined as follows:
// 1.) Square Brackets [ ]: Indicate optional elements. For example, in command [option], the option is not required for the command to run.
// 2.) Angle Brackets < >: Often used to denote placeholders for user-supplied values, such as command <filename>.
// 3.) Curly Braces { }: Indicate a set of choices, where you must choose one. For example, command {start|stop|restart} means you must choose one of start, stop, or restart.
// 4.) Vertical Bar |: Used within curly braces to separate choices, as shown above.
// 5.) Ellipsis ...: Indicates that the preceding element can be repeated multiple times. For example, command [option]... means you can use multiple options.
// 6.) Parentheses ( ): Sometimes used to group elements together, though less common in man pages.
// 7.) Bold Text: Typically used to show the command itself or mandatory elements.
// 8.) Italic Text: Used for arguments or variables that the user must replace with actual values.

// formatUsage - Generates a formatted usage string for a single command option.
// Used for every option in the cobra.Command.Use field.
// This function is meant to be called multiple times to consistently generate a complete string that represents the usage for a single cli command option.
// String format(using all available features): [--option | -o <entry>]
func formatUsage(option, shorthand, entry string, optional bool) string {
	usage := fmt.Sprintf("{--%s | -%s}", option, shorthand)

	if entry != "" {
		usage = fmt.Sprintf("%s <%s>", usage, entry)
	}

	if optional {
		usage = fmt.Sprintf("[%s]", usage)
	} // else {
	// 	usage = fmt.Sprintf("(%s)", usage)
	// }

	return usage
}

func GetOptionUsageGroupServiceTcp(optional bool) string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionUsageServiceIp(optional),
		GetOptionUsageServicePort(optional),
		GetOptionUsageServiceInsecure(optional),
		GetOptionUsageServiceProtocol(optional))
}

func GetOptionUsageServiceIp(optional bool) string {
	return formatUsage(SERVICE_NET_IP, SERVICE_NET_IP_SH, "ip_address", optional)
}

func GetOptionUsageServicePort(optional bool) string {
	return formatUsage(SERVICE_NET_PORT, SERVICE_NET_PORT_SH, PORT, optional)
}

func GetOptionUsageServiceInsecure(optional bool) string {
	return formatUsage(SERVICE_INSECURE, SERVICE_INSECURE_SH, "", optional)
}

func GetOptionUsageServiceProtocol(optional bool) string {
	return formatUsage(SERVICE_PROTOCOL, SERVICE_PROTOCOL_SH, PROTOCOL, optional)
}

func GetOptionUsageApplianceId(optional bool) string {
	return formatUsage(APPLIANCE_ID, APPLIANCE_ID_SH, "applianceId", optional)
}

func GetOptionUsageBladeId(optional bool) string {
	return formatUsage(BLADE_ID, BLADE_ID_SH, "bladeId", optional)
}

func GetOptionUsageHostId(optional bool) string {
	return formatUsage(HOST_ID, HOST_ID_SH, "hostId", optional)
}

func GetOptionUsageNewId(optional bool) string {
	return formatUsage(NEW_ID, NEW_ID_SH, "newId", optional)
}

func GetOptionUsageApplianceUsername(optional bool) string {
	return formatUsage(APPLIANCE_USERNAME, APPLIANCE_USERNAME_SH, USERNAME, optional)
}

func GetOptionUsageAppliancePassword(optional bool) string {
	return formatUsage(APPLIANCE_PASSWORD, APPLIANCE_PASSWORD_SH, PASSWORD, optional)
}

func GetOptionUsageGroupApplianceTcp(optional bool) string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionUsageApplianceIp(optional),
		GetOptionUsageAppliancePort(optional),
		GetOptionUsageApplianceInsecure(optional),
		GetOptionUsageApplianceProtocol(optional))
}

func GetOptionUsageApplianceIp(optional bool) string {
	return formatUsage(APPLIANCE_NET_IP, APPLIANCE_NET_IP_SH, "ip_address", optional)
}

func GetOptionUsageAppliancePort(optional bool) string {
	return formatUsage(APPLIANCE_NET_PORT, APPLIANCE_NET_PORT_SH, PORT, optional)
}

func GetOptionUsageApplianceInsecure(optional bool) string {
	return formatUsage(APPLIANCE_INSECURE, APPLIANCE_INSECURE_SH, "", optional)
}

func GetOptionUsageApplianceProtocol(optional bool) string {
	return formatUsage(APPLIANCE_PROTOCOL, APPLIANCE_PROTOCOL_SH, PROTOCOL, optional)
}

func GetOptionUsageBladeUsername(optional bool) string {
	return formatUsage(BLADE_USERNAME, BLADE_USERNAME_SH, USERNAME, optional)
}

func GetOptionUsageBladePassword(optional bool) string {
	return formatUsage(BLADE_PASSWORD, BLADE_PASSWORD_SH, PASSWORD, optional)
}

func GetOptionUsageGroupBladeTcp(optional bool) string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionUsageBladeIp(optional),
		GetOptionUsageBladePort(optional),
		GetOptionUsageBladeInsecure(optional),
		GetOptionUsageBladeProtocol(optional))
}

func GetOptionUsageBladeIp(optional bool) string {
	return formatUsage(BLADE_NET_IP, BLADE_NET_IP_SH, "ip_address", optional)
}

func GetOptionUsageBladePort(optional bool) string {
	return formatUsage(BLADE_NET_PORT, BLADE_NET_PORT_SH, PORT, optional)
}

func GetOptionUsageBladeInsecure(optional bool) string {
	return formatUsage(BLADE_INSECURE, BLADE_INSECURE_SH, "", optional)
}

func GetOptionUsageBladeProtocol(optional bool) string {
	return formatUsage(BLADE_PROTOCOL, BLADE_PROTOCOL_SH, PROTOCOL, optional)
}

func GetOptionUsageHostUsername(optional bool) string {
	return formatUsage(HOST_USERNAME, HOST_USERNAME_SH, USERNAME, optional)
}

func GetOptionUsageHostPassword(optional bool) string {
	return formatUsage(HOST_PASSWORD, HOST_PASSWORD_SH, PASSWORD, optional)
}

func GetOptionUsageGroupHostTcp(optional bool) string {
	return fmt.Sprintf("%s %s %s %s",
		GetOptionUsageHostIp(optional),
		GetOptionUsageHostPort(optional),
		GetOptionUsageHostInsecure(optional),
		GetOptionUsageHostProtocol(optional))
}

func GetOptionUsageHostIp(optional bool) string {
	return formatUsage(HOST_NET_IP, HOST_NET_IP_SH, "ip_address", optional)
}

func GetOptionUsageHostPort(optional bool) string {
	return formatUsage(HOST_NET_PORT, HOST_NET_PORT_SH, PORT, optional)
}

func GetOptionUsageHostInsecure(optional bool) string {
	return formatUsage(HOST_INSECURE, HOST_INSECURE_SH, "", optional)
}

func GetOptionUsageHostProtocol(optional bool) string {
	return formatUsage(HOST_PROTOCOL, HOST_PROTOCOL_SH, PROTOCOL, optional)
}

func GetOptionUsageMemoryId(optional bool) string {
	return formatUsage(MEMORY_ID, MEMORY_ID_SH, "memoryId", optional)
}

func GetOptionUsageMemoryQos(optional bool) string {
	return formatUsage(MEMORY_QOS, MEMORY_QOS_SH, QOS, optional)
}

func GetOptionUsageMemoryDeviceId(optional bool) string {
	return formatUsage(MEMORY_DEVICE_ID, MEMORY_DEVICE_ID_SH, "memdevId", optional)
}

func GetOptionUsagePortId(optional bool) string {
	return formatUsage(PORT_ID, PORT_ID_SH, "portId", optional)
}

func GetOptionUsageResourceId(optional bool) string {
	return formatUsage(RESOURCE_ID, RESOURCE_ID_SH, "resourceId", optional)
}

func GetOptionUsageResourceSize(optional bool) string {
	return formatUsage(RESOURCE_SIZE, RESOURCE_SIZE_SH, SIZE, optional)
}
