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

// checks if the RulesTestCaseMessagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCaseMessagesInner{}

// RulesTestCaseMessagesInner struct for RulesTestCaseMessagesInner
type RulesTestCaseMessagesInner struct {
	Req InspectorReqMsg `json:"req"`
	Res InspectorResMsg `json:"res"`
	// The number of times to include the message in the test case.
	Count int32 `json:"count"`
	// The delay in milliseconds between message iterations.
	Delay int32 `json:"delay"`
	// The delay in milliseconds after a set of message iterations.
	PostDelay int32 `json:"post_delay"`
	// The description of the test case message.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCaseMessagesInner RulesTestCaseMessagesInner

// NewRulesTestCaseMessagesInner instantiates a new RulesTestCaseMessagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCaseMessagesInner(req InspectorReqMsg, res InspectorResMsg, count int32, delay int32, postDelay int32) *RulesTestCaseMessagesInner {
	this := RulesTestCaseMessagesInner{}
	this.Req = req
	this.Res = res
	this.Count = count
	this.Delay = delay
	this.PostDelay = postDelay
	return &this
}

// NewRulesTestCaseMessagesInnerWithDefaults instantiates a new RulesTestCaseMessagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseMessagesInnerWithDefaults() *RulesTestCaseMessagesInner {
	this := RulesTestCaseMessagesInner{}
	var count int32 = 1
	this.Count = count
	var delay int32 = 0
	this.Delay = delay
	var postDelay int32 = 0
	this.PostDelay = postDelay
	return &this
}

// GetReq returns the Req field value
func (o *RulesTestCaseMessagesInner) GetReq() InspectorReqMsg {
	if o == nil {
		var ret InspectorReqMsg
		return ret
	}

	return o.Req
}

// GetReqOk returns a tuple with the Req field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseMessagesInner) GetReqOk() (*InspectorReqMsg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Req, true
}

// SetReq sets field value
func (o *RulesTestCaseMessagesInner) SetReq(v InspectorReqMsg) {
	o.Req = v
}

// GetRes returns the Res field value
func (o *RulesTestCaseMessagesInner) GetRes() InspectorResMsg {
	if o == nil {
		var ret InspectorResMsg
		return ret
	}

	return o.Res
}

// GetResOk returns a tuple with the Res field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseMessagesInner) GetResOk() (*InspectorResMsg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Res, true
}

// SetRes sets field value
func (o *RulesTestCaseMessagesInner) SetRes(v InspectorResMsg) {
	o.Res = v
}

// GetCount returns the Count field value
func (o *RulesTestCaseMessagesInner) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseMessagesInner) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *RulesTestCaseMessagesInner) SetCount(v int32) {
	o.Count = v
}

// GetDelay returns the Delay field value
func (o *RulesTestCaseMessagesInner) GetDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Delay
}

// GetDelayOk returns a tuple with the Delay field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseMessagesInner) GetDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delay, true
}

// SetDelay sets field value
func (o *RulesTestCaseMessagesInner) SetDelay(v int32) {
	o.Delay = v
}

// GetPostDelay returns the PostDelay field value
func (o *RulesTestCaseMessagesInner) GetPostDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostDelay
}

// GetPostDelayOk returns a tuple with the PostDelay field value
// and a boolean to check if the value has been set.
func (o *RulesTestCaseMessagesInner) GetPostDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostDelay, true
}

// SetPostDelay sets field value
func (o *RulesTestCaseMessagesInner) SetPostDelay(v int32) {
	o.PostDelay = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RulesTestCaseMessagesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesTestCaseMessagesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RulesTestCaseMessagesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RulesTestCaseMessagesInner) SetDescription(v string) {
	o.Description = &v
}

func (o RulesTestCaseMessagesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCaseMessagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["req"] = o.Req
	toSerialize["res"] = o.Res
	toSerialize["count"] = o.Count
	toSerialize["delay"] = o.Delay
	toSerialize["post_delay"] = o.PostDelay
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCaseMessagesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"req",
		"res",
		"count",
		"delay",
		"post_delay",
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

	varRulesTestCaseMessagesInner := _RulesTestCaseMessagesInner{}

	err = json.Unmarshal(data, &varRulesTestCaseMessagesInner)

	if err != nil {
		return err
	}

	*o = RulesTestCaseMessagesInner(varRulesTestCaseMessagesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "req")
		delete(additionalProperties, "res")
		delete(additionalProperties, "count")
		delete(additionalProperties, "delay")
		delete(additionalProperties, "post_delay")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCaseMessagesInner struct {
	value *RulesTestCaseMessagesInner
	isSet bool
}

func (v NullableRulesTestCaseMessagesInner) Get() *RulesTestCaseMessagesInner {
	return v.value
}

func (v *NullableRulesTestCaseMessagesInner) Set(val *RulesTestCaseMessagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCaseMessagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCaseMessagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCaseMessagesInner(val *RulesTestCaseMessagesInner) *NullableRulesTestCaseMessagesInner {
	return &NullableRulesTestCaseMessagesInner{value: val, isSet: true}
}

func (v NullableRulesTestCaseMessagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCaseMessagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
