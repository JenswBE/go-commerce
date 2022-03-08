# ProductAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewProductAllOf

`func NewProductAllOf(name string, price int64, ) *ProductAllOf`

NewProductAllOf instantiates a new ProductAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAllOfWithDefaults

`func NewProductAllOfWithDefaults() *ProductAllOf`

NewProductAllOfWithDefaults instantiates a new ProductAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescriptionShort

`func (o *ProductAllOf) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *ProductAllOf) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *ProductAllOf) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.

### HasDescriptionShort

`func (o *ProductAllOf) HasDescriptionShort() bool`

HasDescriptionShort returns a boolean if a field has been set.

### GetDescriptionLong

`func (o *ProductAllOf) GetDescriptionLong() string`

GetDescriptionLong returns the DescriptionLong field if non-nil, zero value otherwise.

### GetDescriptionLongOk

`func (o *ProductAllOf) GetDescriptionLongOk() (*string, bool)`

GetDescriptionLongOk returns a tuple with the DescriptionLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLong

`func (o *ProductAllOf) SetDescriptionLong(v string)`

SetDescriptionLong sets DescriptionLong field to given value.

### HasDescriptionLong

`func (o *ProductAllOf) HasDescriptionLong() bool`

HasDescriptionLong returns a boolean if a field has been set.

### GetPrice

`func (o *ProductAllOf) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductAllOf) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductAllOf) SetPrice(v int64)`

SetPrice sets Price field to given value.


### GetCategoryIds

`func (o *ProductAllOf) GetCategoryIds() []string`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *ProductAllOf) GetCategoryIdsOk() (*[]string, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *ProductAllOf) SetCategoryIds(v []string)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *ProductAllOf) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductAllOf) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductAllOf) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductAllOf) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductAllOf) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetStatus

`func (o *ProductAllOf) GetStatus() ProductStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductAllOf) GetStatusOk() (*ProductStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductAllOf) SetStatus(v ProductStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStockCount

`func (o *ProductAllOf) GetStockCount() int64`

GetStockCount returns the StockCount field if non-nil, zero value otherwise.

### GetStockCountOk

`func (o *ProductAllOf) GetStockCountOk() (*int64, bool)`

GetStockCountOk returns a tuple with the StockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockCount

`func (o *ProductAllOf) SetStockCount(v int64)`

SetStockCount sets StockCount field to given value.

### HasStockCount

`func (o *ProductAllOf) HasStockCount() bool`

HasStockCount returns a boolean if a field has been set.

### GetImageUrls

`func (o *ProductAllOf) GetImageUrls() []map[string]string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *ProductAllOf) GetImageUrlsOk() (*[]map[string]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *ProductAllOf) SetImageUrls(v []map[string]string)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *ProductAllOf) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


