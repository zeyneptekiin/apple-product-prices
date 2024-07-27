from pymongo import MongoClient, errors
import requests
from bs4 import BeautifulSoup
from product_structure import create_product

uri = "mongodb://zeyneptekin:123456@mongodb:27017"
client = MongoClient(uri)

try:
        database = client["apple"]
        products = database["products"]

        URL = "https://www.apple.com/tr/iphone/"
        response = requests.get(URL)
        soup = BeautifulSoup(response.content, 'html.parser')

        index = 0
        while True:
            try:
                nameSelector = f'#gallery-item-{index} > div:nth-of-type(1) > div:nth-of-type(1) > h3 > p'
                priceSelector = f'#gallery-item-{index} > p:nth-of-type(2) > span'

                elementName = soup.select_one(nameSelector)
                elementPrice = soup.select_one(priceSelector)
                print(elementPrice)

                if not elementName or not elementPrice:
                    break

                product_name = elementName.get_text(strip=True)
                product_price = elementPrice.get_text(strip=True)
                product = create_product(product_name, "tr", product_price)

                result = products.insert_one(product)
                index += 1

            except Exception as e:
                print(f"Error: {e}")
                break

        client.close()

except Exception as e:
    print(f"MongoDB connection error: {e}")


