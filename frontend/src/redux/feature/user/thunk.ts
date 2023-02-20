import { createAsyncThunk } from "@reduxjs/toolkit";
import { LoginRequest } from "../../../lib/auth/auth_pb";
import { authClient } from "../../../utils/grpc";

export const getMeThunk = createAsyncThunk("auth/getMe", async () => {
  return await authClient.getMe();
});

export const loginThunk = createAsyncThunk(
  "auth/loginThunk",
  async (param: LoginRequest.AsObject) => {
    return await authClient.logIn(param);
  }
);

export const logoutThunk = createAsyncThunk("auth/logout", async () => {
  return await authClient.logOut();
});
