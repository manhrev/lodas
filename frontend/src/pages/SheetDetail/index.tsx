import { Grid, Box } from "@mui/material";

import { Helmet } from "react-helmet-async";
import { useParams } from "react-router";
import { selectSheetSlice } from "src/redux/feature/sheet/slice";
import { useAppSelector } from "src/redux/store";
import CreateForm from "./CreateForm";

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
      <Box sx={{ mt: 2 }}>
        <Grid container spacing={2}>
          <Grid container xs={12} md={4} spacing={1} item>
            <CreateForm sheet={sheet} />
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
