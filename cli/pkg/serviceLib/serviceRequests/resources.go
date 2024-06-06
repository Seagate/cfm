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

type ServiceRequestListResources struct {
	ServiceTcp  *TcpInfo
	ApplianceId *Id
	BladeId     *Id
	ResourceId  *Id
}

// If no appliance-id specified, default to listing ALL appliances
func (r *ServiceRequestListResources) AllAppliances() bool {
	return !r.ApplianceId.HasId()
}

// If no blade-id specified, default to listing ALL blades
func (r *ServiceRequestListResources) AllBlades() bool {
	return !r.BladeId.HasId()
}

// If no resource-id specified, default to listing ALL resources
func (r *ServiceRequestListResources) AllResources() bool {
	return !r.ResourceId.HasId()
}

func NewServiceRequestListResources(cmd *cobra.Command) *ServiceRequestListResources {

	return &ServiceRequestListResources{
		ServiceTcp:  NewTcpInfo(cmd, flags.SERVICE),
		ApplianceId: NewId(cmd, flags.APPLIANCE),
		BladeId:     NewId(cmd, flags.BLADE),
		ResourceId:  NewId(cmd, flags.RESOURCE),
	}
}

func (r *ServiceRequestListResources) Execute() (*serviceWrap.ResourceBlockSummary, error) {
	var summary *serviceWrap.ResourceBlockSummary
	var err error

	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ServiceTcp", fmt.Sprintf("%+v", *r.ServiceTcp))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ApplianceId", fmt.Sprintf("%+v", *r.ApplianceId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "BladeId", fmt.Sprintf("%+v", *r.BladeId))
	klog.V(4).InfoS(fmt.Sprintf("%T", *r), "ResourceId", fmt.Sprintf("%+v", *r.ResourceId))

	serviceClient := serviceWrap.GetServiceClient(r.ServiceTcp.ip, r.ServiceTcp.port)

	if r.AllAppliances() && r.AllBlades() && r.AllResources() {
		summary, err = serviceWrap.GetResourceBlocks_AllApplsAllBlades(serviceClient)
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(all appls, all blades, all resources): %s", err)
		}
	} else if r.AllAppliances() && r.AllBlades() && !r.AllResources() {
		summary, err = serviceWrap.FindResourceBlock_AllApplsAllBlades(serviceClient, r.ResourceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(all appls, all blades, 1 resource): %s", err)
		}
	} else if r.AllAppliances() && !r.AllBlades() && r.AllResources() {
		summary, err = serviceWrap.GetResourceBlocks_AllApplsSingleBlade(serviceClient, r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(all appls, 1 blades, all resources): %s", err)
		}
	} else if r.AllAppliances() && !r.AllBlades() && !r.AllResources() {
		summary, err = serviceWrap.FindResourceBlock_AllApplsSingleBlade(serviceClient, r.BladeId.GetId(), r.ResourceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(all appls, 1 blade, 1 resource): %s", err)
		}
	} else if !r.AllAppliances() && r.AllBlades() && r.AllResources() {
		summary, err = serviceWrap.GetResourceBlocks_SingleApplAllBlades(serviceClient, r.ApplianceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(1 appl, all blades, all resources): %s", err)
		}
	} else if !r.AllAppliances() && r.AllBlades() && !r.AllResources() {
		summary, err = serviceWrap.FindResourceBlock_SingleApplAllBlades(serviceClient, r.ApplianceId.GetId(), r.ResourceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(1 appl, all blades, 1 resource): %s", err)
		}
	} else if !r.AllAppliances() && !r.AllBlades() && r.AllResources() {
		summary, err = serviceWrap.GetResourceBlocks_SingleApplSingleBlade(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(1 appl, 1 blade, all resources): %s", err)
		}
	} else {
		summary, err = serviceWrap.FindResourceBlock_SingleApplSingleBlade(serviceClient, r.ApplianceId.GetId(), r.BladeId.GetId(), r.ResourceId.GetId())
		if err != nil {
			return nil, fmt.Errorf("failure: list resources(1 appl, 1 blade, 1 resource): %s", err)
		}
	}

	return summary, nil
}

func (r *ServiceRequestListResources) OutputResults(s *serviceWrap.ResourceBlockSummary) {
	var applianceId, bladeId string

	fmt.Printf("\n%-25s %-15s %-15s %-15s %-15s %-10s %-10s\n",
		"Appliance ID", "Blade ID", "Resource ID", "Size (MiB)", "State", "Chn ID", "Chn Index")
	fmt.Printf("%s %s %s %s %s %s %s\n",
		strings.Repeat("-", 25), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 10), strings.Repeat("-", 10))

	if len(s.ApplToBladeMap) == 0 {
		fmt.Printf("No Appliances Found!\n\n")
	}

	for key, resources := range s.ApplToBladeToResourceBlockMap {
		if len(*resources) == 0 {
			fmt.Printf("%-25s %-15s No Resources Found!\n\n", key.ApplianceId, key.BladeId)
			continue
		}

		for i, resBlk := range *resources {
			status := resBlk.GetCompositionStatus()
			state := status.GetCompositionState()

			if i == 0 {
				applianceId = key.ApplianceId
				bladeId = key.BladeId
			} else {
				applianceId = ""
				bladeId = ""
			}

			fmt.Printf("%-25s %-15s %-15s %-15d %-15s %-10d %-10d\n",
				applianceId, bladeId, resBlk.GetId(), resBlk.GetCapacityMiB(), state, resBlk.GetChannelId(), resBlk.GetChannelResourceIndex())
		}
	}

	fmt.Printf("\n")
}
