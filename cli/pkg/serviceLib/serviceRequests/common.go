/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package serviceRequests

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceWrap"
	"fmt"
	"strconv"
	"strings"

	service "cfm/pkg/client"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

const ERR_MSG_MISSING_FLAG string = "Missing cli flag"
const ERR_MSG_INVALID_FLAG_SETTING string = "Invalid flag setting"
const ERR_MSG_INVALID_SIZE_FORMAT string = "Invalid size format"
const ERR_MSG_INVALID_SIZE_UNIT_VALUE string = "Invalid size unit value"

const HELP_MSG_SIZE_INPUT_FORMAT string = "size format: a string containing a non-zero integer followed by single letter unit('G','g')"

const MIN_MEMORY_SIZE_INPUT_LENGTH int = 2

type Id struct {
	id string
}

func NewId(cmd *cobra.Command, componentName string) *Id {
	flagName := fmt.Sprintf("%s-%s", componentName, flags.ID)

	id_, err := cmd.Flags().GetString(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}

	return &Id{
		id: id_,
	}
}

func (i *Id) GetId() string {
	return i.id
}

func (i *Id) HasId() bool {
	var dflt string

	return i.id != dflt
}

func (i *Id) SetId(id string) {
	i.id = id
}

// TODO: How secure\encrypt\hide this information??
type DeviceCredentials struct {
	username string
	password string
}

func NewDeviceCredentials(cmd *cobra.Command, componentName string) *DeviceCredentials {
	var flagName string

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.USERNAME)
	} else {
		flagName = flags.USERNAME
	}
	username_, err := cmd.Flags().GetString(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.PASSWORD)
	} else {
		flagName = flags.PASSWORD
	}
	password_, _ := cmd.Flags().GetString(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}

	return &DeviceCredentials{
		username: username_,
		password: password_,
	}
}

func (d *DeviceCredentials) GetUsername() string {
	return d.username
}

func (d *DeviceCredentials) GetPassword() string {
	return d.password
}

type TcpInfo struct {
	ip       string
	port     uint16
	insecure bool
	protocol string
}

func NewTcpInfo(cmd *cobra.Command, componentName string) *TcpInfo {
	var flagName string

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.NET_IP)
	} else {
		flagName = flags.NET_IP
	}
	ip_, err := cmd.Flags().GetString(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}
	err = serviceWrap.ValidateIPAddress(ip_)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_INVALID_FLAG_SETTING, "flag", flagName, "value", ip_)
		cobra.CheckErr(err)
	}

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.NET_PORT)
	} else {
		flagName = flags.NET_PORT
	}
	port_, err := cmd.Flags().GetUint16(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}
	err = serviceWrap.ValidatePort(port_)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_INVALID_FLAG_SETTING, "flag", flagName, "value", port_)
		cobra.CheckErr(err)
	}

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.INSECURE)
	} else {
		flagName = flags.INSECURE
	}
	insecure_, err := cmd.Flags().GetBool(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.PROTOCOL)
	} else {
		flagName = flags.PROTOCOL
	}
	protocol_, err := cmd.Flags().GetString(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}

	return &TcpInfo{
		ip:       ip_,
		port:     port_,
		insecure: insecure_,
		protocol: protocol_,
	}
}

func (t *TcpInfo) GetIp() string {
	return t.ip
}

func (t *TcpInfo) GetPort() uint16 {
	return t.port
}

func (t *TcpInfo) GetInsecure() bool {
	return t.insecure
}

func (t *TcpInfo) GetProtocol() string {
	return t.protocol
}

func (t *TcpInfo) GetIpPort() string {
	return fmt.Sprintf("%s:%d", t.ip, t.port)
}

type Size struct {
	sizeGiB int32
	units   string
}

func NewSize(cmd *cobra.Command, componentName string) *Size {
	var flagName string

	if len(componentName) > 0 {
		flagName = fmt.Sprintf("%s-%s", componentName, flags.SIZE)
	} else {
		flagName = flags.SIZE
	}
	s, err := cmd.Flags().GetString(flagName)
	if err != nil {
		klog.ErrorS(err, ERR_MSG_MISSING_FLAG, "flag", flagName)
		cobra.CheckErr(err)
	}

	size, units := ParseSizeFlagString(&s)

	if size < 1 {
		err := fmt.Errorf("failure: " + ERR_MSG_INVALID_FLAG_SETTING)
		klog.ErrorS(err, "size must be >= 1", "found size", size)
		cobra.CheckErr(err)
	}

	return &Size{
		sizeGiB: size,
		units:   *units,
	}
}

