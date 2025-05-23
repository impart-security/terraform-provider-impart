/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
	"fmt"
)

// checks if the ExternalLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalLinks{}

// ExternalLinks struct for ExternalLinks
type ExternalLinks struct {
	// A list of external links.
	Items                []ExternalLink `json:"items"`
	Meta                 CollectionMeta `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _ExternalLinks ExternalLinks

// NewExternalLinks instantiates a new ExternalLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalLinks(items []ExternalLink, meta CollectionMeta) *ExternalLinks {
	this := ExternalLinks{}
	this.Items = items
	this.Meta = meta
	return &this
}

// NewExternalLinksWithDefaults instantiates a new ExternalLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalLinksWithDefaults() *ExternalLinks {
	this := ExternalLinks{}
	return &this
}

// GetItems returns the Items field value
func (o *ExternalLinks) GetItems() []ExternalLink {
	if o == nil {
		var ret []ExternalLink
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ExternalLinks) GetItemsOk() ([]ExternalLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ExternalLinks) SetItems(v []ExternalLink) {
	o.Items = v
}

// GetMeta returns the Meta field value
func (o *ExternalLinks) GetMeta() CollectionMeta {
	if o == nil {
		var ret CollectionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ExternalLinks) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ExternalLinks) SetMeta(v CollectionMeta) {
	o.Meta = v
}

func (o ExternalLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalLinks) UnmarshalJSON(data []byte) (err error) {
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

	varExternalLinks := _ExternalLinks{}

	err = json.Unmarshal(data, &varExternalLinks)

	if err != nil {
		return err
	}

	*o = ExternalLinks(varExternalLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalLinks struct {
	value *ExternalLinks
	isSet bool
}

func (v NullableExternalLinks) Get() *ExternalLinks {
	return v.value
}

func (v *NullableExternalLinks) Set(val *ExternalLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalLinks(val *ExternalLinks) *NullableExternalLinks {
	return &NullableExternalLinks{value: val, isSet: true}
}

func (v NullableExternalLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
