package main

import (
	"math"

	"github.com/google/uuid"
)

type LimitOrder struct {
	size  int
	limit float64
}

type LimitOrderBook struct {
	buyOrders  map[uuid.UUID]LimitOrder
	sellOrders map[uuid.UUID]LimitOrder
}

func NewLimitOrderBook() *LimitOrderBook {
	return &LimitOrderBook{
		buyOrders:  map[uuid.UUID]LimitOrder{},
		sellOrders: map[uuid.UUID]LimitOrder{},
	}
}

func (l *LimitOrderBook) AddBuyOrder(order LimitOrder) error {
	l.buyOrders[uuid.New()] = order
	return nil
}

func (l *LimitOrderBook) AddSellOrder(order LimitOrder) error {
	l.sellOrders[uuid.New()] = order
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
