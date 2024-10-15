// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"encoding/json"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

func AddHost(client *service.APIClient, creds *service.Credentials) (*service.Host, error) {
	request := client.DefaultAPI.HostsPost(context.Background())
	request = request.Credentials(*creds)
	host, response, err := request.Execute()
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

	klog.V(3).InfoS("success: AddHost", "hostId", host.GetId())

	return host, nil
}

func DeleteHostById(client *service.APIClient, hostId string) (*service.Host, error) {
	request := client.DefaultAPI.HostsDeleteById(context.Background(), hostId)
	host, response, err := request.Execute()
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

	klog.V(3).InfoS("success: DeleteHostById", "hostId", host.GetId())

	return host, nil
}

func GetAllHosts(client *service.APIClient) (*[]*service.Host, error) {
	var hosts []*service.Host

	//Get existing hosts
	requestGetHosts := client.DefaultAPI.HostsGet(context.Background())
	collection, response, err := requestGetHosts.Execute()
	if err != nil {
		// Decode the JSON response into a struct
		var status service.StatusMessage
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestGetHosts, err)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
			requestGetHosts, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
		klog.V(4).Info(newErr)
		return nil, newErr
	}

	klog.V(4).InfoS("success: HostsGet", "count", collection.GetMemberCount())

	for _, member := range collection.GetMembers() {
		id := ReadLastItemFromUri(member.GetUri())
		requestGetHostById := client.DefaultAPI.HostsGetById(context.Background(), id)
		host, response, err := requestGetHostById.Execute()
		if err != nil {
			// Decode the JSON response into a struct
			var status service.StatusMessage
			if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
				newErr := fmt.Errorf("failure: Execute(%T): err(%s), error decoding response JSON", requestGetHostById, err)
				klog.V(4).Info(newErr)
				return nil, newErr
			}

			newErr := fmt.Errorf("failure: Execute(%T): err(%s), uri(%s), details(%s), code(%d), message(%s)",
				requestGetHostById, err, status.Uri, status.Details, status.Status.Code, status.Status.Message)
			klog.V(4).Info(newErr)
			return nil, newErr
		}

		klog.V(4).InfoS("success: HostsGetById", "hostId", host.GetId())

		hosts = append(hosts, host)
	}

	klog.V(3).InfoS("success: GetAllHosts", "Host Count", len(hosts))

	return &hosts, nil
}

func ResyncHostById(client *service.APIClient, hostId string) (*service.Host, error) {
	request := client.DefaultAPI.HostsResyncById(context.Background(), hostId)
	host, response, err := request.Execute()
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

	klog.V(3).InfoS("success: ResyncHostById", "hostID", host.GetId())

	return host, nil
}
