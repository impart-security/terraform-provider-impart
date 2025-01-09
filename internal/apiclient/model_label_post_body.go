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

// checks if the LabelPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelPostBody{}

// LabelPostBody struct for LabelPostBody
type LabelPostBody struct {
	// Slug of the label.
	Slug string `json:"slug"`
	// The display name of the label.
	DisplayName *string `json:"display_name,omitempty"`
	// The description of the label.
	Description          *string     `json:"description,omitempty"`
	Color                *LabelColor `json:"color,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LabelPostBody LabelPostBody

// NewLabelPostBody instantiates a new LabelPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelPostBody(slug string) *LabelPostBody {
	this := LabelPostBody{}
	this.Slug = slug
	var color LabelColor = GRAY
	this.Color = &color
	return &this
}

// NewLabelPostBodyWithDefaults instantiates a new LabelPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelPostBodyWithDefaults() *LabelPostBody {
	this := LabelPostBody{}
	var color LabelColor = GRAY
	this.Color = &color
	return &this
}

// GetSlug returns the Slug field value
func (o *LabelPostBody) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *LabelPostBody) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *LabelPostBody) SetSlug(v string) {
	o.Slug = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LabelPostBody) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelPostBody) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LabelPostBody) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LabelPostBody) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LabelPostBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelPostBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LabelPostBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LabelPostBody) SetDescription(v string) {
	o.Description = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LabelPostBody) GetColor() LabelColor {
	if o == nil || IsNil(o.Color) {
		var ret LabelColor
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelPostBody) GetColorOk() (*LabelColor, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LabelPostBody) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given LabelColor and assigns it to the Color field.
func (o *LabelPostBody) SetColor(v LabelColor) {
	o.Color = &v
}

func (o LabelPostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slug"] = o.Slug
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

func (o *LabelPostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"slug",
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

	varLabelPostBody := _LabelPostBody{}

	err = json.Unmarshal(data, &varLabelPostBody)

	if err != nil {
		return err
	}

	*o = LabelPostBody(varLabelPostBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "slug")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelPostBody struct {
	value *LabelPostBody
	isSet bool
}

func (v NullableLabelPostBody) Get() *LabelPostBody {
	return v.value
}

func (v *NullableLabelPostBody) Set(val *LabelPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelPostBody(val *LabelPostBody) *NullableLabelPostBody {
	return &NullableLabelPostBody{value: val, isSet: true}
}

func (v NullableLabelPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}