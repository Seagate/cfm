// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

type ApplToBladeToPortsMapType map[ApplianceBladeKey]*[]*service.PortInformation
type BladePortsSummary struct {
	ApplToBladeMap        ApplToBladeMapType
	ApplToBladeToPortsMap ApplToBladeToPortsMapType
}

// Of the 2 summary maps, this completely fills the 1st map and then fully initializes, but doesn't fill, the 2nd map
func NewBladePortsSummary(m ApplToBladeMapType) *BladePortsSummary {
	var summary BladePortsSummary

	summary.ApplToBladeMap = m
	summary.ApplToBladeToPortsMap = make(ApplToBladeToPortsMapType)

	for applId, blades := range m {
		for _, blade := range *blades {
			var ports []*service.PortInformation

			key := NewApplianceBladeKey(applId, blade.GetId())
			summary.ApplToBladeToPortsMap[*key] = &ports
		}
	}

	return &summary
}

// Add port to ports map.  *Blade map is assumed already filled
func (s *BladePortsSummary) AddPort(applId, bladeId string, port *service.PortInformation) {
	key := NewApplianceBladeKey(applId, bladeId)

	*s.ApplToBladeToPortsMap[*key] = append(*s.ApplToBladeToPortsMap[*key], port)
}

// Add multiple ports to ports map.  *Blade map is assumed already filled
func (s *BladePortsSummary) AddPortSlice(applId, bladeId string, ports *[]*service.PortInformation) {
	key := NewApplianceBladeKey(applId, bladeId)

	*s.ApplToBladeToPortsMap[*key] = append(*s.ApplToBladeToPortsMap[*key], *ports...)
}

func FindPortOnBlade(client *service.APIClient, applId, bladeId, portId string) (*service.PortInformation, error) {
	var port *service.PortInformation

	request := client.DefaultAPI.BladesGetPortById(context.Background(), applId, bladeId, portId)
	port, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: FindPortOnBlade")

			return nil, fmt.Errorf("failure: FindPortOnBlade: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FindPortOnBlade")

		return nil, fmt.Errorf("failure: FindPortOnBlade: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("success: FindPortOnBlade", "applId", applId, "bladeId", bladeId, "portId", port.GetId())

	return port, nil
}

func GetAllPortsForBlade(client *service.APIClient, applId, bladeId string) (*[]*service.PortInformation, error) {
	var ports []*service.PortInformation

	requestPorts := client.DefaultAPI.BladesGetPorts(context.Background(), applId, bladeId)
	portColl, response, err := requestPorts.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestPorts, err)
			klog.ErrorS(newErr, "failure: GetAllPortsForBlade")

			return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestPorts, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: GetAllPortsForBlade")

		// return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s (%s)", status.Status.Message, err)
		return &ports, nil //TODO: Error here instead?
	}

	klog.V(4).InfoS("success: BladesGetPorts", "applId", applId, "bladeId", bladeId, "portColl", portColl.GetMemberCount())

	for _, res := range portColl.GetMembers() {
		portId := ReadLastItemFromUri(res.GetUri())
		requestPortById := client.DefaultAPI.BladesGetPortById(context.Background(), applId, bladeId, portId)
		port, response, err := requestPortById.Execute()
		if err != nil {
			// Decode the JSON response into a struct
			var status service.StatusMessage
			if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
				newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestPortById, err)
				klog.ErrorS(newErr, "failure: GetAllPortsForBlade")

				return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s", newErr)
			}

			newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
				requestPortById, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
			klog.ErrorS(newErr, "failure: GetAllPortsForBlade")

			// return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s (%s)", status.Status.Message, err)
			continue //TODO: Error here instead?
		}

		klog.V(4).InfoS("success: BladesGetPortById", "applId", applId, "bladeId", bladeId, "portId", port.GetId())

		ports = append(ports, port)
	}

	return &ports, nil
}

// /////////////////
// Gather all available Ports from all available connected appliances and blades
func GetPorts_AllApplsAllBlades(client *service.APIClient) (*BladePortsSummary, error) {
	bladesSummary, err := GetAllBlades_AllAppls(client)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_AllAppls: %s", err)
	}

	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	for applId, blades := range summary.ApplToBladeMap {
		for _, blade := range *blades {
			ports, err := GetAllPortsForBlade(client, applId, blade.GetId())
			if err != nil {
				// return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s", err)
				continue //TODO: Error here instead?
			}

			summary.AddPortSlice(applId, blade.GetId(), ports)
		}
	}

	return summary, nil
}

