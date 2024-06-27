# MemoryResourceBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this resource | 
**CompositionStatus** | [**MemoryResourceBlockCompositionStatus**](MemoryResourceBlockCompositionStatus.md) |  | 
**CapacityMiB** | Pointer to **int32** | The number of compositions in which this resource block is currently participating | [optional] 
**DataWidthBits** | Pointer to **int32** | The number of compositions in which this resource block is currently participating | [optional] 
**MemoryType** | Pointer to **string** | The type of memory device | [optional] 
**MemoryDeviceType** | Pointer to **string** | Type details of the memory device | [optional] 
**Manufacturer** | Pointer to **string** | The memory device manufacturer | [optional] 
**OperatingSpeedMhz** | Pointer to **int32** | Operating speed of the memory device in MHz | [optional] 
**PartNumber** | Pointer to **string** | The product part number of this device | [optional] 
**SerialNumber** | Pointer to **string** | The product serial number of this device | [optional] 
**RankCount** | Pointer to **int32** | Number of ranks available in the memory device | [optional] 
**ChannelId** | **int32** | Id of the channel(dimm) associated with this resource | 
**ChannelResourceIndex** | **int32** | Position index for this resource within the designated channel(dimm) | 

## Methods

### NewMemoryResourceBlock

`func NewMemoryResourceBlock(id string, compositionStatus MemoryResourceBlockCompositionStatus, channelId int32, channelResourceIndex int32, ) *MemoryResourceBlock`

NewMemoryResourceBlock instantiates a new MemoryResourceBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryResourceBlockWithDefaults

`func NewMemoryResourceBlockWithDefaults() *MemoryResourceBlock`

NewMemoryResourceBlockWithDefaults instantiates a new MemoryResourceBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemoryResourceBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemoryResourceBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemoryResourceBlock) SetId(v string)`

SetId sets Id field to given value.


### GetCompositionStatus

`func (o *MemoryResourceBlock) GetCompositionStatus() MemoryResourceBlockCompositionStatus`

GetCompositionStatus returns the CompositionStatus field if non-nil, zero value otherwise.

### GetCompositionStatusOk

`func (o *MemoryResourceBlock) GetCompositionStatusOk() (*MemoryResourceBlockCompositionStatus, bool)`

GetCompositionStatusOk returns a tuple with the CompositionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositionStatus

`func (o *MemoryResourceBlock) SetCompositionStatus(v MemoryResourceBlockCompositionStatus)`

SetCompositionStatus sets CompositionStatus field to given value.


### GetCapacityMiB

`func (o *MemoryResourceBlock) GetCapacityMiB() int32`

GetCapacityMiB returns the CapacityMiB field if non-nil, zero value otherwise.

### GetCapacityMiBOk

`func (o *MemoryResourceBlock) GetCapacityMiBOk() (*int32, bool)`

GetCapacityMiBOk returns a tuple with the CapacityMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMiB

`func (o *MemoryResourceBlock) SetCapacityMiB(v int32)`

SetCapacityMiB sets CapacityMiB field to given value.

### HasCapacityMiB

`func (o *MemoryResourceBlock) HasCapacityMiB() bool`

HasCapacityMiB returns a boolean if a field has been set.

### GetDataWidthBits

`func (o *MemoryResourceBlock) GetDataWidthBits() int32`

GetDataWidthBits returns the DataWidthBits field if non-nil, zero value otherwise.

### GetDataWidthBitsOk

`func (o *MemoryResourceBlock) GetDataWidthBitsOk() (*int32, bool)`

GetDataWidthBitsOk returns a tuple with the DataWidthBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWidthBits

`func (o *MemoryResourceBlock) SetDataWidthBits(v int32)`

SetDataWidthBits sets DataWidthBits field to given value.

### HasDataWidthBits

`func (o *MemoryResourceBlock) HasDataWidthBits() bool`

HasDataWidthBits returns a boolean if a field has been set.

### GetMemoryType

`func (o *MemoryResourceBlock) GetMemoryType() string`

GetMemoryType returns the MemoryType field if non-nil, zero value otherwise.

### GetMemoryTypeOk

`func (o *MemoryResourceBlock) GetMemoryTypeOk() (*string, bool)`

GetMemoryTypeOk returns a tuple with the MemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryType

`func (o *MemoryResourceBlock) SetMemoryType(v string)`

SetMemoryType sets MemoryType field to given value.

### HasMemoryType

`func (o *MemoryResourceBlock) HasMemoryType() bool`

HasMemoryType returns a boolean if a field has been set.

### GetMemoryDeviceType

`func (o *MemoryResourceBlock) GetMemoryDeviceType() string`

GetMemoryDeviceType returns the MemoryDeviceType field if non-nil, zero value otherwise.

### GetMemoryDeviceTypeOk

`func (o *MemoryResourceBlock) GetMemoryDeviceTypeOk() (*string, bool)`

GetMemoryDeviceTypeOk returns a tuple with the MemoryDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDeviceType

`func (o *MemoryResourceBlock) SetMemoryDeviceType(v string)`

SetMemoryDeviceType sets MemoryDeviceType field to given value.

### HasMemoryDeviceType

`func (o *MemoryResourceBlock) HasMemoryDeviceType() bool`

HasMemoryDeviceType returns a boolean if a field has been set.

### GetManufacturer

`func (o *MemoryResourceBlock) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MemoryResourceBlock) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MemoryResourceBlock) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MemoryResourceBlock) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetOperatingSpeedMhz

