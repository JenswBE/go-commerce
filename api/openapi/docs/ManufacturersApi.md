# \ManufacturersApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetManufacturer**](ManufacturersApi.md#GetManufacturer) | **Get** /manufacturers/{id}/ | 
[**ListManufacturers**](ManufacturersApi.md#ListManufacturers) | **Get** /manufacturers/ | 



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

