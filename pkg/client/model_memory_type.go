/*
Composer and Fabric Manager Service OpenAPI

This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// MemoryType The type of constructed memory.
type MemoryType string

// List of MemoryType
const (
	MEMORY_TYPE_UNKNOWN MemoryType = "MemoryTypeUnknown"
	MEMORY_TYPE_REGION  MemoryType = "MemoryTypeRegion"
	MEMORY_TYPE_LOCAL   MemoryType = "MemoryTypeLocal"
	MEMORY_TYPE_CXL     MemoryType = "MemoryTypeCXL"
)

// All allowed values of MemoryType enum
var AllowedMemoryTypeEnumValues = []MemoryType{
	"MemoryTypeUnknown",
	"MemoryTypeRegion",
	"MemoryTypeLocal",
	"MemoryTypeCXL",
}

func (v *MemoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MemoryType(value)
	for _, existing := range AllowedMemoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MemoryType", value)
}

// NewMemoryTypeFromValue returns a pointer to a valid MemoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMemoryTypeFromValue(v string) (*MemoryType, error) {
	ev := MemoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MemoryType: valid values are %v", v, AllowedMemoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MemoryType) IsValid() bool {
	for _, existing := range AllowedMemoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MemoryType value
func (v MemoryType) Ptr() *MemoryType {
	return &v
}

type NullableMemoryType struct {
	value *MemoryType
	isSet bool
}

func (v NullableMemoryType) Get() *MemoryType {
	return v.value
}

func (v *NullableMemoryType) Set(val *MemoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryType(val *MemoryType) *NullableMemoryType {
	return &NullableMemoryType{value: val, isSet: true}
}

func (v NullableMemoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
