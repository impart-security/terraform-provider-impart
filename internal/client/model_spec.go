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
	"time"
)

// checks if the Spec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Spec{}

// Spec struct for Spec
type Spec struct {
	// The ID for a spec.
	Id string `json:"id"`
	// The name for a spec.
	Name string `json:"name"`
	// A specification in base64 encoding. Can be Swagger 2.0 or OAS 3.0.
	Spec string `json:"spec"`
	// A learning spec in base64 encoding.
	LearningSpec string `json:"learning_spec"`
	// The current revision of the spec.
	Revision int32 `json:"revision"`
	// The analysis score for a spec.
	Score NullableFloat32 `json:"score,omitempty"`
	// ID of the  member who created the spec.
	CreatedBy string `json:"created_by"`
	// The date the spec was created.
	CreatedAt time.Time `json:"created_at"`
	// ID of the member who last updated the spec.
	UpdatedBy NullableString `json:"updated_by"`
	// The date of when the spec was last updated.
	UpdatedAt NullableTime `json:"updated_at"`
}

// NewSpec instantiates a new Spec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpec(id string, name string, spec string, learningSpec string, revision int32, createdBy string, createdAt time.Time, updatedBy NullableString, updatedAt NullableTime) *Spec {
	this := Spec{}
	this.Id = id
	this.Name = name
	this.Spec = spec
	this.LearningSpec = learningSpec
	this.Revision = revision
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.UpdatedBy = updatedBy
	this.UpdatedAt = updatedAt
	return &this
}

// NewSpecWithDefaults instantiates a new Spec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecWithDefaults() *Spec {
	this := Spec{}
	return &this
}

// GetId returns the Id field value
func (o *Spec) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Spec) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Spec) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Spec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Spec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Spec) SetName(v string) {
	o.Name = v
}

// GetSpec returns the Spec field value
func (o *Spec) GetSpec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *Spec) GetSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *Spec) SetSpec(v string) {
	o.Spec = v
}

// GetLearningSpec returns the LearningSpec field value
func (o *Spec) GetLearningSpec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LearningSpec
}

// GetLearningSpecOk returns a tuple with the LearningSpec field value
// and a boolean to check if the value has been set.
func (o *Spec) GetLearningSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LearningSpec, true
}

// SetLearningSpec sets field value
func (o *Spec) SetLearningSpec(v string) {
	o.LearningSpec = v
}

// GetRevision returns the Revision field value
func (o *Spec) GetRevision() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *Spec) GetRevisionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *Spec) SetRevision(v int32) {
	o.Revision = v
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Spec) GetScore() float32 {
	if o == nil || IsNil(o.Score.Get()) {
		var ret float32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Spec) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *Spec) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableFloat32 and assigns it to the Score field.
func (o *Spec) SetScore(v float32) {
	o.Score.Set(&v)
}

// SetScoreNil sets the value for Score to be an explicit nil
func (o *Spec) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *Spec) UnsetScore() {
	o.Score.Unset()
}

// GetCreatedBy returns the CreatedBy field value
func (o *Spec) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Spec) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Spec) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Spec) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Spec) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Spec) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Spec) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Spec) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// SetUpdatedBy sets field value
func (o *Spec) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Spec) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Spec) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *Spec) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

func (o Spec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Spec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["spec"] = o.Spec
	toSerialize["learning_spec"] = o.LearningSpec
	toSerialize["revision"] = o.Revision
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_by"] = o.UpdatedBy.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	return toSerialize, nil
}

type NullableSpec struct {
	value *Spec
	isSet bool
}

func (v NullableSpec) Get() *Spec {
	return v.value
}

func (v *NullableSpec) Set(val *Spec) {
	v.value = val
	v.isSet = true
}

func (v NullableSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpec(val *Spec) *NullableSpec {
	return &NullableSpec{value: val, isSet: true}
}

func (v NullableSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
