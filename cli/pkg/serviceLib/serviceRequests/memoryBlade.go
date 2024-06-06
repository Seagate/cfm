//Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

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

// ServiceRequestBladesAssignMemory - This request structure supports BOTH "assign" and "unassign" of a memory region to\from a port
// This is an artifact of the way the cfm-service client api is setup.
type ServiceRequestBladesAssignMemory struct {
	serviceTcp  *TcpInfo
	applianceId *Id
	bladeId     *Id
	memoryId    *Id
	portId      *Id
	operation   string //"assign" or "unassign"
}

func NewServiceRequestBladesAssignMemory(cmd *cobra.Command, operation string) *ServiceRequestBladesAssignMemory {

	if operation != "assign" && operation != "unassign" {
		newErr := fmt.Errorf("Error: NewServiceRequestBladesAssignMemory: operation options: 'assign', 'unassign'")
		klog.ErrorS(newErr, "Invalid parameter value", "operation", operation)
		cobra.CheckErr(newErr)
	}

	return &ServiceRequestBladesAssignMemory{
		serviceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		applianceId: NewId(cmd, flags.APPLIANCE),
		bladeId:     NewId(cmd, flags.BLADE),
		memoryId:    NewId(cmd, flags.MEMORY),
		portId:      NewId(cmd, flags.PORT),
		operation:   operation,
	}
}

func (r *ServiceRequestBladesAssignMemory) GetServiceIp() string {
	return r.serviceTcp.GetIp()
}

func (r *ServiceRequestBladesAssignMemory) GetServicePort() uint16 {
	return r.serviceTcp.GetPort()
}

func (r *ServiceRequestBladesAssignMemory) GetApplianceId() string {
	return r.applianceId.GetId()
}

func (r *ServiceRequestBladesAssignMemory) GetBladeId() string {
	return r.bladeId.GetId()
}

func (r *ServiceRequestBladesAssignMemory) GetMemoryId() string {
	return r.memoryId.GetId()
}

func (r *ServiceRequestBladesAssignMemory) GetPortId() string {
	return r.portId.GetId()
}

func (r *ServiceRequestBladesAssignMemory) GetOperation() string {
	return r.operation
}

func (r *ServiceRequestBladesAssignMemory) Execute() (*service.MemoryRegion, error) {
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.serviceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.applianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.bladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "MemoryId", fmt.Sprintf("%+v", *r.memoryId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "PortId", fmt.Sprintf("%+v", *r.portId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "Operation", r.operation)

	serviceClient := serviceWrap.GetServiceClient(r.GetServiceIp(), r.GetServicePort())

	region, err := serviceWrap.BladesAssignMemory(serviceClient, r.GetApplianceId(), r.GetBladeId(), r.GetMemoryId(), r.GetPortId(), r.GetOperation())
	if err != nil {
		newErr := fmt.Errorf("failure: assign blade memory region to blade port: %w", err)
		klog.ErrorS(newErr, "Execute failure", "memoryId", r.GetMemoryId(), "portId", r.GetPortId(), "applId", r.GetApplianceId(), "bladeId", r.GetBladeId())
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
