# \AccountApi

All URIs are relative to *https://engine.coss.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountBalancesHMACSHA256Get**](AccountApi.md#AccountBalancesHMACSHA256Get) | **Get** /account/balances (HMAC SHA256) | Retrieves account balances information.
[**AccountDetailsHMACSHA256Get**](AccountApi.md#AccountDetailsHMACSHA256Get) | **Get** /account/details (HMAC SHA256) | Retrieves account details information.


# **AccountBalancesHMACSHA256Get**
> []AccountBalance AccountBalancesHMACSHA256Get(ctx, timestamp, optional)
Retrieves account balances information.

This is a signed function. User must provide the public api key in Authorization header and signed payload in Signature header. On a Linux machine following command will generate the signed pay load echo -n 'recvWindow=5000×tamp=1540203005798’ | openssl dgst -sha256 -hmac ‘replace this by your private key value’ Example URL: https://trade.coss.io/c/api/v1/account/balances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **int64**| Mandatory field for retrieving account balances in query string ?recvWindow&#x3D;5000×tamp&#x3D;1540203005798. The future release of API generate use timestamp provided by user in conjunction with recvWindow parameter provided by the user to check if request has reached in time. If the server timestamp is later than sun of value of timestamp and recvWindow by the user  then request will be rejected. | 
 **optional** | ***AccountBalancesHMACSHA256GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountBalancesHMACSHA256GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recvWindow** | **optional.Int32**| Optional field for retrieving account balances in query string ?recvWindow&#x3D;5000×tamp&#x3D;1540203005798.  | 

### Return type

[**[]AccountBalance**](AccountBalance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountDetailsHMACSHA256Get**
> AccountDetails AccountDetailsHMACSHA256Get(ctx, timestamp, optional)
Retrieves account details information.

This is a signed function. User must provide the public api key in Authorisation header and signed payload in Signature header. On a Linux machine following command will generate the signed pay load echo -n ''recvWindow=5000×tamp=1540203005798’’ | openssl dgst -sha256 -hmac ‘replace this by your private key value’ Example URL: https://trade.coss.io/c/api/v1/account/details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **int64**| Mandatory field for retrieving account details in query string ?recvWindow&#x3D;5000×tamp&#x3D;1540203005798.If the server timestamp is later than sum of value of timestamp and recvWindow provided by the user  then request will be rejected. | 
 **optional** | ***AccountDetailsHMACSHA256GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountDetailsHMACSHA256GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recvWindow** | **optional.Int32**| Optional field for retrieving account balances in query string ?timestamp&#x3D;1540203005798&amp;recvWindow&#x3D;5000.  | 

### Return type

[**AccountDetails**](AccountDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

