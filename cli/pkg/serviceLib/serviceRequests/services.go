// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package serviceRequests

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceWrap"
	"context"
	"fmt"

	service "cfm/pkg/client"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

type ServiceRequestListServiceInfo struct {
	ServiceTcp *TcpInfo
}

func NewServiceRequestListServiceInfo(cmd *cobra.Command) *ServiceRequestListServiceInfo {
	return &ServiceRequestListServiceInfo{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
	}
}

func (s *ServiceRequestListServiceInfo) Execute() (*service.ServiceInformation, error) {
	klog.V(4).InfoS(fmt.Sprintf("%T", *s), "ServiceTcp", fmt.Sprintf("%+v", *s.ServiceTcp))

	serviceClient := serviceWrap.GetServiceClient(s.ServiceTcp.GetIp(), s.ServiceTcp.GetPort())

	serviceInfo, response, err := serviceClient.DefaultAPI.CfmV1Get(context.Background()).Execute()
	if err != nil {
		msg := fmt.Sprintf("%T: Execute FAILURE", serviceInfo)
		klog.ErrorS(err, msg, "response", response)
		return nil, fmt.Errorf("failure: get service info")
	}

	return serviceInfo, nil
}

func (s *ServiceRequestListServiceInfo) OutputResults(info *service.ServiceInformation) {

	fmt.Printf("\nService Information @ %s\n", s.ServiceTcp.GetIpPort())
	fmt.Printf("Version: %s\n", info.GetVersion())

	fmt.Printf("\n")
}
