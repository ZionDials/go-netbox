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

// checks if the BriefDeviceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefDeviceType{}

// BriefDeviceType Adds support for custom fields and tags.
type BriefDeviceType struct {
	Id                   int32             `json:"id"`
	Url                  string            `json:"url"`
	Display              string            `json:"display"`
	Manufacturer         BriefManufacturer `json:"manufacturer"`
	Model                string            `json:"model"`
	Slug                 string            `json:"slug"`
	Description          *string           `json:"description,omitempty"`
	DeviceCount          int64             `json:"device_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefDeviceType BriefDeviceType

// NewBriefDeviceType instantiates a new BriefDeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefDeviceType(id int32, url string, display string, manufacturer BriefManufacturer, model string, slug string, deviceCount int64) *BriefDeviceType {
	this := BriefDeviceType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	this.DeviceCount = deviceCount
	return &this
}

// NewBriefDeviceTypeWithDefaults instantiates a new BriefDeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefDeviceTypeWithDefaults() *BriefDeviceType {
	this := BriefDeviceType{}
	return &this
}

// GetId returns the Id field value
func (o *BriefDeviceType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefDeviceType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefDeviceType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefDeviceType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefDeviceType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefDeviceType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *BriefDeviceType) GetManufacturer() BriefManufacturer {
	if o == nil {
		var ret BriefManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *BriefDeviceType) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *BriefDeviceType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BriefDeviceType) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *BriefDeviceType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefDeviceType) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefDeviceType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefDeviceType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefDeviceType) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceCount returns the DeviceCount field value
func (o *BriefDeviceType) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *BriefDeviceType) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *BriefDeviceType) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

func (o BriefDeviceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefDeviceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["device_count"] = o.DeviceCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefDeviceType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"manufacturer",
		"model",
		"slug",
		"device_count",
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

	varBriefDeviceType := _BriefDeviceType{}

	err = json.Unmarshal(data, &varBriefDeviceType)

	if err != nil {
		return err
	}

	*o = BriefDeviceType(varBriefDeviceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefDeviceType struct {
	value *BriefDeviceType
	isSet bool
}

func (v NullableBriefDeviceType) Get() *BriefDeviceType {
	return v.value
}

func (v *NullableBriefDeviceType) Set(val *BriefDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefDeviceType(val *BriefDeviceType) *NullableBriefDeviceType {
	return &NullableBriefDeviceType{value: val, isSet: true}
}

func (v NullableBriefDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
