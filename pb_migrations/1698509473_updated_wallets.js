/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  // remove
  collection.schema.removeField("syvikb87")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "syvikb87",
    "name": "wallet_id",
    "type": "text",
    "required": true,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": 20,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
})
