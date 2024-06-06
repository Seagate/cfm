// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package backend

import (
	"context"
	"fmt"

	//"github.com/Seagate/gofish"
	"github.com/google/uuid"
	"github.com/stmcginnis/gofish"
	"k8s.io/klog/v2"
)

// This variable holds http clients and backend information.
type connectionEntry struct {
	redfishConnection *gofish.APIClient // The active http.Client information is embedded here.
	name              string
	uuid              string
}

// A map to old the connections
var activeConnections map[string]*connectionEntry

// This should be called only once at program invocation.  Allocated space for the active connections.
func init() {
	activeConnections = make(map[string]*connectionEntry)
}

// Methods are passed the IP address, but the gofish library uses the URL.
// This function builds a string for the full URL from that IP address.
func (req *CreateSessionRequest) constructEndpoint() string {
	if req.Protocol != "" {
		return req.Protocol + "://" + req.Ip + ":" + fmt.Sprint(req.Port)
	} else {
		return fmt.Sprintf("%s://%s:%d", req.Protocol, req.Ip, req.Port)
	}
}

// CreateSession: Create a new session with an endpoint service
func (service *gofishService) CreateSession(ctx context.Context, settings *ConfigurationSettings, req *CreateSessionRequest) (*CreateSessionResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== CreateSession ======")
	logger.V(4).Info("create session", "request", req)

	// Passing the provided request settings into a format that library function gofish.Connect (in the client.go file) is expecting.
	// Most of these values are the current defaults for the gofish library.  Default values are being specified to avoid assumptions.
	config := gofish.ClientConfig{
		Endpoint:            req.constructEndpoint(), // Expected Endpoint value has the Endpoints URI crafted with leading 'http://'. This function takes IP provided by End-User and creates that string.
		Username:            req.Username,            // Username provided by End-User.
		Password:            req.Password,            // Password provided by End-User
		Session:             &gofish.Session{},       // Session is an optional session ID+token obtained from a previous session.
		Insecure:            req.Insecure,            // Should gofish library verify SSL certificates? If True, then do not verify SSL certificates.
		TLSHandshakeTimeout: 0,                       // Instructs the gofish library to use its default TLS Handshake timeout value.  This is set on line 104 of client.go and is currently set to 10 seconds.
		HTTPClient:          nil,                     // Optional pointer to which http.client to be used by gofish library for connections.
		DumpWriter:          nil,                     // Used by gofish library. Optional io.Writer to use. nil is default value.
		BasicAuth:           true,                    // If true tells gofish library's APIClient to use basic auth. If false tells gofish library's APIClient to use token based auth.
	}

	// Pass the connection information to another function that establishes a connection to the BMC.
	// Capture errors.
	c, err := gofish.Connect(config)

	name := ""

	if err == nil && c.Service != nil {
		name = fmt.Sprintf("%s:%s", c.Service.Description, c.Service.UUID)
	}

	uuid := uuid.New().String()

	// Storing "entry" into the activeConnections Map.
	activeConnections[uuid] = &connectionEntry{
		redfishConnection: c,
		uuid:              uuid,
		name:              name,
	}

	service.service.session = activeConnections[uuid]

	// If connecting to the BMC's Redfish API does not fail (succeeds), then pass
	// the connection pointer into map and pass response to calling function.
	if err == nil {
		// Error information returned to calling function via 'err'.
		return &CreateSessionResponse{SessionId: uuid, Status: "Success", ServiceError: nil}, err
	}

	// If the connection failed, then alert the higher levels.
	//  NOTE: this is my own failure response.  Copy from a tested one once we have one.
	return &CreateSessionResponse{SessionId: uuid, Status: "Failure", ServiceError: err}, err
}

// DeleteSession: Delete a session previously established with an endpoint service
func (service *gofishService) DeleteSession(ctx context.Context, settings *ConfigurationSettings, req *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== DeleteSession ======")
	logger.V(4).Info("delete session", "request", req)

	session := service.service.session.(*connectionEntry)

	// Retrieve the pointer of entry based on the UUID.
	// If the entry does not exist, then "okay" will be false.
	// If the entry does exist, then "okay" will be true.
	entry, okay := activeConnections[session.uuid]
	if okay {
		// Logout function does not return anything.
		entry.redfishConnection.Logout()
		// Remove entry from the map
		delete(activeConnections, session.uuid)
	}

	// As gofish.APIClient.Logout does not return anything, not even an error, there is nothing to verify the exit against. A return code is needed for the calling function, so 'nil' value is provided.
	//  Return Success.
	return &DeleteSessionResponse{Status: "Success", ServiceError: nil}, nil
}

