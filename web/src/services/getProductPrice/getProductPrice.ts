export const getProductPrice = async (currency: string) => {
    const res = await fetch(`http://localhost:8080/exchange?base=${currency}`);

    return await res.json();
};