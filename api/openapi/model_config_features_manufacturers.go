/*
GoCommerce

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConfigFeaturesManufacturers struct for ConfigFeaturesManufacturers
type ConfigFeaturesManufacturers struct {
	Enabled bool `json:"enabled"`
}

// NewConfigFeaturesManufacturers instantiates a new ConfigFeaturesManufacturers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigFeaturesManufacturers(enabled bool) *ConfigFeaturesManufacturers {
	this := ConfigFeaturesManufacturers{}
	this.Enabled = enabled
	return &this
}

// NewConfigFeaturesManufacturersWithDefaults instantiates a new ConfigFeaturesManufacturers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigFeaturesManufacturersWithDefaults() *ConfigFeaturesManufacturers {
	this := ConfigFeaturesManufacturers{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ConfigFeaturesManufacturers) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConfigFeaturesManufacturers) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConfigFeaturesManufacturers) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ConfigFeaturesManufacturers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableConfigFeaturesManufacturers struct {
	value *ConfigFeaturesManufacturers
	isSet bool
}

func (v NullableConfigFeaturesManufacturers) Get() *ConfigFeaturesManufacturers {
	return v.value
}

func (v *NullableConfigFeaturesManufacturers) Set(val *ConfigFeaturesManufacturers) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigFeaturesManufacturers) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigFeaturesManufacturers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigFeaturesManufacturers(val *ConfigFeaturesManufacturers) *NullableConfigFeaturesManufacturers {
	return &NullableConfigFeaturesManufacturers{value: val, isSet: true}
}

func (v NullableConfigFeaturesManufacturers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigFeaturesManufacturers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


