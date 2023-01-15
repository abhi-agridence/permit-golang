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

// InviteCreate struct for InviteCreate
type InviteCreate struct {
	// The invited member's email address
	Email string `json:"email"`
	// The role the member will be assigned with
	Role string `json:"role"`
}

// NewInviteCreate instantiates a new InviteCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteCreate(email string, role string) *InviteCreate {
	this := InviteCreate{}
	this.Email = email
	this.Role = role
	return &this
}

// NewInviteCreateWithDefaults instantiates a new InviteCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteCreateWithDefaults() *InviteCreate {
	this := InviteCreate{}
	return &this
}

// GetEmail returns the Email field value
func (o *InviteCreate) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InviteCreate) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InviteCreate) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *InviteCreate) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InviteCreate) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InviteCreate) SetRole(v string) {
	o.Role = v
}

func (o InviteCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInviteCreate struct {
	value *InviteCreate
	isSet bool
}

func (v NullableInviteCreate) Get() *InviteCreate {
	return v.value
}

func (v *NullableInviteCreate) Set(val *InviteCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteCreate(val *InviteCreate) *NullableInviteCreate {
	return &NullableInviteCreate{value: val, isSet: true}
}

func (v NullableInviteCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
