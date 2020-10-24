package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglePriceCall(t *testing.T) {
	exchange := NewExchange()
	exchange.CreateSecurity(XOM)

	// From Trading & Exchanges Chapter 6
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

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   1,
		limit:  19.8,
		type_:  ORDER_TYPE_SELL,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   5,
		limit:  20.2,
		type_:  ORDER_TYPE_SELL,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   4,
		limit:  20.3,
		type_:  ORDER_TYPE_BUY,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   2,
		limit:  20.1,
		type_:  ORDER_TYPE_BUY,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   6,
		limit:  20.0,
		type_:  ORDER_TYPE_SELL,
	})

	exchange.SubmitLimitOrder(LimitOrderRequest{
		symbol: XOM,
		size:   7,
		limit:  19.8,
		type_:  ORDER_TYPE_BUY,
	})

	security := exchange.GetSecurity(XOM)
	callResult := security.SinglePriceCall()

	assert.InDelta(t, 20.0, callResult.Price, 0.01)
	assert.InDelta(t, 1.6, callResult.TraderSurplus, 0.01)
}
