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

// ResolvedProductAllOf struct for ResolvedProductAllOf
type ResolvedProductAllOf struct {
	Manufacturer *Manufacturer `json:"manufacturer,omitempty"`
	Categories []Category `json:"categories,omitempty"`
}

// NewResolvedProductAllOf instantiates a new ResolvedProductAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolvedProductAllOf() *ResolvedProductAllOf {
	this := ResolvedProductAllOf{}
	return &this
}

// NewResolvedProductAllOfWithDefaults instantiates a new ResolvedProductAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolvedProductAllOfWithDefaults() *ResolvedProductAllOf {
	this := ResolvedProductAllOf{}
	return &this
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *ResolvedProductAllOf) GetManufacturer() Manufacturer {
	if o == nil || o.Manufacturer == nil {
		var ret Manufacturer
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolvedProductAllOf) GetManufacturerOk() (*Manufacturer, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *ResolvedProductAllOf) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given Manufacturer and assigns it to the Manufacturer field.
func (o *ResolvedProductAllOf) SetManufacturer(v Manufacturer) {
	o.Manufacturer = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ResolvedProductAllOf) GetCategories() []Category {
	if o == nil || o.Categories == nil {
		var ret []Category
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolvedProductAllOf) GetCategoriesOk() ([]Category, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ResolvedProductAllOf) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *ResolvedProductAllOf) SetCategories(v []Category) {
	o.Categories = v
}

func (o ResolvedProductAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

type NullableResolvedProductAllOf struct {
	value *ResolvedProductAllOf
	isSet bool
}

func (v NullableResolvedProductAllOf) Get() *ResolvedProductAllOf {
	return v.value
}

func (v *NullableResolvedProductAllOf) Set(val *ResolvedProductAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResolvedProductAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResolvedProductAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolvedProductAllOf(val *ResolvedProductAllOf) *NullableResolvedProductAllOf {
	return &NullableResolvedProductAllOf{value: val, isSet: true}
}

func (v NullableResolvedProductAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolvedProductAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