`func (o *MemoryResourceBlock) GetOperatingSpeedMhz() int32`

GetOperatingSpeedMhz returns the OperatingSpeedMhz field if non-nil, zero value otherwise.

### GetOperatingSpeedMhzOk

`func (o *MemoryResourceBlock) GetOperatingSpeedMhzOk() (*int32, bool)`

GetOperatingSpeedMhzOk returns a tuple with the OperatingSpeedMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSpeedMhz

`func (o *MemoryResourceBlock) SetOperatingSpeedMhz(v int32)`

SetOperatingSpeedMhz sets OperatingSpeedMhz field to given value.

### HasOperatingSpeedMhz

`func (o *MemoryResourceBlock) HasOperatingSpeedMhz() bool`

HasOperatingSpeedMhz returns a boolean if a field has been set.

### GetPartNumber

`func (o *MemoryResourceBlock) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *MemoryResourceBlock) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *MemoryResourceBlock) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *MemoryResourceBlock) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MemoryResourceBlock) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MemoryResourceBlock) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MemoryResourceBlock) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MemoryResourceBlock) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetRankCount

`func (o *MemoryResourceBlock) GetRankCount() int32`

GetRankCount returns the RankCount field if non-nil, zero value otherwise.

### GetRankCountOk

`func (o *MemoryResourceBlock) GetRankCountOk() (*int32, bool)`

GetRankCountOk returns a tuple with the RankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankCount

`func (o *MemoryResourceBlock) SetRankCount(v int32)`

SetRankCount sets RankCount field to given value.

### HasRankCount

`func (o *MemoryResourceBlock) HasRankCount() bool`

HasRankCount returns a boolean if a field has been set.

### GetChannelId

`func (o *MemoryResourceBlock) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MemoryResourceBlock) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MemoryResourceBlock) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelResourceIndex

`func (o *MemoryResourceBlock) GetChannelResourceIndex() int32`

GetChannelResourceIndex returns the ChannelResourceIndex field if non-nil, zero value otherwise.

### GetChannelResourceIndexOk

`func (o *MemoryResourceBlock) GetChannelResourceIndexOk() (*int32, bool)`

GetChannelResourceIndexOk returns a tuple with the ChannelResourceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelResourceIndex

`func (o *MemoryResourceBlock) SetChannelResourceIndex(v int32)`

SetChannelResourceIndex sets ChannelResourceIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


