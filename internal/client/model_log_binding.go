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
	"time"
)

// checks if the LogBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogBinding{}

// LogBinding struct for LogBinding
type LogBinding struct {
	// The ID for a log binding.
	Id string `json:"id"`
	// The name of a log binding.
	Name string `json:"name"`
	// The logstream ID of a log binding.
	LogstreamId string `json:"logstream_id"`
	// The pattern type of a log binding.
	PatternType string `json:"pattern_type"`
	// The grok/json pattern of a log binding.
	Pattern string `json:"pattern"`
	// The spec ID associated with a log binding.
	SpecId string `json:"spec_id"`
	// ID of the member that created the Logq binding.
	CreatedBy string `json:"created_by"`
	// The date the log binding was created.
	CreatedAt time.Time `json:"created_at"`
}

type _LogBinding LogBinding

// NewLogBinding instantiates a new LogBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogBinding(id string, name string, logstreamId string, patternType string, pattern string, specId string, createdBy string, createdAt time.Time) *LogBinding {
	this := LogBinding{}
	this.Id = id
	this.Name = name
	this.LogstreamId = logstreamId
	this.PatternType = patternType
	this.Pattern = pattern
	this.SpecId = specId
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	return &this
}

// NewLogBindingWithDefaults instantiates a new LogBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogBindingWithDefaults() *LogBinding {
	this := LogBinding{}
	return &this
}

// GetId returns the Id field value
func (o *LogBinding) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogBinding) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LogBinding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogBinding) SetName(v string) {
	o.Name = v
}

// GetLogstreamId returns the LogstreamId field value
func (o *LogBinding) GetLogstreamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogstreamId
}

// GetLogstreamIdOk returns a tuple with the LogstreamId field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetLogstreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogstreamId, true
}

// SetLogstreamId sets field value
func (o *LogBinding) SetLogstreamId(v string) {
	o.LogstreamId = v
}

// GetPatternType returns the PatternType field value
func (o *LogBinding) GetPatternType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetPatternTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value
func (o *LogBinding) SetPatternType(v string) {
	o.PatternType = v
}

// GetPattern returns the Pattern field value
func (o *LogBinding) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *LogBinding) SetPattern(v string) {
	o.Pattern = v
}

// GetSpecId returns the SpecId field value
func (o *LogBinding) GetSpecId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecId
}

// GetSpecIdOk returns a tuple with the SpecId field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetSpecIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecId, true
}

// SetSpecId sets field value
func (o *LogBinding) SetSpecId(v string) {
	o.SpecId = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *LogBinding) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *LogBinding) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LogBinding) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LogBinding) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LogBinding) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o LogBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["logstream_id"] = o.LogstreamId
	toSerialize["pattern_type"] = o.PatternType
	toSerialize["pattern"] = o.Pattern
	toSerialize["spec_id"] = o.SpecId
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *LogBinding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"logstream_id",
		"pattern_type",
		"pattern",
		"spec_id",
		"created_by",
		"created_at",
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

	varLogBinding := _LogBinding{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	err = decoder.Decode(&varLogBinding)

	if err != nil {
		return err
	}

	*o = LogBinding(varLogBinding)

	return err
}

type NullableLogBinding struct {
	value *LogBinding
	isSet bool
}

func (v NullableLogBinding) Get() *LogBinding {
	return v.value
}

func (v *NullableLogBinding) Set(val *LogBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableLogBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableLogBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogBinding(val *LogBinding) *NullableLogBinding {
	return &NullableLogBinding{value: val, isSet: true}
}

func (v NullableLogBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
