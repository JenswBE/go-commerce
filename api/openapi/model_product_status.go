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
	"fmt"
)

// ProductStatus the model 'ProductStatus'
type ProductStatus string

// List of ProductStatus
const (
	PRODUCTSTATUS_AVAILABLE ProductStatus = "AVAILABLE"
	PRODUCTSTATUS_ARCHIVED ProductStatus = "ARCHIVED"
)

var allowedProductStatusEnumValues = []ProductStatus{
	"AVAILABLE",
	"ARCHIVED",
}

func (v *ProductStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductStatus(value)
	for _, existing := range allowedProductStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductStatus", value)
}

// NewProductStatusFromValue returns a pointer to a valid ProductStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductStatusFromValue(v string) (*ProductStatus, error) {
	ev := ProductStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductStatus: valid values are %v", v, allowedProductStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductStatus) IsValid() bool {
	for _, existing := range allowedProductStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductStatus value
func (v ProductStatus) Ptr() *ProductStatus {
	return &v
}

type NullableProductStatus struct {
	value *ProductStatus
	isSet bool
}

func (v NullableProductStatus) Get() *ProductStatus {
	return v.value
}

func (v *NullableProductStatus) Set(val *ProductStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProductStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProductStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductStatus(val *ProductStatus) *NullableProductStatus {
	return &NullableProductStatus{value: val, isSet: true}
}

func (v NullableProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
