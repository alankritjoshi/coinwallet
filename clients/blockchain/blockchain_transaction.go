package blockchain

type BlockchainTransaction struct {
	Hash        string `json:"hash"`
	Ver         int64  `json:"ver"`
	VinSz       int64  `json:"vin_sz"`
	VoutSz      int64  `json:"vout_sz"`
	LockTime    int64  `json:"lock_time"`
	Size        int64  `json:"size"`
	RelayedBy   string `json:"relayed_by"`
	BlockHeight int64  `json:"block_height"`
	TxIndex     int64  `json:"tx_index"`
	Inputs      []struct {
		PrevOut struct {
			Hash  string `json:"hash"`
			Value int64  `json:"value"`
			TxIdx int64  `json:"tx_index"`
			N     int64  `json:"n"`
		} `json:"prev_out"`
		Script string `json:"script"`
	} `json:"inputs"`
	Out []struct {
		Value  int64  `json:"value"`
		Hash   string `json:"hash"`
		Script string `json:"script"`
	} `json:"out"`
}

func (t *BlockchainTransaction) GetFromAddresses() []string {
	var hashes []string

	for _, input := range t.Inputs {
		hashes = append(hashes, input.PrevOut.Hash)
	}

	return hashes
}

func (t *BlockchainTransaction) GetToAddresses() []string {
	var hashes []string

	for _, output := range t.Out {
		hashes = append(hashes, output.Hash)
	}

	return hashes
}
