# MemoryResourceBlockCompositionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositionState** | **string** | The current state of the resource block from a composition perspective | 
**MaxCompositions** | Pointer to **int32** | The maximum number of compositions in which this resource block can participate simultaneously | [optional] 
**NumberOfCompositions** | Pointer to **int32** | The number of compositions in which this resource block is currently participating | [optional] 

## Methods

### NewMemoryResourceBlockCompositionStatus

`func NewMemoryResourceBlockCompositionStatus(compositionState string, ) *MemoryResourceBlockCompositionStatus`

NewMemoryResourceBlockCompositionStatus instantiates a new MemoryResourceBlockCompositionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryResourceBlockCompositionStatusWithDefaults

`func NewMemoryResourceBlockCompositionStatusWithDefaults() *MemoryResourceBlockCompositionStatus`

NewMemoryResourceBlockCompositionStatusWithDefaults instantiates a new MemoryResourceBlockCompositionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositionState

`func (o *MemoryResourceBlockCompositionStatus) GetCompositionState() string`

GetCompositionState returns the CompositionState field if non-nil, zero value otherwise.

### GetCompositionStateOk

`func (o *MemoryResourceBlockCompositionStatus) GetCompositionStateOk() (*string, bool)`

GetCompositionStateOk returns a tuple with the CompositionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositionState

`func (o *MemoryResourceBlockCompositionStatus) SetCompositionState(v string)`

SetCompositionState sets CompositionState field to given value.


### GetMaxCompositions

`func (o *MemoryResourceBlockCompositionStatus) GetMaxCompositions() int32`

GetMaxCompositions returns the MaxCompositions field if non-nil, zero value otherwise.

### GetMaxCompositionsOk

`func (o *MemoryResourceBlockCompositionStatus) GetMaxCompositionsOk() (*int32, bool)`

GetMaxCompositionsOk returns a tuple with the MaxCompositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompositions

`func (o *MemoryResourceBlockCompositionStatus) SetMaxCompositions(v int32)`

SetMaxCompositions sets MaxCompositions field to given value.

### HasMaxCompositions

`func (o *MemoryResourceBlockCompositionStatus) HasMaxCompositions() bool`

HasMaxCompositions returns a boolean if a field has been set.

### GetNumberOfCompositions

`func (o *MemoryResourceBlockCompositionStatus) GetNumberOfCompositions() int32`

GetNumberOfCompositions returns the NumberOfCompositions field if non-nil, zero value otherwise.

### GetNumberOfCompositionsOk

`func (o *MemoryResourceBlockCompositionStatus) GetNumberOfCompositionsOk() (*int32, bool)`

GetNumberOfCompositionsOk returns a tuple with the NumberOfCompositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCompositions

`func (o *MemoryResourceBlockCompositionStatus) SetNumberOfCompositions(v int32)`

SetNumberOfCompositions sets NumberOfCompositions field to given value.

### HasNumberOfCompositions

`func (o *MemoryResourceBlockCompositionStatus) HasNumberOfCompositions() bool`

HasNumberOfCompositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


