import {
  Area,
  BetType,
  Province,
  Prize,
  SheetStatus,
} from "src/lib/lodas/lodas_pb";

export const ProvinceStr = {
  [Province.UNSPECIFIED]: "?",
  [Province.PHU_YEN]: "Phú Yên",
  [Province.HUE]: "Huế",
  [Province.DAK_LAK]: "Đắk Lắk",
  [Province.QUANG_NAM]: "Quảng Nam",
  [Province.DA_NANG]: "Đà Nẵng",
  [Province.KHANH_HOA]: "Khánh Hòa",
  [Province.BINH_DINH]: "Bình Định",
  [Province.QUANG_BINH]: "Quảng Bình",
  [Province.QUANG_TRI]: "Quảng Trị",
  [Province.GIA_LAI]: "Gia Lai",
  [Province.NINH_THUAN]: "Ninh Thuận",
  [Province.DAK_NONG]: "Đắk Nông",
  [Province.QUANG_NGAI]: "Quảng Ngãi",
  [Province.KON_TUM]: "Kon Tum",
  [Province.MIEN_BAC]: "Miền Bắc",
};

export const AreaStr = {
  [Area.AREA_UNSPECIFIED]: "?",
  [Area.AREA_BAC]: "Miền Bắc",
  [Area.AREA_TRUNG]: "Miền Trung",
};

export const BetTypeStr = {
  [BetType.BET_TYPE_UNSPECIFIED]: "?",
  [BetType.LO2]: "Lô 2",
  [BetType.LO3]: "Lô 3",
  [BetType.DE2]: "Đề 2",
  [BetType.DE3]: "Đề 3",
  [BetType.DA]: "Đá",
};

export const PrizeStr = {
  [Prize.PRIZE_UNSPECIFIED]: "?",
  [Prize.PRIZE_1DB]: "DB",
  [Prize.PRIZE_1G1]: "1G1",
  [Prize.PRIZE_1G2]: "1G2",
  [Prize.PRIZE_2G2]: "2G2",
  [Prize.PRIZE_1G3]: "1G3",
  [Prize.PRIZE_2G3]: "2G3",
  [Prize.PRIZE_3G3]: "3G3",
  [Prize.PRIZE_4G3]: "4G3",
  [Prize.PRIZE_5G3]: "5G3",
  [Prize.PRIZE_6G3]: "6G3",
  [Prize.PRIZE_1G4]: "1G4",
  [Prize.PRIZE_2G4]: "2G4",
  [Prize.PRIZE_3G4]: "3G4",
  [Prize.PRIZE_4G4]: "4G4",
  [Prize.PRIZE_5G4]: "5G4",
  [Prize.PRIZE_6G4]: "6G4",
  [Prize.PRIZE_7G4]: "6G4",
  [Prize.PRIZE_1G5]: "1G5",
  [Prize.PRIZE_2G5]: "2G5",
  [Prize.PRIZE_3G5]: "3G5",
  [Prize.PRIZE_4G5]: "4G5",
  [Prize.PRIZE_5G5]: "5G5",
  [Prize.PRIZE_6G5]: "6G5",
  [Prize.PRIZE_1G6]: "1G6",
  [Prize.PRIZE_2G6]: "2G6",
  [Prize.PRIZE_3G6]: "3G6",
  [Prize.PRIZE_1G7]: "1G7",
  [Prize.PRIZE_2G7]: "2G7",
  [Prize.PRIZE_3G7]: "3G7",
  [Prize.PRIZE_4G7]: "4G7",
  [Prize.PRIZE_1G8]: "1G8",
};

export const SheetStatusStr = {
  [SheetStatus.SHEET_STATUS_UNSPECIFIED]: "?",
  [SheetStatus.SHEET_STATUS_NOT_SUBMITTED]: "Chưa gửi",
  [SheetStatus.SHEET_STATUS_SUBMITTED]: "Đã gửi",
  [SheetStatus.SHEET_STATUS_GOT_RESULT]: "Đã có kết quả",
};
