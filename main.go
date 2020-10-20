package main

import (
	"fmt"
)

func main() {
	exchange := NewExchange()
	exchange.CreateSecurity(XOM)

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   1,
		limit:  30.0,
		type_:  ORDER_TYPE_BUY,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   1,
		limit:  20.0,
		type_:  ORDER_TYPE_BUY,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   1,
		limit:  40.0,
		type_:  ORDER_TYPE_BUY,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   1,
		limit:  50.0,
		type_:  ORDER_TYPE_SELL,
	})

	security := exchange.GetSecurity(XOM)
	fmt.Printf("bid/ask spread: %v\n", security.limitOrderBook.BidAskSpread())
	fmt.Printf("estimate of value: %v\n", security.limitOrderBook.BestEstimateOfValue())
}
