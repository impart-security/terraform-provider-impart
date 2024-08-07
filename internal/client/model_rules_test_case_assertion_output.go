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

// checks if the RulesTestCaseAssertionOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseAssertionOutput{}

// RulesTestCaseAssertionOutput struct for RulesTestCaseAssertionOutput
type RulesTestCaseAssertionOutput struct {
	// The indexes of the messages in the test case the assertion applies to.
	MessageIndexes []int32                                 `json:"message_indexes"`
	AssertionType  string                                  `json:"assertion_type" validate:"regexp=^output$"`
	Condition      RulesTestCaseAssertionConditionPresence `json:"condition"`
	// The expected value of the assertion.
	Expected             string `json:"expected"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCaseAssertionOutput RulesTestCaseAssertionOutput

// NewRulesTestCaseAssertionOutput instantiates a new RulesTestCaseAssertionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseAssertionOutput(messageIndexes []int32, assertionType string, condition RulesTestCaseAssertionConditionPresence, expected string) *RulesTestCaseAssertionOutput {
	this := RulesTestCaseAssertionOutput{}
	this.MessageIndexes = messageIndexes
	this.AssertionType = assertionType
	this.Condition = condition
	this.Expected = expected
	return &this
}

// NewRulesTestCaseAssertionOutputWithDefaults instantiates a new RulesTestCaseAssertionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseAssertionOutputWithDefaults() *RulesTestCaseAssertionOutput {
	this := RulesTestCaseAssertionOutput{}
	return &this
}

// GetMessageIndexes returns the MessageIndexes field value
func (o *RulesTestCaseAssertionOutput) GetMessageIndexes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MessageIndexes
}

// GetMessageIndexesOk returns a tuple with the MessageIndexes field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionOutput) GetMessageIndexesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageIndexes, true
}

// SetMessageIndexes sets field value
func (o *RulesTestCaseAssertionOutput) SetMessageIndexes(v []int32) {
	o.MessageIndexes = v
}

// GetAssertionType returns the AssertionType field value
func (o *RulesTestCaseAssertionOutput) GetAssertionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssertionType
}

// GetAssertionTypeOk returns a tuple with the AssertionType field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionOutput) GetAssertionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssertionType, true
}

// SetAssertionType sets field value
func (o *RulesTestCaseAssertionOutput) SetAssertionType(v string) {
	o.AssertionType = v
}

// GetCondition returns the Condition field value
func (o *RulesTestCaseAssertionOutput) GetCondition() RulesTestCaseAssertionConditionPresence {
	if o == nil {
		var ret RulesTestCaseAssertionConditionPresence
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionOutput) GetConditionOk() (*RulesTestCaseAssertionConditionPresence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *RulesTestCaseAssertionOutput) SetCondition(v RulesTestCaseAssertionConditionPresence) {
	o.Condition = v
}

// GetExpected returns the Expected field value
func (o *RulesTestCaseAssertionOutput) GetExpected() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionOutput) GetExpectedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expected, true
}

// SetExpected sets field value
func (o *RulesTestCaseAssertionOutput) SetExpected(v string) {
	o.Expected = v
}

func (o RulesTestCaseAssertionOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseAssertionOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_indexes"] = o.MessageIndexes
	toSerialize["assertion_type"] = o.AssertionType
	toSerialize["condition"] = o.Condition
	toSerialize["expected"] = o.Expected

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseAssertionOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_indexes",
		"assertion_type",
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

	varRulesTestCaseAssertionOutput := _RulesTestCaseAssertionOutput{}

	err = json.Unmarshal(data, &varRulesTestCaseAssertionOutput)

	if err != nil {
		return err
	}

	*o = RulesTestCaseAssertionOutput(varRulesTestCaseAssertionOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message_indexes")
		delete(additionalProperties, "assertion_type")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "expected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseAssertionOutput struct {
	value *RulesTestCaseAssertionOutput
	isSet bool
}

func (v NullableRulesTestCaseAssertionOutput) Get() *RulesTestCaseAssertionOutput {
	return v.value
}

func (v *NullableRulesTestCaseAssertionOutput) Set(val *RulesTestCaseAssertionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseAssertionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseAssertionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseAssertionOutput(val *RulesTestCaseAssertionOutput) *NullableRulesTestCaseAssertionOutput {
	return &NullableRulesTestCaseAssertionOutput{value: val, isSet: true}
}

func (v NullableRulesTestCaseAssertionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseAssertionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
