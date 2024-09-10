# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TerminationState** | Pointer to [**NullableContainerTerminationState**](ContainerTerminationState.md) |  | [optional] 

## Methods

### NewContainer

`func NewContainer(name string, ) *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Container) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Container) SetName(v string)`

SetName sets Name field to given value.


### GetTerminationState

`func (o *Container) GetTerminationState() ContainerTerminationState`

GetTerminationState returns the TerminationState field if non-nil, zero value otherwise.

### GetTerminationStateOk

`func (o *Container) GetTerminationStateOk() (*ContainerTerminationState, bool)`

GetTerminationStateOk returns a tuple with the TerminationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationState

`func (o *Container) SetTerminationState(v ContainerTerminationState)`

SetTerminationState sets TerminationState field to given value.

### HasTerminationState

`func (o *Container) HasTerminationState() bool`

HasTerminationState returns a boolean if a field has been set.

### SetTerminationStateNil

`func (o *Container) SetTerminationStateNil(b bool)`

 SetTerminationStateNil sets the value for TerminationState to be an explicit nil

### UnsetTerminationState
`func (o *Container) UnsetTerminationState()`

UnsetTerminationState ensures that no value is present for TerminationState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


