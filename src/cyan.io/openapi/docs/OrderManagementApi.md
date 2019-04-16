# \OrderManagementApi

All URIs are relative to *https://engine.coss.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderAddHMACSHA256Post**](OrderManagementApi.md#OrderAddHMACSHA256Post) | **Post** /order/add/ (HMAC SHA256) | Place a new order
[**OrderCancelHMACSHA256Delete**](OrderManagementApi.md#OrderCancelHMACSHA256Delete) | **Delete** /order/cancel/ (HMAC SHA256) | Cancel the open order
[**OrderDetailsHMACSHA256Post**](OrderManagementApi.md#OrderDetailsHMACSHA256Post) | **Post** /order/details/ (HMAC SHA256) | Get order detail for specific order.
[**OrderListAllHMACSHA256Post**](OrderManagementApi.md#OrderListAllHMACSHA256Post) | **Post** /order/list/all (HMAC SHA256) | Get the list of all orders for user.
[**OrderListCompletedHMACSHA256Post**](OrderManagementApi.md#OrderListCompletedHMACSHA256Post) | **Post** /order/list/completed (HMAC SHA256) | Get the list of completed orders for user.
[**OrderListOpenHMACSHA256Post**](OrderManagementApi.md#OrderListOpenHMACSHA256Post) | **Post** /order/list/open (HMAC SHA256) | Get the list of open orders for user.
[**OrderTradeDetailHMACSHA256Post**](OrderManagementApi.md#OrderTradeDetailHMACSHA256Post) | **Post** /order/trade-detail (HMAC SHA256) | Get order&#39;s trade details.


# **OrderAddHMACSHA256Post**
> OrderResponse OrderAddHMACSHA256Post(ctx, order)
Place a new order

Place a new order for order side(BUY/SELL) and order type (market/limit). This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/add

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | [**Order**](Order.md)| The new order to be created | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCancelHMACSHA256Delete**
> CancelOrderResponse OrderCancelHMACSHA256Delete(ctx, cancelOrder)
Cancel the open order

Cancel an open order.  This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/cancel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cancelOrder** | [**CancelOrder**](CancelOrder.md)| The order to be cancelled | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderDetailsHMACSHA256Post**
> OrderResponse OrderDetailsHMACSHA256Post(ctx, orderDetail)
Get order detail for specific order.

Get order details for specific order.  This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderDetail** | [**OrderDetail**](OrderDetail.md)| The order to retrieve the details | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderListAllHMACSHA256Post**
> []OrderResponse OrderListAllHMACSHA256Post(ctx, orderAllRequest)
Get the list of all orders for user.

Get the list of all orders for user. URL Example: https://trade.coss.io/c/api/v1/order/list/all

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderAllRequest** | [**OrderAllRequest**](OrderAllRequest.md)| The order request for retrieving all orders. This API requires signing. | 

### Return type

[**[]OrderResponse**](OrderResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderListCompletedHMACSHA256Post**
> OrderListResponse OrderListCompletedHMACSHA256Post(ctx, orderListRequest)
Get the list of completed orders for user.

Get completed orders for specific symbol where order status is  FILLED, PARTIAL_FILL, or CANCELED URL Example: https://trade.coss.io/c/api/v1/order/list/completed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderListRequest** | [**OrderListRequest**](OrderListRequest.md)| The order request for retrieving open orders. This API requires signing. | 

### Return type

[**OrderListResponse**](OrderListResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderListOpenHMACSHA256Post**
> OrderListResponse OrderListOpenHMACSHA256Post(ctx, orderListRequest)
Get the list of open orders for user.

Get current open orders for specific symbol OPEN and CANCELING.  This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/list/open

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderListRequest** | [**OrderListRequest**](OrderListRequest.md)| The order request for retrieving open orders. This API requires signing. | 

### Return type

[**OrderListResponse**](OrderListResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderTradeDetailHMACSHA256Post**
> []TradeDetail OrderTradeDetailHMACSHA256Post(ctx, tradeDetailRequest)
Get order's trade details.

Get trade details for an order. API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/trade-detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tradeDetailRequest** | [**TradeDetailRequest**](TradeDetailRequest.md)| The order to retrieve the trade details | 

### Return type

[**[]TradeDetail**](TradeDetail.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [APISignatureHeader](../README.md#APISignatureHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

