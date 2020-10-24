package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglePriceCall(t *testing.T) {
	book := NewLimitOrderBook()

	// From Trading & Exchanges Chapter 6
	book.AddBuyOrder(LimitOrder{
		size:  3,
		limit: 20.0,
	})

	book.AddSellOrder(LimitOrder{
		size:  2,
		limit: 20.1,
	})

	book.AddBuyOrder(LimitOrder{
		size:  2,
		limit: 20.0,
	})

	book.AddSellOrder(LimitOrder{
		size:  1,
		limit: 19.8,
	})

	book.AddSellOrder(LimitOrder{
		size:  5,
		limit: 20.2,
	})

	book.AddBuyOrder(LimitOrder{
		size:  4,
		limit: 20.3,
	})

	book.AddBuyOrder(LimitOrder{
		size:  2,
		limit: 20.1,
	})

	book.AddSellOrder(LimitOrder{
		size:  6,
		limit: 20.0,
	})

	book.AddBuyOrder(LimitOrder{
		size:  7,
		limit: 19.8,
	})

	callResult := book.SinglePriceCall()

	assert.InDelta(t, 20.0, callResult.TradePrice, 0.01)
	assert.InDelta(t, 1.6, callResult.TraderSurplus(), 0.01)
}
