// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceRequests

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceWrap"
	"fmt"
	"strings"

	service "cfm/pkg/client"

	"github.com/facette/natsort"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

type ServiceRequestListMemoryRegions struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
	MemoryId    *Id
}

// If no appliance-id specified, default to listing ALL appliances
func (r *ServiceRequestListMemoryRegions) AllAppliances() bool {
	return !r.ApplianceId.HasId()
}

// If no blade-id specified, default to listing ALL blades
func (r *ServiceRequestListMemoryRegions) AllBlades() bool {
	return !r.BladeId.HasId()
}

// If no memory-id specified, default to listing ALL memoryRegions
func (r *ServiceRequestListMemoryRegions) AllMemoryRegions() bool {
	return !r.MemoryId.HasId()
}

func NewServiceRequestListMemoryRegions(cmd *cobra.Command) *ServiceRequestListMemoryRegions {

	return &ServiceRequestListMemoryRegions{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
		MemoryId:    NewId(cmd, flags.MEMORY),
	}
}

func (r *ServiceRequestListMemoryRegions) Execute() (*serviceWrap.BladeMemoryRegionSummary, error) {
	var summary *serviceWrap.BladeMemoryRegionSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "MemoryId", fmt.Sprintf("%+v", *r.MemoryId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	if r.AllAppliances() && r.AllBlades() && r.AllMemoryRegions() {
		summary, err = serviceWrap.GetMemoryRegions_AllApplsAllBlades(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(all appls, all blades, all memoryRegions): %s", err)
		}
	} else if r.AllAppliances() && r.AllBlades() && !r.AllMemoryRegions() {
		summary, err = serviceWrap.FindMemoryRegion_AllApplsAllBlades(serviceClient, r.MemoryId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(all appls, all blades, 1 memoryRegion): %s", err)
		}
	} else if r.AllAppliances() && !r.AllBlades() && r.AllMemoryRegions() {
		summary, err = serviceWrap.GetMemoryRegions_AllApplsSingleBlade(serviceClient, r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(all appls, 1 blades, all memoryRegions): %s", err)
		}
	} else if r.AllAppliances() && !r.AllBlades() && !r.AllMemoryRegions() {
		summary, err = serviceWrap.FindMemoryRegion_AllApplsSingleBlade(serviceClient, r.BladeId.GetId(), r.MemoryId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(all appls, 1 blade, 1 memoryRegion): %s", err)
		}
	} else if !r.AllAppliances() && r.AllBlades() && r.AllMemoryRegions() {
		summary, err = serviceWrap.GetMemoryRegions_SingleApplAllBlades(serviceClient, r.ApplianceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(1 appl, all blades, all memoryRegions): %s", err)
		}
	} else if !r.AllAppliances() && r.AllBlades() && !r.AllMemoryRegions() {
		summary, err = serviceWrap.FindMemoryRegion_SingleApplAllBlades(serviceClient, r.ApplianceId.GetId(), r.MemoryId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(1 appl, all blades, 1 memoryRegion): %s", err)
		}
	} else if !r.AllAppliances() && !r.AllBlades() && r.AllMemoryRegions() {
		summary, err = serviceWrap.GetMemoryRegions_SingleApplSingleBlade(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(1 appl, 1 blade, all memoryRegions): %s", err)
		}
	} else {
		summary, err = serviceWrap.FindMemoryRegion_SingleApplSingleBlade(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId(), r.MemoryId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(1 appl, 1 blade, 1 memoryRegion): %s", err)
		}
	}

	return summary, nil
}

func (r *ServiceRequestListMemoryRegions) OutputSummaryListMemory(s *serviceWrap.BladeMemoryRegionSummary) {
	var applianceId, bladeId string

	fmt.Printf("\n%-25s %-15s %-15s %-15s %-15s\n",
		"Appliance ID", "Blade ID", "Memory ID", "Size(MiB)", "Port ID")
	fmt.Printf("%s %s %s %s %s\n",
		strings.Repeat("-", 25), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15))

	if len(s.ApplToBladeMap) == 0 {
		fmt.Printf("No Appliances Found!\n\n")
	}

	for key, memoryRegions := range s.ApplToBladeToMemoryRegionsMap {
		num := len(*memoryRegions)
		if num == 0 {
			fmt.Printf("%-25s %-15s No MemoryRegions Found!\n\n", key.ApplianceId, key.BladeId)
			continue
		}

		memoryMap := make(map[string]*service.MemoryRegion, num)
		var memoryIds []string
		for _, memoryRegion := range *memoryRegions {
			id := memoryRegion.GetId()
			memoryIds = append(memoryIds, id)
			memoryMap[id] = memoryRegion
		}

		natsort.Sort(memoryIds)

		for i, id := range memoryIds {
			memoryRegion := memoryMap[id]
			if i == 0 {
				applianceId = key.ApplianceId
				bladeId = key.BladeId
			} else {
				applianceId = ""
				bladeId = ""
			}

			fmt.Printf("%-25s %-15s %-15s %-15d %-15s\n",
				applianceId, bladeId, memoryRegion.GetId(), memoryRegion.GetSizeMiB(), memoryRegion.GetMemoryAppliancePort())
		}
	}

	fmt.Printf("\n")
}

