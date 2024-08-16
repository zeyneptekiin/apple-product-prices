import os

from pymongo import MongoClient
from typing import Dict
from models.product_model import Product
from helpers.add_product_images import add_product_images
from helpers.product_images_data import products


def create_database(items: Dict[str, dict], keyword: str):
    uri = "mongodb://zeyneptekin:123456@mongodb:27017/"
    client = MongoClient(uri)
    db = client['apple']
    collection = db['products']
    country_file = os.path.join('/app', 'countries.txt')

    with open(country_file, 'r') as file:
        country_codes = file.read().splitlines()

    for product_id, item in items.items():
        price_dict = {country_code: [] for country_code in country_codes}

        product = Product(
            product_name=item.get('id', ''),
            name=item.get('baseName', ''),
            price=price_dict,
            category=keyword,
        )

        existing_product = collection.find_one({"product_name": product.product_name})

        if existing_product:
            print(f"Product already exists: {product.model_dump()}")
        else:
            print(f"Inserting product: {product.model_dump()}")
            result = collection.insert_one(product.model_dump())
            print(f"Insert result: {result.inserted_id}")

            add_product_images(products, product.product_name)
