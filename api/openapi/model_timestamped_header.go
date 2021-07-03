/*
 * GoCommerce
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TimestampedHeader struct for TimestampedHeader
type TimestampedHeader struct {
	// Compressed representation of ID
	Id        *string    `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewTimestampedHeader instantiates a new TimestampedHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestampedHeader() *TimestampedHeader {
	this := TimestampedHeader{}
	return &this
}

// NewTimestampedHeaderWithDefaults instantiates a new TimestampedHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampedHeaderWithDefaults() *TimestampedHeader {
	this := TimestampedHeader{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimestampedHeader) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampedHeader) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimestampedHeader) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimestampedHeader) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TimestampedHeader) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampedHeader) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TimestampedHeader) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TimestampedHeader) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TimestampedHeader) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampedHeader) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TimestampedHeader) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TimestampedHeader) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TimestampedHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTimestampedHeader struct {
	value *TimestampedHeader
	isSet bool
}

func (v NullableTimestampedHeader) Get() *TimestampedHeader {
	return v.value
}

func (v *NullableTimestampedHeader) Set(val *TimestampedHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestampedHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestampedHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestampedHeader(val *TimestampedHeader) *NullableTimestampedHeader {
	return &NullableTimestampedHeader{value: val, isSet: true}
}

func (v NullableTimestampedHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestampedHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}