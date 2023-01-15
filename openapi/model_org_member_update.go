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

// OrgMemberUpdate struct for OrgMemberUpdate
type OrgMemberUpdate struct {
	// Custom permit.io dashboard settings, such as preferred theme, etc.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// updates the onboarding step (optional)
	OnboardingStep *OnboardingStep `json:"onboarding_step,omitempty"`
}

// NewOrgMemberUpdate instantiates a new OrgMemberUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgMemberUpdate() *OrgMemberUpdate {
	this := OrgMemberUpdate{}
	return &this
}

// NewOrgMemberUpdateWithDefaults instantiates a new OrgMemberUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgMemberUpdateWithDefaults() *OrgMemberUpdate {
	this := OrgMemberUpdate{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *OrgMemberUpdate) GetSettings() map[string]interface{} {
	if o == nil || isNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMemberUpdate) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *OrgMemberUpdate) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *OrgMemberUpdate) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetOnboardingStep returns the OnboardingStep field value if set, zero value otherwise.
func (o *OrgMemberUpdate) GetOnboardingStep() OnboardingStep {
	if o == nil || isNil(o.OnboardingStep) {
		var ret OnboardingStep
		return ret
	}
	return *o.OnboardingStep
}

// GetOnboardingStepOk returns a tuple with the OnboardingStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMemberUpdate) GetOnboardingStepOk() (*OnboardingStep, bool) {
	if o == nil || isNil(o.OnboardingStep) {
		return nil, false
	}
	return o.OnboardingStep, true
}

// HasOnboardingStep returns a boolean if a field has been set.
func (o *OrgMemberUpdate) HasOnboardingStep() bool {
	if o != nil && !isNil(o.OnboardingStep) {
		return true
	}

	return false
}

// SetOnboardingStep gets a reference to the given OnboardingStep and assigns it to the OnboardingStep field.
func (o *OrgMemberUpdate) SetOnboardingStep(v OnboardingStep) {
	o.OnboardingStep = &v
}

func (o OrgMemberUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.OnboardingStep) {
		toSerialize["onboarding_step"] = o.OnboardingStep
	}
	return json.Marshal(toSerialize)
}

type NullableOrgMemberUpdate struct {
	value *OrgMemberUpdate
	isSet bool
}

func (v NullableOrgMemberUpdate) Get() *OrgMemberUpdate {
	return v.value
}

func (v *NullableOrgMemberUpdate) Set(val *OrgMemberUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgMemberUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgMemberUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgMemberUpdate(val *OrgMemberUpdate) *NullableOrgMemberUpdate {
	return &NullableOrgMemberUpdate{value: val, isSet: true}
}

func (v NullableOrgMemberUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgMemberUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
