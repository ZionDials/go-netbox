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

// checks if the BriefRack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefRack{}

// BriefRack Adds support for custom fields and tags.
type BriefRack struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	DeviceCount          int64   `json:"device_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefRack BriefRack

// NewBriefRack instantiates a new BriefRack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefRack(id int32, url string, display string, name string, deviceCount int64) *BriefRack {
	this := BriefRack{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.DeviceCount = deviceCount
	return &this
}

// NewBriefRackWithDefaults instantiates a new BriefRack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefRackWithDefaults() *BriefRack {
	this := BriefRack{}
	return &this
}

// GetId returns the Id field value
func (o *BriefRack) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefRack) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefRack) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefRack) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefRack) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefRack) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefRack) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefRack) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefRack) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefRack) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefRack) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefRack) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefRack) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefRack) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefRack) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefRack) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceCount returns the DeviceCount field value
func (o *BriefRack) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *BriefRack) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *BriefRack) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

func (o BriefRack) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefRack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["device_count"] = o.DeviceCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefRack) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
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

	varBriefRack := _BriefRack{}

	err = json.Unmarshal(data, &varBriefRack)

	if err != nil {
		return err
	}

	*o = BriefRack(varBriefRack)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefRack struct {
	value *BriefRack
	isSet bool
}

func (v NullableBriefRack) Get() *BriefRack {
	return v.value
}

func (v *NullableBriefRack) Set(val *BriefRack) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefRack) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefRack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefRack(val *BriefRack) *NullableBriefRack {
	return &NullableBriefRack{value: val, isSet: true}
}

func (v NullableBriefRack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefRack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
