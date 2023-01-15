/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// EnvironmentRead struct for EnvironmentRead
type EnvironmentRead struct {
	// A URL-friendly name of the environment (i.e: slug). You will be able to query later using this key instead of the id (UUID) of the environment.
	Key string `json:"key"`
	// Unique id of the environment
	Id string `json:"id"`
	// Unique id of the organization that the environment belongs to.
	OrganizationId string `json:"organization_id"`
	// Unique id of the project that the environment belongs to.
	ProjectId string `json:"project_id"`
	// Date and time when the environment was created (ISO_8601 format).
	CreatedAt time.Time `json:"created_at"`
	// Date and time when the environment was last updated/modified (ISO_8601 format).
	UpdatedAt time.Time `json:"updated_at"`
	// The name of the environment
	Name string `json:"name"`
	// an optional longer description of the environment
	Description *string `json:"description,omitempty"`
	// when using gitops feature, an optional branch name for the environment
	CustomBranchName *string `json:"custom_branch_name,omitempty"`
}

// NewEnvironmentRead instantiates a new EnvironmentRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentRead(key string, id string, organizationId string, projectId string, createdAt time.Time, updatedAt time.Time, name string) *EnvironmentRead {
	this := EnvironmentRead{}
	this.Key = key
	this.Id = id
	this.OrganizationId = organizationId
	this.ProjectId = projectId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	return &this
}

// NewEnvironmentReadWithDefaults instantiates a new EnvironmentRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentReadWithDefaults() *EnvironmentRead {
	this := EnvironmentRead{}
	return &this
}

// GetKey returns the Key field value
func (o *EnvironmentRead) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentRead) SetKey(v string) {
	o.Key = v
}

// GetId returns the Id field value
func (o *EnvironmentRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentRead) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *EnvironmentRead) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *EnvironmentRead) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *EnvironmentRead) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *EnvironmentRead) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentRead) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentRead) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EnvironmentRead) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EnvironmentRead) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *EnvironmentRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentRead) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentRead) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentRead) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentRead) SetDescription(v string) {
	o.Description = &v
}

// GetCustomBranchName returns the CustomBranchName field value if set, zero value otherwise.
func (o *EnvironmentRead) GetCustomBranchName() string {
	if o == nil || isNil(o.CustomBranchName) {
		var ret string
		return ret
	}
	return *o.CustomBranchName
}

// GetCustomBranchNameOk returns a tuple with the CustomBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRead) GetCustomBranchNameOk() (*string, bool) {
	if o == nil || isNil(o.CustomBranchName) {
		return nil, false
	}
	return o.CustomBranchName, true
}

// HasCustomBranchName returns a boolean if a field has been set.
func (o *EnvironmentRead) HasCustomBranchName() bool {
	if o != nil && !isNil(o.CustomBranchName) {
		return true
	}

	return false
}

// SetCustomBranchName gets a reference to the given string and assigns it to the CustomBranchName field.
func (o *EnvironmentRead) SetCustomBranchName(v string) {
	o.CustomBranchName = &v
}

func (o EnvironmentRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
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
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CustomBranchName) {
		toSerialize["custom_branch_name"] = o.CustomBranchName
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentRead struct {
	value *EnvironmentRead
	isSet bool
}

func (v NullableEnvironmentRead) Get() *EnvironmentRead {
	return v.value
}

func (v *NullableEnvironmentRead) Set(val *EnvironmentRead) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRead) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRead(val *EnvironmentRead) *NullableEnvironmentRead {
	return &NullableEnvironmentRead{value: val, isSet: true}
}

func (v NullableEnvironmentRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
