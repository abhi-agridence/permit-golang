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

// EnvironmentCreate struct for EnvironmentCreate
type EnvironmentCreate struct {
	// A URL-friendly name of the environment (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the environment.
	Key string `json:"key"`
	// The name of the environment
	Name string `json:"name"`
	// an optional longer description of the environment
	Description *string `json:"description,omitempty"`
	// when using gitops feature, an optional branch name for the environment
	CustomBranchName *string `json:"custom_branch_name,omitempty"`
}

// NewEnvironmentCreate instantiates a new EnvironmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentCreate(key string, name string) *EnvironmentCreate {
	this := EnvironmentCreate{}
	this.Key = key
	this.Name = name
	return &this
}

// NewEnvironmentCreateWithDefaults instantiates a new EnvironmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentCreateWithDefaults() *EnvironmentCreate {
	this := EnvironmentCreate{}
	return &this
}

// GetKey returns the Key field value
func (o *EnvironmentCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentCreate) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *EnvironmentCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentCreate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomBranchName returns the CustomBranchName field value if set, zero value otherwise.
func (o *EnvironmentCreate) GetCustomBranchName() string {
	if o == nil || isNil(o.CustomBranchName) {
		var ret string
		return ret
	}
	return *o.CustomBranchName
}

// GetCustomBranchNameOk returns a tuple with the CustomBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreate) GetCustomBranchNameOk() (*string, bool) {
	if o == nil || isNil(o.CustomBranchName) {
		return nil, false
	}
	return o.CustomBranchName, true
}

// HasCustomBranchName returns a boolean if a field has been set.
func (o *EnvironmentCreate) HasCustomBranchName() bool {
	if o != nil && !isNil(o.CustomBranchName) {
		return true
	}

	return false
}

// SetCustomBranchName gets a reference to the given string and assigns it to the CustomBranchName field.
func (o *EnvironmentCreate) SetCustomBranchName(v string) {
	o.CustomBranchName = &v
}

func (o EnvironmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
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

type NullableEnvironmentCreate struct {
	value *EnvironmentCreate
	isSet bool
}

func (v NullableEnvironmentCreate) Get() *EnvironmentCreate {
	return v.value
}

func (v *NullableEnvironmentCreate) Set(val *EnvironmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentCreate(val *EnvironmentCreate) *NullableEnvironmentCreate {
	return &NullableEnvironmentCreate{value: val, isSet: true}
}

func (v NullableEnvironmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
