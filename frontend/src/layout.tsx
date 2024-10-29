import { Outlet } from "react-router-dom";
import { Link } from "./Link";
function Layout() {
  return (
    <div className="bg-black text-white h-screen overflow-hidden">
      <ul className="flex justify-center border border-gray-800 h-[50px] items-center font-bold">
        <li>
          <Link className="px-4 py-2" to="/">
            Home
          </Link>
        </li>
        <li>
          <Link to="/about">About</Link>
        </li>
      </ul>
      <Outlet />
    </div>
  );
}

export default Layout;
