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

type LimitOrderBook struct {
	buyOrders  []LimitOrder
	sellOrders []LimitOrder
}

func NewLimitOrderBook() *LimitOrderBook {
	return &LimitOrderBook{
		buyOrders:  []LimitOrder{},
		sellOrders: []LimitOrder{},
	}
}

func (l *LimitOrderBook) AddBuyOrder(order LimitOrder) error {
	l.buyOrders = append(l.buyOrders, order)
	return nil
}

func (l *LimitOrderBook) AddSellOrder(order LimitOrder) error {
	l.sellOrders = append(l.sellOrders, order)
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

type SinglePriceCallResult struct {
	Price         float64
	TraderSurplus float64
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
	var lastPrice float64

	type TransactedUnit struct {
		size        int
		traderLimit float64
	}
	boughtUnits := []TransactedUnit{}
	soldUnits := []TransactedUnit{}
	for i < len(l.buyOrders) && j < len(l.sellOrders) {
		// buy price is less than sell price, we are done
		if l.buyOrders[i].limit < l.sellOrders[j].limit {
			break
		}
		var transactionSize int
		boughtUnit := TransactedUnit{
			traderLimit: l.buyOrders[i].limit,
		}
		soldUnit := TransactedUnit{
			traderLimit: l.sellOrders[j].limit,
		}
		if l.buyOrders[i].size == l.sellOrders[j].size {
			transactionSize = l.buyOrders[i].size
			lastPrice = l.buyOrders[i].limit // for now, using the buy price
			i++
			j++
		} else if l.buyOrders[i].size < l.sellOrders[j].size {
			transactionSize = l.buyOrders[i].size
			l.sellOrders[j].size -= transactionSize
			lastPrice = l.buyOrders[i].limit
			i++
		} else {
			transactionSize = l.sellOrders[j].size
			l.buyOrders[i].size -= transactionSize
			lastPrice = l.buyOrders[i].limit
			j++
		}

		boughtUnit.size = transactionSize
		soldUnit.size = transactionSize

		boughtUnits = append(boughtUnits, boughtUnit)
		soldUnits = append(soldUnits, soldUnit)
	}

	var traderSurplus float64
	for _, unit := range boughtUnits {
		traderSurplus += (unit.traderLimit - lastPrice) * float64(unit.size)
	}
	for _, unit := range soldUnits {
		traderSurplus += (lastPrice - unit.traderLimit) * float64(unit.size)
	}

	return SinglePriceCallResult{
		Price:         lastPrice,
		TraderSurplus: traderSurplus,
	}
}
