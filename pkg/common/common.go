package common

import "cfm/pkg/openapi"

type ConnectionStatus string

const (
	ONLINE         ConnectionStatus = "online"
	FOUND          ConnectionStatus = "found"
	OFFLINE        ConnectionStatus = "offline"
	NOT_APPLICABLE ConnectionStatus = "n\\a"
	UNAVAILABLE        ConnectionStatus = "unavailable"
)

var DefaultApplianceCredentials = &openapi.Credentials{
	Username:  "root",
	Password:  "0penBmc",
	IpAddress: "127.0.0.1",
	Port:      8443,
	Insecure:  true,
	Protocol:  "https",
	CustomId:  "CMA_Discovered_Blades",
}

var DefaultBladeCredentials = &openapi.Credentials{
	Username:  "root",
	Password:  "0penBmc",
	IpAddress: "127.0.0.1",
	Port:      443,
	Insecure:  true,
	Protocol:  "https",
	CustomId:  "Discoverd_Blade_",
}

var DefaultHostCredentials = &openapi.Credentials{
	Username:  "admin",
	Password:  "admin12345",
	IpAddress: "127.0.0.1",
	Port:      8082,
	Insecure:  true,
	Protocol:  "http",
	CustomId:  "Discoverd_Host_",
}
