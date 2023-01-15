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

// Settings struct for Settings
type Settings struct {
	bool   *bool
	int32  *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Settings) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &dst.bool)
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			return nil // data stored in dst.bool, return on the first match
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32)
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Settings)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Settings) MarshalJSON() ([]byte, error) {
	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSettings struct {
	value *Settings
	isSet bool
}

func (v NullableSettings) Get() *Settings {
	return v.value
}

func (v *NullableSettings) Set(val *Settings) {
	v.value = val
	v.isSet = true
}

func (v NullableSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettings(val *Settings) *NullableSettings {
	return &NullableSettings{value: val, isSet: true}
}

func (v NullableSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
