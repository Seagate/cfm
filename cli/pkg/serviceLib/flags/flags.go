/*
Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

*/

package flags

// CLI flag component descriptor
const (
	SERVICE       string = "serv"
	DEVICE        string = "dev" //generic "device" for appliance OR host - future deprecation
	APPLIANCE     string = "appliance"
	BLADE         string = "blade"
	HOST          string = "host"
	MEMORY        string = "memory"
	MEMORY_DEVICE string = "memory-device"
	PORT          string = "port"
	RESOURCE      string = "resource"
)

// CLI flag detail descriptor
const (
	ID       string = "id"
	USERNAME string = "username"
	PASSWORD string = "password"
	NET_IP   string = "ip"
	NET_PORT string = "net-port"
	INSECURE string = "insecure"
	PROTOCOL string = "protocol"
	SIZE     string = "size"
	QOS      string = "qos" //quality of service
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
	SERVICE_INSECURE_SH string = "" //"s"
	SERVICE_PROTOCOL    string = SERVICE + "-" + PROTOCOL
	SERVICE_PROTOCOL_SH string = "" //"t"

	APPLIANCE_ID    string = APPLIANCE + "-" + ID
	APPLIANCE_ID_SH string = "L"
	BLADE_ID        string = BLADE + "-" + ID
	BLADE_ID_SH     string = "B"
	HOST_ID         string = HOST + "-" + ID
	HOST_ID_SH      string = "H"

	DEVICE_USERNAME    string = DEVICE + "-" + USERNAME
	DEVICE_USERNAME_SH string = "R"
	DEVICE_PASSWORD    string = DEVICE + "-" + PASSWORD
	DEVICE_PASSWORD_SH string = "W"
	DEVICE_NET_IP      string = DEVICE + "-" + NET_IP
	DEVICE_NET_IP_SH   string = "A"
	DEVICE_NET_PORT    string = DEVICE + "-" + NET_PORT
	DEVICE_NET_PORT_SH string = "P"
	DEVICE_INSECURE    string = DEVICE + "-" + INSECURE
	DEVICE_INSECURE_SH string = "S"
	DEVICE_PROTOCOL    string = DEVICE + "-" + PROTOCOL
	DEVICE_PROTOCOL_SH string = "T"

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
	SERVICE_PROTOCOL_DFLT string = "http"

	APPLIANCE_NET_IP_DFLT   string = "127.0.0.1"
	APPLIANCE_NET_PORT_DFLT uint16 = 443
	APPLIANCE_INSECURE_DFLT bool   = false
	APPLIANCE_PROTOCOL_DFLT string = "https"
	APPLIANCE_USERNAME_DFLT string = "dummyuser"
	APPLIANCE_PASSWORD_DFLT string = "dummypswd"

	BLADE_NET_IP_DFLT   string = "127.0.0.1"
	BLADE_NET_PORT_DFLT uint16 = 443
	BLADE_INSECURE_DFLT bool   = false
	BLADE_PROTOCOL_DFLT string = "https"
	BLADE_USERNAME_DFLT string = "root"
	BLADE_PASSWORD_DFLT string = "0penBmc"

	HOST_NET_IP_DFLT   string = "127.0.0.1"
	HOST_NET_PORT_DFLT uint16 = 8082
	HOST_INSECURE_DFLT bool   = false
	HOST_PROTOCOL_DFLT string = "http"
	HOST_USERNAME_DFLT string = "admin"
	HOST_PASSWORD_DFLT string = "admin12345"

	SIZE_DFLT string = "0"

	MEMORY_QOS_DFLT = 4

	VERBOSITY_DFLT = 0 // Using logging levels 0-4, with 0 being "OFF" and 4 the most verbose
)
