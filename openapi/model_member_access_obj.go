/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MemberAccessObj An enumeration.
type MemberAccessObj string

// List of MemberAccessObj
const (
	ORG     MemberAccessObj = "org"
	PROJECT MemberAccessObj = "project"
	ENV     MemberAccessObj = "env"
)

// All allowed values of MemberAccessObj enum
var AllowedMemberAccessObjEnumValues = []MemberAccessObj{
	"org",
	"project",
	"env",
}

func (v *MemberAccessObj) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MemberAccessObj(value)
	for _, existing := range AllowedMemberAccessObjEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MemberAccessObj", value)
}

// NewMemberAccessObjFromValue returns a pointer to a valid MemberAccessObj
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMemberAccessObjFromValue(v string) (*MemberAccessObj, error) {
	ev := MemberAccessObj(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MemberAccessObj: valid values are %v", v, AllowedMemberAccessObjEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MemberAccessObj) IsValid() bool {
	for _, existing := range AllowedMemberAccessObjEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MemberAccessObj value
func (v MemberAccessObj) Ptr() *MemberAccessObj {
	return &v
}

type NullableMemberAccessObj struct {
	value *MemberAccessObj
	isSet bool
}

func (v NullableMemberAccessObj) Get() *MemberAccessObj {
	return v.value
}

func (v *NullableMemberAccessObj) Set(val *MemberAccessObj) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberAccessObj) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberAccessObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberAccessObj(val *MemberAccessObj) *NullableMemberAccessObj {
	return &NullableMemberAccessObj{value: val, isSet: true}
}

func (v NullableMemberAccessObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberAccessObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
