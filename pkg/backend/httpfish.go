// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package backend

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

// Redfish and version
const redfish_serviceroot = "/redfish/v1/"

// Type of connection
const connection_type = "Memory"

type HTTPOperationType string

// Enumeration of HTTP operations
var HTTPOperation = struct {
	POST   HTTPOperationType
	GET    HTTPOperationType
	PUT    HTTPOperationType
	DELETE HTTPOperationType
}{
	POST:   "POST",
	GET:    "GET",
	PUT:    "PUT",
	DELETE: "DELETE",
}

type RedfishPath string

// This struct holds the Session and related info
type Session struct {
	redfishPaths      map[RedfishPath]string
	applianceResource *MemoryApplianceResources
	memoryChunkPath   map[string]string
	RedfishSessionId  string
	SessionId         string
	uuid              string
	client            *http.Client

	ip       string // IP address of the client
	port     uint16 // port address of the client
	username string // user name of the client
	password string // password of the client
	protocol string // http or https
	insecure bool   // ignore secure flag in https
	xToken   string // Authentication token
}

// Map of UUID to Session object
var activeSessions map[string]*Session

// Initialize the map of active sessions
func init() {
	activeSessions = make(map[string]*Session)
}

// This struct holds the Response and error
type Response struct {
	StatusCode   int
	header       http.Header
	err          error
	jsonRespBody map[string]interface{}
}

// Member function of Response that extracts a value from JSON
func (resp *Response) valueFromJSON(key string) (interface{}, error) {
	var jsonError error
	var value interface{}
	var exists bool
	if jsonError != nil {
		return value, fmt.Errorf("error reading JSON, error: %v", jsonError)
	}
	value, exists = resp.jsonRespBody[key]
	if !exists {
		return value, fmt.Errorf("key (%s) does not exist in JSON", key)
	}
	return value, nil
}

// Member function of Response that extracts an array from JSON
func (resp *Response) arrayFromJSON(key string) ([]interface{}, error) {
	value, err := resp.valueFromJSON(key)
	returnValue, ok := value.([]interface{})
	if !ok {
		return []interface{}{}, err
	}
	return returnValue, err
}

// Member function of Response that extracts a string from JSON
func (resp *Response) stringFromJSON(key string) (string, error) {
	value, err := resp.valueFromJSON(key)
	returnValue, ok := value.(string)
	if !ok {
		return "", err
	}
	return returnValue, err
}

// Member function of Response that extracts a float from JSON
func (resp *Response) floatFromJSON(key string) (float64, error) {
	value, err := resp.valueFromJSON(key)
	if err != nil {
		return 0, err
	}
	returnValue, ok := value.(float64)
	if !ok {
		return 0, err
	}
	return returnValue, err
}

// Member function of Response that extracts a boolean from JSON
func (resp *Response) booleanFromJSON(key string) (bool, error) {
	value, err := resp.valueFromJSON(key)
	returnValue, ok := value.(bool)
	if !ok {
		return false, err
	}
	return returnValue, err
}

// Member function of Response that extracts a string from JSON
func (resp *Response) odataStringFromJSON(key string) (string, error) {
	value, err := resp.valueFromJSON(key)
	if err != nil {
		return "", err
	}
	returnValue, ok := value.(map[string]interface{})["@odata.id"].(string)
	if !ok {
		return "", err
	}
	return returnValue, err
}

// Return value of function fixedLengthSubArrayWithReservedStatus
type subSet struct {
	reserved bool
	subArr   []int
}

// Get the fixed length array, the length equal to the blocks needed. And Check the reserved status of each array
// For example, an original array is {1,2,3,4}, 2-length subArray with reserved status should look like:
/*
{
		{
	reserved: true
	subArr: {1,2}
	}
		{
	reserved: true
	subArr: {2,3}
	}
		{
	reserved: false
	subArr: {3,4}
	}

}
*/
func (res *MemoryApplianceResources) fixedLengthSubArrayWithReservedStatus(length int, originArray []int) []subSet {
	var results []subSet
	var subArr []int
	for i := 0; i <= len(originArray)-length; i++ {
		var subArrayCollection subSet
		subArr = originArray[i : i+length]
		for _, i := range subArr {
			if res.resourceBlockArray[i].reserved == true {
				subArrayCollection.reserved = true
				break
			}
		}
		subArrayCollection.subArr = subArr
		results = append(results, subArrayCollection)
	}
	return results
}

// Pick up resource blocks from one DIMM, make sure resource blocks are contigious and not reserved
func (res *MemoryApplianceResources) pickNeededBlocksFromOneDIMM(neededBlocks int, DimmIndex int) []int {
	var indexList []int
	startBlockIndex := res.numberOfBlocksInEachDIMM * DimmIndex

	var originArray []int
	for i := startBlockIndex; i < startBlockIndex+res.numberOfBlocksInEachDIMM; i++ {
		originArray = append(originArray, i)
	}

	NeededBlocks := res.fixedLengthSubArrayWithReservedStatus(neededBlocks, originArray)
	for _, i := range NeededBlocks {
		if i.reserved == false {
			indexList = i.subArr
			break
		}
	}

	return indexList
}

// Member function of Response that returns odata of the member by index
func (resp *Response) memberOdataIndex(i int) (string, error) {
	members, err := resp.arrayFromJSON("Members")
	if len(members) < i {
		return "", fmt.Errorf("index %d out of array range", i)
	}

	if err == nil {
		return members[i].(map[string]interface{})["@odata.id"].(string), nil
	}
	return "", err
}

// Member function of Response that returns odata of the member array
func (resp *Response) memberOdataArray() ([]string, error) {
	var returnValue []string
	members, err := resp.arrayFromJSON("Members")
	if err != nil {
		return []string{""}, err
	}
	for _, m := range members {
		returnValue = append(returnValue, m.(map[string]interface{})["@odata.id"].(string))

	}

	return returnValue, nil
}

func (resp *Response) isMemoryResourceBlock() bool {
	blockTypes, err := resp.arrayFromJSON("ResourceBlockType")

	if err != nil {
		return false
	}

	for _, blockType := range blockTypes {
		if blockType == "Memory" {
			return true
		}
	}

	return false
}

// Member function of Session that performs the HTTP operation on the specified path
// Operation can be POST, GET, PUT, or DELETE
// Path is the Redfish URI
// Returns Response object
func (session *Session) query(operation HTTPOperationType, path string) Response {
	return session.queryWithJSON(operation, path, nil)
}

// Member function of Session that performs the HTTP operation on the specified path
// Operation can be POST, or PUT
// Path is the Redfish URI
// Returns Response object
func (session *Session) queryWithJSON(operation HTTPOperationType, path string, jsonData map[string]interface{}) Response {
	var response Response
	url := fmt.Sprintf("%s://%s:%d%s", session.protocol, session.ip, session.port, path)

	jsonByteData, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("HTTP: Error creating request")
		response.err = fmt.Errorf("http error: input json marshal fail")
		return response
	}

	request, err := http.NewRequest(string(operation), url, bytes.NewBuffer(jsonByteData))
	if err != nil {
		fmt.Println("HTTP: Error creating request")
		response.err = fmt.Errorf("http error: Error creating request")
		return response
	}
	if session.xToken != "" {
		request.Header.Set("X-Auth-Token", session.xToken)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if session.client == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: session.insecure},
		}

		session.client = &http.Client{Transport: tr, Timeout: 10 * time.Second}
	}
	httpresponse, err := session.client.Do(request)
	if err != nil {
		fmt.Println("HTTP: Error sending request", url)
		response.err = fmt.Errorf("http error: Error sending request")
		return response
	}
	defer httpresponse.Body.Close()

	response.StatusCode = httpresponse.StatusCode
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		response.err = fmt.Errorf(http.StatusText(response.StatusCode))
	}
	response.header = httpresponse.Header

	err = json.NewDecoder(httpresponse.Body).Decode(&response.jsonRespBody)
	if err != nil && httpresponse.ContentLength != 0 {
		fmt.Println("HTTP: Error decode json response")
		response.err = fmt.Errorf("http error: Error decode json response")
		return response
	}

	// check for error due to session timeout ( service would return error code 401)
	if session.xToken != "" && response.StatusCode == http.StatusUnauthorized {
		// Re-authenticate
		fmt.Print("Redfish session might have timed out. Re-authenticate. Warning! infinite loop might occur if the issue is from the redfish server.\n")
		session.client = nil
		session.xToken = ""
		err := session.auth()
		if err == nil {
			path = session.buildPath(SessionServiceKey, session.RedfishSessionId)
			response = session.queryWithJSON(operation, path, jsonData)
		}
	}

	return response
}

