import {
  Autocomplete,
  TextField,
  Chip,
  Button,
  createFilterOptions,
} from "@mui/material";

import { Helmet } from "react-helmet-async";
import { useParams } from "react-router";
import { useState } from "react";
import { Area, BetType, Prize, Province } from "src/lib/lodas/lodas_pb";

interface KieuDanhType {
  inputValue?: string;
  title: string;
  value?: number;
}

interface GiaiDanhType {
  title: string;
  value: string[];
}

const SheetDetail = () => {
  let { id } = useParams();

  const [area, setArea] = useState<Area>(Area.AREA_BAC);
  const [soDanhValue, setSoDanhValue] = useState<string[]>([]);
  const [kieuDanhValue, setKieuDanhValue] = useState<KieuDanhType | null>(null);
  const [kieuDanhError, setKieuDanhError] = useState(false);
  const [giaiDanhValue, setGiaiDanhValue] = useState<string[]>([]);
  const [giaiDanhMap, setGiaiDanhMap] = useState<Record<string, Prize[]>>({});
  return (
    <>
      <Helmet>
        <title>Bảng ghi 1</title>
      </Helmet>
      {/* <span>current sheetId is: {id}</span> */}
      <Autocomplete
        value={kieuDanhValue}
        onChange={(event, newValue) => {
          if (typeof newValue === "string") {
            let ret = false;
            danhOptions.forEach((value) => {
              if (value.title === newValue) {
                setKieuDanhValue({
                  title: newValue,
                  value: value.value,
                });
                setKieuDanhError(false);
                ret = true;
              }
            });
            if (ret) return;
            setKieuDanhError(true);
          } else if (newValue && newValue.inputValue) {
            setKieuDanhValue(newValue);
            setKieuDanhError(false);
          } else {
            setKieuDanhValue(newValue);
            setKieuDanhError(false);
          }
        }}
        selectOnFocus
        clearOnBlur
        handleHomeEndKeys
        id="Kiểu đánh"
        options={danhOptions}
        getOptionLabel={(option) => {
          // Value selected with enter, right from the input
          if (typeof option === "string") {
            return option;
          }
          return option.title;
        }}
        renderOption={(props, option) => <li {...props}>{option.title}</li>}
        sx={{ width: 200 }}
        freeSolo
        renderInput={(params) => (
          <TextField
            sx={{ backgroundColor: "white" }}
            {...params}
            label="Kiểu đánh"
            error={kieuDanhError}
            helperText={
              kieuDanhError
                ? "Kiểu đánh bị lỗi"
                : "" || kieuDanhValue
                ? ""
                : "Chọn kiểu đánh"
            }
          />
        )}
      />

      <Autocomplete
        freeSolo
        multiple
        id="fixed-tags-demo"
        value={soDanhValue}
        onChange={(event, newValue: string[]) => {
          if (Array.isArray(newValue)) {
            if (newValue.length === 0) return setSoDanhValue(newValue);
            if (checkSoDanh(kieuDanhValue?.value || 0, newValue.at(-1)))
              setSoDanhValue(newValue);
          }
        }}
        options={[]}
        getOptionLabel={(option) => {
          if (typeof option === "string") {
            return option;
          }
          return option[0];
        }}
        renderTags={(tagValue, getTagProps) =>
          tagValue.map((option, index) => (
            <Chip label={option} {...getTagProps({ index })} />
          ))
        }
        style={{ width: 500 }}
        renderInput={(params) => (
          <TextField
            {...params}
            label="Số đánh"
            placeholder="Số đánh"
            sx={{ backgroundColor: "white" }}
            disabled={!kieuDanhValue}
            helperText={soDanhValue.length > 0 ? "" : "Chọn số đánh"}
          />
        )}
      />

      <Autocomplete
        freeSolo
        multiple
        id="fixed-tags-demo"
        value={giaiDanhValue}
        onChange={(event, newValue: string[]) => {
          if (Array.isArray(newValue)) {
            let isOk = false; // if entered value correct

            // if remove all, clear map
            if (newValue.length === 0) {
              setGiaiDanhMap({});
              return setGiaiDanhValue(newValue);
            }

            // removed or added key
            let difference = newValue
              .filter((x) => !giaiDanhValue.includes(x))
              .concat(giaiDanhValue.filter((x) => !newValue.includes(x)));
            console.log(difference);
            // check giai danh
            const listPrize = checkGiai(
              difference[0],
              area,
              kieuDanhValue.value
            );
            console.log(area, listPrize);
            if (listPrize) isOk = true;
            // if new value added, add new map key
            if (newValue.length > giaiDanhValue.length) {
              if (isOk)
                setGiaiDanhMap((prevstate) => {
                  const tempstate = { ...prevstate };

                  tempstate[difference[0]] = listPrize;

                  return tempstate;
                });
            }

            // if old value removed, remove removed map key
            if (newValue.length < giaiDanhValue.length) {
              if (isOk)
                setGiaiDanhMap((prevstate) => {
                  const tempstate = { ...prevstate };
                  difference.forEach((value) => {
                    delete tempstate[value];
                  });
                  return tempstate;
                });
            }

            if (isOk) setGiaiDanhValue(newValue);
          }
        }}
        options={giaiOptions.map((option) => option)}
        getOptionLabel={(option) => {
          if (typeof option === "string") {
            return option;
          }
          return option[0];
        }}
        renderTags={(tagValue, getTagProps) =>
          tagValue.map((option, index) => {
            return <Chip label={option} {...getTagProps({ index })} />;
          })
        }
        style={{ width: 500 }}
        renderInput={(params) => (
          <TextField
            {...params}
            label="Giải"
            placeholder="Giải"
            sx={{ backgroundColor: "white" }}
            disabled={!kieuDanhValue}
            helperText={giaiDanhValue.length > 0 ? "" : "Chọn giải"}
          />
        )}
      />
      <Button variant="contained" onClick={() => console.log(giaiDanhMap)}>
        Ok
      </Button>
    </>
  );
};

