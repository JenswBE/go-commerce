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

// ProductList struct for ProductList
type ProductList struct {
	Products []Product `json:"products"`
}

// NewProductList instantiates a new ProductList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductList(products []Product) *ProductList {
	this := ProductList{}
	this.Products = products
	return &this
}

// NewProductListWithDefaults instantiates a new ProductList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListWithDefaults() *ProductList {
	this := ProductList{}
	return &this
}

// GetProducts returns the Products field value
func (o *ProductList) GetProducts() []Product {
	if o == nil {
		var ret []Product
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *ProductList) GetProductsOk() ([]Product, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *ProductList) SetProducts(v []Product) {
	o.Products = v
}

func (o ProductList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["products"] = o.Products
	}
	return json.Marshal(toSerialize)
}

type NullableProductList struct {
	value *ProductList
	isSet bool
}

func (v NullableProductList) Get() *ProductList {
	return v.value
}

func (v *NullableProductList) Set(val *ProductList) {
	v.value = val
	v.isSet = true
}

func (v NullableProductList) IsSet() bool {
	return v.isSet
}

func (v *NullableProductList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductList(val *ProductList) *NullableProductList {
	return &NullableProductList{value: val, isSet: true}
}

func (v NullableProductList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


