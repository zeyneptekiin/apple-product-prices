import Layout from "@/pages/layout";
import Products from "@/components/products";

export default function Home() {
  return (
      <Layout>
          <section className="flex justify-center items-center mt-16">
              <div className="w-[1440px] px-8">
                  <div className="grid grid-cols-7 gap-4 mt-10 justify-center items-center text-lg text-orange-500">
                      <div className="flex justify-center items-center bg-blue-50 rounded-2xl py-2">Product Info</div>
                      <div className="flex justify-center items-center bg-blue-50 rounded-2xl py-2">Türkiye</div>
                      <div className="col-span-5 bg-blue-50 pl-16 rounded-2xl py-2">Countries</div>
                  </div>
                  <div className="bg-black opacity-10 w-full h-[1px] mt-4"/>
                  <Products/>
              </div>
          </section>
      </Layout>
  );
}
