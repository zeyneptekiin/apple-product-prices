export const getProductDetails = async (product_name: string) => {
    const res = await fetch(`http://localhost:8080/products?product_name=${product_name}`);

    return await res.json();
};
