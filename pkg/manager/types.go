// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"fmt"
)

type SocketDetails struct {
	IpAddress string
	Port      uint16
}

func NewSocketDetails(ip string, port uint16) *SocketDetails {
	return &SocketDetails{
		IpAddress: ip,
		Port:      port,
	}
}

func (s SocketDetails) GetIpAndPort() (string, uint16) {
	return s.IpAddress, s.Port
}

func (s SocketDetails) String() string {
	return fmt.Sprintf("%s:%d", s.IpAddress, s.Port)
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
