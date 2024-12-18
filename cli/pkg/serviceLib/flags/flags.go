// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package flags

// Constants
const (
	PROTOCOL_HTTP  = "http"
	PROTOCOL_HTTPS = "https"
)

// CLI flag component descriptor
const (
	SERVICE       string = "service"
	APPLIANCE     string = "appliance"
	BLADE         string = "blade"
	HOST          string = "host"
	MEMORY        string = "memory"
	MEMORY_DEVICE string = "memory-device"
	PORT          string = "port"
	RESOURCE      string = "resource"

	APPLIANCES     string = "appliances"
	BLADES         string = "blades"
	HOSTS          string = "hosts"
	MEMORY_DEVICES string = "memory-devices"
	PORTS          string = "ports"
	RESOURCES      string = "resources"
)

// CLI flag detail descriptor
const (
	ID       string = "id"
	USERNAME string = "username"
	PASSWORD string = "password"
	NET_IP   string = "net-ip"
	NET_PORT string = "net-port"
	INSECURE string = "insecure"
	PROTOCOL string = "protocol"
	SIZE     string = "size"
	QOS      string = "qos" //quality of service
	NEW      string = "new"
)

//
// Fully formed flag names generally try to follow a pattern of: component-detail
//

// CLI flag names and shorthand (SH) names
const (
	SERVICE_NET_IP      string = SERVICE + "-" + NET_IP
	SERVICE_NET_IP_SH   string = "a"
	SERVICE_NET_PORT    string = SERVICE + "-" + NET_PORT
	SERVICE_NET_PORT_SH string = "p"
	SERVICE_INSECURE    string = SERVICE + "-" + INSECURE
	SERVICE_INSECURE_SH string = "s"
	SERVICE_PROTOCOL    string = SERVICE + "-" + PROTOCOL
	SERVICE_PROTOCOL_SH string = "t"

	APPLIANCE_ID    string = APPLIANCE + "-" + ID
	APPLIANCE_ID_SH string = "L"
	BLADE_ID        string = BLADE + "-" + ID
	BLADE_ID_SH     string = "B"
	HOST_ID         string = HOST + "-" + ID
	HOST_ID_SH      string = "H"
	NEW_ID          string = NEW + "-" + ID
	NEW_ID_SH       string = "N"

	COMMON_USERNAME_SH string = "R"
	COMMON_PASSWORD_SH string = "W"
	COMMON_NET_IP_SH   string = "A"
	COMMON_NET_PORT_SH string = "P"
	COMMON_INSECURE_SH string = "S"
	COMMON_PROTOCOL_SH string = "T"

	APPLIANCE_USERNAME    string = APPLIANCE + "-" + USERNAME
	APPLIANCE_USERNAME_SH string = COMMON_USERNAME_SH
	APPLIANCE_PASSWORD    string = APPLIANCE + "-" + PASSWORD
	APPLIANCE_PASSWORD_SH string = COMMON_PASSWORD_SH
	APPLIANCE_NET_IP      string = APPLIANCE + "-" + NET_IP
	APPLIANCE_NET_IP_SH   string = COMMON_NET_IP_SH
	APPLIANCE_NET_PORT    string = APPLIANCE + "-" + NET_PORT
	APPLIANCE_NET_PORT_SH string = COMMON_NET_PORT_SH
	APPLIANCE_INSECURE    string = APPLIANCE + "-" + INSECURE
	APPLIANCE_INSECURE_SH string = COMMON_INSECURE_SH
	APPLIANCE_PROTOCOL    string = APPLIANCE + "-" + PROTOCOL
	APPLIANCE_PROTOCOL_SH string = COMMON_PROTOCOL_SH

	BLADE_USERNAME    string = BLADE + "-" + USERNAME
	BLADE_USERNAME_SH string = COMMON_USERNAME_SH
	BLADE_PASSWORD    string = BLADE + "-" + PASSWORD
	BLADE_PASSWORD_SH string = COMMON_PASSWORD_SH
	BLADE_NET_IP      string = BLADE + "-" + NET_IP
	BLADE_NET_IP_SH   string = COMMON_NET_IP_SH
	BLADE_NET_PORT    string = BLADE + "-" + NET_PORT
	BLADE_NET_PORT_SH string = COMMON_NET_PORT_SH
	BLADE_INSECURE    string = BLADE + "-" + INSECURE
	BLADE_INSECURE_SH string = COMMON_INSECURE_SH
	BLADE_PROTOCOL    string = BLADE + "-" + PROTOCOL
	BLADE_PROTOCOL_SH string = COMMON_PROTOCOL_SH

	HOST_USERNAME    string = HOST + "-" + USERNAME
	HOST_USERNAME_SH string = COMMON_USERNAME_SH
	HOST_PASSWORD    string = HOST + "-" + PASSWORD
	HOST_PASSWORD_SH string = COMMON_PASSWORD_SH
	HOST_NET_IP      string = HOST + "-" + NET_IP
	HOST_NET_IP_SH   string = COMMON_NET_IP_SH
	HOST_NET_PORT    string = HOST + "-" + NET_PORT
	HOST_NET_PORT_SH string = COMMON_NET_PORT_SH
	HOST_INSECURE    string = HOST + "-" + INSECURE
	HOST_INSECURE_SH string = COMMON_INSECURE_SH
	HOST_PROTOCOL    string = HOST + "-" + PROTOCOL
	HOST_PROTOCOL_SH string = COMMON_PROTOCOL_SH

	MEMORY_ID           string = MEMORY + "-" + ID
	MEMORY_ID_SH        string = "m"
	MEMORY_QOS          string = MEMORY + "-" + QOS
	MEMORY_QOS_SH       string = "q"
	MEMORY_DEVICE_ID    string = MEMORY_DEVICE + "-" + ID
	MEMORY_DEVICE_ID_SH string = "d"

	PORT_ID    string = PORT + "-" + ID
	PORT_ID_SH string = "o"

	RESOURCE_ID      string = RESOURCE + "-" + ID
	RESOURCE_ID_SH   string = "r"
	RESOURCE_SIZE    string = RESOURCE + "-" + SIZE
	RESOURCE_SIZE_SH string = "z"
)

