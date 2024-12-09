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

type ServiceRequestAddAppliance struct {
	ServiceTcp    *TcpInfo
	ApplianceId   *Id
	ApplianceCred *DeviceCredentials
	ApplianceTcp  *TcpInfo
}

// Create new add appliance request
// NOTE: The appliance cred and tcp and NOT USED within cfm-service but are currently REQUIRED by cfm-service to have something in them.
// Currently, just letting flag defaults pass through. This is an artifact of the Blade re-architecture.  Should remove.
func NewServiceRequestAddAppliance(cmd *cobra.Command) *ServiceRequestAddAppliance {
	return &ServiceRequestAddAppliance{
		ServiceTcp:    NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId:   NewId(cmd, flags.APPLIANCE),
		ApplianceCred: NewDeviceCredentials(cmd, flags.APPLIANCE),
		ApplianceTcp:  NewTcpInfo(cmd, flags.APPLIANCE),
	}
}

func (r *ServiceRequestAddAppliance) Execute() (*service.Appliance, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "AppliancesCred", fmt.Sprintf("%+v", *r.ApplianceCred))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceTcp", fmt.Sprintf("%+v", *r.ApplianceTcp))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	insecure := r.ApplianceTcp.GetInsecure()
	protocol := r.ApplianceTcp.GetProtocol()
	id := r.ApplianceId.GetId()

	creds := service.Credentials{
		Username:  r.ApplianceCred.GetUsername(),
		Password:  r.ApplianceCred.GetPassword(),
		IpAddress: r.ApplianceTcp.GetIp(),
		Port:      int32(r.ApplianceTcp.GetPort()),
		Insecure:  &insecure,
		Protocol:  &protocol,
		CustomId:  &id,
	}

	addedAppl, err := serviceWrap.AddAppliance(serviceClient, &creds)
	if err != nil {
		return nil, fmt.Errorf("failure: add appliance: %s", err)
	}

	return addedAppl, nil
}

type ServiceRequestDeleteAppliance struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
}

func NewServiceRequestDeleteAppliance(cmd *cobra.Command) *ServiceRequestDeleteAppliance {
	return &ServiceRequestDeleteAppliance{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
	}
}

func (r *ServiceRequestDeleteAppliance) Execute() (*service.Appliance, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	deletedAppliance, err := serviceWrap.DeleteApplianceById(serviceClient, r.ApplianceId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: delete appliance: %s", err)
	}

	return deletedAppliance, err
}

type ServiceRequestListAppliances ServiceRequestListExternalDevices

func NewServiceRequestListAppliances(cmd *cobra.Command) *ServiceRequestListAppliances {
	request := NewServiceRequestListExternalDevices(cmd)

	return &ServiceRequestListAppliances{
		ServiceTcp: request.ServiceTcp,
	}
}

func (r *ServiceRequestListAppliances) Execute() (*[]*service.Appliance, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.ip, r.ServiceTcp.port)

	appliances, err := serviceWrap.GetAllAppliances(serviceClient)
	if err != nil {
		return nil, fmt.Errorf("failure: list appliances: %s", err)
	}

	return appliances, err
}

func (r *ServiceRequestListAppliances) OutputResults(a *[]*service.Appliance) {

	if len(*a) == 0 {
		fmt.Printf("\nNo Appliances found\n\n")
		return
	}

	fmt.Printf("\n%-25s %-20s %-25s\n", "Appliance ID", "Free Memory (MiB)", "Composed Memory (MiB)")
	fmt.Printf("%s %s %s\n", strings.Repeat("-", 25), strings.Repeat("-", 20), strings.Repeat("-", 25))
	for _, appl := range *a {
		fmt.Printf("%-25s %-20d %-25d\n", appl.GetId(), appl.GetTotalMemoryAvailableMiB(), appl.GetTotalMemoryAllocatedMiB())
	}

	fmt.Printf("\n")
}

type ServiceRequestRenameAppliance struct {
	ServiceTcp     *TcpInfo
	ApplianceId    *Id
	NewApplianceId *Id
}

func NewServiceRequestRenameAppliance(cmd *cobra.Command) *ServiceRequestRenameAppliance {
	return &ServiceRequestRenameAppliance{
		ServiceTcp:     NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId:    NewId(cmd, flags.APPLIANCE),
		NewApplianceId: NewId(cmd, flags.NEW),
	}
}

func (r *ServiceRequestRenameAppliance) Execute() (*service.Appliance, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "NewApplianceId", fmt.Sprintf("%+v", *r.NewApplianceId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	appliance, err := serviceWrap.RenameApplianceById(serviceClient, r.ApplianceId.GetId(), r.NewApplianceId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: rename appliance: %s", err)
	}

	return appliance, err
}

type ServiceRequestResyncAppliance struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
}

func NewServiceRequestResyncAppliance(cmd *cobra.Command) *ServiceRequestResyncAppliance {
	return &ServiceRequestResyncAppliance{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
	}
}

func (r *ServiceRequestResyncAppliance) Execute() (*service.Appliance, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	appliance, err := serviceWrap.ResyncApplianceById(serviceClient, r.ApplianceId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: resync appliance: %s", err)
	}

	return appliance, err
}
