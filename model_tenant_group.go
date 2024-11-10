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
	"time"
)

// checks if the TenantGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantGroup{}

// TenantGroup Extends PrimaryModelSerializer to include MPTT support.
type TenantGroup struct {
	Id                   int32                     `json:"id"`
	Url                  string                    `json:"url"`
	DisplayUrl           string                    `json:"display_url"`
	Display              string                    `json:"display"`
	Name                 string                    `json:"name"`
	Slug                 string                    `json:"slug"`
	Parent               NullableNestedTenantGroup `json:"parent,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	Tags                 []NestedTag               `json:"tags,omitempty"`
	CustomFields         map[string]interface{}    `json:"custom_fields,omitempty"`
	Created              NullableTime              `json:"created"`
	LastUpdated          NullableTime              `json:"last_updated"`
	TenantCount          int32                     `json:"tenant_count"`
	Depth                int32                     `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _TenantGroup TenantGroup

// NewTenantGroup instantiates a new TenantGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantGroup(id int32, url string, displayUrl string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, tenantCount int32, depth int32) *TenantGroup {
	this := TenantGroup{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.TenantCount = tenantCount
	this.Depth = depth
	return &this
}

// NewTenantGroupWithDefaults instantiates a new TenantGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantGroupWithDefaults() *TenantGroup {
	this := TenantGroup{}
	return &this
}

// GetId returns the Id field value
func (o *TenantGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenantGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *TenantGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TenantGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *TenantGroup) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *TenantGroup) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *TenantGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *TenantGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *TenantGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *TenantGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *TenantGroup) SetSlug(v string) {
	o.Slug = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantGroup) GetParent() NestedTenantGroup {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret NestedTenantGroup
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantGroup) GetParentOk() (*NestedTenantGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *TenantGroup) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableNestedTenantGroup and assigns it to the Parent field.
func (o *TenantGroup) SetParent(v NestedTenantGroup) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *TenantGroup) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *TenantGroup) UnsetParent() {
	o.Parent.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TenantGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TenantGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TenantGroup) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TenantGroup) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TenantGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *TenantGroup) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TenantGroup) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TenantGroup) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *TenantGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TenantGroup) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *TenantGroup) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TenantGroup) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *TenantGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetTenantCount returns the TenantCount field value
func (o *TenantGroup) GetTenantCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantCount
}

// GetTenantCountOk returns a tuple with the TenantCount field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetTenantCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantCount, true
}

// SetTenantCount sets field value
func (o *TenantGroup) SetTenantCount(v int32) {
	o.TenantCount = v
}

// GetDepth returns the Depth field value
func (o *TenantGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *TenantGroup) SetDepth(v int32) {
	o.Depth = v
}

func (o TenantGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["tenant_count"] = o.TenantCount
	toSerialize["_depth"] = o.Depth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TenantGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"slug",
		"created",
		"last_updated",
		"tenant_count",
		"_depth",
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

	varTenantGroup := _TenantGroup{}

	err = json.Unmarshal(data, &varTenantGroup)

	if err != nil {
		return err
	}

	*o = TenantGroup(varTenantGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "tenant_count")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantGroup struct {
	value *TenantGroup
	isSet bool
}

func (v NullableTenantGroup) Get() *TenantGroup {
	return v.value
}

func (v *NullableTenantGroup) Set(val *TenantGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantGroup(val *TenantGroup) *NullableTenantGroup {
	return &NullableTenantGroup{value: val, isSet: true}
}

func (v NullableTenantGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
