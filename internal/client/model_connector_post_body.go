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

// checks if the ConnectorPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorPostBody{}

// ConnectorPostBody struct for ConnectorPostBody
type ConnectorPostBody struct {
	// The name of the connector and what it represents.
	Name string `json:"name"`
	// The ID of the connector type.
	ConnectorTypeId string `json:"connector_type_id"`
	// A JSON blob of the fields that will be used to configure the connector
	Config string `json:"config"`
}

type _ConnectorPostBody ConnectorPostBody

// NewConnectorPostBody instantiates a new ConnectorPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorPostBody(name string, connectorTypeId string, config string) *ConnectorPostBody {
	this := ConnectorPostBody{}
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	this.Config = config
	return &this
}

// NewConnectorPostBodyWithDefaults instantiates a new ConnectorPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorPostBodyWithDefaults() *ConnectorPostBody {
	this := ConnectorPostBody{}
	return &this
}

// GetName returns the Name field value
func (o *ConnectorPostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorPostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorPostBody) SetName(v string) {
	o.Name = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *ConnectorPostBody) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *ConnectorPostBody) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *ConnectorPostBody) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetConfig returns the Config field value
func (o *ConnectorPostBody) GetConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ConnectorPostBody) GetConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ConnectorPostBody) SetConfig(v string) {
	o.Config = v
}

func (o ConnectorPostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["connector_type_id"] = o.ConnectorTypeId
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

func (o *ConnectorPostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"connector_type_id",
		"config",
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

	varConnectorPostBody := _ConnectorPostBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	err = decoder.Decode(&varConnectorPostBody)

	if err != nil {
		return err
	}

	*o = ConnectorPostBody(varConnectorPostBody)

	return err
}

type NullableConnectorPostBody struct {
	value *ConnectorPostBody
	isSet bool
}

func (v NullableConnectorPostBody) Get() *ConnectorPostBody {
	return v.value
}

func (v *NullableConnectorPostBody) Set(val *ConnectorPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorPostBody(val *ConnectorPostBody) *NullableConnectorPostBody {
	return &NullableConnectorPostBody{value: val, isSet: true}
}

func (v NullableConnectorPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
