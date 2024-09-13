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

// checks if the RulesTestCaseRunPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseRunPostBody{}

// RulesTestCaseRunPostBody struct for RulesTestCaseRunPostBody
type RulesTestCaseRunPostBody struct {
	// The messages of the test case.
	Messages []RulesTestCaseMessagesInner `json:"messages"`
	// Assertions for the test case.
	Assertions []RulesTestCaseAssertion `json:"assertions,omitempty"`
	// Define rule overrides for a test run.
	EnabledOverrides []RulesTestCaseRunPostBodyEnabledOverridesInner `json:"enabled_overrides,omitempty"`
	// A list of rule ids that should be disabled for the test run.
	DisabledOverrides []RulesTestCaseRunPostBodyDisabledOverridesInner `json:"disabled_overrides,omitempty"`
	// A list of rule ids that should have their blocking effect overridden for the test run.
	BlockingEffectOverrides []RulesTestCaseRunPostBodyBlockingEffectOverridesInner `json:"blocking_effect_overrides,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _RulesTestCaseRunPostBody RulesTestCaseRunPostBody

// NewRulesTestCaseRunPostBody instantiates a new RulesTestCaseRunPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseRunPostBody(messages []RulesTestCaseMessagesInner) *RulesTestCaseRunPostBody {
	this := RulesTestCaseRunPostBody{}
	this.Messages = messages
	return &this
}

// NewRulesTestCaseRunPostBodyWithDefaults instantiates a new RulesTestCaseRunPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseRunPostBodyWithDefaults() *RulesTestCaseRunPostBody {
	this := RulesTestCaseRunPostBody{}
	return &this
}

// GetMessages returns the Messages field value
func (o *RulesTestCaseRunPostBody) GetMessages() []RulesTestCaseMessagesInner {
	if o == nil {
		var ret []RulesTestCaseMessagesInner
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBody) GetMessagesOk() ([]RulesTestCaseMessagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *RulesTestCaseRunPostBody) SetMessages(v []RulesTestCaseMessagesInner) {
	o.Messages = v
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *RulesTestCaseRunPostBody) GetAssertions() []RulesTestCaseAssertion {
	if o == nil || IsNil(o.Assertions) {
		var ret []RulesTestCaseAssertion
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBody) GetAssertionsOk() ([]RulesTestCaseAssertion, bool) {
	if o == nil || IsNil(o.Assertions) {
		return nil, false
	}
	return o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *RulesTestCaseRunPostBody) HasAssertions() bool {
	if o != nil && !IsNil(o.Assertions) {
		return true
	}

	return false
}

// SetAssertions gets a reference to the given []RulesTestCaseAssertion and assigns it to the Assertions field.
func (o *RulesTestCaseRunPostBody) SetAssertions(v []RulesTestCaseAssertion) {
	o.Assertions = v
}

// GetEnabledOverrides returns the EnabledOverrides field value if set, zero value otherwise.
func (o *RulesTestCaseRunPostBody) GetEnabledOverrides() []RulesTestCaseRunPostBodyEnabledOverridesInner {
	if o == nil || IsNil(o.EnabledOverrides) {
		var ret []RulesTestCaseRunPostBodyEnabledOverridesInner
		return ret
	}
	return o.EnabledOverrides
}

// GetEnabledOverridesOk returns a tuple with the EnabledOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBody) GetEnabledOverridesOk() ([]RulesTestCaseRunPostBodyEnabledOverridesInner, bool) {
	if o == nil || IsNil(o.EnabledOverrides) {
		return nil, false
	}
	return o.EnabledOverrides, true
}

// HasEnabledOverrides returns a boolean if a field has been set.
func (o *RulesTestCaseRunPostBody) HasEnabledOverrides() bool {
	if o != nil && !IsNil(o.EnabledOverrides) {
		return true
	}

	return false
}

// SetEnabledOverrides gets a reference to the given []RulesTestCaseRunPostBodyEnabledOverridesInner and assigns it to the EnabledOverrides field.
func (o *RulesTestCaseRunPostBody) SetEnabledOverrides(v []RulesTestCaseRunPostBodyEnabledOverridesInner) {
	o.EnabledOverrides = v
}

// GetDisabledOverrides returns the DisabledOverrides field value if set, zero value otherwise.
func (o *RulesTestCaseRunPostBody) GetDisabledOverrides() []RulesTestCaseRunPostBodyDisabledOverridesInner {
	if o == nil || IsNil(o.DisabledOverrides) {
		var ret []RulesTestCaseRunPostBodyDisabledOverridesInner
		return ret
	}
	return o.DisabledOverrides
}

// GetDisabledOverridesOk returns a tuple with the DisabledOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBody) GetDisabledOverridesOk() ([]RulesTestCaseRunPostBodyDisabledOverridesInner, bool) {
	if o == nil || IsNil(o.DisabledOverrides) {
		return nil, false
	}
	return o.DisabledOverrides, true
}

// HasDisabledOverrides returns a boolean if a field has been set.
func (o *RulesTestCaseRunPostBody) HasDisabledOverrides() bool {
	if o != nil && !IsNil(o.DisabledOverrides) {
		return true
	}

	return false
}

// SetDisabledOverrides gets a reference to the given []RulesTestCaseRunPostBodyDisabledOverridesInner and assigns it to the DisabledOverrides field.
func (o *RulesTestCaseRunPostBody) SetDisabledOverrides(v []RulesTestCaseRunPostBodyDisabledOverridesInner) {
	o.DisabledOverrides = v
}

// GetBlockingEffectOverrides returns the BlockingEffectOverrides field value if set, zero value otherwise.
func (o *RulesTestCaseRunPostBody) GetBlockingEffectOverrides() []RulesTestCaseRunPostBodyBlockingEffectOverridesInner {
	if o == nil || IsNil(o.BlockingEffectOverrides) {
		var ret []RulesTestCaseRunPostBodyBlockingEffectOverridesInner
		return ret
	}
	return o.BlockingEffectOverrides
}

// GetBlockingEffectOverridesOk returns a tuple with the BlockingEffectOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBody) GetBlockingEffectOverridesOk() ([]RulesTestCaseRunPostBodyBlockingEffectOverridesInner, bool) {
	if o == nil || IsNil(o.BlockingEffectOverrides) {
		return nil, false
	}
	return o.BlockingEffectOverrides, true
}

// HasBlockingEffectOverrides returns a boolean if a field has been set.
func (o *RulesTestCaseRunPostBody) HasBlockingEffectOverrides() bool {
	if o != nil && !IsNil(o.BlockingEffectOverrides) {
		return true
	}

	return false
}

// SetBlockingEffectOverrides gets a reference to the given []RulesTestCaseRunPostBodyBlockingEffectOverridesInner and assigns it to the BlockingEffectOverrides field.
func (o *RulesTestCaseRunPostBody) SetBlockingEffectOverrides(v []RulesTestCaseRunPostBodyBlockingEffectOverridesInner) {
	o.BlockingEffectOverrides = v
}

func (o RulesTestCaseRunPostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseRunPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	if !IsNil(o.Assertions) {
		toSerialize["assertions"] = o.Assertions
	}
	if !IsNil(o.EnabledOverrides) {
		toSerialize["enabled_overrides"] = o.EnabledOverrides
	}
	if !IsNil(o.DisabledOverrides) {
		toSerialize["disabled_overrides"] = o.DisabledOverrides
	}
	if !IsNil(o.BlockingEffectOverrides) {
		toSerialize["blocking_effect_overrides"] = o.BlockingEffectOverrides
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseRunPostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
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

	varRulesTestCaseRunPostBody := _RulesTestCaseRunPostBody{}

	err = json.Unmarshal(data, &varRulesTestCaseRunPostBody)

	if err != nil {
		return err
	}

	*o = RulesTestCaseRunPostBody(varRulesTestCaseRunPostBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "messages")
		delete(additionalProperties, "assertions")
		delete(additionalProperties, "enabled_overrides")
		delete(additionalProperties, "disabled_overrides")
		delete(additionalProperties, "blocking_effect_overrides")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseRunPostBody struct {
	value *RulesTestCaseRunPostBody
	isSet bool
}

func (v NullableRulesTestCaseRunPostBody) Get() *RulesTestCaseRunPostBody {
	return v.value
}

func (v *NullableRulesTestCaseRunPostBody) Set(val *RulesTestCaseRunPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseRunPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseRunPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseRunPostBody(val *RulesTestCaseRunPostBody) *NullableRulesTestCaseRunPostBody {
	return &NullableRulesTestCaseRunPostBody{value: val, isSet: true}
}

func (v NullableRulesTestCaseRunPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseRunPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
