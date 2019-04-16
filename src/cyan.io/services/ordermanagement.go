package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"cyan.io/openapi"
	"cyan.io/utils"
)

func GetTradingHistory(symbol string) openapi.OrderListResponse {
	log.Printf("Checking trading history for symbol %s", symbol)

	orderListRequest := openapi.OrderListRequest{
		Limit:      int32(10),
		Page:       0,
		Symbol:     symbol,
		Timestamp:  utils.GetCurrentTimestamp(),
		RecvWindow: recvWindow,
	}

	SignedPostRequest(&orderListRequest)
	orderListResponse, httpResponse, err := apiClient.OrderManagementApi.OrderListCompletedHMACSHA256Post(ctx, orderListRequest)
	if err != nil {
		PrintHTTPResponse(httpResponse, err)
	}

	return orderListResponse
}

func PlaceOrder(order openapi.Order) {
	log.Println("Gonna place a buying order now!")
	apiClient = openapi.NewAPIClient(cfg)

	orderResponse, httpResponse, err := apiClient.OrderManagementApi.OrderAddHMACSHA256Post(ctx, order)
	if err == nil {
		fmt.Printf("Placed an order successfully: %+v\n", orderResponse)
	} else {
		fmt.Printf("Failed to place an order: %+v\n", orderResponse)
		PrintHTTPResponse(httpResponse, err)
	}
}

func GetLastOpenOrdersForSymbol(symbol string) {

	orderListRequest := openapi.OrderListRequest{
		Limit:      int32(10),
		Page:       0,
		Symbol:     symbol,
		Timestamp:  utils.GetCurrentTimestamp(),
		RecvWindow: recvWindow,
	}

	// Re-calculate the time due to user interaction
	var localVarPostBody interface{}
	localVarPostBody = &orderListRequest
	bodyBuf := &bytes.Buffer{}
	json.NewEncoder(bodyBuf).Encode(localVarPostBody)

	log.Printf("Body Order is %+s", bodyBuf.String())

	SignedGetRequest([]byte(bodyBuf.String()))
	orderListResponse, httpResponse, err := apiClient.OrderManagementApi.OrderListOpenHMACSHA256Post(ctx, orderListRequest)

	if err != nil {
		PrintHTTPResponse(httpResponse, err)
	}

	log.Printf("Open orders for %s is %d", symbol, len(orderListResponse.List))
	if orderListResponse.Total > 0 {
		log.Printf("%d open orders for %s found: ", orderListResponse.Total, symbol)
		for i, openOrder := range orderListResponse.List {
			fmt.Printf("No.%d open orders with %s is: \n", i+1, symbol)
			fmt.Printf("\tOrderSide:\t%s\n", openOrder.OrderSide)
			fmt.Printf("\tStatus:\t%s\n", openOrder.Status)
			fmt.Printf("\tCreateTime:\t%d\n", openOrder.CreateTime)
			fmt.Printf("\tType:\t%s\n", openOrder.Type)
			fmt.Printf("\tOrderPrice:\t%s\n", openOrder.OrderPrice)
			fmt.Printf("\tExecuted:\t%s\n", openOrder.Executed)
			fmt.Printf("\tAvg:\t%v\n", openOrder.Avg)
			fmt.Printf("\tTotal:\t%v\n", openOrder.Total)
		}
	}
}
