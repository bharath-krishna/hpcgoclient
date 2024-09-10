# PVCs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [default to "non-existing-pvc"]
**Path** | Pointer to **NullableString** |  | [optional] 
**MountPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPVCs

`func NewPVCs() *PVCs`

NewPVCs instantiates a new PVCs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPVCsWithDefaults

`func NewPVCsWithDefaults() *PVCs`

NewPVCsWithDefaults instantiates a new PVCs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PVCs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PVCs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PVCs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PVCs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *PVCs) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PVCs) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PVCs) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PVCs) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *PVCs) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *PVCs) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMountPath

`func (o *PVCs) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *PVCs) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *PVCs) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *PVCs) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *PVCs) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *PVCs) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


