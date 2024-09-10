# DataTransferStatusList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**PaginationMetadata**](PaginationMetadata.md) |  | 
**Items** | [**[]DataTransferStatus**](DataTransferStatus.md) |  | 

## Methods

### NewDataTransferStatusList

`func NewDataTransferStatusList(metadata PaginationMetadata, items []DataTransferStatus, ) *DataTransferStatusList`

NewDataTransferStatusList instantiates a new DataTransferStatusList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferStatusListWithDefaults

`func NewDataTransferStatusListWithDefaults() *DataTransferStatusList`

NewDataTransferStatusListWithDefaults instantiates a new DataTransferStatusList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DataTransferStatusList) GetMetadata() PaginationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataTransferStatusList) GetMetadataOk() (*PaginationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataTransferStatusList) SetMetadata(v PaginationMetadata)`

SetMetadata sets Metadata field to given value.


### GetItems

`func (o *DataTransferStatusList) GetItems() []DataTransferStatus`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DataTransferStatusList) GetItemsOk() (*[]DataTransferStatus, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DataTransferStatusList) SetItems(v []DataTransferStatus)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


