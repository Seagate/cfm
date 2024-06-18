// Copyright (c) 2022 Seagate Technology LLC and/or its Affiliates

package cxl_host

import (
	"cfm/pkg/accounts"
	"cfm/pkg/redfishapi"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"syscall"

	"github.com/pbnjay/memory"
	"golang.org/x/exp/slices"
	"k8s.io/klog/v2"
)

// CxlHostApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type CxlHostApiService struct {
}

// NewCxlHostApiService creates a default api service
func NewCxlHostApiService() redfishapi.DefaultAPIServicer {
	return &CxlHostApiService{}
}

var enabled = true
var resourcehealthOk = redfishapi.RESOURCEHEALTH_OK
var resourcestateEnabled = redfishapi.RESOURCESTATE_ENABLED

// RedfishV1AccountServiceAccountsGet -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsGet(ctx context.Context) (redfishapi.ImplResponse, error) {
	logger := klog.FromContext(ctx)

	collection := redfishapi.ManagerAccountCollectionManagerAccountCollection{
		OdataContext: "/redfish/v1/$metadata#ManagerAccountCollection.ManagerAccountCollection",
		OdataId:      "/redfish/v1/AccountService/Accounts",
		OdataType:    "#ManagerAccountCollection.ManagerAccountCollection",
		Description:  "Accounts Collection",
		Name:         "Accounts Collection",
		Oem:          nil,
	}

	// Iterate over all account usernames adding each to the collection
	usernames := accounts.GetAccountUsernames()
	logger.V(2).Info("AccountServiceAccountsGet", "usernames", usernames)
	collection.MembersodataCount = int64(len(usernames))
	for _, username := range usernames {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/AccountService/Accounts/%s", username)})
	}

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1AccountServiceAccountsManagerAccountIdDelete -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsManagerAccountIdDelete(ctx context.Context, managerAccountId string) (redfishapi.ImplResponse, error) {

	account := accounts.GetAccount(managerAccountId)
	if account == nil {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Account id (%s) does not exist", managerAccountId)
	}

	err := accounts.AccountsHandler().DeleteAccount(managerAccountId)
	if err != nil {
		return redfishapi.Response(http.StatusBadRequest, nil), err
	}

	resource := fillAccountResource(managerAccountId, account)
	return redfishapi.Response(http.StatusOK, resource), nil
}

func fillAccountResource(managerAccountId string, account *accounts.AccountSingle) redfishapi.ManagerAccountV1120ManagerAccount {

	resource := redfishapi.ManagerAccountV1120ManagerAccount{
		OdataContext: "/redfish/v1/$metadata#ManagerAccount.ManagerAccount",
		OdataId:      "/redfish/v1/AccountService/Accounts/" + managerAccountId,
		OdataType:    ManagerAccountVersion,
		AccountTypes: []redfishapi.ManagerAccountAccountTypes{redfishapi.MANAGERACCOUNTACCOUNTTYPES_REDFISH},
		Description:  "User Account",
		Name:         "UserAccount",
		Password:     nil,
	}

	if account != nil {
		resource.Enabled = account.Enabled
		resource.Id = account.Username
		resource.Links = redfishapi.ManagerAccountV1120Links{
			Role: redfishapi.OdataV4IdRef{
				OdataId: "/redfish/v1/AccountService/Roles/" + account.Role,
			},
		}
		resource.Locked = account.Locked
		resource.PasswordChangeRequired = &account.PasswordChangedRequired
		resource.RoleId = account.Role
		resource.UserName = account.Username
	}

	return resource
}

// RedfishV1AccountServiceAccountsManagerAccountIdGet -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsManagerAccountIdGet(ctx context.Context, managerAccountId string) (redfishapi.ImplResponse, error) {

	account := accounts.GetAccount(managerAccountId)
	if account == nil {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Account id (%s) does not exist", managerAccountId)
	}

	resource := fillAccountResource(managerAccountId, account)

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1AccountServiceAccountsManagerAccountIdPatch -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsManagerAccountIdPatch(ctx context.Context, managerAccountId string, managerAccountV1120ManagerAccount redfishapi.ManagerAccountV1120ManagerAccount) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceAccountsManagerAccountIdPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// Retrieve existing account, or nil indicating it does not exist
	account := accounts.GetAccount(managerAccountId)
	if account == nil {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Account id (%s) does not exist", managerAccountId)
	}

	// Allow the client to change the password for this user.
	_, err := accounts.AccountsHandler().UpdateAccount(account.Username, *managerAccountV1120ManagerAccount.Password, "")
	if err != nil {
		return redfishapi.Response(http.StatusBadRequest, nil), err
	}

	resource := fillAccountResource(managerAccountId, account)
	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1AccountServiceAccountsPost -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsPost(ctx context.Context, managerAccountV1120ManagerAccount redfishapi.ManagerAccountV1120ManagerAccount) (redfishapi.ImplResponse, error) {

	// Create a new user account
	account, err := accounts.AccountsHandler().AddAccount(managerAccountV1120ManagerAccount.UserName, *managerAccountV1120ManagerAccount.Password, managerAccountV1120ManagerAccount.RoleId)
	if err != nil {
		return redfishapi.Response(http.StatusBadRequest, nil), err
	}

	resource := fillAccountResource(account.Username, account)
	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1AccountServiceGet -
func (s *CxlHostApiService) RedfishV1AccountServiceGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	resource := redfishapi.AccountServiceV1150AccountService{
		OdataContext: "/redfish/v1/$metadata#AccountService.AccountService",
		OdataId:      "/redfish/v1/AccountService",
		OdataType:    AccountServiceVersion,
		Accounts: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/AccountService/Accounts",
		},
		Description:       "Account Service",
		Id:                "AccountService",
		MinPasswordLength: accounts.MinPasswordLength,
		Name:              "Account Service",
		Roles: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/AccountService/Roles",
		},
		ServiceEnabled: &enabled,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1AccountServiceRolesGet -
func (s *CxlHostApiService) RedfishV1AccountServiceRolesGet(ctx context.Context) (redfishapi.ImplResponse, error) {
	logger := klog.FromContext(ctx)

	collection := redfishapi.RoleCollectionRoleCollection{
		OdataContext: "/redfish/v1/$metadata#RoleCollection.RoleCollection",
		OdataId:      "/redfish/v1/AccountService/Roles",
		OdataType:    "#RoleCollection.RoleCollection",
		Description:  "Roles Collection",
		Name:         "Roles Collection",
		Oem:          nil,
	}

	// Iterate over all roles adding each to the collection
	roles := accounts.GetRoles()
	logger.V(2).Info("AccountServiceRolesGet", "roles", roles)
	collection.MembersodataCount = int64(len(roles))
	for _, role := range roles {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/AccountService/Roles/%s", role)})
	}

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1AccountServiceRolesRoleIdGet -
func (s *CxlHostApiService) RedfishV1AccountServiceRolesRoleIdGet(ctx context.Context, roleId string) (redfishapi.ImplResponse, error) {

	if !slices.Contains(accounts.RoleTypes, roleId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Role id (%s) is not valid", roleId)
	}

	resource := redfishapi.RoleV131Role{
		OdataContext: "/redfish/v1/$metadata#Role.Role",
		OdataId:      "/redfish/v1/AccountService/Roles/" + roleId,
		OdataType:    RoleVersion,
		AssignedPrivileges: []redfishapi.PrivilegesPrivilegeType{
			redfishapi.PRIVILEGESPRIVILEGETYPE_LOGIN,
			redfishapi.PRIVILEGESPRIVILEGETYPE_CONFIGURE_SELF,
		},
		Description:  "Role Account",
		Name:         "RoleAccount",
		IsPredefined: true,
		RoleId:       roleId,
	}

	switch roleId {
	case accounts.RoleAdministrator:
		resource.AssignedPrivileges = append(resource.AssignedPrivileges, redfishapi.PRIVILEGESPRIVILEGETYPE_CONFIGURE_MANAGER)
		resource.AssignedPrivileges = append(resource.AssignedPrivileges, redfishapi.PRIVILEGESPRIVILEGETYPE_CONFIGURE_USERS)
		resource.AssignedPrivileges = append(resource.AssignedPrivileges, redfishapi.PRIVILEGESPRIVILEGETYPE_CONFIGURE_COMPONENTS)

	case accounts.RoleOperator:
		resource.AssignedPrivileges = append(resource.AssignedPrivileges, redfishapi.PRIVILEGESPRIVILEGETYPE_CONFIGURE_COMPONENTS)

	case accounts.RoleReadOnly:
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1ChassisChassisIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdGet(ctx context.Context, chassisId string) (redfishapi.ImplResponse, error) {

	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}

	tag := GetChassisTag()
	uuid := GetChassisUUID()
	version := GetChassisVersion()

	chassis := redfishapi.ChassisV1250Chassis{
		OdataContext: "/redfish/v1/$metadata#Chassis.Chassis",
		OdataId:      "/redfish/v1/Chassis/" + GetChassisName(),
		OdataType:    ChassisVersion,
		AssetTag:     &tag,
		Description:  "CXL Server Chassis",
		Id:           GetChassisTag(),
		Memory: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/Memory",
		},
		MemoryDomains: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/MemoryDomains",
		},
		Name: GetChassisTag(),
		Oem:  nil,
		PCIeDevices: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/PCIeDevices",
		},
		// PowerSubsystem: redfishapi.OdataV4IdRef{
		// 	OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/PowerSubsystem",
		// },
		SerialNumber: &uuid,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
		UUID:    uuid,
		Version: &version,
	}

	return redfishapi.Response(http.StatusOK, chassis), nil
}

// RedfishV1ChassisChassisIdMemoryGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdMemoryGet(ctx context.Context, chassisId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}

	memory := redfishapi.MemoryCollectionMemoryCollection{
		OdataContext: "/redfish/v1/$metadata#MemoryCollection.MemoryCollection",
		OdataId:      "/redfish/v1/Chassis/" + GetChassisName() + "/Memory",
		OdataType:    "#MemoryCollection.MemoryCollection",
		Description:  "Memory DIMM Collection",
		Members: []redfishapi.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/Memory/" + GetMemoryName(),
			},
		},
		Name: "Memory DIMM Collection",
		Oem:  nil,
	}
	for i := 1; i <= GetCXLDevCnt(); i++ {
		memory.Members = append(memory.Members,
			redfishapi.OdataV4IdRef{
				OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/Memory/CXL" + fmt.Sprint(i),
			})
	}
	memory.MembersodataCount = int64(len(memory.Members))

	return redfishapi.Response(http.StatusOK, memory), nil
}

// RedfishV1ChassisChassisIdMemoryMemoryIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdMemoryMemoryIdGet(ctx context.Context, chassisId string, memoryId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}

	if !CheckMemoryName(memoryId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory id (%s) does not exist", memoryId)
	}

	total := int64(0)
	if memoryId == GetMemoryName() {
		// Total system memory in bytes
		total = int64(memory.TotalMemory())
		// Total system memory in megabytes
		total = total / (1024 * 1024)
	} else {
		_, num, _ := IdParse(memoryId)
		dev := GetCXLDevInfoByIndex(num)
		// Total system memory in bytes
		total = dev.GetMemorySize()
		// Total system memory in megabytes
		total = total / (1024 * 1024)
	}

	memory := redfishapi.MemoryV1190Memory{
		OdataContext: "/redfish/v1/$metadata#Memory.Memory",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/Memory/" + memoryId,
		OdataType:    MemoryVersion,
		CapacityMiB:  &total,
		Id:           memoryId,
		Enabled:      true,
		Name:         memoryId,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
	}

	return redfishapi.Response(http.StatusOK, memory), nil

}

// RedfishV1ChassisGet -
func (s *CxlHostApiService) RedfishV1ChassisGet(ctx context.Context) (redfishapi.ImplResponse, error) {
	chassis := redfishapi.ChassisCollectionChassisCollection{
		OdataContext: "/redfish/v1/$metadata#ChassisCollection.ChassisCollection",
		OdataId:      "/redfish/v1/Chassis",
		OdataType:    "#ChassisCollection.ChassisCollection",
		Description:  "Chassis Collection",
		Members: []redfishapi.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Chassis/" + GetChassisName(),
			},
		},
		MembersodataCount:    1,
		MembersodataNextLink: "",
		Name:                 "Chassis Collection",
		Oem:                  nil,
	}

	return redfishapi.Response(http.StatusOK, chassis), nil
}

// RedfishV1FabricsFabricIdGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdGet(ctx context.Context, fabricId string) (redfishapi.ImplResponse, error) {
	if !IsFabricSupported(fabricId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Fabric id (%s) is not supported", fabricId)
	}

	zones := int64(ZonesMax)

	fabric := redfishapi.FabricV131Fabric{
		OdataContext: "/redfish/v1/$metadata#Fabric.Fabric",
		OdataId:      "/redfish/v1/Fabrics/" + fabricId,
		OdataType:    FabricVersion,
		Description:  fabricId + " Fabric",
		FabricType:   redfishapi.PROTOCOLPROTOCOL_CXL,
		Id:           fabricId,
		MaxZones:     &zones,
		Name:         fabricId + " Fabric",
		Oem:          nil,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
		Switches: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Fabrics/" + fabricId + "/Switches",
		},
		UUID: GetFabricUUID(fabricId),
	}

	return redfishapi.Response(http.StatusOK, fabric), nil
}

// RedfishV1FabricsFabricIdSwitchesGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdSwitchesGet(ctx context.Context, fabricId string) (redfishapi.ImplResponse, error) {
	if !IsFabricSupported(fabricId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Fabric id (%s) is not supported", fabricId)
	}

	collection := redfishapi.SwitchCollectionSwitchCollection{
		OdataContext: "/redfish/v1/$metadata#SwitchCollection.SwitchCollection",
		OdataId:      "/redfish/v1/Fabrics/" + fabricId + "/Switches",
		OdataType:    "#SwitchCollection.SwitchCollection",
		Description:  fabricId + " Switch Collection",
		Members: []redfishapi.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Fabrics/" + fabricId + "/Switches/" + GetSwitchName(),
			},
		},
		MembersodataCount: 1,
		Name:              fabricId + " Switch Collection",
		Oem:               nil,
	}

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1FabricsFabricIdSwitchesSwitchIdGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdSwitchesSwitchIdGet(ctx context.Context, fabricId string, switchId string) (redfishapi.ImplResponse, error) {
	if !IsFabricSupported(fabricId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Fabric id (%s) is not supported", fabricId)
	}

	if switchId != GetSwitchName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Switch id (%s) does not exist", switchId)
	}

	serial := GetSwitchUUID()

	resource := redfishapi.SwitchV192Switch{
		OdataContext:         "/redfish/v1/$metadata#Switch.Switch",
		OdataId:              "/redfish/v1/Fabrics/" + fabricId + "/Switches/" + switchId,
		OdataType:            SwitchVersion,
		Description:          fabricId + " Switch",
		Enabled:              true,
		Id:                   switchId,
		Name:                 switchId,
		Oem:                  nil,
		Ports:                redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Fabrics/" + fabricId + "/Switches/" + switchId + "/Ports"},
		PowerState:           redfishapi.RESOURCEPOWERSTATE_ON,
		Redundancy:           []redfishapi.RedundancyRedundancy{},
		RedundancyodataCount: 1,
		SerialNumber:         &serial,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
		SupportedProtocols: []redfishapi.ProtocolProtocol{redfishapi.PROTOCOLPROTOCOL_CXL},
		SwitchType:         redfishapi.PROTOCOLPROTOCOL_CXL,
		UUID:               GetSwitchUUID(),
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1FabricsFabricIdSwitchesSwitchIdPortsGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdSwitchesSwitchIdPortsGet(ctx context.Context, fabricId string, switchId string) (redfishapi.ImplResponse, error) {
	if !IsFabricSupported(fabricId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Fabric id (%s) is not supported", fabricId)
	}

	if switchId != GetSwitchName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Switch id (%s) does not exist", switchId)
	}

	collection := redfishapi.PortCollectionPortCollection{
		OdataContext: "/redfish/v1/$metadata#PortCollection.PortCollection",
		OdataId:      "/redfish/v1/Fabrics/" + fabricId + "/Switches/" + switchId + "/Ports",
		OdataType:    "#PortCollection.PortCollection",
		Description:  fabricId + " " + switchId + " Ports",
		Name:         fabricId + " " + switchId + " Ports",
		Oem:          nil,
	}

	// Iterate over all possible Fabrics adding each to the collection
	ports := GetPorts()
	collection.MembersodataCount = int64(len(ports))
	for f := range ports {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Fabrics/%s/Switches/%s/Ports/%s", fabricId, switchId, f)})
	}

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1FabricsFabricIdSwitchesSwitchIdPortsPortIdGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdSwitchesSwitchIdPortsPortIdGet(ctx context.Context, fabricId string, switchId string, portId string) (redfishapi.ImplResponse, error) {
	if !IsFabricSupported(fabricId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Fabric id (%s) is not supported", fabricId)
	}

	if switchId != GetSwitchName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Switch id (%s) does not exist", switchId)
	}

	if !IsPortSupported(portId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Port id (%s) is not supported", portId)
	}

	enabled := true
	detected := true

	resource := redfishapi.PortV1110Port{
		OdataContext:     "/redfish/v1/$metadata#Chassis.Port",
		OdataId:          "/redfish/v1/Fabrics/" + fabricId + "/Switches/" + switchId + "/Ports/" + portId,
		OdataType:        PortVersion,
		Description:      fabricId + " " + portId,
		Enabled:          true,
		Id:               portId,
		InterfaceEnabled: &enabled,
		LinkState:        redfishapi.PORTV1110LINKSTATE_ENABLED,
		LinkStatus:       redfishapi.PORTV1110LINKSTATUS_LINK_UP,
		Name:             portId,
		Oem:              nil,
		PortProtocol:     redfishapi.PROTOCOLPROTOCOL_CXL,
		PortType:         redfishapi.PORTV1110PORTTYPE_BIDIRECTIONAL_PORT,
		SignalDetected:   &detected,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1FabricsGet -
func (s *CxlHostApiService) RedfishV1FabricsGet(ctx context.Context) (redfishapi.ImplResponse, error) {
	collection := redfishapi.FabricCollectionFabricCollection{
		OdataContext: "/redfish/v1/$metadata#FabricCollection.FabricCollection",
		OdataId:      "/redfish/v1/Fabrics",
		OdataType:    "#FabricCollection.FabricCollection",
		Description:  "Fabrics Collection",
		Name:         "Fabrics Collection",
		Oem:          nil,
	}

	// Iterate over all possible Fabrics adding each to the collection
	fabrics := GetFabrics()
	collection.MembersodataCount = int64(len(fabrics))
	for f := range fabrics {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Fabrics/%s", f)})
	}

	return redfishapi.Response(http.StatusOK, collection), nil
}

// ServiceVersion - The ServiceVersion schema describes the version of the Redfish Service, located at the '/redfish' URI.
type ServiceVersion struct {
	Version string `json:"v1"`
}

// RedfishGet -
func (s *CxlHostApiService) RedfishGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	version := ServiceVersion{
		Version: "/redfish/v1",
	}

	return redfishapi.Response(http.StatusOK, version), nil
}

// RedfishV1Get -
func (s *CxlHostApiService) RedfishV1Get(ctx context.Context) (redfishapi.ImplResponse, error) {
	vendor := "Seagate"
	product := "CXL Host Redfish Service"

	oem := map[string]interface{}{}
	oem["ServiceVersion"] = Version
	uuid := GetServiceUUID()

	root := redfishapi.ServiceRootV1161ServiceRoot{
		OdataContext:   "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
		OdataId:        "/redfish/v1",
		OdataType:      ServiceRootVersion,
		AccountService: redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/AccountService"},
		Chassis:        redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Chassis"},
		Description:    "CXL Host ServiceRoot",
		// Fabrics:               redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Fabrics"},
		Id:                    "CXL Host ServiceRoot",
		Links:                 redfishapi.ServiceRootV1161Links{Oem: nil, Sessions: redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/SessionService/Sessions"}},
		Name:                  "CXL Host",
		Oem:                   oem,
		Product:               &product,
		RedfishVersion:        "1.16.0",
		ServiceIdentification: "",
		SessionService:        redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/SessionService"},
		Systems:               redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Systems"},
		UUID:                  &uuid,
		Vendor:                &vendor,
	}

	return redfishapi.Response(http.StatusOK, root), nil
}

// RedfishV1SessionServiceGet -
func (s *CxlHostApiService) RedfishV1SessionServiceGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	enabled := true

	resource := redfishapi.SessionServiceV118SessionService{
		OdataContext:   "/redfish/v1/$metadata#SessionService.SessionService",
		OdataId:        "/redfish/v1/SessionService",
		OdataType:      SessionServiceVersion,
		Description:    "Session Service",
		Id:             "Session Service",
		Name:           "Session Service",
		Oem:            nil,
		ServiceEnabled: &enabled,
		SessionTimeout: int64(SessionTimeout),
		Sessions:       redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/SessionService/Sessions"},
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1SessionServiceSessionsGet -
func (s *CxlHostApiService) RedfishV1SessionServiceSessionsGet(ctx context.Context) (redfishapi.ImplResponse, error) {
	collection := redfishapi.SessionCollectionSessionCollection{
		OdataContext: "/redfish/v1/$metadata#SessionCollection.SessionCollection",
		OdataId:      "/redfish/v1/SessionService/Sessions",
		OdataType:    "#SessionCollection.SessionCollection",
		Description:  "Sessions",
		Name:         "Sessions",
		Oem:          nil,
	}

	// Iterate over all possible sessions adding each to the collection
	sessions := accounts.GetSessions()
	collection.MembersodataCount = int64(len(sessions))
	if collection.MembersodataCount > 0 {
		for id := range sessions {
			collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/SessionService/Sessions/%s", id)})
		}
	} else {
		collection.Members = []redfishapi.OdataV4IdRef{}
	}

	return redfishapi.Response(http.StatusOK, collection), nil
}

// FillInSessionResource: Function to fill a Redfish Session resource consistently
func FillInSessionResource(sessionId string, session *accounts.SessionInformation, fillPassword bool) redfishapi.SessionV171Session {

	resource := redfishapi.SessionV171Session{
		OdataId:     "/redfish/v1/SessionService/Sessions/" + sessionId,
		OdataType:   SessionVersion,
		Description: "User Session",
		Id:          sessionId,
		Name:        "User Session " + sessionId,
		Password:    nil,
		SessionType: redfishapi.SESSIONV171SESSIONTYPES_REDFISH,
	}

	if session != nil {
		resource.UserName = &session.Username
		resource.CreatedTime = &session.Created
		if fillPassword {
			resource.Password = &session.Token
		}
	}

	return resource
}

// RedfishV1SessionServiceSessionsPost -
func (s *CxlHostApiService) RedfishV1SessionServiceSessionsPost(ctx context.Context, sessionV171Session redfishapi.SessionV171Session) (redfishapi.ImplResponse, error) {

	// Validate that a username and password were supplied
	if *sessionV171Session.UserName == "" {
		return redfishapi.Response(http.StatusBadRequest, nil), errors.New("UserName is required for creating a new session")
	}

	if *sessionV171Session.Password == "" {
		return redfishapi.Response(http.StatusBadRequest, nil), errors.New("Password is required for creating a new session")
	}

	// Create the session
	session := accounts.CreateSession(ctx, *sessionV171Session.UserName, *sessionV171Session.Password)
	if session == nil {
		return redfishapi.Response(http.StatusUnauthorized, nil), errors.New("Invalid user credentials")
	}

	resource := FillInSessionResource(session.Id, session, true)

	return redfishapi.Response(http.StatusCreated, resource), nil
}

// RedfishV1SessionServiceSessionsSessionIdDelete -
func (s *CxlHostApiService) RedfishV1SessionServiceSessionsSessionIdDelete(ctx context.Context, sessionId string) (redfishapi.ImplResponse, error) {

	if !accounts.IsSessionIdActive(ctx, sessionId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Session id (%s) is not active", sessionId)
	}

	session := accounts.DeleteSession(ctx, sessionId)

	if session == nil {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Session id (%s) was not found in the system", sessionId)
	}

	resource := FillInSessionResource(sessionId, session, false)

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1SessionServiceSessionsSessionIdGet -
func (s *CxlHostApiService) RedfishV1SessionServiceSessionsSessionIdGet(ctx context.Context, sessionId string) (redfishapi.ImplResponse, error) {

	if !accounts.IsSessionIdActive(ctx, sessionId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Session id (%s) is not active", sessionId)
	}

	session := accounts.GetSessionInformation(ctx, sessionId)

	resource := FillInSessionResource(sessionId, session, false)

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1ChassisChassisIdMemoryDomainsGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdMemoryDomainsGet(ctx context.Context, chassisId string) (redfishapi.ImplResponse, error) {
	// Hardcode to two memory domains
	// DIMMs domain contains memory chunks with each chunk from one CPU socket
	// CXL domain contains memory chunks with each chunk from a type 3 CXL device ( only show up when CXL device presents)

	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}
	collection := redfishapi.MemoryDomainCollectionMemoryDomainCollection{
		OdataId:   "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains",
		OdataType: "#MemoryDomainCollection.MemoryDomainCollection",
		Name:      "MemoryDomain Collection",
	}
	collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/MemoryDomains/DIMMs", chassisId)})
	if CheckMemoryDomainName("CXL") {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/MemoryDomains/CXL", chassisId)})
	}
	collection.MembersodataCount = int64(len(collection.Members))
	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1ChassisChassisIdMemoryDomainsMemoryDomainIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdMemoryDomainsMemoryDomainIdGet(ctx context.Context, chassisId string, memoryDomainId string) (redfishapi.ImplResponse, error) {

	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}
	if !CheckMemoryDomainName(memoryDomainId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory Domain id (%s) does not exist", memoryDomainId)
	}

	resource := redfishapi.MemoryDomainV150MemoryDomain{
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/" + memoryDomainId,
		OdataType:    MemoryDomainVersion,
		MemoryChunks: redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/MemoryDomains/%s/MemoryChunks", chassisId, memoryDomainId)},
		Id:           memoryDomainId,
		Name:         "Memory Domain",
	}
	if memoryDomainId == "CXL" {
		for _, bdf := range GetCXLDevList() {
			for _, cxlFuncId := range GetCXLLogicalDeviceIds(BDFtoBD(bdf)) {
				resource.Links.CXLLogicalDevices = append(resource.Links.CXLLogicalDevices, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/PCIeDevices/%s/CXLLogicalDevices/%s", chassisId, BDFtoBD(bdf), cxlFuncId)})
			}
			for _, pcieFuncId := range GetPCIeFunctionIds(BDFtoBD(bdf)) {
				resource.Links.PCIeFunctions = append(resource.Links.PCIeFunctions, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/PCIeDevices/%s/PCIeFunctions/%s", chassisId, BDFtoBD(bdf), pcieFuncId)})
			}
		}
		resource.Links.CXLLogicalDevicesodataCount = int64(len(resource.Links.CXLLogicalDevices))
		resource.Links.PCIeFunctionsodataCount = int64(len(resource.Links.PCIeFunctions))
		resource.Name = "CXL Memory Domain"
		resource.Description = "Local CXL Memory Domain"
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1ChassisChassisIdMemoryDomainsMemoryDomainIdMemoryChunksGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdMemoryDomainsMemoryDomainIdMemoryChunksGet(ctx context.Context, chassisId string, memoryDomainId string) (redfishapi.ImplResponse, error) {

	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}
	if !CheckMemoryDomainName(memoryDomainId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory Domain id (%s) does not exist", memoryDomainId)
	}

	collection := redfishapi.MemoryChunksCollectionMemoryChunksCollection{
		OdataId:   "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/" + memoryDomainId + "/MemoryChunks",
		OdataType: "#MemoryChunksCollection.MemoryChunksCollection",
		Name:      "Memory Chunks Collection",
	}

	if memoryDomainId == "DIMMs" { // each numa node with dimm attached ( numa node with CPU ) is mapped to a memory chunk
		localNodes := GetDimmAttachedNumaNodes()
		for _, node := range localNodes {
			collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/MemoryDomains/DIMMs/MemoryChunks/%s", chassisId, node)})
		}

	} else { // each type 3 CXL device is mapped to a memory chunk
		cxlDevList := GetCXLWithMemDevList()
		for _, dev := range cxlDevList {
			collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/MemoryDomains/CXL/MemoryChunks/%s", chassisId, BDFtoBD(dev))})
		}
	}
	collection.MembersodataCount = int64(len(collection.Members))

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1ChassisChassisIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdGet(ctx context.Context, chassisId string, memoryDomainId string, memoryChunksId string) (redfishapi.ImplResponse, error) {

	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}
	if !CheckMemoryDomainName(memoryDomainId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory Domain id (%s) does not exist", memoryDomainId)
	}
	if !CheckMemoryChunkName(memoryDomainId, memoryChunksId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory chunk id (%s) does not exist", memoryChunksId)
	}

	resource := redfishapi.MemoryChunksV161MemoryChunks{
		OdataId:          "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/" + memoryDomainId + "/MemoryChunks/" + memoryChunksId,
		OdataType:        MemoryChunkVersion,
		Name:             "Memory Chunk",
		Id:               memoryChunksId,
		AddressRangeType: redfishapi.MEMORYCHUNKSV161ADDRESSRANGETYPE_VOLATILE,
	}
	if memoryDomainId == "DIMMs" {
		mem := GetNumaMemInfo(memoryChunksId)
		memMiB := mem.MemTotal >> 10 // MemTotal is in KiB
		resource.MemoryChunkSizeMiB = &memMiB
	} else {
		memAddr := GetCxlAddressInfoMiB(BDtoBDF(memoryChunksId))
		resource.AddressRangeOffsetMiB = &memAddr.BaseAddress
		resource.MemoryChunkSizeMiB = &memAddr.Size

		resource.Links.Endpoints = append(resource.Links.Endpoints, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/Memory/CXL%d", chassisId, CxlDevBDFtoIndex(BDtoBDF(memoryChunksId)))})
		resource.Links.EndpointsodataCount = 1

		resource.Links.CXLLogicalDevices = append(resource.Links.CXLLogicalDevices, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/PCIeDevices/%s/CXLLogicalDevices/0", chassisId, memoryChunksId)})
		resource.Links.CXLLogicalDevicesodataCount = 1

	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1ChassisChassisIdPCIeDevicesGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdPCIeDevicesGet(ctx context.Context, chassisId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("chassis id (%s) does not exist", chassisId)
	}

	collection := redfishapi.PcieDeviceCollectionPcieDeviceCollection{
		OdataContext: "/redfish/v1/$metadata#PCIeDeviceCollection.PCIeDeviceCollection",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices",
		OdataType:    "#PCIeDeviceCollection.PCIeDeviceCollection",
		Description:  "PCIeDevice Collection",
		Name:         "PCIeDevice Collection",
		Oem:          nil,
	}

	// Only show CXL device in the PCIe device collection. Use CXL dev name as PCIe dev name
	cxlDevList := GetCXLDevList()
	for _, dev := range cxlDevList {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/PCIeDevices/%s", chassisId, BDFtoBD(dev))})
	}

	collection.MembersodataCount = int64(len(collection.Members))

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdCXLLogicalDevicesCXLLogicalDeviceIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdCXLLogicalDevicesCXLLogicalDeviceIdGet(ctx context.Context, chassisId string, pCIeDeviceId string, cXLLogicalDeviceId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("chassis id (%s) does not exist", chassisId)
	}

	if !CheckCxlChunk(pCIeDeviceId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the PCIe id (%s) does not exist", pCIeDeviceId)
	}

	if !slices.Contains(GetCXLLogicalDeviceIds(pCIeDeviceId), cXLLogicalDeviceId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the CXL Logical Device id (%s) does not exist", cXLLogicalDeviceId)
	}

	dev := GetCXLDevInfo(BDtoBDF(pCIeDeviceId))
	gcxlid := FormatGCXLID(dev)
	fmt := redfishapi.ResourceV1190DurableNameFormat("GCXLID")
	resource := redfishapi.CxlLogicalDeviceV111CxlLogicalDevice{
		OdataContext: "/redfish/v1/$metadata#CXLLogicalDevice.CXLLogicalDevice",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices/" + cXLLogicalDeviceId,
		OdataType:    CXLLogicalDevice,
		Description:  "CXL Logical Device " + string(dev.GetCxlType()),
		Id:           cXLLogicalDeviceId,
		Links: redfishapi.CxlLogicalDeviceV111Links{
			MemoryChunks: []redfishapi.OdataV4IdRef{
				{OdataId: "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/CXL/MemoryChunks/" + pCIeDeviceId},
			},
			MemoryChunksodataCount: 1,
			MemoryDomains: []redfishapi.OdataV4IdRef{
				{OdataId: "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/CXL"},
			},
			MemoryDomainsodataCount: 1,
			PCIeFunctions: []redfishapi.OdataV4IdRef{
				{OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/PCIeFunctions/" + cXLLogicalDeviceId},
			},
			PCIeFunctionsodataCount: 1,
		},
		MemorySizeMiB:      dev.GetMemorySize() >> 20,
		Name:               "Locally attached CXL Logical Device " + string(dev.GetCxlType()),
		SemanticsSupported: []redfishapi.CxlLogicalDeviceV111CxlSemantic{},
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
		Identifiers: []redfishapi.ResourceIdentifier{{
			DurableName:       &gcxlid,
			DurableNameFormat: &fmt,
		}},
	}
	devCap := dev.GetCxlCap()
	if devCap.IO_En {
		resource.SemanticsSupported = append(resource.SemanticsSupported, redfishapi.CXLLOGICALDEVICEV111CXLSEMANTIC_CX_LIO)
	}
	if devCap.Cache_En {
		resource.SemanticsSupported = append(resource.SemanticsSupported, redfishapi.CXLLOGICALDEVICEV111CXLSEMANTIC_CX_LCACHE)
	}
	if devCap.Mem_En {
		resource.SemanticsSupported = append(resource.SemanticsSupported, redfishapi.CXLLOGICALDEVICEV111CXLSEMANTIC_CX_LMEM)
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdCXLLogicalDevicesGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdCXLLogicalDevicesGet(ctx context.Context, chassisId string, pCIeDeviceId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("chassis id (%s) does not exist", chassisId)
	}

	if !CheckCxlChunk(pCIeDeviceId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the PCIe id (%s) does not exist", pCIeDeviceId)
	}

	collection := redfishapi.CxlLogicalDeviceCollectionCxlLogicalDeviceCollection{
		OdataContext:         "/redfish/v1/$metadata#CXLLogicalDeviceCollection.CXLLogicalDeviceCollection",
		OdataId:              "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices",
		OdataType:            "#CXLLogicalDeviceCollection.CXLLogicalDeviceCollection",
		Description:          "CXL Logical Device Collection",
		MembersodataNextLink: "",
		Name:                 "CXL Logical Device Collection",
	}
	for _, id := range GetCXLLogicalDeviceIds(pCIeDeviceId) {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices/" + id,
		})
	}
	collection.MembersodataCount = int64(len(collection.Members))
	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdGet(ctx context.Context, chassisId string, pCIeDeviceId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("chassis id (%s) does not exist", chassisId)
	}

	if !CheckCxlChunk(pCIeDeviceId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the PCIe id (%s) does not exist", pCIeDeviceId)
	}

	dev := GetCXLDevInfo(BDtoBDF(pCIeDeviceId))
	SN := dev.GetSerialNumber()
	resource := redfishapi.PcieDeviceV1130PcieDevice{
		OdataContext: "/redfish/v1/$metadata#PCIeDevice.PCIeDevice",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId,
		OdataType:    PCIeDevice,
		Actions:      redfishapi.PcieDeviceV1130Actions{},
		Assembly:     redfishapi.OdataV4IdRef{},
		CXLDevice: redfishapi.PcieDeviceV1130CxlDevice{
			DeviceType:              redfishapi.PcieDeviceV1130CxlDeviceType(dev.GetCxlType()),
			MaxNumberLogicalDevices: new(int64),
		},
		CXLLogicalDevices: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices",
		},
		Description: "PCIe Device " + pCIeDeviceId,
		DeviceType:  redfishapi.PCIEDEVICEV1130DEVICETYPE_SINGLE_FUNCTION,
		EnvironmentMetrics: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId,
		},
		Id: pCIeDeviceId,
		Links: redfishapi.PcieDeviceV1130Links{
			Chassis: []redfishapi.OdataV4IdRef{
				{
					OdataId: "/redfish/v1/Chassis/" + chassisId,
				},
			},
			ChassisodataCount: 1,
		},
		Name:         "PCIe Device " + pCIeDeviceId,
		SerialNumber: &SN,
		Oem:          nil,
		PCIeFunctions: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/PCIeFunctions",
		},
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdPCIeFunctionsGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdPCIeFunctionsGet(ctx context.Context, chassisId string, pCIeDeviceId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("chassis id (%s) does not exist", chassisId)
	}

	if !CheckCxlChunk(pCIeDeviceId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the PCIe id (%s) does not exist", pCIeDeviceId)
	}

	collection := redfishapi.PcieFunctionCollectionPcieFunctionCollection{
		OdataContext: "/redfish/v1/$metadata#PCIeFunctionCollection.PCIeFunctionCollection",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/PCIeFunctions",
		OdataType:    "#PCIeFunctionCollection.PCIeFunctionCollection",
		Description:  "PCIeFunction Collection",
		Name:         "PCIe Function Collection",
		Oem:          nil,
	}

	for _, id := range GetPCIeFunctionIds(pCIeDeviceId) {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Chassis/%s/PCIeDevices/%s/PCIeFunctions/%s", chassisId, pCIeDeviceId, id)})
	}

	collection.MembersodataCount = int64(len(collection.Members))

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdPCIeFunctionsPCIeFunctionIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdPCIeDevicesPCIeDeviceIdPCIeFunctionsPCIeFunctionIdGet(ctx context.Context, chassisId string, pCIeDeviceId string, pCIeFunctionId string) (redfishapi.ImplResponse, error) {
	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("chassis id (%s) does not exist", chassisId)
	}

	if !CheckCxlChunk(pCIeDeviceId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the PCIe id (%s) does not exist", pCIeDeviceId)
	}

	if !slices.Contains(GetPCIeFunctionIds(pCIeDeviceId), pCIeFunctionId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("the PCIe Function id (%s) does not exist", pCIeFunctionId)
	}

	dev := GetCXLDevInfo(BDtoBDF(pCIeDeviceId))
	revisionId := fmt.Sprintf("%d", dev.GetPcieHdr().Rev_ID)
	classCode := fmt.Sprintf("%X", dev.GetPcieHdr().Class_Code)
	functionId, _ := strconv.Atoi(pCIeFunctionId)
	functionid := int64(functionId)
	vendorId := fmt.Sprintf("%X", dev.GetPcieHdr().Vendor_ID)
	resource := redfishapi.PcieFunctionV151PcieFunction{
		OdataContext:     "/redfish/v1/$metadata#PCIeFunction.PCIeFunction",
		OdataId:          "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/PCIeFunctions/" + pCIeFunctionId,
		OdataType:        PCIeFunction,
		Description:      "PCIe Function",
		ClassCode:        &classCode,
		DeviceClass:      redfishapi.PCIEFUNCTIONV151DEVICECLASS_UNCLASSIFIED_DEVICE,
		DeviceId:         &pCIeDeviceId,
		FunctionId:       &functionid,
		FunctionProtocol: redfishapi.PCIEFUNCTIONV151FUNCTIONPROTOCOL_CXL,
		FunctionType:     redfishapi.PCIEFUNCTIONV151FUNCTIONTYPE_PHYSICAL,
		Id:               pCIeDeviceId + pCIeFunctionId,
		Links: redfishapi.PcieFunctionV151Links{
			CXLLogicalDevice: redfishapi.OdataV4IdRef{
				OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices/" + pCIeFunctionId,
			},
			MemoryDomains: []redfishapi.OdataV4IdRef{
				{
					OdataId: "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/CXL",
				},
			},
			MemoryDomainsodataCount: 1,
			PCIeDevice: redfishapi.OdataV4IdRef{
				OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId,
			},
		},
		Name:       "PCIe Function",
		RevisionId: &revisionId,
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
		VendorId: &vendorId,
	}
	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1MetadataGet -
func (s *CxlHostApiService) RedfishV1MetadataGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	metadata := `<?xml version="1.0" encoding="UTF-8"?>
	<edmx:Edmx xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx" Version="4.0">
		<edmx:DataServices>
			<Schema xmlns="http://docs.oasis-open.org/odata/ns/edm" Namespace="Service">
				<EntityContainer Name="Service" Extends="ServiceRoot.v1_2_0.ServiceContainer" />
			</Schema>
		</edmx:DataServices>
		<edmx:Reference Uri="https://redfish.dmtf.org/schemas/v1/AccountService_v1.xml">
			<edmx:Include Namespace="AccountService" />
				<edmx:Include Namespace="AccountService.v1_11_1" />
		</edmx:Reference>
		<edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/Chassis_v1.xml">
		   <edmx:Include Namespace="Chassis" />
			   <edmx:Include Namespace="Chassis.v1_21_0" />
		</edmx:Reference>
		<edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/ChassisCollection_v1.xml">
		  <edmx:Include Namespace="ChassisCollection" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/Fabric_v1.xml">
		  <edmx:Include Namespace="Fabric" />
			  <edmx:Include Namespace="Fabric.v1_3_0" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/FabricCollection_v1.xml">
		  <edmx:Include Namespace="FabricCollection" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/Memory_v1.xml">
		  <edmx:Include Namespace="Memory" />
			  <edmx:Include Namespace="Memory.v1_16_0" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/MemoryCollection_v1.xml">
		  <edmx:Include Namespace="MemoryCollection" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/Port_v1.xml">
		  <edmx:Include Namespace="Port" />
			  <edmx:Include Namespace="Port.v1_7_0" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/PortCollection_v1.xml">
		  <edmx:Include Namespace="PortCollection" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/ServiceRoot_v1.xml">
		  <edmx:Include Namespace="ServiceRoot" />
			  <edmx:Include Namespace="ServiceRoot.v1_14_0" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/Session_v1.xml">
		  <edmx:Include Namespace="Session" />
			  <edmx:Include Namespace="Session.v1_5_0" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/SessionCollection_v1.xml">
		  <edmx:Include Namespace="SessionCollection" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/SessionService_v1.xml">
		  <edmx:Include Namespace="SessionService" />
		  <edmx:Include Namespace="SessionService.v1_1_8" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/Switch_v1.xml">
		  <edmx:Include Namespace="Switch" />
			  <edmx:Include Namespace="Switch.v1_8_0" />
	   </edmx:Reference>
	   <edmx:Reference Uri="http://redfish.dmtf.org/schemas/v1/SwitchCollection_v1.xml">
		  <edmx:Include Namespace="SwitchCollection" />
	   </edmx:Reference>
	</edmx:Edmx>`

	return redfishapi.Response(http.StatusOK, metadata), nil
}

// RedfishV1OdataGet -
func (s *CxlHostApiService) RedfishV1OdataGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	resource := redfishapi.RedfishV1OdataGet200Response{
		OdataContext: "/redfish/v1/odata",
	}

	resource.Value = append(resource.Value, redfishapi.RedfishV1OdataGet200ResponseValueInner{
		Kind: "Singleton",
		Name: "Service",
		Url:  "/redfish/v1",
	})

	resource.Value = append(resource.Value, redfishapi.RedfishV1OdataGet200ResponseValueInner{
		Kind: "Singleton",
		Name: "AccountService",
		Url:  "/redfish/v1/AccountService",
	})

	resource.Value = append(resource.Value, redfishapi.RedfishV1OdataGet200ResponseValueInner{
		Kind: "Singleton",
		Name: "Chassis",
		Url:  "/redfish/v1/Chassis",
	})

	resource.Value = append(resource.Value, redfishapi.RedfishV1OdataGet200ResponseValueInner{
		Kind: "Singleton",
		Name: "Fabrics",
		Url:  "/redfish/v1/Fabrics",
	})

	resource.Value = append(resource.Value, redfishapi.RedfishV1OdataGet200ResponseValueInner{
		Kind: "Singleton",
		Name: "SessionService",
		Url:  "/redfish/v1/SessionService",
	})

	resource.Value = append(resource.Value, redfishapi.RedfishV1OdataGet200ResponseValueInner{
		Kind: "Singleton",
		Name: "Sessions",
		Url:  "/redfish/v1/SessionService/Sessions",
	})

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1SystemsComputerSystemIdActionsComputerSystemResetPost -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdActionsComputerSystemResetPost(ctx context.Context, computerSystemId string, computerSystemV1220ResetRequestBody redfishapi.ComputerSystemV1220ResetRequestBody) (redfishapi.ImplResponse, error) {
	ResetType := computerSystemV1220ResetRequestBody.ResetType
	switch ResetType {
	case redfishapi.RESOURCERESETTYPE_FORCE_OFF, redfishapi.RESOURCERESETTYPE_GRACEFUL_SHUTDOWN:
		AsyncReboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
	case redfishapi.RESOURCERESETTYPE_GRACEFUL_RESTART, redfishapi.RESOURCERESETTYPE_FORCE_RESTART, redfishapi.RESOURCERESETTYPE_POWER_CYCLE:
		AsyncReboot(syscall.LINUX_REBOOT_CMD_RESTART)
	case redfishapi.RESOURCERESETTYPE_SUSPEND:
		AsyncReboot(syscall.LINUX_REBOOT_CMD_SW_SUSPEND)
	// case redfishapi.RESOURCERESETTYPE_ON:
	// case redfishapi.RESOURCERESETTYPE_NMI:
	// case redfishapi.RESOURCERESETTYPE_FORCE_ON:
	// case redfishapi.RESOURCERESETTYPE_PUSH_POWER_BUTTON:
	// case redfishapi.RESOURCERESETTYPE_PAUSE:
	default:
		return redfishapi.Response(http.StatusBadRequest, nil), fmt.Errorf("Reset Type %s is not supported", ResetType)
	}
	return redfishapi.Response(http.StatusNoContent, nil), nil
}

// RedfishV1SystemsComputerSystemIdGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdGet(ctx context.Context, computerSystemId string) (redfishapi.ImplResponse, error) {
	if computerSystemId != GetSystemName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("System id (%s) does not exist", computerSystemId)
	}

	system := redfishapi.ComputerSystemV1220ComputerSystem{
		OdataContext: "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		OdataId:      "/redfish/v1/Systems/" + GetSystemName(),
		OdataType:    SystemVersion,
		Actions: redfishapi.ComputerSystemV1220Actions{
			ComputerSystemReset: redfishapi.ComputerSystemV1220Reset{
				Target: "/redfish/v1/Systems/" + GetSystemName() + "/Actions/ComputerSystem.Reset",
			},
		},
		Memory: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/" + GetSystemName() + "/Memory",
		},
		MemoryDomains: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Systems/" + GetSystemName() + "/MemoryDomains",
		},
		PCIeDevices:   []redfishapi.OdataV4IdRef{},
		PCIeFunctions: []redfishapi.OdataV4IdRef{},
	}

	return redfishapi.Response(http.StatusOK, system), nil
}

// RedfishV1SystemsComputerSystemIdMemoryDomainsGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryDomainsGet(ctx context.Context, computerSystemId string) (redfishapi.ImplResponse, error) {
	// Hardcode to two memory domains
	// DIMMs domain contains memory chunks with each chunk from one CPU socket
	// CXL domain contains memory chunks with each chunk from a type 3 CXL device ( only show up when CXL device presents)

	if computerSystemId != GetSystemName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("System id (%s) does not exist", computerSystemId)
	}
	collection := redfishapi.MemoryDomainCollectionMemoryDomainCollection{
		OdataId:   "/redfish/v1/Systems/" + computerSystemId + "/MemoryDomains",
		OdataType: "#MemoryDomainCollection.MemoryDomainCollection",
		Name:      "MemoryDomain Collection",
	}
	collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Systems/%s/MemoryDomains/DIMMs", computerSystemId)})
	if CheckMemoryDomainName("CXL") {
		collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Systems/%s/MemoryDomains/CXL", computerSystemId)})
	}
	collection.MembersodataCount = int64(len(collection.Members))
	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdGet(ctx context.Context, computerSystemId string, memoryDomainId string) (redfishapi.ImplResponse, error) {

	if computerSystemId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Systems id (%s) does not exist", computerSystemId)
	}
	if !CheckMemoryDomainName(memoryDomainId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory Domain id (%s) does not exist", memoryDomainId)
	}

	resource := redfishapi.MemoryDomainV150MemoryDomain{
		OdataId:      "/redfish/v1/Systems/" + computerSystemId + "/MemoryDomains/" + memoryDomainId,
		OdataType:    MemoryDomainVersion,
		MemoryChunks: redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Systems/%s/MemoryDomains/%s/MemoryChunks", computerSystemId, memoryDomainId)},
		Id:           memoryDomainId,
		Name:         "Memory Domain",
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksGet(ctx context.Context, computerSystemId string, memoryDomainId string) (redfishapi.ImplResponse, error) {

	if computerSystemId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Systems id (%s) does not exist", computerSystemId)
	}
	if !CheckMemoryDomainName(memoryDomainId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory Domain id (%s) does not exist", memoryDomainId)
	}

	collection := redfishapi.MemoryChunksCollectionMemoryChunksCollection{
		OdataId:   "/redfish/v1/Systems/" + computerSystemId + "/MemoryDomains/" + memoryDomainId + "/MemoryChunks",
		OdataType: "#MemoryChunksCollection.MemoryChunksCollection",
		Name:      "Memory Chunks Collection",
	}

	if memoryDomainId == "DIMMs" { // each numa node with dimm attached ( numa node with CPU ) is mapped to a memory chunk
		localNodes := GetDimmAttachedNumaNodes()
		for _, node := range localNodes {
			collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Systems/%s/MemoryDomains/DIMMs/MemoryChunks/%s", computerSystemId, node)})
		}

	} else { // each numa node with CXL memory attached
		for node := range GetCXLNumaNodes() {
			collection.Members = append(collection.Members, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Systems/%s/MemoryDomains/CXL/MemoryChunks/%s", computerSystemId, node)})
		}
	}
	collection.MembersodataCount = int64(len(collection.Members))

	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdGet(ctx context.Context, computerSystemId string, memoryDomainId string, memoryChunksId string) (redfishapi.ImplResponse, error) {

	if computerSystemId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Systems id (%s) does not exist", computerSystemId)
	}
	if !CheckMemoryDomainName(memoryDomainId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory Domain id (%s) does not exist", memoryDomainId)
	}
	if !CheckMemoryChunkNumaName(memoryDomainId, memoryChunksId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory chunk id (%s) does not exist", memoryChunksId)
	}

	resource := redfishapi.MemoryChunksV161MemoryChunks{
		OdataId:          "/redfish/v1/Systems/" + computerSystemId + "/MemoryDomains/" + memoryDomainId + "/MemoryChunks/" + memoryChunksId,
		OdataType:        MemoryChunkVersion,
		Name:             "Memory Chunk " + memoryChunksId,
		Id:               memoryChunksId,
		AddressRangeType: redfishapi.MEMORYCHUNKSV161ADDRESSRANGETYPE_VOLATILE,
	}
	if memoryDomainId == "DIMMs" {
		mem := GetNumaMemInfo(memoryChunksId)
		memMiB := mem.MemTotal >> 10 // MemTotal is in KiB
		resource.MemoryChunkSizeMiB = &memMiB

	} else {
		bdf := CxlDevNodeToBDF(memoryChunksId)
		memAddr := GetCxlAddressInfoMiB(bdf)
		resource.AddressRangeOffsetMiB = &memAddr.BaseAddress
		resource.MemoryChunkSizeMiB = &memAddr.Size
		resource.Links.Endpoints = append(resource.Links.Endpoints, redfishapi.OdataV4IdRef{OdataId: fmt.Sprintf("/redfish/v1/Systems/%s/Memory/CXL%d", computerSystemId, CxlDevBDFtoIndex(CxlDevNodeToBDF(memoryChunksId)))})
		resource.Links.EndpointsodataCount = 1

		perfMetric := GetCXLMemPerf(bdf)
		if perfMetric != nil {
			resource.Oem = map[string]interface{}{"Seagate": map[string]interface{}{"Bandwidth": perfMetric.Bandwidth, "Latency": perfMetric.Latency}}
		}
	}

	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1SystemsComputerSystemIdMemoryGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryGet(ctx context.Context, computerSystemId string) (redfishapi.ImplResponse, error) {
	if computerSystemId != GetSystemName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("System id (%s) does not exist", computerSystemId)
	}
	collection := redfishapi.MemoryCollectionMemoryCollection{
		OdataContext:         "/redfish/v1/$metadata#MemoryCollection.MemoryCollection",
		OdataId:              "/redfish/v1/Systems/" + computerSystemId + "/Memory",
		OdataType:            "#MemoryCollection.MemoryCollection",
		Description:          "Memory Collection",
		MembersodataNextLink: "",
		Name:                 "Memory Collection",
	}
	for i := 1; i <= GetCXLDevCnt(); i++ {
		collection.Members = append(collection.Members,
			redfishapi.OdataV4IdRef{
				OdataId: "/redfish/v1/Systems/" + computerSystemId + "/Memory/CXL" + fmt.Sprint(i),
			})
	}
	collection.MembersodataCount = int64(len(collection.Members))
	return redfishapi.Response(http.StatusOK, collection), nil
}

// RedfishV1SystemsComputerSystemIdMemoryMemoryIdGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryMemoryIdGet(ctx context.Context, computerSystemId string, memoryId string) (redfishapi.ImplResponse, error) {
	if computerSystemId != GetSystemName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("System id (%s) does not exist", computerSystemId)
	}

	if !CheckMemoryName(memoryId) {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Memory id (%s) does not exist", memoryId)
	}

	_, num, _ := IdParse(memoryId)
	dev := GetCXLDevInfoByIndex(num)
	// Get memory size in bytes and convert to MiB
	total := dev.GetMemorySize() / (1024 * 1024)

	resource := redfishapi.MemoryV1190Memory{
		OdataContext: "/redfish/v1/$metadata#Memory.Memory",
		OdataId:      "/redfish/v1/Systems/" + computerSystemId + "/Memory/" + memoryId,
		OdataType:    MemoryVersionForSystem,
		CapacityMiB:  &total,
		Id:           memoryId,
		Name:         "CXL Device memory",
		MemoryMedia:  []redfishapi.MemoryV1190MemoryMedia{redfishapi.MEMORYV1190MEMORYMEDIA_PROPRIETARY},
		Status: redfishapi.ResourceStatus{
			Health: &resourcehealthOk,
			State:  &resourcestateEnabled,
		},
		Links: redfishapi.MemoryV1190Links{
			MemoryMediaSources: []redfishapi.OdataV4IdRef{
				{OdataId: "/redfish/v1/Chassis/" + GetChassisName() + "/MemoryDomains/CXL/MemoryChunks/" + BDFtoBD(CxlDevIndexToBDF(num))},
			},
		},
	}
	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1SystemsGet -
func (s *CxlHostApiService) RedfishV1SystemsGet(ctx context.Context) (redfishapi.ImplResponse, error) {
	system := redfishapi.ComputerSystemCollectionComputerSystemCollection{
		OdataContext: "/redfish/v1/$metadata#ComputerSystemCollection.ComputerSystemCollection",
		OdataId:      "/redfish/v1/Systems",
		OdataType:    "#ComputerSystemCollection.ComputerSystemCollection",
		Description:  "ComputerSystem Collection",
		Members: []redfishapi.OdataV4IdRef{
			{
				OdataId: "/redfish/v1/Systems/" + GetSystemName(),
			},
		},
		MembersodataCount:    1,
		MembersodataNextLink: "",
		Name:                 "ComputerSystem Collection",
		Oem:                  nil,
	}

	return redfishapi.Response(http.StatusOK, system), nil
}

// RedfishV1AccountServiceAccountsManagerAccountIdPut -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsManagerAccountIdPut(ctx context.Context, managerAccountId string, managerAccountV1120ManagerAccount redfishapi.ManagerAccountV1120ManagerAccount) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceAccountsManagerAccountIdPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.ManagerAccountV1120ManagerAccount{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.ManagerAccountV1120ManagerAccount{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServiceAccountsManagerAccountIdPut method not implemented")
}

// RedfishV1AccountServicePatch -
func (s *CxlHostApiService) RedfishV1AccountServicePatch(ctx context.Context, accountServiceV1150AccountService redfishapi.AccountServiceV1150AccountService) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServicePatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.AccountServiceV1150AccountService{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.AccountServiceV1150AccountService{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServicePatch method not implemented")
}

// RedfishV1AccountServicePut -
func (s *CxlHostApiService) RedfishV1AccountServicePut(ctx context.Context, accountServiceV1150AccountService redfishapi.AccountServiceV1150AccountService) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServicePut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.AccountServiceV1150AccountService{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.AccountServiceV1150AccountService{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServicePut method not implemented")
}

// RedfishV1AccountServiceRolesPost -
func (s *CxlHostApiService) RedfishV1AccountServiceRolesPost(ctx context.Context, roleV131Role redfishapi.RoleV131Role) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceRolesPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, redfishapi.RoleV131Role{}) or use other options such as http.Ok ...
	// return redfishapi.Response(201, redfishapi.RoleV131Role{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServiceRolesPost method not implemented")
}

// RedfishV1AccountServiceRolesRoleIdDelete -
func (s *CxlHostApiService) RedfishV1AccountServiceRolesRoleIdDelete(ctx context.Context, roleId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceRolesRoleIdDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.RoleV131Role{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.RoleV131Role{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServiceRolesRoleIdDelete method not implemented")
}

// RedfishV1AccountServiceRolesRoleIdPatch -
func (s *CxlHostApiService) RedfishV1AccountServiceRolesRoleIdPatch(ctx context.Context, roleId string, roleV131Role redfishapi.RoleV131Role) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceRolesRoleIdPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.RoleV131Role{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.RoleV131Role{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServiceRolesRoleIdPatch method not implemented")
}

// RedfishV1AccountServiceRolesRoleIdPut -
func (s *CxlHostApiService) RedfishV1AccountServiceRolesRoleIdPut(ctx context.Context, roleId string, roleV131Role redfishapi.RoleV131Role) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceRolesRoleIdPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.RoleV131Role{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.RoleV131Role{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServiceRolesRoleIdPut method not implemented")
}

// RedfishV1SessionServicePatch -
func (s *CxlHostApiService) RedfishV1SessionServicePatch(ctx context.Context, sessionServiceV118SessionService redfishapi.SessionServiceV118SessionService) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1SessionServicePatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.SessionServiceV118SessionService{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.SessionServiceV118SessionService{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServicePatch method not implemented")
}

// RedfishV1SessionServicePut -
func (s *CxlHostApiService) RedfishV1SessionServicePut(ctx context.Context, sessionServiceV118SessionService redfishapi.SessionServiceV118SessionService) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1SessionServicePut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.SessionServiceV118SessionService{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.SessionServiceV118SessionService{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SessionServicePut method not implemented")
}

// RedfishV1FabricsFabricIdConnectionsConnectionIdDelete -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdConnectionsConnectionIdDelete(ctx context.Context, fabricId string, connectionId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1FabricsFabricIdConnectionsConnectionIdDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ConnectionV131Connection{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, ConnectionV131Connection{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1FabricsFabricIdConnectionsConnectionIdDelete method not implemented")
}

// RedfishV1FabricsFabricIdConnectionsConnectionIdGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdConnectionsConnectionIdGet(ctx context.Context, fabricId string, connectionId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1FabricsFabricIdConnectionsConnectionIdGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ConnectionV131Connection{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, ConnectionV131Connection{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1FabricsFabricIdConnectionsConnectionIdGet method not implemented")
}

// RedfishV1FabricsFabricIdConnectionsGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdConnectionsGet(ctx context.Context, fabricId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1FabricsFabricIdConnectionsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ConnectionCollectionConnectionCollection{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, ConnectionCollectionConnectionCollection{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1FabricsFabricIdConnectionsGet method not implemented")
}

// RedfishV1FabricsFabricIdConnectionsPost -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdConnectionsPost(ctx context.Context, fabricId string, connectionV131Connection redfishapi.ConnectionV131Connection) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1FabricsFabricIdConnectionsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, ConnectionV131Connection{}) or use other options such as http.Ok ...
	// return redfishapi.Response(201, ConnectionV131Connection{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1FabricsFabricIdConnectionsPost method not implemented")
}

// RedfishV1FabricsFabricIdEndpointsEndpointIdGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdEndpointsEndpointIdGet(ctx context.Context, fabricId string, endpointId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1FabricsFabricIdEndpointsEndpointIdGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, EndpointV181Endpoint{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, EndpointV181Endpoint{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1FabricsFabricIdEndpointsEndpointIdGet method not implemented")
}

// RedfishV1FabricsFabricIdEndpointsGet -
func (s *CxlHostApiService) RedfishV1FabricsFabricIdEndpointsGet(ctx context.Context, fabricId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1FabricsFabricIdEndpointsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, EndpointCollectionEndpointCollection{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, EndpointCollectionEndpointCollection{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1FabricsFabricIdEndpointsGet method not implemented")
}

// RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdDelete -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdDelete(ctx context.Context, computerSystemId string, memoryDomainId string, memoryChunksId string) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MemoryChunksV161MemoryChunks{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, MemoryChunksV161MemoryChunks{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksMemoryChunksIdDelete method not implemented")
}

// RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksPost -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksPost(ctx context.Context, computerSystemId string, memoryDomainId string, memoryChunksV161MemoryChunks redfishapi.MemoryChunksV161MemoryChunks) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MemoryChunksV161MemoryChunks{}) or use other options such as http.Ok ...
	// return redfishapi.Response(201, MemoryChunksV161MemoryChunks{}), nil

	// TODO: Uncomment the next line to return response Response(0, RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1SystemsComputerSystemIdMemoryDomainsMemoryDomainIdMemoryChunksPost method not implemented")
}
