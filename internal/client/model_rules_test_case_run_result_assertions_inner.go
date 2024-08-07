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

// checks if the RulesTestCaseRunResultAssertionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseRunResultAssertionsInner{}

// RulesTestCaseRunResultAssertionsInner struct for RulesTestCaseRunResultAssertionsInner
type RulesTestCaseRunResultAssertionsInner struct {
	// The results of the assertion.
	Results              []RulesTestCaseRunResultAssertionsInnerResultsInner `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCaseRunResultAssertionsInner RulesTestCaseRunResultAssertionsInner

// NewRulesTestCaseRunResultAssertionsInner instantiates a new RulesTestCaseRunResultAssertionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseRunResultAssertionsInner(results []RulesTestCaseRunResultAssertionsInnerResultsInner) *RulesTestCaseRunResultAssertionsInner {
	this := RulesTestCaseRunResultAssertionsInner{}
	this.Results = results
	return &this
}

// NewRulesTestCaseRunResultAssertionsInnerWithDefaults instantiates a new RulesTestCaseRunResultAssertionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseRunResultAssertionsInnerWithDefaults() *RulesTestCaseRunResultAssertionsInner {
	this := RulesTestCaseRunResultAssertionsInner{}
	return &this
}

// GetResults returns the Results field value
func (o *RulesTestCaseRunResultAssertionsInner) GetResults() []RulesTestCaseRunResultAssertionsInnerResultsInner {
	if o == nil {
		var ret []RulesTestCaseRunResultAssertionsInnerResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunResultAssertionsInner) GetResultsOk() ([]RulesTestCaseRunResultAssertionsInnerResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *RulesTestCaseRunResultAssertionsInner) SetResults(v []RulesTestCaseRunResultAssertionsInnerResultsInner) {
	o.Results = v
}

func (o RulesTestCaseRunResultAssertionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseRunResultAssertionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseRunResultAssertionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varRulesTestCaseRunResultAssertionsInner := _RulesTestCaseRunResultAssertionsInner{}

	err = json.Unmarshal(data, &varRulesTestCaseRunResultAssertionsInner)

	if err != nil {
		return err
	}

	*o = RulesTestCaseRunResultAssertionsInner(varRulesTestCaseRunResultAssertionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseRunResultAssertionsInner struct {
	value *RulesTestCaseRunResultAssertionsInner
	isSet bool
}

func (v NullableRulesTestCaseRunResultAssertionsInner) Get() *RulesTestCaseRunResultAssertionsInner {
	return v.value
}

func (v *NullableRulesTestCaseRunResultAssertionsInner) Set(val *RulesTestCaseRunResultAssertionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseRunResultAssertionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseRunResultAssertionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseRunResultAssertionsInner(val *RulesTestCaseRunResultAssertionsInner) *NullableRulesTestCaseRunResultAssertionsInner {
	return &NullableRulesTestCaseRunResultAssertionsInner{value: val, isSet: true}
}

func (v NullableRulesTestCaseRunResultAssertionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseRunResultAssertionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
