package blockchain

type BlockchainAddress struct {
	Address                string                  `json:"address"`
	NumberOfTransactions   int                     `json:"n_tx"`
	TotalReceived          int64                   `json:"total_received"`
	TotalSent              int64                   `json:"total_sent"`
	FinalBalance           int64                   `json:"final_balance"`
	BlockchainTransactions []BlockchainTransaction `json:"txs"`
}
