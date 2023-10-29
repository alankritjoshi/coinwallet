package handlers

import (
	"fmt"
	"net/http"

	"github.com/alankritjoshi/coinwallet/models"
	"github.com/alankritjoshi/coinwallet/workers"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func registerSyncHandler(app *pocketbase.PocketBase) {

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/collections/addresses/records/:id/sync", func(c echo.Context) error {
			id := c.PathParam("id")

			// Check if address exists
			address := models.Address{}
			err := app.Dao().DB().Select("id", "blockchain_address").From("addresses").One(&address)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("address_id %s not found", id))
			}

			// Begin fetching and storing blockchain balance and transactions
			go func() {
				workers.SyncTransactionsForAddress(app, address.BlockchainAddress)
			}()

			return c.JSON(http.StatusAccepted, map[string]string{"message": "sync started for address_id " + id})
		})

		return nil
	})

}
