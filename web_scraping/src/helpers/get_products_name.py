from bs4 import BeautifulSoup
import requests
from typing import List, Dict


def get_products_name() -> Dict[str, List[str]]:
    base_url = "https://www.apple.com"
    keywords = ["iphone", "ipad", "ipad-pro", "ipad/compare/", "ipad-air", "airpods-pro", "airpods-2nd-generation", "airpods-3rd-generation", "ipad-10.9", "iphone-15-pro", "iphone-15", "iphone-se", "ipad-mini", "macbook", "mac", "macbook-air", "macbook-pro", "imac", "mac-mini", "mac-studio", "displays", "airpods", "watch"]
    content_values = {}

    for keyword in keywords:
        url = f"{base_url}/{keyword}"
        try:
            response = requests.get(url)
            response.raise_for_status()

            if response.status_code == 200:
                soup = BeautifulSoup(response.text, 'html.parser')
                meta_tags = soup.find_all('meta', {'name': 'ac:pricing-alias'})
                feature_items_divs = soup.find_all('div', {'data-type': 'featureItems'})

                keyword_values = []

                for div in feature_items_divs:
                    inner_div = div.find('div', {'data-store-value': ''})
                    if inner_div:
                        text_value = inner_div.get_text(strip=True)
                        cleaned_text = text_value.strip('{}')
                        if 'IPAD' in cleaned_text:
                            keyword_values.append(cleaned_text)

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
