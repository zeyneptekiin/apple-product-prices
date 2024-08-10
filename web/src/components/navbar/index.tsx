export default function Navbar() {
    return (
        <nav className="w-full h-16 fixed top-0 bg-orange-50 flex justify-center z-10">
        <div className="w-[1440px] flex items-center px-8 justify-between text-xl font-medium text-blue-600">
            <p>Global Price Tracking</p>
            <ul className="flex gap-3">
                <li>About</li>
            </ul>
        </div>
        </nav>
    )
}