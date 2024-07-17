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
	"fmt"
	"time"
)

// checks if the ListItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListItemsInner{}

// ListItemsInner struct for ListItemsInner
type ListItemsInner struct {
	// The value of the item.
	Value string `json:"value"`
	// The expiration date of the item.
	Expiration           NullableTime `json:"expiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListItemsInner ListItemsInner

// NewListItemsInner instantiates a new ListItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItemsInner(value string) *ListItemsInner {
	this := ListItemsInner{}
	this.Value = value
	return &this
}

// NewListItemsInnerWithDefaults instantiates a new ListItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemsInnerWithDefaults() *ListItemsInner {
	this := ListItemsInner{}
	return &this
}

// GetValue returns the Value field value
func (o *ListItemsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListItemsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListItemsInner) SetValue(v string) {
	o.Value = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsInner) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsInner) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *ListItemsInner) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *ListItemsInner) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *ListItemsInner) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *ListItemsInner) UnsetExpiration() {
	o.Expiration.Unset()
}

func (o ListItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varListItemsInner := _ListItemsInner{}

	err = json.Unmarshal(data, &varListItemsInner)

	if err != nil {
		return err
	}

	*o = ListItemsInner(varListItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "expiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListItemsInner struct {
	value *ListItemsInner
	isSet bool
}

func (v NullableListItemsInner) Get() *ListItemsInner {
	return v.value
}

func (v *NullableListItemsInner) Set(val *ListItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItemsInner(val *ListItemsInner) *NullableListItemsInner {
	return &NullableListItemsInner{value: val, isSet: true}
}

func (v NullableListItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
