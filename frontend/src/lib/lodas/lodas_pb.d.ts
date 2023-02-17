import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class CreateSheetRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateSheetRequest;

  getArea(): Area;
  setArea(value: Area): CreateSheetRequest;

  getProvince(): Province;
  setProvince(value: Province): CreateSheetRequest;

  getRatio(): number;
  setRatio(value: number): CreateSheetRequest;

  getResultTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setResultTime(value?: google_protobuf_timestamp_pb.Timestamp): CreateSheetRequest;
  hasResultTime(): boolean;
  clearResultTime(): CreateSheetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSheetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSheetRequest): CreateSheetRequest.AsObject;
  static serializeBinaryToWriter(message: CreateSheetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSheetRequest;
  static deserializeBinaryFromReader(message: CreateSheetRequest, reader: jspb.BinaryReader): CreateSheetRequest;
}

export namespace CreateSheetRequest {
  export type AsObject = {
    name: string,
    area: Area,
    province: Province,
    ratio: number,
    resultTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateSheetReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSheetReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSheetReply): CreateSheetReply.AsObject;
  static serializeBinaryToWriter(message: CreateSheetReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSheetReply;
  static deserializeBinaryFromReader(message: CreateSheetReply, reader: jspb.BinaryReader): CreateSheetReply;
}

export namespace CreateSheetReply {
  export type AsObject = {
  }
}

export class DeleteSheetRequest extends jspb.Message {
  getIdsList(): Array<number>;
  setIdsList(value: Array<number>): DeleteSheetRequest;
  clearIdsList(): DeleteSheetRequest;
  addIds(value: number, index?: number): DeleteSheetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSheetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSheetRequest): DeleteSheetRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteSheetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSheetRequest;
  static deserializeBinaryFromReader(message: DeleteSheetRequest, reader: jspb.BinaryReader): DeleteSheetRequest;
}

export namespace DeleteSheetRequest {
  export type AsObject = {
    idsList: Array<number>,
  }
}

export class DeleteSheetReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSheetReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSheetReply): DeleteSheetReply.AsObject;
  static serializeBinaryToWriter(message: DeleteSheetReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSheetReply;
  static deserializeBinaryFromReader(message: DeleteSheetReply, reader: jspb.BinaryReader): DeleteSheetReply;
}

export namespace DeleteSheetReply {
  export type AsObject = {
  }
}

export class UpdateSheetRequest extends jspb.Message {
  getSheetId(): number;
  setSheetId(value: number): UpdateSheetRequest;

  getName(): string;
  setName(value: string): UpdateSheetRequest;

  getArea(): Area;
  setArea(value: Area): UpdateSheetRequest;

  getProvince(): Province;
  setProvince(value: Province): UpdateSheetRequest;

  getRatio(): number;
  setRatio(value: number): UpdateSheetRequest;

  getResultTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setResultTime(value?: google_protobuf_timestamp_pb.Timestamp): UpdateSheetRequest;
  hasResultTime(): boolean;
  clearResultTime(): UpdateSheetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSheetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSheetRequest): UpdateSheetRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateSheetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSheetRequest;
  static deserializeBinaryFromReader(message: UpdateSheetRequest, reader: jspb.BinaryReader): UpdateSheetRequest;
}

export namespace UpdateSheetRequest {
  export type AsObject = {
    sheetId: number,
    name: string,
    area: Area,
    province: Province,
    ratio: number,
    resultTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class UpdateSheetReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSheetReply.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSheetReply): UpdateSheetReply.AsObject;
  static serializeBinaryToWriter(message: UpdateSheetReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSheetReply;
  static deserializeBinaryFromReader(message: UpdateSheetReply, reader: jspb.BinaryReader): UpdateSheetReply;
}

export namespace UpdateSheetReply {
  export type AsObject = {
  }
}

export class ListSheetsRequest extends jspb.Message {
  getLimit(): number;
  setLimit(value: number): ListSheetsRequest;

