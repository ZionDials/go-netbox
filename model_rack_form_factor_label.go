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

// RackFormFactorLabel the model 'RackFormFactorLabel'
type RackFormFactorLabel string

// List of Rack_form_factor_label
const (
	RACKFORMFACTORLABEL__2_POST_FRAME                  RackFormFactorLabel = "2-post frame"
	RACKFORMFACTORLABEL__4_POST_FRAME                  RackFormFactorLabel = "4-post frame"
	RACKFORMFACTORLABEL__4_POST_CABINET                RackFormFactorLabel = "4-post cabinet"
	RACKFORMFACTORLABEL_WALL_MOUNTED_FRAME             RackFormFactorLabel = "Wall-mounted frame"
	RACKFORMFACTORLABEL_WALL_MOUNTED_FRAME__VERTICAL   RackFormFactorLabel = "Wall-mounted frame (vertical)"
	RACKFORMFACTORLABEL_WALL_MOUNTED_CABINET           RackFormFactorLabel = "Wall-mounted cabinet"
	RACKFORMFACTORLABEL_WALL_MOUNTED_CABINET__VERTICAL RackFormFactorLabel = "Wall-mounted cabinet (vertical)"
)

// All allowed values of RackFormFactorLabel enum
var AllowedRackFormFactorLabelEnumValues = []RackFormFactorLabel{
	"2-post frame",
	"4-post frame",
	"4-post cabinet",
	"Wall-mounted frame",
	"Wall-mounted frame (vertical)",
	"Wall-mounted cabinet",
	"Wall-mounted cabinet (vertical)",
}

func (v *RackFormFactorLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackFormFactorLabel(value)
	for _, existing := range AllowedRackFormFactorLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackFormFactorLabel", value)
}

// NewRackFormFactorLabelFromValue returns a pointer to a valid RackFormFactorLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackFormFactorLabelFromValue(v string) (*RackFormFactorLabel, error) {
	ev := RackFormFactorLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackFormFactorLabel: valid values are %v", v, AllowedRackFormFactorLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackFormFactorLabel) IsValid() bool {
	for _, existing := range AllowedRackFormFactorLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_form_factor_label value
func (v RackFormFactorLabel) Ptr() *RackFormFactorLabel {
	return &v
}

type NullableRackFormFactorLabel struct {
	value *RackFormFactorLabel
	isSet bool
}

func (v NullableRackFormFactorLabel) Get() *RackFormFactorLabel {
	return v.value
}

func (v *NullableRackFormFactorLabel) Set(val *RackFormFactorLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRackFormFactorLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRackFormFactorLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackFormFactorLabel(val *RackFormFactorLabel) *NullableRackFormFactorLabel {
	return &NullableRackFormFactorLabel{value: val, isSet: true}
}

func (v NullableRackFormFactorLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackFormFactorLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
