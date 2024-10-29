import { useState } from "react";
import reactLogo from "../assets/react.svg";
import GolangLogo from "../assets/go.svg";
import LuanLogo from "../assets/luna.svg";
import EchoLogo from "../assets/echo.svg";
function Home() {
  const { name } = globalThis.props[window.location.pathname];
  const [count, setCount] = useState(0);
  return (
    <>
      <div className="flex justify-center items-center flex-col">
        <img
          width={"200px"}
          src={LuanLogo}
          className="logo react"
          alt="React logo"
        />
        <h1 className="text-[42px] text-white">Luna</h1>
      </div>
      <div className="flex justify-center items-center">
        <a href="https://go.dev/" target="_blank" className="mr-5">
          <img
            width={"100px"}
            src={GolangLogo}
            className="logo react"
            alt="React logo"
          />
        </a>
        <a href="https://reactjs.org" target="_blank" className="mr-5">
          <img
            width={"120px"}
            src={EchoLogo}
            className="logo react"
            alt="React logo"
          />
        </a>
        <a href="https://reactjs.org" target="_blank" className="mr-5">
          <img
            width={"60px"}
            src={reactLogo}
            className="logo react"
            alt="React logo"
          />
        </a>
      </div>
      <div className="flex justify-center items-center flex-col">
        <div className="flex justify-center flex-col items-center">
          <button
            className="px-4 py-2 bg-blue-600 hover:bg-blue-500 rounded max-w-[200px] mb-5"
            onClick={() => setCount((count) => count + 1)}
          >
            count is {count}
          </button>
          <p>Dynamic content from the props: {name}</p>
        </div>
        <p className="read-the-docs">
          Click on the Go and React logos to learn more
        </p>
      </div>
    </>
  );
}
export default Home;
