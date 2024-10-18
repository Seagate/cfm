// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

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

type ServiceRequestAddBlade struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
	BladeCred   *DeviceCredentials
	BladeTcp    *TcpInfo
}

func NewServiceRequestAddBlade(cmd *cobra.Command) *ServiceRequestAddBlade {
	return &ServiceRequestAddBlade{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
		BladeCred:   NewDeviceCredentials(cmd, flags.DEVICE),
		BladeTcp:    NewTcpInfo(cmd, flags.DEVICE),
	}
}

func (r *ServiceRequestAddBlade) Execute() (*service.Blade, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeCred", fmt.Sprintf("%+v", *r.BladeCred))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeTcp", fmt.Sprintf("%+v", *r.BladeTcp))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	insecure := r.BladeTcp.GetInsecure()
	protocol := r.BladeTcp.GetProtocol()
	id := r.BladeId.GetId()

	bladeCreds := service.Credentials{
		Username:  r.BladeCred.GetUsername(),
		Password:  r.BladeCred.GetPassword(),
		IpAddress: r.BladeTcp.GetIp(),
		Port:      int32(r.BladeTcp.GetPort()),
		Insecure:  &insecure,
		Protocol:  &protocol,
		CustomId:  &id,
	}

	addedBlade, err := serviceWrap.AddBlade(serviceClient, r.ApplianceId.GetId(), &bladeCreds)
	if err != nil {
		return nil, fmt.Errorf("failure: add blade: %s", err)
	}

	return addedBlade, nil
}

type ServiceRequestDeleteBlade struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
}

func NewServiceRequestDeleteBlade(cmd *cobra.Command) *ServiceRequestDeleteBlade {
	return &ServiceRequestDeleteBlade{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
	}
}

func (r *ServiceRequestDeleteBlade) Execute() (*service.Blade, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	blade, err := serviceWrap.DeleteBladeById(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: delete blade: %s", err)
	}

	return blade, err
}

type ServiceRequestListBlades struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
}

func NewServiceRequestListBlades(cmd *cobra.Command) *ServiceRequestListBlades {
	return &ServiceRequestListBlades{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
	}
}

func (r *ServiceRequestListBlades) AllAppliances() bool {
	// Absence of appliance id indicates accessing all possible appliances
	return !r.ApplianceId.HasId()
}

func (r *ServiceRequestListBlades) AllBlades() bool {
	// Absence of blade id indicates accessing all possible blades
	return !r.BladeId.HasId()
}

func (r *ServiceRequestListBlades) Execute() (*serviceWrap.ApplianceBladeSummary, error) {
	var summary *serviceWrap.ApplianceBladeSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	if !r.AllAppliances() && !r.AllBlades() {
		blade, err := serviceWrap.FindBladeById_SingleAppl(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list blades(1 appl, 1 blade): %s", err)
		}
		summary = serviceWrap.NewApplianceBladeSummary()
		summary.AddBlade(r.ApplianceId.GetId(), blade)

	} else if !r.AllAppliances() && r.AllBlades() {
		blades, err := serviceWrap.GetAllBlades_SingleAppl(serviceClient, r.ApplianceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list blades(1 appl, all blade): %s", err)
		}
		summary = serviceWrap.NewApplianceBladeSummary()
		summary.AddBladeSlice(r.ApplianceId.GetId(), blades)

	} else if r.AllAppliances() && !r.AllBlades() {
		summary, err = serviceWrap.FindBladeById_AllAppls(serviceClient, r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list blades(all appl, 1 blade): %s", err)
		}

	} else {
		summary, err = serviceWrap.GetAllBlades_AllAppls(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list blades(all appl, all blade): %s", err)
		}
	}

	return summary, err
}

func (r *ServiceRequestListBlades) OutputResults(s *serviceWrap.ApplianceBladeSummary) {

	fmt.Printf("\n%-25s %-15s %-20s %-25s %-10s\n", "Appliance ID", "Blade ID", "Free Memory (MiB)", "Composed Memory (MiB)", "Status")
	fmt.Printf("%s %s %s %s %s\n", strings.Repeat("-", 25), strings.Repeat("-", 15), strings.Repeat("-", 20), strings.Repeat("-", 25), strings.Repeat("-", 10))
	if len(s.ApplToBladeMap) == 0 {
		fmt.Printf("\nNo Appliances found\n\n")
		return
	}

	for applId, blades := range s.ApplToBladeMap {
		if len(*blades) == 0 {
			fmt.Printf("%-25s %-15s\n", applId, "No Blades found")
			continue
		}
		for _, blade := range *blades {
			fmt.Printf("%-25s %-15s %-20d %-25d %-10s\n", applId, blade.GetId(), blade.GetTotalMemoryAvailableMiB(), blade.GetTotalMemoryAllocatedMiB(), blade.GetStatus())
		}
	}

	fmt.Printf("\n")
}

type ServiceRequestResyncBlade struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
}

func NewServiceRequestResyncBlade(cmd *cobra.Command) *ServiceRequestResyncBlade {
	return &ServiceRequestResyncBlade{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
	}
}

func (r *ServiceRequestResyncBlade) Execute() (*service.Blade, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	blade, err := serviceWrap.ResyncBladeById(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: resync blade: %s", err)
	}

	return blade, err
}
