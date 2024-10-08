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

// checks if the InspectorReqMsg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InspectorReqMsg{}

// InspectorReqMsg An payload sent to the inspector to inspect an HTTP request.
type InspectorReqMsg struct {
	// The request URL. The URL must be valid with either `http` or `https` as the schema. The host may be in the format `hostname:port` or `hostname`. If the port is not present, the default port for the schema is used. If no path is provided `/` is used. The URL may include query parameters.
	Url string `json:"url"`
	// The HTTP method requested by the client.
	Method string `json:"method"`
	// Indicates whether the request body was truncated.
	TruncatedBody *bool `json:"truncated_body,omitempty"`
	// The base64 encoded HTTP request body.
	Body *string `json:"body,omitempty"`
	// The HTTP request header keys. Each key should have a corresponding header_values at the matching index.
	HeaderKeys []string `json:"header_keys,omitempty"`
	// The HTTP request header values. Each value should have a corresponding header_keys at the matching index.
	HeaderValues []string `json:"header_values,omitempty"`
	// The HTTP request cookie keys. Each key should have a corresponding cookie_values at the matching index.
	CookieKeys []string `json:"cookie_keys,omitempty"`
	// The HTTP request cookie values. Each value should have a corresponding cookie_keys at the matching index.
	CookieValues []string `json:"cookie_values,omitempty"`
	// The address of the client. This will usually be a load balancer address.
	RemoteAddr           *string `json:"remote_addr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InspectorReqMsg InspectorReqMsg

// NewInspectorReqMsg instantiates a new InspectorReqMsg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInspectorReqMsg(url string, method string) *InspectorReqMsg {
	this := InspectorReqMsg{}
	this.Url = url
	this.Method = method
	return &this
}

// NewInspectorReqMsgWithDefaults instantiates a new InspectorReqMsg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInspectorReqMsgWithDefaults() *InspectorReqMsg {
	this := InspectorReqMsg{}
	return &this
}

// GetUrl returns the Url field value
func (o *InspectorReqMsg) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InspectorReqMsg) SetUrl(v string) {
	o.Url = v
}

// GetMethod returns the Method field value
func (o *InspectorReqMsg) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *InspectorReqMsg) SetMethod(v string) {
	o.Method = v
}

// GetTruncatedBody returns the TruncatedBody field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetTruncatedBody() bool {
	if o == nil || IsNil(o.TruncatedBody) {
		var ret bool
		return ret
	}
	return *o.TruncatedBody
}

// GetTruncatedBodyOk returns a tuple with the TruncatedBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetTruncatedBodyOk() (*bool, bool) {
	if o == nil || IsNil(o.TruncatedBody) {
		return nil, false
	}
	return o.TruncatedBody, true
}

// HasTruncatedBody returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasTruncatedBody() bool {
	if o != nil && !IsNil(o.TruncatedBody) {
		return true
	}

	return false
}

// SetTruncatedBody gets a reference to the given bool and assigns it to the TruncatedBody field.
func (o *InspectorReqMsg) SetTruncatedBody(v bool) {
	o.TruncatedBody = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InspectorReqMsg) SetBody(v string) {
	o.Body = &v
}

// GetHeaderKeys returns the HeaderKeys field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetHeaderKeys() []string {
	if o == nil || IsNil(o.HeaderKeys) {
		var ret []string
		return ret
	}
	return o.HeaderKeys
}

// GetHeaderKeysOk returns a tuple with the HeaderKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetHeaderKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.HeaderKeys) {
		return nil, false
	}
	return o.HeaderKeys, true
}

// HasHeaderKeys returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasHeaderKeys() bool {
	if o != nil && !IsNil(o.HeaderKeys) {
		return true
	}

	return false
}

// SetHeaderKeys gets a reference to the given []string and assigns it to the HeaderKeys field.
func (o *InspectorReqMsg) SetHeaderKeys(v []string) {
	o.HeaderKeys = v
}

// GetHeaderValues returns the HeaderValues field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetHeaderValues() []string {
	if o == nil || IsNil(o.HeaderValues) {
		var ret []string
		return ret
	}
	return o.HeaderValues
}

// GetHeaderValuesOk returns a tuple with the HeaderValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetHeaderValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.HeaderValues) {
		return nil, false
	}
	return o.HeaderValues, true
}

// HasHeaderValues returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasHeaderValues() bool {
	if o != nil && !IsNil(o.HeaderValues) {
		return true
	}

	return false
}

// SetHeaderValues gets a reference to the given []string and assigns it to the HeaderValues field.
func (o *InspectorReqMsg) SetHeaderValues(v []string) {
	o.HeaderValues = v
}

// GetCookieKeys returns the CookieKeys field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetCookieKeys() []string {
	if o == nil || IsNil(o.CookieKeys) {
		var ret []string
		return ret
	}
	return o.CookieKeys
}

// GetCookieKeysOk returns a tuple with the CookieKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetCookieKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.CookieKeys) {
		return nil, false
	}
	return o.CookieKeys, true
}

// HasCookieKeys returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasCookieKeys() bool {
	if o != nil && !IsNil(o.CookieKeys) {
		return true
	}

	return false
}

// SetCookieKeys gets a reference to the given []string and assigns it to the CookieKeys field.
func (o *InspectorReqMsg) SetCookieKeys(v []string) {
	o.CookieKeys = v
}

// GetCookieValues returns the CookieValues field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetCookieValues() []string {
	if o == nil || IsNil(o.CookieValues) {
		var ret []string
		return ret
	}
	return o.CookieValues
}

// GetCookieValuesOk returns a tuple with the CookieValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetCookieValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.CookieValues) {
		return nil, false
	}
	return o.CookieValues, true
}

// HasCookieValues returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasCookieValues() bool {
	if o != nil && !IsNil(o.CookieValues) {
		return true
	}

	return false
}

// SetCookieValues gets a reference to the given []string and assigns it to the CookieValues field.
func (o *InspectorReqMsg) SetCookieValues(v []string) {
	o.CookieValues = v
}

// GetRemoteAddr returns the RemoteAddr field value if set, zero value otherwise.
func (o *InspectorReqMsg) GetRemoteAddr() string {
	if o == nil || IsNil(o.RemoteAddr) {
		var ret string
		return ret
	}
	return *o.RemoteAddr
}

// GetRemoteAddrOk returns a tuple with the RemoteAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectorReqMsg) GetRemoteAddrOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddr) {
		return nil, false
	}
	return o.RemoteAddr, true
}

// HasRemoteAddr returns a boolean if a field has been set.
func (o *InspectorReqMsg) HasRemoteAddr() bool {
	if o != nil && !IsNil(o.RemoteAddr) {
		return true
	}

	return false
}

// SetRemoteAddr gets a reference to the given string and assigns it to the RemoteAddr field.
func (o *InspectorReqMsg) SetRemoteAddr(v string) {
	o.RemoteAddr = &v
}

func (o InspectorReqMsg) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InspectorReqMsg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["method"] = o.Method
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
	if !IsNil(o.CookieKeys) {
		toSerialize["cookie_keys"] = o.CookieKeys
	}
	if !IsNil(o.CookieValues) {
		toSerialize["cookie_values"] = o.CookieValues
	}
	if !IsNil(o.RemoteAddr) {
		toSerialize["remote_addr"] = o.RemoteAddr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InspectorReqMsg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"method",
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

	varInspectorReqMsg := _InspectorReqMsg{}

	err = json.Unmarshal(data, &varInspectorReqMsg)

	if err != nil {
		return err
	}

	*o = InspectorReqMsg(varInspectorReqMsg)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "method")
		delete(additionalProperties, "truncated_body")
		delete(additionalProperties, "body")
		delete(additionalProperties, "header_keys")
		delete(additionalProperties, "header_values")
		delete(additionalProperties, "cookie_keys")
		delete(additionalProperties, "cookie_values")
		delete(additionalProperties, "remote_addr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInspectorReqMsg struct {
	value *InspectorReqMsg
	isSet bool
}

func (v NullableInspectorReqMsg) Get() *InspectorReqMsg {
	return v.value
}

func (v *NullableInspectorReqMsg) Set(val *InspectorReqMsg) {
	v.value = val
	v.isSet = true
}

func (v NullableInspectorReqMsg) IsSet() bool {
	return v.isSet
}

func (v *NullableInspectorReqMsg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInspectorReqMsg(val *InspectorReqMsg) *NullableInspectorReqMsg {
	return &NullableInspectorReqMsg{value: val, isSet: true}
}

func (v NullableInspectorReqMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInspectorReqMsg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
