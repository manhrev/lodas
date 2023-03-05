import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import { LodasClient } from "src/lib/lodas/LodasServiceClientPb";
import {
  CreateRecordReply,
  CreateRecordRequest,
  CreateSheetReply,
  CreateSheetRequest,
  DeleteRecordReply,
  DeleteRecordRequest,
  DeleteSheetReply,
  DeleteSheetRequest,
  ListRecordsReply,
  ListRecordsRequest,
  ListSheetsReply,
  ListSheetsRequest,
} from "src/lib/lodas/lodas_pb";
import gRPCClientAbstract from "../abstract/gRPCClient";
import { GRPCClientConfig } from "../abstract/types";

class rpcLodasClient extends gRPCClientAbstract {
  constructor(config: GRPCClientConfig) {
    config.serviceName = "AUTH";
    super(LodasClient, config);
  }

  async listSheets(param: ListSheetsRequest.AsObject) {
    const req = new ListSheetsRequest();
    req.setAscending(param.ascending);
    req.setLimit(param.limit);
    req.setOffset(param.offset);
    req.setSortBy(param.sortBy);
    req.setFrom(
      param.from ? new Timestamp().setSeconds(param.from.seconds) : undefined
    );
    req.setTo(
      param.to ? new Timestamp().setSeconds(param.to.seconds) : undefined
    );

    return await this.gRPCClientRequest<ListSheetsReply.AsObject>(
      "listSheets",
      req
    );
  }

  async createSheet(param: CreateSheetRequest.AsObject) {
    const req = new CreateSheetRequest();
    req.setName(param.name);
    req.setArea(param.area);
    req.setProvince(param.province);
    req.setRatio(param.ratio);
    req.setResultTime(
      param.resultTime
        ? new Timestamp().setSeconds(param.resultTime.seconds)
        : undefined
    );

    return await this.gRPCClientRequest<CreateSheetReply.AsObject>(
      "createSheet",
      req
    );
  }

  async deleteSheet(param: DeleteSheetRequest.AsObject) {
    const req = new DeleteSheetRequest();
    req.setIdsList(param.idsList);

    return await this.gRPCClientRequest<DeleteSheetReply.AsObject>(
      "deleteSheet",
      req
    );
  }

  async listRecords(param: ListRecordsRequest.AsObject) {
    const req = new ListRecordsRequest();
    req.setAscending(param.ascending);
    req.setLimit(param.limit);
    req.setOffset(param.offset);
    req.setSheetId(param.sheetId);
    req.setSortBy(param.sortBy);
    req.setFrom(
      param.from ? new Timestamp().setSeconds(param.from.seconds) : undefined
    );
    req.setTo(
      param.to ? new Timestamp().setSeconds(param.to.seconds) : undefined
    );

    return await this.gRPCClientRequest<ListRecordsReply.AsObject>(
      "listRecords",
      req
    );
  }

  async createRecord(param: CreateRecordRequest.AsObject) {
    const req = new CreateRecordRequest();
    req.setSheetId(param.sheetId);
    req.setBetType(param.betType);
    req.setNumbersList(param.numbersList);
    req.setCash(param.cash);
    req.setNumbersList(param.numbersList);
    req.setPrizesList(param.prizesList);

    return await this.gRPCClientRequest<CreateRecordReply.AsObject>(
      "createRecord",
      req
    );
  }

  async deleteRecord(param: DeleteRecordRequest.AsObject) {
    const req = new DeleteRecordRequest();
    req.setIdsList(param.idsList);

    return await this.gRPCClientRequest<DeleteRecordReply.AsObject>(
      "deleteRecord",
      req
    );
  }
}

export default rpcLodasClient;
