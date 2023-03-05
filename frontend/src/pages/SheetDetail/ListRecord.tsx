import {
  Button,
  Chip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";
import { RecordSortBy } from "src/lib/lodas/lodas_pb";
import { selectSheetSlice } from "src/redux/feature/sheet/slice";
import {
  deleteRecordThunk,
  listRecordsThunk,
} from "src/redux/feature/sheet/thunk";
import { useAppDispatch, useAppSelector } from "src/redux/store";
import { BetTypeStr, PrizeStr } from "src/utils/helper/enumstr";
import DeleteIcon from "@mui/icons-material/Delete";
import { lodasClient } from "src/utils/grpc";
import { toast } from "react-toastify";
import { Box } from "@mui/system";

interface ListRecordProps {
  sheet_id: number;
}

const ListRecord = ({ sheet_id }: ListRecordProps) => {
  const dispatch = useAppDispatch();
  const { currentRecordList } = useAppSelector(selectSheetSlice);
  const [deleteModalOpen, setDeleteModalOpen] = useState(false);
  const [idDelete, setIdDelete] = useState<number | null>(null);

  const handleCloseDeleteModal = () => {
    setDeleteModalOpen(false);
  };
  const handleDeleteConfirm = async () => {
    const { error } = await dispatch(
      deleteRecordThunk({
        idsList: [idDelete],
      })
    ).unwrap();

    if (error) {
      setDeleteModalOpen(false);
      return toast.error("Lỗi khi xóa, vui lòng thử lại sau");
    }
    setDeleteModalOpen(false);
    return toast.success("Xóa thành công");
  };
  useEffect(() => {
    dispatch(
      listRecordsThunk({
        sheetId: sheet_id,
        ascending: false,
        limit: 999,
        offset: 0,
        sortBy: RecordSortBy.RECORD_SORT_BY_UNSPECIFIED,
        from: undefined,
        to: undefined,
      })
    );
  }, []);
  return (
    <>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell width={10}>ID</TableCell>
              <TableCell>Kiểu đánh</TableCell>
              <TableCell>Giải đánh</TableCell>
              <TableCell>Số đánh</TableCell>
              <TableCell>Số tiền</TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {currentRecordList.length ? (
              currentRecordList.map((row, key) => (
                <TableRow
                  key={row.id}
                  sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                >
                  <TableCell component="th" scope="row" width={10}>
                    {key + 1}
                  </TableCell>
                  <TableCell>
                    <strong>{BetTypeStr[row.betType]}</strong>
                  </TableCell>
                  <TableCell>
                    {row.prizesList
                      .map((prize) => PrizeStr[prize])
                      .toString()
                      .split(",")
                      .map((value) => {
                        return (
                          <Chip
                            label={<strong>{value}</strong>}
                            size="small"
                            color="error"
                          />
                        );
                      })}
                  </TableCell>
                  <TableCell>
                    {row.numbersList
                      .toString()
                      .split(",")
                      .map((value) => {
                        return (
                          <Chip
                            label={<strong>{value}</strong>}
                            size="small"
                            color="primary"
                          />
                        );
                      })}
                  </TableCell>
                  <TableCell>{<strong>{row.cash}</strong>}</TableCell>
                  <TableCell>
                    <IconButton
                      edge="end"
                      aria-label="delete"
                      onClick={() => {
                        setIdDelete(row.id);
                        setDeleteModalOpen(true);
                      }}
                    >
                      <DeleteIcon color="error" />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <TableCell colSpan={6} align="center">
                <strong>Không có dữ liệu</strong>
              </TableCell>
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <Dialog
        open={deleteModalOpen}
        onClose={handleCloseDeleteModal}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{"Cảnh báo"}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Bạn có chắc chắn muốn xóa hàng này?
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

export default ListRecord;
