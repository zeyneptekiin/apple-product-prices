import Navbar from '../components/navbar'
import Footer from '../components/footer'

// @ts-ignore
export default function Layout({ children }: LayoutProps) {
  return (
      <>
        <Navbar />
        <main>{children}</main>
        <Footer />
      </>
  )
}