// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

type HostToMemoryDevicesMapType map[string]*[]*service.MemoryDeviceInformation
type HostMemoryDeviceSummary struct {
	MemoryDevices HostToMemoryDevicesMapType
}

func NewHostMemoryDeviceSummary() *HostMemoryDeviceSummary {
	var summary HostMemoryDeviceSummary

	summary.MemoryDevices = make(HostToMemoryDevicesMapType)

	return &summary
}

func (s *HostMemoryDeviceSummary) AddHost(hostId string) {
	_, found := s.MemoryDevices[hostId]
	if !found {
		var memory []*service.MemoryDeviceInformation
		s.MemoryDevices[hostId] = &memory
	}
}

// Add memoryDevice to memoryDevices map.
func (s *HostMemoryDeviceSummary) AddMemoryDevice(hostId string, memoryDevice *service.MemoryDeviceInformation) {
	s.AddHost(hostId)

	*s.MemoryDevices[hostId] = append(*s.MemoryDevices[hostId], memoryDevice)
}

// Add multiple memoryDevices to memoryDevices map.
func (s *HostMemoryDeviceSummary) AddMemoryDeviceSlice(hostId string, memoryDevices *[]*service.MemoryDeviceInformation) {
	s.AddHost(hostId)

	*s.MemoryDevices[hostId] = append(*s.MemoryDevices[hostId], *memoryDevices...)
}

func (s *HostMemoryDeviceSummary) HostCount() int {

	return len(s.MemoryDevices)
}

func FindMemoryDeviceOnHost(client *service.APIClient, hostId, memoryDeviceId string) (*service.MemoryDeviceInformation, error) {
	var memoryDevice *service.MemoryDeviceInformation

	request := client.DefaultAPI.HostsGetMemoryDeviceById(context.Background(), hostId, memoryDeviceId)
	memoryDevice, response, err := request.Execute()
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

	klog.V(3).InfoS("success: FindMemoryDeviceOnHost", "hostId", hostId, "memoryDeviceId", memoryDevice.GetId())

	return memoryDevice, nil
}

func GetAllMemoryDevicesForHost(client *service.APIClient, hostId string) (*[]*service.MemoryDeviceInformation, error) {
	var memoryDevices []*service.MemoryDeviceInformation

	request := client.DefaultAPI.HostsGetMemoryDevices(context.Background(), hostId)
	memoryDeviceColl, response, err := request.Execute()
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
		// return nil, newErr
		return &memoryDevices, nil //TODO: Error here instead?
	}

	klog.V(4).InfoS("success: HostsGetMemoryDevices", "hostId", hostId, "memoryDeviceColl", memoryDeviceColl.GetMemberCount())

	for _, member := range memoryDeviceColl.GetMembers() {
		memoryDeviceId := ReadLastItemFromUri(member.GetUri())
		request := client.DefaultAPI.HostsGetMemoryDeviceById(context.Background(), hostId, memoryDeviceId)
		memoryDevice, response, err := request.Execute()
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
			// return nil, newErr
			continue //TODO: Error here instead?
		}

		klog.V(4).InfoS("success: HostsGetMemoryById", "hostId", hostId, "memoryDeviceId", memoryDevice.GetId())

		memoryDevices = append(memoryDevices, memoryDevice)
	}

	return &memoryDevices, nil
}

// Gather all available MemoryDevices from all avaiable Hosts.
func GetMemoryDevices_AllHosts(client *service.APIClient) (*HostMemoryDeviceSummary, error) {
	summary := NewHostMemoryDeviceSummary()

	request := client.DefaultAPI.HostsGet(context.Background())
	hostColl, response, err := request.Execute()
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
		// return nil, newErr
		return summary, nil //TODO: Error here instead?
	}

	for _, member := range hostColl.GetMembers() {
		hostId := ReadLastItemFromUri(member.GetUri())
		memoryDevices, err := GetAllMemoryDevicesForHost(client, hostId)
		if err != nil {
			// return nil, fmt.Errorf("failure: GetAllMemoryDevicesForHost: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryDeviceSlice(hostId, memoryDevices)
	}

	return summary, nil
}

// Find a specific MemoryDevice from all available Hosts.
func FindMemoryDevice_AllHosts(client *service.APIClient, memoryDeviceId string) (*HostMemoryDeviceSummary, error) {
	summary := NewHostMemoryDeviceSummary()

	request := client.DefaultAPI.HostsGet(context.Background())
	hostColl, response, err := request.Execute()
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
		// return nil, newErr
		return summary, nil //TODO: Error here instead?
	}

	for _, member := range hostColl.GetMembers() {
		hostId := ReadLastItemFromUri(member.GetUri())
		memoryDevice, err := FindMemoryDeviceOnHost(client, hostId, memoryDeviceId)
		if err != nil {
			// return nil, fmt.Errorf("failure: FindMemoryDeviceOnHost: %s", err)
			continue //TODO: Error here instead?
		}

		summary.AddMemoryDevice(hostId, memoryDevice)
	}

	return summary, nil
}

// Gather all available MemoryDevices from a specific Host.
func GetMemoryDevices_SingleHost(client *service.APIClient, hostId string) (*HostMemoryDeviceSummary, error) {
	summary := NewHostMemoryDeviceSummary()

	memoryDevices, err := GetAllMemoryDevicesForHost(client, hostId)
	if err != nil {
		return nil, fmt.Errorf("failure: GetAllMemoryDevicesForHost: %s", err)
	}

	summary.AddMemoryDeviceSlice(hostId, memoryDevices)

	return summary, nil
}

// Find a specific MemoryDevice from a specific connected Host.
func FindMemoryDevice_SingleHost(client *service.APIClient, hostId, memoryDeviceId string) (*HostMemoryDeviceSummary, error) {
	summary := NewHostMemoryDeviceSummary()

	memoryDevice, err := FindMemoryDeviceOnHost(client, hostId, memoryDeviceId)
	if err != nil {
		return nil, fmt.Errorf("failure: FindMemoryDeviceOnHost: %s", err)
	}

	summary.AddMemoryDevice(hostId, memoryDevice)

	return summary, nil
}
