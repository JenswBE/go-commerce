# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Compressed representation of ID | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**DescriptionShort** | Pointer to **string** |  | [optional] 
**DescriptionLong** | Pointer to **string** |  | [optional] 
**Price** | **int64** | Price in cents | 
**CategoryIds** | Pointer to **[]string** |  | [optional] 
**ManufacturerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ProductStatus**](ProductStatus.md) |  | [optional] 
**StockCount** | Pointer to **int64** |  | [optional] 
**ImageUrls** | Pointer to **[]map[string]string** |  | [optional] [readonly] 

## Methods

### NewProduct

`func NewProduct(name string, price int64, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Product) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Product) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Product) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Product) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetDescriptionShort

`func (o *Product) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *Product) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *Product) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.

### HasDescriptionShort

`func (o *Product) HasDescriptionShort() bool`

HasDescriptionShort returns a boolean if a field has been set.

### GetDescriptionLong

`func (o *Product) GetDescriptionLong() string`

GetDescriptionLong returns the DescriptionLong field if non-nil, zero value otherwise.

### GetDescriptionLongOk

`func (o *Product) GetDescriptionLongOk() (*string, bool)`

GetDescriptionLongOk returns a tuple with the DescriptionLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLong

`func (o *Product) SetDescriptionLong(v string)`

SetDescriptionLong sets DescriptionLong field to given value.

### HasDescriptionLong

`func (o *Product) HasDescriptionLong() bool`

HasDescriptionLong returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v int64)`

SetPrice sets Price field to given value.


### GetCategoryIds

`func (o *Product) GetCategoryIds() []string`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *Product) GetCategoryIdsOk() (*[]string, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *Product) SetCategoryIds(v []string)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *Product) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### GetManufacturerId

`func (o *Product) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *Product) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *Product) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *Product) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetStatus

`func (o *Product) GetStatus() ProductStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Product) GetStatusOk() (*ProductStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Product) SetStatus(v ProductStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Product) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStockCount

`func (o *Product) GetStockCount() int64`

GetStockCount returns the StockCount field if non-nil, zero value otherwise.

### GetStockCountOk

`func (o *Product) GetStockCountOk() (*int64, bool)`

GetStockCountOk returns a tuple with the StockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockCount

`func (o *Product) SetStockCount(v int64)`

SetStockCount sets StockCount field to given value.

### HasStockCount

`func (o *Product) HasStockCount() bool`

HasStockCount returns a boolean if a field has been set.

### GetImageUrls

`func (o *Product) GetImageUrls() []map[string]string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Product) GetImageUrlsOk() (*[]map[string]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Product) SetImageUrls(v []map[string]string)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Product) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


