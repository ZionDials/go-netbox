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

// checks if the TunnelTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelTerminationRequest{}

// TunnelTerminationRequest Adds support for custom fields and tags.
type TunnelTerminationRequest struct {
	Tunnel               BriefTunnelRequest                          `json:"tunnel"`
	Role                 PatchedWritableTunnelTerminationRequestRole `json:"role"`
	TerminationType      string                                      `json:"termination_type"`
	TerminationId        NullableInt64                               `json:"termination_id"`
	OutsideIp            NullableBriefIPAddressRequest               `json:"outside_ip,omitempty"`
	Tags                 []NestedTagRequest                          `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                      `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelTerminationRequest TunnelTerminationRequest

// NewTunnelTerminationRequest instantiates a new TunnelTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelTerminationRequest(tunnel BriefTunnelRequest, role PatchedWritableTunnelTerminationRequestRole, terminationType string, terminationId NullableInt64) *TunnelTerminationRequest {
	this := TunnelTerminationRequest{}
	this.Tunnel = tunnel
	this.Role = role
	this.TerminationType = terminationType
	this.TerminationId = terminationId
	return &this
}

// NewTunnelTerminationRequestWithDefaults instantiates a new TunnelTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelTerminationRequestWithDefaults() *TunnelTerminationRequest {
	this := TunnelTerminationRequest{}
	return &this
}

// GetTunnel returns the Tunnel field value
func (o *TunnelTerminationRequest) GetTunnel() BriefTunnelRequest {
	if o == nil {
		var ret BriefTunnelRequest
		return ret
	}

	return o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRequest) GetTunnelOk() (*BriefTunnelRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tunnel, true
}

// SetTunnel sets field value
func (o *TunnelTerminationRequest) SetTunnel(v BriefTunnelRequest) {
	o.Tunnel = v
}

// GetRole returns the Role field value
func (o *TunnelTerminationRequest) GetRole() PatchedWritableTunnelTerminationRequestRole {
	if o == nil {
		var ret PatchedWritableTunnelTerminationRequestRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *TunnelTerminationRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole) {
	o.Role = v
}

// GetTerminationType returns the TerminationType field value
func (o *TunnelTerminationRequest) GetTerminationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationType
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRequest) GetTerminationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationType, true
}

// SetTerminationType sets field value
func (o *TunnelTerminationRequest) SetTerminationType(v string) {
	o.TerminationType = v
}

// GetTerminationId returns the TerminationId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *TunnelTerminationRequest) GetTerminationId() int64 {
	if o == nil || o.TerminationId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.TerminationId.Get()
}

// GetTerminationIdOk returns a tuple with the TerminationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTerminationRequest) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationId.Get(), o.TerminationId.IsSet()
}

// SetTerminationId sets field value
func (o *TunnelTerminationRequest) SetTerminationId(v int64) {
	o.TerminationId.Set(&v)
}

// GetOutsideIp returns the OutsideIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelTerminationRequest) GetOutsideIp() BriefIPAddressRequest {
	if o == nil || IsNil(o.OutsideIp.Get()) {
		var ret BriefIPAddressRequest
		return ret
	}
	return *o.OutsideIp.Get()
}

// GetOutsideIpOk returns a tuple with the OutsideIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelTerminationRequest) GetOutsideIpOk() (*BriefIPAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutsideIp.Get(), o.OutsideIp.IsSet()
}

// HasOutsideIp returns a boolean if a field has been set.
func (o *TunnelTerminationRequest) HasOutsideIp() bool {
	if o != nil && o.OutsideIp.IsSet() {
		return true
	}

	return false
}

// SetOutsideIp gets a reference to the given NullableBriefIPAddressRequest and assigns it to the OutsideIp field.
func (o *TunnelTerminationRequest) SetOutsideIp(v BriefIPAddressRequest) {
	o.OutsideIp.Set(&v)
}

// SetOutsideIpNil sets the value for OutsideIp to be an explicit nil
func (o *TunnelTerminationRequest) SetOutsideIpNil() {
	o.OutsideIp.Set(nil)
}

// UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
func (o *TunnelTerminationRequest) UnsetOutsideIp() {
	o.OutsideIp.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TunnelTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TunnelTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *TunnelTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TunnelTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TunnelTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *TunnelTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o TunnelTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tunnel"] = o.Tunnel
	toSerialize["role"] = o.Role
	toSerialize["termination_type"] = o.TerminationType
	toSerialize["termination_id"] = o.TerminationId.Get()
	if o.OutsideIp.IsSet() {
		toSerialize["outside_ip"] = o.OutsideIp.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tunnel",
		"role",
		"termination_type",
		"termination_id",
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

	varTunnelTerminationRequest := _TunnelTerminationRequest{}

	err = json.Unmarshal(data, &varTunnelTerminationRequest)

	if err != nil {
		return err
	}

	*o = TunnelTerminationRequest(varTunnelTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tunnel")
		delete(additionalProperties, "role")
		delete(additionalProperties, "termination_type")
		delete(additionalProperties, "termination_id")
		delete(additionalProperties, "outside_ip")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelTerminationRequest struct {
	value *TunnelTerminationRequest
	isSet bool
}

func (v NullableTunnelTerminationRequest) Get() *TunnelTerminationRequest {
	return v.value
}

func (v *NullableTunnelTerminationRequest) Set(val *TunnelTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelTerminationRequest(val *TunnelTerminationRequest) *NullableTunnelTerminationRequest {
	return &NullableTunnelTerminationRequest{value: val, isSet: true}
}

func (v NullableTunnelTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
