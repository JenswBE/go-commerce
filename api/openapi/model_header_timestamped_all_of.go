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

// HeaderTimestampedAllOf struct for HeaderTimestampedAllOf
type HeaderTimestampedAllOf struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewHeaderTimestampedAllOf instantiates a new HeaderTimestampedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderTimestampedAllOf() *HeaderTimestampedAllOf {
	this := HeaderTimestampedAllOf{}
	return &this
}

// NewHeaderTimestampedAllOfWithDefaults instantiates a new HeaderTimestampedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderTimestampedAllOfWithDefaults() *HeaderTimestampedAllOf {
	this := HeaderTimestampedAllOf{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HeaderTimestampedAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderTimestampedAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HeaderTimestampedAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *HeaderTimestampedAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *HeaderTimestampedAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderTimestampedAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *HeaderTimestampedAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *HeaderTimestampedAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o HeaderTimestampedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableHeaderTimestampedAllOf struct {
	value *HeaderTimestampedAllOf
	isSet bool
}

func (v NullableHeaderTimestampedAllOf) Get() *HeaderTimestampedAllOf {
	return v.value
}

func (v *NullableHeaderTimestampedAllOf) Set(val *HeaderTimestampedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderTimestampedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderTimestampedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderTimestampedAllOf(val *HeaderTimestampedAllOf) *NullableHeaderTimestampedAllOf {
	return &NullableHeaderTimestampedAllOf{value: val, isSet: true}
}

func (v NullableHeaderTimestampedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderTimestampedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


