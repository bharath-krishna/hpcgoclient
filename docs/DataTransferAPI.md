# \DataTransferAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Target**](DataTransferAPI.md#Target) | **Get** /api/data/target | Get status of target
[**TransferClone**](DataTransferAPI.md#TransferClone) | **Get** /api/data/transfer_clone | Clone data transfer
[**TransferList**](DataTransferAPI.md#TransferList) | **Get** /api/data/transfer_list | List data transfers
[**TransferStart**](DataTransferAPI.md#TransferStart) | **Post** /api/data/transfer_start | Launch DataTransfer pods in both source and destination clusters/namespaces
[**TransferStatus**](DataTransferAPI.md#TransferStatus) | **Get** /api/data/transfer_status | Get data transfer status
[**TransferTerminate**](DataTransferAPI.md#TransferTerminate) | **Delete** /api/data/transfer_terminate | Terminate data transfer



## Target

> map[string]string Target(ctx).TargetStatus(targetStatus).Execute()

Get status of target

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bharath-krishna/hpcgoclient"
)

func main() {
	targetStatus := *openapiclient.NewTargetStatus("Url_example") // TargetStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTransferAPI.Target(context.Background()).TargetStatus(targetStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTransferAPI.Target``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Target`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DataTransferAPI.Target`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetStatus** | [**TargetStatus**](TargetStatus.md) |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferClone

> DataTransferOutput TransferClone(ctx).TransferId(transferId).Execute()

Clone data transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bharath-krishna/hpcgoclient"
)

func main() {
	transferId := "transferId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTransferAPI.TransferClone(context.Background()).TransferId(transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTransferAPI.TransferClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferClone`: DataTransferOutput
	fmt.Fprintf(os.Stdout, "Response from `DataTransferAPI.TransferClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferId** | **string** |  | 

### Return type

[**DataTransferOutput**](DataTransferOutput.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferList

> DataTransferStatusList TransferList(ctx).Status(status).Start(start).End(end).Execute()

List data transfers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bharath-krishna/hpcgoclient"
)

func main() {
	status := openapiclient.Status("All") // Status |  (optional)
	start := int32(56) // int32 |  (optional) (default to 0)
	end := int32(56) // int32 |  (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTransferAPI.TransferList(context.Background()).Status(status).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTransferAPI.TransferList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferList`: DataTransferStatusList
	fmt.Fprintf(os.Stdout, "Response from `DataTransferAPI.TransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**Status**](Status.md) |  | 
 **start** | **int32** |  | [default to 0]
 **end** | **int32** |  | [default to 15]

### Return type

[**DataTransferStatusList**](DataTransferStatusList.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferStart

> string TransferStart(ctx).DataTransferInput(dataTransferInput).Execute()

Launch DataTransfer pods in both source and destination clusters/namespaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bharath-krishna/hpcgoclient"
)

func main() {
	dataTransferInput := *openapiclient.NewDataTransferInput("TransferId_example", "Description_example", *openapiclient.NewDataTransferEndpoint(openapiclient.Cluster("All"), "Namespace_example", "Pvc_example", "PvcPath_example"), *openapiclient.NewDataTransferEndpoint(openapiclient.Cluster("All"), "Namespace_example", "Pvc_example", "PvcPath_example")) // DataTransferInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTransferAPI.TransferStart(context.Background()).DataTransferInput(dataTransferInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTransferAPI.TransferStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferStart`: string
	fmt.Fprintf(os.Stdout, "Response from `DataTransferAPI.TransferStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataTransferInput** | [**DataTransferInput**](DataTransferInput.md) |  | 

### Return type

**string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferStatus

> DataTransferStatus TransferStatus(ctx).TransferId(transferId).Execute()

Get data transfer status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bharath-krishna/hpcgoclient"
)

func main() {
	transferId := "transferId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTransferAPI.TransferStatus(context.Background()).TransferId(transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTransferAPI.TransferStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferStatus`: DataTransferStatus
	fmt.Fprintf(os.Stdout, "Response from `DataTransferAPI.TransferStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferId** | **string** |  | 

### Return type

[**DataTransferStatus**](DataTransferStatus.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferTerminate

> string TransferTerminate(ctx).TransferId(transferId).Execute()

Terminate data transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bharath-krishna/hpcgoclient"
)

func main() {
	transferId := "transferId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTransferAPI.TransferTerminate(context.Background()).TransferId(transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTransferAPI.TransferTerminate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferTerminate`: string
	fmt.Fprintf(os.Stdout, "Response from `DataTransferAPI.TransferTerminate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferTerminateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferId** | **string** |  | 

### Return type

**string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

