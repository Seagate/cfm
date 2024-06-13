/*
Composable Fabric Manager Service OpenAPI

This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Collection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Collection{}

// Collection struct for Collection
type Collection struct {
	MemberCount int32        `json:"memberCount"`
	Members     []MemberItem `json:"members"`
}

// NewCollection instantiates a new Collection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollection(memberCount int32, members []MemberItem) *Collection {
	this := Collection{}
	this.MemberCount = memberCount
	this.Members = members
	return &this
}

// NewCollectionWithDefaults instantiates a new Collection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionWithDefaults() *Collection {
	this := Collection{}
	return &this
}

// GetMemberCount returns the MemberCount field value
func (o *Collection) GetMemberCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *Collection) GetMemberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *Collection) SetMemberCount(v int32) {
	o.MemberCount = v
}

// GetMembers returns the Members field value
func (o *Collection) GetMembers() []MemberItem {
	if o == nil {
		var ret []MemberItem
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *Collection) GetMembersOk() ([]MemberItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *Collection) SetMembers(v []MemberItem) {
	o.Members = v
}

func (o Collection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Collection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["memberCount"] = o.MemberCount
	toSerialize["members"] = o.Members
	return toSerialize, nil
}

type NullableCollection struct {
	value *Collection
	isSet bool
}

func (v NullableCollection) Get() *Collection {
	return v.value
}

func (v *NullableCollection) Set(val *Collection) {
	v.value = val
	v.isSet = true
}

func (v NullableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollection(val *Collection) *NullableCollection {
	return &NullableCollection{value: val, isSet: true}
}

func (v NullableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
