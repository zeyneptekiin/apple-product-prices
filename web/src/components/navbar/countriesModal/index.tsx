import React from 'react';
import Link from "next/link";

type CountryModalProps = {
    closeModal: () => void;
    countries: { [key: string]: string };
}

export default function CountryModal({ countries }: CountryModalProps) {

    return (
        <div className="fixed flex mt-5">
            <div className="bg-white p-3 rounded shadow-xl">
                <div className="max-h-52 max-w-40 flex">
                <ul className="overflow-x-auto">
                    {Object.entries(countries).map(([code, name]) => (
                        <li key={code} className="py-1">
                            <Link
                                href={{
                                    pathname: `/${code}`,
                                }}
                            >{name}
                            </Link>
                        </li>
                    ))}
                </ul>
                </div>
            </div>
        </div>
    );
}
