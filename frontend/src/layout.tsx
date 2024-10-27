import { Outlet, Link } from "react-router-dom";
function Layout() {
  return (
    <div>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/app">About</Link>
        </li>
      </ul>
      <h1>App</h1>
      <Outlet />
    </div>
  );
}

export default Layout;
