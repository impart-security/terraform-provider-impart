/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RulesTestCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesTestCase{}

// RulesTestCase struct for RulesTestCase
type RulesTestCase struct {
	// The unique identifier of the test case.
	Id string `json:"id"`
	// The name of the test case.
	Name string `json:"name"`
	// The description of the test case.
	Description string `json:"description"`
	// The messages of the test case.
	Messages []RulesTestCaseMessagesInner `json:"messages"`
	// Assertions for the test case.
	Assertions []RulesTestCaseAssertion `json:"assertions"`
	// The unique identifier of the user who created the test case.
	CreatedBy string `json:"created_by"`
	// The date and time in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) format.
	CreatedAt time.Time `json:"created_at"`
	// The unique identifier of the user who last updated the test case.
	UpdatedBy NullableString `json:"updated_by"`
	// The date and time in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) format.
	UpdatedAt            NullableTime `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _RulesTestCase RulesTestCase

// NewRulesTestCase instantiates a new RulesTestCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesTestCase(id string, name string, description string, messages []RulesTestCaseMessagesInner, assertions []RulesTestCaseAssertion, createdBy string, createdAt time.Time, updatedBy NullableString, updatedAt NullableTime) *RulesTestCase {
	this := RulesTestCase{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Messages = messages
	this.Assertions = assertions
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.UpdatedBy = updatedBy
	this.UpdatedAt = updatedAt
	return &this
}

// NewRulesTestCaseWithDefaults instantiates a new RulesTestCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesTestCaseWithDefaults() *RulesTestCase {
	this := RulesTestCase{}
	return &this
}

// GetId returns the Id field value
func (o *RulesTestCase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RulesTestCase) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RulesTestCase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RulesTestCase) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RulesTestCase) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RulesTestCase) SetDescription(v string) {
	o.Description = v
}

// GetMessages returns the Messages field value
func (o *RulesTestCase) GetMessages() []RulesTestCaseMessagesInner {
	if o == nil {
		var ret []RulesTestCaseMessagesInner
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetMessagesOk() ([]RulesTestCaseMessagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *RulesTestCase) SetMessages(v []RulesTestCaseMessagesInner) {
	o.Messages = v
}

// GetAssertions returns the Assertions field value
func (o *RulesTestCase) GetAssertions() []RulesTestCaseAssertion {
	if o == nil {
		var ret []RulesTestCaseAssertion
		return ret
	}

	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetAssertionsOk() ([]RulesTestCaseAssertion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assertions, true
}

// SetAssertions sets field value
func (o *RulesTestCase) SetAssertions(v []RulesTestCaseAssertion) {
	o.Assertions = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RulesTestCase) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RulesTestCase) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RulesTestCase) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RulesTestCase) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RulesTestCase) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RulesTestCase) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RulesTestCase) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// SetUpdatedBy sets field value
func (o *RulesTestCase) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RulesTestCase) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RulesTestCase) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *RulesTestCase) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

func (o RulesTestCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesTestCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["messages"] = o.Messages
	toSerialize["assertions"] = o.Assertions
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_by"] = o.UpdatedBy.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RulesTestCase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"messages",
		"assertions",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
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

	varRulesTestCase := _RulesTestCase{}

	err = json.Unmarshal(data, &varRulesTestCase)

	if err != nil {
		return err
	}

	*o = RulesTestCase(varRulesTestCase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "assertions")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRulesTestCase struct {
	value *RulesTestCase
	isSet bool
}

func (v NullableRulesTestCase) Get() *RulesTestCase {
	return v.value
}

func (v *NullableRulesTestCase) Set(val *RulesTestCase) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesTestCase) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesTestCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesTestCase(val *RulesTestCase) *NullableRulesTestCase {
	return &NullableRulesTestCase{value: val, isSet: true}
}

func (v NullableRulesTestCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesTestCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
