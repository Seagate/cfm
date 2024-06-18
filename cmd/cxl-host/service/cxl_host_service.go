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
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsManagerAccountIdPatch(ctx context.Context, managerAccountId string, managerAccountV1100ManagerAccount redfishapi.ManagerAccountV1100ManagerAccount) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceAccountsManagerAccountIdPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// Retrieve existing account, or nil indicating it does not exist
	account := accounts.GetAccount(managerAccountId)
	if account == nil {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Account id (%s) does not exist", managerAccountId)
	}

	// Allow the client to change the password for this user.
	_, err := accounts.AccountsHandler().UpdateAccount(account.Username, *managerAccountV1100ManagerAccount.Password, "")
	if err != nil {
		return redfishapi.Response(http.StatusBadRequest, nil), err
	}

	resource := fillAccountResource(managerAccountId, account)
	return redfishapi.Response(http.StatusOK, resource), nil
}

// RedfishV1AccountServiceAccountsPost -
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsPost(ctx context.Context, managerAccountV1100ManagerAccount redfishapi.ManagerAccountV1120ManagerAccount) (redfishapi.ImplResponse, error) {

	// Create a new user account
	account, err := accounts.AccountsHandler().AddAccount(managerAccountV1100ManagerAccount.UserName, *managerAccountV1100ManagerAccount.Password, managerAccountV1100ManagerAccount.RoleId)
	if err != nil {
		return redfishapi.Response(http.StatusBadRequest, nil), err
	}

	resource := fillAccountResource(account.Username, account)
	return redfishapi.Response(http.StatusOK, resource), nil
}

// AccountService - The AccountService schema defines an account service.  The properties are common to, and enable management of, all user accounts.  The properties include the password requirements and control features, such as account lockout.  Properties and actions in this service specify general behavior that should be followed for typical accounts, however implementations may override these behaviors for special accounts or situations to avoid denial of service or other deadlock situations.
type AccountService struct {
	OdataContext                      string                  `json:"@odata.context,omitempty"`                    // The OData description of a payload.
	OdataEtag                         string                  `json:"@odata.etag,omitempty"`                       // The current ETag of the resource.
	OdataId                           string                  `json:"@odata.id"`                                   // The unique identifier for a resource.
	OdataType                         string                  `json:"@odata.type"`                                 // The type of a resource.
	AccountLockoutCounterResetAfter   int64                   `json:"AccountLockoutCounterResetAfter,omitempty"`   // The period of time, in seconds, between the last failed login attempt and the reset of the lockout threshold counter.  This value must be less than or equal to the AccountLockoutDuration value.  A reset sets the counter to `0`.
	AccountLockoutCounterResetEnabled bool                    `json:"AccountLockoutCounterResetEnabled,omitempty"` // An indication of whether the threshold counter is reset after AccountLockoutCounterResetAfter expires.  If `true`, it is reset.  If `false`, only a successful login resets the threshold counter and if the user reaches the AccountLockoutThreshold limit, the account will be locked out indefinitely and only an administrator-issued reset clears the threshold counter.  If this property is absent, the default is `true`.
	AccountLockoutDuration            *int64                  `json:"AccountLockoutDuration,omitempty"`            // The period of time, in seconds, that an account is locked after the number of failed login attempts reaches the account lockout threshold, within the period between the last failed login attempt and the reset of the lockout threshold counter.  If this value is `0`, no lockout will occur.  If the AccountLockoutCounterResetEnabled value is `false`, this property is ignored.
	AccountLockoutThreshold           *int64                  `json:"AccountLockoutThreshold,omitempty"`           // The number of allowed failed login attempts before a user account is locked for a specified duration.  If `0`, the account is never locked.
	Accounts                          redfishapi.OdataV4IdRef `json:"Accounts,omitempty"`
	//	Actions                            AccountServiceV1130Actions                 `json:"Actions,omitempty"`
	//	ActiveDirectory                    AccountServiceV1130ExternalAccountProvider `json:"ActiveDirectory,omitempty"`
	//	AdditionalExternalAccountProviders OdataV4IdRef                               `json:"AdditionalExternalAccountProviders,omitempty"`
	AuthFailureLoggingThreshold int64  `json:"AuthFailureLoggingThreshold,omitempty"` // The number of authorization failures per account that are allowed before the failed attempt is logged to the manager log.
	Description                 string `json:"Description,omitempty"`                 // The description of this resource.  Used for commonality in the schema definitions.
	Id                          string `json:"Id"`                                    // The unique identifier for this resource within the collection of similar resources.
	//	LDAP                               AccountServiceV1130ExternalAccountProvider `json:"LDAP,omitempty"`
	LocalAccountAuth  redfishapi.AccountServiceV1130LocalAccountAuth `json:"LocalAccountAuth,omitempty"`
	MaxPasswordLength int64                                          `json:"MaxPasswordLength,omitempty"` // The maximum password length for this account service.
	MinPasswordLength int64                                          `json:"MinPasswordLength,omitempty"` // The minimum password length for this account service.
	Name              string                                         `json:"Name"`                        // The name of the resource or array member.
	//	OAuth2                             AccountServiceV1130ExternalAccountProvider `json:"OAuth2,omitempty"`
	Oem                    map[string]interface{} `json:"Oem,omitempty"`                    // The OEM extension.
	PasswordExpirationDays *int64                 `json:"PasswordExpirationDays,omitempty"` // The number of days before account passwords in this account service will expire.
	//	PrivilegeMap                       OdataV4IdRef                               `json:"PrivilegeMap,omitempty"`
	RestrictedOemPrivileges  []string                                `json:"RestrictedOemPrivileges,omitempty"` // The set of restricted OEM privileges.
	RestrictedPrivileges     []redfishapi.PrivilegesPrivilegeType    `json:"RestrictedPrivileges,omitempty"`    // The set of restricted Redfish privileges.
	Roles                    redfishapi.OdataV4IdRef                 `json:"Roles,omitempty"`
	ServiceEnabled           *bool                                   `json:"ServiceEnabled,omitempty"` // An indication of whether the account service is enabled.  If `true`, it is enabled.  If `false`, it is disabled and users cannot be created, deleted, or modified, and new sessions cannot be started.  However, established sessions might still continue to run.  Any service, such as the session service, that attempts to access the disabled account service fails.  However, this does not affect HTTP Basic Authentication connections.
	Status                   redfishapi.ResourceStatus               `json:"Status,omitempty"`
	SupportedAccountTypes    []redfishapi.ManagerAccountAccountTypes `json:"SupportedAccountTypes,omitempty"`    // The account types supported by the service.
	SupportedOEMAccountTypes []string                                `json:"SupportedOEMAccountTypes,omitempty"` // The OEM account types supported by the service.
	// TACACSplus                         AccountServiceV1130ExternalAccountProvider `json:"TACACSplus,omitempty"`
}

