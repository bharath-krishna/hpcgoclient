/*
HPC Portal - API

An interface for working with HPC clusters.

API version: 1.9.10-9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the JobItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobItem{}

// JobItem struct for JobItem
type JobItem struct {
	Cluster Cluster `json:"cluster"`
	Id string `json:"id"`
	Name string `json:"name"`
	K8sName string `json:"k8s_name"`
	Namespace string `json:"namespace"`
	Kind Kind `json:"kind"`
	Command string `json:"command"`
	Image string `json:"image"`
	GpuCount int32 `json:"gpu_count"`
	CpuLimit int32 `json:"cpu_limit"`
	CpuRequest int32 `json:"cpu_request"`
	MemoryLimitGb int32 `json:"memory_limit_gb"`
	MemoryRequestGb int32 `json:"memory_request_gb"`
	SharedMemoryGb int32 `json:"shared_memory_gb"`
	WorkingDir string `json:"working_dir"`
	AvgGpuUtilization float32 `json:"avg_gpu_utilization"`
	Interactive bool `json:"interactive"`
	Services []Service `json:"services"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
	EndTimestamp NullableTime `json:"end_timestamp"`
	Pods []Pod `json:"pods"`
	Pvcs []PVCs `json:"pvcs"`
	Envs []EnvVar `json:"envs"`
	Sysenvs []EnvVar `json:"sysenvs"`
	JobOwner string `json:"job_owner"`
	JobDescription string `json:"job_description"`
	JobPriority JobPriority `json:"job_priority"`
	Status *Status `json:"status,omitempty"`
	GrafanaEmbedUrl string `json:"grafana_embed_url"`
	GrafanaDetailsUrl string `json:"grafana_details_url"`
}

type _JobItem JobItem

// NewJobItem instantiates a new JobItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobItem(cluster Cluster, id string, name string, k8sName string, namespace string, kind Kind, command string, image string, gpuCount int32, cpuLimit int32, cpuRequest int32, memoryLimitGb int32, memoryRequestGb int32, sharedMemoryGb int32, workingDir string, avgGpuUtilization float32, interactive bool, services []Service, creationTimestamp time.Time, endTimestamp NullableTime, pods []Pod, pvcs []PVCs, envs []EnvVar, sysenvs []EnvVar, jobOwner string, jobDescription string, jobPriority JobPriority, grafanaEmbedUrl string, grafanaDetailsUrl string) *JobItem {
	this := JobItem{}
	this.Cluster = cluster
	this.Id = id
	this.Name = name
	this.K8sName = k8sName
	this.Namespace = namespace
	this.Kind = kind
	this.Command = command
	this.Image = image
	this.GpuCount = gpuCount
	this.CpuLimit = cpuLimit
	this.CpuRequest = cpuRequest
	this.MemoryLimitGb = memoryLimitGb
	this.MemoryRequestGb = memoryRequestGb
	this.SharedMemoryGb = sharedMemoryGb
	this.WorkingDir = workingDir
	this.AvgGpuUtilization = avgGpuUtilization
	this.Interactive = interactive
	this.Services = services
	this.CreationTimestamp = creationTimestamp
	this.EndTimestamp = endTimestamp
	this.Pods = pods
	this.Pvcs = pvcs
	this.Envs = envs
	this.Sysenvs = sysenvs
	this.JobOwner = jobOwner
	this.JobDescription = jobDescription
	this.JobPriority = jobPriority
	this.GrafanaEmbedUrl = grafanaEmbedUrl
	this.GrafanaDetailsUrl = grafanaDetailsUrl
	return &this
}

// NewJobItemWithDefaults instantiates a new JobItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobItemWithDefaults() *JobItem {
	this := JobItem{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *JobItem) GetCluster() Cluster {
	if o == nil {
		var ret Cluster
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetClusterOk() (*Cluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *JobItem) SetCluster(v Cluster) {
	o.Cluster = v
}

// GetId returns the Id field value
func (o *JobItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *JobItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobItem) SetName(v string) {
	o.Name = v
}

// GetK8sName returns the K8sName field value
func (o *JobItem) GetK8sName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.K8sName
}

// GetK8sNameOk returns a tuple with the K8sName field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetK8sNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.K8sName, true
}

// SetK8sName sets field value
func (o *JobItem) SetK8sName(v string) {
	o.K8sName = v
}

// GetNamespace returns the Namespace field value
func (o *JobItem) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *JobItem) SetNamespace(v string) {
	o.Namespace = v
}

// GetKind returns the Kind field value
func (o *JobItem) GetKind() Kind {
	if o == nil {
		var ret Kind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetKindOk() (*Kind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *JobItem) SetKind(v Kind) {
	o.Kind = v
}

// GetCommand returns the Command field value
func (o *JobItem) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *JobItem) SetCommand(v string) {
	o.Command = v
}

// GetImage returns the Image field value
func (o *JobItem) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *JobItem) SetImage(v string) {
	o.Image = v
}

// GetGpuCount returns the GpuCount field value
func (o *JobItem) GetGpuCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GpuCount
}

// GetGpuCountOk returns a tuple with the GpuCount field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetGpuCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GpuCount, true
}

// SetGpuCount sets field value
func (o *JobItem) SetGpuCount(v int32) {
	o.GpuCount = v
}

// GetCpuLimit returns the CpuLimit field value
func (o *JobItem) GetCpuLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetCpuLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuLimit, true
}

// SetCpuLimit sets field value
func (o *JobItem) SetCpuLimit(v int32) {
	o.CpuLimit = v
}

// GetCpuRequest returns the CpuRequest field value
func (o *JobItem) GetCpuRequest() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuRequest
}

// GetCpuRequestOk returns a tuple with the CpuRequest field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetCpuRequestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuRequest, true
}

// SetCpuRequest sets field value
func (o *JobItem) SetCpuRequest(v int32) {
	o.CpuRequest = v
}

// GetMemoryLimitGb returns the MemoryLimitGb field value
func (o *JobItem) GetMemoryLimitGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryLimitGb
}

// GetMemoryLimitGbOk returns a tuple with the MemoryLimitGb field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetMemoryLimitGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryLimitGb, true
}

// SetMemoryLimitGb sets field value
func (o *JobItem) SetMemoryLimitGb(v int32) {
	o.MemoryLimitGb = v
}

// GetMemoryRequestGb returns the MemoryRequestGb field value
func (o *JobItem) GetMemoryRequestGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryRequestGb
}

// GetMemoryRequestGbOk returns a tuple with the MemoryRequestGb field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetMemoryRequestGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryRequestGb, true
}

// SetMemoryRequestGb sets field value
func (o *JobItem) SetMemoryRequestGb(v int32) {
	o.MemoryRequestGb = v
}

// GetSharedMemoryGb returns the SharedMemoryGb field value
func (o *JobItem) GetSharedMemoryGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SharedMemoryGb
}

// GetSharedMemoryGbOk returns a tuple with the SharedMemoryGb field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetSharedMemoryGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharedMemoryGb, true
}

// SetSharedMemoryGb sets field value
func (o *JobItem) SetSharedMemoryGb(v int32) {
	o.SharedMemoryGb = v
}

// GetWorkingDir returns the WorkingDir field value
func (o *JobItem) GetWorkingDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkingDir
}

// GetWorkingDirOk returns a tuple with the WorkingDir field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetWorkingDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkingDir, true
}

// SetWorkingDir sets field value
func (o *JobItem) SetWorkingDir(v string) {
	o.WorkingDir = v
}

// GetAvgGpuUtilization returns the AvgGpuUtilization field value
func (o *JobItem) GetAvgGpuUtilization() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvgGpuUtilization
}

// GetAvgGpuUtilizationOk returns a tuple with the AvgGpuUtilization field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetAvgGpuUtilizationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgGpuUtilization, true
}

// SetAvgGpuUtilization sets field value
func (o *JobItem) SetAvgGpuUtilization(v float32) {
	o.AvgGpuUtilization = v
}

// GetInteractive returns the Interactive field value
func (o *JobItem) GetInteractive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Interactive
}

// GetInteractiveOk returns a tuple with the Interactive field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetInteractiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactive, true
}

// SetInteractive sets field value
func (o *JobItem) SetInteractive(v bool) {
	o.Interactive = v
}

// GetServices returns the Services field value
func (o *JobItem) GetServices() []Service {
	if o == nil {
		var ret []Service
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetServicesOk() ([]Service, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *JobItem) SetServices(v []Service) {
	o.Services = v
}

// GetCreationTimestamp returns the CreationTimestamp field value
func (o *JobItem) GetCreationTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value
func (o *JobItem) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = v
}

// GetEndTimestamp returns the EndTimestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *JobItem) GetEndTimestamp() time.Time {
	if o == nil || o.EndTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndTimestamp.Get()
}

// GetEndTimestampOk returns a tuple with the EndTimestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobItem) GetEndTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTimestamp.Get(), o.EndTimestamp.IsSet()
}

// SetEndTimestamp sets field value
func (o *JobItem) SetEndTimestamp(v time.Time) {
	o.EndTimestamp.Set(&v)
}

// GetPods returns the Pods field value
func (o *JobItem) GetPods() []Pod {
	if o == nil {
		var ret []Pod
		return ret
	}

	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetPodsOk() ([]Pod, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pods, true
}

// SetPods sets field value
func (o *JobItem) SetPods(v []Pod) {
	o.Pods = v
}

// GetPvcs returns the Pvcs field value
func (o *JobItem) GetPvcs() []PVCs {
	if o == nil {
		var ret []PVCs
		return ret
	}

	return o.Pvcs
}

// GetPvcsOk returns a tuple with the Pvcs field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetPvcsOk() ([]PVCs, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pvcs, true
}

// SetPvcs sets field value
func (o *JobItem) SetPvcs(v []PVCs) {
	o.Pvcs = v
}

// GetEnvs returns the Envs field value
func (o *JobItem) GetEnvs() []EnvVar {
	if o == nil {
		var ret []EnvVar
		return ret
	}

	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetEnvsOk() ([]EnvVar, bool) {
	if o == nil {
		return nil, false
	}
	return o.Envs, true
}

// SetEnvs sets field value
func (o *JobItem) SetEnvs(v []EnvVar) {
	o.Envs = v
}

// GetSysenvs returns the Sysenvs field value
func (o *JobItem) GetSysenvs() []EnvVar {
	if o == nil {
		var ret []EnvVar
		return ret
	}

	return o.Sysenvs
}

// GetSysenvsOk returns a tuple with the Sysenvs field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetSysenvsOk() ([]EnvVar, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sysenvs, true
}

// SetSysenvs sets field value
func (o *JobItem) SetSysenvs(v []EnvVar) {
	o.Sysenvs = v
}

// GetJobOwner returns the JobOwner field value
func (o *JobItem) GetJobOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobOwner
}

// GetJobOwnerOk returns a tuple with the JobOwner field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetJobOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobOwner, true
}

// SetJobOwner sets field value
func (o *JobItem) SetJobOwner(v string) {
	o.JobOwner = v
}

// GetJobDescription returns the JobDescription field value
func (o *JobItem) GetJobDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobDescription
}

// GetJobDescriptionOk returns a tuple with the JobDescription field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetJobDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobDescription, true
}

// SetJobDescription sets field value
func (o *JobItem) SetJobDescription(v string) {
	o.JobDescription = v
}

// GetJobPriority returns the JobPriority field value
func (o *JobItem) GetJobPriority() JobPriority {
	if o == nil {
		var ret JobPriority
		return ret
	}

	return o.JobPriority
}

// GetJobPriorityOk returns a tuple with the JobPriority field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetJobPriorityOk() (*JobPriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobPriority, true
}

// SetJobPriority sets field value
func (o *JobItem) SetJobPriority(v JobPriority) {
	o.JobPriority = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JobItem) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobItem) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JobItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *JobItem) SetStatus(v Status) {
	o.Status = &v
}

// GetGrafanaEmbedUrl returns the GrafanaEmbedUrl field value
func (o *JobItem) GetGrafanaEmbedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrafanaEmbedUrl
}

// GetGrafanaEmbedUrlOk returns a tuple with the GrafanaEmbedUrl field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetGrafanaEmbedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrafanaEmbedUrl, true
}

// SetGrafanaEmbedUrl sets field value
func (o *JobItem) SetGrafanaEmbedUrl(v string) {
	o.GrafanaEmbedUrl = v
}

// GetGrafanaDetailsUrl returns the GrafanaDetailsUrl field value
func (o *JobItem) GetGrafanaDetailsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrafanaDetailsUrl
}

// GetGrafanaDetailsUrlOk returns a tuple with the GrafanaDetailsUrl field value
// and a boolean to check if the value has been set.
func (o *JobItem) GetGrafanaDetailsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrafanaDetailsUrl, true
}

// SetGrafanaDetailsUrl sets field value
func (o *JobItem) SetGrafanaDetailsUrl(v string) {
	o.GrafanaDetailsUrl = v
}

func (o JobItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["k8s_name"] = o.K8sName
	toSerialize["namespace"] = o.Namespace
	toSerialize["kind"] = o.Kind
	toSerialize["command"] = o.Command
	toSerialize["image"] = o.Image
	toSerialize["gpu_count"] = o.GpuCount
	toSerialize["cpu_limit"] = o.CpuLimit
	toSerialize["cpu_request"] = o.CpuRequest
	toSerialize["memory_limit_gb"] = o.MemoryLimitGb
	toSerialize["memory_request_gb"] = o.MemoryRequestGb
	toSerialize["shared_memory_gb"] = o.SharedMemoryGb
	toSerialize["working_dir"] = o.WorkingDir
	toSerialize["avg_gpu_utilization"] = o.AvgGpuUtilization
	toSerialize["interactive"] = o.Interactive
	toSerialize["services"] = o.Services
	toSerialize["creation_timestamp"] = o.CreationTimestamp
	toSerialize["end_timestamp"] = o.EndTimestamp.Get()
	toSerialize["pods"] = o.Pods
	toSerialize["pvcs"] = o.Pvcs
	toSerialize["envs"] = o.Envs
	toSerialize["sysenvs"] = o.Sysenvs
	toSerialize["job_owner"] = o.JobOwner
	toSerialize["job_description"] = o.JobDescription
	toSerialize["job_priority"] = o.JobPriority
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["grafana_embed_url"] = o.GrafanaEmbedUrl
	toSerialize["grafana_details_url"] = o.GrafanaDetailsUrl
	return toSerialize, nil
}

func (o *JobItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster",
		"id",
		"name",
		"k8s_name",
		"namespace",
		"kind",
		"command",
		"image",
		"gpu_count",
		"cpu_limit",
		"cpu_request",
		"memory_limit_gb",
		"memory_request_gb",
		"shared_memory_gb",
		"working_dir",
		"avg_gpu_utilization",
		"interactive",
		"services",
		"creation_timestamp",
		"end_timestamp",
		"pods",
		"pvcs",
		"envs",
		"sysenvs",
		"job_owner",
		"job_description",
		"job_priority",
		"grafana_embed_url",
		"grafana_details_url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varJobItem := _JobItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobItem)

	if err != nil {
		return err
	}

	*o = JobItem(varJobItem)

	return err
}

type NullableJobItem struct {
	value *JobItem
	isSet bool
}

func (v NullableJobItem) Get() *JobItem {
	return v.value
}

func (v *NullableJobItem) Set(val *JobItem) {
	v.value = val
	v.isSet = true
}

func (v NullableJobItem) IsSet() bool {
	return v.isSet
}

func (v *NullableJobItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobItem(val *JobItem) *NullableJobItem {
	return &NullableJobItem{value: val, isSet: true}
}

func (v NullableJobItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


