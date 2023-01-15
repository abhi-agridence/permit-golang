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

// AuthnMeAPIKeyRead struct for AuthnMeAPIKeyRead
type AuthnMeAPIKeyRead struct {
	ActorType  *string         `json:"actor_type,omitempty"`
	Id         string          `json:"id"`
	ObjectType MemberAccessObj `json:"object_type"`
	OwnerType  APIKeyOwnerType `json:"owner_type"`
	OrgId      string          `json:"org_id"`
	ProjectId  *string         `json:"project_id,omitempty"`
	EnvId      *string         `json:"env_id,omitempty"`
}

// NewAuthnMeAPIKeyRead instantiates a new AuthnMeAPIKeyRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthnMeAPIKeyRead(id string, objectType MemberAccessObj, ownerType APIKeyOwnerType, orgId string) *AuthnMeAPIKeyRead {
	this := AuthnMeAPIKeyRead{}
	var actorType string = "api_key"
	this.ActorType = &actorType
	this.Id = id
	this.ObjectType = objectType
	this.OwnerType = ownerType
	this.OrgId = orgId
	return &this
}

// NewAuthnMeAPIKeyReadWithDefaults instantiates a new AuthnMeAPIKeyRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthnMeAPIKeyReadWithDefaults() *AuthnMeAPIKeyRead {
	this := AuthnMeAPIKeyRead{}
	var actorType string = "api_key"
	this.ActorType = &actorType
	return &this
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *AuthnMeAPIKeyRead) GetActorType() string {
	if o == nil || isNil(o.ActorType) {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetActorTypeOk() (*string, bool) {
	if o == nil || isNil(o.ActorType) {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *AuthnMeAPIKeyRead) HasActorType() bool {
	if o != nil && !isNil(o.ActorType) {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *AuthnMeAPIKeyRead) SetActorType(v string) {
	o.ActorType = &v
}

// GetId returns the Id field value
func (o *AuthnMeAPIKeyRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthnMeAPIKeyRead) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *AuthnMeAPIKeyRead) GetObjectType() MemberAccessObj {
	if o == nil {
		var ret MemberAccessObj
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetObjectTypeOk() (*MemberAccessObj, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AuthnMeAPIKeyRead) SetObjectType(v MemberAccessObj) {
	o.ObjectType = v
}

// GetOwnerType returns the OwnerType field value
func (o *AuthnMeAPIKeyRead) GetOwnerType() APIKeyOwnerType {
	if o == nil {
		var ret APIKeyOwnerType
		return ret
	}

	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetOwnerTypeOk() (*APIKeyOwnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value
func (o *AuthnMeAPIKeyRead) SetOwnerType(v APIKeyOwnerType) {
	o.OwnerType = v
}

// GetOrgId returns the OrgId field value
func (o *AuthnMeAPIKeyRead) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *AuthnMeAPIKeyRead) SetOrgId(v string) {
	o.OrgId = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AuthnMeAPIKeyRead) GetProjectId() string {
	if o == nil || isNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetProjectIdOk() (*string, bool) {
	if o == nil || isNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AuthnMeAPIKeyRead) HasProjectId() bool {
	if o != nil && !isNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AuthnMeAPIKeyRead) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetEnvId returns the EnvId field value if set, zero value otherwise.
func (o *AuthnMeAPIKeyRead) GetEnvId() string {
	if o == nil || isNil(o.EnvId) {
		var ret string
		return ret
	}
	return *o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnMeAPIKeyRead) GetEnvIdOk() (*string, bool) {
	if o == nil || isNil(o.EnvId) {
		return nil, false
	}
	return o.EnvId, true
}

// HasEnvId returns a boolean if a field has been set.
func (o *AuthnMeAPIKeyRead) HasEnvId() bool {
	if o != nil && !isNil(o.EnvId) {
		return true
	}

	return false
}

// SetEnvId gets a reference to the given string and assigns it to the EnvId field.
func (o *AuthnMeAPIKeyRead) SetEnvId(v string) {
	o.EnvId = &v
}

func (o AuthnMeAPIKeyRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActorType) {
		toSerialize["actor_type"] = o.ActorType
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["object_type"] = o.ObjectType
	}
	if true {
		toSerialize["owner_type"] = o.OwnerType
	}
	if true {
		toSerialize["org_id"] = o.OrgId
	}
	if !isNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !isNil(o.EnvId) {
		toSerialize["env_id"] = o.EnvId
	}
	return json.Marshal(toSerialize)
}

type NullableAuthnMeAPIKeyRead struct {
	value *AuthnMeAPIKeyRead
	isSet bool
}

func (v NullableAuthnMeAPIKeyRead) Get() *AuthnMeAPIKeyRead {
	return v.value
}

func (v *NullableAuthnMeAPIKeyRead) Set(val *AuthnMeAPIKeyRead) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthnMeAPIKeyRead) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthnMeAPIKeyRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthnMeAPIKeyRead(val *AuthnMeAPIKeyRead) *NullableAuthnMeAPIKeyRead {
	return &NullableAuthnMeAPIKeyRead{value: val, isSet: true}
}

func (v NullableAuthnMeAPIKeyRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthnMeAPIKeyRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
