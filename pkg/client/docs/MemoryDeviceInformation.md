# MemoryDeviceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this resource | 
**DeviceType** | Pointer to **string** | The type of the device | [optional] 
**MemorySizeMiB** | Pointer to **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | [optional] 
**LinkStatus** | Pointer to [**MemoryDeviceInformationLinkStatus**](MemoryDeviceInformationLinkStatus.md) |  | [optional] 
**StatusState** | Pointer to **string** | The state of the resource | [optional] 

## Methods

### NewMemoryDeviceInformation

`func NewMemoryDeviceInformation(id string, ) *MemoryDeviceInformation`

NewMemoryDeviceInformation instantiates a new MemoryDeviceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryDeviceInformationWithDefaults

`func NewMemoryDeviceInformationWithDefaults() *MemoryDeviceInformation`

NewMemoryDeviceInformationWithDefaults instantiates a new MemoryDeviceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemoryDeviceInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemoryDeviceInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemoryDeviceInformation) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *MemoryDeviceInformation) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MemoryDeviceInformation) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MemoryDeviceInformation) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *MemoryDeviceInformation) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetMemorySizeMiB

`func (o *MemoryDeviceInformation) GetMemorySizeMiB() int32`

GetMemorySizeMiB returns the MemorySizeMiB field if non-nil, zero value otherwise.

### GetMemorySizeMiBOk

`func (o *MemoryDeviceInformation) GetMemorySizeMiBOk() (*int32, bool)`

GetMemorySizeMiBOk returns a tuple with the MemorySizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMiB

`func (o *MemoryDeviceInformation) SetMemorySizeMiB(v int32)`

SetMemorySizeMiB sets MemorySizeMiB field to given value.

### HasMemorySizeMiB

`func (o *MemoryDeviceInformation) HasMemorySizeMiB() bool`

HasMemorySizeMiB returns a boolean if a field has been set.

### GetLinkStatus

`func (o *MemoryDeviceInformation) GetLinkStatus() MemoryDeviceInformationLinkStatus`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *MemoryDeviceInformation) GetLinkStatusOk() (*MemoryDeviceInformationLinkStatus, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *MemoryDeviceInformation) SetLinkStatus(v MemoryDeviceInformationLinkStatus)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *MemoryDeviceInformation) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetStatusState

`func (o *MemoryDeviceInformation) GetStatusState() string`

GetStatusState returns the StatusState field if non-nil, zero value otherwise.

### GetStatusStateOk

`func (o *MemoryDeviceInformation) GetStatusStateOk() (*string, bool)`

GetStatusStateOk returns a tuple with the StatusState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusState

`func (o *MemoryDeviceInformation) SetStatusState(v string)`

SetStatusState sets StatusState field to given value.

### HasStatusState

`func (o *MemoryDeviceInformation) HasStatusState() bool`

HasStatusState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


