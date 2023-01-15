/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ActionBlockEditable struct for ActionBlockEditable
type ActionBlockEditable struct {
	// a more descriptive name for the action
	Name *string `json:"name,omitempty"`
	// optional description string explaining what this action represents in your system
	Description *string `json:"description,omitempty"`
}

// NewActionBlockEditable instantiates a new ActionBlockEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionBlockEditable() *ActionBlockEditable {
	this := ActionBlockEditable{}
	return &this
}

// NewActionBlockEditableWithDefaults instantiates a new ActionBlockEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionBlockEditableWithDefaults() *ActionBlockEditable {
	this := ActionBlockEditable{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActionBlockEditable) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionBlockEditable) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActionBlockEditable) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActionBlockEditable) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ActionBlockEditable) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionBlockEditable) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ActionBlockEditable) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ActionBlockEditable) SetDescription(v string) {
	o.Description = &v
}

func (o ActionBlockEditable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableActionBlockEditable struct {
	value *ActionBlockEditable
	isSet bool
}

func (v NullableActionBlockEditable) Get() *ActionBlockEditable {
	return v.value
}

func (v *NullableActionBlockEditable) Set(val *ActionBlockEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableActionBlockEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableActionBlockEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionBlockEditable(val *ActionBlockEditable) *NullableActionBlockEditable {
	return &NullableActionBlockEditable{value: val, isSet: true}
}

func (v NullableActionBlockEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionBlockEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
