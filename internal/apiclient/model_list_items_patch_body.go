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
)

// checks if the ListItemsPatchBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListItemsPatchBody{}

// ListItemsPatchBody struct for ListItemsPatchBody
type ListItemsPatchBody struct {
	// The items in the list.
	Add []ListItemsInner `json:"add,omitempty"`
	// Items to remove from the list.
	Remove               []string `json:"remove,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListItemsPatchBody ListItemsPatchBody

// NewListItemsPatchBody instantiates a new ListItemsPatchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItemsPatchBody() *ListItemsPatchBody {
	this := ListItemsPatchBody{}
	return &this
}

// NewListItemsPatchBodyWithDefaults instantiates a new ListItemsPatchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemsPatchBodyWithDefaults() *ListItemsPatchBody {
	this := ListItemsPatchBody{}
	return &this
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *ListItemsPatchBody) GetAdd() []ListItemsInner {
	if o == nil || IsNil(o.Add) {
		var ret []ListItemsInner
		return ret
	}
	return o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItemsPatchBody) GetAddOk() ([]ListItemsInner, bool) {
	if o == nil || IsNil(o.Add) {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *ListItemsPatchBody) HasAdd() bool {
	if o != nil && !IsNil(o.Add) {
		return true
	}

	return false
}

// SetAdd gets a reference to the given []ListItemsInner and assigns it to the Add field.
func (o *ListItemsPatchBody) SetAdd(v []ListItemsInner) {
	o.Add = v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *ListItemsPatchBody) GetRemove() []string {
	if o == nil || IsNil(o.Remove) {
		var ret []string
		return ret
	}
	return o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItemsPatchBody) GetRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.Remove) {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *ListItemsPatchBody) HasRemove() bool {
	if o != nil && !IsNil(o.Remove) {
		return true
	}

	return false
}

// SetRemove gets a reference to the given []string and assigns it to the Remove field.
func (o *ListItemsPatchBody) SetRemove(v []string) {
	o.Remove = v
}

func (o ListItemsPatchBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListItemsPatchBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Add) {
		toSerialize["add"] = o.Add
	}
	if !IsNil(o.Remove) {
		toSerialize["remove"] = o.Remove
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListItemsPatchBody) UnmarshalJSON(data []byte) (err error) {
	varListItemsPatchBody := _ListItemsPatchBody{}

	err = json.Unmarshal(data, &varListItemsPatchBody)

	if err != nil {
		return err
	}

	*o = ListItemsPatchBody(varListItemsPatchBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "add")
		delete(additionalProperties, "remove")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListItemsPatchBody struct {
	value *ListItemsPatchBody
	isSet bool
}

func (v NullableListItemsPatchBody) Get() *ListItemsPatchBody {
	return v.value
}

func (v *NullableListItemsPatchBody) Set(val *ListItemsPatchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListItemsPatchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListItemsPatchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItemsPatchBody(val *ListItemsPatchBody) *NullableListItemsPatchBody {
	return &NullableListItemsPatchBody{value: val, isSet: true}
}

func (v NullableListItemsPatchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListItemsPatchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}