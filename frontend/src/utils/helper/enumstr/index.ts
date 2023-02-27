import { Area, Province } from "src/lib/lodas/lodas_pb";

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
