# ContainerTerminationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | **NullableTime** |  | 
**FinishedAt** | **NullableTime** |  | 
**ExitCode** | **NullableInt32** |  | 
**Signal** | **NullableInt32** |  | 
**Message** | **NullableString** |  | 
**Reason** | **NullableString** |  | 

## Methods

### NewContainerTerminationState

`func NewContainerTerminationState(startedAt NullableTime, finishedAt NullableTime, exitCode NullableInt32, signal NullableInt32, message NullableString, reason NullableString, ) *ContainerTerminationState`

NewContainerTerminationState instantiates a new ContainerTerminationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTerminationStateWithDefaults

`func NewContainerTerminationStateWithDefaults() *ContainerTerminationState`

NewContainerTerminationStateWithDefaults instantiates a new ContainerTerminationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *ContainerTerminationState) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerTerminationState) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerTerminationState) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *ContainerTerminationState) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ContainerTerminationState) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetFinishedAt

`func (o *ContainerTerminationState) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ContainerTerminationState) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ContainerTerminationState) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *ContainerTerminationState) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ContainerTerminationState) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetExitCode

`func (o *ContainerTerminationState) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ContainerTerminationState) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ContainerTerminationState) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### SetExitCodeNil

`func (o *ContainerTerminationState) SetExitCodeNil(b bool)`

 SetExitCodeNil sets the value for ExitCode to be an explicit nil

### UnsetExitCode
`func (o *ContainerTerminationState) UnsetExitCode()`

UnsetExitCode ensures that no value is present for ExitCode, not even an explicit nil
### GetSignal

`func (o *ContainerTerminationState) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *ContainerTerminationState) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *ContainerTerminationState) SetSignal(v int32)`

SetSignal sets Signal field to given value.


### SetSignalNil

`func (o *ContainerTerminationState) SetSignalNil(b bool)`

 SetSignalNil sets the value for Signal to be an explicit nil

### UnsetSignal
`func (o *ContainerTerminationState) UnsetSignal()`

UnsetSignal ensures that no value is present for Signal, not even an explicit nil
### GetMessage

`func (o *ContainerTerminationState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContainerTerminationState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContainerTerminationState) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ContainerTerminationState) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ContainerTerminationState) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetReason

`func (o *ContainerTerminationState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContainerTerminationState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContainerTerminationState) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ContainerTerminationState) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ContainerTerminationState) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


