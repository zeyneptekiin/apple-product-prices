import uuid


def create_product(name, country, price, category):
    return {
        "product_id": str(uuid.uuid4()),
        "name": name,
        "prices": price,
        "country": country,
        "category": category
    }
