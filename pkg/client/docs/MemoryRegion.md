# MemoryRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | 
**Status** | Pointer to **string** | A response string | [optional] 
**Type** | [**MemoryType**](MemoryType.md) |  | 
**SizeMiB** | **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | 
**Bandwidth** | Pointer to **int32** | Memory bandwidth in the unit of GigaBytes per second | [optional] 
**Latency** | Pointer to **int32** | Memory latency in the unit of nanosecond | [optional] 
**MemoryApplianceId** | Pointer to **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | [optional] 
**MemoryBladeId** | Pointer to **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | [optional] 
**MemoryAppliancePort** | Pointer to **string** | The CXL port name on the Memory Appliance | [optional] 
**MappedHostId** | Pointer to **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | [optional] 
**MappedHostPort** | Pointer to **string** | The id uniquely identifies the resource within a resource collection. Since URIs are constructed with ids, must not contain RFC1738 unsafe characters. | [optional] 

## Methods

### NewMemoryRegion

`func NewMemoryRegion(id string, type_ MemoryType, sizeMiB int32, ) *MemoryRegion`

NewMemoryRegion instantiates a new MemoryRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryRegionWithDefaults

`func NewMemoryRegionWithDefaults() *MemoryRegion`

NewMemoryRegionWithDefaults instantiates a new MemoryRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemoryRegion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemoryRegion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemoryRegion) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *MemoryRegion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemoryRegion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemoryRegion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MemoryRegion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MemoryRegion) GetType() MemoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemoryRegion) GetTypeOk() (*MemoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemoryRegion) SetType(v MemoryType)`

SetType sets Type field to given value.


### GetSizeMiB

`func (o *MemoryRegion) GetSizeMiB() int32`

GetSizeMiB returns the SizeMiB field if non-nil, zero value otherwise.

### GetSizeMiBOk

`func (o *MemoryRegion) GetSizeMiBOk() (*int32, bool)`

GetSizeMiBOk returns a tuple with the SizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMiB

`func (o *MemoryRegion) SetSizeMiB(v int32)`

SetSizeMiB sets SizeMiB field to given value.


### GetBandwidth

`func (o *MemoryRegion) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *MemoryRegion) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *MemoryRegion) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *MemoryRegion) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetLatency

`func (o *MemoryRegion) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *MemoryRegion) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *MemoryRegion) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *MemoryRegion) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMemoryApplianceId

`func (o *MemoryRegion) GetMemoryApplianceId() string`

GetMemoryApplianceId returns the MemoryApplianceId field if non-nil, zero value otherwise.

### GetMemoryApplianceIdOk

`func (o *MemoryRegion) GetMemoryApplianceIdOk() (*string, bool)`

GetMemoryApplianceIdOk returns a tuple with the MemoryApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryApplianceId

`func (o *MemoryRegion) SetMemoryApplianceId(v string)`

SetMemoryApplianceId sets MemoryApplianceId field to given value.

### HasMemoryApplianceId

`func (o *MemoryRegion) HasMemoryApplianceId() bool`

HasMemoryApplianceId returns a boolean if a field has been set.

### GetMemoryBladeId

`func (o *MemoryRegion) GetMemoryBladeId() string`

GetMemoryBladeId returns the MemoryBladeId field if non-nil, zero value otherwise.

### GetMemoryBladeIdOk

`func (o *MemoryRegion) GetMemoryBladeIdOk() (*string, bool)`

GetMemoryBladeIdOk returns a tuple with the MemoryBladeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBladeId

`func (o *MemoryRegion) SetMemoryBladeId(v string)`

SetMemoryBladeId sets MemoryBladeId field to given value.

### HasMemoryBladeId

`func (o *MemoryRegion) HasMemoryBladeId() bool`

HasMemoryBladeId returns a boolean if a field has been set.

### GetMemoryAppliancePort

`func (o *MemoryRegion) GetMemoryAppliancePort() string`

GetMemoryAppliancePort returns the MemoryAppliancePort field if non-nil, zero value otherwise.

### GetMemoryAppliancePortOk

`func (o *MemoryRegion) GetMemoryAppliancePortOk() (*string, bool)`

GetMemoryAppliancePortOk returns a tuple with the MemoryAppliancePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAppliancePort

`func (o *MemoryRegion) SetMemoryAppliancePort(v string)`

SetMemoryAppliancePort sets MemoryAppliancePort field to given value.

### HasMemoryAppliancePort

`func (o *MemoryRegion) HasMemoryAppliancePort() bool`

HasMemoryAppliancePort returns a boolean if a field has been set.

### GetMappedHostId

`func (o *MemoryRegion) GetMappedHostId() string`

GetMappedHostId returns the MappedHostId field if non-nil, zero value otherwise.

### GetMappedHostIdOk

`func (o *MemoryRegion) GetMappedHostIdOk() (*string, bool)`

GetMappedHostIdOk returns a tuple with the MappedHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedHostId

`func (o *MemoryRegion) SetMappedHostId(v string)`

SetMappedHostId sets MappedHostId field to given value.

### HasMappedHostId

`func (o *MemoryRegion) HasMappedHostId() bool`

HasMappedHostId returns a boolean if a field has been set.

### GetMappedHostPort

`func (o *MemoryRegion) GetMappedHostPort() string`

GetMappedHostPort returns the MappedHostPort field if non-nil, zero value otherwise.

### GetMappedHostPortOk

`func (o *MemoryRegion) GetMappedHostPortOk() (*string, bool)`

GetMappedHostPortOk returns a tuple with the MappedHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedHostPort

`func (o *MemoryRegion) SetMappedHostPort(v string)`

SetMappedHostPort sets MappedHostPort field to given value.

### HasMappedHostPort

`func (o *MemoryRegion) HasMappedHostPort() bool`

HasMappedHostPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


