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

	"gopkg.in/validator.v2"
)

// RulesTestCaseAssertion - struct for RulesTestCaseAssertion
type RulesTestCaseAssertion struct {
	RulesTestCaseAssertionBlock      *RulesTestCaseAssertionBlock
	RulesTestCaseAssertionOutput     *RulesTestCaseAssertionOutput
	RulesTestCaseAssertionStatusCode *RulesTestCaseAssertionStatusCode
	RulesTestCaseAssertionTags       *RulesTestCaseAssertionTags
}

// RulesTestCaseAssertionBlockAsRulesTestCaseAssertion is a convenience function that returns RulesTestCaseAssertionBlock wrapped in RulesTestCaseAssertion
func RulesTestCaseAssertionBlockAsRulesTestCaseAssertion(v *RulesTestCaseAssertionBlock) RulesTestCaseAssertion {
	return RulesTestCaseAssertion{
		RulesTestCaseAssertionBlock: v,
	}
}

// RulesTestCaseAssertionOutputAsRulesTestCaseAssertion is a convenience function that returns RulesTestCaseAssertionOutput wrapped in RulesTestCaseAssertion
func RulesTestCaseAssertionOutputAsRulesTestCaseAssertion(v *RulesTestCaseAssertionOutput) RulesTestCaseAssertion {
	return RulesTestCaseAssertion{
		RulesTestCaseAssertionOutput: v,
	}
}

// RulesTestCaseAssertionStatusCodeAsRulesTestCaseAssertion is a convenience function that returns RulesTestCaseAssertionStatusCode wrapped in RulesTestCaseAssertion
func RulesTestCaseAssertionStatusCodeAsRulesTestCaseAssertion(v *RulesTestCaseAssertionStatusCode) RulesTestCaseAssertion {
	return RulesTestCaseAssertion{
		RulesTestCaseAssertionStatusCode: v,
	}
}

// RulesTestCaseAssertionTagsAsRulesTestCaseAssertion is a convenience function that returns RulesTestCaseAssertionTags wrapped in RulesTestCaseAssertion
func RulesTestCaseAssertionTagsAsRulesTestCaseAssertion(v *RulesTestCaseAssertionTags) RulesTestCaseAssertion {
	return RulesTestCaseAssertion{
		RulesTestCaseAssertionTags: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RulesTestCaseAssertion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RulesTestCaseAssertionBlock
	err = newStrictDecoder(data).Decode(&dst.RulesTestCaseAssertionBlock)
	if err == nil {
		jsonRulesTestCaseAssertionBlock, _ := json.Marshal(dst.RulesTestCaseAssertionBlock)
		if string(jsonRulesTestCaseAssertionBlock) == "{}" { // empty struct
			dst.RulesTestCaseAssertionBlock = nil
		} else {
			if err = validator.Validate(dst.RulesTestCaseAssertionBlock); err != nil {
				dst.RulesTestCaseAssertionBlock = nil
			} else {
				match++
			}
		}
	} else {
		dst.RulesTestCaseAssertionBlock = nil
	}

	// try to unmarshal data into RulesTestCaseAssertionOutput
	err = newStrictDecoder(data).Decode(&dst.RulesTestCaseAssertionOutput)
	if err == nil {
		jsonRulesTestCaseAssertionOutput, _ := json.Marshal(dst.RulesTestCaseAssertionOutput)
		if string(jsonRulesTestCaseAssertionOutput) == "{}" { // empty struct
			dst.RulesTestCaseAssertionOutput = nil
		} else {
			if err = validator.Validate(dst.RulesTestCaseAssertionOutput); err != nil {
				dst.RulesTestCaseAssertionOutput = nil
			} else {
				match++
			}
		}
	} else {
		dst.RulesTestCaseAssertionOutput = nil
	}

	// try to unmarshal data into RulesTestCaseAssertionStatusCode
	err = newStrictDecoder(data).Decode(&dst.RulesTestCaseAssertionStatusCode)
	if err == nil {
		jsonRulesTestCaseAssertionStatusCode, _ := json.Marshal(dst.RulesTestCaseAssertionStatusCode)
		if string(jsonRulesTestCaseAssertionStatusCode) == "{}" { // empty struct
			dst.RulesTestCaseAssertionStatusCode = nil
		} else {
			if err = validator.Validate(dst.RulesTestCaseAssertionStatusCode); err != nil {
				dst.RulesTestCaseAssertionStatusCode = nil
			} else {
				match++
			}
		}
	} else {
		dst.RulesTestCaseAssertionStatusCode = nil
	}

	// try to unmarshal data into RulesTestCaseAssertionTags
	err = newStrictDecoder(data).Decode(&dst.RulesTestCaseAssertionTags)
	if err == nil {
		jsonRulesTestCaseAssertionTags, _ := json.Marshal(dst.RulesTestCaseAssertionTags)
		if string(jsonRulesTestCaseAssertionTags) == "{}" { // empty struct
			dst.RulesTestCaseAssertionTags = nil
		} else {
			if err = validator.Validate(dst.RulesTestCaseAssertionTags); err != nil {
				dst.RulesTestCaseAssertionTags = nil
			} else {
				match++
			}
		}
	} else {
		dst.RulesTestCaseAssertionTags = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RulesTestCaseAssertionBlock = nil
		dst.RulesTestCaseAssertionOutput = nil
		dst.RulesTestCaseAssertionStatusCode = nil
		dst.RulesTestCaseAssertionTags = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RulesTestCaseAssertion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RulesTestCaseAssertion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RulesTestCaseAssertion) MarshalJSON() ([]byte, error) {
	if src.RulesTestCaseAssertionBlock != nil {
		return json.Marshal(&src.RulesTestCaseAssertionBlock)
	}

	if src.RulesTestCaseAssertionOutput != nil {
		return json.Marshal(&src.RulesTestCaseAssertionOutput)
	}

	if src.RulesTestCaseAssertionStatusCode != nil {
		return json.Marshal(&src.RulesTestCaseAssertionStatusCode)
	}

	if src.RulesTestCaseAssertionTags != nil {
		return json.Marshal(&src.RulesTestCaseAssertionTags)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RulesTestCaseAssertion) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RulesTestCaseAssertionBlock != nil {
		return obj.RulesTestCaseAssertionBlock
	}

	if obj.RulesTestCaseAssertionOutput != nil {
		return obj.RulesTestCaseAssertionOutput
	}

	if obj.RulesTestCaseAssertionStatusCode != nil {
		return obj.RulesTestCaseAssertionStatusCode
	}

	if obj.RulesTestCaseAssertionTags != nil {
		return obj.RulesTestCaseAssertionTags
	}

	// all schemas are nil
	return nil
}

type NullableRulesTestCaseAssertion struct {
	value *RulesTestCaseAssertion
	isSet bool
}

func (v NullableRulesTestCaseAssertion) Get() *RulesTestCaseAssertion {
	return v.value
}

func (v *NullableRulesTestCaseAssertion) Set(val *RulesTestCaseAssertion) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseAssertion) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseAssertion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseAssertion(val *RulesTestCaseAssertion) *NullableRulesTestCaseAssertion {
	return &NullableRulesTestCaseAssertion{value: val, isSet: true}
}

func (v NullableRulesTestCaseAssertion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseAssertion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
