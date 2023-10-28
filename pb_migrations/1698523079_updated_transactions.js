/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("xiuzpu94ad1t4mh")

  // remove
  collection.schema.removeField("ffc6adrc")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "d4df1rsp",
    "name": "outputs",
    "type": "json",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {}
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("xiuzpu94ad1t4mh")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ffc6adrc",
    "name": "amount",
    "type": "number",
    "required": true,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "noDecimal": false
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "d4df1rsp",
    "name": "out",
    "type": "json",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {}
  }))

  return dao.saveCollection(collection)
})
