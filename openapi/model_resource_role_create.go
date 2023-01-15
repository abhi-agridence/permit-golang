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

// ResourceRoleCreate struct for ResourceRoleCreate
type ResourceRoleCreate struct {
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

// NewResourceRoleCreate instantiates a new ResourceRoleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRoleCreate(key string, name string) *ResourceRoleCreate {
	this := ResourceRoleCreate{}
	this.Key = key
	this.Name = name
	return &this
}

// NewResourceRoleCreateWithDefaults instantiates a new ResourceRoleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRoleCreateWithDefaults() *ResourceRoleCreate {
	this := ResourceRoleCreate{}
	return &this
}

// GetKey returns the Key field value
func (o *ResourceRoleCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResourceRoleCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResourceRoleCreate) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ResourceRoleCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceRoleCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceRoleCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceRoleCreate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRoleCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceRoleCreate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceRoleCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ResourceRoleCreate) GetPermissions() []string {
	if o == nil || isNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRoleCreate) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ResourceRoleCreate) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *ResourceRoleCreate) SetPermissions(v []string) {
	o.Permissions = v
}

// GetExtends returns the Extends field value if set, zero value otherwise.
func (o *ResourceRoleCreate) GetExtends() []string {
	if o == nil || isNil(o.Extends) {
		var ret []string
		return ret
	}
	return o.Extends
}

// GetExtendsOk returns a tuple with the Extends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRoleCreate) GetExtendsOk() ([]string, bool) {
	if o == nil || isNil(o.Extends) {
		return nil, false
	}
	return o.Extends, true
}

// HasExtends returns a boolean if a field has been set.
func (o *ResourceRoleCreate) HasExtends() bool {
	if o != nil && !isNil(o.Extends) {
		return true
	}

	return false
}

// SetExtends gets a reference to the given []string and assigns it to the Extends field.
func (o *ResourceRoleCreate) SetExtends(v []string) {
	o.Extends = v
}

func (o ResourceRoleCreate) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !isNil(o.Extends) {
		toSerialize["extends"] = o.Extends
	}
	return json.Marshal(toSerialize)
}

type NullableResourceRoleCreate struct {
	value *ResourceRoleCreate
	isSet bool
}

func (v NullableResourceRoleCreate) Get() *ResourceRoleCreate {
	return v.value
}

func (v *NullableResourceRoleCreate) Set(val *ResourceRoleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRoleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRoleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRoleCreate(val *ResourceRoleCreate) *NullableResourceRoleCreate {
	return &NullableResourceRoleCreate{value: val, isSet: true}
}

func (v NullableResourceRoleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRoleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
