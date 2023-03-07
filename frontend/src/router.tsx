import { Suspense, lazy } from "react";
import { Navigate } from "react-router-dom";
import { RouteObject } from "react-router";

import SidebarLayout from "src/layouts/SidebarLayout";
import BaseLayout from "src/layouts/BaseLayout";

import Home from "./pages/Home";
import Sheets from "./pages/Sheets";
import SheetDetail from "./pages/SheetDetail";
import { useAppSelector } from "./redux/store";
import { selectUserSlice } from "./redux/feature/user/slice";
import Login from "./pages/Login";
import Status404 from "./pages/Status404";
import SheetResult from "./pages/SheetResult";
import SheetResultRecord from "./pages/SheetResult/Detail";

// Pages

// const Overview = Loader(lazy(() => import("src/content/overview")));

const routes = (): RouteObject[] => {
  const { isSignedIn } = useAppSelector(selectUserSlice);
  if (!isSignedIn)
    return [
      {
        path: "/login",
        element: <Login />,
      },
      {
        path: "*",
        element: <Navigate to="/login" />,
      },
    ];
  return [
    {
      path: "",
      element: <SidebarLayout />,
      children: [
        {
          path: "/",
          element: <Navigate to="/trang-chu" replace />,
        },
        {
          path: "trang-chu",
          element: <Home />,
        },
        {
          path: "sheet",
          element: <Sheets />,
        },
        {
          path: "sheet/:id",
          element: <SheetDetail />,
        },
        {
          path: "ket-qua",
          element: <SheetResult />,
        },
        {
          path: "ket-qua/:id",
          element: <SheetResultRecord />,
        },
        {
          path: "*",
          element: <Status404 />,
        },
      ],
    },
  ];
};
export default routes;
