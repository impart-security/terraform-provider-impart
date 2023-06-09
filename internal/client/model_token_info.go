/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the TokenInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenInfo{}

// TokenInfo struct for TokenInfo
type TokenInfo struct {
	// ID of the user.
	UserId string `json:"user_id"`
	// ID of the org.
	OrgId string `json:"org_id"`
}

// NewTokenInfo instantiates a new TokenInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenInfo(userId string, orgId string) *TokenInfo {
	this := TokenInfo{}
	this.UserId = userId
	this.OrgId = orgId
	return &this
}

// NewTokenInfoWithDefaults instantiates a new TokenInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenInfoWithDefaults() *TokenInfo {
	this := TokenInfo{}
	return &this
}

// GetUserId returns the UserId field value
func (o *TokenInfo) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TokenInfo) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TokenInfo) SetUserId(v string) {
	o.UserId = v
}

// GetOrgId returns the OrgId field value
func (o *TokenInfo) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *TokenInfo) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *TokenInfo) SetOrgId(v string) {
	o.OrgId = v
}

func (o TokenInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["org_id"] = o.OrgId
	return toSerialize, nil
}

type NullableTokenInfo struct {
	value *TokenInfo
	isSet bool
}

func (v NullableTokenInfo) Get() *TokenInfo {
	return v.value
}

func (v *NullableTokenInfo) Set(val *TokenInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenInfo(val *TokenInfo) *NullableTokenInfo {
	return &NullableTokenInfo{value: val, isSet: true}
}

func (v NullableTokenInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
