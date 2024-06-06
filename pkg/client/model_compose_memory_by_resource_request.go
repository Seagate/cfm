/*
Composer and Fabric Manager Service OpenAPI

This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ComposeMemoryByResourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComposeMemoryByResourceRequest{}

// ComposeMemoryByResourceRequest struct for ComposeMemoryByResourceRequest
type ComposeMemoryByResourceRequest struct {
	// The CXL port name on the Memory Appliance
	Port            *string  `json:"Port,omitempty"`
	MemoryResources []string `json:"memoryResources"`
}

// NewComposeMemoryByResourceRequest instantiates a new ComposeMemoryByResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComposeMemoryByResourceRequest(memoryResources []string) *ComposeMemoryByResourceRequest {
	this := ComposeMemoryByResourceRequest{}
	this.MemoryResources = memoryResources
	return &this
}

// NewComposeMemoryByResourceRequestWithDefaults instantiates a new ComposeMemoryByResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComposeMemoryByResourceRequestWithDefaults() *ComposeMemoryByResourceRequest {
	this := ComposeMemoryByResourceRequest{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ComposeMemoryByResourceRequest) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposeMemoryByResourceRequest) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ComposeMemoryByResourceRequest) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *ComposeMemoryByResourceRequest) SetPort(v string) {
	o.Port = &v
}

// GetMemoryResources returns the MemoryResources field value
func (o *ComposeMemoryByResourceRequest) GetMemoryResources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MemoryResources
}

// GetMemoryResourcesOk returns a tuple with the MemoryResources field value
// and a boolean to check if the value has been set.
func (o *ComposeMemoryByResourceRequest) GetMemoryResourcesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryResources, true
}

// SetMemoryResources sets field value
func (o *ComposeMemoryByResourceRequest) SetMemoryResources(v []string) {
	o.MemoryResources = v
}

func (o ComposeMemoryByResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComposeMemoryByResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	toSerialize["memoryResources"] = o.MemoryResources
	return toSerialize, nil
}

type NullableComposeMemoryByResourceRequest struct {
	value *ComposeMemoryByResourceRequest
	isSet bool
}

func (v NullableComposeMemoryByResourceRequest) Get() *ComposeMemoryByResourceRequest {
	return v.value
}

func (v *NullableComposeMemoryByResourceRequest) Set(val *ComposeMemoryByResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableComposeMemoryByResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableComposeMemoryByResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComposeMemoryByResourceRequest(val *ComposeMemoryByResourceRequest) *NullableComposeMemoryByResourceRequest {
	return &NullableComposeMemoryByResourceRequest{value: val, isSet: true}
}

func (v NullableComposeMemoryByResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComposeMemoryByResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}