import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import { LodasClient } from "src/lib/lodas/LodasServiceClientPb";
import {
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
}

export default rpcLodasClient;
