package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"cyan.io/models"
	"cyan.io/openapi"
	"github.com/shopspring/decimal"
)

func CommandLineInput(defaultValue string, comment string) (inputVal string) {
	fmt.Print(comment)
	var val string
	fmt.Scanln(&val)

	if val == "" {
		val = defaultValue
	}

	return val
}

func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func PrintOrderBook(symbol string, side string, list [][]string, limit int, sizeLimit decimal.Decimal) (decimal.Decimal, decimal.Decimal) {

	if limit == 0 {
		limit = 100
	}

	// Set a temp upper limit as 1000 for any kind of token
	if sizeLimit.Equals(decimal.Zero) {
		sizeLimit = decimal.New(1000, 0)
	}
	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 16, 8, 0, '\t', 0)

	defer w.Flush()

	total := decimal.Zero
	totalSize := decimal.Zero
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "", "Symbol", "Side", "Price", "Size")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "--------", "--------", "--------", "--------")
	for i, order := range list {
		fmt.Fprintf(w, "\n %d\t%s\t%s\t%s\t%s\t \n", i+1, symbol, side, order[0], order[1])

		price, _ := decimal.NewFromString(order[0])
		size, _ := decimal.NewFromString(order[1])

		if sizeLimit.LessThanOrEqual(size) {
			return total.Add(price.Mul(sizeLimit)), totalSize.Add(sizeLimit)
		}
		totalSize = totalSize.Add(size)

		sizeLimit = sizeLimit.Sub(size)

		total = total.Add((price.Mul(size)))

		if i > limit-2 {
			return total, totalSize
		}
	}

	return decimal.Zero, decimal.Zero
}

func PrintTradingHistory(list []openapi.OrderResponse) {
	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 16, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "", "Symbol", "Side", "Status", "Price", "Size", "Executed", "Total")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "----", "--------", "--------", "--------", "--------", "--------", "--------", "--------")
	for i, order := range list {
		fmt.Fprintf(w, "\n %d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\t \n", i+1, order.OrderSymbol, order.OrderSide, order.Status, order.OrderPrice, order.OrderSize, order.Executed, order.Total)
	}
}

func PrintMatrixTable(list []models.LastPrice) {
	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 16, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "", "Trading_Pair", "Target_Price", "Last_Price", "Diff(%)")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "--------", "--------", "--------", "--------")
	for i, wish := range list {
		var diff decimal.Decimal
		if wish.CurrentPrice.Equals(decimal.Zero) {
			diff = decimal.New(-1, 0)
		} else {
			diff = wish.TargetPrice.Sub(wish.CurrentPrice).DivRound(wish.CurrentPrice, 8).Mul(HUNDRED)
		}
		fmt.Fprintf(w, "\n %d\t%s\t%s\t%s\t%s\t\t \n", i+1, wish.TradingPair, wish.TargetPrice, wish.CurrentPrice, diff.String())
	}
}
