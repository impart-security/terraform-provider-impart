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
	"time"
)

// checks if the ExternalLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalLink{}

// ExternalLink struct for ExternalLink
type ExternalLink struct {
	// The ID for the external link.
	Id string `json:"id"`
	// The name of the external link.
	Name string `json:"name"`
	// The description of the external link.
	Description string `json:"description"`
	// A list of spec IDs this external link applies to (empty means all).
	SpecIds []string `json:"spec_ids"`
	// The entity to which the links should be applied.
	Entity string `json:"entity"`
	// If the external link was created by an integration or custom.
	Source string `json:"source"`
	// A JSONPath to the element for which this link should apply (e.g. `$.client_ip.address`).
	JsonPathElement string `json:"json_path_element"`
	// The external URL template with JSONPath element variables.
	Url string `json:"url"`
	// The vendor for the external link.
	Vendor string `json:"vendor"`
	// ID of the member who created the external link.
	CreatedBy string `json:"created_by"`
	// The date the external link was created.
	CreatedAt time.Time `json:"created_at"`
	// ID of the member that last updated the external link.
	UpdatedBy string `json:"updated_by"`
	// The date of when the external link was last updated.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _ExternalLink ExternalLink

// NewExternalLink instantiates a new ExternalLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalLink(id string, name string, description string, specIds []string, entity string, source string, jsonPathElement string, url string, vendor string, createdBy string, createdAt time.Time, updatedBy string, updatedAt time.Time) *ExternalLink {
	this := ExternalLink{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.SpecIds = specIds
	this.Entity = entity
	this.Source = source
	this.JsonPathElement = jsonPathElement
	this.Url = url
	this.Vendor = vendor
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.UpdatedBy = updatedBy
	this.UpdatedAt = updatedAt
	return &this
}

// NewExternalLinkWithDefaults instantiates a new ExternalLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalLinkWithDefaults() *ExternalLink {
	this := ExternalLink{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalLink) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalLink) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ExternalLink) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalLink) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ExternalLink) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ExternalLink) SetDescription(v string) {
	o.Description = v
}

// GetSpecIds returns the SpecIds field value
func (o *ExternalLink) GetSpecIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SpecIds
}

// GetSpecIdsOk returns a tuple with the SpecIds field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetSpecIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecIds, true
}

// SetSpecIds sets field value
func (o *ExternalLink) SetSpecIds(v []string) {
	o.SpecIds = v
}

// GetEntity returns the Entity field value
func (o *ExternalLink) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *ExternalLink) SetEntity(v string) {
	o.Entity = v
}

// GetSource returns the Source field value
func (o *ExternalLink) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ExternalLink) SetSource(v string) {
	o.Source = v
}

// GetJsonPathElement returns the JsonPathElement field value
func (o *ExternalLink) GetJsonPathElement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonPathElement
}

// GetJsonPathElementOk returns a tuple with the JsonPathElement field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetJsonPathElementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonPathElement, true
}

// SetJsonPathElement sets field value
func (o *ExternalLink) SetJsonPathElement(v string) {
	o.JsonPathElement = v
}

// GetUrl returns the Url field value
func (o *ExternalLink) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExternalLink) SetUrl(v string) {
	o.Url = v
}

// GetVendor returns the Vendor field value
func (o *ExternalLink) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *ExternalLink) SetVendor(v string) {
	o.Vendor = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ExternalLink) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ExternalLink) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ExternalLink) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ExternalLink) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ExternalLink) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ExternalLink) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ExternalLink) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ExternalLink) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ExternalLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["spec_ids"] = o.SpecIds
	toSerialize["entity"] = o.Entity
	toSerialize["source"] = o.Source
	toSerialize["json_path_element"] = o.JsonPathElement
	toSerialize["url"] = o.Url
	toSerialize["vendor"] = o.Vendor
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_by"] = o.UpdatedBy
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"spec_ids",
		"entity",
		"source",
		"json_path_element",
		"url",
		"vendor",
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

	varExternalLink := _ExternalLink{}

	err = json.Unmarshal(data, &varExternalLink)

	if err != nil {
		return err
	}

	*o = ExternalLink(varExternalLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "spec_ids")
		delete(additionalProperties, "entity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "json_path_element")
		delete(additionalProperties, "url")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalLink struct {
	value *ExternalLink
	isSet bool
}

func (v NullableExternalLink) Get() *ExternalLink {
	return v.value
}

func (v *NullableExternalLink) Set(val *ExternalLink) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalLink) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalLink(val *ExternalLink) *NullableExternalLink {
	return &NullableExternalLink{value: val, isSet: true}
}

func (v NullableExternalLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
