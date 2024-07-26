from pymongo import MongoClient, errors

uri = "mongodb://zeyneptekin:123456@mongodb:27017"
client = MongoClient(uri)

try:
    database = client["apple"]
    products = database["products"]

    product = {
        "product_id": "12345",
        "name": "iphone_14_plus",
        "prices": {
            "usa": "7000",
            "eur": "8000",
            "tr": "10000"
        },
        "category": "iphone"
    }

    result = products.insert_one(product)

    client.close()

except Exception as e:
    raise Exception("Error: ", e)