  getOffset(): number;
  setOffset(value: number): ListSheetsRequest;

  getAscending(): boolean;
  setAscending(value: boolean): ListSheetsRequest;

  getSortBy(): SheetSortBy;
  setSortBy(value: SheetSortBy): ListSheetsRequest;

  getFrom(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFrom(value?: google_protobuf_timestamp_pb.Timestamp): ListSheetsRequest;
  hasFrom(): boolean;
  clearFrom(): ListSheetsRequest;

  getTo(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTo(value?: google_protobuf_timestamp_pb.Timestamp): ListSheetsRequest;
  hasTo(): boolean;
  clearTo(): ListSheetsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSheetsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListSheetsRequest): ListSheetsRequest.AsObject;
  static serializeBinaryToWriter(message: ListSheetsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSheetsRequest;
  static deserializeBinaryFromReader(message: ListSheetsRequest, reader: jspb.BinaryReader): ListSheetsRequest;
}

export namespace ListSheetsRequest {
  export type AsObject = {
    limit: number,
    offset: number,
    ascending: boolean,
    sortBy: SheetSortBy,
    from?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    to?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class ListSheetsReply extends jspb.Message {
  getSheetsList(): Array<SheetInfo>;
  setSheetsList(value: Array<SheetInfo>): ListSheetsReply;
  clearSheetsList(): ListSheetsReply;
  addSheets(value?: SheetInfo, index?: number): SheetInfo;

  getTotal(): number;
  setTotal(value: number): ListSheetsReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSheetsReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListSheetsReply): ListSheetsReply.AsObject;
  static serializeBinaryToWriter(message: ListSheetsReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSheetsReply;
  static deserializeBinaryFromReader(message: ListSheetsReply, reader: jspb.BinaryReader): ListSheetsReply;
}

export namespace ListSheetsReply {
  export type AsObject = {
    sheetsList: Array<SheetInfo.AsObject>,
    total: number,
  }
}

export class SheetInfo extends jspb.Message {
  getId(): number;
  setId(value: number): SheetInfo;

  getStatus(): SheetStatus;
  setStatus(value: SheetStatus): SheetInfo;

  getName(): string;
  setName(value: string): SheetInfo;

  getArea(): Area;
  setArea(value: Area): SheetInfo;

  getProvince(): Province;
  setProvince(value: Province): SheetInfo;

  getRatio(): number;
  setRatio(value: number): SheetInfo;

  getResultTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setResultTime(value?: google_protobuf_timestamp_pb.Timestamp): SheetInfo;
  hasResultTime(): boolean;
  clearResultTime(): SheetInfo;

  getCreatedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedTime(value?: google_protobuf_timestamp_pb.Timestamp): SheetInfo;
  hasCreatedTime(): boolean;
  clearCreatedTime(): SheetInfo;

  getUpdatedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedTime(value?: google_protobuf_timestamp_pb.Timestamp): SheetInfo;
  hasUpdatedTime(): boolean;
  clearUpdatedTime(): SheetInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SheetInfo.AsObject;
  static toObject(includeInstance: boolean, msg: SheetInfo): SheetInfo.AsObject;
  static serializeBinaryToWriter(message: SheetInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SheetInfo;
  static deserializeBinaryFromReader(message: SheetInfo, reader: jspb.BinaryReader): SheetInfo;
}

export namespace SheetInfo {
  export type AsObject = {
    id: number,
    status: SheetStatus,
    name: string,
    area: Area,
    province: Province,
    ratio: number,
    resultTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    createdTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateRecordRequest extends jspb.Message {
  getSheetId(): number;
  setSheetId(value: number): CreateRecordRequest;

  getNumbersList(): Array<string>;
  setNumbersList(value: Array<string>): CreateRecordRequest;
  clearNumbersList(): CreateRecordRequest;
  addNumbers(value: string, index?: number): CreateRecordRequest;

  getCash(): number;
  setCash(value: number): CreateRecordRequest;

  getPrizesList(): Array<Prize>;
  setPrizesList(value: Array<Prize>): CreateRecordRequest;
  clearPrizesList(): CreateRecordRequest;
  addPrizes(value: Prize, index?: number): CreateRecordRequest;

  getBetType(): BetType;
  setBetType(value: BetType): CreateRecordRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRecordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRecordRequest): CreateRecordRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRecordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRecordRequest;
  static deserializeBinaryFromReader(message: CreateRecordRequest, reader: jspb.BinaryReader): CreateRecordRequest;
}

export namespace CreateRecordRequest {
  export type AsObject = {
    sheetId: number,
    numbersList: Array<string>,
    cash: number,
    prizesList: Array<Prize>,
    betType: BetType,
  }
}

export class CreateRecordReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRecordReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRecordReply): CreateRecordReply.AsObject;
  static serializeBinaryToWriter(message: CreateRecordReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRecordReply;
  static deserializeBinaryFromReader(message: CreateRecordReply, reader: jspb.BinaryReader): CreateRecordReply;
}

export namespace CreateRecordReply {
  export type AsObject = {
  }
}

export class ListRecordsRequest extends jspb.Message {
  getSheetId(): number;
  setSheetId(value: number): ListRecordsRequest;

  getLimit(): number;
  setLimit(value: number): ListRecordsRequest;

  getOffset(): number;
  setOffset(value: number): ListRecordsRequest;

  getAscending(): boolean;
  setAscending(value: boolean): ListRecordsRequest;

  getSortBy(): RecordSortBy;
  setSortBy(value: RecordSortBy): ListRecordsRequest;

  getFrom(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFrom(value?: google_protobuf_timestamp_pb.Timestamp): ListRecordsRequest;
  hasFrom(): boolean;
  clearFrom(): ListRecordsRequest;

  getTo(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTo(value?: google_protobuf_timestamp_pb.Timestamp): ListRecordsRequest;
  hasTo(): boolean;
  clearTo(): ListRecordsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRecordsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRecordsRequest): ListRecordsRequest.AsObject;
  static serializeBinaryToWriter(message: ListRecordsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRecordsRequest;
  static deserializeBinaryFromReader(message: ListRecordsRequest, reader: jspb.BinaryReader): ListRecordsRequest;
}

export namespace ListRecordsRequest {
  export type AsObject = {
    sheetId: number,
    limit: number,
    offset: number,
    ascending: boolean,
    sortBy: RecordSortBy,
    from?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    to?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class DeleteRecordRequest extends jspb.Message {
  getIdsList(): Array<number>;
  setIdsList(value: Array<number>): DeleteRecordRequest;
  clearIdsList(): DeleteRecordRequest;
  addIds(value: number, index?: number): DeleteRecordRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRecordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRecordRequest): DeleteRecordRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRecordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRecordRequest;
  static deserializeBinaryFromReader(message: DeleteRecordRequest, reader: jspb.BinaryReader): DeleteRecordRequest;
}

export namespace DeleteRecordRequest {
  export type AsObject = {
    idsList: Array<number>,
  }
}

export class DeleteRecordReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRecordReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRecordReply): DeleteRecordReply.AsObject;
  static serializeBinaryToWriter(message: DeleteRecordReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRecordReply;
  static deserializeBinaryFromReader(message: DeleteRecordReply, reader: jspb.BinaryReader): DeleteRecordReply;
}

export namespace DeleteRecordReply {
  export type AsObject = {
  }
}

export class ListRecordsReply extends jspb.Message {
  getRecordsList(): Array<Record>;
  setRecordsList(value: Array<Record>): ListRecordsReply;
  clearRecordsList(): ListRecordsReply;
  addRecords(value?: Record, index?: number): Record;

  getTotal(): number;
  setTotal(value: number): ListRecordsReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRecordsReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListRecordsReply): ListRecordsReply.AsObject;
  static serializeBinaryToWriter(message: ListRecordsReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRecordsReply;
  static deserializeBinaryFromReader(message: ListRecordsReply, reader: jspb.BinaryReader): ListRecordsReply;
}

export namespace ListRecordsReply {
  export type AsObject = {
    recordsList: Array<Record.AsObject>,
    total: number,
  }
}

export class Record extends jspb.Message {
  getId(): number;
  setId(value: number): Record;

  getSheetId(): number;
  setSheetId(value: number): Record;

  getNumbersList(): Array<string>;
  setNumbersList(value: Array<string>): Record;
  clearNumbersList(): Record;
  addNumbers(value: string, index?: number): Record;

  getCash(): number;
  setCash(value: number): Record;

  getPrizesList(): Array<Prize>;
  setPrizesList(value: Array<Prize>): Record;
  clearPrizesList(): Record;
  addPrizes(value: Prize, index?: number): Record;

  getBetType(): BetType;
  setBetType(value: BetType): Record;

  getCashIn(): number;
  setCashIn(value: number): Record;

  getWinInfoMap(): jspb.Map<number, string>;
  clearWinInfoMap(): Record;

  getCashOut(): number;
  setCashOut(value: number): Record;

  getCreateTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreateTime(value?: google_protobuf_timestamp_pb.Timestamp): Record;
  hasCreateTime(): boolean;
  clearCreateTime(): Record;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Record.AsObject;
  static toObject(includeInstance: boolean, msg: Record): Record.AsObject;
  static serializeBinaryToWriter(message: Record, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Record;
  static deserializeBinaryFromReader(message: Record, reader: jspb.BinaryReader): Record;
}

export namespace Record {
  export type AsObject = {
    id: number,
    sheetId: number,
    numbersList: Array<string>,
    cash: number,
    prizesList: Array<Prize>,
    betType: BetType,
    cashIn: number,
    winInfoMap: Array<[number, string]>,
    cashOut: number,
    createTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateBetSettingRequest extends jspb.Message {
  getBetSettingInfo(): BetSettingInfo | undefined;
  setBetSettingInfo(value?: BetSettingInfo): CreateBetSettingRequest;
  hasBetSettingInfo(): boolean;
  clearBetSettingInfo(): CreateBetSettingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateBetSettingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateBetSettingRequest): CreateBetSettingRequest.AsObject;
  static serializeBinaryToWriter(message: CreateBetSettingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateBetSettingRequest;
  static deserializeBinaryFromReader(message: CreateBetSettingRequest, reader: jspb.BinaryReader): CreateBetSettingRequest;
}

export namespace CreateBetSettingRequest {
  export type AsObject = {
    betSettingInfo?: BetSettingInfo.AsObject,
  }
}

export class BetSettingInfo extends jspb.Message {
  getLo2(): number;
  setLo2(value: number): BetSettingInfo;

  getLo3(): number;
  setLo3(value: number): BetSettingInfo;

  getDe2(): number;
  setDe2(value: number): BetSettingInfo;

  getDe3(): number;
  setDe3(value: number): BetSettingInfo;

  getDa(): number;
  setDa(value: number): BetSettingInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BetSettingInfo.AsObject;
  static toObject(includeInstance: boolean, msg: BetSettingInfo): BetSettingInfo.AsObject;
  static serializeBinaryToWriter(message: BetSettingInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BetSettingInfo;
  static deserializeBinaryFromReader(message: BetSettingInfo, reader: jspb.BinaryReader): BetSettingInfo;
}

export namespace BetSettingInfo {
  export type AsObject = {
    lo2: number,
    lo3: number,
    de2: number,
    de3: number,
    da: number,
  }
}

export class CreateBetSettingReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateBetSettingReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateBetSettingReply): CreateBetSettingReply.AsObject;
  static serializeBinaryToWriter(message: CreateBetSettingReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateBetSettingReply;
  static deserializeBinaryFromReader(message: CreateBetSettingReply, reader: jspb.BinaryReader): CreateBetSettingReply;
}

export namespace CreateBetSettingReply {
  export type AsObject = {
  }
}

export class GetBetSettingRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetBetSettingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetBetSettingRequest): GetBetSettingRequest.AsObject;
  static serializeBinaryToWriter(message: GetBetSettingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetBetSettingRequest;
  static deserializeBinaryFromReader(message: GetBetSettingRequest, reader: jspb.BinaryReader): GetBetSettingRequest;
}

export namespace GetBetSettingRequest {
  export type AsObject = {
  }
}

export class GetBetSettingReply extends jspb.Message {
  getBetSettingInfo(): BetSettingInfo | undefined;
  setBetSettingInfo(value?: BetSettingInfo): GetBetSettingReply;
  hasBetSettingInfo(): boolean;
  clearBetSettingInfo(): GetBetSettingReply;

  getCreatedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedTime(value?: google_protobuf_timestamp_pb.Timestamp): GetBetSettingReply;
  hasCreatedTime(): boolean;
  clearCreatedTime(): GetBetSettingReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetBetSettingReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetBetSettingReply): GetBetSettingReply.AsObject;
  static serializeBinaryToWriter(message: GetBetSettingReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetBetSettingReply;
  static deserializeBinaryFromReader(message: GetBetSettingReply, reader: jspb.BinaryReader): GetBetSettingReply;
}

export namespace GetBetSettingReply {
  export type AsObject = {
    betSettingInfo?: BetSettingInfo.AsObject,
    createdTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export enum Area { 
  AREA_UNSPECIFIED = 0,
  AREA_BAC = 1,
  AREA_TRUNG = 2,
}
export enum Province { 
  UNSPECIFIED = 0,
  PHU_YEN = 1,
  HUE = 2,
  DAK_LAK = 3,
  QUANG_NAM = 4,
  DA_NANG = 5,
  KHANH_HOA = 6,
  BINH_DINH = 7,
  QUANG_BINH = 8,
  QUANG_TRI = 9,
  GIA_LAI = 10,
  NINH_THUAN = 11,
  DAK_NONG = 12,
  QUANG_NGAI = 13,
  KON_TUM = 14,
  MIEN_BAC = 15,
}
export enum SheetSortBy { 
  SHEET_SORT_BY_UNSPECIFIED = 0,
}
export enum SheetStatus { 
  SHEET_STATUS_UNSPECIFIED = 0,
  SHEET_STATUS_NOT_SUBMITTED = 1,
  SHEET_STATUS_SUBMITTED = 2,
  SHEET_STATUS_GOT_RESULT = 3,
}
export enum RecordSortBy { 
  RECORD_SORT_BY_UNSPECIFIED = 0,
}
export enum Prize { 
  PRIZE_UNSPECIFIED = 0,
  PRIZE_1DB = 1,
  PRIZE_1G1 = 2,
  PRIZE_1G2 = 3,
  PRIZE_2G2 = 4,
  PRIZE_1G3 = 5,
  PRIZE_2G3 = 6,
  PRIZE_3G3 = 7,
  PRIZE_4G3 = 8,
  PRIZE_5G3 = 9,
  PRIZE_6G3 = 10,
  PRIZE_1G4 = 11,
  PRIZE_2G4 = 12,
  PRIZE_3G4 = 13,
  PRIZE_4G4 = 14,
  PRIZE_5G4 = 15,
  PRIZE_6G4 = 16,
  PRIZE_7G4 = 17,
  PRIZE_1G5 = 18,
  PRIZE_2G5 = 19,
  PRIZE_3G5 = 20,
  PRIZE_4G5 = 21,
  PRIZE_5G5 = 22,
  PRIZE_6G5 = 23,
  PRIZE_1G6 = 24,
  PRIZE_2G6 = 25,
  PRIZE_3G6 = 26,
  PRIZE_1G7 = 27,
  PRIZE_2G7 = 28,
  PRIZE_3G7 = 29,
  PRIZE_4G7 = 30,
  PRIZE_1G8 = 31,
}
export enum BetType { 
  BET_TYPE_UNSPECIFIED = 0,
  LO2 = 1,
  LO3 = 2,
  DE2 = 3,
  DE3 = 4,
  DA = 5,
}
