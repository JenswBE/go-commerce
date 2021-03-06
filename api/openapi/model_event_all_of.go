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

// EventAllOf struct for EventAllOf
type EventAllOf struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Type of event. Types should be defined in GoCommerce config file.
	EventType *string `json:"event_type,omitempty"`
	// Start of the event. In case \"whole_day\" is true, only the date part is considered.
	Start *time.Time `json:"start,omitempty"`
	// End of the event, could be same as start. In case \"whole_day\" is true, only the date part is considered.
	End *time.Time `json:"end,omitempty"`
	WholeDay *bool `json:"whole_day,omitempty"`
}

// NewEventAllOf instantiates a new EventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventAllOf() *EventAllOf {
	this := EventAllOf{}
	return &this
}

// NewEventAllOfWithDefaults instantiates a new EventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventAllOfWithDefaults() *EventAllOf {
	this := EventAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventAllOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventAllOf) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventAllOf) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *EventAllOf) SetEventType(v string) {
	o.EventType = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *EventAllOf) GetStart() time.Time {
	if o == nil || o.Start == nil {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetStartOk() (*time.Time, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *EventAllOf) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *EventAllOf) SetStart(v time.Time) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *EventAllOf) GetEnd() time.Time {
	if o == nil || o.End == nil {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetEndOk() (*time.Time, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *EventAllOf) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *EventAllOf) SetEnd(v time.Time) {
	o.End = &v
}

// GetWholeDay returns the WholeDay field value if set, zero value otherwise.
func (o *EventAllOf) GetWholeDay() bool {
	if o == nil || o.WholeDay == nil {
		var ret bool
		return ret
	}
	return *o.WholeDay
}

// GetWholeDayOk returns a tuple with the WholeDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetWholeDayOk() (*bool, bool) {
	if o == nil || o.WholeDay == nil {
		return nil, false
	}
	return o.WholeDay, true
}

// HasWholeDay returns a boolean if a field has been set.
func (o *EventAllOf) HasWholeDay() bool {
	if o != nil && o.WholeDay != nil {
		return true
	}

	return false
}

// SetWholeDay gets a reference to the given bool and assigns it to the WholeDay field.
func (o *EventAllOf) SetWholeDay(v bool) {
	o.WholeDay = &v
}

func (o EventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
	if o.Start != nil {
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

type NullableEventAllOf struct {
	value *EventAllOf
	isSet bool
}

func (v NullableEventAllOf) Get() *EventAllOf {
	return v.value
}

func (v *NullableEventAllOf) Set(val *EventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAllOf(val *EventAllOf) *NullableEventAllOf {
	return &NullableEventAllOf{value: val, isSet: true}
}

func (v NullableEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