// AllocateMemory: Create a new memory region on a memory appliance.
func (service *gofishService) AllocateMemory(ctx context.Context, settings *ConfigurationSettings, req *AllocateMemoryRequest) (*AllocateMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AllocateMemory ======")
	logger.V(4).Info("allocate memory", "request", req)
	return &AllocateMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// AllocateMemory: Create a new memory region on a memory appliance.
func (service *gofishService) AllocateMemoryByResource(ctx context.Context, settings *ConfigurationSettings, req *AllocateMemoryByResourceRequest) (*AllocateMemoryByResourceResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AllocateMemoryByResource ======")
	logger.V(4).Info("allocate memory by resource", "request", req)
	return &AllocateMemoryByResourceResponse{Status: "Success", ServiceError: nil}, nil
}

// AssignMemory: Assign a memory region on a memory appliance to a CXL Host
func (service *gofishService) AssignMemory(ctx context.Context, settings *ConfigurationSettings, req *AssignMemoryRequest) (*AssignMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AssignMemory ======")
	logger.V(4).Info("assign memory", "request", req)
	return &AssignMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// UnassignMemory: Unassign a memory region on a memory appliance from a CXL Host
func (service *gofishService) UnassignMemory(ctx context.Context, settings *ConfigurationSettings, req *UnassignMemoryRequest) (*UnassignMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AssignMemory ======")
	logger.V(4).Info("assign memory", "request", req)
	return &UnassignMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemoryResourceBlocks: Request Memory Resource Block information from the backends
func (service *gofishService) GetMemoryResourceBlocks(ctx context.Context, settings *ConfigurationSettings, req *MemoryResourceBlocksRequest) (*MemoryResourceBlocksResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryResourceBlocks ======")
	logger.V(4).Info("memory resource blocks", "request", req)
	return &MemoryResourceBlocksResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemoryResourceBlockById: Request a particular Memory Resource Block information by ID from the backends
func (service *gofishService) GetMemoryResourceBlockById(ctx context.Context, settings *ConfigurationSettings, req *MemoryResourceBlockByIdRequest) (*MemoryResourceBlockByIdResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryResourceBlockById ======")
	logger.V(4).Info("memory resource block by id", "request", req)
	return &MemoryResourceBlockByIdResponse{Status: "Success", ServiceError: nil}, nil
}

// GetPorts: Request Ports ids from the backend
func (service *gofishService) GetPorts(ctx context.Context, settings *ConfigurationSettings, req *GetPortsRequest) (*GetPortsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetPorts ======")
	logger.V(4).Info("GetPorts", "req", req)

	var response GetPortsResponse
	response.Status = "Success"
	response.ServiceError = nil
	response.PortIds = append(response.PortIds, "PO")
	response.PortIds = append(response.PortIds, "P1")

	return &response, nil
}

// GetHostPortPcieDevices:
func (service *gofishService) GetHostPortPcieDevices(ctx context.Context, settings *ConfigurationSettings, req *GetPortsRequest) (*GetPortsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetHostPortPcieDevices ======")
	logger.V(4).Info("gGetHostPortPcieDevices", "request", req)
	return &GetPortsResponse{Status: "Success", ServiceError: nil}, nil
}

// GetPortDetails: Request Ports info from the backend
func (service *gofishService) GetPortDetails(ctx context.Context, settings *ConfigurationSettings, req *GetPortDetailsRequest) (*GetPortDetailsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetPorts ======")
	logger.V(4).Info("GetPortDetails", "req", req)

	var response GetPortDetailsResponse
	response.Status = "Success"
	response.Status = "OK/Enabled"
	response.ServiceError = nil
	response.PortInformation.Id = "P0"
	response.PortInformation.PortProtocol = "CXL"
	response.PortInformation.CurrentSpeedGbps = 32
	response.PortInformation.StatusHealth = "OK"
	response.PortInformation.StatusState = "Enabled"

	return &response, nil
}

// GetHostPortSnById:
func (service *gofishService) GetHostPortSnById(ctx context.Context, settings *ConfigurationSettings, req *GetHostPortSnByIdRequest) (*GetHostPortSnByIdResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetHostPortSnById ======")
	logger.V(4).Info("get host port sn by id", "request", req)
	return &GetHostPortSnByIdResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemoryDevices: Delete memory region info by memory id
func (service *gofishService) GetMemoryDevices(ctx context.Context, settings *ConfigurationSettings, req *GetMemoryDevicesRequest) (*GetMemoryDevicesResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryDevices ======")
	logger.V(4).Info("get memory devices", "request", req)
	return &GetMemoryDevicesResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemoryDeviceDetails: Get a specific memory region info by memory id
func (service *gofishService) GetMemoryDeviceDetails(ctx context.Context, setting *ConfigurationSettings, req *GetMemoryDeviceDetailsRequest) (*GetMemoryDeviceDetailsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryDeviceDetails ======")
	logger.V(4).Info("get memory dev by id", "request", req)
	return &GetMemoryDeviceDetailsResponse{Status: "Success", ServiceError: nil}, nil
}

// FreeMemoryById: Delete memory region info by memory id
func (service *gofishService) FreeMemoryById(ctx context.Context, settings *ConfigurationSettings, req *FreeMemoryRequest) (*FreeMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== FreeMemory ======")
	logger.V(4).Info("free memory", "request", req)
	return &FreeMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemoryById: Get a specific memory region info by memory id
func (service *gofishService) GetMemoryById(ctx context.Context, setting *ConfigurationSettings, req *GetMemoryByIdRequest) (*GetMemoryByIdResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryById ======")
	logger.V(4).Info("get memory by id", "request", req)
	return &GetMemoryByIdResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemory: Get the list of memory ids for a particular endpoint
func (service *gofishService) GetMemory(ctx context.Context, settings *ConfigurationSettings, req *GetMemoryRequest) (*GetMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemory ======")
	logger.V(4).Info("get memory", "request", req)
	return &GetMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// GetBackendInfo: Get the information of this backend
func (service *gofishService) GetBackendInfo(ctx context.Context) *GetBackendInfoResponse {
	return &GetBackendInfoResponse{BackendName: "gofish", Version: "0.0"}
}
