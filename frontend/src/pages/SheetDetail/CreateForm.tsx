import { Autocomplete, TextField, Chip, Button, Grid } from "@mui/material";
import { useState } from "react";
import { toast } from "react-toastify";
import { Area, BetType, Prize, SheetInfo } from "src/lib/lodas/lodas_pb";
import { lodasClient } from "src/utils/grpc";
import { checkGiai, checkSoDanh } from "src/utils/helper";

interface CreateFormProps {
  sheet: SheetInfo.AsObject;
}

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

const CreateForm = ({ sheet }: CreateFormProps) => {
  const [area] = useState<Area>(sheet.area);
  const [soDanhValue, setSoDanhValue] = useState<string[]>([]);
  const [soDanhError, setSoDanhError] = useState(false);

  const [kieuDanhValue, setKieuDanhValue] = useState<KieuDanhType | null>(null);
  const [kieuDanhError, setKieuDanhError] = useState(false);

  const [giaiDanhValue, setGiaiDanhValue] = useState<string[]>([]);
  const [giaiDanhError, setGiaiDanhError] = useState(false);
  const [giaiDanhMap, setGiaiDanhMap] = useState<Record<string, Prize[]>>({});

  const [soTien, setSoTien] = useState<string>("");
  const [soTienError, setSoTienError] = useState(false);

  const handleSoTienChange = (event) => {
    setSoTien(event.target.value);
    setSoTienError(false);
  };

  const handleSubmit = async () => {
    // VALIDATE
    let err = false;
    if (kieuDanhValue === null) {
      err = true;
      setKieuDanhError(true);
    }

    if (giaiDanhValue.length === 0) {
      err = true;
      setGiaiDanhError(true);
    }

    if (soDanhValue.length === 0) {
      err = true;
      setSoDanhError(true);
    }
    if (soTien === "") {
      err = true;
      setSoTienError(true);
    }
    if (err) return;

    let prizes = [];
    Object.keys(giaiDanhMap).forEach((key, arr) => {
      prizes = prizes.concat(giaiDanhMap[key]);
    });
    prizes = [...new Set(prizes)];
    const { error } = await lodasClient.createRecord({
      betType: kieuDanhValue.value,
      sheetId: sheet.id,
      cash: parseInt(soTien),
      numbersList: soDanhValue,
      prizesList: prizes,
    });

    if (error) {
      return toast.error("Lỗi khi tạo hàng");
    }

    return toast.success("Tạo hàng thành công");
  };

  return (
    <>
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
          renderOption={(props, option) => <li {...props}>{option.title}</li>}
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
                .concat(giaiDanhValue.filter((x) => !newValue.includes(x)));
              // check giai danh
              const listPrize = checkGiai(
                difference[0],
                area,
                kieuDanhValue?.value || undefined
              );
              if (listPrize) isOk = true;
              // if new value added, add new map key
              if (newValue.length > giaiDanhValue.length) {
                if (isOk)
                  setGiaiDanhMap((prevstate) => {
                    setGiaiDanhError(false);
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
              helperText={giaiDanhValue.length > 0 ? "" : "Chọn giải"}
              error={giaiDanhError}
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
              if (checkSoDanh(kieuDanhValue?.value || 0, newValue.at(-1))) {
                setSoDanhValue(newValue);
                setSoDanhError(false);
              }
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
              error={soDanhError}
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
          value={soTien}
          onChange={handleSoTienChange}
          error={soTienError}
          helperText={!soTien ? "Chọn số tiền" : ""}
        />
      </Grid>

      <Grid item xs={12}>
        <Button variant="contained" onClick={handleSubmit}>
          Ok
        </Button>
      </Grid>
    </>
  );
};
export default CreateForm;
