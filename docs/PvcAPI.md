# \PvcAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedPvcs**](PvcAPI.md#CreateNamespacedPvcs) | **Post** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs | Create PVC
[**DeleteNamespacedPvcs**](PvcAPI.md#DeleteNamespacedPvcs) | **Delete** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs/{pvc_name} | Delete PVC
[**GetNamespacedPvcs**](PvcAPI.md#GetNamespacedPvcs) | **Get** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs | List namespaced PVCs
[**GetNamespacedPvcsDetails**](PvcAPI.md#GetNamespacedPvcsDetails) | **Get** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs_details | Get PVCs
[**GetPvcStatus**](PvcAPI.md#GetPvcStatus) | **Get** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs/{pvc_name}/status | Get PVC&#39;s status
[**StartPvcviewer**](PvcAPI.md#StartPvcviewer) | **Post** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs/{pvc_name}/viewer/start | Start PVCViewer pod
[**StopPvcviewer**](PvcAPI.md#StopPvcviewer) | **Delete** /api/data/clusters/{cluster}/namespaces/{namespace}/pvcs/{pvc_name}/viewer/stop | Stop PVCViewer pod



## CreateNamespacedPvcs

> string CreateNamespacedPvcs(ctx, cluster, namespace).CreatePVC(createPVC).Execute()

Create PVC



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 
	createPVC := *openapiclient.NewCreatePVC("Name_example", "Namespace_example", int32(123)) // CreatePVC | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.CreateNamespacedPvcs(context.Background(), cluster, namespace).CreatePVC(createPVC).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.CreateNamespacedPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedPvcs`: string
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.CreateNamespacedPvcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createPVC** | [**CreatePVC**](CreatePVC.md) |  | 

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


## DeleteNamespacedPvcs

> string DeleteNamespacedPvcs(ctx, cluster, namespace, pvcName).Execute()

Delete PVC



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 
	pvcName := "pvcName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.DeleteNamespacedPvcs(context.Background(), cluster, namespace, pvcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.DeleteNamespacedPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespacedPvcs`: string
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.DeleteNamespacedPvcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 
**pvcName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetNamespacedPvcs

> []PVCItemBasic GetNamespacedPvcs(ctx, cluster, namespace).Execute()

List namespaced PVCs



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.GetNamespacedPvcs(context.Background(), cluster, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.GetNamespacedPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespacedPvcs`: []PVCItemBasic
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.GetNamespacedPvcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacedPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PVCItemBasic**](PVCItemBasic.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespacedPvcsDetails

> []PVCItem GetNamespacedPvcsDetails(ctx, cluster, namespace).Execute()

Get PVCs



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.GetNamespacedPvcsDetails(context.Background(), cluster, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.GetNamespacedPvcsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespacedPvcsDetails`: []PVCItem
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.GetNamespacedPvcsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacedPvcsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PVCItem**](PVCItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPvcStatus

> PVCStatus GetPvcStatus(ctx, cluster, namespace, pvcName).Execute()

Get PVC's status



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 
	pvcName := "pvcName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.GetPvcStatus(context.Background(), cluster, namespace, pvcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.GetPvcStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPvcStatus`: PVCStatus
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.GetPvcStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 
**pvcName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPvcStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PVCStatus**](PVCStatus.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPvcviewer

> string StartPvcviewer(ctx, cluster, namespace, pvcName).Execute()

Start PVCViewer pod



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 
	pvcName := "pvcName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.StartPvcviewer(context.Background(), cluster, namespace, pvcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.StartPvcviewer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartPvcviewer`: string
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.StartPvcviewer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 
**pvcName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartPvcviewerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## StopPvcviewer

> string StopPvcviewer(ctx, cluster, namespace, pvcName).Execute()

Stop PVCViewer pod



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
	cluster := openapiclient.Cluster("All") // Cluster | 
	namespace := "namespace_example" // string | 
	pvcName := "pvcName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvcAPI.StopPvcviewer(context.Background(), cluster, namespace, pvcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvcAPI.StopPvcviewer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopPvcviewer`: string
	fmt.Fprintf(os.Stdout, "Response from `PvcAPI.StopPvcviewer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 
**pvcName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPvcviewerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

