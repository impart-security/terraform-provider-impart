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

// checks if the LabelPutBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelPutBody{}

// LabelPutBody struct for LabelPutBody
type LabelPutBody struct {
	// The display name of the label.
	DisplayName *string `json:"display_name,omitempty"`
	// The description of the label.
	Description          *string     `json:"description,omitempty"`
	Color                *LabelColor `json:"color,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LabelPutBody LabelPutBody

// NewLabelPutBody instantiates a new LabelPutBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelPutBody() *LabelPutBody {
	this := LabelPutBody{}
	var color LabelColor = GRAY
	this.Color = &color
	return &this
}

// NewLabelPutBodyWithDefaults instantiates a new LabelPutBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelPutBodyWithDefaults() *LabelPutBody {
	this := LabelPutBody{}
	var color LabelColor = GRAY
	this.Color = &color
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LabelPutBody) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelPutBody) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LabelPutBody) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LabelPutBody) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LabelPutBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelPutBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LabelPutBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LabelPutBody) SetDescription(v string) {
	o.Description = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LabelPutBody) GetColor() LabelColor {
	if o == nil || IsNil(o.Color) {
		var ret LabelColor
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelPutBody) GetColorOk() (*LabelColor, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LabelPutBody) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given LabelColor and assigns it to the Color field.
func (o *LabelPutBody) SetColor(v LabelColor) {
	o.Color = &v
}

func (o LabelPutBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelPutBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LabelPutBody) UnmarshalJSON(data []byte) (err error) {
	varLabelPutBody := _LabelPutBody{}

	err = json.Unmarshal(data, &varLabelPutBody)

	if err != nil {
		return err
	}

	*o = LabelPutBody(varLabelPutBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelPutBody struct {
	value *LabelPutBody
	isSet bool
}

func (v NullableLabelPutBody) Get() *LabelPutBody {
	return v.value
}

func (v *NullableLabelPutBody) Set(val *LabelPutBody) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelPutBody) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelPutBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelPutBody(val *LabelPutBody) *NullableLabelPutBody {
	return &NullableLabelPutBody{value: val, isSet: true}
}

func (v NullableLabelPutBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelPutBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}