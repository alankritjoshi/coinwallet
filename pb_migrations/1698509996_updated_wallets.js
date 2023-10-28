/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "yihmfppa",
    "name": "type",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "bitcoin"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  // remove
  collection.schema.removeField("yihmfppa")

  return dao.saveCollection(collection)
})
