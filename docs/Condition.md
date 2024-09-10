# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | **NullableTime** |  | 
**Message** | **NullableString** |  | 
**Reason** | **NullableString** |  | 
**Status** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewCondition

`func NewCondition(lastTransitionTime NullableTime, message NullableString, reason NullableString, status string, type_ string, ) *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *Condition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *Condition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *Condition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.


### SetLastTransitionTimeNil

`func (o *Condition) SetLastTransitionTimeNil(b bool)`

 SetLastTransitionTimeNil sets the value for LastTransitionTime to be an explicit nil

### UnsetLastTransitionTime
`func (o *Condition) UnsetLastTransitionTime()`

UnsetLastTransitionTime ensures that no value is present for LastTransitionTime, not even an explicit nil
### GetMessage

`func (o *Condition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Condition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Condition) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *Condition) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *Condition) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetReason

`func (o *Condition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Condition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Condition) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *Condition) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Condition) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetStatus

`func (o *Condition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Condition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Condition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Condition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


