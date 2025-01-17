/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// RoleCreate struct for RoleCreate
type RoleCreate struct {
	// A URL-friendly name of the role (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the role.
	Key string `json:"key"`
	// The name of the role
	Name string `json:"name"`
	// optional description string explaining what this role represents, or what permissions are granted to it.
	Description *string `json:"description,omitempty"`
	// list of action keys that define what actions this resource role is permitted to do
	Permissions []string `json:"permissions,omitempty"`
	// list of role keys that define what roles this role extends. In other words: this role will automatically inherit all the permissions of the given roles in this list.
	Extends []string `json:"extends,omitempty"`
}

// NewRoleCreate instantiates a new RoleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCreate(key string, name string) *RoleCreate {
	this := RoleCreate{}
	this.Key = key
	this.Name = name
	return &this
}

// NewRoleCreateWithDefaults instantiates a new RoleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCreateWithDefaults() *RoleCreate {
	this := RoleCreate{}
	return &this
}

// GetKey returns the Key field value
func (o *RoleCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RoleCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RoleCreate) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *RoleCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleCreate) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreate) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleCreate) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *RoleCreate) SetPermissions(v []string) {
	o.Permissions = v
}

// GetExtends returns the Extends field value if set, zero value otherwise.
func (o *RoleCreate) GetExtends() []string {
	if o == nil || IsNil(o.Extends) {
		var ret []string
		return ret
	}
	return o.Extends
}

// GetExtendsOk returns a tuple with the Extends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreate) GetExtendsOk() ([]string, bool) {
	if o == nil || IsNil(o.Extends) {
		return nil, false
	}
	return o.Extends, true
}

// HasExtends returns a boolean if a field has been set.
func (o *RoleCreate) HasExtends() bool {
	if o != nil && !IsNil(o.Extends) {
		return true
	}

	return false
}

// SetExtends gets a reference to the given []string and assigns it to the Extends field.
func (o *RoleCreate) SetExtends(v []string) {
	o.Extends = v
}

func (o RoleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Extends) {
		toSerialize["extends"] = o.Extends
	}
	return json.Marshal(toSerialize)
}

type NullableRoleCreate struct {
	value *RoleCreate
	isSet bool
}

func (v NullableRoleCreate) Get() *RoleCreate {
	return v.value
}

func (v *NullableRoleCreate) Set(val *RoleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCreate(val *RoleCreate) *NullableRoleCreate {
	return &NullableRoleCreate{value: val, isSet: true}
}

func (v NullableRoleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
