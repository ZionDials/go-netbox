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

// checks if the AvailableASN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableASN{}

// AvailableASN Representation of an ASN which does not exist in the database.
type AvailableASN struct {
	Asn                  int32   `json:"asn"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvailableASN AvailableASN

// NewAvailableASN instantiates a new AvailableASN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableASN(asn int32) *AvailableASN {
	this := AvailableASN{}
	this.Asn = asn
	return &this
}

// NewAvailableASNWithDefaults instantiates a new AvailableASN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableASNWithDefaults() *AvailableASN {
	this := AvailableASN{}
	return &this
}

// GetAsn returns the Asn field value
func (o *AvailableASN) GetAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *AvailableASN) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *AvailableASN) SetAsn(v int32) {
	o.Asn = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AvailableASN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableASN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AvailableASN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AvailableASN) SetDescription(v string) {
	o.Description = &v
}

func (o AvailableASN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableASN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asn"] = o.Asn
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailableASN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asn",
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

	varAvailableASN := _AvailableASN{}

	err = json.Unmarshal(data, &varAvailableASN)

	if err != nil {
		return err
	}

	*o = AvailableASN(varAvailableASN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asn")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailableASN struct {
	value *AvailableASN
	isSet bool
}

func (v NullableAvailableASN) Get() *AvailableASN {
	return v.value
}

func (v *NullableAvailableASN) Set(val *AvailableASN) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableASN) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableASN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableASN(val *AvailableASN) *NullableAvailableASN {
	return &NullableAvailableASN{value: val, isSet: true}
}

func (v NullableAvailableASN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableASN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
