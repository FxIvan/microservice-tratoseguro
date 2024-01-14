package models

type Transaction struct {
	ID              string  `json:"id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	Operation       string  `json:"operation"`
	TransactionDate string  `json:"transaction_date"`
}