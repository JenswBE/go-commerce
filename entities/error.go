package entities

import (
	"errors"
	"fmt"

	"github.com/JenswBE/go-commerce/api/openapi"
)

// GoComError allows to bundle a status with the original error.
// This allows to fine-grained response codes at the API level.
type GoComError struct {
	// HTTP status code
	Status int `json:"status"`

	// Original error
	Err error `json:"-"`

	// Error code
	Code string `json:"code"`

	// Human-readable description of the error
	Message string `json:"message"`

	// Optional - On which object to error occurred
	Instance string `json:"instance"`
}

func (e *GoComError) Error() string {
	return fmt.Sprintf("%d - %s", e.Status, e.Err.Error())
}

// NewError returns a new GoComError
func NewError(status int, code openapi.ErrorCode, instance string, err error) error {
	return &GoComError{
		Status:   status,
		Err:      err,
		Code:     string(code),
		Message:  translateCodeToMessage(code),
		Instance: instance,
	}
}

func translateCodeToMessage(code openapi.ErrorCode) string {
	switch code {
	case openapi.ERRORCODE_CATEGORY_NAME_EMPTY:
		return `Category name is required and cannot be empty`
	case openapi.ERRORCODE_CATEGORY_ORDER_NEGATIVE:
		return `Category order should be a positive integer`
	case openapi.ERRORCODE_CATEGORY_PARENT_ID_INVALID:
		return `Parent ID of the category is invalid`
	case openapi.ERRORCODE_IMAGE_ORDER_NEGATIVE:
		return `Image order should be a positive integer`
	case openapi.ERRORCODE_INVALID_AUTH_TOKEN:
		return `Provided authentication token is invalid`
	case openapi.ERRORCODE_INVALID_ID:
		return `Provided short ID or UUID is invalid`
	case openapi.ERRORCODE_MISSING_ADMIN_ROLE:
		return `Required role "admin" is missing on provided authentication token`
	case openapi.ERRORCODE_PARAMETER_MISSING:
		return `A required URL parameter is missing`
	case openapi.ERRORCODE_PRODUCT_CATEGORY_IDS_INVALID:
		return `Category ID's of product are invalid`
	case openapi.ERRORCODE_PRODUCT_MANUFACTURER_ID_INVALID:
		return `Manufacturer ID of the product is invalid`
	case openapi.ERRORCODE_PRODUCT_NAME_EMPTY:
		return `Product name is required and cannot be empty`
	case openapi.ERRORCODE_PRODUCT_PRICE_NEGATIVE:
		return `Product price should be a positive integer`
	case openapi.ERRORCODE_SINGLE_IMAGE_IN_FORM:
		return `Exactly one image is expected in multipart form, but none or multiple are provided`
	case openapi.ERRORCODE_UNKNOWN_CATEGORY:
		return `The category does not exist`
	case openapi.ERRORCODE_UNKNOWN_ERROR:
		return `An unknown error occurred`
	case openapi.ERRORCODE_UNKNOWN_IMAGE:
		return `The image does not exist`
	case openapi.ERRORCODE_UNKNOWN_MANUFACTURER:
		return `The manufacturer does not exist`
	case openapi.ERRORCODE_UNKNOWN_PRODUCT:
		return `The product does not exist`
	}
	return "" // Covered by exhaustive check
}

// ErrInvalidEntity indicates the provided entity is invalid
var ErrInvalidEntity = errors.New("invalid entity")

// ErrNotFound indicates the requested entity is not found
var ErrNotFound = errors.New("not found")

// ErrInvalidID indicates the provided id is malformed
var ErrInvalidID = errors.New("invalid id")
