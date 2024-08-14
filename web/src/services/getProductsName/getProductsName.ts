export const getProductsName = async () => {
    const res = await fetch(`http://api:8080/products/name`);

    return await res.json();
};
