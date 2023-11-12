/*
GoCommerce

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the Content type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Content{}

// Content struct for Content
type Content struct {
	Name string `json:"name"`
	ContentType ContentType `json:"content_type"`
	Body string `json:"body"`
}

type _Content Content

// NewContent instantiates a new Content object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContent(name string, contentType ContentType, body string) *Content {
	this := Content{}
	this.Name = name
	this.ContentType = contentType
	this.Body = body
	return &this
}

// NewContentWithDefaults instantiates a new Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentWithDefaults() *Content {
	this := Content{}
	return &this
}

// GetName returns the Name field value
func (o *Content) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Content) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Content) SetName(v string) {
	o.Name = v
}

// GetContentType returns the ContentType field value
func (o *Content) GetContentType() ContentType {
	if o == nil {
		var ret ContentType
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *Content) GetContentTypeOk() (*ContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *Content) SetContentType(v ContentType) {
	o.ContentType = v
}

// GetBody returns the Body field value
func (o *Content) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Content) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *Content) SetBody(v string) {
	o.Body = v
}

func (o Content) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Content) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["content_type"] = o.ContentType
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *Content) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"content_type",
		"body",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContent := _Content{}

	err = json.Unmarshal(bytes, &varContent)

	if err != nil {
		return err
	}

	*o = Content(varContent)

	return err
}

type NullableContent struct {
	value *Content
	isSet bool
}

func (v NullableContent) Get() *Content {
	return v.value
}

func (v *NullableContent) Set(val *Content) {
	v.value = val
	v.isSet = true
}

func (v NullableContent) IsSet() bool {
	return v.isSet
}

func (v *NullableContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent(val *Content) *NullableContent {
	return &NullableContent{value: val, isSet: true}
}

func (v NullableContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


