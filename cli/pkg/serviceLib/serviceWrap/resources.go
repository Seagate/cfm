// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

type ApplToBladeToResourceBlockMapType map[ApplianceBladeKey]*[]*service.MemoryResourceBlock
type ResourceBlockSummary struct {
	ApplToBladeMap                ApplToBladeMapType
	ApplToBladeToResourceBlockMap ApplToBladeToResourceBlockMapType
}

// Of the 2 summary maps, this completely fills the 1st map and then fully initializes, but doesn't fill, the 2nd map
func NewResourceBlockSummary(m ApplToBladeMapType) *ResourceBlockSummary {
	var summary ResourceBlockSummary

	summary.ApplToBladeMap = m
	summary.ApplToBladeToResourceBlockMap = make(ApplToBladeToResourceBlockMapType)

	for applId, blades := range m {
		for _, blade := range *blades {
			var resources []*service.MemoryResourceBlock

			key := NewApplianceBladeKey(applId, blade.GetId())
			summary.ApplToBladeToResourceBlockMap[*key] = &resources
		}
	}

	return &summary
}

// Add resource to resources map.  *Blade map is assumed already filled
func (s *ResourceBlockSummary) AddResource(applId, bladeId string, resource *service.MemoryResourceBlock) {
	key := NewApplianceBladeKey(applId, bladeId)

	*s.ApplToBladeToResourceBlockMap[*key] = append(*s.ApplToBladeToResourceBlockMap[*key], resource)
}

// Add multiple resources to resources map.  *Blade map is assumed already filled
func (s *ResourceBlockSummary) AddResourceSlice(applId, bladeId string, resources *[]*service.MemoryResourceBlock) {
	key := NewApplianceBladeKey(applId, bladeId)

	*s.ApplToBladeToResourceBlockMap[*key] = append(*s.ApplToBladeToResourceBlockMap[*key], *resources...)
}

func FindResourceBlockOnBlade(client *service.APIClient, applId, bladeId, resourceId string) (*service.MemoryResourceBlock, error) {
	var resourceBlock *service.MemoryResourceBlock

	requestResourceById := client.DefaultAPI.BladesGetResourceById(context.Background(), applId, bladeId, resourceId)
	resourceBlock, response, err := requestResourceById.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", requestResourceById)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: get appliance blade resource by id: %s", err)
	}

	klog.V(3).InfoS("BladesGetResourceById success", "applId", applId, "bladeId", bladeId, "resourceId", resourceBlock.GetId())

	return resourceBlock, nil
}

func GetAllResourceBlocksForBlade(client *service.APIClient, applId, bladeId string) (*[]*service.MemoryResourceBlock, error) {
	var resources []*service.MemoryResourceBlock

	requestResources := client.DefaultAPI.BladesGetResources(context.Background(), applId, bladeId)
	resourceColl, response, err := requestResources.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", requestResources)
		klog.ErrorS(err, msg, "response", response)
		// return nil, fmt.Errorf("failure: get appliance blade resources: %s", err)
		return &resources, nil //TODO: Error here instead?
	}

	klog.V(3).InfoS("BladesGetResources success", "applId", applId, "bladeId", bladeId, "resourceColl", resourceColl.GetMemberCount())

	for _, res := range resourceColl.GetMembers() {
		resourceId := ReadLastItemFromUri(res.GetUri())
		requestResourceById := client.DefaultAPI.BladesGetResourceById(context.Background(), applId, bladeId, resourceId)
		resourceBlock, response, err := requestResourceById.Execute()
		if err != nil {
			msg := fmt.Sprintf("%T: Execute FAILURE", requestResourceById)
			klog.ErrorS(err, msg, "response", response)
			// return nil, fmt.Errorf("failure: get appliance blade resource by id: %s", err)
			continue //TODO: Error here instead?
		}

		klog.V(3).InfoS("BladesGetResourceById success", "applId", applId, "bladeId", bladeId, "resourceId", resourceBlock.GetId())

		resources = append(resources, resourceBlock)
	}

	return &resources, nil
}

// /////////////////
// Gather all available ResourceBlocks from all available connected appliances and blades
func GetResourceBlocks_AllApplsAllBlades(client *service.APIClient) (*ResourceBlockSummary, error) {
	bladesSummary, err := GetAllBlades_AllAppls(client)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_AllAppls: %s", err)
	}

	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	for applId, blades := range summary.ApplToBladeMap {
		for _, blade := range *blades {
			resources, err := GetAllResourceBlocksForBlade(client, applId, blade.GetId())
			if err != nil {
				// return nil, fmt.Errorf("failure: GetAllResourceBlocksForBlade: %s", err)
				continue //TODO: Error here instead?
			}

			summary.AddResourceSlice(applId, blade.GetId(), resources)
		}
	}

	return summary, nil
}

