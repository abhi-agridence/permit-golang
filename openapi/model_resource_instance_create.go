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

// ResourceInstanceCreate struct for ResourceInstanceCreate
type ResourceInstanceCreate struct {
	// A unique identifier by which Permit will identify the resource instance for permission checks. You will later pass this identifier to the `permit.check()` API. A key can be anything: for example the resource db id, a url slug, a UUID or anything else as long as it's unique on your end. The resource instance key must be url-friendly.
	Key string `json:"key"`
	// the *key* of the tenant that this resource belongs to, used to enforce tenant boundaries in multi-tenant apps.
	Tenant *string `json:"tenant,omitempty"`
	// the *key* of the resource (type) of this resource instance. For example: if this resource instance is the annual budget document, the key of the resource might be `document`.
	Resource string `json:"resource"`
	// Arbitraty resource attributes that will be used to enforce attribute-based access control policies.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// NewResourceInstanceCreate instantiates a new ResourceInstanceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceInstanceCreate(key string, resource string) *ResourceInstanceCreate {
	this := ResourceInstanceCreate{}
	this.Key = key
	this.Resource = resource
	return &this
}

// NewResourceInstanceCreateWithDefaults instantiates a new ResourceInstanceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceInstanceCreateWithDefaults() *ResourceInstanceCreate {
	this := ResourceInstanceCreate{}
	return &this
}

// GetKey returns the Key field value
func (o *ResourceInstanceCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResourceInstanceCreate) SetKey(v string) {
	o.Key = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *ResourceInstanceCreate) GetTenant() string {
	if o == nil || isNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceCreate) GetTenantOk() (*string, bool) {
	if o == nil || isNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *ResourceInstanceCreate) HasTenant() bool {
	if o != nil && !isNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *ResourceInstanceCreate) SetTenant(v string) {
	o.Tenant = &v
}

// GetResource returns the Resource field value
func (o *ResourceInstanceCreate) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceCreate) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ResourceInstanceCreate) SetResource(v string) {
	o.Resource = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ResourceInstanceCreate) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceCreate) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ResourceInstanceCreate) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ResourceInstanceCreate) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o ResourceInstanceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableResourceInstanceCreate struct {
	value *ResourceInstanceCreate
	isSet bool
}

func (v NullableResourceInstanceCreate) Get() *ResourceInstanceCreate {
	return v.value
}

func (v *NullableResourceInstanceCreate) Set(val *ResourceInstanceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceInstanceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceInstanceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceInstanceCreate(val *ResourceInstanceCreate) *NullableResourceInstanceCreate {
	return &NullableResourceInstanceCreate{value: val, isSet: true}
}

func (v NullableResourceInstanceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceInstanceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
