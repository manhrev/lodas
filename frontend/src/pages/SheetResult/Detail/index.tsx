import { useParams } from "react-router";
import { selectSheetSlice } from "src/redux/feature/sheet/slice";
import { useAppDispatch, useAppSelector } from "src/redux/store";
import {
  Chip,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import { BetTypeStr, PrizeStr } from "src/utils/helper/enumstr";
import { useEffect } from "react";
import { listRecordsThunk } from "src/redux/feature/sheet/thunk";
import { RecordSortBy } from "src/lib/lodas/lodas_pb";

const SheetResultRecord = () => {
  let { id: sheet_id } = useParams();
  const dispatch = useAppDispatch();
  const { currentRecordList } = useAppSelector(selectSheetSlice);
  useEffect(() => {
    dispatch(
      listRecordsThunk({
        sheetId: parseInt(sheet_id),
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
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell width={10}>ID</TableCell>
            <TableCell>Kiểu đánh</TableCell>
            <TableCell>Giải đánh</TableCell>
            <TableCell>Số đánh</TableCell>
            <TableCell>Số trúng</TableCell>
            <TableCell>Số tiền</TableCell>
            <TableCell>Tiền vào</TableCell>
            <TableCell>Tiền ra</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {currentRecordList.length ? (
            currentRecordList.map((row, key) => {
              return (
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
                      .map((value, idx) => {
                        return (
                          <Chip
                            key={idx}
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
                      .map((value, idx) => {
                        return (
                          <Chip
                            key={idx}
                            label={<strong>{value}</strong>}
                            size="small"
                            color="primary"
                          />
                        );
                      })}
                  </TableCell>
                  <TableCell>
                    {row.winInfoMap.map((val, idx) => (
                      <Chip
                        key={idx}
                        label={
                          <strong>{`${PrizeStr[val[0]]} | ${val[1]}`}</strong>
                        }
                        size="small"
                        color="default"
                      />
                    ))}
                  </TableCell>
                  <TableCell>{<strong>{row.cash}</strong>}</TableCell>
                  <TableCell>{row.cashIn}</TableCell>
                  <TableCell>{row.cashOut}</TableCell>
                </TableRow>
              );
            })
          ) : (
            <TableRow>
              <TableCell colSpan={6} align="center">
                <strong>Không có dữ liệu</strong>
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </TableContainer>
  );
};
export default SheetResultRecord;
