// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceRequests

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceWrap"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

type ServiceRequestListBladePorts struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
	PortId      *Id
}

// If no appliance-id specified, default to listing ALL appliances
func (r *ServiceRequestListBladePorts) AllAppliances() bool {
	return !r.ApplianceId.HasId()
}

// If no blade-id specified, default to listing ALL blades
func (r *ServiceRequestListBladePorts) AllBlades() bool {
	return !r.BladeId.HasId()
}

// If no port-id specified, default to listing ALL ports
func (r *ServiceRequestListBladePorts) AllPorts() bool {
	return !r.PortId.HasId()
}

func NewServiceRequestListBladePorts(cmd *cobra.Command) *ServiceRequestListBladePorts {

	return &ServiceRequestListBladePorts{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
		PortId:      NewId(cmd, flags.PORT),
	}
}

func (r *ServiceRequestListBladePorts) Execute() (*serviceWrap.BladePortsSummary, error) {
	var summary *serviceWrap.BladePortsSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "PortId", fmt.Sprintf("%+v", *r.PortId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	if r.AllAppliances() && r.AllBlades() && r.AllPorts() {
		summary, err = serviceWrap.GetPorts_AllApplsAllBlades(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(all appls, all blades, all ports): %s", err)
		}
	} else if r.AllAppliances() && r.AllBlades() && !r.AllPorts() {
		summary, err = serviceWrap.FindPort_AllApplsAllBlades(serviceClient, r.PortId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(all appls, all blades, 1 port): %s", err)
		}
	} else if r.AllAppliances() && !r.AllBlades() && r.AllPorts() {
		summary, err = serviceWrap.GetPorts_AllApplsSingleBlade(serviceClient, r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(all appls, 1 blades, all ports): %s", err)
		}
	} else if r.AllAppliances() && !r.AllBlades() && !r.AllPorts() {
		summary, err = serviceWrap.FindPort_AllApplsSingleBlade(serviceClient, r.BladeId.GetId(), r.PortId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(all appls, 1 blade, 1 port): %s", err)
		}
	} else if !r.AllAppliances() && r.AllBlades() && r.AllPorts() {
		summary, err = serviceWrap.GetPorts_SingleApplAllBlades(serviceClient, r.ApplianceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(1 appl, all blades, all ports): %s", err)
		}
	} else if !r.AllAppliances() && r.AllBlades() && !r.AllPorts() {
		summary, err = serviceWrap.FindPort_SingleApplAllBlades(serviceClient, r.ApplianceId.GetId(), r.PortId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(1 appl, all blades, 1 port): %s", err)
		}
	} else if !r.AllAppliances() && !r.AllBlades() && r.AllPorts() {
		summary, err = serviceWrap.GetPorts_SingleApplSingleBlade(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(1 appl, 1 blade, all ports): %s", err)
		}
	} else {
		summary, err = serviceWrap.FindPort_SingleApplSingleBlade(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId(), r.PortId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(1 appl, 1 blade, 1 port): %s", err)
		}
	}

	return summary, nil
}

func (r *ServiceRequestListBladePorts) OutputSummary(s *serviceWrap.BladePortsSummary) {
	var applianceId, bladeId string

	fmt.Printf("\n%-25s %-15s %-15s %-30s %-30s %-10s\n",
		"Appliance ID", "Blade ID", "Port ID", "GCxlId", "Linkage (Host,Port)", "LinkStatus")
	fmt.Printf("%s %s %s %s %s %s\n",
		strings.Repeat("-", 25), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 30), strings.Repeat("-", 30), strings.Repeat("-", 10))

	if len(s.ApplToBladeMap) == 0 {
		fmt.Printf("No Appliances Found!\n\n")
	}

	for key, ports := range s.ApplToBladeToPortsMap {
		if len(*ports) == 0 {
			fmt.Printf("%-25s %-15s No Ports Found!\n\n", key.ApplianceId, key.BladeId)
			continue
		}

		for i, port := range *ports {
			if i == 0 {
				applianceId = key.ApplianceId
				bladeId = key.BladeId
			} else {
				applianceId = ""
				bladeId = ""
			}

			var linkage string
			uri := port.GetLinkedPortUri()
			if len(uri) > 0 {
				tokens := strings.Split(uri, "/")
				if len(tokens) > 6 {
					linkedHost := tokens[4]
					linkedPort := tokens[6]
					linkage = fmt.Sprintf("%s,%s", linkedHost, linkedPort)
				}
			}

			fmt.Printf("%-25s %-15s %-15s %-30s %-30s %-10s\n",
				applianceId, bladeId, port.GetId(), port.GetGCxlId(), linkage, port.GetLinkStatus())
		}
	}

	fmt.Printf("\n")
}

