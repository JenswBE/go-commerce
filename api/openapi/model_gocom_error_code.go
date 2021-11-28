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

// GocomErrorCode - CATEGORY_NAME_EMPTY: Category name is required and cannot be empty - CATEGORY_ORDER_NEGATIVE: Category order should be a positive integer - CATEGORY_PARENT_ID_INVALID: Parent ID of the category is invalid - CONTENT_NAME_EMPTY: Content name is required and cannot be empty - CONTENT_TYPE_INVALID: Content type is empty or has an invalid value - EVENT_END_BEFORE_START: The end date of the event should be equal to or after the start date - IMAGE_ORDER_NEGATIVE: Image order should be a positive integer - INVALID_AUTH_TOKEN: Provided authentication token is invalid - INVALID_ID: Provided short ID or UUID is invalid - MISSING_ADMIN_ROLE: Required role \"admin\" is missing on provided authentication token - PARAMETER_MISSING: A required URL parameter is missing - PRODUCT_CATEGORY_IDS_INVALID: Category ID's of product are invalid - PRODUCT_MANUFACTURER_ID_INVALID: Manufacturer ID of the product is invalid - PRODUCT_NAME_EMPTY: Product name is required and cannot be empty - PRODUCT_PRICE_NEGATIVE: Product price should be a positive integer - SINGLE_IMAGE_IN_FORM: Exactly one image is expected in multipart form, but none or multiple are provided - UNKNOWN_CATEGORY: The category does not exist - UNKNOWN_CONTENT: The content does not exist - UNKNOWN_ERROR: An unknown error occurred - UNKNOWN_EVENT: The event does not exist - UNKNOWN_IMAGE: The image does not exist - UNKNOWN_MANUFACTURER: The manufacturer does not exist - UNKNOWN_PRODUCT: The product does not exist 
type GocomErrorCode string

// List of GocomErrorCode
const (
	GOCOMERRORCODE_CATEGORY_NAME_EMPTY GocomErrorCode = "CATEGORY_NAME_EMPTY"
	GOCOMERRORCODE_CATEGORY_ORDER_NEGATIVE GocomErrorCode = "CATEGORY_ORDER_NEGATIVE"
	GOCOMERRORCODE_CATEGORY_PARENT_ID_INVALID GocomErrorCode = "CATEGORY_PARENT_ID_INVALID"
	GOCOMERRORCODE_CONTENT_NAME_EMPTY GocomErrorCode = "CONTENT_NAME_EMPTY"
	GOCOMERRORCODE_CONTENT_TYPE_INVALID GocomErrorCode = "CONTENT_TYPE_INVALID"
	GOCOMERRORCODE_EVENT_END_BEFORE_START GocomErrorCode = "EVENT_END_BEFORE_START"
	GOCOMERRORCODE_IMAGE_ORDER_NEGATIVE GocomErrorCode = "IMAGE_ORDER_NEGATIVE"
	GOCOMERRORCODE_INVALID_AUTH_TOKEN GocomErrorCode = "INVALID_AUTH_TOKEN"
	GOCOMERRORCODE_INVALID_ID GocomErrorCode = "INVALID_ID"
	GOCOMERRORCODE_MISSING_ADMIN_ROLE GocomErrorCode = "MISSING_ADMIN_ROLE"
	GOCOMERRORCODE_PARAMETER_MISSING GocomErrorCode = "PARAMETER_MISSING"
	GOCOMERRORCODE_PRODUCT_CATEGORY_IDS_INVALID GocomErrorCode = "PRODUCT_CATEGORY_IDS_INVALID"
	GOCOMERRORCODE_PRODUCT_MANUFACTURER_ID_INVALID GocomErrorCode = "PRODUCT_MANUFACTURER_ID_INVALID"
	GOCOMERRORCODE_PRODUCT_NAME_EMPTY GocomErrorCode = "PRODUCT_NAME_EMPTY"
	GOCOMERRORCODE_PRODUCT_PRICE_NEGATIVE GocomErrorCode = "PRODUCT_PRICE_NEGATIVE"
	GOCOMERRORCODE_SINGLE_IMAGE_IN_FORM GocomErrorCode = "SINGLE_IMAGE_IN_FORM"
	GOCOMERRORCODE_UNKNOWN_CATEGORY GocomErrorCode = "UNKNOWN_CATEGORY"
	GOCOMERRORCODE_UNKNOWN_CONTENT GocomErrorCode = "UNKNOWN_CONTENT"
	GOCOMERRORCODE_UNKNOWN_ERROR GocomErrorCode = "UNKNOWN_ERROR"
	GOCOMERRORCODE_UNKNOWN_EVENT GocomErrorCode = "UNKNOWN_EVENT"
	GOCOMERRORCODE_UNKNOWN_IMAGE GocomErrorCode = "UNKNOWN_IMAGE"
	GOCOMERRORCODE_UNKNOWN_MANUFACTURER GocomErrorCode = "UNKNOWN_MANUFACTURER"
	GOCOMERRORCODE_UNKNOWN_PRODUCT GocomErrorCode = "UNKNOWN_PRODUCT"
)

// All allowed values of GocomErrorCode enum
var AllowedGocomErrorCodeEnumValues = []GocomErrorCode{
	"CATEGORY_NAME_EMPTY",
	"CATEGORY_ORDER_NEGATIVE",
	"CATEGORY_PARENT_ID_INVALID",
	"CONTENT_NAME_EMPTY",
	"CONTENT_TYPE_INVALID",
	"EVENT_END_BEFORE_START",
	"IMAGE_ORDER_NEGATIVE",
	"INVALID_AUTH_TOKEN",
	"INVALID_ID",
	"MISSING_ADMIN_ROLE",
	"PARAMETER_MISSING",
	"PRODUCT_CATEGORY_IDS_INVALID",
	"PRODUCT_MANUFACTURER_ID_INVALID",
	"PRODUCT_NAME_EMPTY",
	"PRODUCT_PRICE_NEGATIVE",
	"SINGLE_IMAGE_IN_FORM",
	"UNKNOWN_CATEGORY",
	"UNKNOWN_CONTENT",
	"UNKNOWN_ERROR",
	"UNKNOWN_EVENT",
	"UNKNOWN_IMAGE",
	"UNKNOWN_MANUFACTURER",
	"UNKNOWN_PRODUCT",
}

func (v *GocomErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GocomErrorCode(value)
	for _, existing := range AllowedGocomErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GocomErrorCode", value)
}

// NewGocomErrorCodeFromValue returns a pointer to a valid GocomErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGocomErrorCodeFromValue(v string) (*GocomErrorCode, error) {
	ev := GocomErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GocomErrorCode: valid values are %v", v, AllowedGocomErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GocomErrorCode) IsValid() bool {
	for _, existing := range AllowedGocomErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GocomErrorCode value
func (v GocomErrorCode) Ptr() *GocomErrorCode {
	return &v
}

type NullableGocomErrorCode struct {
	value *GocomErrorCode
	isSet bool
}

func (v NullableGocomErrorCode) Get() *GocomErrorCode {
	return v.value
}

func (v *NullableGocomErrorCode) Set(val *GocomErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableGocomErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableGocomErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGocomErrorCode(val *GocomErrorCode) *NullableGocomErrorCode {
	return &NullableGocomErrorCode{value: val, isSet: true}
}

func (v NullableGocomErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGocomErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
