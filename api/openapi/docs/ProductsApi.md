# \ProductsApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProduct**](ProductsApi.md#AddProduct) | **Post** /products/ | 
[**AddProductImages**](ProductsApi.md#AddProductImages) | **Post** /products/{id}/images/ | 
[**DeleteProduct**](ProductsApi.md#DeleteProduct) | **Delete** /products/{id}/ | 
[**DeleteProductImage**](ProductsApi.md#DeleteProductImage) | **Delete** /products/{id}/images/{image_id}/ | 
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /products/{id}/ | 
[**ListProductImages**](ProductsApi.md#ListProductImages) | **Get** /products/{id}/images/ | 
[**ListProducts**](ProductsApi.md#ListProducts) | **Get** /products/ | 
[**UpdateProduct**](ProductsApi.md#UpdateProduct) | **Put** /products/{id}/ | 
[**UpdateProductImage**](ProductsApi.md#UpdateProductImage) | **Put** /products/{id}/images/{image_id}/ | 



## AddProduct

> Product AddProduct(ctx).Product(product).Execute()





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
    product := *openapiclient.NewProduct("Gezichtsmasker", int64(1500)) // Product | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.AddProduct(context.Background()).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.AddProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.AddProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**Product**](Product.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProductImages

> ImageList AddProductImages(ctx, id).File(file).Img(img).Execute()





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
    file := []*os.File{"TODO"} // []*os.File | 
    img := []string{"300_200_FIT"} // []string | Comma separated list of ImageConfig. Check ImageConfig for exact format. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.AddProductImages(context.Background(), id).File(file).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.AddProductImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProductImages`: ImageList
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.AddProductImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProductImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **[]*os.File** |  | 
 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**ImageList**](ImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> DeleteProduct(ctx, id).Execute()





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
    resp, r, err := apiClient.ProductsApi.DeleteProduct(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.DeleteProduct``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


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


## DeleteProductImage

> DeleteProductImage(ctx, id, imageId).Execute()





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
    imageId := "imageId_example" // string | Image ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.DeleteProductImage(context.Background(), id, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.DeleteProductImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductImageRequest struct via the builder pattern


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


## ListProductImages

> ImageList ListProductImages(ctx, id).Img(img).Execute()





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
    resp, r, err := apiClient.ProductsApi.ListProductImages(context.Background(), id).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListProductImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductImages`: ImageList
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListProductImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**ImageList**](ImageList.md)

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


## UpdateProduct

> Product UpdateProduct(ctx, id).Product(product).Execute()





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
    product := *openapiclient.NewProduct("Gezichtsmasker", int64(1500)) // Product | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.UpdateProduct(context.Background(), id).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.UpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | [**Product**](Product.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductImage

> ImageList UpdateProductImage(ctx, id, imageId).Image(image).Execute()





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
    imageId := "imageId_example" // string | Image ID
    image := *openapiclient.NewImage("Id_example", "png", map[string]string{"key": "https://images.jensw.be/.../fill/300/200/.../anBn.png"}, int64(123)) // Image | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.UpdateProductImage(context.Background(), id, imageId).Image(image).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.UpdateProductImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProductImage`: ImageList
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.UpdateProductImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **image** | [**Image**](Image.md) |  | 

### Return type

[**ImageList**](ImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

