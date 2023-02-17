/**
 * @fileoverview gRPC-Web generated client stub for lodas
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as lodas_pb from './lodas_pb';


export class LodasClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreateSheet = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/CreateSheet',
    grpcWeb.MethodType.UNARY,
    lodas_pb.CreateSheetRequest,
    lodas_pb.CreateSheetReply,
    (request: lodas_pb.CreateSheetRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.CreateSheetReply.deserializeBinary
  );

  createSheet(
    request: lodas_pb.CreateSheetRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.CreateSheetReply>;

  createSheet(
    request: lodas_pb.CreateSheetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.CreateSheetReply) => void): grpcWeb.ClientReadableStream<lodas_pb.CreateSheetReply>;

  createSheet(
    request: lodas_pb.CreateSheetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.CreateSheetReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/CreateSheet',
        request,
        metadata || {},
        this.methodDescriptorCreateSheet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/CreateSheet',
    request,
    metadata || {},
    this.methodDescriptorCreateSheet);
  }

  methodDescriptorDeleteSheet = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/DeleteSheet',
    grpcWeb.MethodType.UNARY,
    lodas_pb.DeleteSheetRequest,
    lodas_pb.DeleteSheetReply,
    (request: lodas_pb.DeleteSheetRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.DeleteSheetReply.deserializeBinary
  );

  deleteSheet(
    request: lodas_pb.DeleteSheetRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.DeleteSheetReply>;

  deleteSheet(
    request: lodas_pb.DeleteSheetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.DeleteSheetReply) => void): grpcWeb.ClientReadableStream<lodas_pb.DeleteSheetReply>;

  deleteSheet(
    request: lodas_pb.DeleteSheetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.DeleteSheetReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/DeleteSheet',
        request,
        metadata || {},
        this.methodDescriptorDeleteSheet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/DeleteSheet',
    request,
    metadata || {},
    this.methodDescriptorDeleteSheet);
  }

  methodDescriptorUpdateSheet = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/UpdateSheet',
    grpcWeb.MethodType.UNARY,
    lodas_pb.UpdateSheetRequest,
    lodas_pb.UpdateSheetReply,
    (request: lodas_pb.UpdateSheetRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.UpdateSheetReply.deserializeBinary
  );

  updateSheet(
    request: lodas_pb.UpdateSheetRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.UpdateSheetReply>;

  updateSheet(
    request: lodas_pb.UpdateSheetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.UpdateSheetReply) => void): grpcWeb.ClientReadableStream<lodas_pb.UpdateSheetReply>;

  updateSheet(
    request: lodas_pb.UpdateSheetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.UpdateSheetReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/UpdateSheet',
        request,
        metadata || {},
        this.methodDescriptorUpdateSheet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/UpdateSheet',
    request,
    metadata || {},
    this.methodDescriptorUpdateSheet);
  }

  methodDescriptorListSheets = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/ListSheets',
    grpcWeb.MethodType.UNARY,
    lodas_pb.ListSheetsRequest,
    lodas_pb.ListSheetsReply,
    (request: lodas_pb.ListSheetsRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.ListSheetsReply.deserializeBinary
  );

  listSheets(
    request: lodas_pb.ListSheetsRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.ListSheetsReply>;

  listSheets(
    request: lodas_pb.ListSheetsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.ListSheetsReply) => void): grpcWeb.ClientReadableStream<lodas_pb.ListSheetsReply>;

  listSheets(
    request: lodas_pb.ListSheetsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.ListSheetsReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/ListSheets',
        request,
        metadata || {},
        this.methodDescriptorListSheets,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/ListSheets',
    request,
    metadata || {},
    this.methodDescriptorListSheets);
  }

  methodDescriptorListRecords = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/ListRecords',
    grpcWeb.MethodType.UNARY,
    lodas_pb.ListRecordsRequest,
    lodas_pb.ListRecordsReply,
    (request: lodas_pb.ListRecordsRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.ListRecordsReply.deserializeBinary
  );

  listRecords(
    request: lodas_pb.ListRecordsRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.ListRecordsReply>;

  listRecords(
    request: lodas_pb.ListRecordsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.ListRecordsReply) => void): grpcWeb.ClientReadableStream<lodas_pb.ListRecordsReply>;

  listRecords(
    request: lodas_pb.ListRecordsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.ListRecordsReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/ListRecords',
        request,
        metadata || {},
        this.methodDescriptorListRecords,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/ListRecords',
    request,
    metadata || {},
    this.methodDescriptorListRecords);
  }

  methodDescriptorCreateRecord = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/CreateRecord',
    grpcWeb.MethodType.UNARY,
    lodas_pb.CreateRecordRequest,
    lodas_pb.CreateRecordReply,
    (request: lodas_pb.CreateRecordRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.CreateRecordReply.deserializeBinary
  );

  createRecord(
    request: lodas_pb.CreateRecordRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.CreateRecordReply>;

  createRecord(
    request: lodas_pb.CreateRecordRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.CreateRecordReply) => void): grpcWeb.ClientReadableStream<lodas_pb.CreateRecordReply>;

  createRecord(
    request: lodas_pb.CreateRecordRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.CreateRecordReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/CreateRecord',
        request,
        metadata || {},
        this.methodDescriptorCreateRecord,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/CreateRecord',
    request,
    metadata || {},
    this.methodDescriptorCreateRecord);
  }

  methodDescriptorDeleteRecord = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/DeleteRecord',
    grpcWeb.MethodType.UNARY,
    lodas_pb.DeleteRecordRequest,
    lodas_pb.DeleteRecordReply,
    (request: lodas_pb.DeleteRecordRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.DeleteRecordReply.deserializeBinary
  );

  deleteRecord(
    request: lodas_pb.DeleteRecordRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.DeleteRecordReply>;

  deleteRecord(
    request: lodas_pb.DeleteRecordRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.DeleteRecordReply) => void): grpcWeb.ClientReadableStream<lodas_pb.DeleteRecordReply>;

  deleteRecord(
    request: lodas_pb.DeleteRecordRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.DeleteRecordReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/DeleteRecord',
        request,
        metadata || {},
        this.methodDescriptorDeleteRecord,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/DeleteRecord',
    request,
    metadata || {},
    this.methodDescriptorDeleteRecord);
  }

  methodDescriptorCreateBetSetting = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/CreateBetSetting',
    grpcWeb.MethodType.UNARY,
    lodas_pb.CreateBetSettingRequest,
    lodas_pb.CreateBetSettingReply,
    (request: lodas_pb.CreateBetSettingRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.CreateBetSettingReply.deserializeBinary
  );

  createBetSetting(
    request: lodas_pb.CreateBetSettingRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.CreateBetSettingReply>;

  createBetSetting(
    request: lodas_pb.CreateBetSettingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.CreateBetSettingReply) => void): grpcWeb.ClientReadableStream<lodas_pb.CreateBetSettingReply>;

  createBetSetting(
    request: lodas_pb.CreateBetSettingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.CreateBetSettingReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/CreateBetSetting',
        request,
        metadata || {},
        this.methodDescriptorCreateBetSetting,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/CreateBetSetting',
    request,
    metadata || {},
    this.methodDescriptorCreateBetSetting);
  }

  methodDescriptorGetBetSetting = new grpcWeb.MethodDescriptor(
    '/lodas.Lodas/GetBetSetting',
    grpcWeb.MethodType.UNARY,
    lodas_pb.GetBetSettingRequest,
    lodas_pb.GetBetSettingReply,
    (request: lodas_pb.GetBetSettingRequest) => {
      return request.serializeBinary();
    },
    lodas_pb.GetBetSettingReply.deserializeBinary
  );

  getBetSetting(
    request: lodas_pb.GetBetSettingRequest,
    metadata: grpcWeb.Metadata | null): Promise<lodas_pb.GetBetSettingReply>;

  getBetSetting(
    request: lodas_pb.GetBetSettingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: lodas_pb.GetBetSettingReply) => void): grpcWeb.ClientReadableStream<lodas_pb.GetBetSettingReply>;

  getBetSetting(
    request: lodas_pb.GetBetSettingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: lodas_pb.GetBetSettingReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lodas.Lodas/GetBetSetting',
        request,
        metadata || {},
        this.methodDescriptorGetBetSetting,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lodas.Lodas/GetBetSetting',
    request,
    metadata || {},
    this.methodDescriptorGetBetSetting);
  }

}

