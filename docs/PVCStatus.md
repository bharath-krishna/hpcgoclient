# PVCStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** |  | 
**AllocatedResourceStatuses** | **interface{}** |  | 
**AllocatedResources** | **interface{}** |  | 
**Capacity** | [**Capacity**](Capacity.md) |  | 
**Conditions** | **interface{}** |  | 
**Phase** | **string** |  | 

## Methods

### NewPVCStatus

`func NewPVCStatus(accessModes []string, allocatedResourceStatuses interface{}, allocatedResources interface{}, capacity Capacity, conditions interface{}, phase string, ) *PVCStatus`

NewPVCStatus instantiates a new PVCStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPVCStatusWithDefaults

`func NewPVCStatusWithDefaults() *PVCStatus`

NewPVCStatusWithDefaults instantiates a new PVCStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *PVCStatus) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *PVCStatus) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *PVCStatus) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.


### GetAllocatedResourceStatuses

`func (o *PVCStatus) GetAllocatedResourceStatuses() interface{}`

GetAllocatedResourceStatuses returns the AllocatedResourceStatuses field if non-nil, zero value otherwise.

### GetAllocatedResourceStatusesOk

`func (o *PVCStatus) GetAllocatedResourceStatusesOk() (*interface{}, bool)`

GetAllocatedResourceStatusesOk returns a tuple with the AllocatedResourceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResourceStatuses

`func (o *PVCStatus) SetAllocatedResourceStatuses(v interface{})`

SetAllocatedResourceStatuses sets AllocatedResourceStatuses field to given value.


### SetAllocatedResourceStatusesNil

`func (o *PVCStatus) SetAllocatedResourceStatusesNil(b bool)`

 SetAllocatedResourceStatusesNil sets the value for AllocatedResourceStatuses to be an explicit nil

### UnsetAllocatedResourceStatuses
`func (o *PVCStatus) UnsetAllocatedResourceStatuses()`

UnsetAllocatedResourceStatuses ensures that no value is present for AllocatedResourceStatuses, not even an explicit nil
### GetAllocatedResources

`func (o *PVCStatus) GetAllocatedResources() interface{}`

GetAllocatedResources returns the AllocatedResources field if non-nil, zero value otherwise.

### GetAllocatedResourcesOk

`func (o *PVCStatus) GetAllocatedResourcesOk() (*interface{}, bool)`

GetAllocatedResourcesOk returns a tuple with the AllocatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResources

`func (o *PVCStatus) SetAllocatedResources(v interface{})`

SetAllocatedResources sets AllocatedResources field to given value.


### SetAllocatedResourcesNil

`func (o *PVCStatus) SetAllocatedResourcesNil(b bool)`

 SetAllocatedResourcesNil sets the value for AllocatedResources to be an explicit nil

### UnsetAllocatedResources
`func (o *PVCStatus) UnsetAllocatedResources()`

UnsetAllocatedResources ensures that no value is present for AllocatedResources, not even an explicit nil
### GetCapacity

`func (o *PVCStatus) GetCapacity() Capacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *PVCStatus) GetCapacityOk() (*Capacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *PVCStatus) SetCapacity(v Capacity)`

SetCapacity sets Capacity field to given value.


### GetConditions

`func (o *PVCStatus) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PVCStatus) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PVCStatus) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.


### SetConditionsNil

`func (o *PVCStatus) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PVCStatus) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetPhase

`func (o *PVCStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PVCStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PVCStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


