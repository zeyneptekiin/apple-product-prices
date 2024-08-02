from pymongo import MongoClient
from typing import Dict
from web_scraping.src.models.product_model import Product


def create_database(items: Dict[str, dict], keyword: str):
    uri = "mongodb://zeyneptekin:123456@mongodb:27017/"  # TODO: Move to .env file
    client = MongoClient(uri)
    db = client['apple']
    collection = db['products']
    country_file = 'countries.txt'

    with open(country_file, 'r') as file:
        country_codes = file.read().splitlines()

    for product_id, item in items.items():
        price_dict = {country_code: [] for country_code in country_codes}

        product = Product(
            product_name=item.get('id', ''),
            name=item.get('name', ''),
            price=price_dict,
            category=keyword,
        )

        print(f"Inserting product: {product.product.model_dump()}")
        result = collection.insert_one(product.product.model_dump())
        print(f"Insert result: {result.inserted_id}")
