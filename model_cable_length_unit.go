/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.6 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the CableLengthUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableLengthUnit{}

// CableLengthUnit struct for CableLengthUnit
type CableLengthUnit struct {
	Value                *CableLengthUnitValue `json:"value,omitempty"`
	Label                *CableLengthUnitLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CableLengthUnit CableLengthUnit

// NewCableLengthUnit instantiates a new CableLengthUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableLengthUnit() *CableLengthUnit {
	this := CableLengthUnit{}
	return &this
}

// NewCableLengthUnitWithDefaults instantiates a new CableLengthUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableLengthUnitWithDefaults() *CableLengthUnit {
	this := CableLengthUnit{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CableLengthUnit) GetValue() CableLengthUnitValue {
	if o == nil || IsNil(o.Value) {
		var ret CableLengthUnitValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableLengthUnit) GetValueOk() (*CableLengthUnitValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CableLengthUnit) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CableLengthUnitValue and assigns it to the Value field.
func (o *CableLengthUnit) SetValue(v CableLengthUnitValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CableLengthUnit) GetLabel() CableLengthUnitLabel {
	if o == nil || IsNil(o.Label) {
		var ret CableLengthUnitLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableLengthUnit) GetLabelOk() (*CableLengthUnitLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CableLengthUnit) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CableLengthUnitLabel and assigns it to the Label field.
func (o *CableLengthUnit) SetLabel(v CableLengthUnitLabel) {
	o.Label = &v
}

func (o CableLengthUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableLengthUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CableLengthUnit) UnmarshalJSON(data []byte) (err error) {
	varCableLengthUnit := _CableLengthUnit{}

	err = json.Unmarshal(data, &varCableLengthUnit)

	if err != nil {
		return err
	}

	*o = CableLengthUnit(varCableLengthUnit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCableLengthUnit struct {
	value *CableLengthUnit
	isSet bool
}

func (v NullableCableLengthUnit) Get() *CableLengthUnit {
	return v.value
}

func (v *NullableCableLengthUnit) Set(val *CableLengthUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCableLengthUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCableLengthUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableLengthUnit(val *CableLengthUnit) *NullableCableLengthUnit {
	return &NullableCableLengthUnit{value: val, isSet: true}
}

func (v NullableCableLengthUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableLengthUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
