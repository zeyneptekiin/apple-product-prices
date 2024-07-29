from typing import TypedDict
import uuid


class ProductDict(TypedDict):
    product_id: str
    product_name: str
    name: str
    price: float
    country: str
    category: str


def create_product(product_name: str, name: str, country: str, price: float, category: str) -> ProductDict:
    return {
        "product_id": str(uuid.uuid4()),
        "product_name": product_name,
        "name": name,
        "price": price,
        "country": country,
        "category": category
    }
