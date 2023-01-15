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

// ResourceCreate struct for ResourceCreate
type ResourceCreate struct {
	// A URL-friendly name of the resource (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the resource.
	Key string `json:"key"`
	// The name of the resource
	Name string `json:"name"`
	// The [URN](https://en.wikipedia.org/wiki/Uniform_Resource_Name) (Uniform Resource Name) of the resource
	Urn *string `json:"urn,omitempty"`
	// An optional longer description of what this resource respresents in your system
	Description *string `json:"description,omitempty"`
	//          A actions definition block, typically contained within a resource type definition block.         The actions represents the ways you can interact with a protected resource.
	Actions map[string]ActionBlockEditable `json:"actions"`
	// Attributes that each resource of this type defines, and can be used in your ABAC policies.
	Attributes *map[string]AttributeBlockEditable `json:"attributes,omitempty"`
}

// NewResourceCreate instantiates a new ResourceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCreate(key string, name string, actions map[string]ActionBlockEditable) *ResourceCreate {
	this := ResourceCreate{}
	this.Key = key
	this.Name = name
	this.Actions = actions
	return &this
}

// NewResourceCreateWithDefaults instantiates a new ResourceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCreateWithDefaults() *ResourceCreate {
	this := ResourceCreate{}
	return &this
}

// GetKey returns the Key field value
func (o *ResourceCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResourceCreate) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ResourceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceCreate) SetName(v string) {
	o.Name = v
}

// GetUrn returns the Urn field value if set, zero value otherwise.
func (o *ResourceCreate) GetUrn() string {
	if o == nil || isNil(o.Urn) {
		var ret string
		return ret
	}
	return *o.Urn
}

// GetUrnOk returns a tuple with the Urn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetUrnOk() (*string, bool) {
	if o == nil || isNil(o.Urn) {
		return nil, false
	}
	return o.Urn, true
}

// HasUrn returns a boolean if a field has been set.
func (o *ResourceCreate) HasUrn() bool {
	if o != nil && !isNil(o.Urn) {
		return true
	}

	return false
}

// SetUrn gets a reference to the given string and assigns it to the Urn field.
func (o *ResourceCreate) SetUrn(v string) {
	o.Urn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceCreate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceCreate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceCreate) SetDescription(v string) {
	o.Description = &v
}

// GetActions returns the Actions field value
func (o *ResourceCreate) GetActions() map[string]ActionBlockEditable {
	if o == nil {
		var ret map[string]ActionBlockEditable
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetActionsOk() (*map[string]ActionBlockEditable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *ResourceCreate) SetActions(v map[string]ActionBlockEditable) {
	o.Actions = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ResourceCreate) GetAttributes() map[string]AttributeBlockEditable {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]AttributeBlockEditable
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetAttributesOk() (*map[string]AttributeBlockEditable, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ResourceCreate) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]AttributeBlockEditable and assigns it to the Attributes field.
func (o *ResourceCreate) SetAttributes(v map[string]AttributeBlockEditable) {
	o.Attributes = &v
}

func (o ResourceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Urn) {
		toSerialize["urn"] = o.Urn
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableResourceCreate struct {
	value *ResourceCreate
	isSet bool
}

func (v NullableResourceCreate) Get() *ResourceCreate {
	return v.value
}

func (v *NullableResourceCreate) Set(val *ResourceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCreate(val *ResourceCreate) *NullableResourceCreate {
	return &NullableResourceCreate{value: val, isSet: true}
}

func (v NullableResourceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
