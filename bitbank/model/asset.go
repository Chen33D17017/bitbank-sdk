package model

// data : []Asset

type Asset struct {
	Asset           string `json:"asset"`
	AmountPrecision int    `json:"amount_precision"`
	OnhandAmount    string `json:"onhand_amount"`
	FreeAmount      string `json:"free_amount"`
}
