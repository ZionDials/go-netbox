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

// checks if the WritableCustomFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableCustomFieldRequest{}

// WritableCustomFieldRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableCustomFieldRequest struct {
	ObjectTypes       []string                               `json:"object_types"`
	Type              *PatchedWritableCustomFieldRequestType `json:"type,omitempty"`
	RelatedObjectType NullableString                         `json:"related_object_type,omitempty"`
	// Internal field name
	Name string `json:"name"`
	// Name of the field as displayed to users (if not provided, 'the field's name will be used)
	Label *string `json:"label,omitempty"`
	// Custom fields within the same group will be displayed together
	GroupName   *string `json:"group_name,omitempty"`
	Description *string `json:"description,omitempty"`
	// This field is required when creating new objects or editing an existing object.
	Required *bool `json:"required,omitempty"`
	// The value of this field must be unique for the assigned object
	Unique *bool `json:"unique,omitempty"`
	// Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored.
	SearchWeight *int32                                        `json:"search_weight,omitempty"`
	FilterLogic  *PatchedWritableCustomFieldRequestFilterLogic `json:"filter_logic,omitempty"`
	UiVisible    *PatchedWritableCustomFieldRequestUiVisible   `json:"ui_visible,omitempty"`
	UiEditable   *PatchedWritableCustomFieldRequestUiEditable  `json:"ui_editable,omitempty"`
	// Replicate this value when cloning objects
	IsCloneable *bool `json:"is_cloneable,omitempty"`
	// Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \"Foo\").
	Default interface{} `json:"default,omitempty"`
	// Filter the object selection choices using a query_params dict (must be a JSON value).Encapsulate strings with double quotes (e.g. \"Foo\").
	RelatedObjectFilter interface{} `json:"related_object_filter,omitempty"`
	// Fields with higher weights appear lower in a form.
	Weight *int32 `json:"weight,omitempty"`
	// Minimum allowed value (for numeric fields)
	ValidationMinimum NullableInt64 `json:"validation_minimum,omitempty"`
	// Maximum allowed value (for numeric fields)
	ValidationMaximum NullableInt64 `json:"validation_maximum,omitempty"`
	// Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters.
	ValidationRegex      *string                                  `json:"validation_regex,omitempty"`
	ChoiceSet            NullableBriefCustomFieldChoiceSetRequest `json:"choice_set,omitempty"`
	Comments             *string                                  `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableCustomFieldRequest WritableCustomFieldRequest

// NewWritableCustomFieldRequest instantiates a new WritableCustomFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableCustomFieldRequest(objectTypes []string, name string) *WritableCustomFieldRequest {
	this := WritableCustomFieldRequest{}
	this.ObjectTypes = objectTypes
	this.Name = name
	return &this
}

// NewWritableCustomFieldRequestWithDefaults instantiates a new WritableCustomFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableCustomFieldRequestWithDefaults() *WritableCustomFieldRequest {
	this := WritableCustomFieldRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value
func (o *WritableCustomFieldRequest) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *WritableCustomFieldRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetType() PatchedWritableCustomFieldRequestType {
	if o == nil || IsNil(o.Type) {
		var ret PatchedWritableCustomFieldRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetTypeOk() (*PatchedWritableCustomFieldRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritableCustomFieldRequestType and assigns it to the Type field.
func (o *WritableCustomFieldRequest) SetType(v PatchedWritableCustomFieldRequestType) {
	o.Type = &v
}

// GetRelatedObjectType returns the RelatedObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldRequest) GetRelatedObjectType() string {
	if o == nil || IsNil(o.RelatedObjectType.Get()) {
		var ret string
		return ret
	}
	return *o.RelatedObjectType.Get()
}

// GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldRequest) GetRelatedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelatedObjectType.Get(), o.RelatedObjectType.IsSet()
}

// HasRelatedObjectType returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasRelatedObjectType() bool {
	if o != nil && o.RelatedObjectType.IsSet() {
		return true
	}

	return false
}

// SetRelatedObjectType gets a reference to the given NullableString and assigns it to the RelatedObjectType field.
func (o *WritableCustomFieldRequest) SetRelatedObjectType(v string) {
	o.RelatedObjectType.Set(&v)
}

// SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil
func (o *WritableCustomFieldRequest) SetRelatedObjectTypeNil() {
	o.RelatedObjectType.Set(nil)
}

// UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
func (o *WritableCustomFieldRequest) UnsetRelatedObjectType() {
	o.RelatedObjectType.Unset()
}

// GetName returns the Name field value
func (o *WritableCustomFieldRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableCustomFieldRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableCustomFieldRequest) SetLabel(v string) {
	o.Label = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *WritableCustomFieldRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableCustomFieldRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *WritableCustomFieldRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetUnique() bool {
	if o == nil || IsNil(o.Unique) {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.Unique) {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasUnique() bool {
	if o != nil && !IsNil(o.Unique) {
		return true
	}

	return false
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *WritableCustomFieldRequest) SetUnique(v bool) {
	o.Unique = &v
}

// GetSearchWeight returns the SearchWeight field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetSearchWeight() int32 {
	if o == nil || IsNil(o.SearchWeight) {
		var ret int32
		return ret
	}
	return *o.SearchWeight
}

// GetSearchWeightOk returns a tuple with the SearchWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetSearchWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SearchWeight) {
		return nil, false
	}
	return o.SearchWeight, true
}

// HasSearchWeight returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasSearchWeight() bool {
	if o != nil && !IsNil(o.SearchWeight) {
		return true
	}

	return false
}

// SetSearchWeight gets a reference to the given int32 and assigns it to the SearchWeight field.
func (o *WritableCustomFieldRequest) SetSearchWeight(v int32) {
	o.SearchWeight = &v
}

// GetFilterLogic returns the FilterLogic field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetFilterLogic() PatchedWritableCustomFieldRequestFilterLogic {
	if o == nil || IsNil(o.FilterLogic) {
		var ret PatchedWritableCustomFieldRequestFilterLogic
		return ret
	}
	return *o.FilterLogic
}

// GetFilterLogicOk returns a tuple with the FilterLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetFilterLogicOk() (*PatchedWritableCustomFieldRequestFilterLogic, bool) {
	if o == nil || IsNil(o.FilterLogic) {
		return nil, false
	}
	return o.FilterLogic, true
}

// HasFilterLogic returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasFilterLogic() bool {
	if o != nil && !IsNil(o.FilterLogic) {
		return true
	}

	return false
}

// SetFilterLogic gets a reference to the given PatchedWritableCustomFieldRequestFilterLogic and assigns it to the FilterLogic field.
func (o *WritableCustomFieldRequest) SetFilterLogic(v PatchedWritableCustomFieldRequestFilterLogic) {
	o.FilterLogic = &v
}

// GetUiVisible returns the UiVisible field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetUiVisible() PatchedWritableCustomFieldRequestUiVisible {
	if o == nil || IsNil(o.UiVisible) {
		var ret PatchedWritableCustomFieldRequestUiVisible
		return ret
	}
	return *o.UiVisible
}

// GetUiVisibleOk returns a tuple with the UiVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetUiVisibleOk() (*PatchedWritableCustomFieldRequestUiVisible, bool) {
	if o == nil || IsNil(o.UiVisible) {
		return nil, false
	}
	return o.UiVisible, true
}

// HasUiVisible returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasUiVisible() bool {
	if o != nil && !IsNil(o.UiVisible) {
		return true
	}

	return false
}

// SetUiVisible gets a reference to the given PatchedWritableCustomFieldRequestUiVisible and assigns it to the UiVisible field.
func (o *WritableCustomFieldRequest) SetUiVisible(v PatchedWritableCustomFieldRequestUiVisible) {
	o.UiVisible = &v
}

// GetUiEditable returns the UiEditable field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetUiEditable() PatchedWritableCustomFieldRequestUiEditable {
	if o == nil || IsNil(o.UiEditable) {
		var ret PatchedWritableCustomFieldRequestUiEditable
		return ret
	}
	return *o.UiEditable
}

// GetUiEditableOk returns a tuple with the UiEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetUiEditableOk() (*PatchedWritableCustomFieldRequestUiEditable, bool) {
	if o == nil || IsNil(o.UiEditable) {
		return nil, false
	}
	return o.UiEditable, true
}

// HasUiEditable returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasUiEditable() bool {
	if o != nil && !IsNil(o.UiEditable) {
		return true
	}

	return false
}

// SetUiEditable gets a reference to the given PatchedWritableCustomFieldRequestUiEditable and assigns it to the UiEditable field.
func (o *WritableCustomFieldRequest) SetUiEditable(v PatchedWritableCustomFieldRequestUiEditable) {
	o.UiEditable = &v
}

// GetIsCloneable returns the IsCloneable field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetIsCloneable() bool {
	if o == nil || IsNil(o.IsCloneable) {
		var ret bool
		return ret
	}
	return *o.IsCloneable
}

// GetIsCloneableOk returns a tuple with the IsCloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetIsCloneableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloneable) {
		return nil, false
	}
	return o.IsCloneable, true
}

// HasIsCloneable returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasIsCloneable() bool {
	if o != nil && !IsNil(o.IsCloneable) {
		return true
	}

	return false
}

// SetIsCloneable gets a reference to the given bool and assigns it to the IsCloneable field.
func (o *WritableCustomFieldRequest) SetIsCloneable(v bool) {
	o.IsCloneable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldRequest) GetDefault() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldRequest) GetDefaultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *WritableCustomFieldRequest) SetDefault(v interface{}) {
	o.Default = v
}

// GetRelatedObjectFilter returns the RelatedObjectFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldRequest) GetRelatedObjectFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RelatedObjectFilter
}

// GetRelatedObjectFilterOk returns a tuple with the RelatedObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldRequest) GetRelatedObjectFilterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RelatedObjectFilter) {
		return nil, false
	}
	return &o.RelatedObjectFilter, true
}

// HasRelatedObjectFilter returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasRelatedObjectFilter() bool {
	if o != nil && !IsNil(o.RelatedObjectFilter) {
		return true
	}

	return false
}

// SetRelatedObjectFilter gets a reference to the given interface{} and assigns it to the RelatedObjectFilter field.
func (o *WritableCustomFieldRequest) SetRelatedObjectFilter(v interface{}) {
	o.RelatedObjectFilter = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *WritableCustomFieldRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetValidationMinimum returns the ValidationMinimum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldRequest) GetValidationMinimum() int64 {
	if o == nil || IsNil(o.ValidationMinimum.Get()) {
		var ret int64
		return ret
	}
	return *o.ValidationMinimum.Get()
}

// GetValidationMinimumOk returns a tuple with the ValidationMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldRequest) GetValidationMinimumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMinimum.Get(), o.ValidationMinimum.IsSet()
}

// HasValidationMinimum returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasValidationMinimum() bool {
	if o != nil && o.ValidationMinimum.IsSet() {
		return true
	}

	return false
}

// SetValidationMinimum gets a reference to the given NullableInt64 and assigns it to the ValidationMinimum field.
func (o *WritableCustomFieldRequest) SetValidationMinimum(v int64) {
	o.ValidationMinimum.Set(&v)
}

// SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil
func (o *WritableCustomFieldRequest) SetValidationMinimumNil() {
	o.ValidationMinimum.Set(nil)
}

// UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
func (o *WritableCustomFieldRequest) UnsetValidationMinimum() {
	o.ValidationMinimum.Unset()
}

// GetValidationMaximum returns the ValidationMaximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldRequest) GetValidationMaximum() int64 {
	if o == nil || IsNil(o.ValidationMaximum.Get()) {
		var ret int64
		return ret
	}
	return *o.ValidationMaximum.Get()
}

// GetValidationMaximumOk returns a tuple with the ValidationMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldRequest) GetValidationMaximumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMaximum.Get(), o.ValidationMaximum.IsSet()
}

// HasValidationMaximum returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasValidationMaximum() bool {
	if o != nil && o.ValidationMaximum.IsSet() {
		return true
	}

	return false
}

// SetValidationMaximum gets a reference to the given NullableInt64 and assigns it to the ValidationMaximum field.
func (o *WritableCustomFieldRequest) SetValidationMaximum(v int64) {
	o.ValidationMaximum.Set(&v)
}

// SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil
func (o *WritableCustomFieldRequest) SetValidationMaximumNil() {
	o.ValidationMaximum.Set(nil)
}

// UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
func (o *WritableCustomFieldRequest) UnsetValidationMaximum() {
	o.ValidationMaximum.Unset()
}

// GetValidationRegex returns the ValidationRegex field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetValidationRegex() string {
	if o == nil || IsNil(o.ValidationRegex) {
		var ret string
		return ret
	}
	return *o.ValidationRegex
}

// GetValidationRegexOk returns a tuple with the ValidationRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetValidationRegexOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationRegex) {
		return nil, false
	}
	return o.ValidationRegex, true
}

// HasValidationRegex returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasValidationRegex() bool {
	if o != nil && !IsNil(o.ValidationRegex) {
		return true
	}

	return false
}

// SetValidationRegex gets a reference to the given string and assigns it to the ValidationRegex field.
func (o *WritableCustomFieldRequest) SetValidationRegex(v string) {
	o.ValidationRegex = &v
}

// GetChoiceSet returns the ChoiceSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldRequest) GetChoiceSet() BriefCustomFieldChoiceSetRequest {
	if o == nil || IsNil(o.ChoiceSet.Get()) {
		var ret BriefCustomFieldChoiceSetRequest
		return ret
	}
	return *o.ChoiceSet.Get()
}

// GetChoiceSetOk returns a tuple with the ChoiceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldRequest) GetChoiceSetOk() (*BriefCustomFieldChoiceSetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChoiceSet.Get(), o.ChoiceSet.IsSet()
}

// HasChoiceSet returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasChoiceSet() bool {
	if o != nil && o.ChoiceSet.IsSet() {
		return true
	}

	return false
}

// SetChoiceSet gets a reference to the given NullableBriefCustomFieldChoiceSetRequest and assigns it to the ChoiceSet field.
func (o *WritableCustomFieldRequest) SetChoiceSet(v BriefCustomFieldChoiceSetRequest) {
	o.ChoiceSet.Set(&v)
}

// SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil
func (o *WritableCustomFieldRequest) SetChoiceSetNil() {
	o.ChoiceSet.Set(nil)
}

// UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
func (o *WritableCustomFieldRequest) UnsetChoiceSet() {
	o.ChoiceSet.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableCustomFieldRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableCustomFieldRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableCustomFieldRequest) SetComments(v string) {
	o.Comments = &v
}

func (o WritableCustomFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableCustomFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_types"] = o.ObjectTypes
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.RelatedObjectType.IsSet() {
		toSerialize["related_object_type"] = o.RelatedObjectType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Unique) {
		toSerialize["unique"] = o.Unique
	}
	if !IsNil(o.SearchWeight) {
		toSerialize["search_weight"] = o.SearchWeight
	}
	if !IsNil(o.FilterLogic) {
		toSerialize["filter_logic"] = o.FilterLogic
	}
	if !IsNil(o.UiVisible) {
		toSerialize["ui_visible"] = o.UiVisible
	}
	if !IsNil(o.UiEditable) {
		toSerialize["ui_editable"] = o.UiEditable
	}
	if !IsNil(o.IsCloneable) {
		toSerialize["is_cloneable"] = o.IsCloneable
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.RelatedObjectFilter != nil {
		toSerialize["related_object_filter"] = o.RelatedObjectFilter
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if o.ValidationMinimum.IsSet() {
		toSerialize["validation_minimum"] = o.ValidationMinimum.Get()
	}
	if o.ValidationMaximum.IsSet() {
		toSerialize["validation_maximum"] = o.ValidationMaximum.Get()
	}
	if !IsNil(o.ValidationRegex) {
		toSerialize["validation_regex"] = o.ValidationRegex
	}
	if o.ChoiceSet.IsSet() {
		toSerialize["choice_set"] = o.ChoiceSet.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableCustomFieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_types",
		"name",
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

	varWritableCustomFieldRequest := _WritableCustomFieldRequest{}

	err = json.Unmarshal(data, &varWritableCustomFieldRequest)

	if err != nil {
		return err
	}

	*o = WritableCustomFieldRequest(varWritableCustomFieldRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "type")
		delete(additionalProperties, "related_object_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "required")
		delete(additionalProperties, "unique")
		delete(additionalProperties, "search_weight")
		delete(additionalProperties, "filter_logic")
		delete(additionalProperties, "ui_visible")
		delete(additionalProperties, "ui_editable")
		delete(additionalProperties, "is_cloneable")
		delete(additionalProperties, "default")
		delete(additionalProperties, "related_object_filter")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "validation_minimum")
		delete(additionalProperties, "validation_maximum")
		delete(additionalProperties, "validation_regex")
		delete(additionalProperties, "choice_set")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableCustomFieldRequest struct {
	value *WritableCustomFieldRequest
	isSet bool
}

func (v NullableWritableCustomFieldRequest) Get() *WritableCustomFieldRequest {
	return v.value
}

func (v *NullableWritableCustomFieldRequest) Set(val *WritableCustomFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableCustomFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableCustomFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableCustomFieldRequest(val *WritableCustomFieldRequest) *NullableWritableCustomFieldRequest {
	return &NullableWritableCustomFieldRequest{value: val, isSet: true}
}

func (v NullableWritableCustomFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableCustomFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