const (
	SystemsKey               = "Systems"
	SystemMemoryChunksCXLKey = "SystemMemoryChunksCXL"
	SystemMemoryDomainKey    = "SystemMemoryDomain"
	SystemMemoryChunksKey    = "SystemMemoryChunks"
	ChassisKey               = "Chassis"
	ChassisMemoryKey         = "ChassisMemory"
	ChassisPcieDevKey        = "ChassisPcieDev"
	FabricKey                = "Fabric"
	FabricZonesKey           = "FabricZones"
	FabricEndpointsKey       = "FabricEndpoints"
	FabricConnectionsKey     = "FabricConnections"
	FabricSwitchesKey        = "FabricSwitches"
	FabricPortsKey           = "FabricPorts"
	ResourceZonesKey         = "ResourceZones"
	ResourceBlocksKey        = "ResourceBlocks"
	PostResourceKey          = "PostResource"
	SessionServiceKey        = "SessionService"
)

// Session.pathInit(): initialize all usable paths for redfish and constant info
func (session *Session) pathInit() {
	var err error
	var path string
	// Service root
	serviceroot_response := session.query(HTTPOperation.GET, redfish_serviceroot)
	session.uuid, err = serviceroot_response.stringFromJSON("UUID")
	if session.uuid == "" || err != nil {
		session.uuid = uuid.New().String()
	}

	// System
	path, err = serviceroot_response.odataStringFromJSON("Systems")

	if err == nil {
		response := session.query(HTTPOperation.GET, path)
		session.redfishPaths[SystemsKey], err = response.memberOdataIndex(0)
		if err == nil {
			// System/{systemId}/MemoryDomains/{memoryDomainId}/MemoryChunks
			response = session.query(HTTPOperation.GET, session.redfishPaths[SystemsKey]+"/MemoryDomains")
			DomainArray, err2 := response.memberOdataArray()
			if err2 == nil {
				for _, domainPath := range DomainArray {
					if strings.Contains(domainPath, "CXL") {
						session.redfishPaths[SystemMemoryChunksCXLKey] = domainPath + "/MemoryChunks"
					} else {
						session.redfishPaths[SystemMemoryDomainKey] = domainPath
						session.redfishPaths[SystemMemoryChunksKey] = session.redfishPaths[SystemMemoryDomainKey] + "/MemoryChunks"

					}
				}
			} else {
				fmt.Println("init SystemMemoryChunks path err", err)
			}
			session.redfishPaths[SystemMemoryDomainKey], err = response.memberOdataIndex(0)
			if err == nil {
				session.redfishPaths[SystemMemoryChunksKey] = session.redfishPaths[SystemMemoryDomainKey] + "/MemoryChunks"
			} else {
				fmt.Println("init SystemMemoryChunks path err", err)
			}
		} else {
			fmt.Println("init SystemMemoryDomain path err", err)
		}

	} else {
		fmt.Println("init Systems path err", err)
	}

	// Chassis
	path, err = serviceroot_response.odataStringFromJSON("Chassis")

	if err == nil {
		response := session.query(HTTPOperation.GET, path)
		session.redfishPaths[ChassisKey], err = response.memberOdataIndex(0)

		if err == nil {
			session.redfishPaths[ChassisMemoryKey] = session.redfishPaths[ChassisKey] + "/Memory"
			response := session.query(HTTPOperation.GET, session.redfishPaths[ChassisKey])
			session.redfishPaths[ChassisPcieDevKey], err = response.odataStringFromJSON("PCIeDevices")
			if err != nil {
				fmt.Println("init ChassisPcieDev path err", err)
			}
		} else {
			fmt.Println("init ChassisMemory path err", err)
		}

	} else {
		fmt.Println("init Chassis path err", err)
	}

	// Fabric
	path, err = serviceroot_response.odataStringFromJSON("Fabrics")

	if err == nil {
		response := session.query(HTTPOperation.GET, path)
		session.redfishPaths[FabricKey], err = response.memberOdataIndex(0)

		if err == nil {
			session.redfishPaths[FabricZonesKey] = session.redfishPaths[FabricKey] + "/Zones"
			session.redfishPaths[FabricEndpointsKey] = session.redfishPaths[FabricKey] + "/Endpoints"
			session.redfishPaths[FabricConnectionsKey] = session.redfishPaths[FabricKey] + "/Connections"
			response = session.query(HTTPOperation.GET, session.redfishPaths[FabricKey]+"/Switches")
			session.redfishPaths[FabricSwitchesKey], err = response.memberOdataIndex(0)
			if err == nil {
				session.redfishPaths[FabricPortsKey] = session.redfishPaths[FabricSwitchesKey] + "/Ports"
			} else {
				fmt.Println("init FabricPorts path err", err)
			}
		} else {
			fmt.Println("init FabricZones path err", err)
		}
	} else {
		fmt.Println("init Fabrics path err", err)
	}

	// CompositionService
	session.redfishPaths[ResourceZonesKey] = redfish_serviceroot + "CompositionService/ResourceZones"
	session.redfishPaths[ResourceBlocksKey] = redfish_serviceroot + "CompositionService/ResourceBlocks"
	session.redfishPaths[PostResourceKey] = redfish_serviceroot + "Systems"

	// session service
	session.redfishPaths[SessionServiceKey] = redfish_serviceroot + "SessionService/Sessions"

}

// Session.BuildPath: add a string to existing redfish path ( such as a member )
func (session *Session) buildPath(base RedfishPath, addon string) string {
	return fmt.Sprintf("%s/%s", session.redfishPaths[base], addon)
}

// Session.auth(): authenticate the redfish session with user credential.
func (session *Session) auth() error {
	authData := make(map[string]interface{})
	authData["UserName"] = session.username
	authData["Password"] = session.password
	response := session.queryWithJSON(HTTPOperation.POST, redfish_serviceroot+"SessionService/Sessions", authData)

	if response.err == nil {
		session.xToken = response.header.Get("X-Auth-Token")
		session.RedfishSessionId, response.err = response.stringFromJSON("Id")
	}
	return response.err
}

// CreateSession: Create a new session with an endpoint service
func (service *httpfishService) CreateSession(ctx context.Context, settings *ConfigurationSettings, req *CreateSessionRequest) (*CreateSessionResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== CreateSession ======")
	logger.V(4).Info("create session", "request", req)

	var session = Session{
		redfishPaths:    make(map[RedfishPath]string),
		memoryChunkPath: make(map[string]string),
		uuid:            "",

		ip:       req.Ip,
		port:     uint16(req.Port),
		username: req.Username,
		password: req.Password,
		protocol: req.Protocol,
		insecure: req.Insecure,
	}

	err := session.auth()
	if err != nil {
		var tlsCertErr *tls.CertificateVerificationError
		protocolErrStr := "http: server gave HTTP response to HTTPS client" // match hardcoded error message from net/http package

		if req.Protocol == "https" && strings.Contains(err.Error(), protocolErrStr) { // http server with https request
			logger.V(2).Info("Create Session protocol retry", "Error", err.Error())
			req.Protocol = "http"
			return service.CreateSession(ctx, settings, req)
		} else if req.Insecure == false && errors.As(errors.Unwrap(err), &tlsCertErr) {
			logger.V(2).Info("Create Session SSL retry", "Error", err.Error())
			req.Insecure = true
			return service.CreateSession(ctx, settings, req)
		} else {
			return &CreateSessionResponse{SessionId: session.SessionId, Status: "Failure", ServiceError: err}, err
		}
	}
	logger.V(4).Info("Session Created", "X-Auth-Token", session.xToken, "RedfishSessionId", session.RedfishSessionId)

	// walk redfish path and store the path in session struct
	session.pathInit()

	// Create DeviceId from uuid
	session.SessionId = session.uuid

	_, exist := activeSessions[session.SessionId]
	if exist {
		return &CreateSessionResponse{SessionId: session.SessionId, Status: "Duplicated", ServiceError: nil}, fmt.Errorf("endpoint already exist")
	}
	activeSessions[session.SessionId] = &session
	service.service.session = &session

	return &CreateSessionResponse{SessionId: session.SessionId, Status: "Success", ServiceError: nil}, nil
}

// DeleteSession: Delete a session previously established with an endpoint service
func (service *httpfishService) DeleteSession(ctx context.Context, settings *ConfigurationSettings, req *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== DeleteSession ======")
	logger.V(4).Info("delete session", "request", req)

	session := service.service.session.(*Session)

	response := session.query(HTTPOperation.DELETE, session.buildPath(SessionServiceKey, session.RedfishSessionId))

	// CloseIdleConnections closes the idle connections that a session client may make use of
	// session.CloseIdleConnections()
	delete(activeSessions, session.SessionId)

	// Let user know of delete backend failure.
	if response.err != nil {
		return &DeleteSessionResponse{SessionId: session.SessionId, IpAddress: session.ip, Port: int32(session.port), Status: "Failure", ServiceError: response.err}, response.err
	}

	return &DeleteSessionResponse{SessionId: session.SessionId, IpAddress: session.ip, Port: int32(session.port), Status: "Success", ServiceError: nil}, nil
}

// This struct holds the detail info of a specific resource block
type ResourceBlockItem struct {
	id               string
	index            int
	miBSize          float64
	memoryId         string
	memoryUri        string
	dimm             int
	dimmId           string
	dimmUri          string
	dimmMiB          float64
	compositionState string
	reserved         bool
}

