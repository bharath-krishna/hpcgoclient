# DataTransferEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | 
**Namespace** | **string** |  | 
**Pvc** | **string** |  | 
**PvcPath** | **string** |  | 

## Methods

### NewDataTransferEndpoint

`func NewDataTransferEndpoint(cluster Cluster, namespace string, pvc string, pvcPath string, ) *DataTransferEndpoint`

NewDataTransferEndpoint instantiates a new DataTransferEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferEndpointWithDefaults

`func NewDataTransferEndpointWithDefaults() *DataTransferEndpoint`

NewDataTransferEndpointWithDefaults instantiates a new DataTransferEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DataTransferEndpoint) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DataTransferEndpoint) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DataTransferEndpoint) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.


### GetNamespace

`func (o *DataTransferEndpoint) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DataTransferEndpoint) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DataTransferEndpoint) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPvc

`func (o *DataTransferEndpoint) GetPvc() string`

GetPvc returns the Pvc field if non-nil, zero value otherwise.

### GetPvcOk

`func (o *DataTransferEndpoint) GetPvcOk() (*string, bool)`

GetPvcOk returns a tuple with the Pvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvc

`func (o *DataTransferEndpoint) SetPvc(v string)`

SetPvc sets Pvc field to given value.


### GetPvcPath

`func (o *DataTransferEndpoint) GetPvcPath() string`

GetPvcPath returns the PvcPath field if non-nil, zero value otherwise.

### GetPvcPathOk

`func (o *DataTransferEndpoint) GetPvcPathOk() (*string, bool)`

GetPvcPathOk returns a tuple with the PvcPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcPath

`func (o *DataTransferEndpoint) SetPvcPath(v string)`

SetPvcPath sets PvcPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


