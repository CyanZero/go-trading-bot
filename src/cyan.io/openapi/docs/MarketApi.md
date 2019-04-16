# \MarketApi

All URIs are relative to *https://engine.coss.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DpGet**](MarketApi.md#DpGet) | **Get** /dp/ | Retrieve market depth information (Order book) for given symbol
[**GetmarketsummariesGet**](MarketApi.md#GetmarketsummariesGet) | **Get** /getmarketsummaries | Provides information about market summaries for symbols.
[**HtGet**](MarketApi.md#HtGet) | **Get** /ht/ | Retrieve market information  for given symbol
[**MarketPriceGet**](MarketApi.md#MarketPriceGet) | **Get** /market-price/ | Retrieves market price information


# **DpGet**
> Depth DpGet(ctx, symbol)
Retrieve market depth information (Order book) for given symbol

Specific pair id for retrieving depth in query string e.g. ?symbol=ETH_BTC. This API does not require signing.<br/> Example URL: https://engine.coss.io/api/v1/dp?symbol=ETH_BTC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| pair id to retrieve market depth | 

### Return type

[**Depth**](Depth.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetmarketsummariesGet**
> MarketSummariesResponse GetmarketsummariesGet(ctx, )
Provides information about market summaries for symbols.

This is a public function and does not require signing. Retrieves market summaries for all symbols.<br>   URL Example : https://exchange.coss.io/api/getmarketsummaries

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MarketSummariesResponse**](MarketSummariesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HtGet**
> TradeHistoryResponse HtGet(ctx, symbol)
Retrieve market information  for given symbol

Specific pair id for retrieving market information stream in path e.g. ETH_BTC. This API does not require signing.<br/>  Example URL: https://engine.coss.io/api/v1/ht?symbol=ETH_BTC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| pair id to retrieve market information. | 

### Return type

[**TradeHistoryResponse**](TradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarketPriceGet**
> []MarketPrice MarketPriceGet(ctx, optional)
Retrieves market price information

This is a public function and does not require signing. Retrieves market price for all symbols if no symbol is provided as query string parameter. If a symbol is provided then retrieves market-price for the symbol   URL Example : https://trade.coss.io/c/api/v1//market-price?symbol=ETH_BTC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketPriceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketPriceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **optional.String**| Retrieves market-price for all symbols if symbol is not provide in parameter otherwise retrieves information for symbol provided in | 

### Return type

[**[]MarketPrice**](MarketPrice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

