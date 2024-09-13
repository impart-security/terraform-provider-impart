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

// checks if the RulesTestCaseRunPostBodyBlockingEffectOverridesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseRunPostBodyBlockingEffectOverridesInner{}

// RulesTestCaseRunPostBodyBlockingEffectOverridesInner struct for RulesTestCaseRunPostBodyBlockingEffectOverridesInner
type RulesTestCaseRunPostBodyBlockingEffectOverridesInner struct {
	RuleId               *string             `json:"rule_id,omitempty"`
	BlockingEffect       *BlockingEffectType `json:"blocking_effect,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCaseRunPostBodyBlockingEffectOverridesInner RulesTestCaseRunPostBodyBlockingEffectOverridesInner

// NewRulesTestCaseRunPostBodyBlockingEffectOverridesInner instantiates a new RulesTestCaseRunPostBodyBlockingEffectOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseRunPostBodyBlockingEffectOverridesInner() *RulesTestCaseRunPostBodyBlockingEffectOverridesInner {
	this := RulesTestCaseRunPostBodyBlockingEffectOverridesInner{}
	var blockingEffect BlockingEffectType = BLOCK
	this.BlockingEffect = &blockingEffect
	return &this
}

// NewRulesTestCaseRunPostBodyBlockingEffectOverridesInnerWithDefaults instantiates a new RulesTestCaseRunPostBodyBlockingEffectOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseRunPostBodyBlockingEffectOverridesInnerWithDefaults() *RulesTestCaseRunPostBodyBlockingEffectOverridesInner {
	this := RulesTestCaseRunPostBodyBlockingEffectOverridesInner{}
	var blockingEffect BlockingEffectType = BLOCK
	this.BlockingEffect = &blockingEffect
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) SetRuleId(v string) {
	o.RuleId = &v
}

// GetBlockingEffect returns the BlockingEffect field value if set, zero value otherwise.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) GetBlockingEffect() BlockingEffectType {
	if o == nil || IsNil(o.BlockingEffect) {
		var ret BlockingEffectType
		return ret
	}
	return *o.BlockingEffect
}

// GetBlockingEffectOk returns a tuple with the BlockingEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) GetBlockingEffectOk() (*BlockingEffectType, bool) {
	if o == nil || IsNil(o.BlockingEffect) {
		return nil, false
	}
	return o.BlockingEffect, true
}

// HasBlockingEffect returns a boolean if a field has been set.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) HasBlockingEffect() bool {
	if o != nil && !IsNil(o.BlockingEffect) {
		return true
	}

	return false
}

// SetBlockingEffect gets a reference to the given BlockingEffectType and assigns it to the BlockingEffect field.
func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) SetBlockingEffect(v BlockingEffectType) {
	o.BlockingEffect = &v
}

func (o RulesTestCaseRunPostBodyBlockingEffectOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseRunPostBodyBlockingEffectOverridesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.BlockingEffect) {
		toSerialize["blocking_effect"] = o.BlockingEffect
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) UnmarshalJSON(data []byte) (err error) {
	varRulesTestCaseRunPostBodyBlockingEffectOverridesInner := _RulesTestCaseRunPostBodyBlockingEffectOverridesInner{}

	err = json.Unmarshal(data, &varRulesTestCaseRunPostBodyBlockingEffectOverridesInner)

	if err != nil {
		return err
	}

	*o = RulesTestCaseRunPostBodyBlockingEffectOverridesInner(varRulesTestCaseRunPostBodyBlockingEffectOverridesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rule_id")
		delete(additionalProperties, "blocking_effect")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner struct {
	value *RulesTestCaseRunPostBodyBlockingEffectOverridesInner
	isSet bool
}

func (v NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner) Get() *RulesTestCaseRunPostBodyBlockingEffectOverridesInner {
	return v.value
}

func (v *NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner) Set(val *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner(val *RulesTestCaseRunPostBodyBlockingEffectOverridesInner) *NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner {
	return &NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner{value: val, isSet: true}
}

func (v NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseRunPostBodyBlockingEffectOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
