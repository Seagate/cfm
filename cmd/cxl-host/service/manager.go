// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cxl_host

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Seagate/cxl-lib/pkg/cxl"

	"github.com/google/uuid"
	"github.com/zcalusic/sysinfo"
)

const (
	ServiceRootVersion     = "#ServiceRoot.v1_14_0.ServiceRoot"
	AccountServiceVersion  = "#AccountService.v1_11_1.AccountService"
	ChassisVersion         = "#Chassis.v1_21_0.Chassis"
	FabricVersion          = "#Fabric.v1_3_0.Fabric"
	ManagerAccountVersion  = "#ManagerAccount.v1_9_0.ManagerAccount"
	MemoryName             = "DIMM1"
	MemoryVersion          = "#Memory.v1_16_0.Memory"
	PortVersion            = "#Port.v1_7_0.Port"
	PowerSubsystemVersion  = "#PowerSubsystem.v1_1_0.PowerSubsystem"
	PowerSupplyVersion     = "#PowerSupply.v1_5_0.PowerSupply"
	RoleVersion            = "#Role.v1_3_1.Role"
	SessionTimeout         = 60 * 30
	SessionServiceVersion  = "#SessionService.v1_1_8.SessionService"
	SessionVersion         = "#Session.v1_5_0.Session"
	SwitchName             = "Switch1"
	SwitchVersion          = "#Switch.v1_8_0.Switch"
	SystemVersion          = "#ComputerSystem.v1_20_0.ComputerSystem"
	ZonesMax               = 1
	MemoryDomainVersion    = "#MemoryDomain.v1_5_0.MemoryDomain"
	MemoryChunkVersion     = "#MemoryChunks.V1_5_0.MemoryChunks"
	PCIeDevice             = "#PCIeDevice.v1_11_0.PCIeDevice"
	PCIeFunction           = "#PCIeFunction.v1_5_0.PCIeFunction"
	MemoryVersionForSystem = "#Memory.v1_17_0.Memory"
	CXLLogicalDevice       = "#CXLLogicalDevice.v1_0_0.CXLLogicalDevice"
)

const (
	systemPrefix = "cxl-host"
)

type serviceIdentification struct {
	initialized bool
	uuid        uuid.UUID
	set         time.Time
}

var serviceId serviceIdentification

var chassisName string = ""
var chassisUUID string = ""
var fabricUUIDs map[string]string
var switchUUID string = ""
var portUUIDs map[string]string
var systemName string = ""

func init() {
	fabricUUIDs = map[string]string{
		"CXL": uuid.New().String(),
	}
	portUUIDs = map[string]string{
		"P0": uuid.New().String(),
		"P1": uuid.New().String(),
	}
}

////  ********** RedFish helper functions **********  ////

// CheckMemoryName: return true if the memory name is valid
func CheckMemoryName(id string) bool {
	text, num, err := IdParse(id)
	if err == nil {
		if text == "CXL" {
			if num != 0 && num <= GetCXLDevCnt() {
				return true
			}
		}
		if id == GetMemoryName() {
			return true
		}
	}
	return false
}

