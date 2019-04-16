package utils

import "github.com/shopspring/decimal"

const PORT = ":8080"

const BaseURL = "https://trade.coss.io/c/api/v1"

// TODO some custom trading preferences which are hardcoded at the moment
var TradingPairs = []string{"COSS_USDT", "COSS_BTC", "COSS_ETH", "BTC_USDT", "ETH_USDT", "NEO_USDT"}
var TargetPrices = []float32{0.065, 0.0000165, 0.000445, 3900, 145, 8.61}
var OrderSizes = []float64{100, 100, 100, 0.02, 0.2, 1}

// Margin is the default profile margin used in the auto-trading
var Margin = 0.05

const BUY string = "BUY"
const SELL string = "SELL"

var HUNDRED = decimal.New(100, 0)

// Authorization is API auth key for coss
var Authorization string

// Secret is API secret for coss
var Secret []byte
