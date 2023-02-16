import {
  Autocomplete,
  TextField,
  Chip,
  createFilterOptions,
} from "@mui/material";

import { Helmet } from "react-helmet-async";
import { useParams } from "react-router";
import { useState } from "react";

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

  const [soDanhValue, setSoDanhValue] = useState<string[]>([]);
  const [kieuDanhValue, setKieuDanhValue] = useState<KieuDanhType | null>(null);
  const [kieuDanhError, setKieuDanhError] = useState(false);
  const [giaiDanhValue, setGiaiDanhValue] = useState<string[]>([]);
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
            console.log(soDanhValue);
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
            if (newValue.length === 0) return setGiaiDanhValue(newValue);

            setGiaiDanhValue(newValue);
            console.log(giaiDanhValue);
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
            label="Giải"
            placeholder="Giải"
            sx={{ backgroundColor: "white" }}
            disabled={!kieuDanhValue}
            helperText={giaiDanhValue.length > 0 ? "" : "Chọn giải"}
          />
        )}
      />
      {/* <Autocomplete
        value={value}
        onChange={(event, newValue) => {
          setValue(newValue);
        }}
        multiple
        id="tags-filled"
        options={[]}
        freeSolo
        renderTags={(value: string[], getTagProps) =>
          value.map((option: string, index: number) => (
            <Chip
              variant="outlined"
              label={option}
              {...getTagProps({ index })}
            />
          ))
        }
        renderInput={(params) => (
          <TextField {...params} label="Checkboxes" placeholder="Favorites" />
        )}
      /> */}
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

export default SheetDetail;
