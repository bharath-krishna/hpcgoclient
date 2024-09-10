# JobItemFlat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**K8sName** | **string** |  | 
**Namespace** | **string** |  | 
**Kind** | [**Kind**](Kind.md) |  | 
**GpuCount** | **int32** |  | 
**CpuCount** | **int32** |  | 
**MemoryCount** | **int32** |  | 
**AvgGpuUtilization** | **float32** |  | 
**CreationTimestamp** | **time.Time** |  | 
**JobOwner** | **string** |  | 
**EndTimestamp** | **NullableTime** |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewJobItemFlat

`func NewJobItemFlat(cluster Cluster, id string, name string, k8sName string, namespace string, kind Kind, gpuCount int32, cpuCount int32, memoryCount int32, avgGpuUtilization float32, creationTimestamp time.Time, jobOwner string, endTimestamp NullableTime, ) *JobItemFlat`

NewJobItemFlat instantiates a new JobItemFlat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobItemFlatWithDefaults

`func NewJobItemFlatWithDefaults() *JobItemFlat`

NewJobItemFlatWithDefaults instantiates a new JobItemFlat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *JobItemFlat) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JobItemFlat) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JobItemFlat) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.


### GetId

`func (o *JobItemFlat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobItemFlat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobItemFlat) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JobItemFlat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobItemFlat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobItemFlat) SetName(v string)`

SetName sets Name field to given value.


### GetK8sName

`func (o *JobItemFlat) GetK8sName() string`

GetK8sName returns the K8sName field if non-nil, zero value otherwise.

### GetK8sNameOk

`func (o *JobItemFlat) GetK8sNameOk() (*string, bool)`

GetK8sNameOk returns a tuple with the K8sName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sName

`func (o *JobItemFlat) SetK8sName(v string)`

SetK8sName sets K8sName field to given value.


### GetNamespace

`func (o *JobItemFlat) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JobItemFlat) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JobItemFlat) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetKind

`func (o *JobItemFlat) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JobItemFlat) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JobItemFlat) SetKind(v Kind)`

SetKind sets Kind field to given value.


### GetGpuCount

`func (o *JobItemFlat) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *JobItemFlat) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *JobItemFlat) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetCpuCount

`func (o *JobItemFlat) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *JobItemFlat) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *JobItemFlat) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.


### GetMemoryCount

`func (o *JobItemFlat) GetMemoryCount() int32`

GetMemoryCount returns the MemoryCount field if non-nil, zero value otherwise.

### GetMemoryCountOk

`func (o *JobItemFlat) GetMemoryCountOk() (*int32, bool)`

GetMemoryCountOk returns a tuple with the MemoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCount

`func (o *JobItemFlat) SetMemoryCount(v int32)`

SetMemoryCount sets MemoryCount field to given value.


### GetAvgGpuUtilization

`func (o *JobItemFlat) GetAvgGpuUtilization() float32`

GetAvgGpuUtilization returns the AvgGpuUtilization field if non-nil, zero value otherwise.

### GetAvgGpuUtilizationOk

`func (o *JobItemFlat) GetAvgGpuUtilizationOk() (*float32, bool)`

GetAvgGpuUtilizationOk returns a tuple with the AvgGpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgGpuUtilization

`func (o *JobItemFlat) SetAvgGpuUtilization(v float32)`

SetAvgGpuUtilization sets AvgGpuUtilization field to given value.


### GetCreationTimestamp

`func (o *JobItemFlat) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *JobItemFlat) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *JobItemFlat) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetJobOwner

`func (o *JobItemFlat) GetJobOwner() string`

GetJobOwner returns the JobOwner field if non-nil, zero value otherwise.

### GetJobOwnerOk

`func (o *JobItemFlat) GetJobOwnerOk() (*string, bool)`

GetJobOwnerOk returns a tuple with the JobOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOwner

`func (o *JobItemFlat) SetJobOwner(v string)`

SetJobOwner sets JobOwner field to given value.


### GetEndTimestamp

`func (o *JobItemFlat) GetEndTimestamp() time.Time`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *JobItemFlat) GetEndTimestampOk() (*time.Time, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *JobItemFlat) SetEndTimestamp(v time.Time)`

SetEndTimestamp sets EndTimestamp field to given value.


### SetEndTimestampNil

`func (o *JobItemFlat) SetEndTimestampNil(b bool)`

 SetEndTimestampNil sets the value for EndTimestamp to be an explicit nil

### UnsetEndTimestamp
`func (o *JobItemFlat) UnsetEndTimestamp()`

UnsetEndTimestamp ensures that no value is present for EndTimestamp, not even an explicit nil
### GetStatus

`func (o *JobItemFlat) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobItemFlat) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobItemFlat) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobItemFlat) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


