package services

import (
	"context"
	"log"
	"strings"

	"cyan.io/openapi"
	"cyan.io/utils"
	"github.com/shopspring/decimal"
)

// TODO

var recvWindow = int32(5000)

var cfg = &openapi.Configuration{
	BasePath:      utils.BaseURL,
	DefaultHeader: make(map[string]string),
	UserAgent:     "openapi-Codeopenapi/1.0.0/go",
}

// Default value of apiClient, should re-create per usage
var apiClient = openapi.NewAPIClient(cfg)

var ctx = context.Background()

func GetAccountBalances(tradingPair string) (accountBalance openapi.AccountBalance) {
	tradingPair = strings.ToUpper(tradingPair)
	apiClient = openapi.NewAPIClient(cfg)
	//accountBalanceOpt := openapi.AccountBalancesHMACSHA256GetOpts{optional.NewInt32(recvWindow)}

	timeNow := utils.GetCurrentTimestamp()
	ComposeHeaderForGetRequest(timeNow)
	accountBalances, _, err := apiClient.AccountApi.AccountBalancesHMACSHA256Get(ctx, timeNow, nil)
	if err != nil {
		log.Fatal("accountBalances Error:", err)
	}

	log.Printf("===========> Account wallets count: %d", len(accountBalances))

	for _, accountBalance := range accountBalances {

		if tradingPair == accountBalance.CurrencyCode {
			balance, _ := decimal.NewFromString(accountBalance.Total)

			log.Printf("Balance of currency %s is %s", accountBalance.CurrencyCode, balance)
			log.Printf("Availabe balance of currency %s is %s", accountBalance.CurrencyCode, accountBalance.Available)
			return accountBalance
		}

	}

	var a openapi.AccountBalance
	return a
}
