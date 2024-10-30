type Props = {
  [key: string]: {
    product: {
      title: string;
    };
  };
};

// Extend the global scope to include `props` with the `Props` type
declare global {
  interface Window {
    props: Props;
  }
  // Declare props on globalThis
  var props: Props;
}
function Params() {
  const product = globalThis.props[window.location.pathname]?.product;
  return (
    <>
      <h1>Product</h1>

      {product.title}
    </>
  );
}

export default Params;
