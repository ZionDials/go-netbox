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

// PatchedWritableTunnelTerminationRequestRole * `peer` - Peer * `hub` - Hub * `spoke` - Spoke
type PatchedWritableTunnelTerminationRequestRole string

// List of PatchedWritableTunnelTerminationRequest_role
const (
	PATCHEDWRITABLETUNNELTERMINATIONREQUESTROLE_PEER  PatchedWritableTunnelTerminationRequestRole = "peer"
	PATCHEDWRITABLETUNNELTERMINATIONREQUESTROLE_HUB   PatchedWritableTunnelTerminationRequestRole = "hub"
	PATCHEDWRITABLETUNNELTERMINATIONREQUESTROLE_SPOKE PatchedWritableTunnelTerminationRequestRole = "spoke"
)

// All allowed values of PatchedWritableTunnelTerminationRequestRole enum
var AllowedPatchedWritableTunnelTerminationRequestRoleEnumValues = []PatchedWritableTunnelTerminationRequestRole{
	"peer",
	"hub",
	"spoke",
}

func (v *PatchedWritableTunnelTerminationRequestRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableTunnelTerminationRequestRole(value)
	for _, existing := range AllowedPatchedWritableTunnelTerminationRequestRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableTunnelTerminationRequestRole", value)
}

// NewPatchedWritableTunnelTerminationRequestRoleFromValue returns a pointer to a valid PatchedWritableTunnelTerminationRequestRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableTunnelTerminationRequestRoleFromValue(v string) (*PatchedWritableTunnelTerminationRequestRole, error) {
	ev := PatchedWritableTunnelTerminationRequestRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableTunnelTerminationRequestRole: valid values are %v", v, AllowedPatchedWritableTunnelTerminationRequestRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableTunnelTerminationRequestRole) IsValid() bool {
	for _, existing := range AllowedPatchedWritableTunnelTerminationRequestRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableTunnelTerminationRequest_role value
func (v PatchedWritableTunnelTerminationRequestRole) Ptr() *PatchedWritableTunnelTerminationRequestRole {
	return &v
}

type NullablePatchedWritableTunnelTerminationRequestRole struct {
	value *PatchedWritableTunnelTerminationRequestRole
	isSet bool
}

func (v NullablePatchedWritableTunnelTerminationRequestRole) Get() *PatchedWritableTunnelTerminationRequestRole {
	return v.value
}

func (v *NullablePatchedWritableTunnelTerminationRequestRole) Set(val *PatchedWritableTunnelTerminationRequestRole) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableTunnelTerminationRequestRole) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableTunnelTerminationRequestRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableTunnelTerminationRequestRole(val *PatchedWritableTunnelTerminationRequestRole) *NullablePatchedWritableTunnelTerminationRequestRole {
	return &NullablePatchedWritableTunnelTerminationRequestRole{value: val, isSet: true}
}

func (v NullablePatchedWritableTunnelTerminationRequestRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableTunnelTerminationRequestRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
