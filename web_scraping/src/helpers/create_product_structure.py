import requests
from typing import List, Dict
from helpers.create_database import create_database


def create_product_structure(content_values: Dict[str, List[str]]):
    for keyword, values in content_values.items():
        parts = ','.join(values)
        url = f"https://www.apple.com/us/shop/mcm/product-price?parts={parts}"

        try:
            response = requests.get(url)
            response.raise_for_status()

            result_data = response.json()

            items = result_data.get('items', {})
            if not isinstance(items, dict):
                print("Unexpected response format.")
                continue

            create_database(items, keyword)

        except requests.RequestException as e:
            print(f"Request error for URL {url}: {e}")
        except ValueError as e:
            print(f"JSON decoding error for URL {url}: {e}")