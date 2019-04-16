package apis

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetCurrentBalance",
		"GET",
		"/balance/{token}",
		GetCurrentBalance,
	},
	Route{
		"ManualTrading",
		"GET",
		"/trade/manual",
		ManualTrading,
	},
	Route{
		"OrderAdd",
		"POST",
		"/trade/order_add",
		OrderAdd,
	},
}
