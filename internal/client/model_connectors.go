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

// checks if the Connectors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connectors{}

// Connectors struct for Connectors
type Connectors struct {
	// A list of connectors.
	Items []Connector    `json:"items"`
	Meta  CollectionMeta `json:"meta"`
}

type _Connectors Connectors

// NewConnectors instantiates a new Connectors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectors(items []Connector, meta CollectionMeta) *Connectors {
	this := Connectors{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewConnectorsWithDefaults instantiates a new Connectors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorsWithDefaults() *Connectors {
	this := Connectors{}
	return &this
}

// GetItems returns the Items field value
func (o *Connectors) GetItems() []Connector {
	if o == nil {
		var ret []Connector
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Connectors) GetItemsOk() ([]Connector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Connectors) SetItems(v []Connector) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *Connectors) GetMeta() CollectionMeta {
	if o == nil {
		var ret CollectionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *Connectors) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *Connectors) SetMeta(v CollectionMeta) {
	o.Meta = v
}

func (o Connectors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connectors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *Connectors) UnmarshalJSON(data []byte) (err error) {
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

	varConnectors := _Connectors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectors)

	if err != nil {
		return err
	}

	*o = Connectors(varConnectors)

	return err
}

type NullableConnectors struct {
	value *Connectors
	isSet bool
}

func (v NullableConnectors) Get() *Connectors {
	return v.value
}

func (v *NullableConnectors) Set(val *Connectors) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectors) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectors(val *Connectors) *NullableConnectors {
	return &NullableConnectors{value: val, isSet: true}
}

func (v NullableConnectors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}