package main

type Security struct {
	limitOrderBook *LimitOrderBook
}

func NewSecurity(symbol Symbol) *Security {
	return &Security{
		limitOrderBook: NewLimitOrderBook(),
	}
}

func (s *Security) ProcessLimitOrderRequest(request LimitOrderRequest) error {
	order := LimitOrder{
		size:  request.size,
		limit: request.limit,
	}
	if request.type_ == ORDER_TYPE_BUY {
		return s.limitOrderBook.AddBuyOrder(order)
	} else if request.type_ == ORDER_TYPE_SELL {
		return s.limitOrderBook.AddSellOrder(order)
	}
	return nil
}

func (s *Security) SinglePriceCall() SinglePriceCallResult {
	return s.limitOrderBook.SinglePriceCall()
}

// look in market order book.
// should be able to buy any sell order or sell any buy order, even if it's not the best
func (s *Security) ProcessMarketOrderRequest(request MarketOrderRequest) error {
	return nil
}
