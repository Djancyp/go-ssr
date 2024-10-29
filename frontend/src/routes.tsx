import { Routes, Route } from "react-router-dom";
import Home from "./Pages/Home";
import About from "./Pages/About";
import Params from "./Pages/Params";
import Layout from "./layout";

function R() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/about/:id" element={<Params />} />
      </Route>
    </Routes>
  );
}

export default R;
