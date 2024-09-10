# JobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**K8sName** | **string** |  | 
**Namespace** | **string** |  | 
**Kind** | [**Kind**](Kind.md) |  | 
**Command** | **string** |  | 
**Image** | **string** |  | 
**GpuCount** | **int32** |  | 
**CpuLimit** | **int32** |  | 
**CpuRequest** | **int32** |  | 
**MemoryLimitGb** | **int32** |  | 
**MemoryRequestGb** | **int32** |  | 
**SharedMemoryGb** | **int32** |  | 
**WorkingDir** | **string** |  | 
**AvgGpuUtilization** | **float32** |  | 
**Interactive** | **bool** |  | 
**Services** | [**[]Service**](Service.md) |  | 
**CreationTimestamp** | **time.Time** |  | 
**EndTimestamp** | **NullableTime** |  | 
**Pods** | [**[]Pod**](Pod.md) |  | 
**Pvcs** | [**[]PVCs**](PVCs.md) |  | 
**Envs** | [**[]EnvVar**](EnvVar.md) |  | 
**Sysenvs** | [**[]EnvVar**](EnvVar.md) |  | 
**JobOwner** | **string** |  | 
**JobDescription** | **string** |  | 
**JobPriority** | [**JobPriority**](JobPriority.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**GrafanaEmbedUrl** | **string** |  | 
**GrafanaDetailsUrl** | **string** |  | 

## Methods

### NewJobItem

`func NewJobItem(cluster Cluster, id string, name string, k8sName string, namespace string, kind Kind, command string, image string, gpuCount int32, cpuLimit int32, cpuRequest int32, memoryLimitGb int32, memoryRequestGb int32, sharedMemoryGb int32, workingDir string, avgGpuUtilization float32, interactive bool, services []Service, creationTimestamp time.Time, endTimestamp NullableTime, pods []Pod, pvcs []PVCs, envs []EnvVar, sysenvs []EnvVar, jobOwner string, jobDescription string, jobPriority JobPriority, grafanaEmbedUrl string, grafanaDetailsUrl string, ) *JobItem`

NewJobItem instantiates a new JobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobItemWithDefaults

`func NewJobItemWithDefaults() *JobItem`

NewJobItemWithDefaults instantiates a new JobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *JobItem) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JobItem) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JobItem) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.


### GetId

`func (o *JobItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobItem) SetName(v string)`

SetName sets Name field to given value.


### GetK8sName

`func (o *JobItem) GetK8sName() string`

GetK8sName returns the K8sName field if non-nil, zero value otherwise.

### GetK8sNameOk

`func (o *JobItem) GetK8sNameOk() (*string, bool)`

GetK8sNameOk returns a tuple with the K8sName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sName

`func (o *JobItem) SetK8sName(v string)`

SetK8sName sets K8sName field to given value.


### GetNamespace

`func (o *JobItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JobItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JobItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetKind

`func (o *JobItem) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JobItem) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JobItem) SetKind(v Kind)`

SetKind sets Kind field to given value.


### GetCommand

`func (o *JobItem) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *JobItem) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *JobItem) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetImage

`func (o *JobItem) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *JobItem) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *JobItem) SetImage(v string)`

SetImage sets Image field to given value.


### GetGpuCount

`func (o *JobItem) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *JobItem) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *JobItem) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetCpuLimit

`func (o *JobItem) GetCpuLimit() int32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *JobItem) GetCpuLimitOk() (*int32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *JobItem) SetCpuLimit(v int32)`

SetCpuLimit sets CpuLimit field to given value.


### GetCpuRequest

`func (o *JobItem) GetCpuRequest() int32`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *JobItem) GetCpuRequestOk() (*int32, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *JobItem) SetCpuRequest(v int32)`

SetCpuRequest sets CpuRequest field to given value.


### GetMemoryLimitGb

`func (o *JobItem) GetMemoryLimitGb() int32`

GetMemoryLimitGb returns the MemoryLimitGb field if non-nil, zero value otherwise.

### GetMemoryLimitGbOk

`func (o *JobItem) GetMemoryLimitGbOk() (*int32, bool)`

GetMemoryLimitGbOk returns a tuple with the MemoryLimitGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitGb

`func (o *JobItem) SetMemoryLimitGb(v int32)`

SetMemoryLimitGb sets MemoryLimitGb field to given value.


### GetMemoryRequestGb

`func (o *JobItem) GetMemoryRequestGb() int32`

GetMemoryRequestGb returns the MemoryRequestGb field if non-nil, zero value otherwise.

### GetMemoryRequestGbOk

`func (o *JobItem) GetMemoryRequestGbOk() (*int32, bool)`

GetMemoryRequestGbOk returns a tuple with the MemoryRequestGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequestGb

`func (o *JobItem) SetMemoryRequestGb(v int32)`

SetMemoryRequestGb sets MemoryRequestGb field to given value.


### GetSharedMemoryGb

`func (o *JobItem) GetSharedMemoryGb() int32`

GetSharedMemoryGb returns the SharedMemoryGb field if non-nil, zero value otherwise.

### GetSharedMemoryGbOk

`func (o *JobItem) GetSharedMemoryGbOk() (*int32, bool)`

GetSharedMemoryGbOk returns a tuple with the SharedMemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMemoryGb

`func (o *JobItem) SetSharedMemoryGb(v int32)`

SetSharedMemoryGb sets SharedMemoryGb field to given value.


