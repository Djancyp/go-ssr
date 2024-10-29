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
