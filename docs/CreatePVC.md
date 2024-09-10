# CreatePVC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Storage** | **int32** |  | 
**StorageUnit** | Pointer to **string** |  | [optional] [default to "Gi"]

## Methods

### NewCreatePVC

`func NewCreatePVC(name string, namespace string, storage int32, ) *CreatePVC`

NewCreatePVC instantiates a new CreatePVC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePVCWithDefaults

`func NewCreatePVCWithDefaults() *CreatePVC`

NewCreatePVCWithDefaults instantiates a new CreatePVC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePVC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePVC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePVC) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreatePVC) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreatePVC) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreatePVC) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetStorage

`func (o *CreatePVC) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreatePVC) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreatePVC) SetStorage(v int32)`

SetStorage sets Storage field to given value.


### GetStorageUnit

`func (o *CreatePVC) GetStorageUnit() string`

GetStorageUnit returns the StorageUnit field if non-nil, zero value otherwise.

### GetStorageUnitOk

`func (o *CreatePVC) GetStorageUnitOk() (*string, bool)`

GetStorageUnitOk returns a tuple with the StorageUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUnit

`func (o *CreatePVC) SetStorageUnit(v string)`

SetStorageUnit sets StorageUnit field to given value.

### HasStorageUnit

`func (o *CreatePVC) HasStorageUnit() bool`

HasStorageUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


