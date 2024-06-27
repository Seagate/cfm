# StatusMessageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | A unique status code value | 
**Message** | **string** | A description of the status code | 

## Methods

### NewStatusMessageStatus

`func NewStatusMessageStatus(code int32, message string, ) *StatusMessageStatus`

NewStatusMessageStatus instantiates a new StatusMessageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMessageStatusWithDefaults

`func NewStatusMessageStatusWithDefaults() *StatusMessageStatus`

NewStatusMessageStatusWithDefaults instantiates a new StatusMessageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StatusMessageStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatusMessageStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatusMessageStatus) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *StatusMessageStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusMessageStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusMessageStatus) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


