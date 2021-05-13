package model

type GeneralRes struct {
	Msg  int8        `json:"success"`
	Data interface{} `json:"data"`
}
