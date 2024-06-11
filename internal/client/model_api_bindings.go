/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApiBindings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBindings{}

// ApiBindings struct for ApiBindings
type ApiBindings struct {
	// A list of API bindings.
	Items []ApiBinding   `json:"items"`
	Meta  CollectionMeta `json:"meta"`
}

type _ApiBindings ApiBindings

// NewApiBindings instantiates a new ApiBindings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBindings(items []ApiBinding, meta CollectionMeta) *ApiBindings {
	this := ApiBindings{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewApiBindingsWithDefaults instantiates a new ApiBindings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBindingsWithDefaults() *ApiBindings {
	this := ApiBindings{}
	return &this
}

// GetItems returns the Items field value
func (o *ApiBindings) GetItems() []ApiBinding {
	if o == nil {
		var ret []ApiBinding
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ApiBindings) GetItemsOk() ([]ApiBinding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ApiBindings) SetItems(v []ApiBinding) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *ApiBindings) GetMeta() CollectionMeta {
	if o == nil {
		var ret CollectionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ApiBindings) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ApiBindings) SetMeta(v CollectionMeta) {
	o.Meta = v
}

func (o ApiBindings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBindings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *ApiBindings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"meta",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiBindings := _ApiBindings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiBindings)

	if err != nil {
		return err
	}

	*o = ApiBindings(varApiBindings)

	return err
}

type NullableApiBindings struct {
	value *ApiBindings
	isSet bool
}

func (v NullableApiBindings) Get() *ApiBindings {
	return v.value
}

func (v *NullableApiBindings) Set(val *ApiBindings) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBindings) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBindings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBindings(val *ApiBindings) *NullableApiBindings {
	return &NullableApiBindings{value: val, isSet: true}
}

func (v NullableApiBindings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBindings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
