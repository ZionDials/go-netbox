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

// checks if the NestedVMInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVMInterface{}

// NestedVMInterface Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVMInterface struct {
	Id                   int32                `json:"id"`
	Url                  string               `json:"url"`
	DisplayUrl           string               `json:"display_url"`
	Display              string               `json:"display"`
	VirtualMachine       NestedVirtualMachine `json:"virtual_machine"`
	Name                 string               `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedVMInterface NestedVMInterface

// NewNestedVMInterface instantiates a new NestedVMInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVMInterface(id int32, url string, displayUrl string, display string, virtualMachine NestedVirtualMachine, name string) *NestedVMInterface {
	this := NestedVMInterface{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.VirtualMachine = virtualMachine
	this.Name = name
	return &this
}

// NewNestedVMInterfaceWithDefaults instantiates a new NestedVMInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVMInterfaceWithDefaults() *NestedVMInterface {
	this := NestedVMInterface{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVMInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVMInterface) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedVMInterface) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterface) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVMInterface) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *NestedVMInterface) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterface) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *NestedVMInterface) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *NestedVMInterface) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterface) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedVMInterface) SetDisplay(v string) {
	o.Display = v
}

// GetVirtualMachine returns the VirtualMachine field value
func (o *NestedVMInterface) GetVirtualMachine() NestedVirtualMachine {
	if o == nil {
		var ret NestedVirtualMachine
		return ret
	}

	return o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterface) GetVirtualMachineOk() (*NestedVirtualMachine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualMachine, true
}

// SetVirtualMachine sets field value
func (o *NestedVMInterface) SetVirtualMachine(v NestedVirtualMachine) {
	o.VirtualMachine = v
}

// GetName returns the Name field value
func (o *NestedVMInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVMInterface) SetName(v string) {
	o.Name = v
}

func (o NestedVMInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVMInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["virtual_machine"] = o.VirtualMachine
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVMInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"virtual_machine",
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

	varNestedVMInterface := _NestedVMInterface{}

	err = json.Unmarshal(data, &varNestedVMInterface)

	if err != nil {
		return err
	}

	*o = NestedVMInterface(varNestedVMInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVMInterface struct {
	value *NestedVMInterface
	isSet bool
}

func (v NullableNestedVMInterface) Get() *NestedVMInterface {
	return v.value
}

func (v *NullableNestedVMInterface) Set(val *NestedVMInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVMInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVMInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVMInterface(val *NestedVMInterface) *NullableNestedVMInterface {
	return &NullableNestedVMInterface{value: val, isSet: true}
}

func (v NullableNestedVMInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVMInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
