import {
  Avatar,
  Box,
  List,
  ListItemButton,
  ListItemAvatar,
  ListItemText,
  Tab,
  Tabs,
  Typography,
  Divider,
  IconButton,
  ListItem,
  Stack,
  Button,
  Grid,
} from "@mui/material";
import { SyntheticEvent, useState } from "react";
import AddTwoToneIcon from "@mui/icons-material/AddTwoTone";
import ImageIcon from "@mui/icons-material/Image";
import WorkIcon from "@mui/icons-material/Work";
import BeachAccessIcon from "@mui/icons-material/BeachAccess";
import DeleteIcon from "@mui/icons-material/Delete";
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import { Helmet } from "react-helmet-async";
import { useParams } from "react-router";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

const Sheets = () => {
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
          <Button
            variant="contained"
            startIcon={<AddTwoToneIcon fontSize="small" />}
          >
            Thêm mới
          </Button>
        </Grid>
        <Grid item xs={12}>
          <List sx={{ width: "100%", bgcolor: "background.paper" }}>
            <Stack direction="row" spacing={1}>
              <ListItemButton>
                <ListItemAvatar>
                  <Avatar>
                    <ImageIcon />
                  </Avatar>
                </ListItemAvatar>

                <ListItemText
                  primary="Sheet mot ngay"
                  secondary="Jan 9, 2014"
                />
              </ListItemButton>
              <IconButton edge="end" aria-label="delete">
                <ModeEditIcon />
              </IconButton>

              <IconButton edge="end" aria-label="delete">
                <DeleteIcon />
              </IconButton>
            </Stack>
            <Divider variant="fullWidth" component="li" />

            <Stack direction="row" spacing={1}>
              <ListItemButton>
                <ListItemAvatar>
                  <Avatar>
                    <ImageIcon />
                  </Avatar>
                </ListItemAvatar>

                <ListItemText
                  primary="Sheet mot ngay"
                  secondary="Jan 9, 2014"
                />
              </ListItemButton>
              <IconButton edge="end" aria-label="delete">
                <ModeEditIcon />
              </IconButton>

              <IconButton edge="end" aria-label="delete">
                <DeleteIcon />
              </IconButton>
            </Stack>
            <Divider variant="fullWidth" component="li" />

            <Stack direction="row" spacing={1}>
              <ListItemButton>
                <ListItemAvatar>
                  <Avatar>
                    <ImageIcon />
                  </Avatar>
                </ListItemAvatar>

                <ListItemText
                  primary="Sheet mot ngay"
                  secondary="Jan 9, 2014"
                />
              </ListItemButton>
              <IconButton edge="end" aria-label="delete">
                <ModeEditIcon />
              </IconButton>

              <IconButton edge="end" aria-label="delete">
                <DeleteIcon />
              </IconButton>
            </Stack>
            <Divider variant="fullWidth" component="li" />
          </List>
        </Grid>
      </Grid>
    </>
  );
};

export default Sheets;