type ServiceRequestComposeMemory struct {
	ServiceTcp   *TcpInfo
	ApplianceId  *Id
	BladeId      *Id
	PortId       *Id
	ResourceSize *Size
	Qos          int32
}

func (r *ServiceRequestComposeMemory) GetQos() int32 {
	return r.Qos
}

func NewServiceRequestComposeMemory(cmd *cobra.Command) *ServiceRequestComposeMemory {

	qos, err := cmd.Flags().GetInt32(flags.MEMORY_QOS)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flags.MEMORY_QOS)
		cobra.CheckErr(err)
	}

	return &ServiceRequestComposeMemory{
		ServiceTcp:   NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId:  NewId(cmd, flags.APPLIANCE),
		BladeId:      NewId(cmd, flags.BLADE),
		PortId:       NewId(cmd, flags.PORT),
		ResourceSize: NewSize(cmd, flags.RESOURCE),
		Qos:          qos,
	}
}

func (r *ServiceRequestComposeMemory) Execute() (*service.MemoryRegion, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "PortId", fmt.Sprintf("%+v", *r.PortId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ResourceSize", fmt.Sprintf("%+v", r.ResourceSize))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	region, err := serviceWrap.ComposeMemory(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId(), r.PortId.GetId(), r.ResourceSize.GetSizeGiB()*1024, r.GetQos())
	if err != nil {
		return nil, fmt.Errorf("failure: compose memory: %s", err)
	}

	return region, err
}

func (r *ServiceRequestComposeMemory) OutputResultsComposedMemory(m *service.MemoryRegion) {
	fmt.Printf("\nCOMPOSE Memory Region Summary\n")
	fmt.Printf("Status: %s\n\n", m.GetStatus())
	fmt.Printf("%-15s %-15s %-15s %-15s %-25s\n", "Memory ID", "Size(MiB)", "Port ID", "Blade ID", "Appliance ID")
	fmt.Printf("%s %s %s %s %s\n", strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 25))
	fmt.Printf("%-15s %-15d %-15s %-15s %-25s\n", m.GetId(), m.GetSizeMiB(), m.GetMemoryAppliancePort(), m.GetMemoryBladeId(), m.GetMemoryApplianceId())

	fmt.Printf("\n")
}

type ServiceRequestFreeMemory struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
	MemoryId    *Id
}

func NewServiceRequestFreeMemory(cmd *cobra.Command) *ServiceRequestFreeMemory {

	return &ServiceRequestFreeMemory{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
		MemoryId:    NewId(cmd, flags.MEMORY),
	}
}

func (r *ServiceRequestFreeMemory) Execute() (*service.MemoryRegion, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "MemoryId", fmt.Sprintf("%+v", *r.MemoryId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	region, err := serviceWrap.FreeMemory(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId(), r.MemoryId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: free memory: %s", err)
	}

	return region, err
}

func (r *ServiceRequestFreeMemory) OutputResultsFreeMemory(m *service.MemoryRegion) {
	fmt.Printf("\nFREE Memory Region Summary\n")
	fmt.Printf("Status: %s\n\n", m.GetStatus())
	fmt.Printf("%-15s %-15s %-15s %-25s\n", "Memory ID", "Size(MiB)", "Blade ID", "Appliance ID")
	fmt.Printf("%s %s %s %s\n", strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 25))
	fmt.Printf("%-15s %-15d %-15s %-25s\n", m.GetId(), m.GetSizeMiB(), m.GetMemoryBladeId(), m.GetMemoryApplianceId())

	fmt.Printf("\n")
}

type ServiceRequestListHostMemoryRegions struct {
	ServiceTcp *TcpInfo
	HostId     *Id
	MemoryId   *Id
}

// If no host-id specified, default to listing ALL hosts
func (r *ServiceRequestListHostMemoryRegions) AllHosts() bool {
	return !r.HostId.HasId()
}

// If no memory-id specified, default to listing ALL memoryRegions
func (r *ServiceRequestListHostMemoryRegions) AllMemoryRegions() bool {
	return !r.MemoryId.HasId()
}

func NewServiceRequestListHostMemoryRegions(cmd *cobra.Command) *ServiceRequestListHostMemoryRegions {

	return &ServiceRequestListHostMemoryRegions{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
		HostId:     NewId(cmd, flags.HOST),
		MemoryId:   NewId(cmd, flags.MEMORY),
	}
}

