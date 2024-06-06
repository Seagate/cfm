/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package serviceRequests

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceWrap"
	"fmt"
	"strings"

	service "cfm/pkg/client"

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

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.ip, r.ServiceTcp.port)

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
		if len(*memoryRegions) == 0 {
			fmt.Printf("%-25s %-15s No MemoryRegions Found!\n\n", key.ApplianceId, key.BladeId)
			continue
		}

		for i, memoryRegion := range *memoryRegions {
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
	serviceTcp   *TcpInfo
	applianceId  *Id
	bladeId      *Id
	portId       *Id
	resourceSize *Size
	qos          int32
}

//TODO: Should I propogate: private struct variables, forcing Getter usage -and- adding pointer safety to all Getter's (like generated cfm-service client)

func (r *ServiceRequestComposeMemory) GetServiceIp() string {
	return r.serviceTcp.GetIp()
}

func (r *ServiceRequestComposeMemory) GetServicePort() uint16 {
	return r.serviceTcp.GetPort()
}

func (r *ServiceRequestComposeMemory) GetApplianceId() string {
	return r.applianceId.GetId()
}

func (r *ServiceRequestComposeMemory) GetBladeId() string {
	return r.bladeId.GetId()
}

func (r *ServiceRequestComposeMemory) GetPortId() string {
	return r.portId.GetId()
}

func (r *ServiceRequestComposeMemory) GetResourceSizeGiB() int32 {
	return r.resourceSize.GetSizeGiB()
}

func (r *ServiceRequestComposeMemory) GetQos() int32 {
	return r.qos
}

func NewServiceRequestComposeMemory(cmd *cobra.Command) *ServiceRequestComposeMemory {

	qos, err := cmd.Flags().GetInt32(flags.MEMORY_QOS)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flags.MEMORY_QOS)
		cobra.CheckErr(err)
	}

	return &ServiceRequestComposeMemory{
		serviceTcp:   NewTcpInfo(cmd, flags.SERVICE),
		applianceId:  NewId(cmd, flags.APPLIANCE),
		bladeId:      NewId(cmd, flags.BLADE),
		portId:       NewId(cmd, flags.PORT),
		resourceSize: NewSize(cmd, flags.RESOURCE),
		qos:          qos,
	}
}

func (r *ServiceRequestComposeMemory) Execute() (*service.MemoryRegion, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "serviceTcp", fmt.Sprintf("%+v", *r.serviceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "applianceId", fmt.Sprintf("%+v", *r.applianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "bladeId", fmt.Sprintf("%+v", *r.bladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "portId", fmt.Sprintf("%+v", *r.portId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "resourceSize", fmt.Sprintf("%+v", r.resourceSize))

	serviceClient := serviceWrap.GetServiceClient(r.GetServiceIp(), r.GetServicePort())

	region, err := serviceWrap.ComposeMemory(serviceClient, r.GetApplianceId(), r.GetBladeId(), r.GetPortId(), r.GetResourceSizeGiB()*1024, r.GetQos())
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
	serviceTcp  *TcpInfo
	applianceId *Id
	bladeId     *Id
	memoryId    *Id
}

func (r *ServiceRequestFreeMemory) GetServiceIp() string {
	return r.serviceTcp.GetIp()
}

func (r *ServiceRequestFreeMemory) GetServicePort() uint16 {
	return r.serviceTcp.GetPort()
}

func (r *ServiceRequestFreeMemory) GetApplianceId() string {
	return r.applianceId.GetId()
}

func (r *ServiceRequestFreeMemory) GetBladeId() string {
	return r.bladeId.GetId()
}

func (r *ServiceRequestFreeMemory) GetMemoryId() string {
	return r.memoryId.GetId()
}

func NewServiceRequestFreeMemory(cmd *cobra.Command) *ServiceRequestFreeMemory {

	return &ServiceRequestFreeMemory{
		serviceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		applianceId: NewId(cmd, flags.APPLIANCE),
		bladeId:     NewId(cmd, flags.BLADE),
		memoryId:    NewId(cmd, flags.MEMORY),
	}
}

func (r *ServiceRequestFreeMemory) Execute() (*service.MemoryRegion, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "serviceTcp", fmt.Sprintf("%+v", *r.serviceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "applianceId", fmt.Sprintf("%+v", *r.applianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "bladeId", fmt.Sprintf("%+v", *r.bladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "memoryId", fmt.Sprintf("%+v", *r.memoryId))

	serviceClient := serviceWrap.GetServiceClient(r.GetServiceIp(), r.GetServicePort())

	region, err := serviceWrap.FreeMemory(serviceClient, r.GetApplianceId(), r.GetBladeId(), r.GetMemoryId())
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

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.ip, r.ServiceTcp.port)

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
		if len(*memoryRegions) == 0 {
			fmt.Printf("%-15s No MemoryRegions Found!\n\n", hostId)
			continue
		}

		for i, memoryRegion := range *memoryRegions {
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
