package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cyan.io/models"
	"cyan.io/openapi"
	"cyan.io/services"
	"cyan.io/utils"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetCurrentBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	if token != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		accountBalance := services.GetAccountBalances(token)
		if err := json.NewEncoder(w).Encode(accountBalance); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func ManualTrading(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	list := []models.LastPrice{}

	for i := 0; i < len(utils.TradingPairs); i++ {
		w := models.LastPrice{
			TradingPair:  utils.TradingPairs[i],
			TargetPrice:  decimal.NewFromFloat32(utils.TargetPrices[i]),
			CurrentPrice: services.GetMarketPrice(utils.TradingPairs[i]),
		}

		list = append(list, w)
	}
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
	return

}

func OrderAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Body)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var order openapi.Order
	_ = json.NewDecoder(r.Body).Decode(&order)

	log.Printf("Recieved order: %v", order)

	if order.OrderSymbol != "" {

		// Re-calculate the time due to user interaction
		order.Timestamp = utils.GetCurrentTimestamp()
		var localVarPostBody interface{}
		localVarPostBody = &order
		bodyBuf := &bytes.Buffer{}
		json.NewEncoder(bodyBuf).Encode(localVarPostBody)

		log.Printf("Body Order is %+s", bodyBuf.String())

		services.SignedGetRequest([]byte(bodyBuf.String()))
		services.PlaceOrder(order)
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Invalid input"}); err != nil {
		panic(err)
	}

}
