from bs4 import BeautifulSoup
import requests
from typing import List, Dict


def get_products_name() -> Dict[str, List[str]]:
    base_url = "https://www.apple.com"
    keywords = ["iphone", "ipad", "macbook", "airpods", "watch"]
    content_values = {}

    for keyword in keywords:
        url = f"{base_url}/{keyword}"
        try:
            response = requests.get(url)
            response.raise_for_status()

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
