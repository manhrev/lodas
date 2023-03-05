import { Grid, Box } from "@mui/material";

import { Helmet } from "react-helmet-async";
import { useParams } from "react-router";
import { selectSheetSlice } from "src/redux/feature/sheet/slice";
import { useAppSelector } from "src/redux/store";
import CreateForm from "./CreateForm";
import ListRecord from "./ListRecord";

const SheetDetail = () => {
  let { id } = useParams();
  const { sheetList } = useAppSelector(selectSheetSlice);
  const sheet = sheetList.find((sheet) => sheet.id === parseInt(id));

  return (
    <>
      <Helmet>
        <title>Bảng ghi {sheet.name}</title>
      </Helmet>
      {/* <span>current sheetId is: {id}</span> */}
      <Box sx={{ mt: 2, ml: 1 }}>
        <Grid container spacing={1}>
          <Grid container xs={12} md={3} item>
            <Grid item container spacing={1} maxHeight={350}>
              <CreateForm sheet={sheet} />
            </Grid>
          </Grid>
          <Grid container xs={12} md={9} item>
            <ListRecord sheet_id={sheet.id} />
          </Grid>
        </Grid>
      </Box>
    </>
  );
};

export default SheetDetail;
