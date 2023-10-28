package handlers

import (
	"fmt"
	"net/http"

	"github.com/alankritjoshi/coinwallet/clients/blockchain"
	"github.com/alankritjoshi/coinwallet/models"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func registerSyncHandler(app *pocketbase.PocketBase) {

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/collections/addresses/:id/sync", func(c echo.Context) error {
			id := c.PathParam("id")

			// Check if address exists
			address := models.Address{}
			err := app.Dao().DB().Select("id", "blockchain_address").From("addresses").One(&address)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("address_id %s not found", id))
			}

			// Fetch blockchain address
			blockchain_address, err := blockchain.GetBlockchainAddress(address.BlockchainAddress)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to fetch address from blockchain: %s", err))
			}

			return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("balance found: %d", blockchain_address.FinalBalance)})
		})

		return nil
	})

}
