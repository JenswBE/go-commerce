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

// ManufacturerList struct for ManufacturerList
type ManufacturerList struct {
	Manufacturers []Manufacturer `json:"manufacturers"`
}

// NewManufacturerList instantiates a new ManufacturerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManufacturerList(manufacturers []Manufacturer) *ManufacturerList {
	this := ManufacturerList{}
	this.Manufacturers = manufacturers
	return &this
}

// NewManufacturerListWithDefaults instantiates a new ManufacturerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManufacturerListWithDefaults() *ManufacturerList {
	this := ManufacturerList{}
	return &this
}

// GetManufacturers returns the Manufacturers field value
func (o *ManufacturerList) GetManufacturers() []Manufacturer {
	if o == nil {
		var ret []Manufacturer
		return ret
	}

	return o.Manufacturers
}

// GetManufacturersOk returns a tuple with the Manufacturers field value
// and a boolean to check if the value has been set.
func (o *ManufacturerList) GetManufacturersOk() ([]Manufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturers, true
}

// SetManufacturers sets field value
func (o *ManufacturerList) SetManufacturers(v []Manufacturer) {
	o.Manufacturers = v
}

func (o ManufacturerList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["manufacturers"] = o.Manufacturers
	}
	return json.Marshal(toSerialize)
}

type NullableManufacturerList struct {
	value *ManufacturerList
	isSet bool
}

func (v NullableManufacturerList) Get() *ManufacturerList {
	return v.value
}

func (v *NullableManufacturerList) Set(val *ManufacturerList) {
	v.value = val
	v.isSet = true
}

func (v NullableManufacturerList) IsSet() bool {
	return v.isSet
}

func (v *NullableManufacturerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManufacturerList(val *ManufacturerList) *NullableManufacturerList {
	return &NullableManufacturerList{value: val, isSet: true}
}

func (v NullableManufacturerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManufacturerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


