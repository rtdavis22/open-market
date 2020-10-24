package main

import (
	"fmt"
)

func main() {
	exchange := NewExchange()
	exchange.CreateSecurity(XOM)

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   3,
		limit:  20.0,
		type_:  ORDER_TYPE_BUY,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   2,
		limit:  20.1,
		type_:  ORDER_TYPE_SELL,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   2,
		limit:  20.0,
		type_:  ORDER_TYPE_BUY,
	})

	security := exchange.GetSecurity(XOM)

	callResult := security.SinglePriceCall()
	fmt.Printf("single price call price: %v\n", callResult.TradePrice)
	fmt.Printf("trader surplus: %v\n", callResult.BuyerSurplus()+callResult.SellerSurplus())
	fmt.Printf("bid/ask spread: %v\n", security.limitOrderBook.BidAskSpread())
	fmt.Printf("estimate of value: %v\n", security.limitOrderBook.BestEstimateOfValue())
}
