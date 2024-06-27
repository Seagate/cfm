# ComposeMemoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | The CXL port name on the Memory Appliance | [optional] 
**MemorySizeMiB** | **int32** | A mebibyte equals 2**20 or 1,048,576 bytes. | 
**QoS** | [**Qos**](Qos.md) |  | 

## Methods

### NewComposeMemoryRequest

`func NewComposeMemoryRequest(memorySizeMiB int32, qoS Qos, ) *ComposeMemoryRequest`

NewComposeMemoryRequest instantiates a new ComposeMemoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeMemoryRequestWithDefaults

`func NewComposeMemoryRequestWithDefaults() *ComposeMemoryRequest`

NewComposeMemoryRequestWithDefaults instantiates a new ComposeMemoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ComposeMemoryRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ComposeMemoryRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ComposeMemoryRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ComposeMemoryRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMemorySizeMiB

`func (o *ComposeMemoryRequest) GetMemorySizeMiB() int32`

GetMemorySizeMiB returns the MemorySizeMiB field if non-nil, zero value otherwise.

### GetMemorySizeMiBOk

`func (o *ComposeMemoryRequest) GetMemorySizeMiBOk() (*int32, bool)`

GetMemorySizeMiBOk returns a tuple with the MemorySizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMiB

`func (o *ComposeMemoryRequest) SetMemorySizeMiB(v int32)`

SetMemorySizeMiB sets MemorySizeMiB field to given value.


### GetQoS

`func (o *ComposeMemoryRequest) GetQoS() Qos`

GetQoS returns the QoS field if non-nil, zero value otherwise.

### GetQoSOk

`func (o *ComposeMemoryRequest) GetQoSOk() (*Qos, bool)`

GetQoSOk returns a tuple with the QoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoS

`func (o *ComposeMemoryRequest) SetQoS(v Qos)`

SetQoS sets QoS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


