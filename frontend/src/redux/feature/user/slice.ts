import { createSlice } from "@reduxjs/toolkit";
import { KEY_ACCESS_TOKEN } from "../../../utils/grpc";
import { CommonState } from "../../common/types";
import { StatusEnum } from "../../constant";
import { RootState } from "../../reducers";
import { getMeThunk, loginThunk, logoutThunk } from "./thunk";

type UserState = {
  isSignedIn: boolean;
  username: string;
  displayName: string;
  email: string;
  phoneNumber: string;
} & CommonState;

export const initialState: UserState = {
  isSignedIn: false,
  status: StatusEnum.LOADING,
  email: "",
  username: "",
  displayName: "",
  phoneNumber: "",
};

const slice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(getMeThunk.pending, (state) => {
      state.status = StatusEnum.LOADING;
    });
    builder.addCase(getMeThunk.fulfilled, (state, { payload }) => {
      const { response, error } = payload;
      if (error) {
        state.status = StatusEnum.SUCCEEDED;
        state.isSignedIn = false;
        return;
      }
      state.displayName = response?.user?.displayName || "";
      state.email = response?.user?.email || "";
      state.phoneNumber = response?.user?.phoneNumber || "";
      state.username = response?.user?.username || "";

      state.status = StatusEnum.SUCCEEDED;
      state.isSignedIn = true;
    });
    builder.addCase(loginThunk.fulfilled, (state, { payload }) => {
      const { response, error } = payload;
      if (error) {
        state.isSignedIn = false;
        return;
      }
      const token = response?.accessToken || "";
      localStorage.setItem(KEY_ACCESS_TOKEN, token);
      state.isSignedIn = true;
    });
    builder.addCase(logoutThunk.fulfilled, (state, { payload }) => {
      state.isSignedIn = false;
      localStorage.removeItem(KEY_ACCESS_TOKEN);
    });
  },
});

export const selectUserSlice = (state: RootState) => state.user;
export const isUserSliceLoading = (state: RootState) =>
  state.user.status === StatusEnum.LOADING;

export default slice.reducer;