// Find a specific Port (identified by port-id) from all connected blades and appliances
func FindPort_AllApplsAllBlades(client *service.APIClient, portId string) (*BladePortsSummary, error) {
	bladesSummary, err := GetAllBlades_AllAppls(client)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_AllAppls: %s", err)
	}

	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	for applId, blades := range summary.ApplToBladeMap {
		for _, blade := range *blades {
			port, err := FindPortOnBlade(client, applId, blade.GetId(), portId)
			if err != nil {
				// return nil, fmt.Errorf("failure: FindPortOnBlade: %s", err)
				continue //TODO: Error here instead?
			}

			summary.AddPort(applId, blade.GetId(), port)
		}
	}

	return summary, nil
}

// Gather all available Ports from a specific Blade (identified by blade-id) that is present on all connected appliances.
func GetPorts_AllApplsSingleBlade(client *service.APIClient, bladeId string) (*BladePortsSummary, error) {
	bladesSummary, err := FindBladeById_AllAppls(client, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: %s", err)
	}

	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	for applId := range summary.ApplToBladeMap {
		ports, err := GetAllPortsForBlade(client, applId, bladeId)
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddPortSlice(applId, bladeId, ports)
	}

	return summary, nil
}

// Find a specific Port (identified by port-id) from a specific connected Blade (identified by blade-id) that is present on all available Appliances.
func FindPort_AllApplsSingleBlade(client *service.APIClient, bladeId, portId string) (*BladePortsSummary, error) {
	bladesSummary, err := FindBladeById_AllAppls(client, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_AllAppls: %s", err)
	}

	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	for applId := range summary.ApplToBladeMap {
		port, err := FindPortOnBlade(client, applId, bladeId, portId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindPortOnBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddPort(applId, bladeId, port)
	}

	return summary, nil
}

// Gather all available Ports from all connected Blades from a specific Appliance (identified by appliance-id).
func GetPorts_SingleApplAllBlades(client *service.APIClient, applianceId string) (*BladePortsSummary, error) {
	blades, err := GetAllBlades_SingleAppl(client, applianceId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBladeSlice(applianceId, blades)
	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	for _, blade := range *summary.ApplToBladeMap[applianceId] {
		ports, err := GetAllPortsForBlade(client, applianceId, blade.GetId())
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddPortSlice(applianceId, blade.GetId(), ports)
	}

	return summary, nil
}

// Find a specific Port (identified by port-id) from all connected Blades from a specific Appliance (identified by appliance-id).
func FindPort_SingleApplAllBlades(client *service.APIClient, applianceId, portId string) (*BladePortsSummary, error) {
	blades, err := GetAllBlades_SingleAppl(client, applianceId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllBlades_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBladeSlice(applianceId, blades)
	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	for _, blade := range *summary.ApplToBladeMap[applianceId] {
		port, err := FindPortOnBlade(client, applianceId, blade.GetId(), portId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindPortOnBlade: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddPort(applianceId, blade.GetId(), port)
	}

	return summary, nil
}

// Gather all available Ports from a specific connected Blade (identified by blade-id) from a specific Appliance (identified by appliance-id).
func GetPorts_SingleApplSingleBlade(client *service.APIClient, applianceId, bladeId string) (*BladePortsSummary, error) {
	blade, err := FindBladeById_SingleAppl(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBlade(applianceId, blade)
	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	ports, err := GetAllPortsForBlade(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllPortsForBlade: %s", err)
	}

	summary.AddPortSlice(applianceId, blade.GetId(), ports)

	return summary, nil
}

// Find a specific Port (identified by port-id) from a specific connected Blade (identified by blade-id) from a specific Appliance (identified by appliance-id).
func FindPort_SingleApplSingleBlade(client *service.APIClient, applianceId, bladeId, portId string) (*BladePortsSummary, error) {
	blade, err := FindBladeById_SingleAppl(client, applianceId, bladeId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindBladeById_SingleAppl: %s", err)
	}

	bladesSummary := NewApplianceBladeSummary()
	bladesSummary.AddBlade(applianceId, blade)
	summary := NewBladePortsSummary(bladesSummary.ApplToBladeMap)

	port, err := FindPortOnBlade(client, applianceId, bladeId, portId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindPortOnBlade: %s", err)
	}

	summary.AddPort(applianceId, bladeId, port)

	return summary, nil
}

type HostToPortMapType map[string]*[]*service.PortInformation
type HostPortSummary struct {
	Ports HostToPortMapType
}

func NewHostPortSummary() *HostPortSummary {
	var summary HostPortSummary

	summary.Ports = make(HostToPortMapType)

	return &summary
}

func (s *HostPortSummary) AddHost(hostId string) {
	_, found := s.Ports[hostId]
	if !found {
		var ports []*service.PortInformation
		s.Ports[hostId] = &ports
	}
}

func (s *HostPortSummary) AddPort(hostId string, port *service.PortInformation) {
	s.AddHost(hostId)

	*s.Ports[hostId] = append(*s.Ports[hostId], port)
}

func (s *HostPortSummary) AddPortSlice(hostId string, ports *[]*service.PortInformation) {
	s.AddHost(hostId)

	*s.Ports[hostId] = append(*s.Ports[hostId], *ports...)
}

func (s *HostPortSummary) HostCount() int {

	return len(s.Ports)
}

// Find a specific Port by ID on a specific host
func FindPortById_SingleHost(client *service.APIClient, hostId, portId string) (*service.PortInformation, error) {

	request := client.DefaultAPI.HostsGetPortById(context.Background(), hostId, portId)
	port, response, err := request.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", request, err)
			klog.ErrorS(newErr, "failure: FindPortById_SingleHost")

			return nil, fmt.Errorf("failure: FindPortById_SingleHost: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			request, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FindPortById_SingleHost")

		return nil, fmt.Errorf("failure: FindPortById_SingleHost: %s (%s)", status.Status.Message, err)
	}

	klog.V(3).InfoS("success: HostsGetPortById", "hostId", hostId, "portId", port.GetId())

	return port, nil
}

// Find a specific Port by ID over 1 or more hosts
func FindPortById_AllHosts(client *service.APIClient, portId string) (*HostPortSummary, error) {
	summary := NewHostPortSummary()

	//Get all existing hosts
	requestHosts := client.DefaultAPI.HostsGet(context.Background())
	hostsColl, response, err := requestHosts.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestHosts, err)
			klog.ErrorS(newErr, "failure: FindPortById_AllHosts")

			return nil, fmt.Errorf("failure: FindPortById_AllHosts: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestHosts, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: FindPortById_AllHosts")

		return nil, fmt.Errorf("failure: FindPortById_AllHosts: %s (%s)", status.Status.Message, err)
	}

	klog.V(4).InfoS("success: HostsGet", "hostsColl", hostsColl.GetMemberCount())

	if hostsColl.GetMemberCount() == 0 {
		klog.V(3).InfoS("FindPortById_AllHosts: no hosts found")
		return nil, fmt.Errorf("failure: FindPortById_AllHosts: no hosts found")
	}

	//Scan collection members for target port id
	for _, hostMemeber := range hostsColl.GetMembers() {
		hostId := ReadLastItemFromUri(hostMemeber.GetUri())
		port, err := FindPortById_SingleHost(client, hostId, portId)
		if err != nil {
			continue
		}

		summary.AddPort(hostId, port)
	}

	return summary, nil
}

// Gather all available Ports from a specific Host
// Ports array is EMPTY if no ports found
func GetAllPorts_SingleHost(client *service.APIClient, hostId string) (*[]*service.PortInformation, error) {
	var ports []*service.PortInformation

	requestPorts := client.DefaultAPI.HostsGetPorts(context.Background(), hostId)
	portsColl, response, err := requestPorts.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestPorts, err)
			klog.ErrorS(newErr, "failure: GetAllPorts_SingleHost")

			return nil, fmt.Errorf("failure: GetAllPorts_SingleHost: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestPorts, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: GetAllPorts_SingleHost")

		return nil, fmt.Errorf("failure: GetAllPorts_SingleHost: %s (%s)", status.Status.Message, err)
	}

	klog.V(4).InfoS("success: PortsGet", "hostId", hostId, "portsColl", portsColl.GetMemberCount())

	for _, portMember := range portsColl.GetMembers() {
		portId := ReadLastItemFromUri(portMember.GetUri())
		port, err := FindPortById_SingleHost(client, hostId, portId)
		if err != nil {
			continue
		}

		ports = append(ports, port)
	}

	return &ports, nil
}

// Gather all available Ports from all available Hosts
// For each Host, Ports array is EMPTY if no Ports found
func GetAllPorts_AllHosts(client *service.APIClient) (*HostPortSummary, error) {
	summary := NewHostPortSummary()

	//Get all existing hosts
	requestHosts := client.DefaultAPI.HostsGet(context.Background())
	hostColl, response, err := requestHosts.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestHosts, err)
			klog.ErrorS(newErr, "failure: GetAllPorts_AllHosts")

			return nil, fmt.Errorf("failure: GetAllPorts_AllHosts: %s", newErr)
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestHosts, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.ErrorS(newErr, "failure: GetAllPorts_AllHosts")

		return nil, fmt.Errorf("failure: GetAllPorts_AllHosts: %s (%s)", status.Status.Message, err)
	}

	klog.V(4).InfoS("success: HostsGet", "hostColl", hostColl.GetMemberCount())

	//Scan collection members for target host id
	for _, host := range hostColl.GetMembers() {
		hostId := ReadLastItemFromUri(host.GetUri())
		ports, err := GetAllPorts_SingleHost(client, hostId)
		if err != nil {
			continue
			// return nil, err
		}

		summary.AddPortSlice(hostId, ports)
	}

	return summary, nil
}
