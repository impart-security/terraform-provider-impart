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

// checks if the Lists type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lists{}

// Lists struct for Lists
type Lists struct {
	// A list of lists.
	Items []ListsItemsInner `json:"items"`
	Meta  CollectionMeta    `json:"meta"`
}

type _Lists Lists

// NewLists instantiates a new Lists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLists(items []ListsItemsInner, meta CollectionMeta) *Lists {
	this := Lists{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewListsWithDefaults instantiates a new Lists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsWithDefaults() *Lists {
	this := Lists{}
	return &this
}

// GetItems returns the Items field value
func (o *Lists) GetItems() []ListsItemsInner {
	if o == nil {
		var ret []ListsItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Lists) GetItemsOk() ([]ListsItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Lists) SetItems(v []ListsItemsInner) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *Lists) GetMeta() CollectionMeta {
	if o == nil {
		var ret CollectionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *Lists) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *Lists) SetMeta(v CollectionMeta) {
	o.Meta = v
}

func (o Lists) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lists) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *Lists) UnmarshalJSON(data []byte) (err error) {
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

	varLists := _Lists{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLists)

	if err != nil {
		return err
	}

	*o = Lists(varLists)

	return err
}

type NullableLists struct {
	value *Lists
	isSet bool
}

func (v NullableLists) Get() *Lists {
	return v.value
}

func (v *NullableLists) Set(val *Lists) {
	v.value = val
	v.isSet = true
}

func (v NullableLists) IsSet() bool {
	return v.isSet
}

func (v *NullableLists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLists(val *Lists) *NullableLists {
	return &NullableLists{value: val, isSet: true}
}

func (v NullableLists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
