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

// checks if the Connector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connector{}

// Connector struct for Connector
type Connector struct {
	// The ID for a connector.
	Id string `json:"id"`
	// The name of the connector.
	Name string `json:"name"`
	// The ID of the connector type.
	ConnectorTypeId string `json:"connector_type_id"`
	// A boolean flag to indicate whether the connector has been authenticated.
	IsConnected bool `json:"is_connected"`
	// A list of connector destinations.
	Destinations []ConnectorDestination `json:"destinations"`
}

type _Connector Connector

// NewConnector instantiates a new Connector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnector(id string, name string, connectorTypeId string, isConnected bool, destinations []ConnectorDestination) *Connector {
	this := Connector{}
	this.Id = id
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	this.IsConnected = isConnected
	this.Destinations = destinations
	return &this
}

// NewConnectorWithDefaults instantiates a new Connector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorWithDefaults() *Connector {
	this := Connector{}
	return &this
}

// GetId returns the Id field value
func (o *Connector) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Connector) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Connector) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Connector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Connector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Connector) SetName(v string) {
	o.Name = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *Connector) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *Connector) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *Connector) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetIsConnected returns the IsConnected field value
func (o *Connector) GetIsConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value
// and a boolean to check if the value has been set.
func (o *Connector) GetIsConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsConnected, true
}

// SetIsConnected sets field value
func (o *Connector) SetIsConnected(v bool) {
	o.IsConnected = v
}

// GetDestinations returns the Destinations field value
func (o *Connector) GetDestinations() []ConnectorDestination {
	if o == nil {
		var ret []ConnectorDestination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *Connector) GetDestinationsOk() ([]ConnectorDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *Connector) SetDestinations(v []ConnectorDestination) {
	o.Destinations = v
}

func (o Connector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["connector_type_id"] = o.ConnectorTypeId
	toSerialize["is_connected"] = o.IsConnected
	toSerialize["destinations"] = o.Destinations
	return toSerialize, nil
}

func (o *Connector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"connector_type_id",
		"is_connected",
		"destinations",
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

	varConnector := _Connector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnector)

	if err != nil {
		return err
	}

	*o = Connector(varConnector)

	return err
}

type NullableConnector struct {
	value *Connector
	isSet bool
}

func (v NullableConnector) Get() *Connector {
	return v.value
}

func (v *NullableConnector) Set(val *Connector) {
	v.value = val
	v.isSet = true
}

func (v NullableConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnector(val *Connector) *NullableConnector {
	return &NullableConnector{value: val, isSet: true}
}

func (v NullableConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}