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

// ConditionSetRuleCreate struct for ConditionSetRuleCreate
type ConditionSetRuleCreate struct {
	// The userset that will be given permission, i.e: all the users matching this rule will be given the specified permission
	UserSet string `json:"user_set"`
	// The permission that will be granted to the userset *on* the resourceset. The permission can be either a resource action id, or `{resource_key}:{action_key}`, i.e: the \"permission name\".
	Permission string `json:"permission"`
	// The resourceset that represents the resources that are granted for access, i.e: all the resources matching this rule can be accessed by the userset to perform the granted *permission*
	ResourceSet string `json:"resource_set"`
	// if True, will set the condition set rule to the role's autogen user-set.
	IsRole *bool `json:"is_role,omitempty"`
	// if True, will set the condition set rule to the resource's autogen resource-set.
	IsResource *bool `json:"is_resource,omitempty"`
}

// NewConditionSetRuleCreate instantiates a new ConditionSetRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionSetRuleCreate(userSet string, permission string, resourceSet string) *ConditionSetRuleCreate {
	this := ConditionSetRuleCreate{}
	this.UserSet = userSet
	this.Permission = permission
	this.ResourceSet = resourceSet
	var isRole bool = false
	this.IsRole = &isRole
	var isResource bool = false
	this.IsResource = &isResource
	return &this
}

// NewConditionSetRuleCreateWithDefaults instantiates a new ConditionSetRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionSetRuleCreateWithDefaults() *ConditionSetRuleCreate {
	this := ConditionSetRuleCreate{}
	var isRole bool = false
	this.IsRole = &isRole
	var isResource bool = false
	this.IsResource = &isResource
	return &this
}

// GetUserSet returns the UserSet field value
func (o *ConditionSetRuleCreate) GetUserSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserSet
}

// GetUserSetOk returns a tuple with the UserSet field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRuleCreate) GetUserSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserSet, true
}

// SetUserSet sets field value
func (o *ConditionSetRuleCreate) SetUserSet(v string) {
	o.UserSet = v
}

// GetPermission returns the Permission field value
func (o *ConditionSetRuleCreate) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRuleCreate) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *ConditionSetRuleCreate) SetPermission(v string) {
	o.Permission = v
}

// GetResourceSet returns the ResourceSet field value
func (o *ConditionSetRuleCreate) GetResourceSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRuleCreate) GetResourceSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSet, true
}

// SetResourceSet sets field value
func (o *ConditionSetRuleCreate) SetResourceSet(v string) {
	o.ResourceSet = v
}

// GetIsRole returns the IsRole field value if set, zero value otherwise.
func (o *ConditionSetRuleCreate) GetIsRole() bool {
	if o == nil || IsNil(o.IsRole) {
		var ret bool
		return ret
	}
	return *o.IsRole
}

// GetIsRoleOk returns a tuple with the IsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRuleCreate) GetIsRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRole) {
		return nil, false
	}
	return o.IsRole, true
}

// HasIsRole returns a boolean if a field has been set.
func (o *ConditionSetRuleCreate) HasIsRole() bool {
	if o != nil && !IsNil(o.IsRole) {
		return true
	}

	return false
}

// SetIsRole gets a reference to the given bool and assigns it to the IsRole field.
func (o *ConditionSetRuleCreate) SetIsRole(v bool) {
	o.IsRole = &v
}

// GetIsResource returns the IsResource field value if set, zero value otherwise.
func (o *ConditionSetRuleCreate) GetIsResource() bool {
	if o == nil || IsNil(o.IsResource) {
		var ret bool
		return ret
	}
	return *o.IsResource
}

// GetIsResourceOk returns a tuple with the IsResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRuleCreate) GetIsResourceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsResource) {
		return nil, false
	}
	return o.IsResource, true
}

// HasIsResource returns a boolean if a field has been set.
func (o *ConditionSetRuleCreate) HasIsResource() bool {
	if o != nil && !IsNil(o.IsResource) {
		return true
	}

	return false
}

// SetIsResource gets a reference to the given bool and assigns it to the IsResource field.
func (o *ConditionSetRuleCreate) SetIsResource(v bool) {
	o.IsResource = &v
}

func (o ConditionSetRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_set"] = o.UserSet
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if true {
		toSerialize["resource_set"] = o.ResourceSet
	}
	if !IsNil(o.IsRole) {
		toSerialize["is_role"] = o.IsRole
	}
	if !IsNil(o.IsResource) {
		toSerialize["is_resource"] = o.IsResource
	}
	return json.Marshal(toSerialize)
}

type NullableConditionSetRuleCreate struct {
	value *ConditionSetRuleCreate
	isSet bool
}

func (v NullableConditionSetRuleCreate) Get() *ConditionSetRuleCreate {
	return v.value
}

func (v *NullableConditionSetRuleCreate) Set(val *ConditionSetRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionSetRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionSetRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionSetRuleCreate(val *ConditionSetRuleCreate) *NullableConditionSetRuleCreate {
	return &NullableConditionSetRuleCreate{value: val, isSet: true}
}

func (v NullableConditionSetRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionSetRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
