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

// ContentList struct for ContentList
type ContentList struct {
	Content []Content `json:"content"`
}

// NewContentList instantiates a new ContentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentList(content []Content) *ContentList {
	this := ContentList{}
	this.Content = content
	return &this
}

// NewContentListWithDefaults instantiates a new ContentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentListWithDefaults() *ContentList {
	this := ContentList{}
	return &this
}

// GetContent returns the Content field value
func (o *ContentList) GetContent() []Content {
	if o == nil {
		var ret []Content
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ContentList) GetContentOk() ([]Content, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content, true
}

// SetContent sets field value
func (o *ContentList) SetContent(v []Content) {
	o.Content = v
}

func (o ContentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableContentList struct {
	value *ContentList
	isSet bool
}

func (v NullableContentList) Get() *ContentList {
	return v.value
}

func (v *NullableContentList) Set(val *ContentList) {
	v.value = val
	v.isSet = true
}

func (v NullableContentList) IsSet() bool {
	return v.isSet
}

func (v *NullableContentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentList(val *ContentList) *NullableContentList {
	return &NullableContentList{value: val, isSet: true}
}

func (v NullableContentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


