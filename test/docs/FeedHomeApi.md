# \FeedHomeApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFeedHome**](FeedHomeApi.md#FetchFeedHome) | **Get** /feed-home | Fetch feed home



## FetchFeedHome

> []FeedHome FetchFeedHome(ctx).Execute()

Fetch feed home

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedHomeApi.FetchFeedHome(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedHomeApi.FetchFeedHome``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFeedHome`: []FeedHome
    fmt.Fprintf(os.Stdout, "Response from `FeedHomeApi.FetchFeedHome`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeedHomeRequest struct via the builder pattern


### Return type

[**[]FeedHome**](FeedHome.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

