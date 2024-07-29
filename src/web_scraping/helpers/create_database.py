from pymongo import MongoClient
from pymongo.errors import PyMongoError
from .create_product import create_product


def create_database(items, country_code, keyword):
    uri = "mongodb://zeyneptekin:123456@mongodb:27017"
    client = MongoClient(uri)
    db = client['apple']
    collection = db['products']

    for product_id, item in items.items():
        if isinstance(item, dict) and 'price' in item:
            product = create_product(
                product_name=item.get('id', ''),
                name=item.get('name', ''),
                country=country_code,
                price=item['price'].get('value', 0),
                category=keyword
            )
            try:
                existing_product = collection.find_one({
                    "name": product["name"],
                    "country": product["country"],
                    "category": product["category"]
                })

                if existing_product:
                    if existing_product["price"] != product["price"]:
                        collection.update_one(
                            {"_id": existing_product["_id"]},
                            {"$set": {"price": product["price"]}}
                        )
                        print(f"Updated product: {product['name']} in {country_code}")
                else:
                    collection.insert_one(product)
                    print(f"Inserted new product: {product['name']} in {country_code}")

            except PyMongoError as e:
                print(f"MongoDB operation failed: {e}")