// CLI misc
const (
	VERBOSITY string = "verbosity"

	// CLI flag default values
	ID_DFLT string = ""

	SERVICE_NET_IP_DFLT   string = "127.0.0.1"
	SERVICE_NET_PORT_DFLT uint16 = 8080
	SERVICE_INSECURE_DFLT bool   = false
	SERVICE_PROTOCOL_DFLT string = PROTOCOL_HTTPS

	APPLIANCE_NET_IP_DFLT   string = "127.0.0.1"
	APPLIANCE_NET_PORT_DFLT uint16 = 443
	APPLIANCE_INSECURE_DFLT bool   = false
	APPLIANCE_PROTOCOL_DFLT string = PROTOCOL_HTTPS
	APPLIANCE_USERNAME_DFLT string = "dummyuser"
	APPLIANCE_PASSWORD_DFLT string = "dummypswd"

	BLADE_NET_IP_DFLT   string = "127.0.0.1"
	BLADE_NET_PORT_DFLT uint16 = 443
	BLADE_INSECURE_DFLT bool   = false
	BLADE_PROTOCOL_DFLT string = PROTOCOL_HTTPS
	BLADE_USERNAME_DFLT string = "root"
	BLADE_PASSWORD_DFLT string = "0penBmc"

	HOST_NET_IP_DFLT   string = "127.0.0.1"
	HOST_NET_PORT_DFLT uint16 = 8082
	HOST_INSECURE_DFLT bool   = false
	HOST_PROTOCOL_DFLT string = PROTOCOL_HTTP
	HOST_USERNAME_DFLT string = "admin"
	HOST_PASSWORD_DFLT string = "admin12345"

	SIZE_DFLT string = "8g"

	MEMORY_QOS_DFLT = 4

	VERBOSITY_DFLT = 0 // Using logging levels 0-4, with 0 being "OFF" and 4 the most verbose
)
