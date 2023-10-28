package models

type Wallet struct {
	Id     string `db:"id" json:"id"`
	UserID string `db:"user_id" json:"user_id"`
	Label  string `db:"label" json:"label"`
}
