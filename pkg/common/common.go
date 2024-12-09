package common

import "cfm/pkg/openapi"

type ConnectionStatus string

const (
	ONLINE         ConnectionStatus = "online"      // avahi-found, found root service, created backend session, added to service map
	UNAVAILABLE    ConnectionStatus = "unavailable" // avahi-found AND detect root service AND !created backend session AND !added to service map
	OFFLINE        ConnectionStatus = "offline"     // !avahi-found (after previously adding it)
	NOT_APPLICABLE ConnectionStatus = "n\\a"
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
	CustomId:  "",
}

var DefaultHostCredentials = &openapi.Credentials{
	Username:  "admin",
	Password:  "admin12345",
	IpAddress: "127.0.0.1",
	Port:      8082,
	Insecure:  true,
	Protocol:  "http",
	CustomId:  "",
}
