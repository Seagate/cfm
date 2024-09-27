// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

type ApplToBladeToMemoryRegionsMapType map[ApplianceBladeKey]*[]*service.MemoryRegion
type BladeMemoryRegionSummary struct {
	ApplToBladeMap                ApplToBladeMapType
	ApplToBladeToMemoryRegionsMap ApplToBladeToMemoryRegionsMapType
}

// Of the 2 summary maps, this completely fills the 1st map and then fully initializes, but doesn't fill, the 2nd map
func NewBladeMemoryRegionsSummary(m ApplToBladeMapType) *BladeMemoryRegionSummary {
	var summary BladeMemoryRegionSummary

	summary.ApplToBladeMap = m
	summary.ApplToBladeToMemoryRegionsMap = make(ApplToBladeToMemoryRegionsMapType)

	for applId, blades := range m {
		for _, blade := range *blades {
			var memoryRegions []*service.MemoryRegion

			key := NewApplianceBladeKey(applId, blade.GetId())
			summary.ApplToBladeToMemoryRegionsMap[*key] = &memoryRegions
		}
	}

	return &summary
}

// Add memoryRegion to memoryRegions map.  *Blade map is assumed already filled
func (s *BladeMemoryRegionSummary) AddMemoryRegion(applId, bladeId string, memoryRegion *service.MemoryRegion) {
	key := NewApplianceBladeKey(applId, bladeId)

	*s.ApplToBladeToMemoryRegionsMap[*key] = append(*s.ApplToBladeToMemoryRegionsMap[*key], memoryRegion)
}

// Add multiple memoryRegions to memoryRegions map.  *Blade map is assumed already filled
func (s *BladeMemoryRegionSummary) AddMemoryRegionSlice(applId, bladeId string, memoryRegions *[]*service.MemoryRegion) {
	key := NewApplianceBladeKey(applId, bladeId)

	*s.ApplToBladeToMemoryRegionsMap[*key] = append(*s.ApplToBladeToMemoryRegionsMap[*key], *memoryRegions...)
}

func FindMemoryRegionOnBlade(client *service.APIClient, applId, bladeId, memoryId string) (*service.MemoryRegion, error) {
	var memoryRegion *service.MemoryRegion

	request := client.DefaultAPI.BladesGetMemoryById(context.Background(), applId, bladeId, memoryId)
	memoryRegion, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: FindMemoryRegionOnBlade")

			return nil, fmt.Errorf("failure: FindMemoryRegionOnBlade: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FindMemoryRegionOnBlade")

		return nil, fmt.Errorf("failure: FindMemoryRegionOnBlade: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("success: FindMemoryRegionOnBlade", "applId", applId, "bladeId", bladeId, "memoryId", memoryRegion.GetId())

	return memoryRegion, nil
}

func GetAllMemoryRegionsForBlade(client *service.APIClient, applId, bladeId string) (*[]*service.MemoryRegion, error) {
	var memoryRegions []*service.MemoryRegion

	requestMemoryRegions := client.DefaultAPI.BladesGetMemory(context.Background(), applId, bladeId)
	memoryRegionColl, response, err := requestMemoryRegions.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestMemoryRegions, err)
			klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForBlade")

			return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestMemoryRegions, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForBlade")

		// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s (%s)", status.Status.Message, err)
		return &memoryRegions, nil //TODO: Error here instead?
	}

	klog.V(4).InfoS("success: BladesGetMemory", "applId", applId, "bladeId", bladeId, "memoryRegionColl", memoryRegionColl.GetMemberCount())

	for _, res := range memoryRegionColl.GetMembers() {
		memoryId := ReadLastItemFromUri(res.GetUri())
		requestMemoryRegionById := client.DefaultAPI.BladesGetMemoryById(context.Background(), applId, bladeId, memoryId)
		memoryRegion, response, err := requestMemoryRegionById.Execute()
		if err != nil {
			// Decode the JSON response into a struct
			var status service.StatusMessage
			if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
				newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestMemoryRegionById, err)
				klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForBlade")

				return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s", newErr)
			}

			newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
				requestMemoryRegionById, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
			klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForBlade")

			// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s (%s)", status.Status.Message, err)
			continue //TODO: Error here instead?
		}

		klog.V(4).InfoS("success: BladesGetMemoryById", "applId", applId, "bladeId", bladeId, "memoryId", memoryRegion.GetId())

		memoryRegions = append(memoryRegions, memoryRegion)
	}

	return &memoryRegions, nil
}

