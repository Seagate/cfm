// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceWrap

import (
	"context"
	"fmt"

	service "cfm/pkg/client"

	"k8s.io/klog/v2"
)

func AddHost(client *service.APIClient, creds *service.Credentials) (*service.Host, error) {
	newReqAddHost := client.DefaultAPI.HostsPost(context.Background())
	newReqAddHost = newReqAddHost.Credentials(*creds)
	addedHost, response, err := newReqAddHost.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", newReqAddHost)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: hosts post: %s", err)
	}

	klog.V(3).InfoS("HostsPost success", "hostId", addedHost.GetId())

	return addedHost, nil
}

func DeleteHostById(client *service.APIClient, hostId string) (*service.Host, error) {
	newReqDelHostById := client.DefaultAPI.HostsDeleteById(context.Background(), hostId)
	deletedHost, response, err := newReqDelHostById.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", newReqDelHostById)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: delete host by id failure")
	}

	klog.V(3).InfoS("HostsDeleteById success", "hostId", deletedHost.GetId())

	return deletedHost, nil
}

func GetAllHosts(client *service.APIClient) (*[]*service.Host, error) {
	var hosts []*service.Host

	//Get existing hosts
	newReqGetHosts := client.DefaultAPI.HostsGet(context.Background())
	collection, response, err := newReqGetHosts.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", newReqGetHosts)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: hosts get: %s", err)
	}

	klog.V(3).InfoS("HostsGet success", "count", collection.GetMemberCount())

	for _, member := range collection.GetMembers() {
		id := ReadLastItemFromUri(member.GetUri())
		newReqGetHostById := client.DefaultAPI.HostsGetById(context.Background(), id)
		host, response, err := newReqGetHostById.Execute()
		if err != nil {
			msg := fmt.Sprintf("%T: Execute FAILURE", newReqGetHostById)
			klog.ErrorS(err, msg, "response", response)
			return nil, fmt.Errorf("failure: hosts get by id: %s", err)
		}

		klog.V(3).InfoS("HostsGetById success", "hostId", host.GetId())

		hosts = append(hosts, host)
	}

	klog.V(3).InfoS("Discovered hosts", "count", len(hosts))

	return &hosts, nil
}

func ResyncHostById(client *service.APIClient, hostId string) (*service.Host, error) {
	request := client.DefaultAPI.HostsResyncById(context.Background(), hostId)
	host, response, err := request.Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", host)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: resync host by id failure")
	}

	klog.V(3).InfoS("BladesResyncById success", "hostID", host.GetId())

	return host, nil
}
