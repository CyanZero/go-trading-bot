package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cyan.io/models"
	"cyan.io/openapi"
	"cyan.io/services"
	"cyan.io/utils"
	"github.com/shopspring/decimal"
)

func CommandLineMode() {

	fmt.Print("Welcome to Calvin's trading bot! Please choose(Default is 1. Manual): \n")
	fmt.Print("\t1. Manual\n\t2. Auto\n\t3. Arbitrage\n\t4. Check Trading History\n\t5. Get the market depth\n")
	var option string
	fmt.Scanln(&option)

	if option == "2" {
		LoopBot()
	} else if option == "3" {
		//TriangleArbitrage()
	} else if option == "4" {
		CheckTradingHistory()
	} else if option == "5" {
		// CheckMarketDepth()
	} else {
		ManualTrading()
	}
}

func CheckMarketDepth() {
	symbol := utils.CommandLineInput("", "Key in the symbol to check(Default is COSS/BTC) ")

	if symbol == "" {
		symbol = "COSS_BTC"
	}

	asks, bids, _ := services.GetDepthForSymbol(symbol)
	// total := orderListResponse.Total

	utils.PrintOrderBook(symbol, "SELL", asks, 5, decimal.Zero)
	utils.PrintOrderBook(symbol, "BUY", bids, 5, decimal.Zero)
}

func CheckTradingHistory() {
	symbol := utils.CommandLineInput("", "Key in the symbol to check(Default is COSS/BTC) ")

	if symbol == "" {
		symbol = "COSS_BTC"
	}

	orderListResponse := services.GetTradingHistory(symbol)
	// total := orderListResponse.Total
	orders := orderListResponse.List

	utils.PrintTradingHistory(orders)
}

func ManualTrading() {

	list := []models.LastPrice{}

	for i := 0; i < len(utils.TradingPairs); i++ {
		w := models.LastPrice{
			TradingPair:  utils.TradingPairs[i],
			TargetPrice:  decimal.NewFromFloat32(utils.TargetPrices[i]),
			CurrentPrice: services.GetMarketPrice(utils.TradingPairs[i]),
		}

		list = append(list, w)
	}
	utils.PrintMatrixTable(list)

	n := utils.CommandLineInput("", "Choose a number(Exit if invalid/non-selection): ")

	if n == "" {
		return
	}

	number, _ := strconv.Atoi(n)

	chosenPair := utils.TradingPairs[int(number)-1]
	fmt.Printf("The chosen pair: %s\n", chosenPair)

	trade := list[number-1]

	side := "BUY"
	calPrice := trade.CurrentPrice.Mul(decimal.NewFromFloat(0.99))
	baseToken := strings.Split(chosenPair, "_")[0]
	quoteToken := strings.Split(chosenPair, "_")[1]
	token := quoteToken
	convertedToken := baseToken
	if trade.CurrentPrice.GreaterThan(trade.TargetPrice) {
		side = "SELL"
		calPrice = trade.CurrentPrice.Mul(decimal.NewFromFloat(1.01))
		token = baseToken
		convertedToken = quoteToken
	}

	order := ComposeOrder(chosenPair, side, calPrice.String(), token)

	ConfirmOrder(order, convertedToken, token)

	// Repeat
	ManualTrading()

}

func ComposeOrder(chosenPair string, side string, calPrice string, placeOrderToken string) openapi.Order {

	price := utils.CommandLineInput(calPrice, "Key in "+side+" price for "+chosenPair+"(Press enter if keep the calculated price "+calPrice+"): ")

	accountBalance := services.GetAccountBalances(placeOrderToken)

	accountBalanceOfPair, _ := decimal.NewFromString(accountBalance.Available)
	fmt.Printf("Available balance for token %s is %v\n", placeOrderToken, accountBalanceOfPair)

	confirmedPrice, _ := decimal.NewFromString(price)
	fmt.Printf("Price is %v\n", confirmedPrice)

	var maxOrderSize decimal.Decimal
	if side == "BUY" {
		maxOrderSize = accountBalanceOfPair.DivRound(confirmedPrice, 6)
	} else {
		maxOrderSize = accountBalanceOfPair
	}

	orderSize := utils.CommandLineInput(maxOrderSize.String(), "Key in order_size(Max "+maxOrderSize.String()+"):")

	order := openapi.Order{
		OrderSymbol: chosenPair,
		OrderPrice:  price,
		OrderSide:   side,
		OrderSize:   orderSize,
		StopPrice:   "",
		Type:        "limit",
		// RecvWindow:  int32(recvWindow),
	}

	return order
}

