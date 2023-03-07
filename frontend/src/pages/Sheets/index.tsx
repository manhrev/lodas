import {
  Avatar,
  Box,
  List,
  ListItemButton,
  ListItemAvatar,
  ListItemText,
  Typography,
  Divider,
  IconButton,
  Stack,
  Button,
  Grid,
  CircularProgress,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import ImageIcon from "@mui/icons-material/Image";
import DeleteIcon from "@mui/icons-material/Delete";
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import { Helmet } from "react-helmet-async";
import { useAppDispatch, useAppSelector } from "src/redux/store";
import { listSheetsThunk } from "src/redux/feature/sheet/thunk";
import { SheetSortBy, SheetStatus } from "src/lib/lodas/lodas_pb";
import {
  isSheetSliceLoading,
  selectSheetSlice,
} from "src/redux/feature/sheet/slice";
import CreateDialog from "./CreateDialog";
import { formatDate } from "src/utils/helper/datetime";
import { ProvinceStr } from "src/utils/helper/enumstr";
import { lodasClient } from "src/utils/grpc";
import { toast } from "react-toastify";
import { useNavigate } from "react-router";

const Sheets = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const fetchSheets = () => {
    dispatch(
      listSheetsThunk({
        ascending: false,
        limit: 100,
        offset: 0,
        sortBy: SheetSortBy.SHEET_SORT_BY_UNSPECIFIED,
        from: undefined,
        to: undefined,
        idsList: [],
        status: SheetStatus.SHEET_STATUS_NOT_SUBMITTED,
      })
    );
  };
  useEffect(() => {
    fetchSheets();
  }, []);
  const { sheetList } = useAppSelector(selectSheetSlice);
  const loading = useAppSelector(isSheetSliceLoading);
  const [deleteModalOpen, setDeleteModalOpen] = useState(false);
  const handleCloseDeleteModal = () => setDeleteModalOpen(false);
  const [idDelete, setIdDelete] = useState<number | null>(null);
  const handleDeleteClick = (id: number) => {
    setDeleteModalOpen(true);
    setIdDelete(id);
  };
  const handleDeleteConfirm = async () => {
    const { error } = await lodasClient.deleteSheet({ idsList: [idDelete] });
    if (error) {
      toast.error("Không thể xóa, vui lòng thử lại sau!");
    } else {
      fetchSheets();
      toast.success("Xóa thành công!");
    }
    setDeleteModalOpen(false);
  };
  return (
    <>
      <Helmet>
        <title>Bảng ghi</title>
      </Helmet>
      <Grid
        container
        spacing={1}
        sx={{ marginTop: 1, marginLeft: 1, width: "98%" }}
      >
        <Grid item xs={12} display="flex" justifyContent="flex-end">
          <CreateDialog fetchSheets={fetchSheets} />
        </Grid>
        {loading ? (
          <Box sx={{ display: "flex" }}>
            <CircularProgress />
          </Box>
        ) : (
          <Grid item xs={12}>
            <List sx={{ width: "100%", bgcolor: "background.paper" }}>
              {sheetList.map((sheet, id) => (
                <React.Fragment key={id}>
                  <Stack direction="row" spacing={1}>
                    <ListItemButton
                      onClick={() => {
                        navigate("/sheet/" + sheet.id);
                      }}
                    >
                      <ListItemAvatar>
                        <Avatar>
                          <ImageIcon />
                        </Avatar>
                      </ListItemAvatar>

                      <ListItemText
                        primary={
                          <div style={{ display: "flex" }}>
                            <div style={{ minWidth: "70px" }}>
                              <Typography variant="h4">{sheet.name}</Typography>
                            </div>
                            <Typography variant="body2">
                              {`| ${ProvinceStr[sheet.province]}`}
                            </Typography>
                          </div>
                        }
                        secondary={`Tạo: ${formatDate(
                          sheet.createdTime
                        )} | KQ: ${formatDate(sheet.resultTime)}`}
                      />
                    </ListItemButton>
                    <IconButton edge="end" aria-label="delete">
                      <ModeEditIcon />
                    </IconButton>

                    <IconButton
                      edge="end"
                      aria-label="delete"
                      onClick={() => handleDeleteClick(sheet.id)}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </Stack>
                  <Divider variant="fullWidth" component="li" />
                </React.Fragment>
              ))}
            </List>
          </Grid>
        )}
      </Grid>
      <Dialog
        open={deleteModalOpen}
        onClose={handleCloseDeleteModal}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{"Cảnh báo"}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Bạn có chắc chắn muốn xóa bảng ghi này?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDeleteModal} variant="outlined">
            Không đồng ý
          </Button>
          <Button onClick={handleDeleteConfirm} variant="contained" autoFocus>
            Đồng ý
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default Sheets;
