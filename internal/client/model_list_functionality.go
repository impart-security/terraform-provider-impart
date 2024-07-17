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

// ListFunctionality The functionality of the list.
type ListFunctionality string

// List of list_functionality
const (
	ADD        ListFunctionality = "add"
	ADD_REMOVE ListFunctionality = "add/remove"
	NONE       ListFunctionality = "none"
)

// All allowed values of ListFunctionality enum
var AllowedListFunctionalityEnumValues = []ListFunctionality{
	"add",
	"add/remove",
	"none",
}

func (v *ListFunctionality) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListFunctionality(value)
	for _, existing := range AllowedListFunctionalityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListFunctionality", value)
}

// NewListFunctionalityFromValue returns a pointer to a valid ListFunctionality
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListFunctionalityFromValue(v string) (*ListFunctionality, error) {
	ev := ListFunctionality(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListFunctionality: valid values are %v", v, AllowedListFunctionalityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListFunctionality) IsValid() bool {
	for _, existing := range AllowedListFunctionalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to list_functionality value
func (v ListFunctionality) Ptr() *ListFunctionality {
	return &v
}

type NullableListFunctionality struct {
	value *ListFunctionality
	isSet bool
}

func (v NullableListFunctionality) Get() *ListFunctionality {
	return v.value
}

func (v *NullableListFunctionality) Set(val *ListFunctionality) {
	v.value = val
	v.isSet = true
}

func (v NullableListFunctionality) IsSet() bool {
	return v.isSet
}

func (v *NullableListFunctionality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFunctionality(val *ListFunctionality) *NullableListFunctionality {
	return &NullableListFunctionality{value: val, isSet: true}
}

func (v NullableListFunctionality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFunctionality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
