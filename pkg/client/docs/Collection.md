# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberCount** | **int32** |  | 
**Members** | [**[]MemberItem**](MemberItem.md) |  | 

## Methods

### NewCollection

`func NewCollection(memberCount int32, members []MemberItem, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberCount

`func (o *Collection) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *Collection) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *Collection) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetMembers

`func (o *Collection) GetMembers() []MemberItem`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Collection) GetMembersOk() (*[]MemberItem, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Collection) SetMembers(v []MemberItem)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


