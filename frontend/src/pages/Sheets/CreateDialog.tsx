import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  MenuItem,
  TextField,
  Typography,
} from "@mui/material";
import AddTwoToneIcon from "@mui/icons-material/AddTwoTone";
import { useState } from "react";
import { Area } from "src/lib/lodas/lodas_pb";
import { Province } from "src/lib/lodas/lodas_pb";
import { DesktopDatePicker } from "@mui/x-date-pickers/DesktopDatePicker";
import moment from "moment";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterMoment } from "@mui/x-date-pickers/AdapterMoment";
import { lodasClient } from "src/utils/grpc";
import { toast } from "react-toastify";

interface CreateDialogProps {
  fetchSheets: Function;
}
const CreateDialog = ({ fetchSheets }: CreateDialogProps) => {
  const [open, setOpen] = useState(false);

  // form values
  const [name, setName] = useState<string>(""); // sheet name
  const [mien, setMien] = useState<string>(""); // "bac" | "trung"| "unspecified"
  const [tinh, setTinh] = useState<string>(""); // province enum str
  const [resultTime, setResultTime] = useState<moment.Moment | null>(null);
  const [ratio, setRatio] = useState<string>("");
  const [ratioChia, setRatioChia] = useState<string>("");

  const handleCreateSheet = async () => {
    if (name == "" || mien == "" || tinh == "" || ratio == "" || !resultTime) {
      toast.error("Vui lòng nhập đầy đủ thông tin");
      return;
    }
    const { error } = await lodasClient.createSheet({
      name: name,
      area: mien === "bac" ? Area.AREA_BAC : Area.AREA_TRUNG,
      province: provinceStrMap[tinh][0],
      ratio: Number(ratio),
      resultTime: { seconds: resultTime.unix(), nanos: 0 },
      winRatio: Number(ratioChia),
    });
    if (error) {
      toast.error("Đã có lỗi xảy ra, vui lòng thử lại sau");
      return;
    } else {
      toast.success("Tạo bảng ghi thành công");
    }
    handleClearSheet();
    handleClose();
    fetchSheets();
  };
  const handleClearSheet = () => {
    setName("");
    setMien("");
    setTinh("");
    setResultTime(null);
    setRatio("");
    setOpen(false);
    setRatioChia("");
  };
  const handleResultTimeChange = (date: moment.Moment | null) => {
    setResultTime(date);
  };
  const handleNameChange = (event: any) => {
    setName(event.target.value);
  };
  const handleRatioChange = (event: any) => {
    setRatio(event.target.value);
  };
  const handleClose = () => {
    setOpen(false);
  };
  const handleOpen = () => {
    setOpen(true);
  };
  const handleMienChange = (event: any) => {
    setMien(event.target.value);
    setTinh("");
  };
  const handleTinhChange = (event: any) => {
    setTinh(event.target.value);
  };
  const handleRatioChiaChange = (event: any) => {
    setRatioChia(event.target.value);
  };

  return (
    <LocalizationProvider dateAdapter={AdapterMoment} localeText={{}}>
      <Button
        onClick={handleOpen}
        variant="contained"
        startIcon={<AddTwoToneIcon fontSize="small" />}
      >
        Thêm mới
      </Button>
      <Dialog onClose={handleClose} open={open} fullWidth maxWidth="xs">
        <DialogTitle>
          <Typography variant="h4">Tạo sheet mới</Typography>
        </DialogTitle>
        <DialogContent>
          <TextField
            sx={{ mt: 1 }}
            autoFocus
            id="name"
            label="Tên sheet"
            type="text"
            value={name}
            onChange={handleNameChange}
            fullWidth
          />
          <TextField
            id="area"
            sx={{ mt: 1 }}
            value={mien}
            label="Miền"
            onChange={handleMienChange}
            fullWidth
            select
          >
            <MenuItem value={"bac"}>Miền Bắc</MenuItem>
            <MenuItem value={"trung"}>Miền Trung</MenuItem>
          </TextField>
          <TextField
            id="province"
            sx={{ mt: 1 }}
            value={tinh}
            label="Tỉnh"
            onChange={handleTinhChange}
            disabled={mien == ""}
            fullWidth
            select
          >
            {mien == "bac" && <MenuItem value={"mien_bac"}>Miền Bắc</MenuItem>}
            {mien == "trung" &&
              Object.keys(provinceStrMap).map((key) => {
                if (key == "mien_bac") return "";
                return (
                  <MenuItem value={key} key={key}>
                    {provinceStrMap[key][1]}
                  </MenuItem>
                );
              })}
          </TextField>
          <TextField
            sx={{ mt: 1 }}
            id="ratio"
            label="Tỉ lệ"
            type="number"
            value={ratio}
            onChange={handleRatioChange}
            fullWidth
          />
          <Box sx={{ mt: 1 }}>
            <DesktopDatePicker
              label="Ngày trả kết quả"
              inputFormat="DD/MM/YYYY"
              value={resultTime}
              onChange={handleResultTimeChange}
              renderInput={(params) => <TextField {...params} fullWidth />}
            />
          </Box>
          <TextField
            sx={{ mt: 1 }}
            id="ratio-c"
            label="Tỉ lệ chia"
            type="number"
            value={ratioChia}
            onChange={handleRatioChiaChange}
            fullWidth
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClearSheet} variant="outlined">
            Hủy
          </Button>
          <Button onClick={handleCreateSheet} variant="contained">
            Tạo
          </Button>
        </DialogActions>
      </Dialog>
    </LocalizationProvider>
  );
};

const provinceStrMap = {
  phu_yen: [Province.PHU_YEN, "Phú Yên"],
  hue: [Province.HUE, "Huế"],
  quang_nam: [Province.QUANG_NAM, "Quảng Nam"],
  quang_ngai: [Province.QUANG_NGAI, "Quảng Ngãi"],
  quang_tri: [Province.QUANG_TRI, "Quảng Trị"],
  da_nang: [Province.DA_NANG, "Đà Nẵng"],
  quang_binh: [Province.QUANG_BINH, "Quảng Bình"],
  dak_lak: [Province.DAK_LAK, "Đắk Lắk"],
  dak_nong: [Province.DAK_NONG, "Đắk Nông"],
  kon_tum: [Province.KON_TUM, "Kon Tum"],
  gia_lai: [Province.GIA_LAI, "Gia Lai"],
  binh_dinh: [Province.BINH_DINH, "Bình Định"],
  khanh_hoa: [Province.KHANH_HOA, "Khánh Hòa"],
  ninh_thuan: [Province.NINH_THUAN, "Ninh Thuận"],
  mien_bac: [Province.MIEN_BAC, "Miền Bắc"],
};

export default CreateDialog;