// RedfishV1AccountServiceGet -
func (s *CxlHostApiService) RedfishV1AccountServiceGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	enabled := true

	resource := AccountService{
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
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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

// ChassisType - The Chassis schema represents the physical components of a system.  This resource represents the bare minimum supported by this service.
type ChassisType struct {
	OdataContext string `json:"@odata.context,omitempty"` // The OData description of a payload.
	OdataEtag    string `json:"@odata.etag,omitempty"`    // The current ETag of the resource.
	OdataId      string `json:"@odata.id"`                // The unique identifier for a resource.
	OdataType    string `json:"@odata.type"`              // The type of a resource.
	// Actions ChassisV1210Actions `json:"Actions,omitempty"`
	// Assembly OdataV4IdRef `json:"Assembly,omitempty"`
	AssetTag *string `json:"AssetTag,omitempty"` // The user-assigned asset tag of this chassis.
	// Certificates OdataV4IdRef `json:"Certificates,omitempty"`
	ChassisType redfishapi.ChassisV1230ChassisType `json:"ChassisType"`
	// Controls OdataV4IdRef `json:"Controls,omitempty"`
	// DepthMm *float32 `json:"DepthMm,omitempty"`// The depth of the chassis.
	Description string `json:"Description,omitempty"` // The description of this resource.  Used for commonality in the schema definitions.
	// Drives OdataV4IdRef `json:"Drives,omitempty"`
	// ElectricalSourceManagerURIs []*string `json:"ElectricalSourceManagerURIs,omitempty"`// The URIs of the management interfaces for the upstream electrical source connections for this chassis.
	// ElectricalSourceNames []*string `json:"ElectricalSourceNames,omitempty"`// The names of the upstream electrical sources, such as circuits or outlets, connected to this chassis.
	// EnvironmentMetrics OdataV4IdRef `json:"EnvironmentMetrics,omitempty"`
	// EnvironmentalClass ChassisV1210EnvironmentalClass `json:"EnvironmentalClass,omitempty"`
	// FabricAdapters OdataV4IdRef `json:"FabricAdapters,omitempty"`
	// HeightMm *float32 `json:"HeightMm,omitempty"` // The height of the chassis.
	// HotPluggable *bool `json:"HotPluggable,omitempty"`// An indication of whether this component can be inserted or removed while the equipment is in operation.
	Id string `json:"Id"` // The unique identifier for this resource within the collection of similar resources.
	// IndicatorLED ChassisV1210IndicatorLed `json:"IndicatorLED,omitempty"`
	// Links ChassisV1210Links `json:"Links,omitempty"`
	// Location ResourceLocation `json:"Location,omitempty"`
	// LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`// An indicator allowing an operator to physically locate this resource.
	// LogServices OdataV4IdRef `json:"LogServices,omitempty"`
	// Manufacturer *string `json:"Manufacturer,omitempty"`// The manufacturer of this chassis.
	// MaxPowerWatts *float32 `json:"MaxPowerWatts,omitempty"`// The upper bound of the total power consumed by the chassis.
	// Measurements []SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`// An array of DSP0274-defined measurement blocks. (Deprecated)
	// MediaControllers OdataV4IdRef `json:"MediaControllers,omitempty"`
	Memory        redfishapi.OdataV4IdRef `json:"Memory,omitempty"`
	MemoryDomains redfishapi.OdataV4IdRef `json:"MemoryDomains,omitempty"`
	// MinPowerWatts *float32 `json:"MinPowerWatts,omitempty"`// The lower bound of the total power consumed by the chassis.
	// Model *string `json:"Model,omitempty"`// The model number of the chassis.
	Name string `json:"Name"` // The name of the resource or array member.
	// NetworkAdapters OdataV4IdRef `json:"NetworkAdapters,omitempty"`
	Oem         map[string]interface{}  `json:"Oem,omitempty"` // The OEM extension.
	PCIeDevices redfishapi.OdataV4IdRef `json:"PCIeDevices,omitempty"`
	PCIeSlots   redfishapi.OdataV4IdRef `json:"PCIeSlots,omitempty"`  // PCIeDevices OdataV4IdRef `json:"PCIeDevices,omitempty"`
	PartNumber  *string                 `json:"PartNumber,omitempty"` // The part number of the chassis.
	// PhysicalSecurity ChassisV1210PhysicalSecurity `json:"PhysicalSecurity,omitempty"`
	// Power OdataV4IdRef `json:"Power,omitempty"`
	// PowerState ChassisV1210PowerState `json:"PowerState,omitempty"`
	PowerSubsystem redfishapi.OdataV4IdRef `json:"PowerSubsystem,omitempty"`
	// PoweredByParent *bool `json:"PoweredByParent,omitempty"`// Indicates that the chassis receives power from the containing chassis.
	// Replaceable *bool `json:"Replaceable,omitempty"`// An indication of whether this component can be independently replaced as allowed by the vendor's replacement policy.
	// SKU *string `json:"SKU,omitempty"`// The SKU of the chassis.
	// Sensors OdataV4IdRef `json:"Sensors,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty"` // The serial number of the chassis.
	// SparePartNumber *string `json:"SparePartNumber,omitempty"`// The spare part number of the chassis.
	Status redfishapi.ResourceStatus `json:"Status,omitempty"`
	// Thermal OdataV4IdRef `json:"Thermal,omitempty"`
	// ThermalDirection ChassisV1210ThermalDirection `json:"ThermalDirection,omitempty"`
	// ThermalManagedByParent *bool `json:"ThermalManagedByParent,omitempty"` // Indicates that the chassis is thermally managed by the parent chassis.
	// ThermalSubsystem OdataV4IdRef `json:"ThermalSubsystem,omitempty"`
	// TrustedComponents OdataV4IdRef `json:"TrustedComponents,omitempty"`
	UUID    string  `json:"UUID,omitempty"`
	Version *string `json:"Version,omitempty"` // The hardware version of this chassis.
	// WeightKg *float32 `json:"WeightKg,omitempty"` // The weight of the chassis.
	//WidthMm *float32 `json:"WidthMm,omitempty"`	// The width of the chassis.
}

// SwitchType - The Switch schema contains properties that describe a fabric switch.
type SwitchType struct {
	OdataContext string `json:"@odata.context,omitempty"` // The OData description of a payload.
	OdataEtag    string `json:"@odata.etag,omitempty"`    // The current ETag of the resource.
	OdataId      string `json:"@odata.id"`                // The unique identifier for a resource.
	OdataType    string `json:"@odata.type"`              // The type of a resource.
	//Actions                 SwitchV180Actions                   `json:"Actions,omitempty"`                 //
	AssetTag *string `json:"AssetTag,omitempty"` // The user-assigned asset tag for this switch.
	//Certificates            OdataV4IdRef                        `json:"Certificates,omitempty"`            //
	CurrentBandwidthGbps *float32 `json:"CurrentBandwidthGbps,omitempty"` // The current internal bandwidth of this switch.
	Description          string   `json:"Description,omitempty"`          // The description of this resource.  Used for commonality in the schema definitions.
	DomainID             *int64   `json:"DomainID,omitempty"`             // The domain ID for this switch.
	Enabled              bool     `json:"Enabled,omitempty"`              // An indication of whether this switch is enabled.
	//EnvironmentMetrics      OdataV4IdRef                        `json:"EnvironmentMetrics,omitempty"`      //
	FirmwareVersion *string                         `json:"FirmwareVersion,omitempty"` // The firmware version of this switch.
	Id              string                          `json:"Id"`                        // The unique identifier for this resource within the collection of similar resources.
	IndicatorLED    redfishapi.ResourceIndicatorLed `json:"IndicatorLED,omitempty"`    //
	IsManaged       *bool                           `json:"IsManaged,omitempty"`       // An indication of whether the switch is in a managed or unmanaged state.
	//Links                   SwitchV180Links                     `json:"Links,omitempty"`                   //
	//Location                ResourceLocation                    `json:"Location,omitempty"`                //
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"` // An indicator allowing an operator to physically locate this resource.
	//LogServices             OdataV4IdRef                        `json:"LogServices,omitempty"`             //
	Manufacturer     *string                                        `json:"Manufacturer,omitempty"`     // The manufacturer of this switch.
	MaxBandwidthGbps *float32                                       `json:"MaxBandwidthGbps,omitempty"` // The maximum internal bandwidth of this switch as currently configured.
	Measurements     []redfishapi.SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`     // An array of DSP0274-defined measurement blocks. (Deprecated)
	//Metrics                 OdataV4IdRef                        `json:"Metrics,omitempty"`                 //
	Model                *string                           `json:"Model,omitempty"`                  // The product model number of this switch.
	Name                 string                            `json:"Name"`                             // The name of the resource or array member.
	Oem                  map[string]interface{}            `json:"Oem,omitempty"`                    // The OEM extension.
	PartNumber           *string                           `json:"PartNumber,omitempty"`             // The part number for this switch.
	Ports                redfishapi.OdataV4IdRef           `json:"Ports,omitempty"`                  //
	PowerState           redfishapi.ResourcePowerState     `json:"PowerState,omitempty"`             //
	Redundancy           []redfishapi.RedundancyRedundancy `json:"Redundancy,omitempty"`             // Redundancy information for the switches.
	RedundancyodataCount int64                             `json:"Redundancy@odata.count,omitempty"` // The number of items in a collection.
	SKU                  *string                           `json:"SKU,omitempty"`                    // The SKU for this switch.
	SerialNumber         *string                           `json:"SerialNumber,omitempty"`           // The serial number for this switch.
	Status               redfishapi.ResourceStatus         `json:"Status,omitempty"`                 //
	SupportedProtocols   []redfishapi.ProtocolProtocol     `json:"SupportedProtocols,omitempty"`     // The protocols this switch supports.
	SwitchType           redfishapi.ProtocolProtocol       `json:"SwitchType,omitempty"`             //
	TotalSwitchWidth     *int64                            `json:"TotalSwitchWidth,omitempty"`       // The total number of lanes, phys, or other physical transport links that this switch contains.
	UUID                 string                            `json:"UUID,omitempty"`                   //
}

// RedfishV1ChassisChassisIdGet -
func (s *CxlHostApiService) RedfishV1ChassisChassisIdGet(ctx context.Context, chassisId string) (redfishapi.ImplResponse, error) {

	if chassisId != GetChassisName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("Chassis id (%s) does not exist", chassisId)
	}

	tag := GetChassisTag()
	uuid := GetChassisUUID()
	version := GetChassisVersion()

	chassis := ChassisType{

		//	Actions ChassisV1210Actions `json:"Actions,omitempty"`
		//	Assembly OdataV4IdRef `json:"Assembly,omitempty"`
		//	AssetTag *string `json:"AssetTag,omitempty"`
		//	Certificates OdataV4IdRef `json:"Certificates,omitempty"`
		//	ChassisType ChassisV1210ChassisType `json:"ChassisType"`
		//	Controls OdataV4IdRef `json:"Controls,omitempty"`
		//	Description string `json:"Description,omitempty"`
		//	Drives OdataV4IdRef `json:"Drives,omitempty"`
		//	EnvironmentMetrics OdataV4IdRef `json:"EnvironmentMetrics,omitempty"`
		//	FabricAdapters OdataV4IdRef `json:"FabricAdapters,omitempty"`
		//	Id string `json:"Id"`
		//	Links ChassisV1210Links `json:"Links,omitempty"`
		//	Location ResourceLocation `json:"Location,omitempty"`
		//	LogServices OdataV4IdRef `json:"LogServices,omitempty"`
		//	MediaControllers OdataV4IdRef `json:"MediaControllers,omitempty"`
		//	Memory OdataV4IdRef `json:"Memory,omitempty"`
		//	MemoryDomains OdataV4IdRef `json:"MemoryDomains,omitempty"`
		//	Name string `json:"Name"`
		//	NetworkAdapters OdataV4IdRef `json:"NetworkAdapters,omitempty"`
		//	Oem map[string]interface{} `json:"Oem,omitempty"`
		//	PCIeDevices OdataV4IdRef `json:"PCIeDevices,omitempty"`
		//	PCIeSlots OdataV4IdRef `json:"PCIeSlots,omitempty"`
		//	PhysicalSecurity ChassisV1210PhysicalSecurity `json:"PhysicalSecurity,omitempty"`
		//	Power OdataV4IdRef `json:"Power,omitempty"`
		//	PowerSubsystem OdataV4IdRef `json:"PowerSubsystem,omitempty"`
		//	Sensors OdataV4IdRef `json:"Sensors,omitempty"`
		//	SerialNumber *string `json:"SerialNumber,omitempty"`
		//	Status ResourceStatus `json:"Status,omitempty"`
		//	Thermal OdataV4IdRef `json:"Thermal,omitempty"`
		//	ThermalSubsystem OdataV4IdRef `json:"ThermalSubsystem,omitempty"`
		//	TrustedComponents OdataV4IdRef `json:"TrustedComponents,omitempty"`
		//	UUID string `json:"UUID,omitempty"`
		//	Version *string `json:"Version,omitempty"`

		OdataContext: "/redfish/v1/$metadata#Chassis.Chassis",
		OdataId:      "/redfish/v1/Chassis/" + GetChassisName(),
		OdataType:    ChassisVersion,
		AssetTag:     &tag,
		ChassisType:  redfishapi.CHASSISV1230CHASSISTYPE_RACK_MOUNT,
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
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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

// MemoryType - The Memory schema represents a memory device, such as a DIMM, and its configuration.
type MemoryType struct {
	OdataContext string `json:"@odata.context,omitempty"` // The OData description of a payload.
	OdataEtag    string `json:"@odata.etag,omitempty"`    // The current ETag of the resource.
	OdataId      string `json:"@odata.id"`                // The unique identifier for a resource.
	OdataType    string `json:"@odata.type"`              // The type of a resource.
	// Actions                                 MemoryV1171Actions                  `json:"Actions,omitempty"`                                 //
	AllocationAlignmentMiB *int64  `json:"AllocationAlignmentMiB,omitempty"` // The boundary that memory regions are allocated on, measured in mebibytes (MiB).
	AllocationIncrementMiB *int64  `json:"AllocationIncrementMiB,omitempty"` // The size of the smallest unit of allocation for a memory region in mebibytes (MiB).
	AllowedSpeedsMHz       []int64 `json:"AllowedSpeedsMHz,omitempty"`       // Speeds supported by this memory device.
	// Assembly                                OdataV4IdRef                        `json:"Assembly,omitempty"`                                //
	BaseModuleType redfishapi.MemoryV1171BaseModuleType `json:"BaseModuleType,omitempty"` //
	BusWidthBits   *int64                               `json:"BusWidthBits,omitempty"`   // The bus width, in bits.
	CacheSizeMiB   *int64                               `json:"CacheSizeMiB,omitempty"`   // Total size of the cache portion memory in MiB.
	CapacityMiB    *int64                               `json:"CapacityMiB,omitempty"`    // Memory capacity in mebibytes (MiB).
	// Certificates                            OdataV4IdRef                        `json:"Certificates,omitempty"`                            //
	ConfigurationLocked *bool   `json:"ConfigurationLocked,omitempty"` // An indication of whether the configuration of this memory device is locked and cannot be altered.
	DataWidthBits       *int64  `json:"DataWidthBits,omitempty"`       // Data width in bits.
	Description         string  `json:"Description,omitempty"`         // The description of this resource.  Used for commonality in the schema definitions.
	DeviceID            *string `json:"DeviceID,omitempty"`            // Device ID. (Deprectaed)
	DeviceLocator       *string `json:"DeviceLocator,omitempty"`       // Location of the memory device in the platform. (Deprecated)
	Enabled             bool    `json:"Enabled,omitempty"`             // An indication of whether this memory is enabled.
	// EnvironmentMetrics                      OdataV4IdRef                        `json:"EnvironmentMetrics,omitempty"`                      //
	ErrorCorrection      redfishapi.MemoryV1171ErrorCorrection `json:"ErrorCorrection,omitempty"`      //
	FirmwareApiVersion   *string                               `json:"FirmwareApiVersion,omitempty"`   // Version of API supported by the firmware.
	FirmwareRevision     *string                               `json:"FirmwareRevision,omitempty"`     // Revision of firmware on the memory controller.
	FunctionClasses      []string                              `json:"FunctionClasses,omitempty"`      // Function classes by the memory device. (Deprecated)
	Id                   string                                `json:"Id"`                             // The unique identifier for this resource within the collection of similar resources.
	IsRankSpareEnabled   *bool                                 `json:"IsRankSpareEnabled,omitempty"`   // An indication of whether rank spare is enabled for this memory device.
	IsSpareDeviceEnabled *bool                                 `json:"IsSpareDeviceEnabled,omitempty"` // An indication of whether a spare device is enabled for this memory device.
	// Links                                   MemoryV1171Links                    `json:"Links,omitempty"`                                   //
	// Location                                ResourceLocation                    `json:"Location,omitempty"`                                //
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"` // An indicator allowing an operator to physically locate this resource.
	// Log                                     OdataV4IdRef                        `json:"Log,omitempty"`                                     //
	LogicalSizeMiB   *int64                                         `json:"LogicalSizeMiB,omitempty"`   // Total size of the logical memory in MiB.
	Manufacturer     *string                                        `json:"Manufacturer,omitempty"`     // The memory device manufacturer.
	MaxTDPMilliWatts []int64                                        `json:"MaxTDPMilliWatts,omitempty"` // Set of maximum power budgets supported by the memory device in milliwatts.
	Measurements     []redfishapi.SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`     // An array of DSP0274-defined measurement blocks. (Deprecated)
	MemoryDeviceType redfishapi.MemoryV1171MemoryDeviceType         `json:"MemoryDeviceType,omitempty"` //
	// MemoryLocation                          MemoryV1171MemoryLocation           `json:"MemoryLocation,omitempty"`                          //
	MemoryMedia                             []redfishapi.MemoryV1171MemoryMedia `json:"MemoryMedia,omitempty"`                             // Media of this memory device.
	MemorySubsystemControllerManufacturerID *string                             `json:"MemorySubsystemControllerManufacturerID,omitempty"` // The manufacturer ID of the memory subsystem controller of this memory device.
	MemorySubsystemControllerProductID      *string                             `json:"MemorySubsystemControllerProductID,omitempty"`      // The product ID of the memory subsystem controller of this memory device.
	MemoryType                              redfishapi.MemoryV1171MemoryType    `json:"MemoryType,omitempty"`                              //
	// Metrics                                 OdataV4IdRef                        `json:"Metrics,omitempty"`                                 //
	Model                *string                `json:"Model,omitempty"`                // The product model number of this device.
	ModuleManufacturerID *string                `json:"ModuleManufacturerID,omitempty"` // The manufacturer ID of this memory device.
	ModuleProductID      *string                `json:"ModuleProductID,omitempty"`      // The product ID of this memory device.
	Name                 string                 `json:"Name"`                           // The name of the resource or array member.
	NonVolatileSizeMiB   *int64                 `json:"NonVolatileSizeMiB,omitempty"`   // Total size of the non-volatile portion memory in MiB.
	Oem                  map[string]interface{} `json:"Oem,omitempty"`                  // The OEM extension.
	// OperatingMemoryModes                    []MemoryV1171OperatingMemoryModes   `json:"OperatingMemoryModes,omitempty"`                    // Memory modes supported by the memory device.
	OperatingSpeedMhz *int64 `json:"OperatingSpeedMhz,omitempty"` // Operating speed of the memory device in MHz or MT/s as appropriate.
	// OperatingSpeedRangeMHz       ControlControlRangeExcerpt `json:"OperatingSpeedRangeMHz,omitempty"`       //
	PartNumber                   *string `json:"PartNumber,omitempty"`                   // The product part number of this device.
	PersistentRegionNumberLimit  *int64  `json:"PersistentRegionNumberLimit,omitempty"`  // Total number of persistent regions this memory device can support.
	PersistentRegionSizeLimitMiB *int64  `json:"PersistentRegionSizeLimitMiB,omitempty"` // Total size of persistent regions in mebibytes (MiB).
	PersistentRegionSizeMaxMiB   *int64  `json:"PersistentRegionSizeMaxMiB,omitempty"`   // Maximum size of a single persistent region in mebibytes (MiB).
	// PowerManagementPolicy                   MemoryV1171PowerManagementPolicy    `json:"PowerManagementPolicy,omitempty"`                   //
	RankCount *int64                            `json:"RankCount,omitempty"` // Number of ranks available in the memory device.
	Regions   []redfishapi.MemoryV1171RegionSet `json:"Regions,omitempty"`   // Memory regions information within the memory device.
	// SecurityCapabilities                    MemoryV1171SecurityCapabilities     `json:"SecurityCapabilities,omitempty"`                    //
	SecurityState              redfishapi.MemoryV1171SecurityStates `json:"SecurityState,omitempty"`              //
	SerialNumber               *string                              `json:"SerialNumber,omitempty"`               // The product serial number of this device.
	SpareDeviceCount           *int64                               `json:"SpareDeviceCount,omitempty"`           // Number of unused spare devices available in the memory device.
	SparePartNumber            *string                              `json:"SparePartNumber,omitempty"`            // The spare part number of the memory.
	Status                     redfishapi.ResourceStatus            `json:"Status,omitempty"`                     //
	SubsystemDeviceID          *string                              `json:"SubsystemDeviceID,omitempty"`          // Subsystem device ID. (Deprecated)
	SubsystemVendorID          *string                              `json:"SubsystemVendorID,omitempty"`          // SubSystem vendor ID. (Deprecated)
	VendorID                   *string                              `json:"VendorID,omitempty"`                   // Vendor ID. (Deprecated)
	VolatileRegionNumberLimit  *int64                               `json:"VolatileRegionNumberLimit,omitempty"`  // Total number of volatile regions this memory device can support.
	VolatileRegionSizeLimitMiB *int64                               `json:"VolatileRegionSizeLimitMiB,omitempty"` // Total size of volatile regions in mebibytes (MiB).
	VolatileRegionSizeMaxMiB   *int64                               `json:"VolatileRegionSizeMaxMiB,omitempty"`   // Maximum size of a single volatile region in mebibytes (MiB).
	VolatileSizeMiB            *int64                               `json:"VolatileSizeMiB,omitempty"`            // Total size of the volatile portion memory in MiB.
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

	memory := MemoryType{
		OdataContext: "/redfish/v1/$metadata#Memory.Memory",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/Memory/" + memoryId,
		OdataType:    MemoryVersion,
		CapacityMiB:  &total,
		Id:           memoryId,
		Enabled:      true,
		Name:         memoryId,
		Status: redfishapi.ResourceStatus{
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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

	fabric := redfishapi.FabricV130Fabric{
		OdataContext: "/redfish/v1/$metadata#Fabric.Fabric",
		OdataId:      "/redfish/v1/Fabrics/" + fabricId,
		OdataType:    FabricVersion,
		//Actions FabricV130Actions `json:"Actions,omitempty"`
		//AddressPools OdataV4IdRef `json:"AddressPools,omitempty"`
		//Connections OdataV4IdRef `json:"Connections,omitempty"`
		Description: fabricId + " Fabric",
		//EndpointGroups OdataV4IdRef `json:"EndpointGroups,omitempty"`
		//Endpoints OdataV4IdRef `json:"Endpoints,omitempty"`
		FabricType: redfishapi.PROTOCOLPROTOCOL_CXL,
		Id:         fabricId,
		//Links FabricV130Links `json:"Links,omitempty"`
		MaxZones: &zones,
		Name:     fabricId + " Fabric",
		Oem:      nil,
		Status: redfishapi.ResourceStatus{
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
		},
		Switches: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Fabrics/" + fabricId + "/Switches",
		},
		UUID: GetFabricUUID(fabricId),
		//Zones OdataV4IdRef `json:"Zones,omitempty"`
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

	resource := SwitchType{
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
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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

// PortType - The Port schema contains properties that describe a port of a switch, controller, chassis, or any other device that could be connected to another entity.
type PortType struct {
	OdataContext string `json:"@odata.context,omitempty"` // The OData description of a payload.
	OdataEtag    string `json:"@odata.etag,omitempty"`    // The current ETag of the resource.
	OdataId      string `json:"@odata.id"`                // The unique identifier for a resource.
	OdataType    string `json:"@odata.type"`              // The type of a resource.
	//Actions                 PortV190Actions                `json:"Actions,omitempty"`                 //
	ActiveWidth             int64     `json:"ActiveWidth,omitempty"`             // The number of active lanes for this interface.
	CapableProtocolVersions []*string `json:"CapableProtocolVersions,omitempty"` // The protocol versions capable of being sent over this port.
	CurrentProtocolVersion  *string   `json:"CurrentProtocolVersion,omitempty"`  // The protocol version being sent over this port.
	CurrentSpeedGbps        *float32  `json:"CurrentSpeedGbps,omitempty"`        // The current speed of this port.
	Description             string    `json:"Description,omitempty"`             // The description of this resource.  Used for commonality in the schema definitions.
	Enabled                 bool      `json:"Enabled,omitempty"`                 // An indication of whether this port is enabled.
	//EnvironmentMetrics      OdataV4IdRef                   `json:"EnvironmentMetrics,omitempty"`      //
	//Ethernet                PortV190EthernetProperties     `json:"Ethernet,omitempty"`                //
	//FibreChannel            PortV190FibreChannelProperties `json:"FibreChannel,omitempty"`            //
	FunctionMaxBandwidth []redfishapi.PortV190FunctionMaxBandwidth `json:"FunctionMaxBandwidth,omitempty"` // An array of maximum bandwidth allocation percentages for the functions associated with this port.
	FunctionMinBandwidth []redfishapi.PortV190FunctionMinBandwidth `json:"FunctionMinBandwidth,omitempty"` // An array of minimum bandwidth allocation percentages for the functions associated with this port.
	//GenZ                    PortV190GenZ                   `json:"GenZ,omitempty"`                    //
	Id string `json:"Id"` // The unique identifier for this resource within the collection of similar resources.
	//InfiniBand              PortV190InfiniBandProperties   `json:"InfiniBand,omitempty"`              //
	InterfaceEnabled        *bool                                    `json:"InterfaceEnabled,omitempty"`        // An indication of whether the interface is enabled.
	LinkConfiguration       []redfishapi.PortV190LinkConfiguration   `json:"LinkConfiguration,omitempty"`       // The link configuration of this port.
	LinkNetworkTechnology   redfishapi.PortV190LinkNetworkTechnology `json:"LinkNetworkTechnology,omitempty"`   //
	LinkState               redfishapi.PortV190LinkState             `json:"LinkState,omitempty"`               //
	LinkStatus              redfishapi.PortV190LinkStatus            `json:"LinkStatus,omitempty"`              //
	LinkTransitionIndicator int64                                    `json:"LinkTransitionIndicator,omitempty"` // The number of link state transitions for this interface.
	Links                   redfishapi.PortV190Links                 `json:"Links,omitempty"`                   //
	//Location                ResourceLocation               `json:"Location,omitempty"`                //
	LocationIndicatorActive *bool    `json:"LocationIndicatorActive,omitempty"` // An indicator allowing an operator to physically locate this resource.
	MaxFrameSize            *int64   `json:"MaxFrameSize,omitempty"`            // The maximum frame size supported by the port.
	MaxSpeedGbps            *float32 `json:"MaxSpeedGbps,omitempty"`            // The maximum speed of this port as currently configured.
	//Metrics                 OdataV4IdRef                   `json:"Metrics,omitempty"`                 //
	Name         string                        `json:"Name"`                   // The name of the resource or array member.
	Oem          map[string]interface{}        `json:"Oem,omitempty"`          // The OEM extension.
	PortId       *string                       `json:"PortId,omitempty"`       // The label of this port on the physical package for this port.
	PortMedium   redfishapi.PortV190PortMedium `json:"PortMedium,omitempty"`   //
	PortProtocol redfishapi.ProtocolProtocol   `json:"PortProtocol,omitempty"` //
	PortType     redfishapi.PortV190PortType   `json:"PortType,omitempty"`     //
	//SFP                     PortV190Sfp                    `json:"SFP,omitempty"`                     //
	SignalDetected *bool                     `json:"SignalDetected,omitempty"` // An indication of whether a signal is detected on this interface.
	Status         redfishapi.ResourceStatus `json:"Status,omitempty"`         //
	Width          *int64                    `json:"Width,omitempty"`          // The number of lanes, phys, or other physical transport links that this port contains.
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

	resource := PortType{
		OdataContext:     "/redfish/v1/$metadata#Chassis.Port",
		OdataId:          "/redfish/v1/Fabrics/" + fabricId + "/Switches/" + switchId + "/Ports/" + portId,
		OdataType:        PortVersion,
		Description:      fabricId + " " + portId,
		Enabled:          true,
		Id:               portId,
		InterfaceEnabled: &enabled,
		LinkState:        redfishapi.PORTV190LINKSTATE_ENABLED,
		LinkStatus:       redfishapi.PORTV190LINKSTATUS_LINK_UP,
		Name:             portId,
		Oem:              nil,
		PortProtocol:     redfishapi.PROTOCOLPROTOCOL_CXL,
		PortType:         redfishapi.PORTV190PORTTYPE_BIDIRECTIONAL_PORT,
		SignalDetected:   &detected,
		Status: redfishapi.ResourceStatus{
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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

// ServiceRoot - The ServiceRoot schema describes the root of the Redfish Service, located at the '/redfish/v1' URI.
type ServiceRoot struct {
	OdataContext   string                  `json:"@odata.context,omitempty"` // The OData description of a payload.
	OdataId        string                  `json:"@odata.id"`                // The unique identifier for a resource.
	OdataType      string                  `json:"@odata.type"`              // The type of a resource.
	AccountService redfishapi.OdataV4IdRef `json:"AccountService,omitempty"`
	Chassis        redfishapi.OdataV4IdRef `json:"Chassis,omitempty"`     // Chassis support
	Description    string                  `json:"Description,omitempty"` // The description of this resource.  Used for commonality in the schema definitions.
	//Fabrics               redfishapi.OdataV4IdRef          `json:"Fabrics,omitempty"`               // Fabrics support
	Id                    string                           `json:"Id"`                              // The unique identifier for this resource within the collection of similar resources.
	Links                 redfishapi.ServiceRootV1160Links `json:"Links"`                           // Links
	Name                  string                           `json:"Name"`                            // The name of the resource or array member.
	Oem                   map[string]interface{}           `json:"Oem,omitempty"`                   // The OEM extension.
	Product               *string                          `json:"Product,omitempty"`               // The product associated with this Redfish Service.
	RedfishVersion        string                           `json:"RedfishVersion,omitempty"`        // The version of the Redfish Service.
	ServiceIdentification string                           `json:"ServiceIdentification,omitempty"` // The vendor or user-provided product and service identifier.
	SessionService        redfishapi.OdataV4IdRef          `json:"SessionService,omitempty"`        // SessionService
	Systems               redfishapi.OdataV4IdRef          `json:"Systems,omitempty"`               // System sufpport
	UUID                  string                           `json:"UUID,omitempty"`                  // A UUID for this service
	Vendor                *string                          `json:"Vendor,omitempty"`                // The vendor or manufacturer associated with this Redfish Service.
}

// RedfishV1Get -
func (s *CxlHostApiService) RedfishV1Get(ctx context.Context) (redfishapi.ImplResponse, error) {
	vendor := "Seagate"
	product := "CXL Host Redfish Service"

	oem := map[string]interface{}{}
	oem["ServiceVersion"] = common.Version

	root := ServiceRoot{
		OdataContext:   "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
		OdataId:        "/redfish/v1",
		OdataType:      ServiceRootVersion,
		AccountService: redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/AccountService"},
		Chassis:        redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Chassis"},
		Description:    "CXL Host ServiceRoot",
		// Fabrics:               redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Fabrics"},
		Id:                    "CXL Host ServiceRoot",
		Links:                 redfishapi.ServiceRootV1160Links{Oem: nil, Sessions: redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/SessionService/Sessions"}},
		Name:                  "CXL Host",
		Oem:                   oem,
		Product:               &product,
		RedfishVersion:        "1.16.0",
		ServiceIdentification: "",
		SessionService:        redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/SessionService"},
		Systems:               redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/Systems"},
		UUID:                  GetServiceUUID(),
		Vendor:                &vendor,
	}

	return redfishapi.Response(http.StatusOK, root), nil
}

// RedfishV1SessionServiceGet -
func (s *CxlHostApiService) RedfishV1SessionServiceGet(ctx context.Context) (redfishapi.ImplResponse, error) {

	enabled := true

	resource := redfishapi.SessionServiceV118SessionService{
		OdataContext: "/redfish/v1/$metadata#SessionService.SessionService",
		OdataId:      "/redfish/v1/SessionService",
		OdataType:    SessionServiceVersion,
		//Actions SessionServiceV118Actions `json:"Actions,omitempty"`
		Description:    "Session Service",
		Id:             "Session Service",
		Name:           "Session Service",
		Oem:            nil,
		ServiceEnabled: &enabled,
		SessionTimeout: int64(SessionTimeout),
		Sessions:       redfishapi.OdataV4IdRef{OdataId: "/redfish/v1/SessionService/Sessions"},
		Status: redfishapi.ResourceStatus{
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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
func FillInSessionResource(sessionId string, session *accounts.SessionInformation, fillPassword bool) redfishapi.SessionV160Session {

	resource := redfishapi.SessionV160Session{
		OdataId:     "/redfish/v1/SessionService/Sessions/" + sessionId,
		OdataType:   SessionVersion,
		Description: "User Session",
		Id:          sessionId,
		Name:        "User Session " + sessionId,
		Password:    nil,
		SessionType: redfishapi.SESSIONV160SESSIONTYPES_REDFISH,
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
func (s *CxlHostApiService) RedfishV1SessionServiceSessionsPost(ctx context.Context, sessionV160Session redfishapi.SessionV160Session) (redfishapi.ImplResponse, error) {

	// Validate that a username and password were supplied
	if *sessionV160Session.UserName == "" {
		return redfishapi.Response(http.StatusBadRequest, nil), errors.New("UserName is required for creating a new session")
	}

	if *sessionV160Session.Password == "" {
		return redfishapi.Response(http.StatusBadRequest, nil), errors.New("Password is required for creating a new session")
	}

	// Create the session
	session := accounts.CreateSession(ctx, *sessionV160Session.UserName, *sessionV160Session.Password)
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

	resource := redfishapi.MemoryChunksV150MemoryChunks{
		OdataId:          "/redfish/v1/Chassis/" + chassisId + "/MemoryDomains/" + memoryDomainId + "/MemoryChunks/" + memoryChunksId,
		OdataType:        MemoryChunkVersion,
		Name:             "Memory Chunk",
		Id:               memoryChunksId,
		AddressRangeType: redfishapi.MEMORYCHUNKSV150ADDRESSRANGETYPE_VOLATILE,
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
	resource := redfishapi.CxlLogicalDeviceV100CxlLogicalDevice{
		OdataContext: "/redfish/v1/$metadata#CXLLogicalDevice.CXLLogicalDevice",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices/" + cXLLogicalDeviceId,
		OdataType:    CXLLogicalDevice,
		Description:  "CXL Logical Device " + string(dev.GetCxlType()),
		Id:           cXLLogicalDeviceId,
		Links: redfishapi.CxlLogicalDeviceV100Links{
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
		SemanticsSupported: []redfishapi.CxlLogicalDeviceV100CxlSemantic{},
		Status: redfishapi.ResourceStatus{
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
		},
		Identifiers: []redfishapi.ResourceIdentifier{{
			DurableName:       &gcxlid,
			DurableNameFormat: "GCXLID",
		}},
	}
	devCap := dev.GetCxlCap()
	if devCap.IO_En {
		resource.SemanticsSupported = append(resource.SemanticsSupported, redfishapi.CXLLOGICALDEVICEV100CXLSEMANTIC_CX_LIO)
	}
	if devCap.Cache_En {
		resource.SemanticsSupported = append(resource.SemanticsSupported, redfishapi.CXLLOGICALDEVICEV100CXLSEMANTIC_CX_LCACHE)
	}
	if devCap.Mem_En {
		resource.SemanticsSupported = append(resource.SemanticsSupported, redfishapi.CXLLOGICALDEVICEV100CXLSEMANTIC_CX_LMEM)
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
	resource := redfishapi.PcieDeviceV1111PcieDevice{
		OdataContext: "/redfish/v1/$metadata#PCIeDevice.PCIeDevice",
		OdataId:      "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId,
		OdataType:    PCIeDevice,
		Actions:      redfishapi.PcieDeviceV1111Actions{},
		Assembly:     redfishapi.OdataV4IdRef{},
		CXLDevice: redfishapi.PcieDeviceV1111CxlDevice{
			DeviceType:              redfishapi.PcieDeviceV1111CxlDeviceType(dev.GetCxlType()),
			MaxNumberLogicalDevices: new(float32),
		},
		CXLLogicalDevices: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/CXLLogicalDevices",
		},
		Description: "PCIe Device " + pCIeDeviceId,
		DeviceType:  redfishapi.PCIEDEVICEV1111DEVICETYPE_SINGLE_FUNCTION,
		EnvironmentMetrics: redfishapi.OdataV4IdRef{
			OdataId: "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId,
		},
		Id: pCIeDeviceId,
		Links: redfishapi.PcieDeviceV1111Links{
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
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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
	resource := redfishapi.PcieFunctionV150PcieFunction{
		OdataContext:     "/redfish/v1/$metadata#PCIeFunction.PCIeFunction",
		OdataId:          "/redfish/v1/Chassis/" + chassisId + "/PCIeDevices/" + pCIeDeviceId + "/PCIeFunctions/" + pCIeFunctionId,
		OdataType:        PCIeFunction,
		Description:      "PCIe Function",
		ClassCode:        &classCode,
		DeviceClass:      redfishapi.PCIEFUNCTIONV150DEVICECLASS_UNCLASSIFIED_DEVICE,
		DeviceId:         &pCIeDeviceId,
		FunctionId:       &functionid,
		FunctionProtocol: redfishapi.PCIEFUNCTIONV150FUNCTIONPROTOCOL_CXL,
		FunctionType:     redfishapi.PCIEFUNCTIONV150FUNCTIONTYPE_PHYSICAL,
		Id:               pCIeDeviceId + pCIeFunctionId,
		Links: redfishapi.PcieFunctionV150Links{
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
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
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
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdActionsComputerSystemResetPost(ctx context.Context, computerSystemId string, computerSystemV1201ResetRequestBody redfishapi.ComputerSystemV1201ResetRequestBody) (redfishapi.ImplResponse, error) {
	ResetType := computerSystemV1201ResetRequestBody.ResetType
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

type SystemType struct {
	OdataContext string `json:"@odata.context,omitempty"`
	// OdataEtag string `json:"@odata.etag,omitempty"`
	OdataId   string                         `json:"@odata.id"`
	OdataType string                         `json:"@odata.type"`
	Actions   ComputerSystemV1201ActionsType `json:"Actions,omitempty"`
	// AssetTag *string `json:"AssetTag,omitempty"`
	// Bios OdataV4IdRef `json:"Bios,omitempty"`
	// BiosVersion *string `json:"BiosVersion,omitempty"`
	// Boot ComputerSystemV1201Boot `json:"Boot,omitempty"`
	// BootProgress ComputerSystemV1201BootProgress `json:"BootProgress,omitempty"`
	// Certificates OdataV4IdRef `json:"Certificates,omitempty"`
	// Composition ComputerSystemV1201Composition `json:"Composition,omitempty"`
	// Description string `json:"Description,omitempty"`
	// EthernetInterfaces OdataV4IdRef `json:"EthernetInterfaces,omitempty"`
	// FabricAdapters OdataV4IdRef `json:"FabricAdapters,omitempty"`
	// GraphicalConsole ComputerSystemV1201HostGraphicalConsole `json:"GraphicalConsole,omitempty"`
	// GraphicsControllers OdataV4IdRef `json:"GraphicsControllers,omitempty"`
	// HostName *string `json:"HostName,omitempty"`
	// HostWatchdogTimer ComputerSystemV1201WatchdogTimer `json:"HostWatchdogTimer,omitempty"`
	// HostedServices ComputerSystemV1201HostedServices `json:"HostedServices,omitempty"`
	// HostingRoles []ComputerSystemV1201HostingRole `json:"HostingRoles,omitempty"`
	// Id string `json:"Id"`
	// IdlePowerSaver ComputerSystemV1201IdlePowerSaver `json:"IdlePowerSaver,omitempty"`
	// IndicatorLED ComputerSystemV1201IndicatorLed `json:"IndicatorLED,omitempty"`
	// KeyManagement ComputerSystemV1201KeyManagement `json:"KeyManagement,omitempty"`
	// LastResetTime time.Time `json:"LastResetTime,omitempty"`
	// Links ComputerSystemV1201Links `json:"Links,omitempty"`
	// LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`
	// LogServices OdataV4IdRef `json:"LogServices,omitempty"`
	// Manufacturer *string `json:"Manufacturer,omitempty"`
	// ManufacturingMode *bool `json:"ManufacturingMode,omitempty"`
	Memory        redfishapi.OdataV4IdRef `json:"Memory,omitempty"`
	MemoryDomains redfishapi.OdataV4IdRef `json:"MemoryDomains,omitempty"`
	// MemorySummary ComputerSystemV1201MemorySummary `json:"MemorySummary,omitempty"`
	// Model *string `json:"Model,omitempty"`
	// Name string `json:"Name"`
	// NetworkInterfaces OdataV4IdRef `json:"NetworkInterfaces,omitempty"`
	// Oem map[string]interface{} `json:"Oem,omitempty"`
	PCIeDevices []redfishapi.OdataV4IdRef `json:"PCIeDevices,omitempty"`
	// PCIeDevicesodataCount int64 `json:"PCIeDevices@odata.count,omitempty"`
	PCIeFunctions []redfishapi.OdataV4IdRef `json:"PCIeFunctions,omitempty"`
	// PCIeFunctionsodataCount int64 `json:"PCIeFunctions@odata.count,omitempty"`
	// PartNumber *string `json:"PartNumber,omitempty"`
	// PowerCycleDelaySeconds *float32 `json:"PowerCycleDelaySeconds,omitempty"`
	// PowerMode ComputerSystemV1201PowerMode `json:"PowerMode,omitempty"`
	// PowerOffDelaySeconds *float32 `json:"PowerOffDelaySeconds,omitempty"`
	// PowerOnDelaySeconds *float32 `json:"PowerOnDelaySeconds,omitempty"`
	// PowerRestorePolicy ComputerSystemV1201PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`
	// PowerState ResourcePowerState `json:"PowerState,omitempty"`
	// ProcessorSummary ComputerSystemV1201ProcessorSummary `json:"ProcessorSummary,omitempty"`
	// Processors OdataV4IdRef `json:"Processors,omitempty"`
	// Redundancy []RedundancyRedundancy `json:"Redundancy,omitempty"`
	// RedundancyodataCount int64 `json:"Redundancy@odata.count,omitempty"`
	// SKU *string `json:"SKU,omitempty"`
	// SecureBoot OdataV4IdRef `json:"SecureBoot,omitempty"`
	// SerialConsole ComputerSystemV1201HostSerialConsole `json:"SerialConsole,omitempty"`
	// SerialNumber *string `json:"SerialNumber,omitempty"`
	// SimpleStorage OdataV4IdRef `json:"SimpleStorage,omitempty"`
	// Status ResourceStatus `json:"Status,omitempty"`
	// Storage OdataV4IdRef `json:"Storage,omitempty"`
	// SubModel *string `json:"SubModel,omitempty"`
	// SystemType ComputerSystemV1201SystemType `json:"SystemType,omitempty"`
	// USBControllers OdataV4IdRef `json:"USBControllers,omitempty"`
	// UUID string `json:"UUID,omitempty"`
	// VirtualMedia OdataV4IdRef `json:"VirtualMedia,omitempty"`
	// VirtualMediaConfig ComputerSystemV1201VirtualMediaConfig `json:"VirtualMediaConfig,omitempty"`
}

type ComputerSystemV1201ActionsType struct {
	// ComputerSystemAddResourceBlock ComputerSystemV1201AddResourceBlock `json:"#ComputerSystem.AddResourceBlock,omitempty"`
	// ComputerSystemRemoveResourceBlock ComputerSystemV1201RemoveResourceBlock `json:"#ComputerSystem.RemoveResourceBlock,omitempty"`
	ComputerSystemReset redfishapi.ComputerSystemV1201Reset `json:"#ComputerSystem.Reset,omitempty"`
	// ComputerSystemSetDefaultBootOrder ComputerSystemV1201SetDefaultBootOrder `json:"#ComputerSystem.SetDefaultBootOrder,omitempty"`
	// Oem map[string]interface{} `json:"Oem,omitempty"`
}

// RedfishV1SystemsComputerSystemIdGet -
func (s *CxlHostApiService) RedfishV1SystemsComputerSystemIdGet(ctx context.Context, computerSystemId string) (redfishapi.ImplResponse, error) {
	if computerSystemId != GetSystemName() {
		return redfishapi.Response(http.StatusNotFound, nil), fmt.Errorf("System id (%s) does not exist", computerSystemId)
	}

	system := SystemType{
		OdataContext: "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
		OdataId:      "/redfish/v1/Systems/" + GetSystemName(),
		OdataType:    SystemVersion,
		Actions: ComputerSystemV1201ActionsType{
			ComputerSystemReset: redfishapi.ComputerSystemV1201Reset{
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

	resource := redfishapi.MemoryChunksV150MemoryChunks{
		OdataId:          "/redfish/v1/Systems/" + computerSystemId + "/MemoryDomains/" + memoryDomainId + "/MemoryChunks/" + memoryChunksId,
		OdataType:        MemoryChunkVersion,
		Name:             "Memory Chunk " + memoryChunksId,
		Id:               memoryChunksId,
		AddressRangeType: redfishapi.MEMORYCHUNKSV150ADDRESSRANGETYPE_VOLATILE,
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

	resource := redfishapi.MemoryV1171Memory{
		OdataContext: "/redfish/v1/$metadata#Memory.Memory",
		OdataId:      "/redfish/v1/Systems/" + computerSystemId + "/Memory/" + memoryId,
		OdataType:    MemoryVersionForSystem,
		CapacityMiB:  &total,
		Id:           memoryId,
		Name:         "CXL Device memory",
		MemoryMedia:  []redfishapi.MemoryV1171MemoryMedia{redfishapi.MEMORYV1171MEMORYMEDIA_CXL},
		Status: redfishapi.ResourceStatus{
			Health: redfishapi.RESOURCEHEALTH_OK,
			State:  redfishapi.RESOURCESTATE_ENABLED,
		},
		Links: redfishapi.MemoryV1171Links{
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
func (s *CxlHostApiService) RedfishV1AccountServiceAccountsManagerAccountIdPut(ctx context.Context, managerAccountId string, managerAccountV1100ManagerAccount redfishapi.ManagerAccountV1100ManagerAccount) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServiceAccountsManagerAccountIdPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.ManagerAccountV1100ManagerAccount{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.ManagerAccountV1100ManagerAccount{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServiceAccountsManagerAccountIdPut method not implemented")
}

// RedfishV1AccountServicePatch -
func (s *CxlHostApiService) RedfishV1AccountServicePatch(ctx context.Context, accountServiceV1130AccountService redfishapi.AccountServiceV1130AccountService) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServicePatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.AccountServiceV1130AccountService{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.AccountServiceV1130AccountService{}), nil

	// TODO: Uncomment the next line to return response Response(202, redfishapi.TaskV171Task{}) or use other options such as http.Ok ...
	// return redfishapi.Response(202, redfishapi.TaskV171Task{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return redfishapi.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, redfishapi.RedfishError{}) or use other options such as http.Ok ...
	// return redfishapi.Response(0, redfishapi.RedfishError{}), nil

	return redfishapi.Response(http.StatusNotImplemented, nil), errors.New("RedfishV1AccountServicePatch method not implemented")
}

// RedfishV1AccountServicePut -
func (s *CxlHostApiService) RedfishV1AccountServicePut(ctx context.Context, accountServiceV1130AccountService redfishapi.AccountServiceV1130AccountService) (redfishapi.ImplResponse, error) {
	// TODO - update RedfishV1AccountServicePut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, redfishapi.AccountServiceV1130AccountService{}) or use other options such as http.Ok ...
	// return redfishapi.Response(200, redfishapi.AccountServiceV1130AccountService{}), nil

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