// /////////////////
// Gather all available MemoryRegions from all available connected appliances and blades
func GetMemoryRegions_AllApplsAllBlades(client *service.APIClient) (*BladeMemoryRegionSummary, error) {
	bladesSummary, err := GetAllBlades_AllAppls(client)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_AllAppls: %s", err)
	}

	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	for applId, blades := range summary.ApplToBladeMap {
		for _, blade := range *blades {
			memoryRegions, err := GetAllMemoryRegionsForBlade(client, applId, blade.GetId())
			if err != nil {
				// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s", err)
				continue //TODO: Error here instead?
			}

			summary.AddMemoryRegionSlice(applId, blade.GetId(), memoryRegions)
		}
	}

	return summary, nil
}

// Find a specific MemoryRegion (identified by memory-Id) from all connected blades and appliances
func FindMemoryRegion_AllApplsAllBlades(client *service.APIClient, memoryId string) (*BladeMemoryRegionSummary, error) {
	bladesSummary, err := GetAllBlades_AllAppls(client)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_AllAppls: %s", err)
	}

	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	for applId, blades := range summary.ApplToBladeMap {
		for _, blade := range *blades {
			memoryRegion, err := FindMemoryRegionOnBlade(client, applId, blade.GetId(), memoryId)
			if err != nil {
				// return nil, fmt.Errorf("failure: FindMemoryRegionOnBlade: %s", err)
				continue //TODO: Error here instead?
			}

			summary.AddMemoryRegion(applId, blade.GetId(), memoryRegion)
		}
	}

	return summary, nil
}

// Gather all available MemoryRegions from a specific Blade (identified by blade-id) that is present on all connected appliances.
func GetMemoryRegions_AllApplsSingleBlade(client *service.APIClient, bladeId string) (*BladeMemoryRegionSummary, error) {
	bladesSummary, err := FindBladeById_AllAppls(client, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: %s", err)
	}

	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	for applId := range summary.ApplToBladeMap {
		memoryRegions, err := GetAllMemoryRegionsForBlade(client, applId, bladeId)
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryRegionSlice(applId, bladeId, memoryRegions)
	}

	return summary, nil
}

