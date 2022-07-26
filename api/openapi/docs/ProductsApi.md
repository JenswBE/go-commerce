# \ProductsApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /products/{id}/ | 
[**ListProducts**](ProductsApi.md#ListProducts) | **Get** /products/ | 



## GetProduct

> ResolvedProduct GetProduct(ctx, id).Img(img).Resolve(resolve).Execute()





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
    resolve := true // bool | The returned object should include related objects. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.GetProduct(context.Background(), id).Img(img).Resolve(resolve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: ResolvedProduct
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 
 **resolve** | **bool** | The returned object should include related objects. | [default to false]

### Return type

[**ResolvedProduct**](ResolvedProduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ProductList ListProducts(ctx).Img(img).Execute()





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
    resp, r, err := apiClient.ProductsApi.ListProducts(context.Background()).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ProductList
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**ProductList**](ProductList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

