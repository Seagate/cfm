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

type ServiceRequestAddHost struct {
	ServiceTcp *TcpInfo
	HostId     *Id
	HostCred   *DeviceCredentials
	HostTcp    *TcpInfo
}

func NewServiceRequestAddHost(cmd *cobra.Command) *ServiceRequestAddHost {
	return &ServiceRequestAddHost{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
		HostId:     NewId(cmd, flags.HOST),
		HostCred:   NewDeviceCredentials(cmd, flags.DEVICE),
		HostTcp:    NewTcpInfo(cmd, flags.DEVICE),
	}
}

func (r *ServiceRequestAddHost) Execute() (*service.Host, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostCred", fmt.Sprintf("%+v", *r.HostCred))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostTcp", fmt.Sprintf("%+v", *r.HostTcp))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	insecure := r.HostTcp.GetInsecure()
	protocol := r.HostTcp.GetProtocol()
	id := r.HostId.GetId()

	creds := service.Credentials{
		Username:  r.HostCred.GetUsername(),
		Password:  r.HostCred.GetPassword(),
		IpAddress: r.HostTcp.GetIp(),
		Port:      int32(r.HostTcp.GetPort()),
		Insecure:  &insecure,
		Protocol:  &protocol,
		CustomId:  &id,
	}

	addedHost, err := serviceWrap.AddHost(serviceClient, &creds)
	if err != nil {
		return nil, fmt.Errorf("failure: add host: %s", err)
	}

	return addedHost, nil
}

type ServiceRequestDeleteHost struct {
	ServiceTcp *TcpInfo
	HostId     *Id
}

func NewServiceRequestDeleteHost(cmd *cobra.Command) *ServiceRequestDeleteHost {
	return &ServiceRequestDeleteHost{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
		HostId:     NewId(cmd, flags.HOST),
	}
}

func (r *ServiceRequestDeleteHost) Execute() (*service.Host, error) {
	var host *service.Host
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	host, err = serviceWrap.DeleteHostById(serviceClient, r.HostId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: delete host: %s", err)
	}

	return host, err
}

type ServiceRequestListHosts ServiceRequestListExternalDevices

func NewServiceRequestListHosts(cmd *cobra.Command) *ServiceRequestListHosts {
	request := NewServiceRequestListExternalDevices(cmd)

	return &ServiceRequestListHosts{
		ServiceTcp: request.ServiceTcp,
	}
}

func (r *ServiceRequestListHosts) Execute() (*[]*service.Host, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.ip, r.ServiceTcp.port)

	hosts, err := serviceWrap.GetAllHosts(serviceClient)
	if err != nil {
		return nil, fmt.Errorf("failure: list hosts: %s", err)
	}

	return hosts, err
}

func (r *ServiceRequestListHosts) OutputResults(h *[]*service.Host) {

	if len(*h) == 0 {
		fmt.Printf("\nNo Hosts found\n\n")
		return
	}

	fmt.Printf("\n%-15s %-25s %-25s %-10s\n", "Host ID", "Local Memory (MiB)", "Remote Memory (MiB)", "Status")
	fmt.Printf("%s %s %s %s\n", strings.Repeat("-", 15), strings.Repeat("-", 25), strings.Repeat("-", 25), strings.Repeat("-", 10))
	for _, host := range *h {
		fmt.Printf("%-15s %-25d %-25d %-10s\n", host.GetId(), host.GetLocalMemoryMiB(), host.GetRemoteMemoryMiB(), host.GetStatus())
	}

	fmt.Printf("\n")
}

type ServiceRequestRenameHost struct {
	ServiceTcp *TcpInfo
	HostId     *Id
	NewHostId  *Id
}

func NewServiceRequestRenameHost(cmd *cobra.Command) *ServiceRequestRenameHost {
	return &ServiceRequestRenameHost{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
		HostId:     NewId(cmd, flags.HOST),
		NewHostId:  NewId(cmd, flags.NEW),
	}
}

func (r *ServiceRequestRenameHost) Execute() (*service.Host, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "NewHostId", fmt.Sprintf("%+v", *r.NewHostId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	host, err := serviceWrap.RenameHostById(serviceClient, r.HostId.GetId(), r.NewHostId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: rename host: %s", err)
	}

	return host, err
}

type ServiceRequestResyncHost struct {
	ServiceTcp *TcpInfo
	HostId     *Id
}

func NewServiceRequestResyncHost(cmd *cobra.Command) *ServiceRequestResyncHost {
	return &ServiceRequestResyncHost{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
		HostId:     NewId(cmd, flags.HOST),
	}
}

func (r *ServiceRequestResyncHost) Execute() (*service.Host, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.GetIp(), r.ServiceTcp.GetPort())

	host, err := serviceWrap.ResyncHostById(serviceClient, r.HostId.GetId())
	if err != nil {
		return nil, fmt.Errorf("failure: resync host: %s", err)
	}

	return host, err
}