// Find a specific MemoryRegion (identified by memory-Id) from a specific connected Blade (identified by blade-id) that is present on all available Appliances.
func FindMemoryRegion_AllApplsSingleBlade(client *service.APIClient, bladeId, memoryId string) (*BladeMemoryRegionSummary, error) {
	bladesSummary, err := FindBladeById_AllAppls(client, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: %s", err)
	}

	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	for applId := range summary.ApplToBladeMap {
		memoryRegion, err := FindMemoryRegionOnBlade(client, applId, bladeId, memoryId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindMemoryRegionOnBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryRegion(applId, bladeId, memoryRegion)
	}

	return summary, nil
}

// Gather all available MemoryRegions from all connected Blades from a specific Appliance (identified by appliance-id).
func GetMemoryRegions_SingleApplAllBlades(client *service.APIClient, applianceId string) (*BladeMemoryRegionSummary, error) {
	blades, err := GetAllBlades_SingleAppl(client, applianceId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBladeSlice(applianceId, blades)
	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	for _, blade := range *summary.ApplToBladeMap[applianceId] {
		memoryRegions, err := GetAllMemoryRegionsForBlade(client, applianceId, blade.GetId())
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryRegionSlice(applianceId, blade.GetId(), memoryRegions)
	}

	return summary, nil
}

// Find a specific MemoryRegion (identified by memory-Id) from all connected Blades from a specific Appliance (identified by appliance-id).
func FindMemoryRegion_SingleApplAllBlades(client *service.APIClient, applianceId, memoryId string) (*BladeMemoryRegionSummary, error) {
	blades, err := GetAllBlades_SingleAppl(client, applianceId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBladeSlice(applianceId, blades)
	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	for _, blade := range *summary.ApplToBladeMap[applianceId] {
		memoryRegion, err := FindMemoryRegionOnBlade(client, applianceId, blade.GetId(), memoryId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindMemoryRegionOnBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryRegion(applianceId, blade.GetId(), memoryRegion)
	}

	return summary, nil
}

// Gather all available MemoryRegions from a specific connected Blade (identified by blade-id) from a specific Appliance (identified by appliance-id).
func GetMemoryRegions_SingleApplSingleBlade(client *service.APIClient, applianceId, bladeId string) (*BladeMemoryRegionSummary, error) {
	blade, err := FindBladeById_SingleAppl(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBlade(applianceId, blade)
	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	memoryRegions, err := GetAllMemoryRegionsForBlade(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllMemoryRegionsForBlade: %s", err)
	}

	summary.AddMemoryRegionSlice(applianceId, blade.GetId(), memoryRegions)

	return summary, nil
}

// Find a specific MemoryRegion (identified by memory-Id) from a specific connected Blade (identified by blade-id) from a specific Appliance (identified by appliance-id).
func FindMemoryRegion_SingleApplSingleBlade(client *service.APIClient, applianceId, bladeId, memoryId string) (*BladeMemoryRegionSummary, error) {
	blade, err := FindBladeById_SingleAppl(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBlade(applianceId, blade)
	summary := NewBladeMemoryRegionsSummary(bladesSummary.ApplToBladeMap)

	memoryRegion, err := FindMemoryRegionOnBlade(client, applianceId, bladeId, memoryId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindMemoryRegionOnBlade: %s", err)
	}

	summary.AddMemoryRegion(applianceId, bladeId, memoryRegion)

	return summary, nil
}

// Send compose request via cfm-service client
func ComposeMemory(client *service.APIClient, applianceId, bladeId, portId string, resourceSizeMiB int32, qosInt int32) (*service.MemoryRegion, error) {

	// create QoS struct
	qos, _ := service.NewQosFromValue(qosInt)

	// create new ComposeMemoryRequest
	requestCompose := service.NewComposeMemoryRequest(resourceSizeMiB, *qos)
	requestCompose.SetPort(portId)

	// create new ApiBladesComposeMemoryRequest
	requestApiCompose := client.DefaultAPI.BladesComposeMemory(context.Background(), applianceId, bladeId)

	// add ComposeMemoryRequest to ApiAppliancesComposeMemoryRequest
	requestApiCompose = requestApiCompose.ComposeMemoryRequest(*requestCompose)

	// Execute ApiAppliancesComposeMemoryRequest
	region, response, err := requestApiCompose.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestApiCompose, err)
			klog.ErrorS(newErr, "failure: ComposeMemory")

			return nil, fmt.Errorf("failure: ComposeMemory: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestApiCompose, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: ComposeMemory")

		return nil, fmt.Errorf("failure: ComposeMemory: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("success: ComposeMemory", "region", region.GetId(), "size", region.GetSizeMiB(), "applId", applianceId, "bladeId", bladeId)

	return region, nil
}

// Send free memory request
func FreeMemory(client *service.APIClient, applianceId, bladeId, memoryId string) (*service.MemoryRegion, error) {

	// create new ApiAppliancesFreeMemoryByIdRequest
	requestFree := client.DefaultAPI.BladesFreeMemoryById(context.Background(), applianceId, bladeId, memoryId)
	// Execute ApiAppliancesFreeMemoryByIdRequest
	region, response, err := requestFree.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestFree, err)
			klog.ErrorS(newErr, "failure: FreeMemory")

			return nil, fmt.Errorf("failure: FreeMemory: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestFree, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FreeMemory")

		return nil, fmt.Errorf("failure: FreeMemory: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("success: FreeMemory", "region", region.GetId(), "size", region.GetSizeMiB(), "applId", applianceId, "bladeId", bladeId)

	return region, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////

type HostToMemoryRegionsMapType map[string]*[]*service.MemoryRegion
type HostMemoryRegionSummary struct {
	MemoryRegions HostToMemoryRegionsMapType
}

func NewHostMemoryRegionSummary() *HostMemoryRegionSummary {
	var summary HostMemoryRegionSummary

	summary.MemoryRegions = make(HostToMemoryRegionsMapType)

	return &summary
}

func (s *HostMemoryRegionSummary) AddHost(hostId string) {
	_, found := s.MemoryRegions[hostId]
	if !found {
		var memory []*service.MemoryRegion
		s.MemoryRegions[hostId] = &memory
	}
}

// Add memoryRegion to memoryRegions map.
func (s *HostMemoryRegionSummary) AddMemoryRegion(hostId string, memoryRegion *service.MemoryRegion) {
	s.AddHost(hostId)

	*s.MemoryRegions[hostId] = append(*s.MemoryRegions[hostId], memoryRegion)
}

// Add multiple memoryRegions to memoryRegions map.
func (s *HostMemoryRegionSummary) AddMemoryRegionSlice(hostId string, memoryRegions *[]*service.MemoryRegion) {
	s.AddHost(hostId)

	*s.MemoryRegions[hostId] = append(*s.MemoryRegions[hostId], *memoryRegions...)
}

func (s *HostMemoryRegionSummary) HostCount() int {

	return len(s.MemoryRegions)
}

func FindMemoryRegionOnHost(client *service.APIClient, hostId, memoryId string) (*service.MemoryRegion, error) {
	var memoryRegion *service.MemoryRegion

	request := client.DefaultAPI.HostsGetMemoryById(context.Background(), hostId, memoryId)
	memoryRegion, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: FindMemoryRegionOnHost")

			return nil, fmt.Errorf("failure: FindMemoryRegionOnHost: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FindMemoryRegionOnHost")

		return nil, fmt.Errorf("failure: FindMemoryRegionOnHost: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("success: FindMemoryRegionOnHost", "hostId", hostId, "memoryId", memoryRegion.GetId())

	return memoryRegion, nil
}

func GetAllMemoryRegionsForHost(client *service.APIClient, hostId string) (*[]*service.MemoryRegion, error) {
	var memoryRegions []*service.MemoryRegion

	request := client.DefaultAPI.HostGetMemory(context.Background(), hostId)
	memoryRegionColl, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForHost")

			return nil, fmt.Errorf("failure: GetAllMemoryRegionsForHost: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForHost")

		// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForHost: %s (%s)", status.Status.Message, err)
		return &memoryRegions, nil //TODO: Error here instead?
	}

	klog.V(4).InfoS("success: HostGetMemory", "hostId", hostId, "memoryRegionColl", memoryRegionColl.GetMemberCount())

	for _, member := range memoryRegionColl.GetMembers() {
		memoryId := ReadLastItemFromUri(member.GetUri())
		request := client.DefaultAPI.HostsGetMemoryById(context.Background(), hostId, memoryId)
		memoryRegion, response, err := request.Execute()
		if err != nil {
			// Decode the JSON response into a struct
			var status service.StatusMessage
			if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
				newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
				klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForHost")

				return nil, fmt.Errorf("failure: GetAllMemoryRegionsForHost: %s", newErr)
			}

			newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
				request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
			klog.ErrorS(newErr, "failure: GetAllMemoryRegionsForHost")

			// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForHost: %s (%s)", status.Status.Message, err)
			continue //TODO: Error here instead?
		}

		klog.V(4).InfoS("success: HostsGetMemoryById", "hostId", hostId, "memoryId", memoryRegion.GetId())

		memoryRegions = append(memoryRegions, memoryRegion)
	}

	return &memoryRegions, nil
}

// Gather all available MemoryRegions from all avaiable Hosts.
func GetMemoryRegions_AllHosts(client *service.APIClient) (*HostMemoryRegionSummary, error) {
	summary := NewHostMemoryRegionSummary()

	request := client.DefaultAPI.HostsGet(context.Background())
	hostColl, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: GetMemoryRegions_AllHosts")

			return nil, fmt.Errorf("failure: GetMemoryRegions_AllHosts: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: GetMemoryRegions_AllHosts")

		// return nil, fmt.Errorf("failure: GetMemoryRegions_AllHosts: %s (%s)", status.Status.Message, err)
		return summary, nil //TODO: Error here instead?
	}

	for _, member := range hostColl.GetMembers() {
		hostId := ReadLastItemFromUri(member.GetUri())
		memoryRegions, err := GetAllMemoryRegionsForHost(client, hostId)
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllMemoryRegionsForHost: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryRegionSlice(hostId, memoryRegions)
	}

	return summary, nil
}

// Find a specific MemoryRegion from all available Hosts.
func FindMemoryRegion_AllHosts(client *service.APIClient, memoryId string) (*HostMemoryRegionSummary, error) {
	summary := NewHostMemoryRegionSummary()

	request := client.DefaultAPI.HostsGet(context.Background())
	hostColl, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: FindMemoryRegion_AllHosts")

			return nil, fmt.Errorf("failure: FindMemoryRegion_AllHosts: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FindMemoryRegion_AllHosts")

		// return nil, fmt.Errorf("failure: FindMemoryRegion_AllHosts: %s (%s)", status.Status.Message, err)
		return summary, nil //TODO: Error here instead?
	}

	for _, member := range hostColl.GetMembers() {
		hostId := ReadLastItemFromUri(member.GetUri())
		memoryRegion, err := FindMemoryRegionOnHost(client, hostId, memoryId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindMemoryRegionOnHost: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryRegion(hostId, memoryRegion)
	}

	return summary, nil
}

// Gather all available MemoryRegions from a specific Host.
func GetMemoryRegions_SingleHost(client *service.APIClient, hostId string) (*HostMemoryRegionSummary, error) {
	summary := NewHostMemoryRegionSummary()

	memoryRegions, err := GetAllMemoryRegionsForHost(client, hostId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllMemoryRegionsForHost: %s", err)
	}

	summary.AddMemoryRegionSlice(hostId, memoryRegions)

	return summary, nil
}

// Find a specific MemoryRegion from a specific connected Host.
func FindMemoryRegion_SingleHost(client *service.APIClient, hostId, memoryId string) (*HostMemoryRegionSummary, error) {
	summary := NewHostMemoryRegionSummary()

	memoryRegion, err := FindMemoryRegionOnHost(client, hostId, memoryId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindMemoryRegionOnHost: %s", err)
	}

	summary.AddMemoryRegion(hostId, memoryRegion)

	return summary, nil
}
