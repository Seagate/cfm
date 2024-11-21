// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

func AddHost(client *service.APIClient, creds *service.Credentials) (*service.Host, error) {
	request := client.DefaultAPI.HostsPost(context.Background())
	request = request.Credentials(*creds)
	host, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: AddHost", "hostId", host.GetId())

	return host, nil
}

func DeleteHostById(client *service.APIClient, hostId string) (*service.Host, error) {
	request := client.DefaultAPI.HostsDeleteById(context.Background(), hostId)
	host, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: DeleteHostById", "hostId", host.GetId())

	return host, nil
}

func GetAllHosts(client *service.APIClient) (*[]*service.Host, error) {
	var hosts []*service.Host

	//Get existing hosts
	request := client.DefaultAPI.HostsGet(context.Background())
	collection, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(4).InfoS("success: HostsGet", "count", collection.GetMemberCount())

	for _, member := range collection.GetMembers() {
		id := ReadLastItemFromUri(member.GetUri())
		request2 := client.DefaultAPI.HostsGetById(context.Background(), id)
		host, response2, err := request2.Execute()
		if response2 != nil {
			defer response2.Body.Close() // Required by http lib implementation.
		}
		if err != nil {
			newErr := handleServiceError(response2, err)
			return nil, fmt.Errorf("execute failure(%T): %w", request2, newErr)
		}

		klog.V(4).InfoS("success: HostsGetById", "hostId", host.GetId())

		hosts = append(hosts, host)
	}

	klog.V(3).InfoS("success: GetAllHosts", "Host Count", len(hosts))

	return &hosts, nil
}

func RenameHostById(client *service.APIClient, hostId string, newHostId string) (*service.Host, error) {
	request := client.DefaultAPI.HostsUpdateById(context.Background(), hostId)
	request = request.NewHostId(newHostId)
	host, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: RenameHostById", "request", request)

	return host, nil
}

func ResyncHostById(client *service.APIClient, hostId string) (*service.Host, error) {
	request := client.DefaultAPI.HostsResyncById(context.Background(), hostId)
	host, response, err := request.Execute()
	if response != nil {
		defer response.Body.Close() // Required by http lib implementation.
	}
	if err != nil {
		newErr := handleServiceError(response, err)
		return nil, fmt.Errorf("execute failure(%T): %w", request, newErr)
	}

	klog.V(3).InfoS("success: ResyncHostById", "hostID", host.GetId())

	return host, nil
}