// CheckMemoryDomainName: DIMMs domain always exist. CXL domain only exist when CXL device presents
func CheckMemoryDomainName(id string) bool {
	switch id {
	case "DIMMs":
		return true
	case "CXL":
		if len(cxlDevMap) != 0 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

// CheckMemoryChunkName: return true if memChunk exists in the memDomain
func CheckMemoryChunkName(memDomain string, memChunk string) bool {
	switch strings.ToLower(memDomain) {
	case "dimms":
		return CheckDimmChunk(memChunk)
	case "cxl":
		return CheckCxlChunk(memChunk)
	default:
		// return false if memDomain doesn't exist
		return false
	}
}

// CheckMemoryChunkNumaName: return true if node exists in the memDomain
func CheckMemoryChunkNumaName(memDomain string, memChunk string) bool {
	switch strings.ToLower(memDomain) {
	case "dimms":
		return CheckDimmChunk(memChunk)
	case "cxl":
		return CheckCxlNumaChunk(memChunk)
	default:
		// return false if memDomain doesn't exist
		return false
	}
}

// CheckDimmChunk: return ture if memChunk exists in the DIMMs domain
func CheckDimmChunk(memChunk string) bool {
	localNodes := GetDimmAttachedNumaNodes()
	for _, node := range localNodes {
		if node == strings.ToLower(memChunk) {
			return true
		}
	}
	// return false if no matching node
	return false
}

// CheckCxlChunk: return ture if memChunk exists in the CXL domain
func CheckCxlChunk(memChunk string) bool {
	cxlDevList := GetCXLDevList()
	for _, dev := range cxlDevList {
		if dev == BDtoBDF(memChunk) {
			return true
		}
	}
	// return false if no matching cxl dev
	return false
}

// CheckCxlNumaChunk: return ture if memChunk exists in the CXL domain
func CheckCxlNumaChunk(memChunk string) bool {
	cxlNodeMap := GetCXLNumaNodes()
	for node := range cxlNodeMap {
		if node == memChunk {
			return true
		}
	}
	// return false if no matching cxl dev
	return false
}

func GetPCIeFunctionIds(pCIeId string) []string {
	// CXL 1.1 spec defines only one PCIe function with function id 0
	// Use hardcoded value for now
	return []string{"0"}
}

// Get CXL Logical Device Id
func GetCXLLogicalDeviceIds(pCIeId string) []string {
	// CXL 1.1 can only be single logical device
	// Use hardcoded value for now
	return []string{"0"}
}

// GetServiceUUID: return the unique id for this instance of the service application
func GetServiceUUID() string {
	if !serviceId.initialized {
		serviceId.initialized = true
		serviceId.uuid = uuid.New()
		serviceId.set = time.Now()
	}
	return serviceId.uuid.String()
}

// GetChassisName: return the unique name for this chassis
func GetChassisName() string {
	if chassisName == "" {
		chassisName = systemPrefix + "-" + runtime.GOOS
	}
	return chassisName
}

// GetChassisUUID: return the unique id for this chassis
func GetChassisUUID() string {
	if chassisUUID == "" {
		chassisUUID = uuid.New().String()
	}
	return chassisUUID
}

// GetChassisTag: return the unique asset tag for this chassis
func GetChassisTag() string {
	return "Seagate" + "-" + GetChassisUUID()
}

// GetChassisVersion: return the version for this chassis
func GetChassisVersion() string {
	return "1.0.0"
}

// GetMemoryName: return the unique name for this memory part
func GetMemoryName() string {
	return MemoryName
}

// GetFabrics: return a map of supported fabrics with UUIDs
func GetFabrics() map[string]string {
	return fabricUUIDs
}

// IsFabricSupported: return true if Fabric is supported
func IsFabricSupported(id string) bool {
	_, ok := fabricUUIDs[id]
	return ok
}

// GetFabricUUID: return a UUID for a given Fabric
func GetFabricUUID(id string) string {
	uuid := ""
	value, ok := fabricUUIDs[id]
	if ok {
		uuid = value
	}
	return uuid
}

// GetSwitchName: return the unique name of the switch
func GetSwitchName() string {
	return SwitchName
}

// GetSwitchUUID: return the unique id for this switch
func GetSwitchUUID() string {
	if switchUUID == "" {
		switchUUID = uuid.New().String()
	}
	return switchUUID
}

// GetSystemName: return the unique name for this system
func GetSystemName() string {
	if systemName == "" {
		systemName = systemPrefix + "-" + runtime.GOOS
	}
	return systemName
}

// GetPorts: return a map of supported ports with UUIDs
func GetPorts() map[string]string {
	return portUUIDs
}

// IsPortSupported: return true if Port is supported
func IsPortSupported(id string) bool {
	_, ok := portUUIDs[id]
	return ok
}

// GetPortUUID: return a UUID for a given Port
func GetPortUUID(id string) string {
	uuid := ""
	value, ok := portUUIDs[id]
	if ok {
		uuid = value
	}
	return uuid
}

////  ********** CXL helper functions **********  ////

// Map of CXL devices as a string of BUS:DEVICE.FUNCTION ( BDF )
var cxlDevMap = map[string]*cxl.CxlDev{}

// List of CXL BDF to maintain order
var cxlDevOrder []string

func init() {
	cxlDevMap = cxl.InitCxlDevList()
	for bdf := range cxlDevMap {
		cxlDevOrder = append(cxlDevOrder, bdf)
	}
	sort.Strings(cxlDevOrder)
}

// BDFtoBD: return cxl devices as bus-dev ( ":" is not a legal character in URL )
func BDFtoBD(s string) string {
	return strings.ReplaceAll(strings.Split(s, ".")[0], ":", "-")
}

// BDtoBDF: return bus:dev.fun from bus-dev
func BDtoBDF(s string) string {
	return strings.ReplaceAll(s, "-", ":") + ".0"
}

// CxlDevIndexToBDF: return cxl device BDF by index ( 1 based )
func CxlDevIndexToBDF(i int) string {
	return cxlDevOrder[i-1]
}

// CxlDevBDFtoIndex: return cxl device index by BDF index ( 1 based )
func CxlDevBDFtoIndex(ibdf string) int {
	for i, bdf := range cxlDevOrder {
		if bdf == ibdf {
			return i + 1
		}
	}
	return 0
}

// CxlDevNodeToBDF: return cxl device BDF by numa node
func CxlDevNodeToBDF(s string) string {
	for node, bdf := range GetCXLNumaNodes() {
		if node == s {
			return bdf
		}
	}
	return "Error"
}

// GetCXLDevCnt: return the number of CXL devices on the host
func GetCXLDevCnt() int {
	return len(cxlDevOrder)
}

// GetCXLDevList: return a list of cxl devices as bus:dev.fun
func GetCXLDevList() []string {
	return cxlDevOrder
}

// GetCXLWithMemDevList: return a list of cxl devices as bus:dev.fun
func GetCXLWithMemDevList() []string {
	bdfList := make([]string, 0, len(cxlDevOrder))
	for _, bdf := range cxlDevOrder {
		dev := GetCXLDevInfo(bdf)
		if dev.GetCxlCap().Mem_En {
			bdfList = append(bdfList, bdf)
		}
	}
	return bdfList
}

// GetCXLDevInfo: return cxl device structure
func GetCXLDevInfo(bdf string) *cxl.CxlDev {
	dev, ok := cxlDevMap[bdf]
	if !ok {
		return nil
	}
	return dev
}

// GetCXLDevInfoByIndex: return cxl device structure by index ( 1 based )
func GetCXLDevInfoByIndex(i int) *cxl.CxlDev {
	dev, ok := cxlDevMap[CxlDevIndexToBDF(i)]
	if !ok {
		return nil
	}
	return dev
}

// GetCXLDevInfoByNode: return cxl device structure by numa node
func GetCXLDevInfoByNode(i int) *cxl.CxlDev {
	dev, ok := cxlDevMap[CxlDevIndexToBDF(i)]
	if !ok {
		return nil
	}
	return dev
}

// GetCXLDevInfoByNode: return cxl device structure by numa node
func FormatGCXLID(dev *cxl.CxlDev) string {
	sn := dev.GetSerialNumber()
	regex := regexp.MustCompile(`(\w{2})(\w{2})(\w{2})(\w{2})(\w{2})(\w{2})(\w{2})(\w{2})`)
	// Index of MLD. Hardcode to 0000 for SLD
	mldSuffix := "0000"
	return regex.ReplaceAllString(sn[2:], `${1}-${2}-${3}-${4}-${5}-${6}-${7}-${8}:`) + mldSuffix
}

type MemoryChunkMiB struct {
	BaseAddress int64
	Size        int64
}

// GetCxlAddressInfoMiB: return the base address and size of the cxl memory
func GetCxlAddressInfoMiB(bdf string) MemoryChunkMiB {
	cxlDev := GetCXLDevInfo(bdf)
	if cxlDev == nil {
		return MemoryChunkMiB{BaseAddress: 0, Size: 0}
	}
	return MemoryChunkMiB{BaseAddress: cxlDev.GetMemoryBaseAddr() >> 20, Size: cxlDev.GetMemorySize() >> 20}
}

// Readsystemfs: helper function to read a systemfs and return a trimed string
func Readsystemfs(f string) (string, error) {
	fileBytes, err := os.ReadFile(f)
	return strings.Trim(string(fileBytes), "\n"), err
}

// GetDimmAttachedNumaNodes: return a list of numa node name strings, where the node contains CPU ( meaning it is not a memory only node )
func GetDimmAttachedNumaNodes() []string {
	localNodes := []string{}
	entries, err := os.ReadDir(PlatformPath.SysDev)
	if err == nil {
		for _, entry := range entries { // iterate all numa nodes
			FileName := entry.Name()
			if strings.ToLower(FileName[:4]) == "node" {
				cpulist, err2 := Readsystemfs(fmt.Sprintf("%s%s/cpulist", PlatformPath.SysDev, FileName)) // check has cpu
				if err2 == nil && cpulist != "" {
					// Now this node has cpu. It is time to check for memory
					mem := GetNumaMemInfo(FileName)
					if mem.MemTotal != 0 {
						localNodes = append(localNodes, FileName)
					}
				}
			}
		}
	}
	return localNodes
}

// GetMemOnlyNumaNodes: return a list of numa node name strings, where the node contains NO CPU ( meaning it is a memory only node )
func GetMemOnlyNumaNodes() []string {
	localNodes := []string{}
	entries, err := os.ReadDir(PlatformPath.SysDev)
	if err == nil {
		for _, entry := range entries { // iterate all numa nodes
			FileName := entry.Name()
			if strings.ToLower(FileName[:4]) == "node" {
				cpulist, err2 := Readsystemfs(fmt.Sprintf("%s%s/cpulist", PlatformPath.SysDev, FileName)) // check has cpu
				if err2 == nil && cpulist == "" {
					// Now this node has NO cpu. It is time to check for memory
					mem := GetNumaMemInfo(FileName)
					if mem.MemTotal != 0 {
						localNodes = append(localNodes, FileName)
					}
				}
			}
		}
	}
	return localNodes
}

// GetCXLNumaNodes: return a map of [node]BDF
func GetCXLNumaNodes() map[string]string {
	nodeMap := make(map[string]string)
	for _, node := range GetMemOnlyNumaNodes() { // Initialize Node Map
		nodeMap[node] = ""
	}
	if blockSizeByte != -1 {
		cxlDevList := GetCXLWithMemDevList()
		for _, bdf := range cxlDevList {
			dev := GetCXLDevInfo(bdf)
			base_addr := dev.GetMemoryBaseAddr()
			if base_addr != 0 {
				blockIndex := base_addr / blockSizeByte
				entries, err := os.ReadDir(fmt.Sprintf("%smemory%d", PlatformPath.SysMem, blockIndex))
				if err == nil {
					for _, entry := range entries {
						if strings.HasPrefix(entry.Name(), "node") {
							nodeMap[entry.Name()] = bdf
						}
					}
				}
			}
		}
	}
	return nodeMap
}

type MemMetrics struct {
	Bandwidth string
	Latency   string
}

var measuredMemMetrics map[string]MemMetrics

// MeasurePerfMetrics: Measure the performance
func MeasurePerfMetrics(cxlDev *cxl.CxlDev, to chan bool) {
	// Report Priority: Measured>CDAT READ>CDAT access
	bw, _ := cxlDev.MeasureBandwidth()
	lat, _ := cxlDev.MeasureLatency()
	if (bw == 0 || lat == 0) && cxlDev.Cdat != nil {
		cdatPerf := cxlDev.Cdat.Get_CDAT_DSLBIS_performance()

		if bw == 0 {
			if cdatPerf.ReadBandwidthMBs != 0 {
				bw = float64(cdatPerf.ReadBandwidthMBs) / 1000
			} else {
				bw = float64(cdatPerf.AccessBandwidthMBs) / 1000
			}
		}

		if lat == 0 {
			if cdatPerf.ReadLatencyPs != 0 {
				lat = cdatPerf.ReadLatencyPs / 1000
			} else {
				lat = cdatPerf.AccessLatencyPs / 1000
			}
		}
	}

	measuredMemMetrics[cxlDev.GetBdfString()] = MemMetrics{Bandwidth: fmt.Sprintf("%.2f GiB/s", bw), Latency: fmt.Sprintf("%d ns", lat)}
	to <- true
}

// GetCXLMemPerf: Return performance data for CXL memory node
func GetCXLMemPerf(bdf string) *MemMetrics {
	memMetric, ok := measuredMemMetrics[bdf]
	if ok {
		return &memMetric
	}
	return nil
}

type memInfoKB struct {
	MemTotal int64
	MemFree  int64
	MemUsed  int64
}

// GetNumaMemInfo: return memory info for given numa node
func GetNumaMemInfo(node string) memInfoKB {
	memInfo := memInfoKB{}
	if node[:4] != "node" { // input either full node or just the number
		node = "node" + node
	}
	memInfoRaw, err := Readsystemfs(fmt.Sprintf("%s%s/meminfo", PlatformPath.SysDev, node))
	if err == nil {
		for _, line := range strings.Split(memInfoRaw, "\n") {
			lineSplit := strings.Fields(line)
			switch lineSplit[2] {
			case "MemTotal:":
				memInfo.MemTotal, _ = strconv.ParseInt(lineSplit[3], 10, 64)
			case "MemFree:":
				memInfo.MemFree, _ = strconv.ParseInt(lineSplit[3], 10, 64)
			case "MemUsed:":
				memInfo.MemUsed, _ = strconv.ParseInt(lineSplit[3], 10, 64)
			}
		}
	}
	return memInfo
}

////  ********** Platform Support **********  ////

var blockSizeByte int64

type PlatformConstants struct {
	SysDev  string
	SysMem  string
	BlkSize string
}

// Platform default to ubuntu
var PlatformPath = PlatformConstants{
	SysDev:  "/sys/bus/node/devices/",
	SysMem:  "/sys/devices/system/memory/",
	BlkSize: "/sys/devices/system/memory/block_size_bytes",
}

// TODO: Optimize this function when CICD is enabled and test under multiple distributions
// Initialize platform related constants based on system info
func init() {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	switch si.OS.Vendor {
	case "ubuntu":
		PlatformPath.SysDev = "/sys/bus/node/devices/"
	default:
		PlatformPath.SysDev = "/sys/bus/node/devices/"
	}
	blockSizeByte = initBlockSize()

	// Place for the memory performance tests
	measuredMemMetrics = map[string]MemMetrics{}
	for _, cxlDev := range cxlDevMap {
		fmt.Printf("Get Performance data for dev %s \n", cxlDev.GetBdfString())
		// Measure Performance; timeout after 10s
		to := make(chan bool, 1)
		go MeasurePerfMetrics(cxlDev, to)
		select {
		case <-to:
			fmt.Printf("Performance test done\n")
		case <-time.After(10 * time.Second):
			fmt.Printf("Performance test timeout\n")
		}
	}
}

// initBlockSize: Get the correct block size from sysfs
func initBlockSize() int64 {
	f, err := Readsystemfs(PlatformPath.BlkSize)
	if err != nil {
		fmt.Print("Failed to read system memory block size. Please check if there is pemission issue. CXL memory chunk will not be displayed..\n")
		return -1
	}
	size, err2 := strconv.ParseInt(f, 16, 64)
	if err2 != nil {
		fmt.Print("Failed to parse system memory block size. CXL memory chunk will not be displayed..\n")
		return -1
	}
	return size
}

// AsyncReboot: reboot the sytem asynchronously. The actual reboot system call has to be issued after the client receives the http response.
func AsyncReboot(bootType int) <-chan error {
	r := make(chan error)

	go func() {
		defer close(r)

		// Delay 5 seconds and issue the system call.
		time.Sleep(time.Second * 5)
		syscall.Sync()
		r <- syscall.Reboot(bootType)
	}()

	return r
}

////  ********** General helper functions **********  ////

// IdParse: Separate the text part and digital part. Eg. "DIMM1" -> "DIMM",1,nil
func IdParse(s string) (string, int, error) {
	var i int
	for i = 0; ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z'); i++ {
		continue
	}
	text := s[:i]
	num, err := strconv.Atoi(s[i:])
	return text, num, err
}
