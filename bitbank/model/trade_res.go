package model

type TradesResponse struct {
	Msg  int8 `json:"success"`
	Data struct {
		Trades []Trade `json:"trades"`
	} `json:"data"`
}

type Trades struct []Trade `json:"trades"`

type Trade struct {
	TradeId        int64  `json:"trade_id"`
	Pair           string `json:"pair"`
	OrderId        int64  `json:"order_id"`
	Side           string `json:"side"`
	Type           string `json:"type"`
	Amount         string `json:"amount"`
	Price          string `json:"price"`
	MakerTaker     string `json:"maker_taker"`
	FeeAmountBase  string `json:"fee_amount_base"`
	FeeAmountQuote string `json:"fee_amount_quote"`
	ExecutedAt     int64  `json:"executed_at"`
}