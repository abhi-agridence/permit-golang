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

// ElementsType An enumeration.
type ElementsType string

// List of ElementsType
const (
	USER_MANAGEMENT ElementsType = "user_management"
	AUDIT_LOG       ElementsType = "audit_log"
	APPROVAL_FLOW   ElementsType = "approval_flow"
)

// All allowed values of ElementsType enum
var AllowedElementsTypeEnumValues = []ElementsType{
	"user_management",
	"audit_log",
	"approval_flow",
}

func (v *ElementsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ElementsType(value)
	for _, existing := range AllowedElementsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ElementsType", value)
}

// NewElementsTypeFromValue returns a pointer to a valid ElementsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewElementsTypeFromValue(v string) (*ElementsType, error) {
	ev := ElementsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ElementsType: valid values are %v", v, AllowedElementsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ElementsType) IsValid() bool {
	for _, existing := range AllowedElementsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElementsType value
func (v ElementsType) Ptr() *ElementsType {
	return &v
}

type NullableElementsType struct {
	value *ElementsType
	isSet bool
}

func (v NullableElementsType) Get() *ElementsType {
	return v.value
}

func (v *NullableElementsType) Set(val *ElementsType) {
	v.value = val
	v.isSet = true
}

func (v NullableElementsType) IsSet() bool {
	return v.isSet
}

func (v *NullableElementsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElementsType(val *ElementsType) *NullableElementsType {
	return &NullableElementsType{value: val, isSet: true}
}

func (v NullableElementsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElementsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
