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

// checks if the NestedRegionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRegionRequest{}

// NestedRegionRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedRegionRequest struct {
	Name                 string `json:"name"`
	Slug                 string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	AdditionalProperties map[string]interface{}
}

type _NestedRegionRequest NestedRegionRequest

// NewNestedRegionRequest instantiates a new NestedRegionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRegionRequest(name string, slug string) *NestedRegionRequest {
	this := NestedRegionRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedRegionRequestWithDefaults instantiates a new NestedRegionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRegionRequestWithDefaults() *NestedRegionRequest {
	this := NestedRegionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedRegionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRegionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRegionRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedRegionRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedRegionRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedRegionRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedRegionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRegionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRegionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varNestedRegionRequest := _NestedRegionRequest{}

	err = json.Unmarshal(data, &varNestedRegionRequest)

	if err != nil {
		return err
	}

	*o = NestedRegionRequest(varNestedRegionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRegionRequest struct {
	value *NestedRegionRequest
	isSet bool
}

func (v NullableNestedRegionRequest) Get() *NestedRegionRequest {
	return v.value
}

func (v *NullableNestedRegionRequest) Set(val *NestedRegionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRegionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRegionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRegionRequest(val *NestedRegionRequest) *NullableNestedRegionRequest {
	return &NullableNestedRegionRequest{value: val, isSet: true}
}

func (v NullableNestedRegionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRegionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
