# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Compressed representation of ID | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EventType** | **string** | Type of event. Types should be defined in GoCommerce config file. | 
**Start** | **time.Time** | Start of the event. In case \&quot;whole_day\&quot; is true, only the date part is considered. | 
**End** | **time.Time** | End of the event, could be same as start. In case \&quot;whole_day\&quot; is true, only the date part is considered. | 
**WholeDay** | Pointer to **bool** |  | [optional] 

## Methods

### NewEvent

`func NewEvent(name string, eventType string, start time.Time, end time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Event) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Event) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Event) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Event) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Event) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Event) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Event) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetStart

`func (o *Event) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Event) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Event) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Event) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Event) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Event) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetWholeDay

`func (o *Event) GetWholeDay() bool`

GetWholeDay returns the WholeDay field if non-nil, zero value otherwise.

### GetWholeDayOk

`func (o *Event) GetWholeDayOk() (*bool, bool)`

GetWholeDayOk returns a tuple with the WholeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholeDay

`func (o *Event) SetWholeDay(v bool)`

SetWholeDay sets WholeDay field to given value.

### HasWholeDay

`func (o *Event) HasWholeDay() bool`

HasWholeDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


