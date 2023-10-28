/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  collection.listRule = "@request.auth.id != id"
  collection.viewRule = "@request.auth.id != id"
  collection.createRule = "@request.auth.id != id"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  collection.listRule = null
  collection.viewRule = null
  collection.createRule = null

  return dao.saveCollection(collection)
})
