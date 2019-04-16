package services

import (
	"log"

	"cyan.io/openapi"
	"cyan.io/utils"
	"github.com/antihax/optional"
	"github.com/shopspring/decimal"
)

func GetMarketPrice(tradingPair string) decimal.Decimal {
	var cfg = &openapi.Configuration{
		BasePath:      utils.BaseURL,
		DefaultHeader: make(map[string]string),
		UserAgent:     "openapi-Codeopenapi/1.0.0/go",
	}

	apiClient = openapi.NewAPIClient(cfg)

	marketPriceGetOpts := openapi.MarketPriceGetOpts{optional.NewString(tradingPair)}
	marketPrices, _, _ := apiClient.MarketApi.MarketPriceGet(ctx, &marketPriceGetOpts)
	for _, marketPrice := range marketPrices {
		price, _ := decimal.NewFromString(marketPrice.Price)
		// log.Printf("Buying margin price for %s is %v", tradingPair, price.Mul(decimal.NewFromFloat(1+margin)))
		// log.Printf("Selling margin price for %s is %v", tradingPair, price.Mul(decimal.NewFromFloat(1-margin)))

		return price
	}

	return decimal.Zero
}

func GetDepthForSymbol(tradingPair string) ([][]string, [][]string, error) {
	var cfg = &openapi.Configuration{
		BasePath:      "https://engine.coss.io/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "openapi-Codeopenapi/1.0.0/go",
	}

	apiClient = openapi.NewAPIClient(cfg)

	res, httpResponse, error := apiClient.MarketApi.DpGet(ctx, tradingPair)

	if error != nil {
		PrintHTTPResponse(httpResponse, error)
	}

	return res.Asks, res.Bids, error
}

func GetMarketSummary() (result []openapi.MarketSummary) {
	var cfg = &openapi.Configuration{
		BasePath:      "https://exchange.coss.io/api",
		DefaultHeader: make(map[string]string),
		UserAgent:     "openapi-Codeopenapi/1.0.0/go",
	}

	apiClient = openapi.NewAPIClient(cfg)

	ms, _, _ := apiClient.MarketApi.GetmarketsummariesGet(ctx)

	log.Printf("MarketSummary %d", len(ms.Result))

	return ms.Result
}
