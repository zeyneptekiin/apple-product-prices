from typing import TypedDict
import uuid


class ProductDict(TypedDict):
    product_id: str
    name: str
    price: float
    country: str
    category: str


def create_product(name: str, country: str, price: float, category: str) -> ProductDict:
    return {
        "product_id": str(uuid.uuid4()),
        "name": name,
        "price": price,
        "country": country,
        "category": category
    }
