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

// checks if the WritableJournalEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableJournalEntryRequest{}

// WritableJournalEntryRequest Adds support for custom fields and tags.
type WritableJournalEntryRequest struct {
	AssignedObjectType   string                 `json:"assigned_object_type"`
	AssignedObjectId     int64                  `json:"assigned_object_id"`
	CreatedBy            NullableInt32          `json:"created_by,omitempty"`
	Kind                 *JournalEntryKindValue `json:"kind,omitempty"`
	Comments             string                 `json:"comments"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableJournalEntryRequest WritableJournalEntryRequest

// NewWritableJournalEntryRequest instantiates a new WritableJournalEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableJournalEntryRequest(assignedObjectType string, assignedObjectId int64, comments string) *WritableJournalEntryRequest {
	this := WritableJournalEntryRequest{}
	this.AssignedObjectType = assignedObjectType
	this.AssignedObjectId = assignedObjectId
	this.Comments = comments
	return &this
}

// NewWritableJournalEntryRequestWithDefaults instantiates a new WritableJournalEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableJournalEntryRequestWithDefaults() *WritableJournalEntryRequest {
	this := WritableJournalEntryRequest{}
	return &this
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *WritableJournalEntryRequest) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *WritableJournalEntryRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *WritableJournalEntryRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *WritableJournalEntryRequest) GetAssignedObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *WritableJournalEntryRequest) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *WritableJournalEntryRequest) SetAssignedObjectId(v int64) {
	o.AssignedObjectId = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableJournalEntryRequest) GetCreatedBy() int32 {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret int32
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableJournalEntryRequest) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WritableJournalEntryRequest) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableInt32 and assigns it to the CreatedBy field.
func (o *WritableJournalEntryRequest) SetCreatedBy(v int32) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *WritableJournalEntryRequest) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *WritableJournalEntryRequest) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *WritableJournalEntryRequest) GetKind() JournalEntryKindValue {
	if o == nil || IsNil(o.Kind) {
		var ret JournalEntryKindValue
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableJournalEntryRequest) GetKindOk() (*JournalEntryKindValue, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *WritableJournalEntryRequest) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given JournalEntryKindValue and assigns it to the Kind field.
func (o *WritableJournalEntryRequest) SetKind(v JournalEntryKindValue) {
	o.Kind = &v
}

// GetComments returns the Comments field value
func (o *WritableJournalEntryRequest) GetComments() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *WritableJournalEntryRequest) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *WritableJournalEntryRequest) SetComments(v string) {
	o.Comments = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableJournalEntryRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableJournalEntryRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableJournalEntryRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableJournalEntryRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableJournalEntryRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableJournalEntryRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableJournalEntryRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableJournalEntryRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableJournalEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableJournalEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object_id"] = o.AssignedObjectId
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["comments"] = o.Comments
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

func (o *WritableJournalEntryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assigned_object_type",
		"assigned_object_id",
		"comments",
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

	varWritableJournalEntryRequest := _WritableJournalEntryRequest{}

	err = json.Unmarshal(data, &varWritableJournalEntryRequest)

	if err != nil {
		return err
	}

	*o = WritableJournalEntryRequest(varWritableJournalEntryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableJournalEntryRequest struct {
	value *WritableJournalEntryRequest
	isSet bool
}

func (v NullableWritableJournalEntryRequest) Get() *WritableJournalEntryRequest {
	return v.value
}

func (v *NullableWritableJournalEntryRequest) Set(val *WritableJournalEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableJournalEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableJournalEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableJournalEntryRequest(val *WritableJournalEntryRequest) *NullableWritableJournalEntryRequest {
	return &NullableWritableJournalEntryRequest{value: val, isSet: true}
}

func (v NullableWritableJournalEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableJournalEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
