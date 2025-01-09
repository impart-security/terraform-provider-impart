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

// checks if the TagPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagPostBody{}

// TagPostBody struct for TagPostBody
type TagPostBody struct {
	// The tag name.
	Name string `json:"name"`
	// A description of the tag.
	Description *string `json:"description,omitempty"`
	// A risk statement for the tag.
	RiskStatement *string          `json:"risk_statement,omitempty"`
	Remediations  []TagRemediation `json:"remediations,omitempty"`
	// An external URL for the tag.
	ExternalUrl *string `json:"external_url,omitempty"`
	// The applied labels.
	Labels               []string `json:"labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TagPostBody TagPostBody

// NewTagPostBody instantiates a new TagPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagPostBody(name string) *TagPostBody {
	this := TagPostBody{}
	this.Name = name
	return &this
}

// NewTagPostBodyWithDefaults instantiates a new TagPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagPostBodyWithDefaults() *TagPostBody {
	this := TagPostBody{}
	return &this
}

// GetName returns the Name field value
func (o *TagPostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagPostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagPostBody) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagPostBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPostBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagPostBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagPostBody) SetDescription(v string) {
	o.Description = &v
}

// GetRiskStatement returns the RiskStatement field value if set, zero value otherwise.
func (o *TagPostBody) GetRiskStatement() string {
	if o == nil || IsNil(o.RiskStatement) {
		var ret string
		return ret
	}
	return *o.RiskStatement
}

// GetRiskStatementOk returns a tuple with the RiskStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPostBody) GetRiskStatementOk() (*string, bool) {
	if o == nil || IsNil(o.RiskStatement) {
		return nil, false
	}
	return o.RiskStatement, true
}

// HasRiskStatement returns a boolean if a field has been set.
func (o *TagPostBody) HasRiskStatement() bool {
	if o != nil && !IsNil(o.RiskStatement) {
		return true
	}

	return false
}

// SetRiskStatement gets a reference to the given string and assigns it to the RiskStatement field.
func (o *TagPostBody) SetRiskStatement(v string) {
	o.RiskStatement = &v
}

// GetRemediations returns the Remediations field value if set, zero value otherwise.
func (o *TagPostBody) GetRemediations() []TagRemediation {
	if o == nil || IsNil(o.Remediations) {
		var ret []TagRemediation
		return ret
	}
	return o.Remediations
}

// GetRemediationsOk returns a tuple with the Remediations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPostBody) GetRemediationsOk() ([]TagRemediation, bool) {
	if o == nil || IsNil(o.Remediations) {
		return nil, false
	}
	return o.Remediations, true
}

// HasRemediations returns a boolean if a field has been set.
func (o *TagPostBody) HasRemediations() bool {
	if o != nil && !IsNil(o.Remediations) {
		return true
	}

	return false
}

// SetRemediations gets a reference to the given []TagRemediation and assigns it to the Remediations field.
func (o *TagPostBody) SetRemediations(v []TagRemediation) {
	o.Remediations = v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *TagPostBody) GetExternalUrl() string {
	if o == nil || IsNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPostBody) GetExternalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUrl) {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *TagPostBody) HasExternalUrl() bool {
	if o != nil && !IsNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *TagPostBody) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TagPostBody) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPostBody) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *TagPostBody) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *TagPostBody) SetLabels(v []string) {
	o.Labels = v
}

func (o TagPostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RiskStatement) {
		toSerialize["risk_statement"] = o.RiskStatement
	}
	if !IsNil(o.Remediations) {
		toSerialize["remediations"] = o.Remediations
	}
	if !IsNil(o.ExternalUrl) {
		toSerialize["external_url"] = o.ExternalUrl
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TagPostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTagPostBody := _TagPostBody{}

	err = json.Unmarshal(data, &varTagPostBody)

	if err != nil {
		return err
	}

	*o = TagPostBody(varTagPostBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "risk_statement")
		delete(additionalProperties, "remediations")
		delete(additionalProperties, "external_url")
		delete(additionalProperties, "labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTagPostBody struct {
	value *TagPostBody
	isSet bool
}

func (v NullableTagPostBody) Get() *TagPostBody {
	return v.value
}

func (v *NullableTagPostBody) Set(val *TagPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTagPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTagPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagPostBody(val *TagPostBody) *NullableTagPostBody {
	return &NullableTagPostBody{value: val, isSet: true}
}

func (v NullableTagPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}