// This struct holds the collection of all resource blocks
type MemoryApplianceResources struct {
	numberOfDIMMs            int
	numberOfBlocksInEachDIMM int
	capacityOfOneBlock       float64
	numberOfBlocks           int
	resourceBlockArray       []ResourceBlockItem
}

// Helper function to extracte id from odata.id
func getIdFromOdataId(odataId string) string {
	var id string
	components := strings.Split(odataId, "/")
	id = components[len(components)-1]
	return id
}

// Acquire all resource blocks from a memory appliance and initialize local data structure
func (session *Session) ResourcesInit(qos QoS) (*MemoryApplianceResources, error) {
	var resources MemoryApplianceResources
	var resourcesItem []ResourceBlockItem // store info of resourceBlockArray to return

	//TODO: Implement QoS for 1, 2 and 8 dimms
	if int(qos) != 4 {
		return nil, fmt.Errorf("qos=%d not yet supported. Only qos=4 supported", int(qos))
	}

	// Acquire the info of dimms and assign the number of it to resources.numberOfDIMMs
	responseOfDimms := session.query(HTTPOperation.GET, session.redfishPaths[ChassisMemoryKey])
	if responseOfDimms.err != nil {
		return nil, fmt.Errorf("failed to get dimm info")
	}

	collectionMembers, dimmErr := responseOfDimms.arrayFromJSON("Members")
	if dimmErr != nil {
		return nil, dimmErr
	}
	if len(collectionMembers) != 4 {
		return nil, fmt.Errorf("invalid dimm count found (%d)", len(collectionMembers))
	}

	resources.numberOfDIMMs = len(collectionMembers)
	uriOfResources := session.redfishPaths[ResourceBlocksKey]
	// Get collection of all resource blocks
	responseOfResources := session.query(HTTPOperation.GET, uriOfResources)
	if responseOfResources.err != nil {
		return nil, fmt.Errorf("failed to get resource blocks")
	}

	odataOfResources, err := responseOfResources.memberOdataArray()
	if err != nil {
		return nil, fmt.Errorf("failed to get resource block urls")
	}

	resources.numberOfBlocks = len(odataOfResources)
	resources.numberOfBlocksInEachDIMM = resources.numberOfBlocks / resources.numberOfDIMMs

	startIndex := 0

	for _, uriOfResource := range odataOfResources {
		// Get a specific resource block
		responseOfResource := session.query(HTTPOperation.GET, uriOfResource)
		if responseOfResource.err != nil {
			return nil, fmt.Errorf("failed to get resource block")
		}
		// Declare a template ResourceBlock to store info for one resource and append to resourceItem
		var OneResourceBlock ResourceBlockItem

		// resource block id should from the last compoenent of the @odata.id property
		OneResourceBlock.id = getIdFromOdataId(uriOfResource)
		// Increase the index for resources
		OneResourceBlock.index = startIndex
		startIndex++

		// Get reserved and compositionState info
		compositionStatus, _ := responseOfResource.valueFromJSON("CompositionStatus")
		OneResourceBlock.reserved = compositionStatus.(map[string]interface{})["Reserved"].(bool)
		OneResourceBlock.compositionState = compositionStatus.(map[string]interface{})["CompositionState"].(string)

		memoryElements, _ := responseOfResource.arrayFromJSON("Memory")

		for _, memoryElement := range memoryElements {
			OneResourceBlock.memoryUri = memoryElement.(map[string]interface{})["@odata.id"].(string)
			OneResourceBlock.memoryId = getIdFromOdataId(OneResourceBlock.memoryUri)
			blockDimm := session.query(HTTPOperation.GET, OneResourceBlock.memoryUri)
			OneResourceBlock.miBSize, _ = blockDimm.floatFromJSON("CapacityMiB")

			blockdimmId, _ := blockDimm.stringFromJSON("Id")
			// blockdimmId looks like "block0dimm0", extract "dimm" and the floolwing number
			OneResourceBlock.dimmId = blockdimmId[strings.Index(blockdimmId, "dimm"):]
		}

		OneResourceBlock.dimmUri = session.buildPath(ChassisMemoryKey, OneResourceBlock.dimmId)
		responseOfDimm := session.query(HTTPOperation.GET, OneResourceBlock.dimmUri)
		OneResourceBlock.dimmMiB, _ = responseOfDimm.floatFromJSON("CapacityMiB")

		resourcesItem = append(resourcesItem, OneResourceBlock)
	}

	// Assume all resource block have equal capacity (MiB)
	resources.capacityOfOneBlock = resourcesItem[0].miBSize
	// Checkthe the capacity of one block, if it is less than 0, can't round up targetMebibytes in getResourceBlocks
	if resources.capacityOfOneBlock <= 0 {
		return nil, fmt.Errorf("resource block is empty")
	}

	resources.resourceBlockArray = resourcesItem

	return &resources, nil
}

// Helper function to make sure there are no empty resource array in BlockArrayFromDIMMsCollection
func hasEmptyArray(arr [][]int) bool {
	for _, subArr := range arr {
		if len(subArr) == 0 {
			return true
		}
	}
	return false
}

// Takes the bladeId to use and mebibytes of capacity needed
// Returns the allocatedMebibytes and the JSON of resource blocks used in allocating memory
func getResourceBlocks(session *Session, targetMebibytes int32, qos QoS) (int32, []map[string]interface{}, error) {
	var returnValue []map[string]interface{}

	memoryApplianceResources, err := session.ResourcesInit(qos)
	if err != nil {
		return 0, nil, err
	}
	session.applianceResource = memoryApplianceResources

	if memoryApplianceResources.numberOfBlocks == 0 || targetMebibytes <= 0 {
		return 0, nil, fmt.Errorf("unable to allocate memory for blade [%s]", session.SessionId)
	}

	// Store the id of to be allocated resource blocks
	var resourceListId []int
	// Round up the target size to the nearest integer multiply of the resource block size
	neededBlocks := int(math.Ceil(float64(targetMebibytes) / memoryApplianceResources.capacityOfOneBlock))
	// Round up the neededBlocks to the nearest integer multiply of number of dimm
	blocksFromEachDIMM := int(math.Ceil(float64(neededBlocks) / float64(memoryApplianceResources.numberOfDIMMs)))

	// Utility BlockArrayFromDIMMsCollection to store all index from all dimms, for example:
	// 8 blocks from 4 DIMMs looks like: [0,1][4,5][9,10][13,14]
	var BlockArrayFromDIMMsCollection [][]int

	resourceBlockCollection := session.query(HTTPOperation.GET, session.redfishPaths[ResourceBlocksKey])
	collectionMembers, _ := resourceBlockCollection.arrayFromJSON("Members")

	// Only 4-way interleave accepted, make sure there are enough blocks in each dimm
	if blocksFromEachDIMM <= memoryApplianceResources.numberOfBlocksInEachDIMM {

		for i := 0; i < memoryApplianceResources.numberOfDIMMs; i++ {
			BlockArrayFromDIMMsCollection = append(BlockArrayFromDIMMsCollection, session.applianceResource.pickNeededBlocksFromOneDIMM(blocksFromEachDIMM, i))
		}

		// if exist at least one empty pickNeededBlocksFromOneDIMM(sub array of BlockArrayFromDIMMsCollection ) in BlockArrayFromDIMMsCollection
		// that means at least one DIMM could not provide enough blocks because the available block(s) has been reserved
		if hasEmptyArray(BlockArrayFromDIMMsCollection) {
			return 0, nil, fmt.Errorf("unable to allocate memory for blade [%s]", session.SessionId)
		}

		for _, blocksCollection := range BlockArrayFromDIMMsCollection {
			resourceListId = append(resourceListId, blocksCollection...)
		}

	} else {
		return 0, nil, fmt.Errorf("unable to allocate memory for blade [%s]", session.SessionId)
	}

	for _, i := range resourceListId {
		returnValue = append(returnValue, collectionMembers[i].(map[string]interface{}))
	}
	allocatedMebibytes := int32(blocksFromEachDIMM * memoryApplianceResources.numberOfDIMMs * int(memoryApplianceResources.capacityOfOneBlock))
	return allocatedMebibytes, returnValue, nil
}

