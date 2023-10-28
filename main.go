package main

import (
	"log"

	"github.com/alankritjoshi/coinwallet/handlers"
	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	// add custom handlers
	handlers.RegisterHandlers(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
