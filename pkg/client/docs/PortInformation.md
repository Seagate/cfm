# PortInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this resource | 
**GCxlId** | Pointer to **string** | The global CXL identifier | [optional] 
**LinkedPortUri** | Pointer to **string** | A full path to the resource with id as the last component | [optional] 
**PortProtocol** | Pointer to **string** | The protocol being sent over this port | [optional] 
**PortMedium** | Pointer to **string** | The physical connection medium for this port | [optional] 
**CurrentSpeedGbps** | Pointer to **int32** | The current speed of this port | [optional] 
**StatusHealth** | **string** | The health of the resource | 
**StatusState** | **string** | The state of the resource | 
**Width** | Pointer to **int32** | The number of lanes, phys, or other physical transport links that this port contains | [optional] 
**LinkStatus** | Pointer to **string** | Status of the link, such as LinkUp or LinkDown | [optional] 

## Methods

### NewPortInformation

`func NewPortInformation(id string, statusHealth string, statusState string, ) *PortInformation`

NewPortInformation instantiates a new PortInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortInformationWithDefaults

`func NewPortInformationWithDefaults() *PortInformation`

NewPortInformationWithDefaults instantiates a new PortInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortInformation) SetId(v string)`

SetId sets Id field to given value.


### GetGCxlId

`func (o *PortInformation) GetGCxlId() string`

GetGCxlId returns the GCxlId field if non-nil, zero value otherwise.

### GetGCxlIdOk

`func (o *PortInformation) GetGCxlIdOk() (*string, bool)`

GetGCxlIdOk returns a tuple with the GCxlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGCxlId

`func (o *PortInformation) SetGCxlId(v string)`

SetGCxlId sets GCxlId field to given value.

### HasGCxlId

`func (o *PortInformation) HasGCxlId() bool`

HasGCxlId returns a boolean if a field has been set.

### GetLinkedPortUri

`func (o *PortInformation) GetLinkedPortUri() string`

GetLinkedPortUri returns the LinkedPortUri field if non-nil, zero value otherwise.

### GetLinkedPortUriOk

`func (o *PortInformation) GetLinkedPortUriOk() (*string, bool)`

GetLinkedPortUriOk returns a tuple with the LinkedPortUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedPortUri

`func (o *PortInformation) SetLinkedPortUri(v string)`

SetLinkedPortUri sets LinkedPortUri field to given value.

### HasLinkedPortUri

`func (o *PortInformation) HasLinkedPortUri() bool`

HasLinkedPortUri returns a boolean if a field has been set.

### GetPortProtocol

`func (o *PortInformation) GetPortProtocol() string`

GetPortProtocol returns the PortProtocol field if non-nil, zero value otherwise.

### GetPortProtocolOk

`func (o *PortInformation) GetPortProtocolOk() (*string, bool)`

GetPortProtocolOk returns a tuple with the PortProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortProtocol

`func (o *PortInformation) SetPortProtocol(v string)`

SetPortProtocol sets PortProtocol field to given value.

### HasPortProtocol

`func (o *PortInformation) HasPortProtocol() bool`

HasPortProtocol returns a boolean if a field has been set.

### GetPortMedium

`func (o *PortInformation) GetPortMedium() string`

GetPortMedium returns the PortMedium field if non-nil, zero value otherwise.

### GetPortMediumOk

`func (o *PortInformation) GetPortMediumOk() (*string, bool)`

GetPortMediumOk returns a tuple with the PortMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMedium

`func (o *PortInformation) SetPortMedium(v string)`

SetPortMedium sets PortMedium field to given value.

### HasPortMedium

`func (o *PortInformation) HasPortMedium() bool`

HasPortMedium returns a boolean if a field has been set.

### GetCurrentSpeedGbps

`func (o *PortInformation) GetCurrentSpeedGbps() int32`

GetCurrentSpeedGbps returns the CurrentSpeedGbps field if non-nil, zero value otherwise.

### GetCurrentSpeedGbpsOk

`func (o *PortInformation) GetCurrentSpeedGbpsOk() (*int32, bool)`

GetCurrentSpeedGbpsOk returns a tuple with the CurrentSpeedGbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpeedGbps

`func (o *PortInformation) SetCurrentSpeedGbps(v int32)`

SetCurrentSpeedGbps sets CurrentSpeedGbps field to given value.

### HasCurrentSpeedGbps

`func (o *PortInformation) HasCurrentSpeedGbps() bool`

HasCurrentSpeedGbps returns a boolean if a field has been set.

### GetStatusHealth

`func (o *PortInformation) GetStatusHealth() string`

GetStatusHealth returns the StatusHealth field if non-nil, zero value otherwise.

### GetStatusHealthOk

`func (o *PortInformation) GetStatusHealthOk() (*string, bool)`

GetStatusHealthOk returns a tuple with the StatusHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHealth

`func (o *PortInformation) SetStatusHealth(v string)`

SetStatusHealth sets StatusHealth field to given value.


### GetStatusState

`func (o *PortInformation) GetStatusState() string`

GetStatusState returns the StatusState field if non-nil, zero value otherwise.

### GetStatusStateOk

`func (o *PortInformation) GetStatusStateOk() (*string, bool)`

GetStatusStateOk returns a tuple with the StatusState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusState

`func (o *PortInformation) SetStatusState(v string)`

SetStatusState sets StatusState field to given value.


### GetWidth

`func (o *PortInformation) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PortInformation) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PortInformation) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PortInformation) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLinkStatus

`func (o *PortInformation) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *PortInformation) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *PortInformation) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *PortInformation) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