// AllocateMemory: Create a new memory region.
func (service *httpfishService) AllocateMemory(ctx context.Context, settings *ConfigurationSettings, req *AllocateMemoryRequest) (*AllocateMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AllocateMemory ======")
	logger.V(4).Info("allocate memory", "request", req)

	session := service.service.session.(*Session)

	// allocated memory size returned here may not be the same as req.SizeMiB due to rounding
	allocatedMebibytes, resourceBlockInterface, calcErr := getResourceBlocks(session, req.SizeMiB, req.Qos)
	if allocatedMebibytes == 0 || calcErr != nil {
		newErr := fmt.Errorf("problem during resource capacity calculations: %w", calcErr)
		logger.Error(newErr, "failure: allocate memory", "allocatedMebibytes", allocatedMebibytes, "req", req)
		return &AllocateMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	jsonData := make(map[string]interface{})
	jsonData["Links"] = make(map[string]interface{})
	jsonData["Links"].(map[string]interface{})["ResourceBlocks"] = resourceBlockInterface

	response := session.queryWithJSON(HTTPOperation.POST, session.redfishPaths[PostResourceKey], jsonData)
	if response.err != nil {
		newErr := fmt.Errorf("backend session post failure(%s): %w", session.redfishPaths[PostResourceKey], response.err)
		logger.Error(newErr, "failure: allocate memory", "req", req, "allocatedMebibytes", allocatedMebibytes)
		return &AllocateMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	//extract the memorychunk Id
	uriOfMemorychunkId := response.header.Values("Location")
	memoryId := getIdFromOdataId(uriOfMemorychunkId[0])
	session.memoryChunkPath[memoryId] = uriOfMemorychunkId[0]

	return &AllocateMemoryResponse{SizeMiB: allocatedMebibytes, MemoryId: memoryId, Status: "Success", ServiceError: nil}, nil
}

// AllocateMemoryByResource: Create a new memory region using user-specified resource blocks
func (service *httpfishService) AllocateMemoryByResource(ctx context.Context, settings *ConfigurationSettings, req *AllocateMemoryByResourceRequest) (*AllocateMemoryByResourceResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AllocateMemoryByResource ======")
	logger.V(4).Info("allocate memory by resource", "request", req)

	session := service.service.session.(*Session)

	var backendResourceUris []string
	for _, resourceId := range req.MemoryResoureIds {
		// check for a valid resource
		backendResourceUri := session.buildPath(ResourceBlocksKey, resourceId)
		response := session.query(HTTPOperation.GET, backendResourceUri)
		if response.err != nil {
			newErr := fmt.Errorf("backend session get failure(%s): %w", backendResourceUri, response.err)
			logger.Error(newErr, "failure: allocate memory by resource", "req", req)
			return &AllocateMemoryByResourceResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		// check for a valid composition state
		compositionStatus, statusErr := response.valueFromJSON("CompositionStatus")
		if statusErr != nil {
			newErr := fmt.Errorf("CompositionStatus not found(%s): %w", backendResourceUri, statusErr)
			logger.Error(newErr, "failure: allocate memory by resource", "req", req)
			return &AllocateMemoryByResourceResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		compositionState := compositionStatus.(map[string]interface{})["CompositionState"].(string)
		reserved := compositionStatus.(map[string]interface{})["Reserved"].(bool)

		resourceState := findResourceState(&compositionState, reserved)
		if *resourceState != ResourceUnused {
			newErr := fmt.Errorf("resource not accessible(%s): session [%s] state [%s] ", backendResourceUri, session.SessionId, *resourceState)
			logger.Error(newErr, "failure: allocate memory by resource")
			return &AllocateMemoryByResourceResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		// save off the backend resource uri, used later in POST
		backendResourceUris = append(backendResourceUris, backendResourceUri)
	}

	var listOfUriMaps []map[string]interface{}
	for _, uri := range backendResourceUris {
		newMap := make(map[string]interface{})
		newMap["@odata.id"] = uri
		listOfUriMaps = append(listOfUriMaps, newMap)
	}

	jsonData := make(map[string]interface{})
	jsonData["Links"] = make(map[string]interface{})
	jsonData["Links"].(map[string]interface{})["ResourceBlocks"] = listOfUriMaps

	response := session.queryWithJSON(HTTPOperation.POST, session.redfishPaths[PostResourceKey], jsonData)
	if response.err != nil {
		newErr := fmt.Errorf("backend session post failure(%s): %w", session.redfishPaths[PostResourceKey], response.err)
		logger.Error(newErr, "failure: allocate memory by resource", "req", req)
		return &AllocateMemoryByResourceResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	//extract the memorychunk Id
	uriOfMemorychunkId := response.header.Values("Location")
	memoryId := getIdFromOdataId(uriOfMemorychunkId[0])
	session.memoryChunkPath[memoryId] = uriOfMemorychunkId[0]

	return &AllocateMemoryByResourceResponse{MemoryId: memoryId, Status: "Success", ServiceError: nil}, nil
}

// Extract the corresponding endpoint uri of the input port
func (session *Session) getEndpointUriFromPort(portId string) (*string, error) {
	// Get the target port
	portResponse := session.query(HTTPOperation.GET, session.buildPath(FabricPortsKey, portId))
	if portResponse.err != nil {
		return nil, fmt.Errorf("failed to get port [%s]: %w", portId, portResponse.err)
	}

	// Extract the port's corresponding endpoint URI
	links, _ := portResponse.valueFromJSON("Links")
	endptUri := links.(map[string]interface{})["AssociatedEndpoints"].([]interface{})[0].(map[string]interface{})["@odata.id"].(string)

	return &endptUri, nil
}

// Queries the endpoint uri to see if the endpoint is avaliable.  Error if connection found.
func (session *Session) isEndpointAvailable(endptUri string) error {

	// Get the target endpoint
	endptResponse := session.query(HTTPOperation.GET, endptUri)
	if endptResponse.err != nil {
		return fmt.Errorf("failed to get endpoint [%s] for uri [%s]: %w", endptUri, endptUri, endptResponse.err)
	}

	// Search for any existing connections
	endptLinks, _ := endptResponse.valueFromJSON("Links")
	endptConnections := endptLinks.(map[string]interface{})["Connections"].([]interface{})
	if len(endptConnections) > 0 {
		// Any connections indicate that the port is currently in use
		endptConnectionUri := endptConnections[0].(map[string]interface{})["@odata.id"].(string)
		return fmt.Errorf("existing connection found [%s] for endpt [%s]", endptConnectionUri, endptUri)
	}

	return nil
}

// Get the available endpoint, search all endpoints and select an available one
// Not used right but keep it in case needed in the future
func (session *Session) getAvailableEndpoint() (*string, error) {
	var AvailableEndpoint string

	// Get all endpoints
	responseOfEndpoints := session.query(HTTPOperation.GET, session.redfishPaths[FabricEndpointsKey])
	if responseOfEndpoints.err != nil {
		return nil, fmt.Errorf("failed to get endpoints")
	}

	// Get all member(endpoint) uri
	memberCollection, err := responseOfEndpoints.memberOdataArray()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint")
	}

	// Get a specific endpoint
	for _, uriOfEndpoint := range memberCollection {
		responseOfEndpoint := session.query(HTTPOperation.GET, uriOfEndpoint)
		if responseOfEndpoint.err != nil {
			return nil, fmt.Errorf("failed to get endpoint")
		}

		// Check the "ConnectedEntities" of a endpoint, empty means available
		connectedEntities, _ := responseOfEndpoint.arrayFromJSON("ConnectedEntities")
		if len(connectedEntities) == 0 {
			AvailableEndpoint = uriOfEndpoint
			break
		}
	}

	return &AvailableEndpoint, nil
}

// AssignMemory: Establish(Assign) a connection between a memory region and a local hardware port
func (service *httpfishService) AssignMemory(ctx context.Context, settings *ConfigurationSettings, req *AssignMemoryRequest) (*AssignMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== AssignMemory ======")
	logger.V(4).Info("assign memory", "request", req)

	session := service.service.session.(*Session)

	jsonData := make(map[string]interface{})

	jsonData["ConnectionType"] = connection_type

	jsonData["MemoryChunkInfo"] = make([]map[string]interface{}, 1)
	jsonData["MemoryChunkInfo"].([]map[string]interface{})[0] = make(map[string]interface{})
	jsonData["MemoryChunkInfo"].([]map[string]interface{})[0]["MemoryChunk"] = make(map[string]interface{})
	jsonData["MemoryChunkInfo"].([]map[string]interface{})[0]["MemoryChunk"].(map[string]interface{})["@odata.id"] =
		session.buildPath(SystemMemoryChunksKey, string(req.MemoryId))

	jsonData["Links"] = make(map[string]interface{})

	// Extract and validate the availability of the target endpoint
	// Assume there is only one target endpoint, so we always check targetEndpoint[0]
	uriOfTargetEndpoint, errOfTargetEndpoint := session.getEndpointUriFromPort(req.PortId)
	if errOfTargetEndpoint != nil {
		newErr := fmt.Errorf("backend session failure(%s): %w", session.redfishPaths[FabricConnectionsKey], errOfTargetEndpoint)
		logger.Error(newErr, "failure: assign memory", "req", req)
		return &AssignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	availableErr := session.isEndpointAvailable(*uriOfTargetEndpoint)
	if availableErr != nil {
		newErr := fmt.Errorf("backend session failure(%s): %w", session.redfishPaths[FabricConnectionsKey], availableErr)
		logger.Error(newErr, "failure: assign memory", "req", req)
		return &AssignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	jsonDataOfTargetEndpoint := make([]map[string]interface{}, 1)
	jsonDataOfTargetEndpoint[0] = map[string]interface{}{
		"@odata.id": uriOfTargetEndpoint,
	}

	jsonData["Links"].(map[string]interface{})["TargetEndpoints"] = jsonDataOfTargetEndpoint

	// Assign memory
	response := session.queryWithJSON(HTTPOperation.POST, session.redfishPaths[FabricConnectionsKey], jsonData)
	if response.err != nil {
		newErr := fmt.Errorf("backend session post failure(%s): %w", session.redfishPaths[FabricConnectionsKey], response.err)
		logger.Error(newErr, "failure: assign memory", "req", req)
		return &AssignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	return &AssignMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// UnassignMemory: Delete(Unassign) a connection between a memory region and it's local hardware port.  If no connection found, no action taken.
func (service *httpfishService) UnassignMemory(ctx context.Context, settings *ConfigurationSettings, req *UnassignMemoryRequest) (*UnassignMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== UnassignMemory ======")
	logger.V(4).Info("unassign memory", "request", req)

	session := service.service.session.(*Session)

	// Find available connections
	responseGetAllConnections := session.query(HTTPOperation.GET, session.redfishPaths[FabricConnectionsKey])
	if responseGetAllConnections.err != nil {
		newErr := fmt.Errorf("backend session get failure(%s): %w", session.redfishPaths[FabricConnectionsKey], responseGetAllConnections.err)
		logger.Error(newErr, "failure: unassign memory", "req", req)
		return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	connections, err := responseGetAllConnections.arrayFromJSON("Members")
	if err != nil {
		newErr := fmt.Errorf("response parsing failure('Members'): %w", err)
		logger.Error(newErr, "failure: unassign memory", "req", req)
		return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	// Search all session connections for the specified memoryId
	// If found, verify that requested portId is connected to memoryId before deleting the connection
	var foundMemory bool
	var foundPort bool
	for _, connection := range connections {
		connectionUri := connection.(map[string]interface{})["@odata.id"].(string)

		// Get one of the connections
		responseGetConnection := session.query(HTTPOperation.GET, connectionUri)
		if responseGetConnection.err != nil {
			newErr := fmt.Errorf("backend session get failure(%s): %w", connectionUri, responseGetConnection.err)
			logger.Error(newErr, "failure: unassign memory", "req", req)
			return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		mcInfos, mciErr := responseGetConnection.arrayFromJSON("MemoryChunkInfo")
		if mciErr != nil {
			newErr := fmt.Errorf("response parsing failure('MemoryChunkInfo'): %w", mciErr)
			logger.Error(newErr, "failure: unassign memory", "req", req)
			return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		// Search for memoryId on the current connection
		for _, info := range mcInfos {
			chunk := info.(map[string]interface{})["MemoryChunk"]
			chunkUri := chunk.(map[string]interface{})["@odata.id"].(string)
			if strings.Contains(chunkUri, req.MemoryId) {

				foundMemory = true

				//Verify that the requested portId also matches
				connectionLinks, _ := responseGetConnection.valueFromJSON("Links")
				connectionTargetEndpts := connectionLinks.(map[string]interface{})["TargetEndpoints"]
				endpts := connectionTargetEndpts.([]interface{})

				for _, endpt := range endpts {
					endptUri := endpt.(map[string]interface{})["@odata.id"].(string)
					responseGetEndpt := session.query(HTTPOperation.GET, endptUri)
					if responseGetEndpt.err != nil {
						newErr := fmt.Errorf("backend session get failure(%s): %w", endptUri, responseGetEndpt.err)
						logger.Error(newErr, "failure: unassign memory", "req", req)
						return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
					}

					endptLinks, _ := responseGetEndpt.valueFromJSON("Links")
					endptConnectedPorts := endptLinks.(map[string]interface{})["ConnectedPorts"]
					endptPorts := endptConnectedPorts.([]interface{})

					for _, port := range endptPorts {
						portUri := port.(map[string]interface{})["@odata.id"].(string)
						if strings.Contains(portUri, req.PortId) {
							foundPort = true

							// Delete the matching connection
							responseDeleteConnection := session.query(HTTPOperation.DELETE, connectionUri)
							if responseDeleteConnection.err != nil {
								newErr := fmt.Errorf("backend session delete failure(%s): %w", connectionUri, responseDeleteConnection.err)
								logger.Error(newErr, "failure: unassign memory", "req", req)
								return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
							}

							logger.V(4).Info("unassign memory: connection deleted", "connectionUri", connectionUri, "req", req)
							break
						}
					}

					if foundPort {
						break
					}
				}

				// Found the memoryId but didn't find the matching portId
				if !foundPort {
					newErr := fmt.Errorf("connection mismatch: portId(%s) not connected to memoryId(%s)", req.MemoryId, req.PortId)
					logger.Error(newErr, "failure: unassign memory", "req", req)
					return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
				}

				break
			}
		}

		if foundMemory {
			break
		}
	}

	if !foundMemory {
		newErr := fmt.Errorf("memoryId(%s) not found", req.MemoryId)
		logger.Error(newErr, "failure: unassign memory", "req", req)
		return &UnassignMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	return &UnassignMemoryResponse{Status: "Success", ServiceError: nil}, nil
}

// GetMemoryResourceBlocks: Request Memory Resource Block information from the backends
func (service *httpfishService) GetMemoryResourceBlocks(ctx context.Context, settings *ConfigurationSettings, req *MemoryResourceBlocksRequest) (*MemoryResourceBlocksResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryResourceBlocks ======")
	logger.V(4).Info("memory resource blocks", "request", req)

	memoryResources := make([]string, 0)

	session := service.service.session.(*Session)

	response := session.query(HTTPOperation.GET, session.redfishPaths[ResourceBlocksKey])

	if response.err != nil {
		return &MemoryResourceBlocksResponse{Status: "Failure", ServiceError: response.err}, response.err
	}

	resourceBlocks, _ := response.arrayFromJSON("Members")

	for _, resourceBlock := range resourceBlocks {
		blockDetails := session.query(HTTPOperation.GET, resourceBlock.(map[string]interface{})["@odata.id"].(string))

		if !(blockDetails.err == nil && blockDetails.isMemoryResourceBlock()) {
			continue
		}

		uri, _ := blockDetails.stringFromJSON("@odata.id")

		memoryResources = append(memoryResources, getIdFromOdataId(uri))
	}

	return &MemoryResourceBlocksResponse{MemoryResources: memoryResources, Status: "Success", ServiceError: nil}, nil
}

// GetMemoryResourceBlockById: Request a particular Memory Resource Block information by ID from the backends
func (service *httpfishService) GetMemoryResourceBlockById(ctx context.Context, settings *ConfigurationSettings, req *MemoryResourceBlockByIdRequest) (*MemoryResourceBlockByIdResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryResourceBlockById ======")
	logger.V(4).Info("memory resource block by id", "request", req)

	memoryResourceBlock := MemoryResourceBlock{
		CompositionStatus: MemoryResourceBlockCompositionStatus{},
		Id:                req.ResourceId,
	}

	session := service.service.session.(*Session)

	uri := session.buildPath(ResourceBlocksKey, req.ResourceId)
	response := session.query(HTTPOperation.GET, uri)
	if response.err != nil {
		newErr := fmt.Errorf("backend session get failure(%s): %w", uri, response.err)
		logger.Error(newErr, "failure: get resource by id", "req", req)
		return &MemoryResourceBlockByIdResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	// Update CompositionState using the Reserved and CompositionState values from Redfish
	compositionStatus, err := response.valueFromJSON("CompositionStatus")
	if err == nil {
		compositionState := compositionStatus.(map[string]interface{})["CompositionState"].(string)
		reserved := compositionStatus.(map[string]interface{})["Reserved"].(bool)

		resourceState := findResourceState(&compositionState, reserved)

		memoryResourceBlock.CompositionStatus.CompositionState = *resourceState
	}

	memoryElements, _ := response.arrayFromJSON("Memory")

	var totalMebibytes float64
	firstDimm := true

	for _, memoryElement := range memoryElements {
		channel := session.query(HTTPOperation.GET, memoryElement.(map[string]interface{})["@odata.id"].(string))

		if channel.err != nil {
			continue
		}

		// Technically all of these values could vary by DIMM
		// For simplicity, we just use the first DIMM
		if firstDimm {
			dataWidthBits, _ := channel.floatFromJSON("DataWidthBits")
			memoryResourceBlock.DataWidthBits = int32(dataWidthBits)

			memoryDeviceType, _ := channel.stringFromJSON("MemoryDeviceType")
			memoryResourceBlock.MemoryDeviceType = memoryDeviceType

			memoryType, _ := channel.stringFromJSON("MemoryType")
			memoryResourceBlock.MemoryType = memoryType

			operatingSpeedMhz, _ := channel.floatFromJSON("OperatingSpeedMhz")
			memoryResourceBlock.OperatingSpeedMhz = int32(operatingSpeedMhz)

			rankCount, _ := channel.floatFromJSON("RankCount")
			memoryResourceBlock.RankCount = int32(rankCount)

			firstDimm = false
		}

		mebibytes, _ := channel.floatFromJSON("CapacityMiB")
		totalMebibytes += mebibytes

		blockDimmStr := getIdFromOdataId(memoryElement.(map[string]interface{})["@odata.id"].(string))

		channelId, _ := strconv.ParseInt(strings.TrimPrefix(blockDimmStr[strings.Index(blockDimmStr, "dimm"):], "dimm"), 10, 64)
		memoryResourceBlock.ChannelId = int32(channelId)
		channelResourceIdx, _ := strconv.ParseInt(strings.TrimPrefix(blockDimmStr[:strings.Index(blockDimmStr, "dimm")], "block"), 10, 64)
		memoryResourceBlock.ChannelResourceIdx = int32(channelResourceIdx)
	}

	memoryResourceBlock.CapacityMiB = int32(totalMebibytes)

	return &MemoryResourceBlockByIdResponse{MemoryResourceBlock: memoryResourceBlock, Status: "Success", ServiceError: nil}, nil
}

func findResourceState(compositionState *string, reserved bool) *ResourceState {
	resourceState := ResourceUnused

	if reserved {
		resourceState = ResourceReserved
	}

	if *compositionState == COMPOSITION_STATE_COMPOSED || *compositionState == COMPOSITION_STATE_COMPOSED_AND_AVAILABLE {
		resourceState = ResourceComposed
	} else if *compositionState == COMPOSITION_STATE_FAILED {
		resourceState = ResourceFailed
	} else if *compositionState == COMPOSITION_STATE_UNAVAILABLE {
		resourceState = ResourceUnavailable
	}

	return &resourceState
}

// GetPorts: Request Ports ids from the backend
func (service *httpfishService) GetPorts(ctx context.Context, settings *ConfigurationSettings, req *GetPortsRequest) (*GetPortsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetPorts ======")
	logger.V(4).Info("GetPorts", "req", req)

	var ret = GetPortsResponse{
		Status: "Failure",
	}

	session := service.service.session.(*Session)

	// Allow blade sessions only
	_, keyExist := session.redfishPaths[FabricPortsKey]
	if !keyExist {
		ret.Status = "Not Supported"
		ret.ServiceError = fmt.Errorf("session (%s) does not support .../fabrics/.../switches/.../ports", session.SessionId)
		logger.Error(ret.ServiceError, "failure: get ports")
		return &ret, ret.ServiceError
	}

	response := session.query(HTTPOperation.GET, session.redfishPaths[FabricPortsKey])

	if response.err != nil {
		ret.ServiceError = response.err
		logger.Error(ret.ServiceError, "failure: get ports")
		return &ret, ret.ServiceError
	}

	ports, _ := response.arrayFromJSON("Members")

	for _, port := range ports {
		uri := port.(map[string]interface{})["@odata.id"].(string)
		tokens := strings.Split(uri, "/")
		if len(tokens) == 0 {
			continue
		}

		portId := tokens[len(tokens)-1]
		if len(portId) > 0 {
			ret.PortIds = append(ret.PortIds, portId)
		}
	}

	ret.Status = "Success"

	return &ret, nil
}

// GetHostPortPcieDevices: Request pcie devices, each representing a physical host port, from the backend
func (service *httpfishService) GetHostPortPcieDevices(ctx context.Context, settings *ConfigurationSettings, req *GetPortsRequest) (*GetPortsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetHostPortPcieDevices ======")
	logger.V(4).Info("GetHostPortPcieDevices", "req", req)

	var ret = GetPortsResponse{
		Status: "Failure",
	}

	session := service.service.session.(*Session)

	// Allow host sessions only
	_, keyExist := session.redfishPaths[ChassisPcieDevKey]
	if !keyExist {
		ret.Status = "Not Supported"
		ret.ServiceError = fmt.Errorf("session (%s) does not support .../chassis/.../pciedevices", session.SessionId)
		logger.Error(ret.ServiceError, "failure: get port pcie devices")
		return &ret, ret.ServiceError
	}

	response := session.query(HTTPOperation.GET, session.redfishPaths[ChassisPcieDevKey])
	if response.err != nil {
		ret.ServiceError = response.err
		logger.Error(ret.ServiceError, "failure: get ports")
		return &ret, ret.ServiceError
	}

	ports, _ := response.arrayFromJSON("Members")
	for _, port := range ports {
		uri := port.(map[string]interface{})["@odata.id"].(string)
		tokens := strings.Split(uri, "/")
		if len(tokens) == 0 {
			continue
		}

		pcieId := tokens[len(tokens)-1]
		if len(pcieId) > 0 {
			ret.PortIds = append(ret.PortIds, pcieId)
		}
	}

	ret.Status = "Success"

	return &ret, nil
}

// GetPortDetails: Request Ports info from the backend
func (service *httpfishService) GetPortDetails(ctx context.Context, settings *ConfigurationSettings, req *GetPortDetailsRequest) (*GetPortDetailsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetPortDetails ======")
	logger.V(4).Info("GetPortDetails", "req", req)

	var ret = GetPortDetailsResponse{
		PortInformation: PortInformation{},
		Status:          "Failure",
	}

	session := service.service.session.(*Session)

	// Allow blade sessions only
	_, keyExist := session.redfishPaths[FabricPortsKey]
	if !keyExist {
		ret.Status = "Not Supported"
		ret.ServiceError = fmt.Errorf("session (%s) does not support .../fabrics/.../switches/.../ports", session.SessionId)
		logger.Error(ret.ServiceError, "failure: get port details", "req", req)
		return &ret, ret.ServiceError
	}

	response := session.query(HTTPOperation.GET, session.buildPath(FabricPortsKey, req.PortId))
	if response.err != nil {
		ret.ServiceError = response.err
		logger.Error(ret.ServiceError, "failure: get port details", "req", req)
		return &ret, ret.ServiceError
	}

	id, _ := response.stringFromJSON("Id")
	ret.PortInformation.Id = id
	ret.PortInformation.PortProtocol, _ = response.stringFromJSON("PortProtocol")
	ret.PortInformation.PortMedium, _ = response.stringFromJSON("PortMedium")
	width, err := response.floatFromJSON("ActiveWidth")
	if err == nil {
		ret.PortInformation.Width = int32(width)
	}
	ret.PortInformation.LinkStatus, _ = response.stringFromJSON("LinkStatus")
	ret.PortInformation.LinkState, _ = response.stringFromJSON("LinkState")

	status, _ := response.valueFromJSON("Status")

	health := status.(map[string]interface{})["Health"].(string)
	state := status.(map[string]interface{})["State"].(string)
	healthAndState := fmt.Sprintf("%s/%s", health, state)

	ret.PortInformation.StatusHealth = health
	ret.PortInformation.StatusState = state
	ret.Status = healthAndState

	portField, err := response.valueFromJSON("Port")
	if err == nil {
		speedFloat, _ := portField.(map[string]interface{})["CurrentSpeedGbps"].(float64)
		ret.PortInformation.CurrentSpeedGbps = int32(speedFloat)
	}

	// Extract GCXLID from endpoint
	uriOfTargetEndpoint, errOfTargetEndpoint := session.getEndpointUriFromPort(id)
	if errOfTargetEndpoint != nil {
		ret.ServiceError = errOfTargetEndpoint
		logger.Error(ret.ServiceError, "failure: get port details", "req", req)
		return &ret, ret.ServiceError
	}

	response = session.query(HTTPOperation.GET, *uriOfTargetEndpoint)
	if response.err != nil {
		ret.ServiceError = response.err
		logger.Error(ret.ServiceError, "failure: get port details", "req", req)
		return &ret, ret.ServiceError
	}

	identifiers, _ := response.valueFromJSON("Identifiers")
	ret.PortInformation.GCxlId = identifiers.([]interface{})[0].(map[string]interface{})["DurableName"].(string)

	//Note: "PortInformation.LinkedPortUri" can't be determined here.  Handled separately.

	return &ret, ret.ServiceError
}

// GetHostPortSnById: Request the serial number from a specific port (ie - pcie device) and cxl host
func (service *httpfishService) GetHostPortSnById(ctx context.Context, settings *ConfigurationSettings, req *GetHostPortSnByIdRequest) (*GetHostPortSnByIdResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetHostPortSnById ======")
	logger.V(4).Info("GetHostPortSnById", "req", req)

	var ret = GetHostPortSnByIdResponse{
		Status: "Failure",
	}

	session := service.service.session.(*Session)

	// Allow host sessions only
	_, keyExist := session.redfishPaths[ChassisPcieDevKey]
	if !keyExist {
		ret.Status = "Not Supported"
		ret.ServiceError = fmt.Errorf("session (%s) does not support .../chassis/.../pciedevices", session.SessionId)
		logger.Error(ret.ServiceError, "failure: get host port sn by id", "req", req)
		return &ret, ret.ServiceError
	}

	// Query port
	deviceUri := session.buildPath(ChassisPcieDevKey, req.PortId)
	response := session.query(HTTPOperation.GET, deviceUri)
	if response.err != nil {
		ret.ServiceError = fmt.Errorf("session (%s) query failure (%s) for port (%s): %w", session.SessionId, deviceUri, req.PortId, response.err)
		logger.Error(ret.ServiceError, "failure: get host port sn by id")
		return &ret, ret.ServiceError
	}

	// Extract the SN
	sn, err := response.stringFromJSON("SerialNumber")
	if err != nil {
		ret.ServiceError = fmt.Errorf("session (%s) query failure (%s) for port (%s): SerialNumber NOT found: %w", session.SessionId, deviceUri, req.PortId, response.err)
		logger.Error(ret.ServiceError, "failure: get host port sn by id")
		return &ret, ret.ServiceError
	}

	ret.SerialNumber = sn
	ret.Status = "Success"

	return &ret, nil
}

// GetMemoryDevices: Delete memory region info by memory id
func (service *httpfishService) GetMemoryDevices(ctx context.Context, settings *ConfigurationSettings, req *GetMemoryDevicesRequest) (*GetMemoryDevicesResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryDevices ======")
	logger.V(4).Info("get memory devices", "request", req)

	session := service.service.session.(*Session)
	response := session.query(HTTPOperation.GET, session.redfishPaths[ChassisPcieDevKey])
	if response.err != nil {
		newErr := fmt.Errorf("backend get failure [%s]: %w", session.redfishPaths[ChassisPcieDevKey], response.err)
		logger.Error(newErr, "failure: get memory devices")
		return &GetMemoryDevicesResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	// Mapping of physical device IDs (keys) to a slice of logical device IDs (values)
	deviceIdMap := make(map[string][]string)

	physicalDeviceUris, _ := response.memberOdataArray()
	for _, uri := range physicalDeviceUris {
		phyDevId := getIdFromOdataId(uri)
		deviceIdMap[phyDevId] = []string{}
	}

	for phyDevId := range deviceIdMap {
		response := session.query(HTTPOperation.GET, session.buildPath("ChassisPcieDev", phyDevId))
		if response.err != nil {
			newErr := fmt.Errorf("backend get failure [%s]: %w", session.buildPath("ChassisPcieDev", phyDevId), response.err)
			logger.Error(newErr, "failure: get memory devices")
			return &GetMemoryDevicesResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		logicalDevicesUri, err := response.odataStringFromJSON("CXLLogicalDevices")
		if err != nil {
			newErr := fmt.Errorf("backend response key ['CXLLogicalDevices'] not found for uri [%s]: %w", logicalDevicesUri, err)
			logger.Error(newErr, "failure: get memory devices")
			return &GetMemoryDevicesResponse{Status: "Failure", ServiceError: newErr}, newErr
		}

		cxlCollection := session.query(HTTPOperation.GET, logicalDevicesUri)
		logicalDeviceUris, _ := cxlCollection.memberOdataArray()
		for _, uri := range logicalDeviceUris {
			logicalDevId := getIdFromOdataId(uri)
			deviceIdMap[phyDevId] = append(deviceIdMap[phyDevId], logicalDevId)
		}
	}

	return &GetMemoryDevicesResponse{DeviceIdMap: deviceIdMap, Status: "Success", ServiceError: nil}, nil
}

// pcieGenToSpeed: convert PCIe device generation string to speed
func pcieGenToSpeed(gen string) int32 {
	switch gen {
	// CXL works on a minimum of gen3 speed
	case "gen3":
		return 8
	case "gen4":
		return 16
	case "gen5":
		return 32
	default:
		return 0
	}
}

// GetMemoryDeviceDetails: Get a specific memory region info by memory id
func (service *httpfishService) GetMemoryDeviceDetails(ctx context.Context, setting *ConfigurationSettings, req *GetMemoryDeviceDetailsRequest) (*GetMemoryDeviceDetailsResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryDeviceDetails ======")
	logger.V(4).Info("get memory dev by id", "request", req)

	session := service.service.session.(*Session)

	pcieDeviceUri := session.buildPath("ChassisPcieDev", req.PhysicalDeviceId)
	response := session.query(HTTPOperation.GET, pcieDeviceUri)
	if response.err != nil {
		newErr := fmt.Errorf("backend get failure [%s]: %w", pcieDeviceUri, response.err)
		logger.Error(newErr, "failure: get memory device details")
		return &GetMemoryDeviceDetailsResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	status, _ := response.valueFromJSON("Status")
	memDev := GetMemoryDeviceDetailsResponse{
		Status:       status.(map[string]interface{})["State"].(string),
		ServiceError: nil,
	}

	sn, err := response.stringFromJSON("SerialNumber")
	if err != nil {
		newErr := fmt.Errorf("backend response key ['SerialNumber'] not found for uri [%s]: %w", pcieDeviceUri, err)
		logger.Error(newErr, "failure: get memory device details")
		return &GetMemoryDeviceDetailsResponse{Status: "Failure", ServiceError: newErr}, newErr
	}
	memDev.SerialNumber = sn

	// PCIe dev to get link status and device type [optional]
	//CXLDevice DeviceType
	cxlDev, err := response.valueFromJSON("CXLDevice")
	if err == nil {
		if cxlDev.(map[string]interface{})["DeviceType"] != nil {
			memDev.DeviceType = cxlDev.(map[string]interface{})["DeviceType"].(string)
		}
	}
	// memDev.LinkStatus=    MemoryDeviceLinkStatus{},  // PCIeInterface
	linkStatus, err := response.valueFromJSON("PCIeInterface")
	if err == nil {
		if linkStatus.(map[string]interface{})["LanesInUse"] != nil {
			memDev.LinkStatus.CurrentWidth = linkStatus.(map[string]interface{})["LanesInUse"].(int32)
		}
		if linkStatus.(map[string]interface{})["MaxLanes"] != nil {
			memDev.LinkStatus.MaxWidth = linkStatus.(map[string]interface{})["MaxLanes"].(int32)
		}
		if linkStatus.(map[string]interface{})["MaxPCIeType"] != nil {
			memDev.LinkStatus.MaxSpeedGTps = pcieGenToSpeed(linkStatus.(map[string]interface{})["MaxPCIeType"].(string))
		}
		if linkStatus.(map[string]interface{})["PCIeType"] != nil {
			memDev.LinkStatus.CurrentSpeedGTps = pcieGenToSpeed(linkStatus.(map[string]interface{})["PCIeType"].(string))
		}
	}

	// CXL logical dev to get memory size
	logicalDevicesUri, err := response.odataStringFromJSON("CXLLogicalDevices")
	if err != nil {
		newErr := fmt.Errorf("backend response key ['CXLLogicalDevices'] not found for uri [%s]: %w", logicalDevicesUri, err)
		logger.Error(newErr, "failure: get memory device details")
		return &GetMemoryDeviceDetailsResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	logicalDeviceUri := fmt.Sprintf("%s/%s", logicalDevicesUri, req.LogicalDeviceId)
	logicalDevice := session.query(HTTPOperation.GET, logicalDeviceUri)
	if logicalDevice.err != nil {
		newErr := fmt.Errorf("backend get failure [%s]: %w", logicalDeviceUri, logicalDevice.err)
		logger.Error(newErr, "failure: get memory device details")
		return &GetMemoryDeviceDetailsResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	memSize, _ := logicalDevice.valueFromJSON("MemorySizeMiB")
	if memSize != nil {
		memDev.MemorySizeMiB = int32(memSize.(float64))
	}

	logger.V(2).Info("success: GetMemoryDeviceDetails", "memDev", memDev)

	return &memDev, nil
}

// Extract the connection info for a specific endpoint
func (session *Session) getConnectionFromEndpoint(uriEndpoint string) (*string, error) {
	// Get the endpoint
	responseOfEndpoint := session.query(HTTPOperation.GET, uriEndpoint)
	if responseOfEndpoint.err != nil {
		return nil, fmt.Errorf("failed to get endpoint")
	}

	// Extra the corresponding endpoint
	links, _ := responseOfEndpoint.valueFromJSON("Links")
	uriOfConnection := links.(map[string]interface{})["Connections"].([]interface{})[0].(map[string]interface{})["@odata.id"].(string)
	// if the connection is empty, no connection
	if &uriOfConnection == nil {
		return nil, fmt.Errorf("endpoint has no connection")
	}

	// Get the connection id
	connectionId := getIdFromOdataId(uriOfConnection)

	return &connectionId, nil
}

// FreeMemoryById: Delete memory region (memory chunk) by memory id
func (service *httpfishService) FreeMemoryById(ctx context.Context, settings *ConfigurationSettings, req *FreeMemoryRequest) (*FreeMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== FreeMemoryById ======")
	logger.V(4).Info("free memory", "request", req)

	session := service.service.session.(*Session)

	// Deallocate the memory region
	// Currently, a successful delete returns an empty response
	response := session.query(HTTPOperation.DELETE, session.buildPath(SystemMemoryChunksKey, req.MemoryId))
	if response.err != nil {
		newErr := fmt.Errorf("backend session delete failure(%s): %w", session.buildPath(SystemMemoryChunksKey, req.MemoryId), response.err)
		logger.Error(newErr, "failure: free memory by id", "req", req)
		return &FreeMemoryResponse{Status: "Failure", ServiceError: newErr}, newErr
	}

	delete(session.memoryChunkPath, req.MemoryId)

	freeMemoryResponse := &FreeMemoryResponse{
		Status:       "Success",
		ServiceError: nil}

	return freeMemoryResponse, nil
}

// GetMemoryById: Get a specific memory region info by memory id
func (service *httpfishService) GetMemoryById(ctx context.Context, setting *ConfigurationSettings, req *GetMemoryByIdRequest) (*GetMemoryByIdResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemoryById ======")
	logger.V(4).Info("get memory by id", "request", req)

	memoryByIdResponse := GetMemoryByIdResponse{
		MemoryRegion: TypeMemoryRegion{
			MemoryId: req.MemoryId,
			Status:   "Success",
			Type:     MemoryType(MEMORYTYPE_MEMORY_TYPE_REGION),
			SizeMiB:  0,
		},
		Status: "Success",
	}

	session := service.service.session.(*Session)

	path, exist := session.memoryChunkPath[req.MemoryId]
	if !exist {
		// rescan memory collection
		memReq := GetMemoryRequest{}
		service.GetMemory(ctx, setting, &memReq)

		path, exist = session.memoryChunkPath[req.MemoryId]
		if !exist {
			memoryByIdResponse.Status = "Not Found"
			memoryByIdResponse.MemoryRegion.Status = "Failure"
			memoryByIdResponse.ServiceError = fmt.Errorf("memory (%s) does not exist", req.MemoryId)

			return &memoryByIdResponse, memoryByIdResponse.ServiceError
		}
	}
	response := session.query(HTTPOperation.GET, path)

	if response.err != nil {
		memoryByIdResponse.Status = "Failure"
		memoryByIdResponse.MemoryRegion.Status = "Failure"
		memoryByIdResponse.ServiceError = response.err

		return &memoryByIdResponse, memoryByIdResponse.ServiceError
	}
	memoryByIdResponse.MemoryRegion.MemoryId, _ = response.stringFromJSON("Id")
	val, _ := response.valueFromJSON("MemoryChunkSizeMiB")
	memoryByIdResponse.MemoryRegion.SizeMiB = int32(val.(float64))

	if strings.Contains(path, "CXL") { // host cxl memory
		memoryByIdResponse.MemoryRegion.Type = MemoryType(MEMORYTYPE_MEMORY_TYPE_CXL)
		// Check if performacne metric is reported
		oemField, _ := response.valueFromJSON("Oem")
		if oemField != nil {
			// Example response from redfish
			/*     "Oem": {
				  	"Seagate": {
			 		"Bandwidth": "8.34 GiB/s",
			 		"Latency": "514 ns"
						}
					},
			*/
			bwStr := oemField.(map[string]interface{})["Seagate"].(map[string]interface{})["Bandwidth"].(string)
			bwFloat, _ := strconv.ParseFloat(strings.Split(bwStr, " ")[0], 64)
			memoryByIdResponse.MemoryRegion.Bandwidth = int32(bwFloat)
			latStr := oemField.(map[string]interface{})["Seagate"].(map[string]interface{})["Latency"].(string)
			latInt64, _ := strconv.ParseInt(strings.Split(latStr, " ")[0], 10, 64)
			memoryByIdResponse.MemoryRegion.Latency = int32(latInt64)
		}

		links, _ := response.valueFromJSON("Links")
		endpoints, ok := links.(map[string]interface{})["Endpoints"].([]interface{})
		if !ok || len(endpoints) >= 2 {
			return &memoryByIdResponse, fmt.Errorf("invalid endpoints")
		}

		// This entire IF is about finding the host port associated with the requested memoryId
		if len(endpoints) != 0 {
			uriSystemMemory := endpoints[0].(map[string]interface{})["@odata.id"].(string)

			response := session.query(HTTPOperation.GET, uriSystemMemory)
			if response.err != nil {
				return &memoryByIdResponse, fmt.Errorf("get [%s] failure: %w", uriSystemMemory, response.err)
			}

			links, _ = response.valueFromJSON("Links")
			sources := links.(map[string]interface{})["MemoryMediaSources"].([]interface{})
			if !ok || len(sources) >= 2 {
				return &memoryByIdResponse, fmt.Errorf("invalid memory media sources")
			}

			if len(sources) != 0 {
				uriChassisMemoryChunks := sources[0].(map[string]interface{})["@odata.id"].(string)

				response = session.query(HTTPOperation.GET, uriChassisMemoryChunks)
				if response.err != nil {
					return &memoryByIdResponse, fmt.Errorf("get [%s] failure: %w", uriChassisMemoryChunks, response.err)
				}

				links, _ = response.valueFromJSON("Links")
				devices := links.(map[string]interface{})["CXLLogicalDevices"].([]interface{})
				if !ok || len(devices) >= 2 {
					return &memoryByIdResponse, fmt.Errorf("invalid cxl logical devices")
				}

				if len(devices) != 0 {
					uriCxlLogicalDevice := devices[0].(map[string]interface{})["@odata.id"].(string)

					elements := strings.Split(uriCxlLogicalDevice, "/")
					if len(elements) < 8 {
						return &memoryByIdResponse, fmt.Errorf("invalid cxl logical devices uri [%s]", uriCxlLogicalDevice)
					}

					memoryByIdResponse.MemoryRegion.PortId = elements[len(elements)-3]
					memoryByIdResponse.MemoryRegion.LogicalDeviceId = elements[len(elements)-1]
				}
			}
		}

		return &memoryByIdResponse, memoryByIdResponse.ServiceError

	} else if strings.Contains(path, "DIMMs") { // host local memory
		memoryByIdResponse.MemoryRegion.Type = MemoryType(MEMORYTYPE_MEMORY_TYPE_LOCAL)
		return &memoryByIdResponse, memoryByIdResponse.ServiceError

	} else { // memory appliance memory
		links, _ := response.valueFromJSON("Links")
		endpoints, ok := links.(map[string]interface{})["Endpoints"].([]interface{})
		if !ok || len(endpoints) >= 2 {
			return nil, fmt.Errorf("invalid endpoints")
		}

		// This entire IF is about finding the blade port associated with the requested memoryId
		if len(endpoints) != 0 {
			uriEndpoint := endpoints[0].(map[string]interface{})["@odata.id"].(string)

			response = session.query(HTTPOperation.GET, uriEndpoint)
			if response.err != nil {
				return nil, fmt.Errorf("get [%s] failure: %w", uriEndpoint, response.err)
			}

			links, _ = response.valueFromJSON("Links")
			ports, ok := links.(map[string]interface{})["ConnectedPorts"].([]interface{})
			if !ok || len(ports) >= 2 {
				return nil, fmt.Errorf("invalid connected ports")
			}

			if len(ports) != 0 {
				uriPort := ports[0].(map[string]interface{})["@odata.id"].(string)

				elements := strings.Split(uriPort, "/")
				if len(elements) < 8 {
					return nil, fmt.Errorf("invalid port uri [%s]", uriPort)
				}

				memoryByIdResponse.MemoryRegion.PortId = elements[len(elements)-1]
			}
		}

		return &memoryByIdResponse, memoryByIdResponse.ServiceError
	}
}

// GetMemory: Get the list of memory ids for a particular endpoint
func (service *httpfishService) GetMemory(ctx context.Context, settings *ConfigurationSettings, req *GetMemoryRequest) (*GetMemoryResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info("====== GetMemory ======")
	logger.V(4).Info("get memory", "request", req)

	memoryResponse := GetMemoryResponse{
		MemoryIds: make([]string, 0),
		Status:    "Success",
	}

	session := service.service.session.(*Session)

	response := session.query(HTTPOperation.GET, session.redfishPaths[SystemMemoryChunksKey])

	if response.err != nil {
		memoryResponse.Status = "Failure"
		memoryResponse.ServiceError = response.err

		return &memoryResponse, response.err
	}

	members, _ := response.arrayFromJSON("Members")

	for _, member := range members {
		uri := member.(map[string]interface{})["@odata.id"].(string)

		components := strings.Split(uri, "/")

		if len(components) > 0 {
			memoryResponse.MemoryIds = append(memoryResponse.MemoryIds, components[len(components)-1])
			session.memoryChunkPath[components[len(components)-1]] = uri
		}
	}

	//workaround for cxl-host multiple domain name
	cxlMemoryPath, exist := session.redfishPaths[SystemMemoryChunksCXLKey]
	if exist {
		response = session.query(HTTPOperation.GET, cxlMemoryPath)

		if response.err != nil {
			memoryResponse.Status = "Failure"
			memoryResponse.ServiceError = response.err

			return &memoryResponse, response.err
		}

		members, _ = response.arrayFromJSON("Members")

		for _, member := range members {
			uri := member.(map[string]interface{})["@odata.id"].(string)

			components := strings.Split(uri, "/")

			if len(components) > 0 {
				memoryResponse.MemoryIds = append(memoryResponse.MemoryIds, components[len(components)-1])
				session.memoryChunkPath[components[len(components)-1]] = uri
			}
		}
	}

	return &memoryResponse, nil
}

// GetBackendInfo: Get the information of this backend
func (service *httpfishService) GetBackendInfo(ctx context.Context) *GetBackendInfoResponse {
	return &GetBackendInfoResponse{BackendName: "httpfish", Version: "0.1", SessionId: service.service.session.(*Session).SessionId}
}
