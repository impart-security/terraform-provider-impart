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

// checks if the LogBindings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogBindings{}

// LogBindings struct for LogBindings
type LogBindings struct {
	// A list of log bindings.
	Items []LogBinding   `json:"items"`
	Meta  CollectionMeta `json:"meta"`
}

type _LogBindings LogBindings

// NewLogBindings instantiates a new LogBindings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogBindings(items []LogBinding, meta CollectionMeta) *LogBindings {
	this := LogBindings{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewLogBindingsWithDefaults instantiates a new LogBindings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogBindingsWithDefaults() *LogBindings {
	this := LogBindings{}
	return &this
}

// GetItems returns the Items field value
func (o *LogBindings) GetItems() []LogBinding {
	if o == nil {
		var ret []LogBinding
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *LogBindings) GetItemsOk() ([]LogBinding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *LogBindings) SetItems(v []LogBinding) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *LogBindings) GetMeta() CollectionMeta {
	if o == nil {
		var ret CollectionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *LogBindings) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *LogBindings) SetMeta(v CollectionMeta) {
	o.Meta = v
}

func (o LogBindings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogBindings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *LogBindings) UnmarshalJSON(data []byte) (err error) {
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

	varLogBindings := _LogBindings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogBindings)

	if err != nil {
		return err
	}

	*o = LogBindings(varLogBindings)

	return err
}

type NullableLogBindings struct {
	value *LogBindings
	isSet bool
}

func (v NullableLogBindings) Get() *LogBindings {
	return v.value
}

func (v *NullableLogBindings) Set(val *LogBindings) {
	v.value = val
	v.isSet = true
}

func (v NullableLogBindings) IsSet() bool {
	return v.isSet
}

func (v *NullableLogBindings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogBindings(val *LogBindings) *NullableLogBindings {
	return &NullableLogBindings{value: val, isSet: true}
}

func (v NullableLogBindings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogBindings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}