package blockchain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockchainAddress(address string) (*BlockchainAddress, error) {
	// Make HTTP GET request to fetch transactions
	resp, err := http.Get(fmt.Sprintf("https://blockchain.info/rawaddr/%s", address))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for error message in response body
	errorResponse := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{}
	err = json.Unmarshal(body, &errorResponse)
	if err == nil && errorResponse.Error != "" {
		return nil, fmt.Errorf("failed to fetch transactions: %s", errorResponse.Message)
	}

	// Parse response body
	blockchainAddress := &BlockchainAddress{}
	err = json.Unmarshal(body, blockchainAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return blockchainAddress, nil
}
