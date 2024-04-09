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

// checks if the InspectorResMsg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InspectorResMsg{}

// InspectorResMsg An payload sent to the inspector to inspect an HTTP response.
type InspectorResMsg struct {
	// Indicates whether the response body was truncated.
	TruncatedBody *bool `json:"truncated_body,omitempty"`
	// The base64 encoded HTTP response body.
	Body *string `json:"body,omitempty"`
	// The HTTP response header keys. Each key should have a corresponding header_values at the matching index.
	HeaderKeys []string `json:"header_keys,omitempty"`
	// The HTTP response header values. Each value should have a corresponding header_keys at the matching index.
	HeaderValues []string `json:"header_values,omitempty"`
	// The HTTP status code.
	StatusCode int32 `json:"status_code"`
}

type _InspectorResMsg InspectorResMsg

// NewInspectorResMsg instantiates a new InspectorResMsg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInspectorResMsg(statusCode int32) *InspectorResMsg {
	this := InspectorResMsg{}
	this.StatusCode = statusCode
	return &this
}

// NewInspectorResMsgWithDefaults instantiates a new InspectorResMsg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInspectorResMsgWithDefaults() *InspectorResMsg {
	this := InspectorResMsg{}
	return &this
}

// GetTruncatedBody returns the TruncatedBody field value if set, zero value otherwise.
func (o *InspectorResMsg) GetTruncatedBody() bool {
	if o == nil || IsNil(o.TruncatedBody) {
		var ret bool
		return ret
	}
	return *o.TruncatedBody
}

// GetTruncatedBodyOk returns a tuple with the TruncatedBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorResMsg) GetTruncatedBodyOk() (*bool, bool) {
	if o == nil || IsNil(o.TruncatedBody) {
		return nil, false
	}
	return o.TruncatedBody, true
}

// HasTruncatedBody returns a boolean if a field has been set.
func (o *InspectorResMsg) HasTruncatedBody() bool {
	if o != nil && !IsNil(o.TruncatedBody) {
		return true
	}

	return false
}

// SetTruncatedBody gets a reference to the given bool and assigns it to the TruncatedBody field.
func (o *InspectorResMsg) SetTruncatedBody(v bool) {
	o.TruncatedBody = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InspectorResMsg) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorResMsg) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InspectorResMsg) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InspectorResMsg) SetBody(v string) {
	o.Body = &v
}

// GetHeaderKeys returns the HeaderKeys field value if set, zero value otherwise.
func (o *InspectorResMsg) GetHeaderKeys() []string {
	if o == nil || IsNil(o.HeaderKeys) {
		var ret []string
		return ret
	}
	return o.HeaderKeys
}

// GetHeaderKeysOk returns a tuple with the HeaderKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorResMsg) GetHeaderKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.HeaderKeys) {
		return nil, false
	}
	return o.HeaderKeys, true
}

// HasHeaderKeys returns a boolean if a field has been set.
func (o *InspectorResMsg) HasHeaderKeys() bool {
	if o != nil && !IsNil(o.HeaderKeys) {
		return true
	}

	return false
}

// SetHeaderKeys gets a reference to the given []string and assigns it to the HeaderKeys field.
func (o *InspectorResMsg) SetHeaderKeys(v []string) {
	o.HeaderKeys = v
}

// GetHeaderValues returns the HeaderValues field value if set, zero value otherwise.
func (o *InspectorResMsg) GetHeaderValues() []string {
	if o == nil || IsNil(o.HeaderValues) {
		var ret []string
		return ret
	}
	return o.HeaderValues
}

// GetHeaderValuesOk returns a tuple with the HeaderValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorResMsg) GetHeaderValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.HeaderValues) {
		return nil, false
	}
	return o.HeaderValues, true
}

// HasHeaderValues returns a boolean if a field has been set.
func (o *InspectorResMsg) HasHeaderValues() bool {
	if o != nil && !IsNil(o.HeaderValues) {
		return true
	}

	return false
}

// SetHeaderValues gets a reference to the given []string and assigns it to the HeaderValues field.
func (o *InspectorResMsg) SetHeaderValues(v []string) {
	o.HeaderValues = v
}

// GetStatusCode returns the StatusCode field value
func (o *InspectorResMsg) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *InspectorResMsg) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *InspectorResMsg) SetStatusCode(v int32) {
	o.StatusCode = v
}

func (o InspectorResMsg) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InspectorResMsg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TruncatedBody) {
		toSerialize["truncated_body"] = o.TruncatedBody
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.HeaderKeys) {
		toSerialize["header_keys"] = o.HeaderKeys
	}
	if !IsNil(o.HeaderValues) {
		toSerialize["header_values"] = o.HeaderValues
	}
	toSerialize["status_code"] = o.StatusCode
	return toSerialize, nil
}

func (o *InspectorResMsg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status_code",
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

	varInspectorResMsg := _InspectorResMsg{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInspectorResMsg)

	if err != nil {
		return err
	}

	*o = InspectorResMsg(varInspectorResMsg)

	return err
}

type NullableInspectorResMsg struct {
	value *InspectorResMsg
	isSet bool
}

func (v NullableInspectorResMsg) Get() *InspectorResMsg {
	return v.value
}

func (v *NullableInspectorResMsg) Set(val *InspectorResMsg) {
	v.value = val
	v.isSet = true
}

func (v NullableInspectorResMsg) IsSet() bool {
	return v.isSet
}

func (v *NullableInspectorResMsg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInspectorResMsg(val *InspectorResMsg) *NullableInspectorResMsg {
	return &NullableInspectorResMsg{value: val, isSet: true}
}

func (v NullableInspectorResMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInspectorResMsg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
