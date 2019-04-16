package models

import "github.com/shopspring/decimal"

/**
 * The data model for last price
 */
type LastPrice struct {
	TradingPair  string          `json:"TradingPair"`
	TargetPrice  decimal.Decimal `json:"TargetPrice"`
	CurrentPrice decimal.Decimal `json:"CurrentPrice"`
	OrderSize    int             `json:"OrderSize"`
}
