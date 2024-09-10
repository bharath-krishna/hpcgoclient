# Pod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CreationTimestamp** | **time.Time** |  | 
**EndTimestamp** | **NullableTime** |  | 
**Containers** | [**[]Container**](Container.md) |  | 
**Status** | [**PodStatus**](PodStatus.md) |  | 
**GpuCount** | **int32** |  | 
**GpuUtilization** | **float32** |  | 

## Methods

### NewPod

`func NewPod(name string, creationTimestamp time.Time, endTimestamp NullableTime, containers []Container, status PodStatus, gpuCount int32, gpuUtilization float32, ) *Pod`

NewPod instantiates a new Pod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodWithDefaults

`func NewPodWithDefaults() *Pod`

NewPodWithDefaults instantiates a new Pod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Pod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pod) SetName(v string)`

SetName sets Name field to given value.


### GetCreationTimestamp

`func (o *Pod) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Pod) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Pod) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetEndTimestamp

`func (o *Pod) GetEndTimestamp() time.Time`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *Pod) GetEndTimestampOk() (*time.Time, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *Pod) SetEndTimestamp(v time.Time)`

SetEndTimestamp sets EndTimestamp field to given value.


### SetEndTimestampNil

`func (o *Pod) SetEndTimestampNil(b bool)`

 SetEndTimestampNil sets the value for EndTimestamp to be an explicit nil

### UnsetEndTimestamp
`func (o *Pod) UnsetEndTimestamp()`

UnsetEndTimestamp ensures that no value is present for EndTimestamp, not even an explicit nil
### GetContainers

`func (o *Pod) GetContainers() []Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Pod) GetContainersOk() (*[]Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Pod) SetContainers(v []Container)`

SetContainers sets Containers field to given value.


### GetStatus

`func (o *Pod) GetStatus() PodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pod) GetStatusOk() (*PodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pod) SetStatus(v PodStatus)`

SetStatus sets Status field to given value.


### GetGpuCount

`func (o *Pod) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *Pod) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *Pod) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetGpuUtilization

`func (o *Pod) GetGpuUtilization() float32`

GetGpuUtilization returns the GpuUtilization field if non-nil, zero value otherwise.

### GetGpuUtilizationOk

`func (o *Pod) GetGpuUtilizationOk() (*float32, bool)`

GetGpuUtilizationOk returns a tuple with the GpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtilization

`func (o *Pod) SetGpuUtilization(v float32)`

SetGpuUtilization sets GpuUtilization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


