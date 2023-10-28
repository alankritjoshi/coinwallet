package models

type Transaction struct {
	Id                   string  `db:"id" json:"id"`
	From                 string  `db:"from" json:"from"`
	To                   string  `db:"to" json:"to"`
	Amount               float64 `db:"amount" json:"amount"`
	TransactionHash      string  `db:"transaction_hash" json:"transaction_hash"`
	TransactionTimestamp string  `db:"transaction_timestamp" json:"transaction_timestamp"`
}
