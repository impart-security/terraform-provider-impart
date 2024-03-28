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

// checks if the ConnectorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorType{}

// ConnectorType struct for ConnectorType
type ConnectorType struct {
	// The ID for a connector.
	Id string `json:"id"`
	// The name of the connector.
	Name string `json:"name"`
	// A JSON schema for the connector configuration.
	ConfigSchema string `json:"config_schema"`
}

type _ConnectorType ConnectorType

// NewConnectorType instantiates a new ConnectorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorType(id string, name string, configSchema string) *ConnectorType {
	this := ConnectorType{}
	this.Id = id
	this.Name = name
	this.ConfigSchema = configSchema
	return &this
}

// NewConnectorTypeWithDefaults instantiates a new ConnectorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTypeWithDefaults() *ConnectorType {
	this := ConnectorType{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectorType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectorType) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectorType) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConnectorType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorType) SetName(v string) {
	o.Name = v
}

// GetConfigSchema returns the ConfigSchema field value
func (o *ConnectorType) GetConfigSchema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigSchema
}

// GetConfigSchemaOk returns a tuple with the ConfigSchema field value
// and a boolean to check if the value has been set.
func (o *ConnectorType) GetConfigSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSchema, true
}

// SetConfigSchema sets field value
func (o *ConnectorType) SetConfigSchema(v string) {
	o.ConfigSchema = v
}

func (o ConnectorType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["config_schema"] = o.ConfigSchema
	return toSerialize, nil
}

func (o *ConnectorType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"config_schema",
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

	varConnectorType := _ConnectorType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectorType)

	if err != nil {
		return err
	}

	*o = ConnectorType(varConnectorType)

	return err
}

type NullableConnectorType struct {
	value *ConnectorType
	isSet bool
}

func (v NullableConnectorType) Get() *ConnectorType {
	return v.value
}

func (v *NullableConnectorType) Set(val *ConnectorType) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorType) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorType(val *ConnectorType) *NullableConnectorType {
	return &NullableConnectorType{value: val, isSet: true}
}

func (v NullableConnectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
