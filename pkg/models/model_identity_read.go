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

// IdentityRead struct for IdentityRead
type IdentityRead struct {
	// Unique User Id of this identity in the identity provider (including the provider type)
	UserId string `json:"user_id"`
	// The identity provider type this identity came from
	Provider string `json:"provider"`
	// Unique User Id of this identity in the identity provider (NOT including the provider type)
	Sub string `json:"sub"`
	// Email connected to this account identity
	Email string `json:"email"`
	// Whether this email address connected to this account identity is verified or not. For social providers like 'Login with Google' this is done automatically, otherwise we will send the user a verification link in email.
	EmailVerified bool `json:"email_verified"`
	// Raw user info json coming from our identity provider and matching a specific account identity
	Auth0Info map[string]interface{} `json:"auth0_info"`
}

// NewIdentityRead instantiates a new IdentityRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRead(userId string, provider string, sub string, email string, emailVerified bool, auth0Info map[string]interface{}) *IdentityRead {
	this := IdentityRead{}
	this.UserId = userId
	this.Provider = provider
	this.Sub = sub
	this.Email = email
	this.EmailVerified = emailVerified
	this.Auth0Info = auth0Info
	return &this
}

// NewIdentityReadWithDefaults instantiates a new IdentityRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityReadWithDefaults() *IdentityRead {
	this := IdentityRead{}
	return &this
}

// GetUserId returns the UserId field value
func (o *IdentityRead) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *IdentityRead) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *IdentityRead) SetUserId(v string) {
	o.UserId = v
}

// GetProvider returns the Provider field value
func (o *IdentityRead) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *IdentityRead) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *IdentityRead) SetProvider(v string) {
	o.Provider = v
}

// GetSub returns the Sub field value
func (o *IdentityRead) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *IdentityRead) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *IdentityRead) SetSub(v string) {
	o.Sub = v
}

// GetEmail returns the Email field value
func (o *IdentityRead) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IdentityRead) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IdentityRead) SetEmail(v string) {
	o.Email = v
}

// GetEmailVerified returns the EmailVerified field value
func (o *IdentityRead) GetEmailVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
func (o *IdentityRead) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *IdentityRead) SetEmailVerified(v bool) {
	o.EmailVerified = v
}

// GetAuth0Info returns the Auth0Info field value
func (o *IdentityRead) GetAuth0Info() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Auth0Info
}

// GetAuth0InfoOk returns a tuple with the Auth0Info field value
// and a boolean to check if the value has been set.
func (o *IdentityRead) GetAuth0InfoOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Auth0Info, true
}

// SetAuth0Info sets field value
func (o *IdentityRead) SetAuth0Info(v map[string]interface{}) {
	o.Auth0Info = v
}

func (o IdentityRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["sub"] = o.Sub
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if true {
		toSerialize["auth0_info"] = o.Auth0Info
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRead struct {
	value *IdentityRead
	isSet bool
}

func (v NullableIdentityRead) Get() *IdentityRead {
	return v.value
}

func (v *NullableIdentityRead) Set(val *IdentityRead) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRead) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRead(val *IdentityRead) *NullableIdentityRead {
	return &NullableIdentityRead{value: val, isSet: true}
}

func (v NullableIdentityRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}