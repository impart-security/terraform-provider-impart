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

// checks if the RulesTestCaseAssertionTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseAssertionTags{}

// RulesTestCaseAssertionTags struct for RulesTestCaseAssertionTags
type RulesTestCaseAssertionTags struct {
	// The indexes of the messages in the test case the assertion applies to.
	MessageIndexes []int32 `json:"message_indexes"`
	AssertionType  string  `json:"assertion_type" validate:"regexp=^tags$"`
	// The location of the assertion.
	Location  string                                  `json:"location"`
	Condition RulesTestCaseAssertionConditionPresence `json:"condition"`
	// The expected value of the assertion.
	Expected             string `json:"expected"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCaseAssertionTags RulesTestCaseAssertionTags

// NewRulesTestCaseAssertionTags instantiates a new RulesTestCaseAssertionTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseAssertionTags(messageIndexes []int32, assertionType string, location string, condition RulesTestCaseAssertionConditionPresence, expected string) *RulesTestCaseAssertionTags {
	this := RulesTestCaseAssertionTags{}
	this.MessageIndexes = messageIndexes
	this.AssertionType = assertionType
	this.Location = location
	this.Condition = condition
	this.Expected = expected
	return &this
}

// NewRulesTestCaseAssertionTagsWithDefaults instantiates a new RulesTestCaseAssertionTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseAssertionTagsWithDefaults() *RulesTestCaseAssertionTags {
	this := RulesTestCaseAssertionTags{}
	return &this
}

// GetMessageIndexes returns the MessageIndexes field value
func (o *RulesTestCaseAssertionTags) GetMessageIndexes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MessageIndexes
}

// GetMessageIndexesOk returns a tuple with the MessageIndexes field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionTags) GetMessageIndexesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageIndexes, true
}

// SetMessageIndexes sets field value
func (o *RulesTestCaseAssertionTags) SetMessageIndexes(v []int32) {
	o.MessageIndexes = v
}

// GetAssertionType returns the AssertionType field value
func (o *RulesTestCaseAssertionTags) GetAssertionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssertionType
}

// GetAssertionTypeOk returns a tuple with the AssertionType field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionTags) GetAssertionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssertionType, true
}

// SetAssertionType sets field value
func (o *RulesTestCaseAssertionTags) SetAssertionType(v string) {
	o.AssertionType = v
}

// GetLocation returns the Location field value
func (o *RulesTestCaseAssertionTags) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionTags) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *RulesTestCaseAssertionTags) SetLocation(v string) {
	o.Location = v
}

// GetCondition returns the Condition field value
func (o *RulesTestCaseAssertionTags) GetCondition() RulesTestCaseAssertionConditionPresence {
	if o == nil {
		var ret RulesTestCaseAssertionConditionPresence
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionTags) GetConditionOk() (*RulesTestCaseAssertionConditionPresence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *RulesTestCaseAssertionTags) SetCondition(v RulesTestCaseAssertionConditionPresence) {
	o.Condition = v
}

// GetExpected returns the Expected field value
func (o *RulesTestCaseAssertionTags) GetExpected() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionTags) GetExpectedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expected, true
}

// SetExpected sets field value
func (o *RulesTestCaseAssertionTags) SetExpected(v string) {
	o.Expected = v
}

func (o RulesTestCaseAssertionTags) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseAssertionTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_indexes"] = o.MessageIndexes
	toSerialize["assertion_type"] = o.AssertionType
	toSerialize["location"] = o.Location
	toSerialize["condition"] = o.Condition
	toSerialize["expected"] = o.Expected

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseAssertionTags) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_indexes",
		"assertion_type",
		"location",
		"condition",
		"expected",
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

	varRulesTestCaseAssertionTags := _RulesTestCaseAssertionTags{}

	err = json.Unmarshal(data, &varRulesTestCaseAssertionTags)

	if err != nil {
		return err
	}

	*o = RulesTestCaseAssertionTags(varRulesTestCaseAssertionTags)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message_indexes")
		delete(additionalProperties, "assertion_type")
		delete(additionalProperties, "location")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "expected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseAssertionTags struct {
	value *RulesTestCaseAssertionTags
	isSet bool
}

func (v NullableRulesTestCaseAssertionTags) Get() *RulesTestCaseAssertionTags {
	return v.value
}

func (v *NullableRulesTestCaseAssertionTags) Set(val *RulesTestCaseAssertionTags) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseAssertionTags) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseAssertionTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseAssertionTags(val *RulesTestCaseAssertionTags) *NullableRulesTestCaseAssertionTags {
	return &NullableRulesTestCaseAssertionTags{value: val, isSet: true}
}

func (v NullableRulesTestCaseAssertionTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseAssertionTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
