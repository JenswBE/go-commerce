# \ManufacturersApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddManufacturer**](ManufacturersApi.md#AddManufacturer) | **Post** /manufacturers/ | 
[**DeleteManufacturer**](ManufacturersApi.md#DeleteManufacturer) | **Delete** /manufacturers/{id}/ | 
[**DeleteManufacturerImage**](ManufacturersApi.md#DeleteManufacturerImage) | **Delete** /manufacturers/{id}/image/ | 
[**GetManufacturer**](ManufacturersApi.md#GetManufacturer) | **Get** /manufacturers/{id}/ | 
[**ListManufacturers**](ManufacturersApi.md#ListManufacturers) | **Get** /manufacturers/ | 
[**UpdateManufacturer**](ManufacturersApi.md#UpdateManufacturer) | **Put** /manufacturers/{id}/ | 
[**UpsertManufacturerImage**](ManufacturersApi.md#UpsertManufacturerImage) | **Put** /manufacturers/{id}/image/ | 



## AddManufacturer

> Manufacturer AddManufacturer(ctx).Manufacturer(manufacturer).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    manufacturer := *openapiclient.NewManufacturer("Bjoetiek Y") // Manufacturer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.AddManufacturer(context.Background()).Manufacturer(manufacturer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.AddManufacturer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddManufacturer`: Manufacturer
    fmt.Fprintf(os.Stdout, "Response from `ManufacturersApi.AddManufacturer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddManufacturerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manufacturer** | [**Manufacturer**](Manufacturer.md) |  | 

### Return type

[**Manufacturer**](Manufacturer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManufacturer

> DeleteManufacturer(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.DeleteManufacturer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.DeleteManufacturer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManufacturerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManufacturerImage

> DeleteManufacturerImage(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.DeleteManufacturerImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.DeleteManufacturerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManufacturerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManufacturer

> Manufacturer GetManufacturer(ctx, id).Img(img).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID
    img := []string{"300_200_FIT"} // []string | Comma separated list of ImageConfig. Check ImageConfig for exact format. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.GetManufacturer(context.Background(), id).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.GetManufacturer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManufacturer`: Manufacturer
    fmt.Fprintf(os.Stdout, "Response from `ManufacturersApi.GetManufacturer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManufacturerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**Manufacturer**](Manufacturer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManufacturers

> ManufacturerList ListManufacturers(ctx).Img(img).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    img := []string{"300_200_FIT"} // []string | Comma separated list of ImageConfig. Check ImageConfig for exact format. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.ListManufacturers(context.Background()).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.ListManufacturers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManufacturers`: ManufacturerList
    fmt.Fprintf(os.Stdout, "Response from `ManufacturersApi.ListManufacturers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListManufacturersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**ManufacturerList**](ManufacturerList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManufacturer

> Manufacturer UpdateManufacturer(ctx, id).Manufacturer(manufacturer).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID
    manufacturer := *openapiclient.NewManufacturer("Bjoetiek Y") // Manufacturer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.UpdateManufacturer(context.Background(), id).Manufacturer(manufacturer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.UpdateManufacturer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManufacturer`: Manufacturer
    fmt.Fprintf(os.Stdout, "Response from `ManufacturersApi.UpdateManufacturer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManufacturerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manufacturer** | [**Manufacturer**](Manufacturer.md) |  | 

### Return type

[**Manufacturer**](Manufacturer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertManufacturerImage

> Image UpsertManufacturerImage(ctx, id).File(file).Img(img).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID
    file := os.NewFile(1234, "some_file") // *os.File | 
    img := []string{"300_200_FIT"} // []string | Comma separated list of ImageConfig. Check ImageConfig for exact format. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManufacturersApi.UpsertManufacturerImage(context.Background(), id).File(file).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManufacturersApi.UpsertManufacturerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertManufacturerImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `ManufacturersApi.UpsertManufacturerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertManufacturerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

