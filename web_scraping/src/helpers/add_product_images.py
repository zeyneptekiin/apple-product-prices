from pymongo import MongoClient
from typing import Dict


def add_product_images(products: Dict[str, str], product_name: str) -> None:
    uri = "mongodb://zeyneptekin:123456@mongodb:27017/"
    client = MongoClient(uri)
    db = client['apple']
    collection = db['products']

    image_url = products.get(product_name)

    if image_url:
        result = collection.update_one(
            {'product_name': product_name},
            {'$set': {'images': image_url}}
        )
        if result.matched_count:
            print(f"Updated {product_name} with image URL: {image_url}")
        else:
            print(f"No matching product found for {product_name}")
    else:
        print(f"No image URL found for {product_name}")
