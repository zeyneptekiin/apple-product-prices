from pymongo import MongoClient, errors
import requests
from bs4 import BeautifulSoup
from product_structure import create_product

uri = "mongodb://zeyneptekin:123456@mongodb:27017"
client = MongoClient(uri)

try:
    database = client["apple"]
    products = database["products"]

    with open('/app/countries.txt', 'r') as file:
        country_codes = file.read().splitlines()

    for country_code in country_codes:
        try:
            URL = f"https://www.apple.com/{country_code}/iphone/"
            response = requests.get(URL)
            soup = BeautifulSoup(response.content, 'html.parser')

            index = 0
            while True:
                try:
                    nameSelector = f'#gallery-item-{index} > div:nth-of-type(1) > div:nth-of-type(1) > h3 > p'
                    priceSelector = f'#gallery-item-{index} > p:nth-of-type(2) > span'

                    elementName = soup.select_one(nameSelector)
                    elementPrice = soup.select_one(priceSelector)

                    if not elementName or not elementPrice:
                        break

                    product_name = elementName.get_text(strip=True)
                    product_price = elementPrice.get_text(strip=True)

                    existing_product = products.find_one({
                        "name": product_name,
                        "country": country_code,
                        "category": "iphone"
                    })

                    if existing_product:

                        if existing_product["prices"] != product_price:
                            products.update_one(
                                {"_id": existing_product["_id"]},
                                {"$set": {"prices": product_price}}
                            )
                            print(f"Updated product: {product_name} in {country_code}")
                    else:
                        product = create_product(product_name, country_code, product_price)
                        products.insert_one(product)
                        print(f"Inserted new product: {product_name} in {country_code}")

                    index += 1

                except Exception as e:
                    print(f"Error while processing index {index} for country {country_code}: {e}")
                    break

        except requests.RequestException as e:
            print(f"Error fetching URL for country {country_code}: {e}")

    client.close()

except errors.PyMongoError as e:
    print(f"MongoDB connection error: {e}")

except FileNotFoundError as e:
    print(f"File error: {e}")
