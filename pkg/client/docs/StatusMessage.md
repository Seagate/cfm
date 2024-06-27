# StatusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | The URI of the http request | 
**Details** | **string** | Additional information provided to the client regarding this response | 
**Status** | Pointer to [**StatusMessageStatus**](StatusMessageStatus.md) |  | [optional] 

## Methods

### NewStatusMessage

`func NewStatusMessage(uri string, details string, ) *StatusMessage`

NewStatusMessage instantiates a new StatusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMessageWithDefaults

`func NewStatusMessageWithDefaults() *StatusMessage`

NewStatusMessageWithDefaults instantiates a new StatusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *StatusMessage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *StatusMessage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *StatusMessage) SetUri(v string)`

SetUri sets Uri field to given value.


### GetDetails

`func (o *StatusMessage) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *StatusMessage) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *StatusMessage) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetStatus

`func (o *StatusMessage) GetStatus() StatusMessageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusMessage) GetStatusOk() (*StatusMessageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusMessage) SetStatus(v StatusMessageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatusMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


