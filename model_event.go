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

// Event * `object_created` - Object created * `object_updated` - Object updated * `object_deleted` - Object deleted * `job_started` - Job started * `job_completed` - Job completed * `job_failed` - Job failed * `job_errored` - Job errored
type Event string

// List of Event
const (
	EVENT_OBJECT_CREATED Event = "object_created"
	EVENT_OBJECT_UPDATED Event = "object_updated"
	EVENT_OBJECT_DELETED Event = "object_deleted"
	EVENT_JOB_STARTED    Event = "job_started"
	EVENT_JOB_COMPLETED  Event = "job_completed"
	EVENT_JOB_FAILED     Event = "job_failed"
	EVENT_JOB_ERRORED    Event = "job_errored"
)

// All allowed values of Event enum
var AllowedEventEnumValues = []Event{
	"object_created",
	"object_updated",
	"object_deleted",
	"job_started",
	"job_completed",
	"job_failed",
	"job_errored",
}

func (v *Event) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Event(value)
	for _, existing := range AllowedEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Event", value)
}

// NewEventFromValue returns a pointer to a valid Event
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventFromValue(v string) (*Event, error) {
	ev := Event(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Event: valid values are %v", v, AllowedEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Event) IsValid() bool {
	for _, existing := range AllowedEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Event value
func (v Event) Ptr() *Event {
	return &v
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
