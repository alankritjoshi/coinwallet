package handlers

import "github.com/pocketbase/pocketbase"

func RegisterHandlers(app *pocketbase.PocketBase) {
	registerSyncHandler(app)
}