type ServiceRequestListHostPorts struct {
	ServiceTcp *TcpInfo
	HostId     *Id
	PortId     *Id
}

// If no blade-id specified, default to listing ALL blades
func (r *ServiceRequestListHostPorts) AllHosts() bool {
	return !r.HostId.HasId()
}

// If no port-id specified, default to listing ALL ports
func (r *ServiceRequestListHostPorts) AllPorts() bool {
	return !r.PortId.HasId()
}

func NewServiceRequestListHostPorts(cmd *cobra.Command) *ServiceRequestListHostPorts {

	return &ServiceRequestListHostPorts{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
		HostId:     NewId(cmd, flags.HOST),
		PortId:     NewId(cmd, flags.PORT),
	}
}

func (r *ServiceRequestListHostPorts) Execute() (*serviceWrap.HostPortSummary, error) {
	var summary *serviceWrap.HostPortSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "PortId", fmt.Sprintf("%+v", *r.PortId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort(), r.ServiceTcp.GetInsecure(), r.ServiceTcp.GetProtocol())

	if r.AllHosts() && r.AllPorts() {
		summary, err = serviceWrap.GetAllPorts_AllHosts(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(all hosts, all ports): %s", err)
		}
	} else if r.AllHosts() && !r.AllPorts() {
		summary, err = serviceWrap.FindPortById_AllHosts(serviceClient, r.PortId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(all hosts, 1 port): %s", err)
		}
	} else if !r.AllHosts() && r.AllPorts() {
		ports, err := serviceWrap.GetAllPorts_SingleHost(serviceClient, r.HostId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(1 host, all ports): %s", err)
		}
		summary = serviceWrap.NewHostPortSummary()
		summary.AddPortSlice(r.HostId.GetId(), ports)
	} else {
		port, err := serviceWrap.FindPortById_SingleHost(serviceClient, r.HostId.GetId(), r.PortId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list ports(1 host, 1 port): %s", err)
		}
		summary = serviceWrap.NewHostPortSummary()
		summary.AddPort(r.HostId.GetId(), port)
	}

	return summary, nil
}

func (r *ServiceRequestListHostPorts) OutputSummary(s *serviceWrap.HostPortSummary) {
	var outputHostId string

	fmt.Printf("\n%-15s %-15s %-30s %-30s %-10s\n",
		"Host ID", "Port ID", "GCxlId", "Linkage(Appliance,Blade,Port)", "LinkStatus")
	fmt.Printf("%s %s %s %s %s\n",
		strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 30), strings.Repeat("-", 30), strings.Repeat("-", 10))

	if len(s.Ports) == 0 {
		fmt.Printf("No Hosts Found!\n\n")
	}

	for hostId, ports := range s.Ports {
		if len(*ports) == 0 {
			fmt.Printf("%-15s No Ports Found!\n\n", hostId)
			continue
		}

		for i, port := range *ports {
			if i == 0 {
				outputHostId = hostId
			} else {
				outputHostId = ""
			}

			var linkage string
			uri := port.GetLinkedPortUri()
			if len(uri) > 0 {
				tokens := strings.Split(uri, "/")
				if len(tokens) > 8 {
					linkedAppliance := tokens[4]
					linkedBlade := tokens[6]
					linkedPort := tokens[8]
					linkage = fmt.Sprintf("%s,%s,%s", linkedAppliance, linkedBlade, linkedPort)
				}
			}

			fmt.Printf("%-15s %-15s %-30s %-30s %-10s\n",
				outputHostId, port.GetId(), port.GetGCxlId(), linkage, port.GetLinkStatus())
		}
	}

	fmt.Printf("\n")
}
