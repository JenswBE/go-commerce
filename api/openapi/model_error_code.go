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

// ErrorCode - CATEGORY_NAME_EMPTY: Category name is required and cannot be empty - CATEGORY_ORDER_NEGATIVE: Category order should be a positive integer - CATEGORY_PARENT_ID_INVALID: Parent ID of the category is invalid - IMAGE_ORDER_NEGATIVE: Image order should be a positive integer - INVALID_AUTH_TOKEN: Provided authentication token is invalid - INVALID_ID: Provided short ID or UUID is invalid - MISSING_ADMIN_ROLE: Required role \"admin\" is missing on provided authentication token - PRODUCT_CATEGORY_IDS_INVALID: Category ID's of product are invalid - PRODUCT_MANUFACTURER_ID_INVALID: Manufacturer ID of the product is invalid - PRODUCT_NAME_EMPTY: Product name is required and cannot be empty - PRODUCT_PRICE_NEGATIVE: Product price should be a positive integer - UNKNOWN_CATEGORY: The category does not exist - UNKNOWN_ERROR: An unknown error occurred - UNKNOWN_IMAGE: The image does not exist - UNKNOWN_MANUFACTURER: The manufacturer does not exist - UNKNOWN_PRODUCT: The product does not exist 
type ErrorCode string

// List of ErrorCode
const (
	ERRORCODE_CATEGORY_NAME_EMPTY ErrorCode = "CATEGORY_NAME_EMPTY"
	ERRORCODE_CATEGORY_ORDER_NEGATIVE ErrorCode = "CATEGORY_ORDER_NEGATIVE"
	ERRORCODE_CATEGORY_PARENT_ID_INVALID ErrorCode = "CATEGORY_PARENT_ID_INVALID"
	ERRORCODE_IMAGE_ORDER_NEGATIVE ErrorCode = "IMAGE_ORDER_NEGATIVE"
	ERRORCODE_INVALID_AUTH_TOKEN ErrorCode = "INVALID_AUTH_TOKEN"
	ERRORCODE_INVALID_ID ErrorCode = "INVALID_ID"
	ERRORCODE_MISSING_ADMIN_ROLE ErrorCode = "MISSING_ADMIN_ROLE"
	ERRORCODE_PRODUCT_CATEGORY_IDS_INVALID ErrorCode = "PRODUCT_CATEGORY_IDS_INVALID"
	ERRORCODE_PRODUCT_MANUFACTURER_ID_INVALID ErrorCode = "PRODUCT_MANUFACTURER_ID_INVALID"
	ERRORCODE_PRODUCT_NAME_EMPTY ErrorCode = "PRODUCT_NAME_EMPTY"
	ERRORCODE_PRODUCT_PRICE_NEGATIVE ErrorCode = "PRODUCT_PRICE_NEGATIVE"
	ERRORCODE_UNKNOWN_CATEGORY ErrorCode = "UNKNOWN_CATEGORY"
	ERRORCODE_UNKNOWN_ERROR ErrorCode = "UNKNOWN_ERROR"
	ERRORCODE_UNKNOWN_IMAGE ErrorCode = "UNKNOWN_IMAGE"
	ERRORCODE_UNKNOWN_MANUFACTURER ErrorCode = "UNKNOWN_MANUFACTURER"
	ERRORCODE_UNKNOWN_PRODUCT ErrorCode = "UNKNOWN_PRODUCT"
)

var allowedErrorCodeEnumValues = []ErrorCode{
	"CATEGORY_NAME_EMPTY",
	"CATEGORY_ORDER_NEGATIVE",
	"CATEGORY_PARENT_ID_INVALID",
	"IMAGE_ORDER_NEGATIVE",
	"INVALID_AUTH_TOKEN",
	"INVALID_ID",
	"MISSING_ADMIN_ROLE",
	"PRODUCT_CATEGORY_IDS_INVALID",
	"PRODUCT_MANUFACTURER_ID_INVALID",
	"PRODUCT_NAME_EMPTY",
	"PRODUCT_PRICE_NEGATIVE",
	"UNKNOWN_CATEGORY",
	"UNKNOWN_ERROR",
	"UNKNOWN_IMAGE",
	"UNKNOWN_MANUFACTURER",
	"UNKNOWN_PRODUCT",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range allowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, allowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range allowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

