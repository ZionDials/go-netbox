/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.6 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the BriefSiteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefSiteRequest{}

// BriefSiteRequest Adds support for custom fields and tags.
type BriefSiteRequest struct {
	// Full name of the site
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefSiteRequest BriefSiteRequest

// NewBriefSiteRequest instantiates a new BriefSiteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefSiteRequest(name string, slug string) *BriefSiteRequest {
	this := BriefSiteRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefSiteRequestWithDefaults instantiates a new BriefSiteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefSiteRequestWithDefaults() *BriefSiteRequest {
	this := BriefSiteRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefSiteRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefSiteRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefSiteRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefSiteRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefSiteRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefSiteRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefSiteRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefSiteRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefSiteRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefSiteRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefSiteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefSiteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefSiteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varBriefSiteRequest := _BriefSiteRequest{}

	err = json.Unmarshal(data, &varBriefSiteRequest)

	if err != nil {
		return err
	}

	*o = BriefSiteRequest(varBriefSiteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefSiteRequest struct {
	value *BriefSiteRequest
	isSet bool
}

func (v NullableBriefSiteRequest) Get() *BriefSiteRequest {
	return v.value
}

func (v *NullableBriefSiteRequest) Set(val *BriefSiteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefSiteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefSiteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefSiteRequest(val *BriefSiteRequest) *NullableBriefSiteRequest {
	return &NullableBriefSiteRequest{value: val, isSet: true}
}

func (v NullableBriefSiteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefSiteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
