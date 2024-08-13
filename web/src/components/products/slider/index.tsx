import { useEffect, useState } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';

import 'swiper/css';
import 'swiper/css/free-mode';
import 'swiper/css/pagination';

import { Pagination } from 'swiper/modules';
import { getProductDetails } from '@/services/getProductDetails/getProductDetails';
import {getKeyText, CountryCode} from "@/services/getKeyText/getKeyText";

type PriceEntry = {
    price: number;
    vat: number;
    date: string;
}

type ProductData = {
    id: string;
    product_id: string;
    product_name: string;
    price: {
        [country: string]: PriceEntry[];
    };
    name: string;
    category: string;
}

type ProductSliderProps = {
    productName: string;
}

export default function ProductSlider({ productName }: ProductSliderProps) {
    const [productData, setProductData] = useState<ProductData | null>(null);

    useEffect(() => {
        const fetchProductData = async () => {
            try {
                const data = await getProductDetails(productName);
                setProductData(data);
            } catch (error) {
                console.error('Error fetching product details:', error);
            }
        };

        fetchProductData();
    }, [productName]);

    return (
        <>
            <Swiper
                slidesPerView={5}
                spaceBetween={30}
                modules={[Pagination]}
                className="mySwiper"
            >
                {productData?.price && Object.entries(productData.price).map(([country, entries]) => {
                    if (entries.length === 0) return null;

                    const firstEntry = entries[0];

                    return (
                        <SwiperSlide key={`${country}-${firstEntry.date}`}>
                            <div className="text-center">
                                <p>{getKeyText(country as CountryCode)}</p>
                                <p className="mt-3">{firstEntry.price.toLocaleString()}</p>
                            </div>
                        </SwiperSlide>
                    );
                })}
            </Swiper>
        </>
    );
}
