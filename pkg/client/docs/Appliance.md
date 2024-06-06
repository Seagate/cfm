# Appliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | 
**IpAddress** | Pointer to **string** | The IP Address in dot notation of the service | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | A response string | [optional] 
**Blades** | Pointer to [**MemberItem**](MemberItem.md) |  | [optional] 
**TotalMemoryAvailableMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 
**TotalMemoryAllocatedMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 

## Methods

### NewAppliance

`func NewAppliance(id string, ) *Appliance`

NewAppliance instantiates a new Appliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceWithDefaults

`func NewApplianceWithDefaults() *Appliance`

NewApplianceWithDefaults instantiates a new Appliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Appliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Appliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Appliance) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *Appliance) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Appliance) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Appliance) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Appliance) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *Appliance) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Appliance) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Appliance) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Appliance) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *Appliance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Appliance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Appliance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Appliance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBlades

`func (o *Appliance) GetBlades() MemberItem`

GetBlades returns the Blades field if non-nil, zero value otherwise.

### GetBladesOk

`func (o *Appliance) GetBladesOk() (*MemberItem, bool)`

GetBladesOk returns a tuple with the Blades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlades

`func (o *Appliance) SetBlades(v MemberItem)`

SetBlades sets Blades field to given value.

### HasBlades

`func (o *Appliance) HasBlades() bool`

HasBlades returns a boolean if a field has been set.

### GetTotalMemoryAvailableMiB

`func (o *Appliance) GetTotalMemoryAvailableMiB() int32`

GetTotalMemoryAvailableMiB returns the TotalMemoryAvailableMiB field if non-nil, zero value otherwise.

### GetTotalMemoryAvailableMiBOk

`func (o *Appliance) GetTotalMemoryAvailableMiBOk() (*int32, bool)`

GetTotalMemoryAvailableMiBOk returns a tuple with the TotalMemoryAvailableMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryAvailableMiB

`func (o *Appliance) SetTotalMemoryAvailableMiB(v int32)`

SetTotalMemoryAvailableMiB sets TotalMemoryAvailableMiB field to given value.

### HasTotalMemoryAvailableMiB

`func (o *Appliance) HasTotalMemoryAvailableMiB() bool`

HasTotalMemoryAvailableMiB returns a boolean if a field has been set.

### GetTotalMemoryAllocatedMiB

`func (o *Appliance) GetTotalMemoryAllocatedMiB() int32`

GetTotalMemoryAllocatedMiB returns the TotalMemoryAllocatedMiB field if non-nil, zero value otherwise.

### GetTotalMemoryAllocatedMiBOk

`func (o *Appliance) GetTotalMemoryAllocatedMiBOk() (*int32, bool)`

GetTotalMemoryAllocatedMiBOk returns a tuple with the TotalMemoryAllocatedMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryAllocatedMiB

`func (o *Appliance) SetTotalMemoryAllocatedMiB(v int32)`

SetTotalMemoryAllocatedMiB sets TotalMemoryAllocatedMiB field to given value.

### HasTotalMemoryAllocatedMiB

`func (o *Appliance) HasTotalMemoryAllocatedMiB() bool`

HasTotalMemoryAllocatedMiB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


