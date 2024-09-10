# PVCViewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Status** | [**Status**](Status.md) |  | 

## Methods

### NewPVCViewer

`func NewPVCViewer(url string, status Status, ) *PVCViewer`

NewPVCViewer instantiates a new PVCViewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPVCViewerWithDefaults

`func NewPVCViewerWithDefaults() *PVCViewer`

NewPVCViewerWithDefaults instantiates a new PVCViewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PVCViewer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PVCViewer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PVCViewer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *PVCViewer) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PVCViewer) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PVCViewer) SetStatus(v Status)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


