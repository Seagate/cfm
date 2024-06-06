/*
Composer and Fabric Manager Service OpenAPI

This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the MemberItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberItem{}

// MemberItem struct for MemberItem
type MemberItem struct {
	// A full path to the resource with id as the last component
	Uri string `json:"uri"`
}

// NewMemberItem instantiates a new MemberItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberItem(uri string) *MemberItem {
	this := MemberItem{}
	this.Uri = uri
	return &this
}

// NewMemberItemWithDefaults instantiates a new MemberItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberItemWithDefaults() *MemberItem {
	this := MemberItem{}
	return &this
}

// GetUri returns the Uri field value
func (o *MemberItem) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *MemberItem) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *MemberItem) SetUri(v string) {
	o.Uri = v
}

func (o MemberItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableMemberItem struct {
	value *MemberItem
	isSet bool
}

func (v NullableMemberItem) Get() *MemberItem {
	return v.value
}

func (v *NullableMemberItem) Set(val *MemberItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberItem(val *MemberItem) *NullableMemberItem {
	return &NullableMemberItem{value: val, isSet: true}
}

func (v NullableMemberItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}