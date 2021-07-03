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
	"time"
)

// Product struct for Product
type Product struct {
	// Compressed representation of ID
	Id               *string    `json:"id,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	Name             *string    `json:"name,omitempty"`
	DescriptionShort *string    `json:"description_short,omitempty"`
	DescriptionLong  *string    `json:"description_long,omitempty"`
	// Price in cents
	Price          *int32    `json:"price,omitempty"`
	CategoryIds    *[]string `json:"category_ids,omitempty"`
	ManufacturerId *string   `json:"manufacturer_id,omitempty"`
	Status         *string   `json:"status,omitempty"`
	StockCount     *int32    `json:"stock_count,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Product) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Product) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Product) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Product) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Product) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Product) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Product) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Product) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Product) SetName(v string) {
	o.Name = &v
}

// GetDescriptionShort returns the DescriptionShort field value if set, zero value otherwise.
func (o *Product) GetDescriptionShort() string {
	if o == nil || o.DescriptionShort == nil {
		var ret string
		return ret
	}
	return *o.DescriptionShort
}

// GetDescriptionShortOk returns a tuple with the DescriptionShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionShortOk() (*string, bool) {
	if o == nil || o.DescriptionShort == nil {
		return nil, false
	}
	return o.DescriptionShort, true
}

// HasDescriptionShort returns a boolean if a field has been set.
func (o *Product) HasDescriptionShort() bool {
	if o != nil && o.DescriptionShort != nil {
		return true
	}

	return false
}

// SetDescriptionShort gets a reference to the given string and assigns it to the DescriptionShort field.
func (o *Product) SetDescriptionShort(v string) {
	o.DescriptionShort = &v
}

// GetDescriptionLong returns the DescriptionLong field value if set, zero value otherwise.
func (o *Product) GetDescriptionLong() string {
	if o == nil || o.DescriptionLong == nil {
		var ret string
		return ret
	}
	return *o.DescriptionLong
}

// GetDescriptionLongOk returns a tuple with the DescriptionLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionLongOk() (*string, bool) {
	if o == nil || o.DescriptionLong == nil {
		return nil, false
	}
	return o.DescriptionLong, true
}

// HasDescriptionLong returns a boolean if a field has been set.
func (o *Product) HasDescriptionLong() bool {
	if o != nil && o.DescriptionLong != nil {
		return true
	}

	return false
}

// SetDescriptionLong gets a reference to the given string and assigns it to the DescriptionLong field.
func (o *Product) SetDescriptionLong(v string) {
	o.DescriptionLong = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Product) GetPrice() int32 {
	if o == nil || o.Price == nil {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetPriceOk() (*int32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Product) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *Product) SetPrice(v int32) {
	o.Price = &v
}

// GetCategoryIds returns the CategoryIds field value if set, zero value otherwise.
func (o *Product) GetCategoryIds() []string {
	if o == nil || o.CategoryIds == nil {
		var ret []string
		return ret
	}
	return *o.CategoryIds
}

// GetCategoryIdsOk returns a tuple with the CategoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCategoryIdsOk() (*[]string, bool) {
	if o == nil || o.CategoryIds == nil {
		return nil, false
	}
	return o.CategoryIds, true
}

// HasCategoryIds returns a boolean if a field has been set.
func (o *Product) HasCategoryIds() bool {
	if o != nil && o.CategoryIds != nil {
		return true
	}

	return false
}

// SetCategoryIds gets a reference to the given []string and assigns it to the CategoryIds field.
func (o *Product) SetCategoryIds(v []string) {
	o.CategoryIds = &v
}

// GetManufacturerId returns the ManufacturerId field value if set, zero value otherwise.
func (o *Product) GetManufacturerId() string {
	if o == nil || o.ManufacturerId == nil {
		var ret string
		return ret
	}
	return *o.ManufacturerId
}

// GetManufacturerIdOk returns a tuple with the ManufacturerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetManufacturerIdOk() (*string, bool) {
	if o == nil || o.ManufacturerId == nil {
		return nil, false
	}
	return o.ManufacturerId, true
}

// HasManufacturerId returns a boolean if a field has been set.
func (o *Product) HasManufacturerId() bool {
	if o != nil && o.ManufacturerId != nil {
		return true
	}

	return false
}

// SetManufacturerId gets a reference to the given string and assigns it to the ManufacturerId field.
func (o *Product) SetManufacturerId(v string) {
	o.ManufacturerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Product) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Product) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Product) SetStatus(v string) {
	o.Status = &v
}

// GetStockCount returns the StockCount field value if set, zero value otherwise.
func (o *Product) GetStockCount() int32 {
	if o == nil || o.StockCount == nil {
		var ret int32
		return ret
	}
	return *o.StockCount
}

// GetStockCountOk returns a tuple with the StockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStockCountOk() (*int32, bool) {
	if o == nil || o.StockCount == nil {
		return nil, false
	}
	return o.StockCount, true
}

// HasStockCount returns a boolean if a field has been set.
func (o *Product) HasStockCount() bool {
	if o != nil && o.StockCount != nil {
		return true
	}

	return false
}

// SetStockCount gets a reference to the given int32 and assigns it to the StockCount field.
func (o *Product) SetStockCount(v int32) {
	o.StockCount = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DescriptionShort != nil {
		toSerialize["description_short"] = o.DescriptionShort
	}
	if o.DescriptionLong != nil {
		toSerialize["description_long"] = o.DescriptionLong
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.CategoryIds != nil {
		toSerialize["category_ids"] = o.CategoryIds
	}
	if o.ManufacturerId != nil {
		toSerialize["manufacturer_id"] = o.ManufacturerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StockCount != nil {
		toSerialize["stock_count"] = o.StockCount
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}