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

// DeviceStatusValue * `offline` - Offline * `active` - Active * `planned` - Planned * `staged` - Staged * `failed` - Failed * `inventory` - Inventory * `decommissioning` - Decommissioning
type DeviceStatusValue string

// List of Device_status_value
const (
	DEVICESTATUSVALUE_OFFLINE         DeviceStatusValue = "offline"
	DEVICESTATUSVALUE_ACTIVE          DeviceStatusValue = "active"
	DEVICESTATUSVALUE_PLANNED         DeviceStatusValue = "planned"
	DEVICESTATUSVALUE_STAGED          DeviceStatusValue = "staged"
	DEVICESTATUSVALUE_FAILED          DeviceStatusValue = "failed"
	DEVICESTATUSVALUE_INVENTORY       DeviceStatusValue = "inventory"
	DEVICESTATUSVALUE_DECOMMISSIONING DeviceStatusValue = "decommissioning"
)

// All allowed values of DeviceStatusValue enum
var AllowedDeviceStatusValueEnumValues = []DeviceStatusValue{
	"offline",
	"active",
	"planned",
	"staged",
	"failed",
	"inventory",
	"decommissioning",
}

func (v *DeviceStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceStatusValue(value)
	for _, existing := range AllowedDeviceStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceStatusValue", value)
}

// NewDeviceStatusValueFromValue returns a pointer to a valid DeviceStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceStatusValueFromValue(v string) (*DeviceStatusValue, error) {
	ev := DeviceStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceStatusValue: valid values are %v", v, AllowedDeviceStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceStatusValue) IsValid() bool {
	for _, existing := range AllowedDeviceStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Device_status_value value
func (v DeviceStatusValue) Ptr() *DeviceStatusValue {
	return &v
}

type NullableDeviceStatusValue struct {
	value *DeviceStatusValue
	isSet bool
}

func (v NullableDeviceStatusValue) Get() *DeviceStatusValue {
	return v.value
}

func (v *NullableDeviceStatusValue) Set(val *DeviceStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceStatusValue(val *DeviceStatusValue) *NullableDeviceStatusValue {
	return &NullableDeviceStatusValue{value: val, isSet: true}
}

func (v NullableDeviceStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
