import { useState } from 'react';
import CountryModal from "@/components/navbar/countriesModal";
import countries from "../../json_files/countries.json"

export default function Navbar() {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const toggleModal = () => {
        setIsModalOpen(!isModalOpen);
    };

    return (
        <nav className="w-full h-16 fixed top-0 bg-orange-50 flex justify-center items-center z-10">
            <div className="w-[1440px] flex items-center justify-between px-8">
                <p className="text-xl font-medium text-blue-600">Global Price Tracking</p>
                <ul className="flex gap-10">
                    <li onClick={toggleModal} className="cursor-pointer">
                        <p className="text-xl font-medium text-blue-600">Select Country</p>
                        {isModalOpen && <CountryModal closeModal={toggleModal} countries={countries}/>}
                    </li>
                    <li className="text-xl font-medium text-blue-600">About</li>
                </ul>
            </div>
        </nav>
    )
}