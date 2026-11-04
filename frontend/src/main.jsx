import { createRoot } from "react-dom/client";
import Dashboard from "./pages/Dashboard";

// Page registry: the name Go sends -> the React component
const pages = {
  Dashboard,
};

function readPage() {
  try {
    return JSON.parse(document.getElementById("page-data").textContent);
  } catch {
    return null;
  }
}

const page = readPage();
const Page = page && pages[page.component];

createRoot(document.getElementById("root")).render(
  Page ? (
    <Page {...page.props} />
  ) : (
    <p>Unknown page: {page?.component ?? "no page data"}</p>
  ),
);
