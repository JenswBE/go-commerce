# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Ext** | **string** | Extension of the image | [readonly] 
**Urls** | **map[string]string** |  | [readonly] 
**Order** | **int64** | Should be sorted ascending by this column | 

## Methods

### NewImage

`func NewImage(id string, ext string, urls map[string]string, order int64, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.


### GetExt

`func (o *Image) GetExt() string`

GetExt returns the Ext field if non-nil, zero value otherwise.

### GetExtOk

`func (o *Image) GetExtOk() (*string, bool)`

GetExtOk returns a tuple with the Ext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExt

`func (o *Image) SetExt(v string)`

SetExt sets Ext field to given value.


### GetUrls

`func (o *Image) GetUrls() map[string]string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Image) GetUrlsOk() (*map[string]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Image) SetUrls(v map[string]string)`

SetUrls sets Urls field to given value.


### GetOrder

`func (o *Image) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Image) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Image) SetOrder(v int64)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


