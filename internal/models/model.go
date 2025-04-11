package models

type PaymentMethod struct {
	Account string `json:"account"`
	Status  string `json:"status"`
	Balance string `json:"balance"`
	Icon    string `json:"icon"`
}
