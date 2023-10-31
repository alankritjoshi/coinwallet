package workers

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/alankritjoshi/coinwallet/clients/blockchain"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func SyncTransactionsForAddress(app *pocketbase.PocketBase, address string) {
	log.Printf("sync STARTED for %s", address)
	err := fetchTransactionsForAddress(app, address)
	if err != nil {
		log.Printf("failed to fetch transactions for address %s: %v", address, err)
		log.Printf("sync FINISHED for %s", address)
		return
	}
	log.Printf("sync FINISHED for %s", address)
}

func fetchTransactionsForAddress(app *pocketbase.PocketBase, address string) error {
	// Make the first call to fetch the number of transactions and first LIMIT transactions
	blockchainAddress, err := blockchain.GetBlockchainAddress(address, 0)
	if err != nil {
		return fmt.Errorf("failed to fetch first transaction: %w", err)
	}

	// Update balance
	address_id, err := updateAddressBalance(app, address, blockchainAddress.FinalBalance)
	if err != nil {
		return fmt.Errorf("failed to update address balance: %w", err)
	}

	// Create channels for transactions and errors
	transactionsChan := make(chan []blockchain.BlockchainTransaction)
	errorChan := make(chan error)

	defer close(transactionsChan)
	defer close(errorChan)

	// Start fetching transactions in parallel using Goroutines
	for i := 0; i < numPaginatedCalls(blockchainAddress.NumberOfTransactions); i++ {
		go func(offset int) {
			blockchainAddress, err := blockchain.GetBlockchainAddress(address, offset)
			if err != nil {
				log.Printf("failed to fetch transactions for %s", address)
				errorChan <- fmt.Errorf("failed to fetch transactions %d to %d: %w", offset, offset+blockchain.LIMIT, err)
				return
			}

			transactionsChan <- blockchainAddress.BlockchainTransactions
		}(i * blockchain.LIMIT)
	}

	for i := 0; i < numPaginatedCalls(blockchainAddress.NumberOfTransactions); i++ {
		select {
		case transactionsBatch := <-transactionsChan:
			err := createTransactionsBatchInDB(app, address_id, transactionsBatch)
			if err != nil {
				return fmt.Errorf("failed to store transactions in database: %w", err)
			}
		case err := <-errorChan:
			return fmt.Errorf("failed to fetch transactions: %w", err)
		}
	}

	return nil
}

func createTransactionsBatchInDB(app *pocketbase.PocketBase, address_id *string, transactions []blockchain.BlockchainTransaction) error {
	transactionsCollection, err := app.Dao().FindCollectionByNameOrId("transactions")
	if err != nil {
		return fmt.Errorf("failed to find transactions collection: %w", err)
	}

	app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		for _, transaction := range transactions {
			record := models.NewRecord(transactionsCollection)
			// TODO: make sure from and to are set using the transaction addresses
			record.Set("id", transaction.Hash)
			record.Set("to", []string{*address_id})
			record.Set("index", transaction.TxIndex)
			record.Set("relayed_by", transaction.RelayedBy)
			record.Set("block_height", transaction.BlockHeight)
			record.Set("version", transaction.Ver)
			if err := txDao.SaveRecord(record); err != nil {
				if !strings.Contains(err.Error(), "UNIQUE constraint failed: transactions.id") {
					log.Printf("failed to save transaction record for %s: %v", transaction.Hash, err)
					return err
				}
				log.Printf("duplicate record for %s: %v. skipping", transaction.Hash, err)
			} else {
				log.Printf("record saved for %s", transaction.Hash)
			}
		}

		return nil
	})

	return nil
}

func updateAddressBalance(app *pocketbase.PocketBase, address string, balance int64) (*string, error) {
	record, err := app.Dao().FindFirstRecordByFilter("addresses", fmt.Sprintf("blockchain_address = '%s'", address))
	if err != nil {
		return nil, fmt.Errorf("failed to find address record: %w", err)
	}

	record.Set("balance", balance)

	if err := app.Dao().SaveRecord(record); err != nil {
		return nil, fmt.Errorf("failed to update address record: %w", err)
	}

	return &record.Id, nil
}

func numPaginatedCalls(numTransactions int) int {
	return int(math.Ceil(float64(numTransactions) / float64(blockchain.LIMIT)))
}
