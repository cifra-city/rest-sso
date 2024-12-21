# \EmailAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AproveOperationPatch**](EmailAPI.md#AproveOperationPatch) | **Patch** /aprove-operation | Approve operation



## AproveOperationPatch

> AproveOperationPatch200Response AproveOperationPatch(ctx).ApproveOperationReq(approveOperationReq).Execute()

Approve operation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	approveOperationReq := *openapiclient.NewApproveOperationReq(*openapiclient.NewApproveOperationReqData("Type_example", *openapiclient.NewApproveOperationReqDataAttributes("Email_example", "Code_example", "reset_password"))) // ApproveOperationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.AproveOperationPatch(context.Background()).ApproveOperationReq(approveOperationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.AproveOperationPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AproveOperationPatch`: AproveOperationPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.AproveOperationPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAproveOperationPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approveOperationReq** | [**ApproveOperationReq**](ApproveOperationReq.md) |  | 

### Return type

[**AproveOperationPatch200Response**](AproveOperationPatch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/json, application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

