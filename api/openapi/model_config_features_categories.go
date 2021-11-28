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

// ConfigFeaturesCategories struct for ConfigFeaturesCategories
type ConfigFeaturesCategories struct {
	Enabled bool `json:"enabled"`
}

// NewConfigFeaturesCategories instantiates a new ConfigFeaturesCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigFeaturesCategories(enabled bool) *ConfigFeaturesCategories {
	this := ConfigFeaturesCategories{}
	this.Enabled = enabled
	return &this
}

// NewConfigFeaturesCategoriesWithDefaults instantiates a new ConfigFeaturesCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigFeaturesCategoriesWithDefaults() *ConfigFeaturesCategories {
	this := ConfigFeaturesCategories{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ConfigFeaturesCategories) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConfigFeaturesCategories) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConfigFeaturesCategories) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ConfigFeaturesCategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableConfigFeaturesCategories struct {
	value *ConfigFeaturesCategories
	isSet bool
}

func (v NullableConfigFeaturesCategories) Get() *ConfigFeaturesCategories {
	return v.value
}

func (v *NullableConfigFeaturesCategories) Set(val *ConfigFeaturesCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigFeaturesCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigFeaturesCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigFeaturesCategories(val *ConfigFeaturesCategories) *NullableConfigFeaturesCategories {
	return &NullableConfigFeaturesCategories{value: val, isSet: true}
}

func (v NullableConfigFeaturesCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigFeaturesCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

