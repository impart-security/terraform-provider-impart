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

// RulesTestCaseAssertionConditionDefault The condition to be met for the assertion to pass.
type RulesTestCaseAssertionConditionDefault string

// List of rules_test_case_assertion_condition_default
const (
	EQUAL        RulesTestCaseAssertionConditionDefault = "equal"
	NOT_EQUAL    RulesTestCaseAssertionConditionDefault = "not_equal"
	GREATER_THAN RulesTestCaseAssertionConditionDefault = "greater_than"
	LESS_THAN    RulesTestCaseAssertionConditionDefault = "less_than"
	ONE_OF       RulesTestCaseAssertionConditionDefault = "one_of"
)

// All allowed values of RulesTestCaseAssertionConditionDefault enum
var AllowedRulesTestCaseAssertionConditionDefaultEnumValues = []RulesTestCaseAssertionConditionDefault{
	"equal",
	"not_equal",
	"greater_than",
	"less_than",
	"one_of",
}

func (v *RulesTestCaseAssertionConditionDefault) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RulesTestCaseAssertionConditionDefault(value)
	for _, existing := range AllowedRulesTestCaseAssertionConditionDefaultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RulesTestCaseAssertionConditionDefault", value)
}

// NewRulesTestCaseAssertionConditionDefaultFromValue returns a pointer to a valid RulesTestCaseAssertionConditionDefault
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRulesTestCaseAssertionConditionDefaultFromValue(v string) (*RulesTestCaseAssertionConditionDefault, error) {
	ev := RulesTestCaseAssertionConditionDefault(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RulesTestCaseAssertionConditionDefault: valid values are %v", v, AllowedRulesTestCaseAssertionConditionDefaultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RulesTestCaseAssertionConditionDefault) IsValid() bool {
	for _, existing := range AllowedRulesTestCaseAssertionConditionDefaultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to rules_test_case_assertion_condition_default value
func (v RulesTestCaseAssertionConditionDefault) Ptr() *RulesTestCaseAssertionConditionDefault {
	return &v
}

type NullableRulesTestCaseAssertionConditionDefault struct {
	value *RulesTestCaseAssertionConditionDefault
	isSet bool
}

func (v NullableRulesTestCaseAssertionConditionDefault) Get() *RulesTestCaseAssertionConditionDefault {
	return v.value
}

func (v *NullableRulesTestCaseAssertionConditionDefault) Set(val *RulesTestCaseAssertionConditionDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseAssertionConditionDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseAssertionConditionDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseAssertionConditionDefault(val *RulesTestCaseAssertionConditionDefault) *NullableRulesTestCaseAssertionConditionDefault {
	return &NullableRulesTestCaseAssertionConditionDefault{value: val, isSet: true}
}

func (v NullableRulesTestCaseAssertionConditionDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseAssertionConditionDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
