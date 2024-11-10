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

// ServiceProtocolLabel the model 'ServiceProtocolLabel'
type ServiceProtocolLabel string

// List of Service_protocol_label
const (
	SERVICEPROTOCOLLABEL_TCP  ServiceProtocolLabel = "TCP"
	SERVICEPROTOCOLLABEL_UDP  ServiceProtocolLabel = "UDP"
	SERVICEPROTOCOLLABEL_SCTP ServiceProtocolLabel = "SCTP"
)

// All allowed values of ServiceProtocolLabel enum
var AllowedServiceProtocolLabelEnumValues = []ServiceProtocolLabel{
	"TCP",
	"UDP",
	"SCTP",
}

func (v *ServiceProtocolLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceProtocolLabel(value)
	for _, existing := range AllowedServiceProtocolLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceProtocolLabel", value)
}

// NewServiceProtocolLabelFromValue returns a pointer to a valid ServiceProtocolLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceProtocolLabelFromValue(v string) (*ServiceProtocolLabel, error) {
	ev := ServiceProtocolLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceProtocolLabel: valid values are %v", v, AllowedServiceProtocolLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceProtocolLabel) IsValid() bool {
	for _, existing := range AllowedServiceProtocolLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Service_protocol_label value
func (v ServiceProtocolLabel) Ptr() *ServiceProtocolLabel {
	return &v
}

type NullableServiceProtocolLabel struct {
	value *ServiceProtocolLabel
	isSet bool
}

func (v NullableServiceProtocolLabel) Get() *ServiceProtocolLabel {
	return v.value
}

func (v *NullableServiceProtocolLabel) Set(val *ServiceProtocolLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProtocolLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProtocolLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProtocolLabel(val *ServiceProtocolLabel) *NullableServiceProtocolLabel {
	return &NullableServiceProtocolLabel{value: val, isSet: true}
}

func (v NullableServiceProtocolLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProtocolLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
