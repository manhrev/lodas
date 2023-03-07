import { useEffect } from "react";
import { SheetSortBy, SheetStatus } from "src/lib/lodas/lodas_pb";
import { selectSheetSlice } from "src/redux/feature/sheet/slice";
import { listSheetsThunk } from "src/redux/feature/sheet/thunk";
import { useAppDispatch, useAppSelector } from "src/redux/store";
import {
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { AreaStr, ProvinceStr, SheetStatusStr } from "src/utils/helper/enumstr";
import { formatDate, formatDateTime } from "src/utils/helper/datetime";
import { useNavigate } from "react-router";

const SheetResult = () => {
  const dispatch = useAppDispatch();
  const sheets = useAppSelector(selectSheetSlice);
  const navigate = useNavigate();

  useEffect(() => {
    dispatch(
      listSheetsThunk({
        ascending: false,
        idsList: [],
        limit: 100,
        offset: 0,
        sortBy: SheetSortBy.SHEET_SORT_BY_UNSPECIFIED,
        status: SheetStatus.SHEET_STATUS_SUBMITTED,
        from: undefined,
        to: undefined,
      })
    );
  }, []);
  return (
    sheets.sheetList && (
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>
                <strong>Id</strong>
              </TableCell>
              <TableCell>
                <strong>Tên</strong>
              </TableCell>
              <TableCell>
                <strong>Vùng</strong>
              </TableCell>
              <TableCell>
                <strong>Tỉnh</strong>
              </TableCell>
              <TableCell>
                <strong>Tỉ lệ</strong>
              </TableCell>
              <TableCell>
                <strong>Trạng thái</strong>
              </TableCell>
              <TableCell>
                <strong>Thời gian nhận kết quả</strong>
              </TableCell>
              <TableCell>
                <strong>Cập nhật vào</strong>
              </TableCell>
              <TableCell>
                <strong>Hành động</strong>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {sheets.sheetList.map((row) => (
              <TableRow key={row.id}>
                <TableCell component="th" scope="row">
                  {row.id}
                </TableCell>
                <TableCell>{row.name}</TableCell>
                <TableCell>{AreaStr[row.area]}</TableCell>
                <TableCell>{ProvinceStr[row.province]}</TableCell>
                <TableCell>{row.ratio}</TableCell>
                <TableCell>{SheetStatusStr[row.status]}</TableCell>
                <TableCell>{formatDate(row.resultTime)}</TableCell>
                <TableCell>{formatDateTime(row.updatedTime)}</TableCell>
                <TableCell>
                  <Button
                    variant="contained"
                    size="small"
                    onClick={() => {
                      navigate("/ket-qua/" + row.id);
                    }}
                  >
                    Xem
                  </Button>
                  <Button variant="contained" size="small" color="error">
                    Xóa
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    )
  );
};

export default SheetResult;
