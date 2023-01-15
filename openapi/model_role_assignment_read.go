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

// RoleAssignmentRead struct for RoleAssignmentRead
type RoleAssignmentRead struct {
	// Unique id of the role assignment
	Id string `json:"id"`
	// the user the role is assigned to
	User string `json:"user"`
	// the role that is assigned
	Role string `json:"role"`
	// the tenant the role is associated with
	Tenant string `json:"tenant"`
	// Unique id of the user
	UserId string `json:"user_id"`
	// Unique id of the role
	RoleId string `json:"role_id"`
	// Unique id of the tenant
	TenantId string `json:"tenant_id"`
	// Unique id of the organization that the role assignment belongs to.
	OrganizationId string `json:"organization_id"`
	// Unique id of the project that the role assignment belongs to.
	ProjectId string `json:"project_id"`
	// Unique id of the environment that the role assignment belongs to.
	EnvironmentId string `json:"environment_id"`
	// Date and time when the role assignment was created (ISO_8601 format).
	CreatedAt time.Time `json:"created_at"`
}

// NewRoleAssignmentRead instantiates a new RoleAssignmentRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentRead(id string, user string, role string, tenant string, userId string, roleId string, tenantId string, organizationId string, projectId string, environmentId string, createdAt time.Time) *RoleAssignmentRead {
	this := RoleAssignmentRead{}
	this.Id = id
	this.User = user
	this.Role = role
	this.Tenant = tenant
	this.UserId = userId
	this.RoleId = roleId
	this.TenantId = tenantId
	this.OrganizationId = organizationId
	this.ProjectId = projectId
	this.EnvironmentId = environmentId
	this.CreatedAt = createdAt
	return &this
}

// NewRoleAssignmentReadWithDefaults instantiates a new RoleAssignmentRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentReadWithDefaults() *RoleAssignmentRead {
	this := RoleAssignmentRead{}
	return &this
}

// GetId returns the Id field value
func (o *RoleAssignmentRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleAssignmentRead) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *RoleAssignmentRead) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RoleAssignmentRead) SetUser(v string) {
	o.User = v
}

// GetRole returns the Role field value
func (o *RoleAssignmentRead) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleAssignmentRead) SetRole(v string) {
	o.Role = v
}

// GetTenant returns the Tenant field value
func (o *RoleAssignmentRead) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *RoleAssignmentRead) SetTenant(v string) {
	o.Tenant = v
}

// GetUserId returns the UserId field value
func (o *RoleAssignmentRead) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *RoleAssignmentRead) SetUserId(v string) {
	o.UserId = v
}

// GetRoleId returns the RoleId field value
func (o *RoleAssignmentRead) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *RoleAssignmentRead) SetRoleId(v string) {
	o.RoleId = v
}

// GetTenantId returns the TenantId field value
func (o *RoleAssignmentRead) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RoleAssignmentRead) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *RoleAssignmentRead) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *RoleAssignmentRead) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *RoleAssignmentRead) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RoleAssignmentRead) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *RoleAssignmentRead) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *RoleAssignmentRead) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RoleAssignmentRead) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRead) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RoleAssignmentRead) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o RoleAssignmentRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["role_id"] = o.RoleId
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
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
	return json.Marshal(toSerialize)
}

type NullableRoleAssignmentRead struct {
	value *RoleAssignmentRead
	isSet bool
}

func (v NullableRoleAssignmentRead) Get() *RoleAssignmentRead {
	return v.value
}

func (v *NullableRoleAssignmentRead) Set(val *RoleAssignmentRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentRead(val *RoleAssignmentRead) *NullableRoleAssignmentRead {
	return &NullableRoleAssignmentRead{value: val, isSet: true}
}

func (v NullableRoleAssignmentRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
