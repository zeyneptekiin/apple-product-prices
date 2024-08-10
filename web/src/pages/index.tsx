import Layout from "@/pages/layout";
import Products from "@/components/products";
import {getProductsName} from "@/services/getProductsName/getProductsName";

type HomeProps = {
    productsInfo: {
        product_name: string;
        name: string;
    }[];
}

export default function Home({productsInfo}: HomeProps) {
  return (
      <Layout>
          <section className="flex justify-center items-center mt-16">
              <div className="w-[1440px] px-8">
                  <div className="grid grid-cols-7 gap-4 sticky top-0 pt-20 pb-4 bg-white z-[5] justify-center items-center text-lg text-orange-500">
                      <div className="flex justify-center items-center bg-blue-50 rounded-2xl py-2">Product Info</div>
                      <div className="flex justify-center items-center bg-blue-50 rounded-2xl py-2">Türkiye</div>
                      <div className="col-span-5 bg-blue-50 pl-16 rounded-2xl py-2">Countries</div>
                  </div>
                  <div className="bg-black opacity-10 w-full h-[1px] mt-4"/>
                  <Products data={productsInfo}/>
              </div>
          </section>
      </Layout>
  );
}

export async function getServerSideProps() {
    const productsInfo = await getProductsName();

    return { props: { productsInfo } }
}
