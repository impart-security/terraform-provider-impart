/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RulesDependenciesPutBodyInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesDependenciesPutBodyInner{}

// RulesDependenciesPutBodyInner struct for RulesDependenciesPutBodyInner
type RulesDependenciesPutBodyInner struct {
	// The rule ID.
	RuleId string `json:"rule_id"`
	// The IDs of the rule dependencies.
	Dependencies []string `json:"dependencies"`
}

type _RulesDependenciesPutBodyInner RulesDependenciesPutBodyInner

// NewRulesDependenciesPutBodyInner instantiates a new RulesDependenciesPutBodyInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesDependenciesPutBodyInner(ruleId string, dependencies []string) *RulesDependenciesPutBodyInner {
	this := RulesDependenciesPutBodyInner{}
	this.RuleId = ruleId
	this.Dependencies = dependencies
	return &this
}

// NewRulesDependenciesPutBodyInnerWithDefaults instantiates a new RulesDependenciesPutBodyInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesDependenciesPutBodyInnerWithDefaults() *RulesDependenciesPutBodyInner {
	this := RulesDependenciesPutBodyInner{}
	return &this
}

// GetRuleId returns the RuleId field value
func (o *RulesDependenciesPutBodyInner) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *RulesDependenciesPutBodyInner) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *RulesDependenciesPutBodyInner) SetRuleId(v string) {
	o.RuleId = v
}

// GetDependencies returns the Dependencies field value
func (o *RulesDependenciesPutBodyInner) GetDependencies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
func (o *RulesDependenciesPutBodyInner) GetDependenciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// SetDependencies sets field value
func (o *RulesDependenciesPutBodyInner) SetDependencies(v []string) {
	o.Dependencies = v
}

func (o RulesDependenciesPutBodyInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesDependenciesPutBodyInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rule_id"] = o.RuleId
	toSerialize["dependencies"] = o.Dependencies
	return toSerialize, nil
}

func (o *RulesDependenciesPutBodyInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rule_id",
		"dependencies",
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

	varRulesDependenciesPutBodyInner := _RulesDependenciesPutBodyInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	err = decoder.Decode(&varRulesDependenciesPutBodyInner)

	if err != nil {
		return err
	}

	*o = RulesDependenciesPutBodyInner(varRulesDependenciesPutBodyInner)

	return err
}

type NullableRulesDependenciesPutBodyInner struct {
	value *RulesDependenciesPutBodyInner
	isSet bool
}

func (v NullableRulesDependenciesPutBodyInner) Get() *RulesDependenciesPutBodyInner {
	return v.value
}

func (v *NullableRulesDependenciesPutBodyInner) Set(val *RulesDependenciesPutBodyInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesDependenciesPutBodyInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesDependenciesPutBodyInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesDependenciesPutBodyInner(val *RulesDependenciesPutBodyInner) *NullableRulesDependenciesPutBodyInner {
	return &NullableRulesDependenciesPutBodyInner{value: val, isSet: true}
}

func (v NullableRulesDependenciesPutBodyInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesDependenciesPutBodyInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
