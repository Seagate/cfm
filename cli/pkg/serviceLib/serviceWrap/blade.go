// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

type ApplToBladeMapType map[string]*[]*service.Blade
type ApplianceBladeSummary struct {
	ApplToBladeMap ApplToBladeMapType
}

func NewApplianceBladeSummary() *ApplianceBladeSummary {
	var summary ApplianceBladeSummary

	summary.ApplToBladeMap = make(ApplToBladeMapType)

	return &summary
}

func (s *ApplianceBladeSummary) AddAppliance(applId string) {
	_, found := s.ApplToBladeMap[applId]
	if !found {
		var blades []*service.Blade
		s.ApplToBladeMap[applId] = &blades
	}
}

func (s *ApplianceBladeSummary) AddBlade(applId string, blade *service.Blade) {
	s.AddAppliance(applId)

	*s.ApplToBladeMap[applId] = append(*s.ApplToBladeMap[applId], blade)
}

func (s *ApplianceBladeSummary) AddBladeSlice(applId string, blades *[]*service.Blade) {
	s.AddAppliance(applId)

	*s.ApplToBladeMap[applId] = append(*s.ApplToBladeMap[applId], *blades...)
}

func (s *ApplianceBladeSummary) ApplianceCount() int {

	return len(s.ApplToBladeMap)
}

func AddBlade(client *service.APIClient, applianceId string, bladeCreds *service.Credentials) (*service.Blade, error) {
	addRequest := client.DefaultAPI.BladesPost(context.Background(), applianceId)
	addRequest = addRequest.Credentials(*bladeCreds)
	addedBlade, response, err := addRequest.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", addRequest)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: blades post: %s", err)
	}

	klog.V(3).InfoS("BladesPost success", "bladeId", addedBlade.GetId(), "applianceId", applianceId)

	return addedBlade, nil
}

func DeleteBladeById(client *service.APIClient, applId, bladeId string) (*service.Blade, error) {
	deleteRequest := client.DefaultAPI.BladesDeleteById(context.Background(), applId, bladeId)
	deletedBlade, response, err := deleteRequest.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", deleteRequest)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: delete blade by id failure")
	}

	klog.V(3).InfoS("BladesDeleteById success", "ApplianceID", applId, "BladeID", deletedBlade.GetId())

	return deletedBlade, nil
}

// Find a specific Blade by ID on a specific Appliance
func FindBladeById_SingleAppl(client *service.APIClient, applId, bladeId string) (*service.Blade, error) {

	requestBlade := client.DefaultAPI.BladesGetById(context.Background(), applId, bladeId)
	//TODO: What does this api do when the blade is empty
	blade, response, err := requestBlade.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", requestBlade)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: get appliance blade by id: %s", err)
	}

	klog.V(3).InfoS("BladesGetById success", "applId", applId, "bladeId", blade.GetId())

	return blade, nil
}

// Find a specific Blade by ID over 1 or more appliances
func FindBladeById_AllAppls(client *service.APIClient, bladeId string) (*ApplianceBladeSummary, error) {
	summary := NewApplianceBladeSummary()

	//Get all existing appliances
	requestAppliances := client.DefaultAPI.AppliancesGet(context.Background())
	applColl, response, err := requestAppliances.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", requestAppliances)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: get appliances: %s", err)
	}

	klog.V(3).InfoS("AppliancesGet success", "applCount", applColl.GetMemberCount())

	if applColl.GetMemberCount() == 0 {
		klog.V(3).InfoS("FindBladeById_AllAppls: no appliances found")
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: no appliances found")
	}

	//Scan collection members for target blade id
	for _, appl := range applColl.GetMembers() {
		applId := ReadLastItemFromUri(appl.GetUri())
		blade, err := FindBladeById_SingleAppl(client, applId, bladeId)
		if err != nil {
			continue
		}

		summary.AddBlade(applId, blade)
	}

	return summary, nil
}

// Gather all available Blades from a specific Appliance
// Blade array is EMPTY if no blades found
func GetAllBlades_SingleAppl(client *service.APIClient, applId string) (*[]*service.Blade, error) {
	var blades []*service.Blade

	requestBlades := client.DefaultAPI.BladesGet(context.Background(), applId)
	bladeColl, response, err := requestBlades.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", requestBlades)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: get appliance blades: %s", err)
	}

	klog.V(3).InfoS("BladesGet success", "applId", applId, "bladeColl", bladeColl.GetMemberCount())

	for _, bladeMember := range bladeColl.GetMembers() {
		bladeId := ReadLastItemFromUri(bladeMember.GetUri())
		blade, err := FindBladeById_SingleAppl(client, applId, bladeId)
		if err != nil {
			continue
		}

		blades = append(blades, blade)
	}

	return &blades, nil
}

// Gather all available Blades from all available Appliances
// For each Appliance, Blade array is EMPTY if no Blades found
func GetAllBlades_AllAppls(client *service.APIClient) (*ApplianceBladeSummary, error) {
	summary := NewApplianceBladeSummary()

	//Get all existing appliances
	requestAppliances := client.DefaultAPI.AppliancesGet(context.Background())
	applColl, response, err := requestAppliances.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", requestAppliances)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: get appliances: %s", err)
	}

	klog.V(3).InfoS("AppliancesGet success", "applCount", applColl.GetMemberCount())

	//Scan collection members for target appliance id
	for _, appl := range applColl.GetMembers() {
		applId := ReadLastItemFromUri(appl.GetUri())
		blades, err := GetAllBlades_SingleAppl(client, applId)
		if err != nil {
			continue
			// return nil, err
		}

		summary.AddBladeSlice(applId, blades)
	}

	return summary, nil
}

func ResyncBladeById(client *service.APIClient, applId, bladeId string) (*service.Blade, error) {
	request := client.DefaultAPI.BladesResyncById(context.Background(), applId, bladeId)
	blade, response, err := request.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", blade)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: resync blade by id failure")
	}

	klog.V(3).InfoS("BladesResyncById success", "ApplianceID", applId, "BladeID", blade.GetId())

	return blade, nil
}
