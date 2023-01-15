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

// PolicyRepoRead struct for PolicyRepoRead
type PolicyRepoRead struct {
	// Unique id of the policy repo
	Id     string           `json:"id"`
	Status PolicyRepoStatus `json:"status"`
	// A URL-friendly name of the policy repo (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the policy repo.
	Key            string      `json:"key"`
	Url            string      `json:"url"`
	MainBranchName *string     `json:"main_branch_name,omitempty"`
	Credentials    SSHAuthData `json:"credentials"`
	// if you want to change your policy repository to this repo right after it is validated
	ActivateWhenValidated *bool `json:"activate_when_validated,omitempty"`
}

// NewPolicyRepoRead instantiates a new PolicyRepoRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRepoRead(id string, status PolicyRepoStatus, key string, url string, credentials SSHAuthData) *PolicyRepoRead {
	this := PolicyRepoRead{}
	this.Id = id
	this.Status = status
	this.Key = key
	this.Url = url
	var mainBranchName string = "main"
	this.MainBranchName = &mainBranchName
	this.Credentials = credentials
	var activateWhenValidated bool = false
	this.ActivateWhenValidated = &activateWhenValidated
	return &this
}

// NewPolicyRepoReadWithDefaults instantiates a new PolicyRepoRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRepoReadWithDefaults() *PolicyRepoRead {
	this := PolicyRepoRead{}
	var mainBranchName string = "main"
	this.MainBranchName = &mainBranchName
	var activateWhenValidated bool = false
	this.ActivateWhenValidated = &activateWhenValidated
	return &this
}

// GetId returns the Id field value
func (o *PolicyRepoRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PolicyRepoRead) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *PolicyRepoRead) GetStatus() PolicyRepoStatus {
	if o == nil {
		var ret PolicyRepoStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetStatusOk() (*PolicyRepoStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PolicyRepoRead) SetStatus(v PolicyRepoStatus) {
	o.Status = v
}

// GetKey returns the Key field value
func (o *PolicyRepoRead) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PolicyRepoRead) SetKey(v string) {
	o.Key = v
}

// GetUrl returns the Url field value
func (o *PolicyRepoRead) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PolicyRepoRead) SetUrl(v string) {
	o.Url = v
}

// GetMainBranchName returns the MainBranchName field value if set, zero value otherwise.
func (o *PolicyRepoRead) GetMainBranchName() string {
	if o == nil || isNil(o.MainBranchName) {
		var ret string
		return ret
	}
	return *o.MainBranchName
}

// GetMainBranchNameOk returns a tuple with the MainBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetMainBranchNameOk() (*string, bool) {
	if o == nil || isNil(o.MainBranchName) {
		return nil, false
	}
	return o.MainBranchName, true
}

// HasMainBranchName returns a boolean if a field has been set.
func (o *PolicyRepoRead) HasMainBranchName() bool {
	if o != nil && !isNil(o.MainBranchName) {
		return true
	}

	return false
}

// SetMainBranchName gets a reference to the given string and assigns it to the MainBranchName field.
func (o *PolicyRepoRead) SetMainBranchName(v string) {
	o.MainBranchName = &v
}

// GetCredentials returns the Credentials field value
func (o *PolicyRepoRead) GetCredentials() SSHAuthData {
	if o == nil {
		var ret SSHAuthData
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetCredentialsOk() (*SSHAuthData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *PolicyRepoRead) SetCredentials(v SSHAuthData) {
	o.Credentials = v
}

// GetActivateWhenValidated returns the ActivateWhenValidated field value if set, zero value otherwise.
func (o *PolicyRepoRead) GetActivateWhenValidated() bool {
	if o == nil || isNil(o.ActivateWhenValidated) {
		var ret bool
		return ret
	}
	return *o.ActivateWhenValidated
}

// GetActivateWhenValidatedOk returns a tuple with the ActivateWhenValidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRepoRead) GetActivateWhenValidatedOk() (*bool, bool) {
	if o == nil || isNil(o.ActivateWhenValidated) {
		return nil, false
	}
	return o.ActivateWhenValidated, true
}

// HasActivateWhenValidated returns a boolean if a field has been set.
func (o *PolicyRepoRead) HasActivateWhenValidated() bool {
	if o != nil && !isNil(o.ActivateWhenValidated) {
		return true
	}

	return false
}

// SetActivateWhenValidated gets a reference to the given bool and assigns it to the ActivateWhenValidated field.
func (o *PolicyRepoRead) SetActivateWhenValidated(v bool) {
	o.ActivateWhenValidated = &v
}

func (o PolicyRepoRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.MainBranchName) {
		toSerialize["main_branch_name"] = o.MainBranchName
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if !isNil(o.ActivateWhenValidated) {
		toSerialize["activate_when_validated"] = o.ActivateWhenValidated
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyRepoRead struct {
	value *PolicyRepoRead
	isSet bool
}

func (v NullablePolicyRepoRead) Get() *PolicyRepoRead {
	return v.value
}

func (v *NullablePolicyRepoRead) Set(val *PolicyRepoRead) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRepoRead) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRepoRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRepoRead(val *PolicyRepoRead) *NullablePolicyRepoRead {
	return &NullablePolicyRepoRead{value: val, isSet: true}
}

func (v NullablePolicyRepoRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRepoRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
