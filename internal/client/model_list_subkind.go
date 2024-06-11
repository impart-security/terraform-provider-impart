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

// ListSubkind The subkind of list.
type ListSubkind string

// List of list_subkind
const (
	SPEC_UUID     ListSubkind = "spec_uuid"
	ENDPOINT_UUID ListSubkind = "endpoint_uuid"
)

// All allowed values of ListSubkind enum
var AllowedListSubkindEnumValues = []ListSubkind{
	"spec_uuid",
	"endpoint_uuid",
}

func (v *ListSubkind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListSubkind(value)
	for _, existing := range AllowedListSubkindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListSubkind", value)
}

// NewListSubkindFromValue returns a pointer to a valid ListSubkind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListSubkindFromValue(v string) (*ListSubkind, error) {
	ev := ListSubkind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListSubkind: valid values are %v", v, AllowedListSubkindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListSubkind) IsValid() bool {
	for _, existing := range AllowedListSubkindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to list_subkind value
func (v ListSubkind) Ptr() *ListSubkind {
	return &v
}

type NullableListSubkind struct {
	value *ListSubkind
	isSet bool
}

func (v NullableListSubkind) Get() *ListSubkind {
	return v.value
}

func (v *NullableListSubkind) Set(val *ListSubkind) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubkind) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubkind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubkind(val *ListSubkind) *NullableListSubkind {
	return &NullableListSubkind{value: val, isSet: true}
}

func (v NullableListSubkind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubkind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