// Find a specific ResourceBlock (identified by resource-id) from all connected blades and appliances
func FindResourceBlock_AllApplsAllBlades(client *service.APIClient, resourceId string) (*ResourceBlockSummary, error) {
	bladesSummary, err := GetAllBlades_AllAppls(client)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_AllAppls: %s", err)
	}

	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	for applId, blades := range summary.ApplToBladeMap {
		for _, blade := range *blades {
			resource, err := FindResourceBlockOnBlade(client, applId, blade.GetId(), resourceId)
			if err != nil {
				// return nil, fmt.Errorf("failure: FindResourceBlockOnBlade: %s", err)
				continue //TODO: Error here instead?
			}

			summary.AddResource(applId, blade.GetId(), resource)
		}
	}

	return summary, nil
}

// Gather all available ResourceBlocks from a specific Blade (identified by blade-id) that is present on all connected appliances.
func GetResourceBlocks_AllApplsSingleBlade(client *service.APIClient, bladeId string) (*ResourceBlockSummary, error) {
	bladesSummary, err := FindBladeById_AllAppls(client, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: %s", err)
	}

	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	for applId := range summary.ApplToBladeMap {
		resources, err := GetAllResourceBlocksForBlade(client, applId, bladeId)
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllResourceBlocksForBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddResourceSlice(applId, bladeId, resources)
	}

	return summary, nil
}

// Find a specific ResourceBlock (identified by resource-id) from a specific connected Blade (identified by blade-id) that is present on all available Appliances.
func FindResourceBlock_AllApplsSingleBlade(client *service.APIClient, bladeId, resourceId string) (*ResourceBlockSummary, error) {
	bladesSummary, err := FindBladeById_AllAppls(client, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: %s", err)
	}

	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	for applId := range summary.ApplToBladeMap {
		resource, err := FindResourceBlockOnBlade(client, applId, bladeId, resourceId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindResourceBlockOnBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddResource(applId, bladeId, resource)
	}

	return summary, nil
}

// Gather all available ResourceBlocks from all connected Blades from a specific Appliance (identified by appliance-id).
func GetResourceBlocks_SingleApplAllBlades(client *service.APIClient, applianceId string) (*ResourceBlockSummary, error) {
	blades, err := GetAllBlades_SingleAppl(client, applianceId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBladeSlice(applianceId, blades)
	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	for _, blade := range *summary.ApplToBladeMap[applianceId] {
		resources, err := GetAllResourceBlocksForBlade(client, applianceId, blade.GetId())
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllResourceBlocksForBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddResourceSlice(applianceId, blade.GetId(), resources)
	}

	return summary, nil
}

// Find a specific ResourceBlock (identified by resource-id) from all connected Blades from a specific Appliance (identified by appliance-id).
func FindResourceBlock_SingleApplAllBlades(client *service.APIClient, applianceId, resourceId string) (*ResourceBlockSummary, error) {
	blades, err := GetAllBlades_SingleAppl(client, applianceId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBladeSlice(applianceId, blades)
	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	for _, blade := range *summary.ApplToBladeMap[applianceId] {
		resource, err := FindResourceBlockOnBlade(client, applianceId, blade.GetId(), resourceId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindResourceBlockOnBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddResource(applianceId, blade.GetId(), resource)
	}

	return summary, nil
}

// Gather all available ResourceBlocks from a specific connected Blade (identified by blade-id) from a specific Appliance (identified by appliance-id).
func GetResourceBlocks_SingleApplSingleBlade(client *service.APIClient, applianceId, bladeId string) (*ResourceBlockSummary, error) {
	blade, err := FindBladeById_SingleAppl(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBlade(applianceId, blade)
	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	resources, err := GetAllResourceBlocksForBlade(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllResourceBlocksForBlade: %s", err)
	}

	summary.AddResourceSlice(applianceId, blade.GetId(), resources)

	return summary, nil
}

// Find a specific ResourceBlock (identified by resource-id) from a specific connected Blade (identified by blade-id) from a specific Appliance (identified by appliance-id).
func FindResourceBlock_SingleApplSingleBlade(client *service.APIClient, applianceId, bladeId, resourceId string) (*ResourceBlockSummary, error) {
	blade, err := FindBladeById_SingleAppl(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBlade(applianceId, blade)
	summary := NewResourceBlockSummary(bladesSummary.ApplToBladeMap)

	resource, err := FindResourceBlockOnBlade(client, applianceId, bladeId, resourceId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindResourceBlockOnBlade: %s", err)
	}

	summary.AddResource(applianceId, bladeId, resource)

	return summary, nil
}
