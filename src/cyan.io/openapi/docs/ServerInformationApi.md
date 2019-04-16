# \ServerInformationApi

All URIs are relative to *https://engine.coss.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PingGet**](ServerInformationApi.md#PingGet) | **Get** /ping | Test connectivity to API
[**TimeGet**](ServerInformationApi.md#TimeGet) | **Get** /time | Test connectivity to API and get Server time


# **PingGet**
> ServerStatusInfo PingGet(ctx, )
Test connectivity to API

Test connectivity to API.  Example URL: https://trade.coss.io/c/api/v1/ping

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerStatusInfo**](ServerStatusInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TimeGet**
> ServerTimeInfo TimeGet(ctx, )
Test connectivity to API and get Server time

Test connectivity to API and get server time.  Example URL: https://trade.coss.io/c/api/v1/time

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerTimeInfo**](ServerTimeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

