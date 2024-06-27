# ComposeMemoryByResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | The CXL port name on the Memory Appliance | [optional] 
**MemoryResources** | **[]string** |  | 

## Methods

### NewComposeMemoryByResourceRequest

`func NewComposeMemoryByResourceRequest(memoryResources []string, ) *ComposeMemoryByResourceRequest`

NewComposeMemoryByResourceRequest instantiates a new ComposeMemoryByResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeMemoryByResourceRequestWithDefaults

`func NewComposeMemoryByResourceRequestWithDefaults() *ComposeMemoryByResourceRequest`

NewComposeMemoryByResourceRequestWithDefaults instantiates a new ComposeMemoryByResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ComposeMemoryByResourceRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ComposeMemoryByResourceRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ComposeMemoryByResourceRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ComposeMemoryByResourceRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMemoryResources

`func (o *ComposeMemoryByResourceRequest) GetMemoryResources() []string`

GetMemoryResources returns the MemoryResources field if non-nil, zero value otherwise.

### GetMemoryResourcesOk

`func (o *ComposeMemoryByResourceRequest) GetMemoryResourcesOk() (*[]string, bool)`

GetMemoryResourcesOk returns a tuple with the MemoryResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryResources

`func (o *ComposeMemoryByResourceRequest) SetMemoryResources(v []string)`

SetMemoryResources sets MemoryResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


