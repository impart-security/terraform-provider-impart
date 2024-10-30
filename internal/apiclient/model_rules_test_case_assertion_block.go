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

// checks if the RulesTestCaseAssertionBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseAssertionBlock{}

// RulesTestCaseAssertionBlock struct for RulesTestCaseAssertionBlock
type RulesTestCaseAssertionBlock struct {
	// The indexes of the messages in the test case the assertion applies to.
	MessageIndexes []int32 `json:"message_indexes"`
	AssertionType  string  `json:"assertion_type" validate:"regexp=^block$"`
	// The location of the assertion.
	Location string `json:"location"`
	// The expected value of the assertion.
	Expected bool `json:"expected"`
	// The description of the test case assertion.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCaseAssertionBlock RulesTestCaseAssertionBlock

// NewRulesTestCaseAssertionBlock instantiates a new RulesTestCaseAssertionBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseAssertionBlock(messageIndexes []int32, assertionType string, location string, expected bool) *RulesTestCaseAssertionBlock {
	this := RulesTestCaseAssertionBlock{}
	this.MessageIndexes = messageIndexes
	this.AssertionType = assertionType
	this.Location = location
	this.Expected = expected
	return &this
}

// NewRulesTestCaseAssertionBlockWithDefaults instantiates a new RulesTestCaseAssertionBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseAssertionBlockWithDefaults() *RulesTestCaseAssertionBlock {
	this := RulesTestCaseAssertionBlock{}
	return &this
}

// GetMessageIndexes returns the MessageIndexes field value
func (o *RulesTestCaseAssertionBlock) GetMessageIndexes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MessageIndexes
}

// GetMessageIndexesOk returns a tuple with the MessageIndexes field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionBlock) GetMessageIndexesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageIndexes, true
}

// SetMessageIndexes sets field value
func (o *RulesTestCaseAssertionBlock) SetMessageIndexes(v []int32) {
	o.MessageIndexes = v
}

// GetAssertionType returns the AssertionType field value
func (o *RulesTestCaseAssertionBlock) GetAssertionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssertionType
}

// GetAssertionTypeOk returns a tuple with the AssertionType field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionBlock) GetAssertionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssertionType, true
}

// SetAssertionType sets field value
func (o *RulesTestCaseAssertionBlock) SetAssertionType(v string) {
	o.AssertionType = v
}

// GetLocation returns the Location field value
func (o *RulesTestCaseAssertionBlock) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionBlock) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *RulesTestCaseAssertionBlock) SetLocation(v string) {
	o.Location = v
}

// GetExpected returns the Expected field value
func (o *RulesTestCaseAssertionBlock) GetExpected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionBlock) GetExpectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expected, true
}

// SetExpected sets field value
func (o *RulesTestCaseAssertionBlock) SetExpected(v bool) {
	o.Expected = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RulesTestCaseAssertionBlock) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseAssertionBlock) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RulesTestCaseAssertionBlock) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RulesTestCaseAssertionBlock) SetDescription(v string) {
	o.Description = &v
}

func (o RulesTestCaseAssertionBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseAssertionBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_indexes"] = o.MessageIndexes
	toSerialize["assertion_type"] = o.AssertionType
	toSerialize["location"] = o.Location
	toSerialize["expected"] = o.Expected
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseAssertionBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_indexes",
		"assertion_type",
		"location",
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

	varRulesTestCaseAssertionBlock := _RulesTestCaseAssertionBlock{}

	err = json.Unmarshal(data, &varRulesTestCaseAssertionBlock)

	if err != nil {
		return err
	}

	*o = RulesTestCaseAssertionBlock(varRulesTestCaseAssertionBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message_indexes")
		delete(additionalProperties, "assertion_type")
		delete(additionalProperties, "location")
		delete(additionalProperties, "expected")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseAssertionBlock struct {
	value *RulesTestCaseAssertionBlock
	isSet bool
}

func (v NullableRulesTestCaseAssertionBlock) Get() *RulesTestCaseAssertionBlock {
	return v.value
}

func (v *NullableRulesTestCaseAssertionBlock) Set(val *RulesTestCaseAssertionBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseAssertionBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseAssertionBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseAssertionBlock(val *RulesTestCaseAssertionBlock) *NullableRulesTestCaseAssertionBlock {
	return &NullableRulesTestCaseAssertionBlock{value: val, isSet: true}
}

func (v NullableRulesTestCaseAssertionBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseAssertionBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
