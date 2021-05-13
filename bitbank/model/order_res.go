package model

type Order struct {
	OrderId         int64  `json:"order_id"`
	Pair            string `json:"pair"`
	Type            string `json:"type"`
	StartAmount     string `json:"start_amount"`
	RemainingAmount string `json:"remaining_amount"`
	ExecutedAmount  string `json:"executed_amount"`
	Price           string `json:"Price"`
	AveragePrice    string `json:"average_price"`
	OrderedAt       int64  `json:"ordered_at"`
	Status          string `json:"status"`
}
