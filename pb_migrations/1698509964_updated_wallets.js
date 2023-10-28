/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "x9hkxkro",
    "name": "label",
    "type": "text",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": 15,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ome3vnzyajnf2e5")

  // remove
  collection.schema.removeField("x9hkxkro")

  return dao.saveCollection(collection)
})
