import requests
from .create_database import create_database


def get_products_prices(country_file, content_values):
    for keyword, values in content_values.items():
        with open(country_file, 'r') as file:
            country_codes = file.read().splitlines()

        for country_code in country_codes:
            parts = ','.join(values)
            url = f"https://www.apple.com/{country_code}/shop/mcm/product-price?parts={parts}"

            try:
                response = requests.get(url)
                response.raise_for_status()

                print(f"Fetching data from {url}")

                result_data = response.json()
                print(f"Parsed JSON data from {url}: {result_data}")

                items = result_data.get('items', {})
                if not isinstance(items, dict):
                    print("Unexpected response format.")
                    continue

                create_database(items, country_code, keyword)

            except requests.RequestException as e:
                print(f"Request failed for {url}: {e}")
            except ValueError as e:
                print(f"Error parsing JSON from {url}: {e}")