func CalTotal(order openapi.Order) decimal.Decimal {
	confirmedOrderSize, _ := decimal.NewFromString(order.OrderSize)
	confirmedPrice, _ := decimal.NewFromString(order.OrderPrice)

	return confirmedOrderSize.Mul(confirmedPrice)
}

func ConfirmOrder(order openapi.Order, convertedToken string, token string) bool {
	services.GetLastOpenOrdersForSymbol(order.OrderSymbol)

	total := CalTotal(order)

	fmt.Printf("The new order: \n")
	fmt.Printf("\ttrading_pair: \t%s\n\torder_price: \t%v\n\torder_size: \t%s\n", order.OrderSymbol, order.OrderPrice, order.OrderSize)
	if "BUY" == order.OrderSide {
		fmt.Printf("\torder_side: \t%s(%s)\n\ttotal (%s): \t%v\n", order.OrderSide, convertedToken, token, total)
	} else {
		fmt.Printf("\torder_side: \t%s(%s)\n\ttotal (%s): \t%v\n", order.OrderSide, token, convertedToken, total)
	}

	fmt.Print("Confirm the order? (Yes/No) Must key in Yes to execute the order: ")
	var confirmOrder string
	fmt.Scanln(&confirmOrder)
	if confirmOrder != "Yes" {
		confirmOrder = "No"

		fmt.Println()
		return false
	}

	// Re-calculate the time due to user interaction
	order.Timestamp = utils.GetCurrentTimestamp()
	var localVarPostBody interface{}
	localVarPostBody = &order
	bodyBuf := &bytes.Buffer{}
	json.NewEncoder(bodyBuf).Encode(localVarPostBody)

	log.Printf("Body Order is %+s", bodyBuf.String())

	services.SignedGetRequest([]byte(bodyBuf.String()))
	services.PlaceOrder(order)

	return true
}

func LoopBot() {

	for counter := 1; ; counter++ {
		log.Printf("---------->Attempt no %d", counter)
		for i := 0; i < len(utils.TradingPairs); i++ {

			tradingPair := utils.TradingPairs[i]
			targetPrice := decimal.NewFromFloat32(utils.TargetPrices[i])
			buyingPrice := targetPrice.Mul(decimal.NewFromFloat(1 - utils.Margin))
			sellingPrice := targetPrice.Mul(decimal.NewFromFloat(1 + utils.Margin))
			// log.Printf("The buying price for %s is: %v", tradingPair, buyingPrice)
			// log.Printf("The selling price for %s is: %v", tradingPair, sellingPrice)

			currentPrice := services.GetMarketPrice(tradingPair)

			// Reduce 1% further to avoid lost
			price := currentPrice.Mul(decimal.NewFromFloat(0.79))
			orderSize := strconv.FormatFloat(utils.OrderSizes[i], 'f', -1, 64)
			total := price.Mul(decimal.NewFromFloat(utils.OrderSizes[i]))

			order := openapi.Order{
				OrderSymbol: tradingPair,
				OrderSize:   orderSize,
				StopPrice:   "",
				Type:        "limit",
				// RecvWindow:  int32(recvWindow),
			}

			log.Printf("Last price for %s is %v", tradingPair, currentPrice)
			difference := (currentPrice.Sub(targetPrice)).DivRound(targetPrice, 4).Mul(utils.HUNDRED)
			log.Printf("The difference is %v(%%)", (difference))

			baseToken := strings.Split(tradingPair, "_")[0]
			quoteToken := strings.Split(tradingPair, "_")[1]
			token := quoteToken
			convertedToken := baseToken
			if buyingPrice.GreaterThanOrEqual(currentPrice) {
				order.OrderSide = utils.BUY
				order.OrderPrice = currentPrice.Mul(decimal.NewFromFloat(0.99)).String()
			} else if sellingPrice.LessThanOrEqual(currentPrice) {
				order.OrderSide = utils.SELL
				token = baseToken
				convertedToken = quoteToken
				order.OrderPrice = currentPrice.Mul(decimal.NewFromFloat(1.01)).String()
			} else {
				// log.Printf("The last price %v doesn't match profit margin for pair %s", currentPrice, tradingPair)
				continue
			}

			// Check available balance
			accountBalance := services.GetAccountBalances(token)
			accountBalanceOfPair, _ := decimal.NewFromString(accountBalance.Available)
			if accountBalanceOfPair.LessThan(total) {
				log.Printf("There is no sufficient balance(%v) of %s to buy %s", accountBalance, token, convertedToken)
				continue
			}

			ConfirmOrder(order, convertedToken, token)
		}

		time.Sleep(30 * time.Second)
	}
}
