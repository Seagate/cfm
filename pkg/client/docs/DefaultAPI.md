# \DefaultAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppliancesDeleteById**](DefaultAPI.md#AppliancesDeleteById) | **Delete** /cfm/v1/appliances/{applianceId} | 
[**AppliancesGet**](DefaultAPI.md#AppliancesGet) | **Get** /cfm/v1/appliances | 
[**AppliancesGetById**](DefaultAPI.md#AppliancesGetById) | **Get** /cfm/v1/appliances/{applianceId} | 
[**AppliancesPost**](DefaultAPI.md#AppliancesPost) | **Post** /cfm/v1/appliances | 
[**BladesAssignMemoryById**](DefaultAPI.md#BladesAssignMemoryById) | **Patch** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/memory/{memoryId} | 
[**BladesComposeMemory**](DefaultAPI.md#BladesComposeMemory) | **Post** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/memory | 
[**BladesComposeMemoryByResource**](DefaultAPI.md#BladesComposeMemoryByResource) | **Put** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/memory | 
[**BladesDeleteById**](DefaultAPI.md#BladesDeleteById) | **Delete** /cfm/v1/appliances/{applianceId}/blades/{bladeId} | 
[**BladesFreeMemoryById**](DefaultAPI.md#BladesFreeMemoryById) | **Delete** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/memory/{memoryId} | 
[**BladesGet**](DefaultAPI.md#BladesGet) | **Get** /cfm/v1/appliances/{applianceId}/blades | 
[**BladesGetById**](DefaultAPI.md#BladesGetById) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId} | 
[**BladesGetMemory**](DefaultAPI.md#BladesGetMemory) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/memory | 
[**BladesGetMemoryById**](DefaultAPI.md#BladesGetMemoryById) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/memory/{memoryId} | 
[**BladesGetPortById**](DefaultAPI.md#BladesGetPortById) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/ports/{portId} | 
[**BladesGetPorts**](DefaultAPI.md#BladesGetPorts) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/ports | 
[**BladesGetResourceById**](DefaultAPI.md#BladesGetResourceById) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/resources/{resourceId} | 
[**BladesGetResources**](DefaultAPI.md#BladesGetResources) | **Get** /cfm/v1/appliances/{applianceId}/blades/{bladeId}/resources | 
[**BladesPost**](DefaultAPI.md#BladesPost) | **Post** /cfm/v1/appliances/{applianceId}/blades | 
[**CfmGet**](DefaultAPI.md#CfmGet) | **Get** /cfm | 
[**CfmV1Get**](DefaultAPI.md#CfmV1Get) | **Get** /cfm/v1 | 
[**HostGetMemory**](DefaultAPI.md#HostGetMemory) | **Get** /cfm/v1/hosts/{hostId}/memory | 
[**HostsComposeMemory**](DefaultAPI.md#HostsComposeMemory) | **Post** /cfm/v1/hosts/{hostId}/memory | 
[**HostsDeleteById**](DefaultAPI.md#HostsDeleteById) | **Delete** /cfm/v1/hosts/{hostId} | 
[**HostsFreeMemoryById**](DefaultAPI.md#HostsFreeMemoryById) | **Delete** /cfm/v1/hosts/{hostId}/memory/{memoryId} | 
[**HostsGet**](DefaultAPI.md#HostsGet) | **Get** /cfm/v1/hosts | Get CXL Host information.
[**HostsGetById**](DefaultAPI.md#HostsGetById) | **Get** /cfm/v1/hosts/{hostId} | Get information for a single CXL Host.
[**HostsGetMemoryById**](DefaultAPI.md#HostsGetMemoryById) | **Get** /cfm/v1/hosts/{hostId}/memory/{memoryId} | 
[**HostsGetMemoryDeviceById**](DefaultAPI.md#HostsGetMemoryDeviceById) | **Get** /cfm/v1/hosts/{hostId}/memory-devices/{memoryDeviceId} | 
[**HostsGetMemoryDevices**](DefaultAPI.md#HostsGetMemoryDevices) | **Get** /cfm/v1/hosts/{hostId}/memory-devices | 
[**HostsGetPortById**](DefaultAPI.md#HostsGetPortById) | **Get** /cfm/v1/hosts/{hostId}/ports/{portId} | 
[**HostsGetPorts**](DefaultAPI.md#HostsGetPorts) | **Get** /cfm/v1/hosts/{hostId}/ports | 
[**HostsPost**](DefaultAPI.md#HostsPost) | **Post** /cfm/v1/hosts | Add a CXL host to be managed by CFM.
[**RootGet**](DefaultAPI.md#RootGet) | **Get** / | 



## AppliancesDeleteById

> Appliance AppliancesDeleteById(ctx, applianceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AppliancesDeleteById(context.Background(), applianceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppliancesDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppliancesDeleteById`: Appliance
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppliancesDeleteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppliancesDeleteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Appliance**](Appliance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppliancesGet

> Collection AppliancesGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AppliancesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppliancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppliancesGet`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppliancesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppliancesGetRequest struct via the builder pattern


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppliancesGetById

> Appliance AppliancesGetById(ctx, applianceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AppliancesGetById(context.Background(), applianceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppliancesGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppliancesGetById`: Appliance
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppliancesGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppliancesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Appliance**](Appliance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppliancesPost

> Appliance AppliancesPost(ctx).Credentials(credentials).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    credentials := *openapiclient.NewCredentials("User0", "User0password!", "127.0.0.1", int32(80)) // Credentials |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AppliancesPost(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppliancesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppliancesPost`: Appliance
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppliancesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppliancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) |  | 

### Return type

[**Appliance**](Appliance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesAssignMemoryById

> MemoryRegion BladesAssignMemoryById(ctx, applianceId, bladeId, memoryId).AssignMemoryRequest(assignMemoryRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    memoryId := "memoryId_example" // string | A unique identifier for a Memory Region
    assignMemoryRequest := *openapiclient.NewAssignMemoryRequest("P1", "Operation_example") // AssignMemoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesAssignMemoryById(context.Background(), applianceId, bladeId, memoryId).AssignMemoryRequest(assignMemoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesAssignMemoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesAssignMemoryById`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesAssignMemoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 
**memoryId** | **string** | A unique identifier for a Memory Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesAssignMemoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **assignMemoryRequest** | [**AssignMemoryRequest**](AssignMemoryRequest.md) |  | 

### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesComposeMemory

> MemoryRegion BladesComposeMemory(ctx, applianceId, bladeId).ComposeMemoryRequest(composeMemoryRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    composeMemoryRequest := *openapiclient.NewComposeMemoryRequest(int32(123), openapiclient.qos(1)) // ComposeMemoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesComposeMemory(context.Background(), applianceId, bladeId).ComposeMemoryRequest(composeMemoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesComposeMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesComposeMemory`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesComposeMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesComposeMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **composeMemoryRequest** | [**ComposeMemoryRequest**](ComposeMemoryRequest.md) |  | 

### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesComposeMemoryByResource

> MemoryRegion BladesComposeMemoryByResource(ctx, applianceId, bladeId).ComposeMemoryByResourceRequest(composeMemoryByResourceRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    composeMemoryByResourceRequest := *openapiclient.NewComposeMemoryByResourceRequest([]string{"device-1234"}) // ComposeMemoryByResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesComposeMemoryByResource(context.Background(), applianceId, bladeId).ComposeMemoryByResourceRequest(composeMemoryByResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesComposeMemoryByResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesComposeMemoryByResource`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesComposeMemoryByResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesComposeMemoryByResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **composeMemoryByResourceRequest** | [**ComposeMemoryByResourceRequest**](ComposeMemoryByResourceRequest.md) |  | 

### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesDeleteById

> Blade BladesDeleteById(ctx, applianceId, bladeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesDeleteById(context.Background(), applianceId, bladeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesDeleteById`: Blade
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesDeleteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesDeleteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Blade**](Blade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesFreeMemoryById

> MemoryRegion BladesFreeMemoryById(ctx, applianceId, bladeId, memoryId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    memoryId := "memoryId_example" // string | A unique identifier for a Memory Region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesFreeMemoryById(context.Background(), applianceId, bladeId, memoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesFreeMemoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesFreeMemoryById`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesFreeMemoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 
**memoryId** | **string** | A unique identifier for a Memory Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesFreeMemoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGet

> Collection BladesGet(ctx, applianceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGet(context.Background(), applianceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGet`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetById

> Blade BladesGetById(ctx, applianceId, bladeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetById(context.Background(), applianceId, bladeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetById`: Blade
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Blade**](Blade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetMemory

> Collection BladesGetMemory(ctx, applianceId, bladeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetMemory(context.Background(), applianceId, bladeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetMemory`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetMemoryById

> MemoryRegion BladesGetMemoryById(ctx, applianceId, bladeId, memoryId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    memoryId := "memoryId_example" // string | A unique identifier for a Memory Region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetMemoryById(context.Background(), applianceId, bladeId, memoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetMemoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetMemoryById`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetMemoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 
**memoryId** | **string** | A unique identifier for a Memory Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetMemoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetPortById

> PortInformation BladesGetPortById(ctx, applianceId, bladeId, portId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    portId := "portId_example" // string | A unique identifier for a Port

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetPortById(context.Background(), applianceId, bladeId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetPortById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetPortById`: PortInformation
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetPortById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 
**portId** | **string** | A unique identifier for a Port | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetPortByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PortInformation**](PortInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetPorts

> Collection BladesGetPorts(ctx, applianceId, bladeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetPorts(context.Background(), applianceId, bladeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetPorts`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetResourceById

> MemoryResourceBlock BladesGetResourceById(ctx, applianceId, bladeId, resourceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade
    resourceId := "resourceId_example" // string | A unique identifier for a Memory Resource Block

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetResourceById(context.Background(), applianceId, bladeId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetResourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetResourceById`: MemoryResourceBlock
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetResourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 
**resourceId** | **string** | A unique identifier for a Memory Resource Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetResourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MemoryResourceBlock**](MemoryResourceBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesGetResources

> Collection BladesGetResources(ctx, applianceId, bladeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    bladeId := "bladeId_example" // string | A unique identifier for a Memory Blade

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesGetResources(context.Background(), applianceId, bladeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesGetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesGetResources`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesGetResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 
**bladeId** | **string** | A unique identifier for a Memory Blade | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BladesPost

> Blade BladesPost(ctx, applianceId).Credentials(credentials).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    applianceId := "applianceId_example" // string | A unique identifier for a Memory Appliance
    credentials := *openapiclient.NewCredentials("User0", "User0password!", "127.0.0.1", int32(80)) // Credentials | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BladesPost(context.Background(), applianceId).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BladesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BladesPost`: Blade
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BladesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applianceId** | **string** | A unique identifier for a Memory Appliance | 

### Other Parameters

Other parameters are passed through a pointer to a apiBladesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentials** | [**Credentials**](Credentials.md) |  | 

### Return type

[**Blade**](Blade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfmGet

> string CfmGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CfmGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CfmGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CfmGet`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CfmGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfmGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfmV1Get

> ServiceInformation CfmV1Get(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CfmV1Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CfmV1Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CfmV1Get`: ServiceInformation
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CfmV1Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfmV1GetRequest struct via the builder pattern


### Return type

[**ServiceInformation**](ServiceInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostGetMemory

> Collection HostGetMemory(ctx, hostId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostGetMemory(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostGetMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostGetMemory`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostGetMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostGetMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsComposeMemory

> MemoryRegion HostsComposeMemory(ctx, hostId).ComposeMemoryRequest(composeMemoryRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host
    composeMemoryRequest := *openapiclient.NewComposeMemoryRequest(int32(123), openapiclient.qos(1)) // ComposeMemoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsComposeMemory(context.Background(), hostId).ComposeMemoryRequest(composeMemoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsComposeMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsComposeMemory`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsComposeMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsComposeMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **composeMemoryRequest** | [**ComposeMemoryRequest**](ComposeMemoryRequest.md) |  | 

### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsDeleteById

> Host HostsDeleteById(ctx, hostId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsDeleteById(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsDeleteById`: Host
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsDeleteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsDeleteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsFreeMemoryById

> MemoryRegion HostsFreeMemoryById(ctx, hostId, memoryId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host
    memoryId := "memoryId_example" // string | A unique identifier for a Memory Region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsFreeMemoryById(context.Background(), hostId, memoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsFreeMemoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsFreeMemoryById`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsFreeMemoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 
**memoryId** | **string** | A unique identifier for a Memory Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsFreeMemoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGet

> Collection HostsGet(ctx).Execute()

Get CXL Host information.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGet`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetRequest struct via the builder pattern


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGetById

> Host HostsGetById(ctx, hostId).Execute()

Get information for a single CXL Host.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGetById(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGetById`: Host
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGetMemoryById

> MemoryRegion HostsGetMemoryById(ctx, hostId, memoryId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host
    memoryId := "memoryId_example" // string | A unique identifier for a Memory Region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGetMemoryById(context.Background(), hostId, memoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGetMemoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGetMemoryById`: MemoryRegion
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGetMemoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 
**memoryId** | **string** | A unique identifier for a Memory Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetMemoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemoryRegion**](MemoryRegion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGetMemoryDeviceById

> MemoryDeviceInformation HostsGetMemoryDeviceById(ctx, hostId, memoryDeviceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host
    memoryDeviceId := "memoryDeviceId_example" // string | A unique identifier for a memory device

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGetMemoryDeviceById(context.Background(), hostId, memoryDeviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGetMemoryDeviceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGetMemoryDeviceById`: MemoryDeviceInformation
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGetMemoryDeviceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 
**memoryDeviceId** | **string** | A unique identifier for a memory device | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetMemoryDeviceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemoryDeviceInformation**](MemoryDeviceInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGetMemoryDevices

> Collection HostsGetMemoryDevices(ctx, hostId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGetMemoryDevices(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGetMemoryDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGetMemoryDevices`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGetMemoryDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetMemoryDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGetPortById

> PortInformation HostsGetPortById(ctx, hostId, portId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host
    portId := "portId_example" // string | A unique identifier for a Port

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGetPortById(context.Background(), hostId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGetPortById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGetPortById`: PortInformation
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGetPortById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 
**portId** | **string** | A unique identifier for a Port | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetPortByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PortInformation**](PortInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGetPorts

> Collection HostsGetPorts(ctx, hostId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    hostId := "hostId_example" // string | A unique identifier for a CXL Host

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsGetPorts(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGetPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsGetPorts`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGetPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | A unique identifier for a CXL Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsPost

> Host HostsPost(ctx).Credentials(credentials).Execute()

Add a CXL host to be managed by CFM.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    credentials := *openapiclient.NewCredentials("User0", "User0password!", "127.0.0.1", int32(80)) // Credentials | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.HostsPost(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostsPost`: Host
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) |  | 

### Return type

[**Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootGet

> string RootGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.RootGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RootGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootGet`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

