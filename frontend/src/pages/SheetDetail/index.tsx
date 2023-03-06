import {
  Grid,
  Box,
  Typography,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@mui/material";
import { useEffect, useState } from "react";

import { Helmet } from "react-helmet-async";
import { useNavigate, useParams } from "react-router";
import { toast } from "react-toastify";
import { SheetSortBy, SheetStatus } from "src/lib/lodas/lodas_pb";
import { selectSheetSlice } from "src/redux/feature/sheet/slice";
import { listSheetsThunk } from "src/redux/feature/sheet/thunk";
import { useAppDispatch, useAppSelector } from "src/redux/store";
import { lodasClient } from "src/utils/grpc";
import { AreaStr, ProvinceStr } from "src/utils/helper/enumstr";
import CreateForm from "./CreateForm";
import ListRecord from "./ListRecord";

const SheetDetail = () => {
  const navigate = useNavigate();
  let { id } = useParams();
  const dispatch = useAppDispatch();
  const { sheetList } = useAppSelector(selectSheetSlice);

  const [submitModalOpen, setSubmitModalOpen] = useState(false);
  const handleCloseSubmitModal = () => {
    setSubmitModalOpen(false);
  };
  const handleSubmitConfirm = async () => {
    const { error } = await lodasClient.submitSheet({ id: parseInt(id) });
    handleCloseSubmitModal();
    if (error) {
      toast.error("Lỗi khi gửi bảng ghi, vui lòng thử lại sau!");
      navigate("/sheet");
      return;
    }
    navigate("/sheet");
    toast.success("Gửi bảng ghi thành công!");
    return;
  };

  const fetchSheets = () => {
    dispatch(
      listSheetsThunk({
        ascending: false,
        limit: 100,
        offset: 0,
        sortBy: SheetSortBy.SHEET_SORT_BY_UNSPECIFIED,
        from: undefined,
        to: undefined,
        idsList: [parseInt(id)],
        status: SheetStatus.SHEET_STATUS_NOT_SUBMITTED,
      })
    );
  };
  useEffect(() => {
    fetchSheets();
  }, []);
  const sheet = sheetList.find((sheet) => sheet.id === parseInt(id));

  return (
    <>
      <Helmet>
        <title>Bảng ghi {sheet?.name || ""}</title>
      </Helmet>
      {/* <span>current sheetId is: {id}</span> */}
      <Typography variant="h4" sx={{ mt: 2, ml: 1 }}>{`${
        AreaStr[sheet?.area || 0]
      }, tỉnh ${ProvinceStr[sheet?.province || 0]}`}</Typography>
      {sheet && (
        <>
          <Box sx={{ mt: 2, ml: 1 }}>
            <Grid container spacing={1}>
              <Grid container xs={12} md={3} item>
                <Grid item container spacing={1} maxHeight={350}>
                  <CreateForm sheet={sheet} />
                </Grid>
              </Grid>
              <Grid container xs={12} md={9} item>
                <ListRecord sheet_id={sheet.id} />
                <Button
                  variant="contained"
                  onClick={() => {
                    setSubmitModalOpen(true);
                  }}
                >
                  Submit
                </Button>
              </Grid>
            </Grid>
          </Box>
          <Dialog
            open={submitModalOpen}
            onClose={handleCloseSubmitModal}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
          >
            <DialogTitle id="alert-dialog-title">{"Cảnh báo"}</DialogTitle>
            <DialogContent>
              <DialogContentText id="alert-dialog-description">
                Bạn có chắc chắn muốn gửi bảng ghi này?
              </DialogContentText>
            </DialogContent>
            <DialogActions>
              <Button onClick={handleCloseSubmitModal} variant="outlined">
                Không đồng ý
              </Button>
              <Button
                onClick={handleSubmitConfirm}
                variant="contained"
                autoFocus
              >
                Đồng ý
              </Button>
            </DialogActions>
          </Dialog>
        </>
      )}
    </>
  );
};

export default SheetDetail;
