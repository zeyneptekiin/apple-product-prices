from helpers.get_products_name import get_products_name
from helpers.get_products_prices import get_products_prices


def main():
    product_names = get_products_name()

    country_file = '/app/countries.txt'
    get_products_prices(country_file, product_names)


if __name__ == '__main__':
    main()