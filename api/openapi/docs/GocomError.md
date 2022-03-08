# GocomError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | HTTP status code | [optional] 
**Code** | [**GocomErrorCode**](GocomErrorCode.md) |  | 
**Message** | **string** | Human-readable description of the error | 
**Instance** | Pointer to **string** | Object to which this error is related | [optional] 

## Methods

### NewGocomError

`func NewGocomError(code GocomErrorCode, message string, ) *GocomError`

NewGocomError instantiates a new GocomError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGocomErrorWithDefaults

`func NewGocomErrorWithDefaults() *GocomError`

NewGocomErrorWithDefaults instantiates a new GocomError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GocomError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GocomError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GocomError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GocomError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *GocomError) GetCode() GocomErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GocomError) GetCodeOk() (*GocomErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GocomError) SetCode(v GocomErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *GocomError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GocomError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GocomError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInstance

`func (o *GocomError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GocomError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GocomError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GocomError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


