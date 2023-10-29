package models

type Transaction struct {
	Id          string   `db:"id" json:"id"`
	From        []string `db:"from" json:"from"`
	To          []string `db:"to" json:"to"`
	Hash        string   `db:"hash" json:"hash"`
	Index       int64    `db:"index" json:"index"`
	RelayedBy   string   `db:"relayed_by" json:"relayed_by"`
	BlockHeight int64    `db:"block_height" json:"block_height"`
	Version     string   `db:"version" json:"version"`
}
