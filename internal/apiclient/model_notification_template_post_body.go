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

// checks if the NotificationTemplatePostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationTemplatePostBody{}

// NotificationTemplatePostBody struct for NotificationTemplatePostBody
type NotificationTemplatePostBody struct {
	// The ID of the connector.
	ConnectorId string `json:"connector_id"`
	// The name of the notification template.
	Name string `json:"name"`
	// The text payload to be sent in a notification.
	Payload string `json:"payload"`
	// The text for the subject of the message payload.
	Subject *string `json:"subject,omitempty"`
	// An array of ids for where the notification should be sent.
	Destination []string `json:"destination,omitempty"`
	// The payload to sign using a predefined hmac key. (Ex. \"{{timeTriggered}}.{{URL}}.{{REQUEST_BODY}}\")
	HmacPayloadToSign    *string `json:"hmac_payload_to_sign,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationTemplatePostBody NotificationTemplatePostBody

// NewNotificationTemplatePostBody instantiates a new NotificationTemplatePostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTemplatePostBody(connectorId string, name string, payload string) *NotificationTemplatePostBody {
	this := NotificationTemplatePostBody{}
	this.ConnectorId = connectorId
	this.Name = name
	this.Payload = payload
	return &this
}

// NewNotificationTemplatePostBodyWithDefaults instantiates a new NotificationTemplatePostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTemplatePostBodyWithDefaults() *NotificationTemplatePostBody {
	this := NotificationTemplatePostBody{}
	return &this
}

// GetConnectorId returns the ConnectorId field value
func (o *NotificationTemplatePostBody) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *NotificationTemplatePostBody) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *NotificationTemplatePostBody) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetName returns the Name field value
func (o *NotificationTemplatePostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationTemplatePostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationTemplatePostBody) SetName(v string) {
	o.Name = v
}

// GetPayload returns the Payload field value
func (o *NotificationTemplatePostBody) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *NotificationTemplatePostBody) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *NotificationTemplatePostBody) SetPayload(v string) {
	o.Payload = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *NotificationTemplatePostBody) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplatePostBody) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *NotificationTemplatePostBody) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *NotificationTemplatePostBody) SetSubject(v string) {
	o.Subject = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *NotificationTemplatePostBody) GetDestination() []string {
	if o == nil || IsNil(o.Destination) {
		var ret []string
		return ret
	}
	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplatePostBody) GetDestinationOk() ([]string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *NotificationTemplatePostBody) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given []string and assigns it to the Destination field.
func (o *NotificationTemplatePostBody) SetDestination(v []string) {
	o.Destination = v
}

// GetHmacPayloadToSign returns the HmacPayloadToSign field value if set, zero value otherwise.
func (o *NotificationTemplatePostBody) GetHmacPayloadToSign() string {
	if o == nil || IsNil(o.HmacPayloadToSign) {
		var ret string
		return ret
	}
	return *o.HmacPayloadToSign
}

// GetHmacPayloadToSignOk returns a tuple with the HmacPayloadToSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplatePostBody) GetHmacPayloadToSignOk() (*string, bool) {
	if o == nil || IsNil(o.HmacPayloadToSign) {
		return nil, false
	}
	return o.HmacPayloadToSign, true
}

// HasHmacPayloadToSign returns a boolean if a field has been set.
func (o *NotificationTemplatePostBody) HasHmacPayloadToSign() bool {
	if o != nil && !IsNil(o.HmacPayloadToSign) {
		return true
	}

	return false
}

// SetHmacPayloadToSign gets a reference to the given string and assigns it to the HmacPayloadToSign field.
func (o *NotificationTemplatePostBody) SetHmacPayloadToSign(v string) {
	o.HmacPayloadToSign = &v
}

func (o NotificationTemplatePostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationTemplatePostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connector_id"] = o.ConnectorId
	toSerialize["name"] = o.Name
	toSerialize["payload"] = o.Payload
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.HmacPayloadToSign) {
		toSerialize["hmac_payload_to_sign"] = o.HmacPayloadToSign
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationTemplatePostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connector_id",
		"name",
		"payload",
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

	varNotificationTemplatePostBody := _NotificationTemplatePostBody{}

	err = json.Unmarshal(data, &varNotificationTemplatePostBody)

	if err != nil {
		return err
	}

	*o = NotificationTemplatePostBody(varNotificationTemplatePostBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connector_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "destination")
		delete(additionalProperties, "hmac_payload_to_sign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationTemplatePostBody struct {
	value *NotificationTemplatePostBody
	isSet bool
}

func (v NullableNotificationTemplatePostBody) Get() *NotificationTemplatePostBody {
	return v.value
}

func (v *NullableNotificationTemplatePostBody) Set(val *NotificationTemplatePostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTemplatePostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTemplatePostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTemplatePostBody(val *NotificationTemplatePostBody) *NullableNotificationTemplatePostBody {
	return &NullableNotificationTemplatePostBody{value: val, isSet: true}
}

func (v NullableNotificationTemplatePostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTemplatePostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
