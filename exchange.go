package main

type Symbol int

const (
	XOM = iota
)

type OrderType int

const (
	ORDER_TYPE_BUY = iota
	ORDER_TYPE_SELL
)

type MarketOrderRequest struct {
	symbol Symbol
	size   int
}

type MarketOrderResult struct {
	err error
}

type LimitOrderRequest struct {
	symbol Symbol
	size   int
	limit  float64
	type_  OrderType
}

type LimitOrderResult struct {
	err error
}

type Exchange struct {
	securities map[Symbol]*Security
}

func NewExchange() *Exchange {
	return &Exchange{
		securities: map[Symbol]*Security{},
	}
}

func (e *Exchange) CreateSecurity(symbol Symbol) {
	e.securities[symbol] = NewSecurity(symbol)
}

func (e *Exchange) GetSecurity(symbol Symbol) *Security {
	return e.securities[symbol]
}

func (e *Exchange) SubmitLimitOrder(request LimitOrderRequest) LimitOrderResult {
	err := e.securities[request.symbol].ProcessLimitOrderRequest(request)
	return LimitOrderResult{
		err: err,
	}
}

func (e *Exchange) SubmitMarketOrder(request MarketOrderRequest) MarketOrderResult {
	err := e.securities[request.symbol].ProcessMarketOrderRequest(request)
	return MarketOrderResult{
		err: err,
	}
}
