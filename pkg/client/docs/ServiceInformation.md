# ServiceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The cfm-service version | 
**Resources** | [**[]ServiceResource**](ServiceResource.md) |  | 

## Methods

### NewServiceInformation

`func NewServiceInformation(version string, resources []ServiceResource, ) *ServiceInformation`

NewServiceInformation instantiates a new ServiceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInformationWithDefaults

`func NewServiceInformationWithDefaults() *ServiceInformation`

NewServiceInformationWithDefaults instantiates a new ServiceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServiceInformation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceInformation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceInformation) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetResources

`func (o *ServiceInformation) GetResources() []ServiceResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServiceInformation) GetResourcesOk() (*[]ServiceResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServiceInformation) SetResources(v []ServiceResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


