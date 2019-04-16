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
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the COSS public api for trading API v1.2
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccountApi *AccountApiService

	DefaultApi *DefaultApiService

	ExchangeInformationApi *ExchangeInformationApiService

	MarketApi *MarketApiService

	OrderManagementApi *OrderManagementApiService

	ServerInformationApi *ServerInformationApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccountApi = (*AccountApiService)(&c.common)
	c.DefaultApi = (*DefaultApiService)(&c.common)
	c.ExchangeInformationApi = (*ExchangeInformationApiService)(&c.common)
	c.MarketApi = (*MarketApiService)(&c.common)
	c.OrderManagementApi = (*OrderManagementApiService)(&c.common)
	c.ServerInformationApi = (*ServerInformationApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	//dump http request for debugging purpose only

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))

	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}

	// s := strings.Replace(bodyBuf.String(), ":", ": ", -1)
	// s = strings.Replace(s, ",", ", ", -1)
	// b := &bytes.Buffer{}
	// b.WriteString(s)
	// log.Printf("Body is %+s", s)
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
