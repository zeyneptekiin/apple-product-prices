import requests
from pymongo import MongoClient
from pymongo.errors import PyMongoError
from bs4 import BeautifulSoup
from product_structure import create_product


uri = "mongodb://zeyneptekin:123456@mongodb:27017"
client = MongoClient(uri)


def fetch_pricing_aliases():
    base_url = "https://www.apple.com/"
    keywords = ["iphone", "ipad", "macbook", "airpods"]
    content_values = {}

    for keyword in keywords:
        url = f"{base_url}{keyword}"
        try:
            response = requests.get(url)
            response.raise_for_status()  # Raise an exception for HTTP errors

            if response.status_code == 200:
                soup = BeautifulSoup(response.text, 'html.parser')
                meta_tags = soup.find_all('meta', {'name': 'ac:pricing-alias'})

                keyword_values = []
                for meta in meta_tags:
                    content = meta.get('content')
                    if content:
                        value = content.split('=')[-1]
                        keyword_values.append(value)
                content_values[keyword] = keyword_values
            else:
                print(f"Failed to retrieve data from {url}. Status code: {response.status_code}")
        except requests.RequestException as e:
            print(f"Request failed for {url}: {e}")

    return content_values


def fetch_and_store_data(country_file, content_values):
    db = client['apple']
    collection = db['products']

    for keyword, values in content_values.items():
        print(f"Processing keyword: {keyword}")
        with open(country_file, 'r') as file:
            country_codes = file.read().splitlines()

        for country_code in country_codes:
            parts = ','.join(values)
            url = f"https://www.apple.com/{country_code}/shop/mcm/product-price?parts={parts}"

            try:
                response = requests.get(url)
                response.raise_for_status()  # Raise an exception for HTTP errors

                print(f"Fetching data from {url}")

                result_data = response.json()
                print(f"Parsed JSON data from {url}: {result_data}")

                items = result_data.get('items', {})
                if not isinstance(items, dict):
                    print("Unexpected response format.")
                    continue

                for product_id, item in items.items():
                    if isinstance(item, dict) and 'price' in item:
                        product = create_product(
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
            except requests.RequestException as e:
                print(f"Request failed for {url}: {e}")
            except ValueError as e:
                print(f"Error parsing JSON from {url}: {e}")


def main():
    content_values = fetch_pricing_aliases()
    print("Fetched content values:", content_values)

    country_file = '/app/countries.txt'
    fetch_and_store_data(country_file, content_values)


if __name__ == "__main__":
    main()
