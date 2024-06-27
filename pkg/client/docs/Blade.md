# Blade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | 
**IpAddress** | **string** | The IP Address in dot notation of the service | 
**Port** | **int32** |  | 
**Status** | Pointer to **string** | A response string | [optional] 
**Ports** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**Resources** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**Memory** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**TotalMemoryAvailableMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 
**TotalMemoryAllocatedMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 

## Methods

### NewBlade

`func NewBlade(id string, ipAddress string, port int32, ) *Blade`

NewBlade instantiates a new Blade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBladeWithDefaults

`func NewBladeWithDefaults() *Blade`

NewBladeWithDefaults instantiates a new Blade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Blade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Blade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Blade) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *Blade) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Blade) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Blade) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *Blade) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Blade) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Blade) SetPort(v int32)`

SetPort sets Port field to given value.


### GetStatus

`func (o *Blade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Blade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Blade) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Blade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPorts

`func (o *Blade) GetPorts() MemberItem`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Blade) GetPortsOk() (*MemberItem, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Blade) SetPorts(v MemberItem)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Blade) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetResources

`func (o *Blade) GetResources() MemberItem`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Blade) GetResourcesOk() (*MemberItem, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Blade) SetResources(v MemberItem)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Blade) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetMemory

`func (o *Blade) GetMemory() MemberItem`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Blade) GetMemoryOk() (*MemberItem, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Blade) SetMemory(v MemberItem)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Blade) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetTotalMemoryAvailableMiB

`func (o *Blade) GetTotalMemoryAvailableMiB() int32`

GetTotalMemoryAvailableMiB returns the TotalMemoryAvailableMiB field if non-nil, zero value otherwise.

### GetTotalMemoryAvailableMiBOk

`func (o *Blade) GetTotalMemoryAvailableMiBOk() (*int32, bool)`

GetTotalMemoryAvailableMiBOk returns a tuple with the TotalMemoryAvailableMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryAvailableMiB

`func (o *Blade) SetTotalMemoryAvailableMiB(v int32)`

SetTotalMemoryAvailableMiB sets TotalMemoryAvailableMiB field to given value.

### HasTotalMemoryAvailableMiB

`func (o *Blade) HasTotalMemoryAvailableMiB() bool`

HasTotalMemoryAvailableMiB returns a boolean if a field has been set.

### GetTotalMemoryAllocatedMiB

`func (o *Blade) GetTotalMemoryAllocatedMiB() int32`

GetTotalMemoryAllocatedMiB returns the TotalMemoryAllocatedMiB field if non-nil, zero value otherwise.

### GetTotalMemoryAllocatedMiBOk

`func (o *Blade) GetTotalMemoryAllocatedMiBOk() (*int32, bool)`

GetTotalMemoryAllocatedMiBOk returns a tuple with the TotalMemoryAllocatedMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryAllocatedMiB

`func (o *Blade) SetTotalMemoryAllocatedMiB(v int32)`

SetTotalMemoryAllocatedMiB sets TotalMemoryAllocatedMiB field to given value.

### HasTotalMemoryAllocatedMiB

`func (o *Blade) HasTotalMemoryAllocatedMiB() bool`

HasTotalMemoryAllocatedMiB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


