# ServiceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | A full path to the resource with id as the last component | 
**Methods** | **string** | The service(s) available for the specified URI | 
**Description** | **string** | The description of service(s) offered by the URI | 

## Methods

### NewServiceResource

`func NewServiceResource(uri string, methods string, description string, ) *ServiceResource`

NewServiceResource instantiates a new ServiceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResourceWithDefaults

`func NewServiceResourceWithDefaults() *ServiceResource`

NewServiceResourceWithDefaults instantiates a new ServiceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *ServiceResource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ServiceResource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ServiceResource) SetUri(v string)`

SetUri sets Uri field to given value.


### GetMethods

`func (o *ServiceResource) GetMethods() string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ServiceResource) GetMethodsOk() (*string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ServiceResource) SetMethods(v string)`

SetMethods sets Methods field to given value.


### GetDescription

`func (o *ServiceResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceResource) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


