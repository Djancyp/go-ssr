import { StrictMode } from "react";
import { StaticRouter } from "react-router-dom/server";
import { renderToString } from "react-dom/server";
import Routes from "./routes";

export function render(url: any) {
  if (!url) {
    throw new Error("No NO url is required");
  }
  const html = renderToString(
    <StrictMode>
      <StaticRouter location={url}>
        <Routes />
      </StaticRouter>
    </StrictMode>
  );
  return { html };
}
