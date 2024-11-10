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

// BriefFHRPGroupProtocol * `vrrp2` - VRRPv2 * `vrrp3` - VRRPv3 * `carp` - CARP * `clusterxl` - ClusterXL * `hsrp` - HSRP * `glbp` - GLBP * `other` - Other
type BriefFHRPGroupProtocol string

// List of BriefFHRPGroup_protocol
const (
	BRIEFFHRPGROUPPROTOCOL_VRRP2     BriefFHRPGroupProtocol = "vrrp2"
	BRIEFFHRPGROUPPROTOCOL_VRRP3     BriefFHRPGroupProtocol = "vrrp3"
	BRIEFFHRPGROUPPROTOCOL_CARP      BriefFHRPGroupProtocol = "carp"
	BRIEFFHRPGROUPPROTOCOL_CLUSTERXL BriefFHRPGroupProtocol = "clusterxl"
	BRIEFFHRPGROUPPROTOCOL_HSRP      BriefFHRPGroupProtocol = "hsrp"
	BRIEFFHRPGROUPPROTOCOL_GLBP      BriefFHRPGroupProtocol = "glbp"
	BRIEFFHRPGROUPPROTOCOL_OTHER     BriefFHRPGroupProtocol = "other"
)

// All allowed values of BriefFHRPGroupProtocol enum
var AllowedBriefFHRPGroupProtocolEnumValues = []BriefFHRPGroupProtocol{
	"vrrp2",
	"vrrp3",
	"carp",
	"clusterxl",
	"hsrp",
	"glbp",
	"other",
}

func (v *BriefFHRPGroupProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BriefFHRPGroupProtocol(value)
	for _, existing := range AllowedBriefFHRPGroupProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BriefFHRPGroupProtocol", value)
}

// NewBriefFHRPGroupProtocolFromValue returns a pointer to a valid BriefFHRPGroupProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBriefFHRPGroupProtocolFromValue(v string) (*BriefFHRPGroupProtocol, error) {
	ev := BriefFHRPGroupProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BriefFHRPGroupProtocol: valid values are %v", v, AllowedBriefFHRPGroupProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BriefFHRPGroupProtocol) IsValid() bool {
	for _, existing := range AllowedBriefFHRPGroupProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BriefFHRPGroup_protocol value
func (v BriefFHRPGroupProtocol) Ptr() *BriefFHRPGroupProtocol {
	return &v
}

type NullableBriefFHRPGroupProtocol struct {
	value *BriefFHRPGroupProtocol
	isSet bool
}

func (v NullableBriefFHRPGroupProtocol) Get() *BriefFHRPGroupProtocol {
	return v.value
}

func (v *NullableBriefFHRPGroupProtocol) Set(val *BriefFHRPGroupProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefFHRPGroupProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefFHRPGroupProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefFHRPGroupProtocol(val *BriefFHRPGroupProtocol) *NullableBriefFHRPGroupProtocol {
	return &NullableBriefFHRPGroupProtocol{value: val, isSet: true}
}

func (v NullableBriefFHRPGroupProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefFHRPGroupProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
