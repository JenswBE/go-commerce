# CategoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** | Compressed representation of ID | [optional] 
**Order** | **int64** | Should be sorted ascending by this column | 
**ProductIds** | Pointer to **[]string** |  | [optional] [readonly] 
**ImageUrls** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewCategoryAllOf

`func NewCategoryAllOf(name string, order int64, ) *CategoryAllOf`

NewCategoryAllOf instantiates a new CategoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryAllOfWithDefaults

`func NewCategoryAllOfWithDefaults() *CategoryAllOf`

NewCategoryAllOfWithDefaults instantiates a new CategoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CategoryAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *CategoryAllOf) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CategoryAllOf) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CategoryAllOf) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CategoryAllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOrder

`func (o *CategoryAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CategoryAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CategoryAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetProductIds

`func (o *CategoryAllOf) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *CategoryAllOf) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *CategoryAllOf) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *CategoryAllOf) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetImageUrls

`func (o *CategoryAllOf) GetImageUrls() map[string]string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *CategoryAllOf) GetImageUrlsOk() (*map[string]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *CategoryAllOf) SetImageUrls(v map[string]string)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *CategoryAllOf) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