const danhOptions: readonly KieuDanhType[] = [
  { title: "Lô2", value: 1 },
  { title: "Lô3", value: 2 },
  { title: "Đề2", value: 3 },
  { title: "Đề3", value: 4 },
  { title: "Đá", value: 5 },
];

const giaiOptions = []; //["Dsà", "Df", "sdfs"];

function checkSoDanh(kieuDanh: number, num: string) {
  // check e
  const parsedNum = parseInt(num);
  if (isNaN(parsedNum)) return false;
  if (kieuDanh == 1 || kieuDanh == 3 || kieuDanh == 5) {
    if (num.length != 2 || parsedNum > 99 || parsedNum < 0) return false;
    return true;
  }
  if (kieuDanh == 2 || kieuDanh == 4) {
    if (num.length != 3 || parsedNum > 999 || parsedNum < 0) return false;
    return true;
  }
  return false;
}

function checkGiai(giai: string, area: Area, betType: BetType): Prize[] | null {
  // kiem tra cac giai don gian nhuw All, Dc, Db, c (ko co so)
  const simple = checkGiaiSimple(giai, area, betType);
  if (simple) return simple;

  // xu li chuoi
  const str = xuLiChuoi(giai, area, betType);
  if (str !== null) return str;
  return null;
}

function xuLiChuoi(s: string, area: Area, betType: BetType): number[] | null {
  if (!s.includes("G")) return null;
  const indG = s.indexOf("G");

  let truoc = xuLiChuoiCon(s.slice(0, indG), area, betType, true);
  let sau = xuLiChuoiCon(s.slice(indG + 1, s.length), area, betType, false);

  // loi ca 2 chuoi
  if (truoc === null || sau === null) return null;
  // 2 chuoi deu trong
  if (truoc.length === 0 && sau.length === 0) {
    return null;
  }

  const truocLenght = truoc.length;
  const sauLenght = sau.length;

  if (truoc.length === 0) truoc = [1, 2, 3, 4, 5, 6, 7, 8];
  if (sau.length === 0) sau = [1, 2, 3, 4, 5, 6, 7, 8];

  let prizeList = [];
  if (area == Area.AREA_BAC) {
    sau.forEach((giai) => {
      truoc.forEach((so) => {
        const giaiMap = GiaiMB[giai];
        if (giaiMap === undefined) return;
        const phanGiai = giaiMap[so - 1];
        if (phanGiai === undefined) return;
        prizeList.push(phanGiai);
      });
    });
  }
  if (area == Area.AREA_TRUNG) {
    sau.forEach((giai) => {
      truoc.forEach((so) => {
        const phanGiai = GiaiMT[giai][so - 1];
        if (phanGiai !== undefined) {
          prizeList.push(phanGiai);
        }
      });
    });
  }
  if (prizeList.length === 0) return null;

  // truong hop 2 ben deu co 1 so
  if (
    truocLenght > 0 &&
    sauLenght > 0 &&
    truocLenght * sauLenght > prizeList.length
  ) {
    return null;
  }
  return prizeList;
}

