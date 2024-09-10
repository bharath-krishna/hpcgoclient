# JobDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**JobMetadata**](JobMetadata.md) |  | 
**Items** | [**[]JobItemFlat**](JobItemFlat.md) |  | 

## Methods

### NewJobDetailsModel

`func NewJobDetailsModel(metadata JobMetadata, items []JobItemFlat, ) *JobDetailsModel`

NewJobDetailsModel instantiates a new JobDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDetailsModelWithDefaults

`func NewJobDetailsModelWithDefaults() *JobDetailsModel`

NewJobDetailsModelWithDefaults instantiates a new JobDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *JobDetailsModel) GetMetadata() JobMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *JobDetailsModel) GetMetadataOk() (*JobMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *JobDetailsModel) SetMetadata(v JobMetadata)`

SetMetadata sets Metadata field to given value.


### GetItems

`func (o *JobDetailsModel) GetItems() []JobItemFlat`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JobDetailsModel) GetItemsOk() (*[]JobItemFlat, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JobDetailsModel) SetItems(v []JobItemFlat)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


