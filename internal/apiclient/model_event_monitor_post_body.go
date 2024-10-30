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

// checks if the EventMonitorPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventMonitorPostBody{}

// EventMonitorPostBody struct for EventMonitorPostBody
type EventMonitorPostBody struct {
	// The name of the event monitor and what it represents.
	Name string `json:"name"`
	// A human readable string describing what this monitor is for.
	Description string `json:"description"`
	// Array of condition objects that (if all are true) will trigger the monitor.
	Conditions []EventMonitorConditionPostBody `json:"conditions"`
	// The ids of the notification templates that the monitor is associated with.
	NotificationTemplateIds []string `json:"notification_template_ids"`
	// The applied labels.
	Labels               []string `json:"labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventMonitorPostBody EventMonitorPostBody

// NewEventMonitorPostBody instantiates a new EventMonitorPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMonitorPostBody(name string, description string, conditions []EventMonitorConditionPostBody, notificationTemplateIds []string) *EventMonitorPostBody {
	this := EventMonitorPostBody{}
	this.Name = name
	this.Description = description
	this.Conditions = conditions
	this.NotificationTemplateIds = notificationTemplateIds
	return &this
}

// NewEventMonitorPostBodyWithDefaults instantiates a new EventMonitorPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMonitorPostBodyWithDefaults() *EventMonitorPostBody {
	this := EventMonitorPostBody{}
	return &this
}

// GetName returns the Name field value
func (o *EventMonitorPostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventMonitorPostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventMonitorPostBody) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *EventMonitorPostBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EventMonitorPostBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EventMonitorPostBody) SetDescription(v string) {
	o.Description = v
}

// GetConditions returns the Conditions field value
func (o *EventMonitorPostBody) GetConditions() []EventMonitorConditionPostBody {
	if o == nil {
		var ret []EventMonitorConditionPostBody
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *EventMonitorPostBody) GetConditionsOk() ([]EventMonitorConditionPostBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *EventMonitorPostBody) SetConditions(v []EventMonitorConditionPostBody) {
	o.Conditions = v
}

// GetNotificationTemplateIds returns the NotificationTemplateIds field value
func (o *EventMonitorPostBody) GetNotificationTemplateIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NotificationTemplateIds
}

// GetNotificationTemplateIdsOk returns a tuple with the NotificationTemplateIds field value
// and a boolean to check if the value has been set.
func (o *EventMonitorPostBody) GetNotificationTemplateIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationTemplateIds, true
}

// SetNotificationTemplateIds sets field value
func (o *EventMonitorPostBody) SetNotificationTemplateIds(v []string) {
	o.NotificationTemplateIds = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *EventMonitorPostBody) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMonitorPostBody) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *EventMonitorPostBody) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *EventMonitorPostBody) SetLabels(v []string) {
	o.Labels = v
}

func (o EventMonitorPostBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventMonitorPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["conditions"] = o.Conditions
	toSerialize["notification_template_ids"] = o.NotificationTemplateIds
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventMonitorPostBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"conditions",
		"notification_template_ids",
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

	varEventMonitorPostBody := _EventMonitorPostBody{}

	err = json.Unmarshal(data, &varEventMonitorPostBody)

	if err != nil {
		return err
	}

	*o = EventMonitorPostBody(varEventMonitorPostBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "notification_template_ids")
		delete(additionalProperties, "labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventMonitorPostBody struct {
	value *EventMonitorPostBody
	isSet bool
}

func (v NullableEventMonitorPostBody) Get() *EventMonitorPostBody {
	return v.value
}

func (v *NullableEventMonitorPostBody) Set(val *EventMonitorPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMonitorPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMonitorPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMonitorPostBody(val *EventMonitorPostBody) *NullableEventMonitorPostBody {
	return &NullableEventMonitorPostBody{value: val, isSet: true}
}

func (v NullableEventMonitorPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMonitorPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
