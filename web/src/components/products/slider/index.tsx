import React, { useRef, useState } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';

import 'swiper/css';
import 'swiper/css/free-mode';
import 'swiper/css/pagination';

import {Pagination } from 'swiper/modules';

export default function ProductSlider() {
    return (
        <>
            <Swiper
                slidesPerView={5}
                spaceBetween={30}
                modules={[Pagination]}
                className="mySwiper"
            >
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
                <SwiperSlide>
                    <p className="text-center">Country</p>
                    <p className="text-center mt-3">Price</p>
                </SwiperSlide>
            </Swiper>
        </>
    );
}
