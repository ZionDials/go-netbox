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

// checks if the BriefDeviceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefDeviceTypeRequest{}

// BriefDeviceTypeRequest Adds support for custom fields and tags.
type BriefDeviceTypeRequest struct {
	Manufacturer         BriefManufacturerRequest `json:"manufacturer"`
	Model                string                   `json:"model"`
	Slug                 string                   `json:"slug"`
	Description          *string                  `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefDeviceTypeRequest BriefDeviceTypeRequest

// NewBriefDeviceTypeRequest instantiates a new BriefDeviceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefDeviceTypeRequest(manufacturer BriefManufacturerRequest, model string, slug string) *BriefDeviceTypeRequest {
	this := BriefDeviceTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	return &this
}

// NewBriefDeviceTypeRequestWithDefaults instantiates a new BriefDeviceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefDeviceTypeRequestWithDefaults() *BriefDeviceTypeRequest {
	this := BriefDeviceTypeRequest{}
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *BriefDeviceTypeRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil {
		var ret BriefManufacturerRequest
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *BriefDeviceTypeRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *BriefDeviceTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BriefDeviceTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *BriefDeviceTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefDeviceTypeRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefDeviceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefDeviceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefDeviceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefDeviceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefDeviceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefDeviceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"manufacturer",
		"model",
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

	varBriefDeviceTypeRequest := _BriefDeviceTypeRequest{}

	err = json.Unmarshal(data, &varBriefDeviceTypeRequest)

	if err != nil {
		return err
	}

	*o = BriefDeviceTypeRequest(varBriefDeviceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefDeviceTypeRequest struct {
	value *BriefDeviceTypeRequest
	isSet bool
}

func (v NullableBriefDeviceTypeRequest) Get() *BriefDeviceTypeRequest {
	return v.value
}

func (v *NullableBriefDeviceTypeRequest) Set(val *BriefDeviceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefDeviceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefDeviceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefDeviceTypeRequest(val *BriefDeviceTypeRequest) *NullableBriefDeviceTypeRequest {
	return &NullableBriefDeviceTypeRequest{value: val, isSet: true}
}

func (v NullableBriefDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefDeviceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