func (r *ServiceRequestListHostMemoryRegions) Execute() (*serviceWrap.HostMemoryRegionSummary, error) {
	var summary *serviceWrap.HostMemoryRegionSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "MemoryId", fmt.Sprintf("%+v", *r.MemoryId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	if r.AllHosts() && r.AllMemoryRegions() {
		summary, err = serviceWrap.GetMemoryRegions_AllHosts(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(all hosts, all memoryRegions): %s", err)
		}
	} else if r.AllHosts() && !r.AllMemoryRegions() {
		summary, err = serviceWrap.FindMemoryRegion_AllHosts(serviceClient, r.MemoryId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(all hosts, 1 memoryRegion): %s", err)
		}
	} else if !r.AllHosts() && r.AllMemoryRegions() {
		summary, err = serviceWrap.GetMemoryRegions_SingleHost(serviceClient, r.HostId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(1 host, all memoryRegions): %s", err)
		}
	} else if !r.AllHosts() && !r.AllMemoryRegions() {
		summary, err = serviceWrap.FindMemoryRegion_SingleHost(serviceClient, r.HostId.GetId(), r.MemoryId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryRegions(1 host, 1 memoryRegion): %s", err)
		}
	}

	return summary, nil
}

func (r *ServiceRequestListHostMemoryRegions) OutputSummaryListMemory(s *serviceWrap.HostMemoryRegionSummary) {
	var outputHostId string

	fmt.Printf("\n%-15s %-15s %-15s %-15s\n",
		"Host ID", "Memory ID", "Size(MiB)", "Port ID")
	fmt.Printf("%s %s %s %s\n",
		strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15))

	if s.HostCount() == 0 {
		fmt.Printf("No Hosts Found!\n\n")
	}

	for hostId, memoryRegions := range s.MemoryRegions {
		num := len(*memoryRegions)
		if num == 0 {
			fmt.Printf("%-15s No MemoryRegions Found!\n\n", hostId)
			continue
		}

		memoryMap := make(map[string]*service.MemoryRegion, num)
		var memoryIds []string
		for _, memoryRegion := range *memoryRegions {
			id := memoryRegion.GetId()
			memoryIds = append(memoryIds, id)
			memoryMap[id] = memoryRegion
		}

		natsort.Sort(memoryIds)

		for i, id := range memoryIds {
			memoryRegion := memoryMap[id]
			if i == 0 {
				outputHostId = hostId
			} else {
				outputHostId = ""
			}

			fmt.Printf("%-15s %-15s %-15d %-15s\n",
				outputHostId, memoryRegion.GetId(), memoryRegion.GetSizeMiB(), memoryRegion.GetMemoryAppliancePort())
		}
	}

	fmt.Printf("\n")
}

// ServiceRequestBladesAssignMemory - This request structure supports BOTH "assign" and "unassign" of a memory region to\from a port
// This is an artifact of the way the cfm-service client api is setup.
type ServiceRequestBladesAssignMemory struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
	MemoryId    *Id
	PortId      *Id
	Operation   string //"assign" or "unassign"
}

func NewServiceRequestBladesAssignMemory(cmd *cobra.Command, operation string) *ServiceRequestBladesAssignMemory {

	if operation != "assign" && operation != "unassign" {
		newErr := fmt.Errorf("failure: NewServiceRequestBladesAssignMemory: operation options: 'assign', 'unassign'")
		klog.ErrorS(newErr, "Invalid parameter value", "operation", operation)
		cobra.CheckErr(newErr)
	}

	return &ServiceRequestBladesAssignMemory{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
		MemoryId:    NewId(cmd, flags.MEMORY),
		PortId:      NewId(cmd, flags.PORT),
		Operation:   operation,
	}
}

func (r *ServiceRequestBladesAssignMemory) GetOperation() string {
	return r.Operation
}

func (r *ServiceRequestBladesAssignMemory) Execute() (*service.MemoryRegion, error) {
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "MemoryId", fmt.Sprintf("%+v", *r.MemoryId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "PortId", fmt.Sprintf("%+v", *r.PortId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "Operation", r.Operation)

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	region, err := serviceWrap.BladesAssignMemory(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId(), r.MemoryId.GetId(), r.PortId.GetId(), r.GetOperation())
	if err != nil {
		newErr := fmt.Errorf("failure: assign blade memory region to blade port: %w", err)
		klog.ErrorS(newErr, "Execute failure", "memoryId", r.MemoryId.GetId(), "portId", r.PortId.GetId(), "applId", r.ApplianceId.GetId(), "bladeId", r.BladeId.GetId())
		return nil, newErr
	}

	return region, nil
}

func (r *ServiceRequestBladesAssignMemory) OutputSummaryBladesAssignMemory(m *service.MemoryRegion) {
	fmt.Printf("\n%s Memory and Port Summary\n", strings.ToUpper(r.GetOperation()))
	fmt.Printf("Status: %s\n\n", m.GetStatus())
	fmt.Printf("%-15s %-15s %-15s %-25s\n", "Memory ID", "Port ID", "Blade ID", "Appliance ID")
	fmt.Printf("%s %s %s %s\n", strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 25))
	fmt.Printf("%-15s %-15s %-15s %-25s\n", m.GetId(), m.GetMemoryAppliancePort(), m.GetMemoryBladeId(), m.GetMemoryApplianceId())

	fmt.Printf("\n")
}
