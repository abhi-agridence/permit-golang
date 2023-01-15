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

// ResourceAttributeCreate struct for ResourceAttributeCreate
type ResourceAttributeCreate struct {
	// A URL-friendly name of the attribute (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the attribute.
	Key string `json:"key"`
	// The type of the attribute, we currently support: `bool`, `number` (ints, floats), `time` (a timestamp), `string`, and `json`.
	Type AttributeType `json:"type"`
	// An optional longer description of what this attribute respresents in your system
	Description *string `json:"description,omitempty"`
}

// NewResourceAttributeCreate instantiates a new ResourceAttributeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAttributeCreate(key string, type_ AttributeType) *ResourceAttributeCreate {
	this := ResourceAttributeCreate{}
	this.Key = key
	this.Type = type_
	return &this
}

// NewResourceAttributeCreateWithDefaults instantiates a new ResourceAttributeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAttributeCreateWithDefaults() *ResourceAttributeCreate {
	this := ResourceAttributeCreate{}
	return &this
}

// GetKey returns the Key field value
func (o *ResourceAttributeCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResourceAttributeCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResourceAttributeCreate) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value
func (o *ResourceAttributeCreate) GetType() AttributeType {
	if o == nil {
		var ret AttributeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceAttributeCreate) GetTypeOk() (*AttributeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceAttributeCreate) SetType(v AttributeType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceAttributeCreate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAttributeCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceAttributeCreate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceAttributeCreate) SetDescription(v string) {
	o.Description = &v
}

func (o ResourceAttributeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableResourceAttributeCreate struct {
	value *ResourceAttributeCreate
	isSet bool
}

func (v NullableResourceAttributeCreate) Get() *ResourceAttributeCreate {
	return v.value
}

func (v *NullableResourceAttributeCreate) Set(val *ResourceAttributeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAttributeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAttributeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAttributeCreate(val *ResourceAttributeCreate) *NullableResourceAttributeCreate {
	return &NullableResourceAttributeCreate{value: val, isSet: true}
}

func (v NullableResourceAttributeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAttributeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
