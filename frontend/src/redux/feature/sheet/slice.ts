import { createSlice } from "@reduxjs/toolkit";
import { Record, SheetInfo } from "src/lib/lodas/lodas_pb";
import { CommonState } from "src/redux/common/types";
import { StatusEnum } from "src/redux/constant";
import { listRecordsThunk, listSheetsThunk } from "./thunk";

type SheetState = {
  sheetList: SheetInfo.AsObject[];
  currentRecordList: Record.AsObject[];
} & CommonState;

export const initialState: SheetState = {
  sheetList: [],
  currentRecordList: [],
  status: StatusEnum.LOADING,
};

const slice = createSlice({
  name: "sheetList",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(listSheetsThunk.pending, (state) => {
      state.status = StatusEnum.LOADING;
    });
    builder.addCase(listRecordsThunk.pending, (state) => {
      state.status = StatusEnum.LOADING;
    });
    builder.addCase(listSheetsThunk.fulfilled, (state, { payload }) => {
      const { response, error } = payload;
      if (error) {
        state.status = StatusEnum.SUCCEEDED;
        return;
      }
      state.sheetList = response?.sheetsList || [];
      state.status = StatusEnum.SUCCEEDED;
    });
    builder.addCase(listRecordsThunk.fulfilled, (state, { payload }) => {
      const { response, error } = payload;
      if (error) {
        state.status = StatusEnum.SUCCEEDED;
        return;
      }
      state.currentRecordList = response?.recordsList || [];
      state.status = StatusEnum.SUCCEEDED;
    });
  },
});
