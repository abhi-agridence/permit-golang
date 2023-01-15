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

// TenantRead struct for TenantRead
type TenantRead struct {
	// A unique id by which Permit will identify the tenant. The tenant key must be url-friendly (slugified).
	Key string `json:"key"`
	// Unique id of the tenant
	Id string `json:"id"`
	// Unique id of the organization that the tenant belongs to.
	OrganizationId string `json:"organization_id"`
	// Unique id of the project that the tenant belongs to.
	ProjectId string `json:"project_id"`
	// Unique id of the environment that the tenant belongs to.
	EnvironmentId string `json:"environment_id"`
	// Date and time when the tenant was created (ISO_8601 format).
	CreatedAt time.Time `json:"created_at"`
	// Date and time when the tenant was last updated/modified (ISO_8601 format).
	UpdatedAt time.Time `json:"updated_at"`
	// Date and time when the tenant was last active (ISO_8601 format). In other words, this is the last time a permission check was done on a resource belonging to this tenant.
	LastActionAt time.Time `json:"last_action_at"`
	// A descriptive name for the tenant
	Name string `json:"name"`
	// an optional longer description of the tenant
	Description *string `json:"description,omitempty"`
	// Arbitraty tenant attributes that will be used to enforce attribute-based access control policies.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// NewTenantRead instantiates a new TenantRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRead(key string, id string, organizationId string, projectId string, environmentId string, createdAt time.Time, updatedAt time.Time, lastActionAt time.Time, name string) *TenantRead {
	this := TenantRead{}
	this.Key = key
	this.Id = id
	this.OrganizationId = organizationId
	this.ProjectId = projectId
	this.EnvironmentId = environmentId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.LastActionAt = lastActionAt
	this.Name = name
	return &this
}

// NewTenantReadWithDefaults instantiates a new TenantRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantReadWithDefaults() *TenantRead {
	this := TenantRead{}
	return &this
}

// GetKey returns the Key field value
func (o *TenantRead) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TenantRead) SetKey(v string) {
	o.Key = v
}

// GetId returns the Id field value
func (o *TenantRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenantRead) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *TenantRead) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *TenantRead) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *TenantRead) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TenantRead) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *TenantRead) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *TenantRead) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TenantRead) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TenantRead) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TenantRead) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TenantRead) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLastActionAt returns the LastActionAt field value
func (o *TenantRead) GetLastActionAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastActionAt
}

// GetLastActionAtOk returns a tuple with the LastActionAt field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetLastActionAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastActionAt, true
}

// SetLastActionAt sets field value
func (o *TenantRead) SetLastActionAt(v time.Time) {
	o.LastActionAt = v
}

// GetName returns the Name field value
func (o *TenantRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantRead) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantRead) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TenantRead) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRead) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TenantRead) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TenantRead) SetDescription(v string) {
	o.Description = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TenantRead) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRead) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TenantRead) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *TenantRead) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TenantRead) MarshalJSON() ([]byte, error) {
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
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["last_action_at"] = o.LastActionAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableTenantRead struct {
	value *TenantRead
	isSet bool
}

func (v NullableTenantRead) Get() *TenantRead {
	return v.value
}

func (v *NullableTenantRead) Set(val *TenantRead) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRead) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRead(val *TenantRead) *NullableTenantRead {
	return &NullableTenantRead{value: val, isSet: true}
}

func (v NullableTenantRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
