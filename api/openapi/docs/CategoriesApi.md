# \CategoriesApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCategory**](CategoriesApi.md#AddCategory) | **Post** /categories/ | 
[**DeleteCategory**](CategoriesApi.md#DeleteCategory) | **Delete** /categories/{id}/ | 
[**DeleteCategoryImage**](CategoriesApi.md#DeleteCategoryImage) | **Delete** /categories/{id}/image/ | 
[**GetCategory**](CategoriesApi.md#GetCategory) | **Get** /categories/{id}/ | 
[**ListCategories**](CategoriesApi.md#ListCategories) | **Get** /categories/ | 
[**UpdateCategory**](CategoriesApi.md#UpdateCategory) | **Put** /categories/{id}/ | 
[**UpsertCategoryImage**](CategoriesApi.md#UpsertCategoryImage) | **Put** /categories/{id}/image/ | 



## AddCategory

> Category AddCategory(ctx).Category(category).Execute()





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
    category := *openapiclient.NewCategory("Makeup & Cosmetica", int64(123)) // Category | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesApi.AddCategory(context.Background()).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.AddCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.AddCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**Category**](Category.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> DeleteCategory(ctx, id).Execute()





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
    resp, r, err := apiClient.CategoriesApi.DeleteCategory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.DeleteCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## DeleteCategoryImage

> DeleteCategoryImage(ctx, id).Execute()





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
    resp, r, err := apiClient.CategoriesApi.DeleteCategoryImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.DeleteCategoryImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCategoryImageRequest struct via the builder pattern


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


## GetCategory

> Category GetCategory(ctx, id).Img(img).Execute()





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
    resp, r, err := apiClient.CategoriesApi.GetCategory(context.Background(), id).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategories

> CategoryList ListCategories(ctx).Img(img).Execute()





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
    resp, r, err := apiClient.CategoriesApi.ListCategories(context.Background()).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: CategoryList
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.ListCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **img** | **[]string** | Comma separated list of ImageConfig. Check ImageConfig for exact format. | 

### Return type

[**CategoryList**](CategoryList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCategory

> Category UpdateCategory(ctx, id).Category(category).Execute()





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
    category := *openapiclient.NewCategory("Makeup & Cosmetica", int64(123)) // Category | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesApi.UpdateCategory(context.Background(), id).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | [**Category**](Category.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCategoryImage

> Image UpsertCategoryImage(ctx, id).File(file).Img(img).Execute()





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
    resp, r, err := apiClient.CategoriesApi.UpsertCategoryImage(context.Background(), id).File(file).Img(img).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.UpsertCategoryImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertCategoryImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.UpsertCategoryImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCategoryImageRequest struct via the builder pattern


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