// xu lis 1/2/3 or 1-3 hoac 4
function xuLiChuoiCon(
  s: string,
  area: Area,
  betType: BetType,
  truoc: boolean // so truoc ?
): number[] | null {
  const kieu2 = [BetType.LO2, BetType.DE2, BetType.DA];
  const kieu3 = [BetType.LO3, BetType.DE3];

  let pattern1 = /^[1-8]-[1-8]$/;
  let pattern2 = /^[1-8](\/[1-8])+$/;
  let pattern3 = /^[1-8]$/;
  if (kieu2.includes(betType)) {
    if (area == Area.AREA_BAC) {
      pattern1 = /^[1-7]-[1-7]$/;
      pattern2 = /^[1-7](\/[1-7])+$/;
      pattern3 = /^[1-7]$/;
    }
  }
  if (kieu3.includes(betType)) {
    if (area == Area.AREA_BAC) {
      if (!truoc) {
        pattern1 = /^[1-6]-[1-6]$/;
        pattern2 = /^[1-6](\/[1-6])+$/;
        pattern3 = /^[1-6]$/;
      }
    }
    if (area == Area.AREA_TRUNG) {
      if (!truoc) {
        pattern1 = /^[1-7]-[1-7]$/;
        pattern2 = /^[1-7](\/[1-7])+$/;
        pattern3 = /^[1-7]$/;
      }
    }
  }

  // ko co chi ca
  if (s === "") {
    return [];
  }
  // co dau "-"
  else if (pattern1.test(s)) {
    const splited = s.split("-");
    const soTrc = parseInt(splited[0]);
    const soSau = parseInt(splited[1]);
    if (soTrc >= soSau) return null;
    return arrayRange(soTrc, soSau, 1);
  }
  // co dau "/"
  else if (pattern2.test(s)) {
    const splited = s.split("/");
    return splited.map((soStr) => {
      return parseInt(soStr);
    });
  }
  // khong co ca hai cai tren
  else if (pattern3.test(s)) {
    return [parseInt(s)];
  } else return null;
}

function checkGiaiSimple(
  giai: string,
  area: Area,
  betType: BetType
): Prize[] | null {
  if (
    betType == BetType.LO2 ||
    betType == BetType.DE2 ||
    betType == BetType.DA
  ) {
    if (area == Area.AREA_BAC) {
      if (giai === "All") {
        return GiaiMB[Cum.DB].concat(
          GiaiMB[Cum.G1],
          GiaiMB[Cum.G2],
          GiaiMB[Cum.G3],
          GiaiMB[Cum.G4],
          GiaiMB[Cum.G5],
          GiaiMB[Cum.G6],
          GiaiMB[Cum.G7]
        );
      }
      if (giai === "DC") {
        return GiaiMB[Cum.DB].concat(GiaiMB[Cum.G7]);
      }
      if (giai === "DB") {
        return GiaiMB[Cum.DB];
      }
      if (giai === "C") {
        return GiaiMB[Cum.G7];
      }
      return null;
    }
    if (area == Area.AREA_TRUNG) {
      if (giai === "All") {
        return GiaiMT[Cum.DB].concat(
          GiaiMT[Cum.G1],
          GiaiMT[Cum.G2],
          GiaiMT[Cum.G3],
          GiaiMT[Cum.G4],
          GiaiMT[Cum.G5],
          GiaiMT[Cum.G6],
          GiaiMT[Cum.G7],
          GiaiMT[Cum.G8]
        );
      }
      if (giai === "DC") {
        return GiaiMT[Cum.DB].concat(GiaiMT[Cum.G8]);
      }
      if (giai === "DB") {
        GiaiMT[Cum.DB];
      }
      if (giai === "C") {
        GiaiMT[Cum.G8];
      }
      return null;
    }
    return null;
  } else if (betType == BetType.LO3 || betType == BetType.DE3) {
    if (area == Area.AREA_BAC) {
      if (giai === "All") {
        return GiaiMB[Cum.DB].concat(
          GiaiMB[Cum.G1],
          GiaiMB[Cum.G2],
          GiaiMB[Cum.G3],
          GiaiMB[Cum.G4],
          GiaiMB[Cum.G5],
          GiaiMB[Cum.G6]
        );
      }
      if (giai === "DC") {
        return GiaiMB[Cum.DB].concat(GiaiMB[Cum.G6]);
      }
      if (giai === "DB") {
        return GiaiMB[Cum.DB];
      }
      if (giai === "C") {
        return GiaiMB[Cum.G6];
      }
      return null;
    }
    if (area == Area.AREA_TRUNG) {
      if (giai === "All") {
        return GiaiMT[Cum.DB].concat(
          GiaiMT[Cum.G1],
          GiaiMT[Cum.G2],
          GiaiMT[Cum.G3],
          GiaiMT[Cum.G4],
          GiaiMT[Cum.G5],
          GiaiMT[Cum.G6],
          GiaiMT[Cum.G7]
        );
      }
      if (giai === "DC") {
        return GiaiMT[Cum.DB].concat(GiaiMT[Cum.G7]);
      }
      if (giai === "DB") {
        GiaiMT[Cum.DB];
      }
      if (giai === "C") {
        GiaiMT[Cum.G7];
      }
      return null;
    }
    return null;
  }
  return null;
}

