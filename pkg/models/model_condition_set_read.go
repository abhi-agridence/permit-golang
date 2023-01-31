/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)

// ConditionSetRead struct for ConditionSetRead
type ConditionSetRead struct {
	// A unique id by which Permit will identify the condition set. The key will be used as the generated rego rule name.
	Key string `json:"key"`
	// the type of the set: UserSet or ResourceSet
	Type *ConditionSetType `json:"type,omitempty"`
	// whether the set was autogenerated by the system.
	Autogenerated *bool       `json:"autogenerated,omitempty"`
	ResourceId    *ResourceId `json:"resource_id,omitempty"`
	// Unique id of the condition set
	Id string `json:"id"`
	// Unique id of the organization that the condition set belongs to.
	OrganizationId string `json:"organization_id"`
	// Unique id of the project that the condition set belongs to.
	ProjectId string `json:"project_id"`
	// Unique id of the environment that the condition set belongs to.
	EnvironmentId string `json:"environment_id"`
	// Date and time when the condition set was created (ISO_8601 format).
	CreatedAt time.Time `json:"created_at"`
	// Date and time when the condition set was last updated/modified (ISO_8601 format).
	UpdatedAt time.Time     `json:"updated_at"`
	Resource  *ResourceRead `json:"resource,omitempty"`
	// A descriptive name for the set, i.e: 'US based employees' or 'Users behind VPN'
	Name string `json:"name"`
	// an optional longer description of the set
	Description *string `json:"description,omitempty"`
	// a boolean expression that consists of multiple conditions, with and/or logic.
	Conditions map[string]interface{} `json:"conditions,omitempty"`
}

// NewConditionSetRead instantiates a new ConditionSetRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionSetRead(key string, id string, organizationId string, projectId string, environmentId string, createdAt time.Time, updatedAt time.Time, name string) *ConditionSetRead {
	this := ConditionSetRead{}
	this.Key = key
	var type_ ConditionSetType = USERSET
	this.Type = &type_
	var autogenerated bool = false
	this.Autogenerated = &autogenerated
	this.Id = id
	this.OrganizationId = organizationId
	this.ProjectId = projectId
	this.EnvironmentId = environmentId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	return &this
}

// NewConditionSetReadWithDefaults instantiates a new ConditionSetRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionSetReadWithDefaults() *ConditionSetRead {
	this := ConditionSetRead{}
	var type_ ConditionSetType = USERSET
	this.Type = &type_
	var autogenerated bool = false
	this.Autogenerated = &autogenerated
	return &this
}

// GetKey returns the Key field value
func (o *ConditionSetRead) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ConditionSetRead) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConditionSetRead) GetType() ConditionSetType {
	if o == nil || IsNil(o.Type) {
		var ret ConditionSetType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetTypeOk() (*ConditionSetType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConditionSetRead) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConditionSetType and assigns it to the Type field.
func (o *ConditionSetRead) SetType(v ConditionSetType) {
	o.Type = &v
}

// GetAutogenerated returns the Autogenerated field value if set, zero value otherwise.
func (o *ConditionSetRead) GetAutogenerated() bool {
	if o == nil || IsNil(o.Autogenerated) {
		var ret bool
		return ret
	}
	return *o.Autogenerated
}

// GetAutogeneratedOk returns a tuple with the Autogenerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetAutogeneratedOk() (*bool, bool) {
	if o == nil || IsNil(o.Autogenerated) {
		return nil, false
	}
	return o.Autogenerated, true
}

// HasAutogenerated returns a boolean if a field has been set.
func (o *ConditionSetRead) HasAutogenerated() bool {
	if o != nil && !IsNil(o.Autogenerated) {
		return true
	}

	return false
}

// SetAutogenerated gets a reference to the given bool and assigns it to the Autogenerated field.
func (o *ConditionSetRead) SetAutogenerated(v bool) {
	o.Autogenerated = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *ConditionSetRead) GetResourceId() ResourceId {
	if o == nil || IsNil(o.ResourceId) {
		var ret ResourceId
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetResourceIdOk() (*ResourceId, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *ConditionSetRead) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given ResourceId and assigns it to the ResourceId field.
func (o *ConditionSetRead) SetResourceId(v ResourceId) {
	o.ResourceId = &v
}

// GetId returns the Id field value
func (o *ConditionSetRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConditionSetRead) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *ConditionSetRead) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *ConditionSetRead) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *ConditionSetRead) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ConditionSetRead) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *ConditionSetRead) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *ConditionSetRead) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConditionSetRead) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConditionSetRead) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ConditionSetRead) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ConditionSetRead) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ConditionSetRead) GetResource() ResourceRead {
	if o == nil || IsNil(o.Resource) {
		var ret ResourceRead
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetResourceOk() (*ResourceRead, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ConditionSetRead) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ResourceRead and assigns it to the Resource field.
func (o *ConditionSetRead) SetResource(v ResourceRead) {
	o.Resource = &v
}

// GetName returns the Name field value
func (o *ConditionSetRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConditionSetRead) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConditionSetRead) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConditionSetRead) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConditionSetRead) SetDescription(v string) {
	o.Description = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ConditionSetRead) GetConditions() map[string]interface{} {
	if o == nil || IsNil(o.Conditions) {
		var ret map[string]interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionSetRead) GetConditionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return map[string]interface{}{}, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ConditionSetRead) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given map[string]interface{} and assigns it to the Conditions field.
func (o *ConditionSetRead) SetConditions(v map[string]interface{}) {
	o.Conditions = v
}

func (o ConditionSetRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Autogenerated) {
		toSerialize["autogenerated"] = o.Autogenerated
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resource_id"] = o.ResourceId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	if true {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return json.Marshal(toSerialize)
}

type NullableConditionSetRead struct {
	value *ConditionSetRead
	isSet bool
}

func (v NullableConditionSetRead) Get() *ConditionSetRead {
	return v.value
}

func (v *NullableConditionSetRead) Set(val *ConditionSetRead) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionSetRead) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionSetRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionSetRead(val *ConditionSetRead) *NullableConditionSetRead {
	return &NullableConditionSetRead{value: val, isSet: true}
}

func (v NullableConditionSetRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionSetRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}