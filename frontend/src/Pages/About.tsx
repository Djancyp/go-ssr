import { Link } from "../Link";
function About() {
  const products = globalThis.props[window.location.pathname]?.products;

  return (
    <div>
      <h1>About</h1>
      <p>This is the about page.</p>
      {products.length > 0 && (
        <ul>
          {products.map((product: any) => (
            <li key={product.id}>
              <Link to={"/about/" + product.id}>{product.title}</Link>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
export default About;
