import os

from pymongo import MongoClient
import requests
from typing import List, Dict
from datetime import datetime


def add_product_prices(category_products: Dict[str, List[str]]) -> None:
    uri = "mongodb://zeyneptekin:123456@mongodb:27017/"  # TODO: Move to .env file
    client = MongoClient(uri)
    db = client['apple']
    collection = db['products']
    country_file = os.path.join('/app', 'countries.txt')

    all_product_names = [product_name for product_list in category_products.values() for product_name in product_list]
    parts = ','.join(all_product_names)
    print(f"Product names concatenated: {parts}")

    with open(country_file, 'r') as file:
        country_codes = file.read().splitlines()

    for country_code in country_codes:
        url = f"https://www.apple.com/{country_code}/shop/mcm/product-price?parts={parts}"

        try:
            response = requests.get(url)
            response.raise_for_status()

            print(f"Fetching data from {url}")

            result_data = response.json()
            print(f"Parsed JSON data from {url}")

            items = result_data.get('items', {})
            if not isinstance(items, dict):
                print("Unexpected response format.")
                continue

            today_date = datetime.now().strftime('%Y-%m-%d')

            for product_id, item in items.items():
                price_info = item.get('price', {})
                price = price_info.get('value', 0.0)
                vat = 0.0

                collection.update_one(
                    {'product_name': product_id},
                    {'$push': {
                        f'price.{country_code}': {
                            'price': price,
                            'vat': vat,
                            'date': today_date
                        }
                    }}
                )

        except requests.RequestException as e:
            print(f"Request failed for {url}: {e}")
        except ValueError as e:
            print(f"Error parsing JSON from {url}: {e}")
