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

// EnvironmentUpdate struct for EnvironmentUpdate
type EnvironmentUpdate struct {
	// The name of the environment
	Name *string `json:"name,omitempty"`
	// an optional longer description of the environment
	Description *string `json:"description,omitempty"`
	// when using gitops feature, an optional branch name for the environment
	CustomBranchName *string `json:"custom_branch_name,omitempty"`
}

// NewEnvironmentUpdate instantiates a new EnvironmentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentUpdate() *EnvironmentUpdate {
	this := EnvironmentUpdate{}
	return &this
}

// NewEnvironmentUpdateWithDefaults instantiates a new EnvironmentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentUpdateWithDefaults() *EnvironmentUpdate {
	this := EnvironmentUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomBranchName returns the CustomBranchName field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetCustomBranchName() string {
	if o == nil || isNil(o.CustomBranchName) {
		var ret string
		return ret
	}
	return *o.CustomBranchName
}

// GetCustomBranchNameOk returns a tuple with the CustomBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetCustomBranchNameOk() (*string, bool) {
	if o == nil || isNil(o.CustomBranchName) {
		return nil, false
	}
	return o.CustomBranchName, true
}

// HasCustomBranchName returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasCustomBranchName() bool {
	if o != nil && !isNil(o.CustomBranchName) {
		return true
	}

	return false
}

// SetCustomBranchName gets a reference to the given string and assigns it to the CustomBranchName field.
func (o *EnvironmentUpdate) SetCustomBranchName(v string) {
	o.CustomBranchName = &v
}

func (o EnvironmentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CustomBranchName) {
		toSerialize["custom_branch_name"] = o.CustomBranchName
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentUpdate struct {
	value *EnvironmentUpdate
	isSet bool
}

func (v NullableEnvironmentUpdate) Get() *EnvironmentUpdate {
	return v.value
}

func (v *NullableEnvironmentUpdate) Set(val *EnvironmentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentUpdate(val *EnvironmentUpdate) *NullableEnvironmentUpdate {
	return &NullableEnvironmentUpdate{value: val, isSet: true}
}

func (v NullableEnvironmentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
