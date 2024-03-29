/*
GoCommerce

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResolvedServiceCategoryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolvedServiceCategoryList{}

// ResolvedServiceCategoryList struct for ResolvedServiceCategoryList
type ResolvedServiceCategoryList struct {
	ServiceCategories []ResolvedServiceCategory `json:"service_categories"`
}

type _ResolvedServiceCategoryList ResolvedServiceCategoryList

// NewResolvedServiceCategoryList instantiates a new ResolvedServiceCategoryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolvedServiceCategoryList(serviceCategories []ResolvedServiceCategory) *ResolvedServiceCategoryList {
	this := ResolvedServiceCategoryList{}
	this.ServiceCategories = serviceCategories
	return &this
}

// NewResolvedServiceCategoryListWithDefaults instantiates a new ResolvedServiceCategoryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolvedServiceCategoryListWithDefaults() *ResolvedServiceCategoryList {
	this := ResolvedServiceCategoryList{}
	return &this
}

// GetServiceCategories returns the ServiceCategories field value
func (o *ResolvedServiceCategoryList) GetServiceCategories() []ResolvedServiceCategory {
	if o == nil {
		var ret []ResolvedServiceCategory
		return ret
	}

	return o.ServiceCategories
}

// GetServiceCategoriesOk returns a tuple with the ServiceCategories field value
// and a boolean to check if the value has been set.
func (o *ResolvedServiceCategoryList) GetServiceCategoriesOk() ([]ResolvedServiceCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceCategories, true
}

// SetServiceCategories sets field value
func (o *ResolvedServiceCategoryList) SetServiceCategories(v []ResolvedServiceCategory) {
	o.ServiceCategories = v
}

func (o ResolvedServiceCategoryList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolvedServiceCategoryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service_categories"] = o.ServiceCategories
	return toSerialize, nil
}

func (o *ResolvedServiceCategoryList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_categories",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResolvedServiceCategoryList := _ResolvedServiceCategoryList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResolvedServiceCategoryList)

	if err != nil {
		return err
	}

	*o = ResolvedServiceCategoryList(varResolvedServiceCategoryList)

	return err
}

type NullableResolvedServiceCategoryList struct {
	value *ResolvedServiceCategoryList
	isSet bool
}

func (v NullableResolvedServiceCategoryList) Get() *ResolvedServiceCategoryList {
	return v.value
}

func (v *NullableResolvedServiceCategoryList) Set(val *ResolvedServiceCategoryList) {
	v.value = val
	v.isSet = true
}

func (v NullableResolvedServiceCategoryList) IsSet() bool {
	return v.isSet
}

func (v *NullableResolvedServiceCategoryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolvedServiceCategoryList(val *ResolvedServiceCategoryList) *NullableResolvedServiceCategoryList {
	return &NullableResolvedServiceCategoryList{value: val, isSet: true}
}

func (v NullableResolvedServiceCategoryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolvedServiceCategoryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


