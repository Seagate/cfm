# AssignMemoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **string** | The CXL port name on the Memory Appliance | 
**Operation** | **string** | -assign- an existing memory to a port or -unassign- an existing link between memory and port | 

## Methods

### NewAssignMemoryRequest

`func NewAssignMemoryRequest(port string, operation string, ) *AssignMemoryRequest`

NewAssignMemoryRequest instantiates a new AssignMemoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignMemoryRequestWithDefaults

`func NewAssignMemoryRequestWithDefaults() *AssignMemoryRequest`

NewAssignMemoryRequestWithDefaults instantiates a new AssignMemoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *AssignMemoryRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssignMemoryRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssignMemoryRequest) SetPort(v string)`

SetPort sets Port field to given value.


### GetOperation

`func (o *AssignMemoryRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AssignMemoryRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AssignMemoryRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


