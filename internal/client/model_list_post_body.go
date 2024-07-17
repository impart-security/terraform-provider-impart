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
)

// checks if the ListPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPostBody{}

// ListPostBody struct for ListPostBody
type ListPostBody struct {
	// The name of the list.
	Name          string             `json:"name"`
	Kind          ListKind           `json:"kind"`
	Subkind       *ListSubkind       `json:"subkind,omitempty"`
	Functionality *ListFunctionality `json:"functionality,omitempty"`
	// The items in the list.
	Items                []ListItemsInner `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPostBody ListPostBody

// NewListPostBody instantiates a new ListPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPostBody(name string, kind ListKind) *ListPostBody {
	this := ListPostBody{}
	this.Name = name
	this.Kind = kind
	var functionality ListFunctionality = ADD_REMOVE
	this.Functionality = &functionality
	return &this
}

// NewListPostBodyWithDefaults instantiates a new ListPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPostBodyWithDefaults() *ListPostBody {
	this := ListPostBody{}
	var functionality ListFunctionality = ADD_REMOVE
	this.Functionality = &functionality
	return &this
}

// GetName returns the Name field value
func (o *ListPostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListPostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListPostBody) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value
func (o *ListPostBody) GetKind() ListKind {
	if o == nil {
		var ret ListKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ListPostBody) GetKindOk() (*ListKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ListPostBody) SetKind(v ListKind) {
	o.Kind = v
}

// GetSubkind returns the Subkind field value if set, zero value otherwise.
func (o *ListPostBody) GetSubkind() ListSubkind {
	if o == nil || IsNil(o.Subkind) {
		var ret ListSubkind
		return ret
	}
	return *o.Subkind
}

// GetSubkindOk returns a tuple with the Subkind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPostBody) GetSubkindOk() (*ListSubkind, bool) {
	if o == nil || IsNil(o.Subkind) {
		return nil, false
	}
	return o.Subkind, true
}

// HasSubkind returns a boolean if a field has been set.
func (o *ListPostBody) HasSubkind() bool {
	if o != nil && !IsNil(o.Subkind) {
		return true
	}

	return false
}

// SetSubkind gets a reference to the given ListSubkind and assigns it to the Subkind field.
func (o *ListPostBody) SetSubkind(v ListSubkind) {
	o.Subkind = &v
}

// GetFunctionality returns the Functionality field value if set, zero value otherwise.
func (o *ListPostBody) GetFunctionality() ListFunctionality {
	if o == nil || IsNil(o.Functionality) {
		var ret ListFunctionality
		return ret
	}
	return *o.Functionality
}

// GetFunctionalityOk returns a tuple with the Functionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPostBody) GetFunctionalityOk() (*ListFunctionality, bool) {
	if o == nil || IsNil(o.Functionality) {
		return nil, false
	}
	return o.Functionality, true
}

// HasFunctionality returns a boolean if a field has been set.
func (o *ListPostBody) HasFunctionality() bool {
	if o != nil && !IsNil(o.Functionality) {
		return true
	}

	return false
}

// SetFunctionality gets a reference to the given ListFunctionality and assigns it to the Functionality field.
func (o *ListPostBody) SetFunctionality(v ListFunctionality) {
	o.Functionality = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListPostBody) GetItems() []ListItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []ListItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPostBody) GetItemsOk() ([]ListItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListPostBody) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ListItemsInner and assigns it to the Items field.
func (o *ListPostBody) SetItems(v []ListItemsInner) {
	o.Items = v
}

func (o ListPostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Subkind) {
		toSerialize["subkind"] = o.Subkind
	}
	if !IsNil(o.Functionality) {
		toSerialize["functionality"] = o.Functionality
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"kind",
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

	varListPostBody := _ListPostBody{}

	err = json.Unmarshal(data, &varListPostBody)

	if err != nil {
		return err
	}

	*o = ListPostBody(varListPostBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "subkind")
		delete(additionalProperties, "functionality")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPostBody struct {
	value *ListPostBody
	isSet bool
}

func (v NullableListPostBody) Get() *ListPostBody {
	return v.value
}

func (v *NullableListPostBody) Set(val *ListPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPostBody(val *ListPostBody) *NullableListPostBody {
	return &NullableListPostBody{value: val, isSet: true}
}

func (v NullableListPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
