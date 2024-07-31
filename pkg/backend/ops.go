// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package backend

type CreateSessionRequest struct {
	Ip       string // Ip Address
	Port     int32  // Port
	Username string // Service username
	Password string // Service password
	Insecure bool   // Controls the use of a secure connection
	Protocol string // Examples of http vs https
}

type CreateSessionResponse struct {
	SessionId    string // The session id returned form creating a session
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type DeleteSessionRequest struct {
}

type DeleteSessionResponse struct {
	SessionId    string // The session id we are ending
	IpAddress    string // Ip Address
	Port         int32  // Port
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type PortNumber int32

type GetMemoryRegionResponse struct {
	Id           string       // Id of the memory region
	Ports        []PortNumber // Port numbers on fabric
	CapacityMiB  int32        // Allocated memory capacity
	Status       string       // The status of the request
	ServiceError error        // Any error returned by the service
}

type QoS int32

type AllocateMemoryRequest struct {
	SizeMiB int32 // The number of mebibytes of memory to allocate\provision
	Qos     QoS   // Quality of service setting
}

type AllocateMemoryResponse struct {
	SizeMiB      int32  // The allocated number of mebibytes (This may be adjusted to implement dimm interleave)
	MemoryId     string // The id of the memory region allocated\provisioned
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type AllocateMemoryByResourceRequest struct {
	MemoryResoureIds []string // Id's of specific resource blocks to be used to allocate\provision the new memory region
}

type AllocateMemoryByResourceResponse struct {
	SizeMiB      int32  // The allocated number of mebibytes
	MemoryId     string // The id of the memory region allocated\provisioned
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type AssignMemoryRequest struct {
	MemoryId string // The id of the memory region
	PortId   string // The CXL port id to map the memory region to, may be empty to indicate no mapping
}

type AssignMemoryResponse struct {
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type UnassignMemoryRequest struct {
	MemoryId string // The id of the memory region to disconnect(unassign) from the designated port
	PortId   string // The CXL port id to disconnect(unassign) from the designated memory region.
}

type UnassignMemoryResponse struct {
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type MemoryResourceBlocksRequest struct {
}

type MemoryResourceBlocksResponse struct {
	MemoryResources []string // Array to hold ids of memory resources, get from memory appliance
	Status          string   // The status of the request
	ServiceError    error    // Any error returned by the service
}

type GetPortsRequest struct {
}

type GetPortsResponse struct {
	PortIds      []string // Array to hold ids of ports
	Status       string   // The status of the request
	ServiceError error    // Any error returned by the service
}

type GetPortDetailsRequest struct {
	PortId string // Available port id
}

type PortInformation struct {
	Id               string // The id of this resource, for example, "port0"
	GCxlId           string // Unique port serial number which can be used to map with cxl-host port
	LinkedPortUri    string // If physically connected, the port uri on the "other end" of the connection.
	PortProtocol     string // The protocol being sent over this port, for example "CXL"
	PortMedium       string // The physical connection medium for this port
	CurrentSpeedGbps int32  // The current speed of this port
	StatusHealth     string // The health of the resource
	StatusState      string // The state of the resource
	Width            int32  // The number of lanes, phys, or other physical transport links that this port contains
	LinkStatus       string // The CXL link status
	LinkState        string // The CXL link state
}

type GetPortDetailsResponse struct {
	PortInformation PortInformation // Detail info for one port id
	Status          string          // The status of the request
	ServiceError    error           // Any error returned by the service
}

type GetHostPortSnByIdRequest struct {
	PortId string
}

type GetHostPortSnByIdResponse struct {
	SerialNumber string
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type GetMemoryDevicesRequest struct {
}

type GetMemoryDevicesResponse struct {
	DeviceIdMap  map[string][]string // map of physical devices ids (keys) to 1 or more logical device ids (values)
	Status       string              // The status of the request
	ServiceError error               // Any error returned by the service
}

type GetMemoryDeviceDetailsRequest struct {
	PhysicalDeviceId string
	LogicalDeviceId  string
}

type MemoryDeviceLinkStatus struct {
	CurrentSpeedGTps int32 `json:"CurrentSpeedGTps,omitempty"` // The current speed of this link
	MaxSpeedGTps     int32 `json:"MaxSpeedGTps,omitempty"`     // The max speed of this link
	CurrentWidth     int32 `json:"CurrentWidth,omitempty"`     // The current width of this link
	MaxWidth         int32 `json:"MaxWidth,omitempty"`         // The max width of this link
}

type GetMemoryDeviceDetailsResponse struct {
	SerialNumber  string
	DeviceType    string                  // The type of the device
	MemorySizeMiB int32                   // The memory size of the device in MiB
	LinkStatus    *MemoryDeviceLinkStatus // Detail info for one port id
	Status        string                  // The status of the request
	ServiceError  error                   // Any error returned by the service
}

type MemoryResourceBlockByIdRequest struct {
	ResourceId string // Resource ID for a particular memory resource, for example, resource01
}

type ResourceState int32

const (
	ResourceUnused      ResourceState = iota // The resource is not in use and available for reservation or allocation.
	ResourceUnavailable                      // The resource has been made unavailable by the service, such as due to maintenance being performed on the resource.
	ResourceReserved                         // The resource is reserved so that it can be used for allocation and no other request can use it.
	ResourceComposed                         // The resource has been used for allocation and assignment, this is the final state for a resource.
	ResourceFailed                           // A prior allocation step failed and this resource is no longer available for allocation, manual intervention may be required to fix it.
)

func (s ResourceState) String() string {
	switch s {
	case ResourceUnused:
		return RESOURCE_STATE_UNUSED
	case ResourceUnavailable:
		return RESOURCE_STATE_UNAVAILABLE
	case ResourceReserved:
		return RESOURCE_STATE_RESERVED
	case ResourceComposed:
		return RESOURCE_STATE_COMPOSED
	case ResourceFailed:
		return RESOURCE_STATE_FAILED
	}
	return RESOURCE_STATE_UNKNOWN
}

type MemoryResourceBlockCompositionStatus struct {
	CompositionState     ResourceState // The current state of the resource block from a composition perspective
	MaxCompositions      int32         // The maximum number of compositions in which this resource block can participate simultaneously
	NumberOfCompositions int32         // The number of compositions in which this resource block is currently participating
}

type MemoryResourceBlock struct {
	Id                 string // The id of this resource
	CompositionStatus  MemoryResourceBlockCompositionStatus
	CapacityMiB        int32  // The number of compositions in which this resource block is currently participating
	DataWidthBits      int32  // The number of compositions in which this resource block is currently participating
	MemoryType         string // The type of memory device
	MemoryDeviceType   string // Type details of the memory device
	Manufacturer       string // The memory device manufacturer
	OperatingSpeedMhz  int32  // Operating speed of the memory device in MHz
	PartNumber         string // The product part number of this device
	SerialNumber       string // The product serial number of this device
	RankCount          int32  // Number of ranks available in the memory device
	ChannelId          int32  // The id of the hardware channel associated with this resource
	ChannelResourceIdx int32  // The index for this single resource within the given channel (designated by "ChannelId")
}

type MemoryResourceBlockByIdResponse struct {
	MemoryResourceBlock MemoryResourceBlock // Detail info for one memory resource block
	Status              string              // The status of the request
	ServiceError        error               // Any error returned by the service
}

type FreeMemoryRequest struct {
	MemoryId string
}

type FreeMemoryResponse struct {
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type GetMemoryByIdRequest struct {
	SessionId string // The backend session id for the desired BMC target (on appliance or host)
	MemoryId  string
}

type MemoryType string

const (
	MEMORYTYPE_MEMORY_TYPE_UNKNOWN MemoryType = "MemoryTypeUnknown"
	MEMORYTYPE_MEMORY_TYPE_REGION  MemoryType = "MemoryTypeRegion"
	MEMORYTYPE_MEMORY_TYPE_LOCAL   MemoryType = "MemoryTypeLocal"
	MEMORYTYPE_MEMORY_TYPE_CXL     MemoryType = "MemoryTypeCXL"
)

type TypeMemoryRegion struct {
	MemoryId        string // ID for target memory region
	PortId          string // ID from the URI that identifies the physical port device that this memory region is connected to, if any
	LogicalDeviceId string // ID from the URI that identifies the logical device that this memory region is connected to, if any
	Status          string // A response string
	Type            MemoryType
	SizeMiB         int32 // A mebibyte equals 2**20 or 1,048,576 bytes.
	Bandwidth       int32 // Memory bandwidth in the unit of GigaBytes per second
	Latency         int32 // Memory latency in the unit of nanosecond
}

type GetMemoryByIdResponse struct {
	MemoryRegion TypeMemoryRegion
	Status       string // The status of the request
	ServiceError error  // Any error returned by the service
}

type GetMemoryRequest struct {
}

type GetMemoryResponse struct {
	MemoryIds    []string // list of memory ids
	Status       string   // The status of the request
	ServiceError error    // Any error returned by the service
}

type GetBackendInfoResponse struct {
	BackendName string
	Version     string
	SessionId   string
}
