import { createAsyncThunk } from "@reduxjs/toolkit";
import {
  DeleteRecordRequest,
  ListRecordsRequest,
  ListSheetsRequest,
} from "src/lib/lodas/lodas_pb";
import { lodasClient } from "src/utils/grpc";

export const listSheetsThunk = createAsyncThunk(
  "lodas/listSheets",
  async (param: ListSheetsRequest.AsObject) => {
    return await lodasClient.listSheets(param);
  }
);

export const listRecordsThunk = createAsyncThunk(
  "lodas/listRecords",
  async (param: ListRecordsRequest.AsObject) => {
    return await lodasClient.listRecords(param);
  }
);

export const deleteRecordThunk = createAsyncThunk(
  "lodas/deleteRecord",
  async (param: DeleteRecordRequest.AsObject) => {
    return await lodasClient.deleteRecord(param);
  }
);
