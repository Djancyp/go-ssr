import { Routes, Route } from "react-router-dom";
import App from "./App";
import Home from "./Pages/Home";
import Layout from "./layout";

function R() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route path="/" element={<Home />} />
        <Route path="/app" element={<App />} />
      </Route>
    </Routes>
  );
}

export default R;
