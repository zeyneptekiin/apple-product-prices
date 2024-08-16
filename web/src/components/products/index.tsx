import { useState, useEffect } from "react";
import ProductSlider from "@/components/products/slider";
import { getProductDetails } from "@/services/getProductDetails/getProductDetails";
import {getProductPrice} from "@/services/getProductPrice/getProductPrice";
import {getCurrency, CountryCode} from "@/services/getCurrency/getCurrency";
import {CurrencySymbol, getSymbol} from "@/services/getSymbol/getSymbol";

type Product = {
    product_name: string;
    name: string;
    images: string;
}

type ProductsProps = {
    data?: Product[];
    lang: string;
}

export default function Products({ data = [], lang }: ProductsProps) {
    const [searchQuery, setSearchQuery] = useState("");
    const [productDetails, setProductDetails] = useState<Map<string, any>>(new Map());
    const [currentData, setCurrentData] = useState("");

    useEffect(() => {
        async function fetchDetails() {
            const detailsMap = new Map<string, any>();
            for (const product of data) {
                try {
                    const details = await getProductDetails(product.product_name);
                    detailsMap.set(product.product_name, details);
                } catch (error) {
                    console.error(`Failed to fetch details for ${product.product_name}:`, error);
                }
            }
            setProductDetails(detailsMap);

            try {
                const currency = getCurrency(lang as CountryCode)
                const prices = await getProductPrice(currency);
                setCurrentData(prices?.conversion_rates);
            } catch (error) {
                console.error(`Failed to fetch prices`, error);
            }
        }

        fetchDetails();
    }, [data]);

    if (!Array.isArray(data)) {
        console.error('Expected data to be an array but received:', data);
        return <p>Invalid data format.</p>;
    }

    const filteredProducts = data.filter(product =>
        product.product_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        product.name.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <>
            {filteredProducts.length > 0 ? (
                filteredProducts.map((product) => {
                    const details = productDetails.get(product.product_name);

                    return (
                        <div key={product.product_name} className="grid grid-cols-7 gap-4 mt-16 justify-center items-center">
                            <div className="justify-center items-center">
                                <div className="w-20 h-20 mx-auto">
                                    <img src={product.images} alt={product.product_name} className="h-[100%] object-contain mx-auto"/>
                                </div>
                                <p className="text-center mt-2">{product.name}</p>
                            </div>
                            <div className="flex justify-center items-center">
                                {details?.price[lang] ? (details.price[lang][0].price + " " + getSymbol(lang as CurrencySymbol)) : "No price for this country"}
                            </div>
                            <div className="col-span-5 justify-center items-center">
                                <ProductSlider productName={product.product_name} currentData={currentData} lang={lang}/>
                            </div>
                        </div>
                    );
                })
            ) : (
                <p>No products found.</p>
            )}
        </>
    );
}
