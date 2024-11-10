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

// Termination the model 'Termination'
type Termination string

// List of Termination
const (
	TERMINATION_A Termination = "A"
	TERMINATION_Z Termination = "Z"
)

// All allowed values of Termination enum
var AllowedTerminationEnumValues = []Termination{
	"A",
	"Z",
}

func (v *Termination) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Termination(value)
	for _, existing := range AllowedTerminationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Termination", value)
}

// NewTerminationFromValue returns a pointer to a valid Termination
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTerminationFromValue(v string) (*Termination, error) {
	ev := Termination(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Termination: valid values are %v", v, AllowedTerminationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Termination) IsValid() bool {
	for _, existing := range AllowedTerminationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Termination value
func (v Termination) Ptr() *Termination {
	return &v
}

type NullableTermination struct {
	value *Termination
	isSet bool
}

func (v NullableTermination) Get() *Termination {
	return v.value
}

func (v *NullableTermination) Set(val *Termination) {
	v.value = val
	v.isSet = true
}

func (v NullableTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermination(val *Termination) *NullableTermination {
	return &NullableTermination{value: val, isSet: true}
}

func (v NullableTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
