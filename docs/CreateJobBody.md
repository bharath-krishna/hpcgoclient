# CreateJobBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Cluster** | [**Cluster**](Cluster.md) |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Image** | **string** |  | 
**JobPriority** | Pointer to [**JobPriority**](JobPriority.md) |  | [optional] 
**Workdir** | Pointer to **string** |  | [optional] [default to "/"]
**Command** | Pointer to **string** |  | [optional] [default to ""]
**Interactive** | Pointer to **bool** |  | [optional] [default to false]
**Distributed** | Pointer to **bool** |  | [optional] [default to false]
**DistributedReplicas** | Pointer to **int32** |  | [optional] [default to 2]
**Gpus** | Pointer to **int32** |  | [optional] [default to 0]
**Cpus** | Pointer to **int32** |  | [optional] [default to 20]
**MemoryGb** | Pointer to **int32** |  | [optional] [default to 200]
**SharedMemoryGb** | Pointer to **int32** |  | [optional] [default to 0]
**Pvcs** | Pointer to [**[]PVCs**](PVCs.md) |  | [optional] [default to []]
**Envs** | Pointer to [**[]EnvVar**](EnvVar.md) |  | [optional] [default to []]
**Sysenvs** | Pointer to [**[]EnvVar**](EnvVar.md) |  | [optional] [default to []]

## Methods

### NewCreateJobBody

`func NewCreateJobBody(name string, namespace string, cluster Cluster, image string, ) *CreateJobBody`

NewCreateJobBody instantiates a new CreateJobBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobBodyWithDefaults

`func NewCreateJobBodyWithDefaults() *CreateJobBody`

NewCreateJobBodyWithDefaults instantiates a new CreateJobBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateJobBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateJobBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateJobBody) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateJobBody) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateJobBody) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateJobBody) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetCluster

`func (o *CreateJobBody) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreateJobBody) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreateJobBody) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.


### GetDescription

`func (o *CreateJobBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateJobBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateJobBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateJobBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *CreateJobBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateJobBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateJobBody) SetImage(v string)`

SetImage sets Image field to given value.


### GetJobPriority

`func (o *CreateJobBody) GetJobPriority() JobPriority`

GetJobPriority returns the JobPriority field if non-nil, zero value otherwise.

### GetJobPriorityOk

`func (o *CreateJobBody) GetJobPriorityOk() (*JobPriority, bool)`

GetJobPriorityOk returns a tuple with the JobPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobPriority

`func (o *CreateJobBody) SetJobPriority(v JobPriority)`

SetJobPriority sets JobPriority field to given value.

### HasJobPriority

`func (o *CreateJobBody) HasJobPriority() bool`

HasJobPriority returns a boolean if a field has been set.

### GetWorkdir

`func (o *CreateJobBody) GetWorkdir() string`

GetWorkdir returns the Workdir field if non-nil, zero value otherwise.

### GetWorkdirOk

`func (o *CreateJobBody) GetWorkdirOk() (*string, bool)`

GetWorkdirOk returns a tuple with the Workdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkdir

`func (o *CreateJobBody) SetWorkdir(v string)`

SetWorkdir sets Workdir field to given value.

### HasWorkdir

`func (o *CreateJobBody) HasWorkdir() bool`

HasWorkdir returns a boolean if a field has been set.

### GetCommand

`func (o *CreateJobBody) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateJobBody) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateJobBody) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CreateJobBody) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetInteractive

`func (o *CreateJobBody) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *CreateJobBody) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *CreateJobBody) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *CreateJobBody) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetDistributed

`func (o *CreateJobBody) GetDistributed() bool`

GetDistributed returns the Distributed field if non-nil, zero value otherwise.

### GetDistributedOk

`func (o *CreateJobBody) GetDistributedOk() (*bool, bool)`

GetDistributedOk returns a tuple with the Distributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributed

`func (o *CreateJobBody) SetDistributed(v bool)`

SetDistributed sets Distributed field to given value.

### HasDistributed

`func (o *CreateJobBody) HasDistributed() bool`

HasDistributed returns a boolean if a field has been set.

### GetDistributedReplicas

`func (o *CreateJobBody) GetDistributedReplicas() int32`

GetDistributedReplicas returns the DistributedReplicas field if non-nil, zero value otherwise.

### GetDistributedReplicasOk

`func (o *CreateJobBody) GetDistributedReplicasOk() (*int32, bool)`

GetDistributedReplicasOk returns a tuple with the DistributedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedReplicas

`func (o *CreateJobBody) SetDistributedReplicas(v int32)`

SetDistributedReplicas sets DistributedReplicas field to given value.

### HasDistributedReplicas

`func (o *CreateJobBody) HasDistributedReplicas() bool`

HasDistributedReplicas returns a boolean if a field has been set.

### GetGpus

`func (o *CreateJobBody) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *CreateJobBody) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *CreateJobBody) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *CreateJobBody) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetCpus

`func (o *CreateJobBody) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CreateJobBody) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CreateJobBody) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CreateJobBody) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemoryGb

`func (o *CreateJobBody) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *CreateJobBody) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *CreateJobBody) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *CreateJobBody) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetSharedMemoryGb

`func (o *CreateJobBody) GetSharedMemoryGb() int32`

GetSharedMemoryGb returns the SharedMemoryGb field if non-nil, zero value otherwise.

### GetSharedMemoryGbOk

`func (o *CreateJobBody) GetSharedMemoryGbOk() (*int32, bool)`

GetSharedMemoryGbOk returns a tuple with the SharedMemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMemoryGb

`func (o *CreateJobBody) SetSharedMemoryGb(v int32)`

SetSharedMemoryGb sets SharedMemoryGb field to given value.

### HasSharedMemoryGb

`func (o *CreateJobBody) HasSharedMemoryGb() bool`

HasSharedMemoryGb returns a boolean if a field has been set.

### GetPvcs

`func (o *CreateJobBody) GetPvcs() []PVCs`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *CreateJobBody) GetPvcsOk() (*[]PVCs, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *CreateJobBody) SetPvcs(v []PVCs)`

SetPvcs sets Pvcs field to given value.

### HasPvcs

`func (o *CreateJobBody) HasPvcs() bool`

HasPvcs returns a boolean if a field has been set.

### GetEnvs

`func (o *CreateJobBody) GetEnvs() []EnvVar`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *CreateJobBody) GetEnvsOk() (*[]EnvVar, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *CreateJobBody) SetEnvs(v []EnvVar)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *CreateJobBody) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetSysenvs

`func (o *CreateJobBody) GetSysenvs() []EnvVar`

GetSysenvs returns the Sysenvs field if non-nil, zero value otherwise.

### GetSysenvsOk

`func (o *CreateJobBody) GetSysenvsOk() (*[]EnvVar, bool)`

GetSysenvsOk returns a tuple with the Sysenvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysenvs

`func (o *CreateJobBody) SetSysenvs(v []EnvVar)`

SetSysenvs sets Sysenvs field to given value.

### HasSysenvs

`func (o *CreateJobBody) HasSysenvs() bool`

HasSysenvs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


