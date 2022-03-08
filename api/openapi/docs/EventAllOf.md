# EventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** | Type of event. Types should be defined in GoCommerce config file. | [optional] 
**Start** | Pointer to **time.Time** | Start of the event. In case \&quot;whole_day\&quot; is true, only the date part is considered. | [optional] 
**End** | Pointer to **time.Time** | End of the event, could be same as start. In case \&quot;whole_day\&quot; is true, only the date part is considered. | [optional] 
**WholeDay** | Pointer to **bool** |  | [optional] 

## Methods

### NewEventAllOf

`func NewEventAllOf() *EventAllOf`

NewEventAllOf instantiates a new EventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAllOfWithDefaults

`func NewEventAllOfWithDefaults() *EventAllOf`

NewEventAllOfWithDefaults instantiates a new EventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EventAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventType

`func (o *EventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetStart

`func (o *EventAllOf) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventAllOf) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventAllOf) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventAllOf) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *EventAllOf) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EventAllOf) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EventAllOf) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EventAllOf) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetWholeDay

`func (o *EventAllOf) GetWholeDay() bool`

GetWholeDay returns the WholeDay field if non-nil, zero value otherwise.

### GetWholeDayOk

`func (o *EventAllOf) GetWholeDayOk() (*bool, bool)`

GetWholeDayOk returns a tuple with the WholeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholeDay

`func (o *EventAllOf) SetWholeDay(v bool)`

SetWholeDay sets WholeDay field to given value.

### HasWholeDay

`func (o *EventAllOf) HasWholeDay() bool`

HasWholeDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


