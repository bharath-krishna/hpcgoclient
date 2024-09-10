# PaginationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** |  | [optional] [default to 0]
**End** | Pointer to **int32** |  | [optional] [default to 15]
**Count** | Pointer to **int32** |  | [optional] [default to 15]
**RemainingItemCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewPaginationMetadata

`func NewPaginationMetadata() *PaginationMetadata`

NewPaginationMetadata instantiates a new PaginationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationMetadataWithDefaults

`func NewPaginationMetadataWithDefaults() *PaginationMetadata`

NewPaginationMetadataWithDefaults instantiates a new PaginationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *PaginationMetadata) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PaginationMetadata) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PaginationMetadata) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *PaginationMetadata) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *PaginationMetadata) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PaginationMetadata) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PaginationMetadata) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PaginationMetadata) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetCount

`func (o *PaginationMetadata) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginationMetadata) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginationMetadata) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginationMetadata) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRemainingItemCount

`func (o *PaginationMetadata) GetRemainingItemCount() int32`

GetRemainingItemCount returns the RemainingItemCount field if non-nil, zero value otherwise.

### GetRemainingItemCountOk

`func (o *PaginationMetadata) GetRemainingItemCountOk() (*int32, bool)`

GetRemainingItemCountOk returns a tuple with the RemainingItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingItemCount

`func (o *PaginationMetadata) SetRemainingItemCount(v int32)`

SetRemainingItemCount sets RemainingItemCount field to given value.

### HasRemainingItemCount

`func (o *PaginationMetadata) HasRemainingItemCount() bool`

HasRemainingItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


