# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Compressed representation of ID | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** | Compressed representation of ID | [optional] 
**Order** | **int64** | Should be sorted ascending by this column | 
**ProductIds** | Pointer to **[]string** |  | [optional] [readonly] 
**ImageUrls** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewCategory

`func NewCategory(name string, order int64, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Category) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Category) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Category) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Category) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Category) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *Category) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Category) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Category) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Category) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOrder

`func (o *Category) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Category) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Category) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetProductIds

`func (o *Category) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *Category) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *Category) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *Category) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetImageUrls

`func (o *Category) GetImageUrls() map[string]string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Category) GetImageUrlsOk() (*map[string]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Category) SetImageUrls(v map[string]string)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Category) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