### GetWorkingDir

`func (o *JobItem) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *JobItem) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *JobItem) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.


### GetAvgGpuUtilization

`func (o *JobItem) GetAvgGpuUtilization() float32`

GetAvgGpuUtilization returns the AvgGpuUtilization field if non-nil, zero value otherwise.

### GetAvgGpuUtilizationOk

`func (o *JobItem) GetAvgGpuUtilizationOk() (*float32, bool)`

GetAvgGpuUtilizationOk returns a tuple with the AvgGpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgGpuUtilization

`func (o *JobItem) SetAvgGpuUtilization(v float32)`

SetAvgGpuUtilization sets AvgGpuUtilization field to given value.


### GetInteractive

`func (o *JobItem) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *JobItem) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *JobItem) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.


### GetServices

`func (o *JobItem) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *JobItem) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *JobItem) SetServices(v []Service)`

SetServices sets Services field to given value.


### GetCreationTimestamp

`func (o *JobItem) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *JobItem) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *JobItem) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetEndTimestamp

`func (o *JobItem) GetEndTimestamp() time.Time`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *JobItem) GetEndTimestampOk() (*time.Time, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *JobItem) SetEndTimestamp(v time.Time)`

SetEndTimestamp sets EndTimestamp field to given value.


### SetEndTimestampNil

`func (o *JobItem) SetEndTimestampNil(b bool)`

 SetEndTimestampNil sets the value for EndTimestamp to be an explicit nil

### UnsetEndTimestamp
`func (o *JobItem) UnsetEndTimestamp()`

UnsetEndTimestamp ensures that no value is present for EndTimestamp, not even an explicit nil
### GetPods

`func (o *JobItem) GetPods() []Pod`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *JobItem) GetPodsOk() (*[]Pod, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *JobItem) SetPods(v []Pod)`

SetPods sets Pods field to given value.


### GetPvcs

`func (o *JobItem) GetPvcs() []PVCs`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *JobItem) GetPvcsOk() (*[]PVCs, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *JobItem) SetPvcs(v []PVCs)`

SetPvcs sets Pvcs field to given value.


### GetEnvs

`func (o *JobItem) GetEnvs() []EnvVar`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *JobItem) GetEnvsOk() (*[]EnvVar, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *JobItem) SetEnvs(v []EnvVar)`

SetEnvs sets Envs field to given value.


### GetSysenvs

`func (o *JobItem) GetSysenvs() []EnvVar`

GetSysenvs returns the Sysenvs field if non-nil, zero value otherwise.

### GetSysenvsOk

`func (o *JobItem) GetSysenvsOk() (*[]EnvVar, bool)`

GetSysenvsOk returns a tuple with the Sysenvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysenvs

`func (o *JobItem) SetSysenvs(v []EnvVar)`

SetSysenvs sets Sysenvs field to given value.


### GetJobOwner

`func (o *JobItem) GetJobOwner() string`

GetJobOwner returns the JobOwner field if non-nil, zero value otherwise.

### GetJobOwnerOk

`func (o *JobItem) GetJobOwnerOk() (*string, bool)`

GetJobOwnerOk returns a tuple with the JobOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOwner

`func (o *JobItem) SetJobOwner(v string)`

SetJobOwner sets JobOwner field to given value.


### GetJobDescription

`func (o *JobItem) GetJobDescription() string`

GetJobDescription returns the JobDescription field if non-nil, zero value otherwise.

### GetJobDescriptionOk

`func (o *JobItem) GetJobDescriptionOk() (*string, bool)`

GetJobDescriptionOk returns a tuple with the JobDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDescription

`func (o *JobItem) SetJobDescription(v string)`

SetJobDescription sets JobDescription field to given value.


### GetJobPriority

`func (o *JobItem) GetJobPriority() JobPriority`

GetJobPriority returns the JobPriority field if non-nil, zero value otherwise.

### GetJobPriorityOk

`func (o *JobItem) GetJobPriorityOk() (*JobPriority, bool)`

GetJobPriorityOk returns a tuple with the JobPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobPriority

`func (o *JobItem) SetJobPriority(v JobPriority)`

SetJobPriority sets JobPriority field to given value.


### GetStatus

`func (o *JobItem) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobItem) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobItem) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGrafanaEmbedUrl

`func (o *JobItem) GetGrafanaEmbedUrl() string`

GetGrafanaEmbedUrl returns the GrafanaEmbedUrl field if non-nil, zero value otherwise.

### GetGrafanaEmbedUrlOk

`func (o *JobItem) GetGrafanaEmbedUrlOk() (*string, bool)`

GetGrafanaEmbedUrlOk returns a tuple with the GrafanaEmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaEmbedUrl

`func (o *JobItem) SetGrafanaEmbedUrl(v string)`

SetGrafanaEmbedUrl sets GrafanaEmbedUrl field to given value.


### GetGrafanaDetailsUrl

`func (o *JobItem) GetGrafanaDetailsUrl() string`

GetGrafanaDetailsUrl returns the GrafanaDetailsUrl field if non-nil, zero value otherwise.

### GetGrafanaDetailsUrlOk

`func (o *JobItem) GetGrafanaDetailsUrlOk() (*string, bool)`

GetGrafanaDetailsUrlOk returns a tuple with the GrafanaDetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaDetailsUrl

`func (o *JobItem) SetGrafanaDetailsUrl(v string)`

SetGrafanaDetailsUrl sets GrafanaDetailsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


