import { createSlice } from "@reduxjs/toolkit";
import { Record, SheetInfo } from "src/lib/lodas/lodas_pb";
import { CommonState } from "src/redux/common/types";
import { StatusEnum } from "src/redux/constant";
import { RootState } from "src/redux/reducers";
import { deleteRecordThunk, listRecordsThunk, listSheetsThunk } from "./thunk";

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
    builder
      .addCase(listSheetsThunk.pending, (state) => {
        state.status = StatusEnum.LOADING;
      })
      .addCase(listRecordsThunk.pending, (state) => {
        state.status = StatusEnum.LOADING;
      })
      .addCase(listSheetsThunk.fulfilled, (state, { payload }) => {
        const { response, error } = payload;
        if (error) {
          state.status = StatusEnum.SUCCEEDED;
          return;
        }
        state.sheetList = response?.sheetsList || [];
        state.status = StatusEnum.SUCCEEDED;
      })
      .addCase(listRecordsThunk.fulfilled, (state, { payload }) => {
        const { response, error } = payload;
        if (error) {
          state.status = StatusEnum.SUCCEEDED;
          return;
        }
        state.currentRecordList = response?.recordsList || [];
        state.status = StatusEnum.SUCCEEDED;
      })
      .addCase(deleteRecordThunk.fulfilled, (state, { payload, meta }) => {
        const { error } = payload;
        if (!error) {
          state.currentRecordList = state.currentRecordList.filter((value) => {
            return value.id !== meta.arg.idsList[0];
          });
        }
      });
  },
});

export const selectSheetSlice = (state: RootState) => state.sheet;
export const isSheetSliceLoading = (state: RootState) =>
  state.sheet.status === StatusEnum.LOADING;

export default slice.reducer;
