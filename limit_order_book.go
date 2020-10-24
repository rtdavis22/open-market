package main

import (
	"math"
	"sort"
	//"github.com/google/uuid"
)

type LimitOrder struct {
	size  int
	limit float64
}

type PendingLimitOrder struct {
	unfilled int
	LimitOrder
}

type LimitOrderBook struct {
	buyOrders  []*PendingLimitOrder
	sellOrders []*PendingLimitOrder
}

func NewLimitOrderBook() *LimitOrderBook {
	return &LimitOrderBook{
		buyOrders:  []*PendingLimitOrder{},
		sellOrders: []*PendingLimitOrder{},
	}
}

func (l *LimitOrderBook) AddBuyOrder(order LimitOrder) error {
	l.buyOrders = append(l.buyOrders, &PendingLimitOrder{
		unfilled:   order.size,
		LimitOrder: order,
	})
	return nil
}

func (l *LimitOrderBook) AddSellOrder(order LimitOrder) error {
	l.sellOrders = append(l.sellOrders, &PendingLimitOrder{
		unfilled:   order.size,
		LimitOrder: order,
	})
	return nil
}

func (l *LimitOrderBook) BestBid() float64 {
	best := 0.0
	for _, order := range l.buyOrders {
		if order.limit > best {
			best = order.limit
		}
	}
	return best
}

func (l *LimitOrderBook) BestAsk() float64 {
	best := math.MaxFloat64
	for _, order := range l.sellOrders {
		if order.limit < best {
			best = order.limit
		}
	}
	return best
}

func (l *LimitOrderBook) BidAskSpread() float64 {
	return l.BestAsk() - l.BestBid()
}

func (l *LimitOrderBook) BestEstimateOfValue() float64 {
	return (l.BestAsk() + l.BestBid()) / 2.0
}

type Trade struct {
	buyOrder  LimitOrder
	sellOrder LimitOrder
	size      int
}

type SinglePriceCallResult struct {
	Trades     []Trade
	TradePrice float64
}

func (s SinglePriceCallResult) BuyerSurplus() float64 {
	var surplus float64
	for _, trade := range s.Trades {
		surplus += (trade.buyOrder.limit - s.TradePrice) * float64(trade.size)
	}
	return surplus
}

func (s SinglePriceCallResult) SellerSurplus() float64 {
	var surplus float64
	for _, trade := range s.Trades {
		surplus += (s.TradePrice - trade.sellOrder.limit) * float64(trade.size)
	}
	return surplus
}

func (s SinglePriceCallResult) TraderSurplus() float64 {
	return s.BuyerSurplus() + s.SellerSurplus()
}

// Conduct a single-price auction and return the price of the auction and the trader surplus.
func (l *LimitOrderBook) SinglePriceCall() SinglePriceCallResult {
	sort.SliceStable(l.buyOrders, func(i int, j int) bool {
		return l.buyOrders[i].limit > l.buyOrders[j].limit
	})
	sort.SliceStable(l.sellOrders, func(i int, j int) bool {
		return l.sellOrders[i].limit < l.sellOrders[j].limit
	})

	var i, j int
	trades := []Trade{}
	for i < len(l.buyOrders) && j < len(l.sellOrders) {
		buyOrder := l.buyOrders[i]
		sellOrder := l.sellOrders[j]

		// No trades left to make.
		if buyOrder.limit < sellOrder.limit {
			break
		}

		var tradeSize int
		if buyOrder.unfilled == sellOrder.unfilled {
			tradeSize = buyOrder.unfilled
		} else if buyOrder.unfilled < sellOrder.unfilled {
			tradeSize = buyOrder.unfilled
			i++
		} else {
			tradeSize = sellOrder.unfilled
			j++
		}

		trades = append(trades, Trade{
			buyOrder:  buyOrder.LimitOrder,
			sellOrder: sellOrder.LimitOrder,
			size:      tradeSize,
		})

		buyOrder.unfilled -= tradeSize
		sellOrder.unfilled -= tradeSize
	}

	return SinglePriceCallResult{
		Trades: trades,
		// Harris says the limits of the last buy and last sell should almost always be the same.
		TradePrice: trades[len(trades)-1].buyOrder.limit,
	}
}
