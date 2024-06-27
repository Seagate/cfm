# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | 
**IpAddress** | **string** | The IP Address in dot notation of the service | 
**Port** | **int32** |  | 
**Status** | Pointer to **string** | A response string | [optional] 
**Ports** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**MemoryDevices** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**Memory** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**LocalMemoryMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 
**RemoteMemoryMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 

## Methods

### NewHost

`func NewHost(id string, ipAddress string, port int32, ) *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Host) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *Host) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Host) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Host) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *Host) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Host) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Host) SetPort(v int32)`

SetPort sets Port field to given value.


### GetStatus

`func (o *Host) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Host) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Host) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Host) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPorts

`func (o *Host) GetPorts() MemberItem`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Host) GetPortsOk() (*MemberItem, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Host) SetPorts(v MemberItem)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Host) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetMemoryDevices

`func (o *Host) GetMemoryDevices() MemberItem`

GetMemoryDevices returns the MemoryDevices field if non-nil, zero value otherwise.

### GetMemoryDevicesOk

`func (o *Host) GetMemoryDevicesOk() (*MemberItem, bool)`

GetMemoryDevicesOk returns a tuple with the MemoryDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDevices

`func (o *Host) SetMemoryDevices(v MemberItem)`

SetMemoryDevices sets MemoryDevices field to given value.

### HasMemoryDevices

`func (o *Host) HasMemoryDevices() bool`

HasMemoryDevices returns a boolean if a field has been set.

### GetMemory

`func (o *Host) GetMemory() MemberItem`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Host) GetMemoryOk() (*MemberItem, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Host) SetMemory(v MemberItem)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Host) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetLocalMemoryMiB

`func (o *Host) GetLocalMemoryMiB() int32`

GetLocalMemoryMiB returns the LocalMemoryMiB field if non-nil, zero value otherwise.

### GetLocalMemoryMiBOk

`func (o *Host) GetLocalMemoryMiBOk() (*int32, bool)`

GetLocalMemoryMiBOk returns a tuple with the LocalMemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMemoryMiB

`func (o *Host) SetLocalMemoryMiB(v int32)`

SetLocalMemoryMiB sets LocalMemoryMiB field to given value.

### HasLocalMemoryMiB

`func (o *Host) HasLocalMemoryMiB() bool`

HasLocalMemoryMiB returns a boolean if a field has been set.

### GetRemoteMemoryMiB

`func (o *Host) GetRemoteMemoryMiB() int32`

GetRemoteMemoryMiB returns the RemoteMemoryMiB field if non-nil, zero value otherwise.

### GetRemoteMemoryMiBOk

`func (o *Host) GetRemoteMemoryMiBOk() (*int32, bool)`

GetRemoteMemoryMiBOk returns a tuple with the RemoteMemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMemoryMiB

`func (o *Host) SetRemoteMemoryMiB(v int32)`

SetRemoteMemoryMiB sets RemoteMemoryMiB field to given value.

### HasRemoteMemoryMiB

`func (o *Host) HasRemoteMemoryMiB() bool`

HasRemoteMemoryMiB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


