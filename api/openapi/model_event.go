/*
GoCommerce

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Event struct for Event
type Event struct {
	// Compressed representation of ID
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// Type of event. Types should be defined in GoCommerce config file.
	EventType string `json:"event_type"`
	// Start of the event. In case \"whole_day\" is true, only the date part is considered.
	Start time.Time `json:"start"`
	// Optional end of the event. In case \"whole_day\" is true, only the date part is considered.
	End *time.Time `json:"end,omitempty"`
	WholeDay *bool `json:"whole_day,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(name string, eventType string, start time.Time) *Event {
	this := Event{}
	this.Name = name
	this.EventType = eventType
	this.Start = start
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Event) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Event) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Event) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Event) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Event) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Event) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Event) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Event) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Event) SetDescription(v string) {
	o.Description = &v
}

// GetEventType returns the EventType field value
func (o *Event) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *Event) SetEventType(v string) {
	o.EventType = v
}

// GetStart returns the Start field value
func (o *Event) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *Event) GetStartOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *Event) SetStart(v time.Time) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Event) GetEnd() time.Time {
	if o == nil || o.End == nil {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEndOk() (*time.Time, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Event) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *Event) SetEnd(v time.Time) {
	o.End = &v
}

// GetWholeDay returns the WholeDay field value if set, zero value otherwise.
func (o *Event) GetWholeDay() bool {
	if o == nil || o.WholeDay == nil {
		var ret bool
		return ret
	}
	return *o.WholeDay
}

// GetWholeDayOk returns a tuple with the WholeDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWholeDayOk() (*bool, bool) {
	if o == nil || o.WholeDay == nil {
		return nil, false
	}
	return o.WholeDay, true
}

// HasWholeDay returns a boolean if a field has been set.
func (o *Event) HasWholeDay() bool {
	if o != nil && o.WholeDay != nil {
		return true
	}

	return false
}

// SetWholeDay gets a reference to the given bool and assigns it to the WholeDay field.
func (o *Event) SetWholeDay(v bool) {
	o.WholeDay = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if true {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.WholeDay != nil {
		toSerialize["whole_day"] = o.WholeDay
	}
	return json.Marshal(toSerialize)
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


