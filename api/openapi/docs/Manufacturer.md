# Manufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Compressed representation of ID | [optional] [readonly] 
**Name** | **string** |  | 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**ImageUrls** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewManufacturer

`func NewManufacturer(name string, ) *Manufacturer`

NewManufacturer instantiates a new Manufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManufacturerWithDefaults

`func NewManufacturerWithDefaults() *Manufacturer`

NewManufacturerWithDefaults instantiates a new Manufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Manufacturer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Manufacturer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Manufacturer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Manufacturer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Manufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manufacturer) SetName(v string)`

SetName sets Name field to given value.


### GetWebsiteUrl

`func (o *Manufacturer) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Manufacturer) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Manufacturer) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Manufacturer) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetImageUrls

`func (o *Manufacturer) GetImageUrls() map[string]string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Manufacturer) GetImageUrlsOk() (*map[string]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Manufacturer) SetImageUrls(v map[string]string)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Manufacturer) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


