package main

import (
	"fmt"
	"testing"

	"../services"
)

func TestGetDepthForSymbol(t *testing.T) {
	symbol := "EOS_USDT"
	asks, bids, err := services.GetDepthForSymbol(symbol)

	fmt.Printf("OrderHisotry for %s is %v %v\n", symbol, asks, bids)

	if err != nil {
		t.Errorf("Failed to get depth orders for symbol: %s", symbol)
	}
}
