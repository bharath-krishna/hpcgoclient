# DataTransferInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | **string** |  | 
**Description** | **string** |  | 
**Source** | [**DataTransferEndpoint**](DataTransferEndpoint.md) |  | 
**Target** | [**DataTransferEndpoint**](DataTransferEndpoint.md) |  | 

## Methods

### NewDataTransferInput

`func NewDataTransferInput(transferId string, description string, source DataTransferEndpoint, target DataTransferEndpoint, ) *DataTransferInput`

NewDataTransferInput instantiates a new DataTransferInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferInputWithDefaults

`func NewDataTransferInputWithDefaults() *DataTransferInput`

NewDataTransferInputWithDefaults instantiates a new DataTransferInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *DataTransferInput) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *DataTransferInput) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *DataTransferInput) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetDescription

`func (o *DataTransferInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataTransferInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataTransferInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSource

`func (o *DataTransferInput) GetSource() DataTransferEndpoint`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataTransferInput) GetSourceOk() (*DataTransferEndpoint, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataTransferInput) SetSource(v DataTransferEndpoint)`

SetSource sets Source field to given value.


### GetTarget

`func (o *DataTransferInput) GetTarget() DataTransferEndpoint`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DataTransferInput) GetTargetOk() (*DataTransferEndpoint, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DataTransferInput) SetTarget(v DataTransferEndpoint)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