enum Cum {
  DB = 0,
  G1,
  G2,
  G3,
  G4,
  G5,
  G6,
  G7,
  G8,
}

const GiaiMT = {
  [Cum.DB]: [Prize.PRIZE_1DB],
  [Cum.G1]: [Prize.PRIZE_1G1],
  [Cum.G2]: [Prize.PRIZE_1G2],
  [Cum.G3]: [Prize.PRIZE_1G3, Prize.PRIZE_2G3],
  [Cum.G4]: [
    Prize.PRIZE_1G4,
    Prize.PRIZE_2G4,
    Prize.PRIZE_3G4,
    Prize.PRIZE_4G4,
    Prize.PRIZE_5G4,
    Prize.PRIZE_6G4,
    Prize.PRIZE_7G4,
  ],
  [Cum.G5]: [Prize.PRIZE_1G5],
  [Cum.G6]: [Prize.PRIZE_1G6, Prize.PRIZE_2G6, Prize.PRIZE_3G6],
  [Cum.G7]: [Prize.PRIZE_1G7],
  [Cum.G8]: [Prize.PRIZE_1G8],
};

const GiaiMB = {
  [Cum.DB]: [Prize.PRIZE_1DB],
  [Cum.G1]: [Prize.PRIZE_1G1],
  [Cum.G2]: [Prize.PRIZE_1G2, Prize.PRIZE_2G2],
  [Cum.G3]: [
    Prize.PRIZE_1G3,
    Prize.PRIZE_2G3,
    Prize.PRIZE_3G3,
    Prize.PRIZE_4G3,
    Prize.PRIZE_5G3,
    Prize.PRIZE_6G3,
  ],
  [Cum.G4]: [
    Prize.PRIZE_1G4,
    Prize.PRIZE_2G4,
    Prize.PRIZE_3G4,
    Prize.PRIZE_4G4,
  ],
  [Cum.G5]: [
    Prize.PRIZE_1G5,
    Prize.PRIZE_2G5,
    Prize.PRIZE_3G5,
    Prize.PRIZE_4G5,
    Prize.PRIZE_5G5,
    Prize.PRIZE_6G5,
  ],
  [Cum.G6]: [Prize.PRIZE_1G6, Prize.PRIZE_2G6, Prize.PRIZE_3G6],
  [Cum.G7]: [
    Prize.PRIZE_1G7,
    Prize.PRIZE_2G7,
    Prize.PRIZE_3G7,
    Prize.PRIZE_4G7,
  ],
};

export default SheetDetail;

function arrayRange(start, stop, step) {
  return Array.from(
    { length: (stop - start) / step + 1 },
    (value, index) => start + index * step
  );
}
