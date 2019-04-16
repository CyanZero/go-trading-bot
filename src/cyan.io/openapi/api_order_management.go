/*
 * COSS public api for trading
 *
 * This is coss public api to facilitate secure trading for registered users. You can find out more about how to enable api and obtain api key for trading under Profile section in coss.io [https://coss.io](https://coss.io) . For APIs which require signing user must provide following headers in request  Authorisation - The public key.  Signature - The HMAC256 hashed payload using private key.    Please refer to [community trading wrappers](https://github.com/coss-exchange) for sample codes       <h1>API news</h1>       <ul>     <li>        <strong><u>January, 16 2019</u></strong> : Document update for <strong>rate limits and order cancellation limits</strong> information         <ul>          <li> The api <strong>https://trade.coss.io/c/api/v1/exchange-info</strong> rate_limits contains objects related to the exchange’s REQUESTS rate limits. <br/><br/>           </li>          <li> A 429 will be returned when rate limit is violated.  <br/><br/>           </li>         <li>             Your account may get <strong>blocked when you place and cancel orders too frequently </strong>as explained below:          <br/><br/>          <font color=\"red\"><strong>          The frequency of your order placement and cancellation was too high. If you placed and cancelled an order within 10 seconds then it is counted as a violation. Five continuous violations would result your account being blocked.   </strong></font> <br/><br/>         <strong> First occurrence</strong>: Your transactions will be blocked for 5 minutes.<br>          <strong>Second occurrence</strong>: You will be blocked for 1 hour.<br>          <strong>Third occurrence</strong>: Your account will be locked for 24 hours.<br/>          <strong>Fourth occurrence</strong>: Your account will be locked for a longer duration. You may need to contact support to unlock your account.         </li>         <br/>       <li>       <strong><u>December, 19th 2018</u></strong> : Changes regarding COSS 1.2         <ul>          <li> A new api has been added to provide trade details for an order <br/>            https://trade.coss.io/c/api/v1/order/trade-detail            <br/>         </li>        <li> For API server status please use following <br/>          https://trade.coss.io/c/api/v1/ping            <br/>         </li>        <li> For API server  time please use following <br/>          https://trade.coss.io/c/api/v1/time            <br/>         </li>        <li> For retrieving market summaries please use          https://exchange.coss.io/api/getmarketsummaries            [1 unit] <br/>         Please note that this api is used by external data providers and the symbol format is different from api for users.            </li>         <li>           Updated Document to reflect paths specific to hosts.         </li>         <li>          Your account may get blocked when you place and cancel orders too frequently as explained below:          <br/><br/>          The frequency of your order placement and cancellation was too high.<br/>          First occurrence: Your transactions will be blocked for 5 minutes.<br>          Second occurrence: You will be blocked for 1 hour.<br>          Third occurrence: Your account will be locked for 24 hours.<br/>          Fourth occurrence: Your account will be locked for a longer duration. You may need to contact support to unlock your account.         </li>                 </ul>   <ul>     <li>       <strong><u>December, 7th 2018</u></strong> : Changes regarding COSS 1.2       <ul>         <li>       Rate Limits have been an issue. To make it fairer and easier to deal with burst data, we’re decreasing the API throttling. We’re assigning a usage limit of 1000 units per “MINUTE”. Different API calls have different weights, heavier calls use more units. We’ve added the weight below in []’s         </li>         <li>           <code>https://api.coss.io/v1/</code> will be depreciated and split into 2 two domains:           <ol>             <li>               <code>https://engine.coss.io/api/v1/</code> - this will handle all our pricing streams               <ul>                 <li>                   <strong>GET</strong> <u>/dp</u> - for depth  [1 unit]                 </li>                 <li>                   <strong>GET</strong> <u>/ht</u> - for trade history [1 unit]                 </li>               </ul>             </li>             <li>               <code>https://trade.coss.io/c/api/v1/</code> - this will handle all account and order requests               <ul>                 <li>                   <strong>POST</strong> <u>/order/add</u> [1 unit]                 </li>                 <li>                   <strong>DELETE</strong> <u>/order/cancel</u> [1 unit]                 </li>                 <li>                   <strong>POST</strong> <u>/order/details</u> [1 unit]                 </li>                 <li>                   <strong>POST</strong> <u>/order/list/open</u> [1 unit]                 </li>                 <li>                   <strong>POST</strong> <u>/order/list/completed</u> [1 unit]                 </li>                 <li>                   <strong>POST</strong> <u>/order/list/all</u> [5 units]                 </li>                 <li>                   <strong>GET</strong> <u>/account/balances</u> [5 units]                 </li>                 <li>                   <strong>GET</strong> <u>/account/details</u> [5 units]                  </li>                 <li>                   <strong>GET</strong> <u>/market-price</u> [1 unit]                 </li>                 <li>                   <strong>GET</strong> <u>/exchange-info</u> [1 unit]                 </li>               </ul>             </li>           </ol>         </li>         <li>           We’ve added price precision on order price and size (similar to other exchanges)           e.g. for ETH_BTC price precision 5 order size precision 3 (full list below)           The precision per pair is available from the /exchange-info API call         </li>         <li>           We’ve added a new websocket price feed for order book depth and trades, it’s efficient as 0 units to use           <ul>             <li>                <strong>GET</strong> <u>wss://engine.coss.io/ws/v1/ht/{symbol}</u>               <blockquote>                 {<br />                   &emsp;\"c\" : 1544064724447,   // Event time<br />                   &emsp;\"e\" : \"history_trade\", // Event type<br />                   &emsp;\"k\" : 461999,       // ID<br />                   &emsp;\"m\" : false,        // Buyer Made Order (buy order)<br />                   &emsp;\"p\" : \"0.02771000\", // Price<br />                   &emsp;\"q\" : \"0.37800000\", // Quantity (Size)<br />                   &emsp;\"s\" : \"ETH_BTC\",    // Symbol<br />                   &emsp;\"t\" : 1544064724247 // Trade Time<br />                 }               </blockquote>             </li>             <li>               <strong>GET</strong> <u>wss://engine.coss.io/ws/v1/dp/{symbol}</u>               <blockquote>                 {<br />                   &emsp;\"a\" : [ // asks<br />                   &emsp;&emsp;[<br />                   &emsp;&emsp;&emsp;&emsp;\"0.02773000\",   // price<br />                   &emsp;&emsp;&emsp;&emsp;\"0.67800000\"    // size<br />                   &emsp;&emsp;]<br />                   &emsp;],<br />                   &emsp;\"b\" : [ // bids<br />                   &emsp;&emsp;[<br />                   &emsp;&emsp;&emsp;\"0.02769000\",   // price<br />                   &emsp;&emsp;&emsp;\"0.97800000\"    // size<br />                   &emsp;&emsp;]<br />                   &emsp;]<br />                   &emsp;\"e\" : \"depthUpdate\" // event_type<br />                   &emsp;\"s\" : \"ETH_BTC\"   // Symbol<br />                   &emsp;\"t\" : 1544064724247 // Time<br />                 }<br />               </blockquote>             </li>           </ul>         </li>         <li>           A list of the symbols and the precision (some samples)           <table>             <thead>               <tr>                 <th>Pairs</th>                 <th>Token</th>                 <th>Base Pair</th>                 <th>Order Amount Limit Decimal</th>                 <th>Order Price Limit Decimal</th>               </tr>             </thead>             <tbody>               <tr>                 <td>COSS_BTC</td>                 <td>COSS</td>                 <td>BTC</td>                 <td>2</td>                 <td>6</td>               </tr>               <tr>                 <td>COSS_ETH</td>                 <td>COSS</td>                 <td>ETH</td>                 <td>2</td>                 <td>6</td>               </tr>               <tr>                 <td>KIN_BTC</td>                 <td>KIN</td>                 <td>BTC</td>                 <td>0</td>                 <td>8</td>               </tr>               <tr>                 <td>KIN_ETH</td>                 <td>KIN</td>                 <td>ETH</td>                 <td>1</td>                 <td>7</td>               </tr>               <tr>                 <td>NEO_BTC</td>                 <td>NEO</td>                 <td>BTC</td>                 <td>3</td>                 <td>5</td>               </tr>               <tr>                 <td>NEO_ETH</td>                 <td>NEO</td>                 <td>ETH</td>                 <td>3</td>                 <td>5</td>               </tr>             </tbody>           </table>         </li>       </ul>     </li>     <li>       <strong><u>November, 16th 2018</u></strong> : Updated API Document</li>     <li>       <strong><u>October, 31st 2018</u></strong> : To match industry convention, changed response code for Order creation from <code>202</code> to <code>200</code></li>     <li>       <strong><u>October, 30th 2018</u></strong> : Reduced Throttling to 1 request per second</li>   </ul> </p> <p>   <h3>Swagger Specs</h3>   <ul>     <li>You can also copy &amp; paste       <a href=\"https://s3-ap-southeast-1.amazonaws.com/coss-dev-s3-static-site-assets/api-gateway/apispecs.json\">coss-api-json</a> into an online swagger editor at       <a href=\"http://swagger.io\">swagger.io</a>     </li> </p> <p>   </ul> </p> <p>   <h3>Sample Code</h3>   <ul>     <li>Please refer to       <a href=\"https://github.com/coss-exchange\">community trading wrapper</a> for sample codes</li>   </ul> </p> <p>   <h3>Note: Known Issues</h3>   <h3>SIGNED GET Method</h3>   <ul>     <li>Please provide the query string (Signed GET methods for account/balances and account/details) in alphabetical order. The cloud provider sends the query strings of parameter names (arranged alphabetically) so the signed pay load may not match at server end. As a work around, please provide parameters in the  following format:       recvWindow=5000&timestamp=12345678</li>   </ul> </p> <h3>Orders</h3> <ul>   <li>After creation of a new order, a response code <code>200</code> is sent when order created successfully.</li>   <li>stop_price in order request and response is not used. The field is for future release.</li>   <li>Completed orders not returning orders which were created before public API release.</li>   <li>Market Orders are not currently supported. Please use Limit Orders.</li> </ul> <p>   <br/> </p> <p>   <h3>General</h3>   <ul>     <li>Timestamp and recvWindow are for future release and currently not used for request timeout, however for signed GET requests the signed payload must be provided: recvWindow=5000&amp;timestamp=12345678</li>     <li>Throttling is now set to 1 request every second, this will change over the coming weeks</li>   </ul> </p> <p>   <br/> </p>
 *
 * API version: 1.2
 * Contact: support@coss.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type OrderManagementApiService service

/*
OrderManagementApiService Place a new order
Place a new order for order side(BUY/SELL) and order type (market/limit). This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/add
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param order The new order to be created
@return OrderResponse
*/
func (a *OrderManagementApiService) OrderAddHMACSHA256Post(ctx context.Context, order Order) (OrderResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OrderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/add"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}
	localVarHeaderParams["X-Requested-With"] = "XMLHttpRequest"

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &order
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OrderResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			log.Printf("VarBody %v", v)
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrderManagementApiService Cancel the open order
Cancel an open order.  This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/cancel
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cancelOrder The order to be cancelled
@return CancelOrderResponse
*/
func (a *OrderManagementApiService) OrderCancelHMACSHA256Delete(ctx context.Context, cancelOrder CancelOrder) (CancelOrderResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CancelOrderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &cancelOrder
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CancelOrderResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrderManagementApiService Get order detail for specific order.
Get order details for specific order.  This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/details
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orderDetail The order to retrieve the details
@return OrderResponse
*/
func (a *OrderManagementApiService) OrderDetailsHMACSHA256Post(ctx context.Context, orderDetail OrderDetail) (OrderResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OrderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/details/ (HMAC SHA256)"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &orderDetail
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OrderResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrderManagementApiService Get the list of all orders for user.
Get the list of all orders for user. URL Example: https://trade.coss.io/c/api/v1/order/list/all
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orderAllRequest The order request for retrieving all orders. This API requires signing.
@return []OrderResponse
*/
func (a *OrderManagementApiService) OrderListAllHMACSHA256Post(ctx context.Context, orderAllRequest OrderAllRequest) ([]OrderResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []OrderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/list/all"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &orderAllRequest
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []OrderResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrderManagementApiService Get the list of completed orders for user.
Get completed orders for specific symbol where order status is  FILLED, PARTIAL_FILL, or CANCELED URL Example: https://trade.coss.io/c/api/v1/order/list/completed
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orderListRequest The order request for retrieving open orders. This API requires signing.
@return OrderListResponse
*/
func (a *OrderManagementApiService) OrderListCompletedHMACSHA256Post(ctx context.Context, orderListRequest OrderListRequest) (OrderListResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OrderListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/list/completed"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &orderListRequest
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OrderListResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrderManagementApiService Get the list of open orders for user.
Get current open orders for specific symbol OPEN and CANCELING.  This API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/list/open
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orderListRequest The order request for retrieving open orders. This API requires signing.
@return OrderListResponse
*/
func (a *OrderManagementApiService) OrderListOpenHMACSHA256Post(ctx context.Context, orderListRequest OrderListRequest) (OrderListResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OrderListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/list/open"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &orderListRequest
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OrderListResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OrderManagementApiService Get order's trade details.
Get trade details for an order. API requires signing of the payload. URL Example: https://trade.coss.io/c/api/v1/order/trade-detail
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tradeDetailRequest The order to retrieve the trade details
@return []TradeDetail
*/
func (a *OrderManagementApiService) OrderTradeDetailHMACSHA256Post(ctx context.Context, tradeDetailRequest TradeDetailRequest) ([]TradeDetail, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TradeDetail
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/order/trade-detail (HMAC SHA256)"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &tradeDetailRequest
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorisation"] = key
		}
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Signature"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TradeDetail
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
