import { useRoutes } from "react-router-dom";
import router from "src/router";

import AdapterDateFns from "@mui/lab/AdapterDateFns";
import LocalizationProvider from "@mui/lab/LocalizationProvider";

import { CssBaseline } from "@mui/material";
import ThemeProvider from "./theme/ThemeProvider";
import { useAppDispatch, useAppSelector } from "./redux/store";
import { useEffect } from "react";
import { getMeThunk } from "./redux/feature/user/thunk";
import {
  isUserSliceLoading,
  selectUserSlice,
} from "./redux/feature/user/slice";
import "react-toastify/dist/ReactToastify.css";
import { ToastContainer } from "react-toastify";

function App() {
  const content = useRoutes(router());
  const dispatch = useAppDispatch();
  const loading = useAppSelector(isUserSliceLoading);
  const { isSignedIn } = useAppSelector(selectUserSlice);
  useEffect(() => {
    dispatch(getMeThunk());
  }, [isSignedIn]);

  return (
    <ThemeProvider>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <CssBaseline />
        {loading ? <></> : content}
        <ToastContainer position="top-right" autoClose={2000} />
      </LocalizationProvider>
    </ThemeProvider>
  );
}
export default App;
