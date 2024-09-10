# PodStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]Condition**](Condition.md) |  | 
**StartTime** | **NullableTime** |  | 
**Reason** | **NullableString** |  | 
**Message** | **NullableString** |  | 
**Phase** | **string** |  | 

## Methods

### NewPodStatus

`func NewPodStatus(conditions []Condition, startTime NullableTime, reason NullableString, message NullableString, phase string, ) *PodStatus`

NewPodStatus instantiates a new PodStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodStatusWithDefaults

`func NewPodStatusWithDefaults() *PodStatus`

NewPodStatusWithDefaults instantiates a new PodStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *PodStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PodStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PodStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### SetConditionsNil

`func (o *PodStatus) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PodStatus) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetStartTime

`func (o *PodStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PodStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PodStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *PodStatus) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *PodStatus) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetReason

`func (o *PodStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PodStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PodStatus) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *PodStatus) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PodStatus) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMessage

`func (o *PodStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PodStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PodStatus) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *PodStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PodStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPhase

`func (o *PodStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PodStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PodStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


