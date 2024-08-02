from typing import List, Dict
from pydantic import BaseModel, Field
import uuid


class ProductPrice(BaseModel):
    country_name: str
    price: float
    vat: float
    date: str


class Product(BaseModel):
    product_id: str = Field(default_factory=lambda: str(uuid.uuid4()))
    product_name: str
    price: Dict[str, List[ProductPrice]]
    name: str
    category: str

    class Config:
        json_encoders = {
            uuid.UUID: str
        }
