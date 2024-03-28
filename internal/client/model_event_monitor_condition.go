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

// checks if the EventMonitorCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventMonitorCondition{}

// EventMonitorCondition struct for EventMonitorCondition
type EventMonitorCondition struct {
	// The number of times a condition can be met before it triggers its monitor.
	Threshold int32 `json:"threshold"`
	// How the condition should compare values (less than or greater than)
	Comparator string `json:"comparator"`
	// The time offset from now() in the past in milliseconds.
	Delay int64 `json:"delay"`
	// Indicates how far in the future from delay we will be looking in milliseconds.
	TimePeriod int64                        `json:"time_period"`
	Details    EventMonitorConditionDetails `json:"details"`
}

type _EventMonitorCondition EventMonitorCondition

// NewEventMonitorCondition instantiates a new EventMonitorCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMonitorCondition(threshold int32, comparator string, delay int64, timePeriod int64, details EventMonitorConditionDetails) *EventMonitorCondition {
	this := EventMonitorCondition{}
	this.Threshold = threshold
	this.Comparator = comparator
	this.Delay = delay
	this.TimePeriod = timePeriod
	this.Details = details
	return &this
}

// NewEventMonitorConditionWithDefaults instantiates a new EventMonitorCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMonitorConditionWithDefaults() *EventMonitorCondition {
	this := EventMonitorCondition{}
	return &this
}

// GetThreshold returns the Threshold field value
func (o *EventMonitorCondition) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *EventMonitorCondition) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *EventMonitorCondition) SetThreshold(v int32) {
	o.Threshold = v
}

// GetComparator returns the Comparator field value
func (o *EventMonitorCondition) GetComparator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *EventMonitorCondition) GetComparatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value
func (o *EventMonitorCondition) SetComparator(v string) {
	o.Comparator = v
}

// GetDelay returns the Delay field value
func (o *EventMonitorCondition) GetDelay() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Delay
}

// GetDelayOk returns a tuple with the Delay field value
// and a boolean to check if the value has been set.
func (o *EventMonitorCondition) GetDelayOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delay, true
}

// SetDelay sets field value
func (o *EventMonitorCondition) SetDelay(v int64) {
	o.Delay = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *EventMonitorCondition) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *EventMonitorCondition) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *EventMonitorCondition) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetDetails returns the Details field value
func (o *EventMonitorCondition) GetDetails() EventMonitorConditionDetails {
	if o == nil {
		var ret EventMonitorConditionDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *EventMonitorCondition) GetDetailsOk() (*EventMonitorConditionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *EventMonitorCondition) SetDetails(v EventMonitorConditionDetails) {
	o.Details = v
}

func (o EventMonitorCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventMonitorCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threshold"] = o.Threshold
	toSerialize["comparator"] = o.Comparator
	toSerialize["delay"] = o.Delay
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

func (o *EventMonitorCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threshold",
		"comparator",
		"delay",
		"time_period",
		"details",
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

	varEventMonitorCondition := _EventMonitorCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventMonitorCondition)

	if err != nil {
		return err
	}

	*o = EventMonitorCondition(varEventMonitorCondition)

	return err
}

type NullableEventMonitorCondition struct {
	value *EventMonitorCondition
	isSet bool
}

func (v NullableEventMonitorCondition) Get() *EventMonitorCondition {
	return v.value
}

func (v *NullableEventMonitorCondition) Set(val *EventMonitorCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMonitorCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMonitorCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMonitorCondition(val *EventMonitorCondition) *NullableEventMonitorCondition {
	return &NullableEventMonitorCondition{value: val, isSet: true}
}

func (v NullableEventMonitorCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMonitorCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
