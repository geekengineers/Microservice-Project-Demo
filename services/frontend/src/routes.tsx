import { createBrowserRouter } from "react-router-dom";
import { HomeComponent } from "./pages/home";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <HomeComponent />
    }
])