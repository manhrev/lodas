import { createSlice } from "@reduxjs/toolkit";
import { RootState } from "../../reducers";

interface CommonState {
  data: string;
}

const initialState: CommonState = {
  data: "heelo",
};

const slice = createSlice({
  name: "common",
  initialState,
  reducers: {
    setData: (state: CommonState) => {
      state.data = "set";
    },
  },
});

export const { setData } = slice.actions;

export const selectCommonSlice = (state: RootState) => state.common;

export default slice.reducer;
