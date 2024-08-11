import os
from bs4 import BeautifulSoup
from pymongo import MongoClient
import requests
from typing import Dict


def add_product_images(name: str, category: str) -> None:
    uri = "mongodb://zeyneptekin:123456@mongodb:27017/"
    client = MongoClient(uri)
    db = client['apple']
    collection = db['products']

    base_url = "https://www.apple.com"
    url = f"{base_url}/{category}"
    print(f"this is : {url}")
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'html.parser')

    image_url = None

    for li in soup.find_all('li', class_='product-tile'):
        aria_label = li.get('aria-label')
        if aria_label == name:
            img_tag = li.find('img')
            if img_tag:
                src = img_tag.get('src')
                if src:
                    image_url = f"{base_url}{src}"
                    break

    if image_url:
        collection.update_one(
            {'name': name},
            {'$set': {
                'images': image_url
            }}
        )
