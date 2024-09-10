# DataTransferStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | **string** |  | 
**Description** | **string** |  | 
**Status** | [**Status**](Status.md) |  | 
**StartDatetime** | **time.Time** |  | 
**EndDatetime** | **NullableTime** |  | 
**Owner** | **string** |  | 
**Source** | [**DataTransferEndpoint**](DataTransferEndpoint.md) |  | 
**Target** | [**DataTransferEndpoint**](DataTransferEndpoint.md) |  | 
**Logs** | [**DataTransferLogs**](DataTransferLogs.md) |  | 

## Methods

### NewDataTransferStatus

`func NewDataTransferStatus(transferId string, description string, status Status, startDatetime time.Time, endDatetime NullableTime, owner string, source DataTransferEndpoint, target DataTransferEndpoint, logs DataTransferLogs, ) *DataTransferStatus`

NewDataTransferStatus instantiates a new DataTransferStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferStatusWithDefaults

`func NewDataTransferStatusWithDefaults() *DataTransferStatus`

NewDataTransferStatusWithDefaults instantiates a new DataTransferStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *DataTransferStatus) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *DataTransferStatus) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *DataTransferStatus) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetDescription

`func (o *DataTransferStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataTransferStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataTransferStatus) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *DataTransferStatus) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataTransferStatus) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataTransferStatus) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetStartDatetime

`func (o *DataTransferStatus) GetStartDatetime() time.Time`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *DataTransferStatus) GetStartDatetimeOk() (*time.Time, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *DataTransferStatus) SetStartDatetime(v time.Time)`

SetStartDatetime sets StartDatetime field to given value.


### GetEndDatetime

`func (o *DataTransferStatus) GetEndDatetime() time.Time`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *DataTransferStatus) GetEndDatetimeOk() (*time.Time, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *DataTransferStatus) SetEndDatetime(v time.Time)`

SetEndDatetime sets EndDatetime field to given value.


### SetEndDatetimeNil

`func (o *DataTransferStatus) SetEndDatetimeNil(b bool)`

 SetEndDatetimeNil sets the value for EndDatetime to be an explicit nil

### UnsetEndDatetime
`func (o *DataTransferStatus) UnsetEndDatetime()`

UnsetEndDatetime ensures that no value is present for EndDatetime, not even an explicit nil
### GetOwner

`func (o *DataTransferStatus) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DataTransferStatus) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DataTransferStatus) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSource

`func (o *DataTransferStatus) GetSource() DataTransferEndpoint`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataTransferStatus) GetSourceOk() (*DataTransferEndpoint, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataTransferStatus) SetSource(v DataTransferEndpoint)`

SetSource sets Source field to given value.


### GetTarget

`func (o *DataTransferStatus) GetTarget() DataTransferEndpoint`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DataTransferStatus) GetTargetOk() (*DataTransferEndpoint, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DataTransferStatus) SetTarget(v DataTransferEndpoint)`

SetTarget sets Target field to given value.


### GetLogs

`func (o *DataTransferStatus) GetLogs() DataTransferLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DataTransferStatus) GetLogsOk() (*DataTransferLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DataTransferStatus) SetLogs(v DataTransferLogs)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


