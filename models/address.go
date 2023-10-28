package models

type Address struct {
	Id                string `db:"id" json:"id"`
	WalletId          string `db:"wallet_id" json:"wallet_id"`
	BlockchainAddress string `db:"blockchain_address" json:"blockchain_address"`
	Label             string `db:"label" json:"label"`
}
