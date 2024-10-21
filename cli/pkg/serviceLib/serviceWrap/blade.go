// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
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
	request := client.DefaultAPI.BladesPost(context.Background(), applianceId)
	request = request.Credentials(*bladeCreds)
	blade, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: AddBlade", "bladeId", blade.GetId(), "applianceId", applianceId)

	return blade, nil
}

func DeleteBladeById(client *service.APIClient, applId, bladeId string) (*service.Blade, error) {
	request := client.DefaultAPI.BladesDeleteById(context.Background(), applId, bladeId)
	blade, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: DeleteBladeById", "applianceId", applId, "bladeID", blade.GetId())

	return blade, nil
}

// Find a specific Blade by ID on a specific Appliance
func FindBladeById_SingleAppl(client *service.APIClient, applId, bladeId string) (*service.Blade, error) {

	request := client.DefaultAPI.BladesGetById(context.Background(), applId, bladeId)
	//TODO: What does this api do when the blade is empty
	blade, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: FindBladeById_SingleAppl", "applianceId", applId, "bladeId", blade.GetId())

	return blade, nil
}

// Find a specific Blade by ID over 1 or more appliances
func FindBladeById_AllAppls(client *service.APIClient, bladeId string) (*ApplianceBladeSummary, error) {
	summary := NewApplianceBladeSummary()

	//Get all existing appliances
	request := client.DefaultAPI.AppliancesGet(context.Background())
	applColl, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(4).InfoS("success: AppliancesGet", "applCount", applColl.GetMemberCount())

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

	request := client.DefaultAPI.BladesGet(context.Background(), applId)
	bladeColl, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(4).InfoS("success: BladesGet", "applianceId", applId, "bladeCount", bladeColl.GetMemberCount())

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
	request := client.DefaultAPI.AppliancesGet(context.Background())
	applColl, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(4).InfoS("success: AppliancesGet", "applCount", applColl.GetMemberCount())

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

func RenameBladeById(client *service.APIClient, applId string, bladeId string, newBladeId string) (*service.Blade, error) {
	request := client.DefaultAPI.BladesUpdateById(context.Background(), applId, bladeId)
	request = request.NewBladeId(newBladeId)
	blade, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: RenameBladeById", "request", request)

	return blade, nil
}

func ResyncBladeById(client *service.APIClient, applId, bladeId string) (*service.Blade, error) {
	request := client.DefaultAPI.BladesResyncById(context.Background(), applId, bladeId)
	blade, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(3).InfoS("success: ResyncBladeById", "applianceId", applId, "bladeID", blade.GetId())

	return blade, nil
}
