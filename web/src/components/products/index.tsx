import ProductSlider from "@/components/products/slider";

interface ProductsProps {
    data: {
        product_name: string;
        name: string;
    }[];
}

export default function Products({ data }: ProductsProps) {
    return (
        <>
            {data.map((product) => (
                <div key={product.product_name} className="grid grid-cols-7 gap-4 mt-16 justify-center items-center">
                    <div className="justify-center items-center">
                        <div className="bg-slate-400 w-20 h-20 mx-auto"></div>
                        <p className="text-center mt-2">{product.name}</p>
                    </div>
                    <div className="flex justify-center items-center">Price</div>
                    <div className="col-span-5 justify-center items-center">
                        <ProductSlider productName={product.product_name} />
                    </div>
                </div>
            ))}
        </>
    );
}
