/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package serviceRequests

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceWrap"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

type ServiceRequestListHostMemoryDevices struct {
	ServiceTcp     *TcpInfo
	HostId         *Id
	MemoryDeviceId *Id
}

// If no host-id specified, default to listing ALL hosts
func (r *ServiceRequestListHostMemoryDevices) AllHosts() bool {
	return !r.HostId.HasId()
}

// If no memory-id specified, default to listing ALL memoryDevices
func (r *ServiceRequestListHostMemoryDevices) AllMemoryDevices() bool {
	return !r.MemoryDeviceId.HasId()
}

func NewServiceRequestListHostMemoryDevices(cmd *cobra.Command) *ServiceRequestListHostMemoryDevices {

	return &ServiceRequestListHostMemoryDevices{
		ServiceTcp:     NewTcpInfo(cmd, flags.SERVICE),
		HostId:         NewId(cmd, flags.HOST),
		MemoryDeviceId: NewId(cmd, flags.MEMORY_DEVICE),
	}
}

func (r *ServiceRequestListHostMemoryDevices) Execute() (*serviceWrap.HostMemoryDeviceSummary, error) {
	var summary *serviceWrap.HostMemoryDeviceSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "HostId", fmt.Sprintf("%+v", *r.HostId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "MemoryDeviceId", fmt.Sprintf("%+v", *r.MemoryDeviceId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.ip, r.ServiceTcp.port)

	if r.AllHosts() && r.AllMemoryDevices() {
		summary, err = serviceWrap.GetMemoryDevices_AllHosts(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryDevices(all hosts, all memoryDevices): %s", err)
		}
	} else if r.AllHosts() && !r.AllMemoryDevices() {
		summary, err = serviceWrap.FindMemoryDevice_AllHosts(serviceClient, r.MemoryDeviceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryDevices(all hosts, 1 memoryDevice): %s", err)
		}
	} else if !r.AllHosts() && r.AllMemoryDevices() {
		summary, err = serviceWrap.GetMemoryDevices_SingleHost(serviceClient, r.HostId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryDevices(1 host, all memoryDevices): %s", err)
		}
	} else if !r.AllHosts() && !r.AllMemoryDevices() {
		summary, err = serviceWrap.FindMemoryDevice_SingleHost(serviceClient, r.HostId.GetId(), r.MemoryDeviceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list memoryDevices(1 host, 1 memoryDevice): %s", err)
		}
	}

	return summary, nil
}

func (r *ServiceRequestListHostMemoryDevices) OutputSummaryListMemoryDevices(s *serviceWrap.HostMemoryDeviceSummary) {
	var outputHostId string

	fmt.Printf("\n%-15s %-15s\n",
		"Host ID", "MemoryDevice ID")
	fmt.Printf("%s %s\n",
		strings.Repeat("-", 15), strings.Repeat("-", 15))

	if s.HostCount() == 0 {
		fmt.Printf("No Hosts Found!\n\n")
	}

	for hostId, memoryDevices := range s.MemoryDevices {
		if len(*memoryDevices) == 0 {
			fmt.Printf("%-15s No MemoryDevices Found!\n\n", hostId)
			continue
		}

		for i, memoryDevice := range *memoryDevices {
			if i == 0 {
				outputHostId = hostId
			} else {
				outputHostId = ""
			}

			fmt.Printf("%-15s %-15s\n",
				outputHostId, memoryDevice.GetId())
		}
	}

	fmt.Printf("\n")
}
