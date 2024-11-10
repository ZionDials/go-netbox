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

// checks if the CustomFieldUiEditable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldUiEditable{}

// CustomFieldUiEditable struct for CustomFieldUiEditable
type CustomFieldUiEditable struct {
	Value                *CustomFieldUiEditableValue `json:"value,omitempty"`
	Label                *CustomFieldUiEditableLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldUiEditable CustomFieldUiEditable

// NewCustomFieldUiEditable instantiates a new CustomFieldUiEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldUiEditable() *CustomFieldUiEditable {
	this := CustomFieldUiEditable{}
	return &this
}

// NewCustomFieldUiEditableWithDefaults instantiates a new CustomFieldUiEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldUiEditableWithDefaults() *CustomFieldUiEditable {
	this := CustomFieldUiEditable{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomFieldUiEditable) GetValue() CustomFieldUiEditableValue {
	if o == nil || IsNil(o.Value) {
		var ret CustomFieldUiEditableValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldUiEditable) GetValueOk() (*CustomFieldUiEditableValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldUiEditable) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CustomFieldUiEditableValue and assigns it to the Value field.
func (o *CustomFieldUiEditable) SetValue(v CustomFieldUiEditableValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomFieldUiEditable) GetLabel() CustomFieldUiEditableLabel {
	if o == nil || IsNil(o.Label) {
		var ret CustomFieldUiEditableLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldUiEditable) GetLabelOk() (*CustomFieldUiEditableLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomFieldUiEditable) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CustomFieldUiEditableLabel and assigns it to the Label field.
func (o *CustomFieldUiEditable) SetLabel(v CustomFieldUiEditableLabel) {
	o.Label = &v
}

func (o CustomFieldUiEditable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldUiEditable) ToMap() (map[string]interface{}, error) {
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

func (o *CustomFieldUiEditable) UnmarshalJSON(data []byte) (err error) {
	varCustomFieldUiEditable := _CustomFieldUiEditable{}

	err = json.Unmarshal(data, &varCustomFieldUiEditable)

	if err != nil {
		return err
	}

	*o = CustomFieldUiEditable(varCustomFieldUiEditable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldUiEditable struct {
	value *CustomFieldUiEditable
	isSet bool
}

func (v NullableCustomFieldUiEditable) Get() *CustomFieldUiEditable {
	return v.value
}

func (v *NullableCustomFieldUiEditable) Set(val *CustomFieldUiEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldUiEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldUiEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldUiEditable(val *CustomFieldUiEditable) *NullableCustomFieldUiEditable {
	return &NullableCustomFieldUiEditable{value: val, isSet: true}
}

func (v NullableCustomFieldUiEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldUiEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
