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

// DcimDeviceTypesListWeightUnitParameter the model 'DcimDeviceTypesListWeightUnitParameter'
type DcimDeviceTypesListWeightUnitParameter string

// List of dcim_device_types_list_weight_unit_parameter
const (
	DCIMDEVICETYPESLISTWEIGHTUNITPARAMETER_G  DcimDeviceTypesListWeightUnitParameter = "g"
	DCIMDEVICETYPESLISTWEIGHTUNITPARAMETER_KG DcimDeviceTypesListWeightUnitParameter = "kg"
	DCIMDEVICETYPESLISTWEIGHTUNITPARAMETER_LB DcimDeviceTypesListWeightUnitParameter = "lb"
	DCIMDEVICETYPESLISTWEIGHTUNITPARAMETER_OZ DcimDeviceTypesListWeightUnitParameter = "oz"
)

// All allowed values of DcimDeviceTypesListWeightUnitParameter enum
var AllowedDcimDeviceTypesListWeightUnitParameterEnumValues = []DcimDeviceTypesListWeightUnitParameter{
	"g",
	"kg",
	"lb",
	"oz",
}

func (v *DcimDeviceTypesListWeightUnitParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimDeviceTypesListWeightUnitParameter(value)
	for _, existing := range AllowedDcimDeviceTypesListWeightUnitParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimDeviceTypesListWeightUnitParameter", value)
}

// NewDcimDeviceTypesListWeightUnitParameterFromValue returns a pointer to a valid DcimDeviceTypesListWeightUnitParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimDeviceTypesListWeightUnitParameterFromValue(v string) (*DcimDeviceTypesListWeightUnitParameter, error) {
	ev := DcimDeviceTypesListWeightUnitParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimDeviceTypesListWeightUnitParameter: valid values are %v", v, AllowedDcimDeviceTypesListWeightUnitParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimDeviceTypesListWeightUnitParameter) IsValid() bool {
	for _, existing := range AllowedDcimDeviceTypesListWeightUnitParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_device_types_list_weight_unit_parameter value
func (v DcimDeviceTypesListWeightUnitParameter) Ptr() *DcimDeviceTypesListWeightUnitParameter {
	return &v
}

type NullableDcimDeviceTypesListWeightUnitParameter struct {
	value *DcimDeviceTypesListWeightUnitParameter
	isSet bool
}

func (v NullableDcimDeviceTypesListWeightUnitParameter) Get() *DcimDeviceTypesListWeightUnitParameter {
	return v.value
}

func (v *NullableDcimDeviceTypesListWeightUnitParameter) Set(val *DcimDeviceTypesListWeightUnitParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimDeviceTypesListWeightUnitParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimDeviceTypesListWeightUnitParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimDeviceTypesListWeightUnitParameter(val *DcimDeviceTypesListWeightUnitParameter) *NullableDcimDeviceTypesListWeightUnitParameter {
	return &NullableDcimDeviceTypesListWeightUnitParameter{value: val, isSet: true}
}

func (v NullableDcimDeviceTypesListWeightUnitParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimDeviceTypesListWeightUnitParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
