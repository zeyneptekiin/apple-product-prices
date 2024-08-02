from .helpers.get_products_name import get_products_name
from .helpers.create_product_structure import create_product_structure
from .helpers.add_product_prices import add_product_prices


def main():
    product_names = get_products_name()

    create_product_structure(product_names)

    add_product_prices(product_names)


if __name__ == '__main__':
    main()
