import {
  Autocomplete,
  TextField,
  Chip,
  Button,
  Grid,
  Container,
  Box,
} from "@mui/material";

import { Helmet } from "react-helmet-async";
import { useParams } from "react-router";
import { useState } from "react";
import { Area, Prize } from "src/lib/lodas/lodas_pb";
import { checkGiai, checkSoDanh } from "src/utils/helper";

interface KieuDanhType {
  inputValue?: string;
  title: string;
  value?: number;
}

const danhOptions: readonly KieuDanhType[] = [
  { title: "Lô2", value: 1 },
  { title: "Lô3", value: 2 },
  { title: "Đề2", value: 3 },
  { title: "Đề3", value: 4 },
  { title: "Đá", value: 5 },
];

const giaiOptions = []; //["Dsà", "Df", "sdfs"];

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
      <Box sx={{ mt: 2 }}>
        <Grid container spacing={2}>
          <Grid container xs={12} md={4} spacing={1} item>
            <Grid item xs={12}>
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
                renderOption={(props, option) => (
                  <li {...props}>{option.title}</li>
                )}
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
            </Grid>
            <Grid item xs={12}>
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
                      .concat(
                        giaiDanhValue.filter((x) => !newValue.includes(x))
                      );
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
            </Grid>
            <Grid item xs={12}>
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
            </Grid>
            <Grid item xs={12}>
              <TextField
                type="number"
                label="Số tiền"
                placeholder="Số tiền"
                sx={{ backgroundColor: "white", width: "100%" }}
                // disabled={!kieuDanhValue}
                // helperText={soDanhValue.length > 0 ? "" : "Chọn số "}
              />
            </Grid>

            <Grid item xs={12}>
              <Button
                variant="contained"
                onClick={() => console.log(giaiDanhMap)}
              >
                Ok
              </Button>
            </Grid>
          </Grid>
          <Grid container xs={12} md={8} spacing={1} item>
            <div style={{ backgroundColor: "red", width: "100%" }}>asdfds</div>
          </Grid>
        </Grid>
      </Box>
    </>
  );
};

export default SheetDetail;
