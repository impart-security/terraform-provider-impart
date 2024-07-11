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

// checks if the CompilationDiagnosticRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompilationDiagnosticRange{}

// CompilationDiagnosticRange struct for CompilationDiagnosticRange
type CompilationDiagnosticRange struct {
	Start  int32  `json:"start"`
	End    int32  `json:"end"`
	Source string `json:"source"`
}

type _CompilationDiagnosticRange CompilationDiagnosticRange

// NewCompilationDiagnosticRange instantiates a new CompilationDiagnosticRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompilationDiagnosticRange(start int32, end int32, source string) *CompilationDiagnosticRange {
	this := CompilationDiagnosticRange{}
	this.Start = start
	this.End = end
	this.Source = source
	return &this
}

// NewCompilationDiagnosticRangeWithDefaults instantiates a new CompilationDiagnosticRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompilationDiagnosticRangeWithDefaults() *CompilationDiagnosticRange {
	this := CompilationDiagnosticRange{}
	return &this
}

// GetStart returns the Start field value
func (o *CompilationDiagnosticRange) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CompilationDiagnosticRange) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *CompilationDiagnosticRange) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *CompilationDiagnosticRange) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *CompilationDiagnosticRange) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *CompilationDiagnosticRange) SetEnd(v int32) {
	o.End = v
}

// GetSource returns the Source field value
func (o *CompilationDiagnosticRange) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CompilationDiagnosticRange) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CompilationDiagnosticRange) SetSource(v string) {
	o.Source = v
}

func (o CompilationDiagnosticRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompilationDiagnosticRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

func (o *CompilationDiagnosticRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"end",
		"source",
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

	varCompilationDiagnosticRange := _CompilationDiagnosticRange{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	err = decoder.Decode(&varCompilationDiagnosticRange)

	if err != nil {
		return err
	}

	*o = CompilationDiagnosticRange(varCompilationDiagnosticRange)

	return err
}

type NullableCompilationDiagnosticRange struct {
	value *CompilationDiagnosticRange
	isSet bool
}

func (v NullableCompilationDiagnosticRange) Get() *CompilationDiagnosticRange {
	return v.value
}

func (v *NullableCompilationDiagnosticRange) Set(val *CompilationDiagnosticRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCompilationDiagnosticRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCompilationDiagnosticRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompilationDiagnosticRange(val *CompilationDiagnosticRange) *NullableCompilationDiagnosticRange {
	return &NullableCompilationDiagnosticRange{value: val, isSet: true}
}

func (v NullableCompilationDiagnosticRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompilationDiagnosticRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
