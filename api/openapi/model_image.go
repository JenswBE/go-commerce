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

// Image struct for Image
type Image struct {
	Id string `json:"id"`
	// Extension of the image
	Ext string `json:"ext"`
	Urls map[string]string `json:"urls"`
	// Should be sorted ascending by this column
	Order int64 `json:"order"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(id string, ext string, urls map[string]string, order int64) *Image {
	this := Image{}
	this.Id = id
	this.Ext = ext
	this.Urls = urls
	this.Order = order
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetId returns the Id field value
func (o *Image) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Image) SetId(v string) {
	o.Id = v
}

// GetExt returns the Ext field value
func (o *Image) GetExt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ext
}

// GetExtOk returns a tuple with the Ext field value
// and a boolean to check if the value has been set.
func (o *Image) GetExtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ext, true
}

// SetExt sets field value
func (o *Image) SetExt(v string) {
	o.Ext = v
}

// GetUrls returns the Urls field value
func (o *Image) GetUrls() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Image) GetUrlsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Image) SetUrls(v map[string]string) {
	o.Urls = v
}

// GetOrder returns the Order field value
func (o *Image) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *Image) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *Image) SetOrder(v int64) {
	o.Order = v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ext"] = o.Ext
	}
	if true {
		toSerialize["urls"] = o.Urls
	}
	if true {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