func (s *Size) GetSizeGiB() int32 {
	return s.sizeGiB
}

func (s *Size) GetUnits() *string {
	if s == nil {
		var ret string
		return &ret
	}
	return &s.units
}

// s: input string from size flag.  Expected format: 8G, 32g
func ParseSizeFlagString(s *string) (int32, *string) {
	var units string

	runes := []rune(*s)
	length := len(runes)
	if length < MIN_MEMORY_SIZE_INPUT_LENGTH {
		err := fmt.Errorf("failure: " + ERR_MSG_INVALID_SIZE_FORMAT)
		errMsg := fmt.Sprintf("minimum size param length >= %d", MIN_MEMORY_SIZE_INPUT_LENGTH)
		klog.ErrorS(err, errMsg)
		cobra.CheckErr(err)
	}

	firstRune := runes[0]
	//Look for number
	size, err := strconv.ParseInt(string(firstRune), 10, 32)
	if err != nil {
		//if conversion failed, first is non-numeric.  fail
		err := fmt.Errorf("failure: " + ERR_MSG_INVALID_SIZE_FORMAT)
		klog.ErrorS(err, HELP_MSG_SIZE_INPUT_FORMAT)
		cobra.CheckErr(err)
	}

	//Look for letter
	lastRune := runes[length-1]
	_, err = strconv.ParseInt(string(lastRune), 10, 32)
	if err == nil {
		//if conversion passed, last is numeric.  fail
		err := fmt.Errorf("failure: " + ERR_MSG_INVALID_SIZE_FORMAT)
		klog.ErrorS(err, HELP_MSG_SIZE_INPUT_FORMAT)
		cobra.CheckErr(err)
	}

	if length > MIN_MEMORY_SIZE_INPUT_LENGTH {
		numericRunes := runes[:length-1]
		size, err = strconv.ParseInt(string(numericRunes), 10, 32)
		if err != nil {
			//if conversion failed, all runes are NOT numeric.  fail
			err := fmt.Errorf("failure: " + ERR_MSG_INVALID_SIZE_FORMAT)
			klog.ErrorS(err, HELP_MSG_SIZE_INPUT_FORMAT)
			cobra.CheckErr(err)
		}
	}

	unitChar := string(lastRune)
	switch unitChar {
	case "G", "g":
		units = "GiB"
	default:
		err := fmt.Errorf("failure: " + ERR_MSG_INVALID_SIZE_UNIT_VALUE)
		klog.ErrorS(err, HELP_MSG_SIZE_INPUT_FORMAT, "found char", unitChar)
		cobra.CheckErr(err)

	}

	return int32(size), &units
}

// CLI REQUESTS
type ServiceRequestListExternalDevices struct {
	ServiceTcp *TcpInfo
}

func NewServiceRequestListExternalDevices(cmd *cobra.Command) *ServiceRequestListExternalDevices {
	return &ServiceRequestListExternalDevices{
		ServiceTcp: NewTcpInfo(cmd, flags.SERVICE),
	}
}

///////////////////////////////
// Output functions

// Output for add and delete of appliance
func OutputResultsAddDeleteAppliance(a *service.Appliance, action string) {
	if a == nil {
		fmt.Printf("\nAppliance %s Status: FAILED\n\n", strings.ToUpper(action))
		return
	}

	fmt.Printf("\nAppliance %s\n", strings.ToUpper(action))
	fmt.Printf("Appliance ID: %s\n\n", a.GetId())
}

// Output for add and delete of appliance blade
func OutputResultsAddDeleteBlade(b *service.Blade, action string) {
	if b == nil {
		fmt.Printf("\nBlade %s Status: FAILED\n\n", strings.ToUpper(action))
		return
	}

	fmt.Printf("\nBlade %s\n", strings.ToUpper(action))
	fmt.Printf("Blade ID: %s\n\n", b.GetId())
}

// Output for add and delete of host
func OutputResultsAddDeleteHost(h *service.Host, action string) {
	if h == nil {
		fmt.Printf("\nHost %s Status: FAILED\n\n", strings.ToUpper(action))
		return
	}

	fmt.Printf("\nHost %s\n", strings.ToUpper(action))
	fmt.Printf("Host ID: %s\n\n", h.GetId())
}
