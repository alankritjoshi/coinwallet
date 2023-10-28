package models

type Transaction struct {
	Id          string `db:"id" json:"id"`
	Address     string `db:"address" json:"address"`
	Hash        string `db:"hash" json:"hash"`
	Size        int64  `db:"size" json:"size"`
	Timestamp   string `db:"timestamp" json:"timestamp"`
	Index       int64  `db:"index" json:"index"`
	RelayedBy   string `db:"relayed_by" json:"relayed_by"`
	BlockHeight int64  `db:"block_height" json:"block_height"`
	InputsSize  int64  `db:"inputs_size" json:"inputs_size"`
	OutputSize  int64  `db:"outputs_size" json:"outputs_size"`
	Version     string `db:"version" json:"version"`
	Inputs      string `db:"inputs" json:"inputs"`
	Outputs     string `db:"outputs" json:"outputs"`
}
