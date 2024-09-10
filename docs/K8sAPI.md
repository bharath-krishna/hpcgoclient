# \K8sAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterImageDefault**](K8sAPI.md#ClusterImageDefault) | **Get** /api/k8s/clusters/{cluster}/image/default | Get default container image
[**CreateNamespacedJob**](K8sAPI.md#CreateNamespacedJob) | **Post** /api/k8s/clusters/{cluster}/namespaces/{namespace}/jobs | Create K8s Job
[**CreateRbac**](K8sAPI.md#CreateRbac) | **Post** /api/k8s/rbac | Create Namespace, Roles and Role Bindings.
[**GetAllJobsDetails**](K8sAPI.md#GetAllJobsDetails) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/kind/{kind}/details | Get list of job/pytorchjob with basic details
[**GetDetailsOfJob**](K8sAPI.md#GetDetailsOfJob) | **Get** /api/k8s/jobs/{job_id}/details | Get basic details of a single job/pytorchjob
[**GetJobDescription**](K8sAPI.md#GetJobDescription) | **Get** /api/k8s/jobs/{job_id}/description | Get job description
[**GetJobPriorityTypes**](K8sAPI.md#GetJobPriorityTypes) | **Get** /api/k8s/job_priority_types | Get supported job priority types
[**GetJobPriorityTypesNew**](K8sAPI.md#GetJobPriorityTypesNew) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/job_priority_types | Get supported job priority types
[**GetLog**](K8sAPI.md#GetLog) | **Get** /api/k8s/clusters/{cluster}/namespaces/{namespace}/pods/{pod_name}/containers/{container}/logs | Get Pod container logs
[**GetLogLoki**](K8sAPI.md#GetLogLoki) | **Get** /api/k8s/jobs/{job_id}/logs | Get logs
[**GetStatusTypes**](K8sAPI.md#GetStatusTypes) | **Get** /api/k8s/status_types | Get supported status types
[**GetTemplateFromJob**](K8sAPI.md#GetTemplateFromJob) | **Get** /api/k8s/jobs/{job_id}/template | Get job template from existing job
[**GetUsersNamespaces**](K8sAPI.md#GetUsersNamespaces) | **Get** /api/k8s/clusters/{cluster}/namespaces | Get user&#39;s namespaces
[**ListClusters**](K8sAPI.md#ListClusters) | **Get** /api/k8s/clusters | List Clusters
[**ListClusters_0**](K8sAPI.md#ListClusters_0) | **Get** /api/k8s/clusters | List Clusters
[**ListJobKinds**](K8sAPI.md#ListJobKinds) | **Get** /api/k8s/clusters/{cluster}/jobs/kinds | List types of jobs supported
[**PutJobDescription**](K8sAPI.md#PutJobDescription) | **Put** /api/k8s/jobs/{job_id}/description | Add job description
[**TerminateJob**](K8sAPI.md#TerminateJob) | **Delete** /api/k8s/jobs/{job_id} | Delete a job



## ClusterImageDefault

> string ClusterImageDefault(ctx, cluster).Execute()

Get default container image



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.ClusterImageDefault(context.Background(), cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ClusterImageDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClusterImageDefault`: string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ClusterImageDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterImageDefaultRequest struct via the builder pattern


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


## CreateNamespacedJob

> CreateJobResp CreateNamespacedJob(ctx, cluster, namespace).CreateJobBody(createJobBody).Execute()

Create K8s Job



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
	createJobBody := *openapiclient.NewCreateJobBody("Name_example", "Namespace_example", openapiclient.Cluster("All"), "Image_example") // CreateJobBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.CreateNamespacedJob(context.Background(), cluster, namespace).CreateJobBody(createJobBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateNamespacedJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedJob`: CreateJobResp
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CreateNamespacedJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createJobBody** | [**CreateJobBody**](CreateJobBody.md) |  | 

### Return type

[**CreateJobResp**](CreateJobResp.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRbac

> interface{} CreateRbac(ctx).OnboardUser(onboardUser).Execute()

Create Namespace, Roles and Role Bindings.



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
	onboardUser := *openapiclient.NewOnboardUser("Namespace_example", "Cluster_example") // OnboardUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.CreateRbac(context.Background()).OnboardUser(onboardUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateRbac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRbac`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CreateRbac`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRbacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onboardUser** | [**OnboardUser**](OnboardUser.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllJobsDetails

> JobDetailsModel GetAllJobsDetails(ctx, cluster, namespace, kind).JobName(jobName).Owner(owner).Status(status).Start(start).End(end).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get list of job/pytorchjob with basic details



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
	kind := openapiclient.Kind("All") // Kind | 
	jobName := "jobName_example" // string |  (optional) (default to "")
	owner := "owner_example" // string |  (optional) (default to "")
	status := openapiclient.Status("All") // Status |  (optional)
	start := int32(56) // int32 |  (optional) (default to 0)
	end := int32(56) // int32 |  (optional) (default to 15)
	sortBy := openapiclient.SortBy("None") // SortBy |  (optional)
	sortOrder := openapiclient.SortOrder("Ascending") // SortOrder |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetAllJobsDetails(context.Background(), cluster, namespace, kind).JobName(jobName).Owner(owner).Status(status).Start(start).End(end).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetAllJobsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllJobsDetails`: JobDetailsModel
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetAllJobsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 
**kind** | [**Kind**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **jobName** | **string** |  | [default to &quot;&quot;]
 **owner** | **string** |  | [default to &quot;&quot;]
 **status** | [**Status**](Status.md) |  | 
 **start** | **int32** |  | [default to 0]
 **end** | **int32** |  | [default to 15]
 **sortBy** | [**SortBy**](SortBy.md) |  | 
 **sortOrder** | [**SortOrder**](SortOrder.md) |  | 

### Return type

[**JobDetailsModel**](JobDetailsModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailsOfJob

> JobItem GetDetailsOfJob(ctx, jobId).Execute()

Get basic details of a single job/pytorchjob



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetDetailsOfJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetDetailsOfJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailsOfJob`: JobItem
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetDetailsOfJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsOfJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobItem**](JobItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDescription

> JobDescription GetJobDescription(ctx, jobId).Execute()

Get job description



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetJobDescription(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetJobDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobDescription`: JobDescription
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetJobDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobDescription**](JobDescription.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobPriorityTypes

> []string GetJobPriorityTypes(ctx).Execute()

Get supported job priority types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetJobPriorityTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetJobPriorityTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobPriorityTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetJobPriorityTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobPriorityTypesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobPriorityTypesNew

> interface{} GetJobPriorityTypesNew(ctx, cluster, namespace).Execute()

Get supported job priority types



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
	resp, r, err := apiClient.K8sAPI.GetJobPriorityTypesNew(context.Background(), cluster, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetJobPriorityTypesNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobPriorityTypesNew`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetJobPriorityTypesNew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobPriorityTypesNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLog

> string GetLog(ctx, cluster, namespace, podName, container).Follow(follow).MaxLines(maxLines).Execute()

Get Pod container logs



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
	podName := "podName_example" // string | 
	container := "container_example" // string | 
	follow := true // bool |  (optional) (default to false)
	maxLines := int32(56) // int32 |  (optional) (default to 10000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetLog(context.Background(), cluster, namespace, podName, container).Follow(follow).MaxLines(maxLines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLog`: string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 
**namespace** | **string** |  | 
**podName** | **string** |  | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **follow** | **bool** |  | [default to false]
 **maxLines** | **int32** |  | [default to 10000]

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


## GetLogLoki

> string GetLogLoki(ctx, jobId).Execute()

Get logs



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetLogLoki(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetLogLoki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogLoki`: string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetLogLoki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogLokiRequest struct via the builder pattern


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


## GetStatusTypes

> []string GetStatusTypes(ctx).JobType(jobType).Execute()

Get supported status types



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
	jobType := openapiclient.Kind("All") // Kind | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetStatusTypes(context.Background()).JobType(jobType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetStatusTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetStatusTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobType** | [**Kind**](Kind.md) |  | 

### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateFromJob

> CreateJobBody GetTemplateFromJob(ctx, jobId).Execute()

Get job template from existing job



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetTemplateFromJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetTemplateFromJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateFromJob`: CreateJobBody
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetTemplateFromJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateFromJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateJobBody**](CreateJobBody.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersNamespaces

> UserNamespacesResp GetUsersNamespaces(ctx, cluster).Execute()

Get user's namespaces



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetUsersNamespaces(context.Background(), cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetUsersNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersNamespaces`: UserNamespacesResp
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetUsersNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserNamespacesResp**](UserNamespacesResp.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> []string ListClusters(ctx).Execute()

List Clusters



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.ListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: []string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters_0

> []string ListClusters_0(ctx).Execute()

List Clusters



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.ListClusters_0(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListClusters_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters_0`: []string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListClusters_0`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClusters_1Request struct via the builder pattern


### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobKinds

> []map[string]string ListJobKinds(ctx, cluster).Execute()

List types of jobs supported



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.ListJobKinds(context.Background(), cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListJobKinds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobKinds`: []map[string]string
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListJobKinds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**Cluster**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJobKindsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]string**](map.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutJobDescription

> JobDescriptionResp PutJobDescription(ctx, jobId).JobDescription(jobDescription).Execute()

Add job description



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
	jobId := "jobId_example" // string | 
	jobDescription := *openapiclient.NewJobDescription("Description_example") // JobDescription | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.PutJobDescription(context.Background(), jobId).JobDescription(jobDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.PutJobDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutJobDescription`: JobDescriptionResp
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.PutJobDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutJobDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobDescription** | [**JobDescription**](JobDescription.md) |  | 

### Return type

[**JobDescriptionResp**](JobDescriptionResp.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateJob

> interface{} TerminateJob(ctx, jobId).Execute()

Delete a job



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.TerminateJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.TerminateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateJob`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.TerminateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

