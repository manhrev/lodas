import {
  Box,
  Card,
  Typography,
  Container,
  Divider,
  Button,
  FormControl,
  TextField,
  Stack,
} from "@mui/material";
import { Helmet } from "react-helmet-async";
import { loginThunk } from "src/redux/feature/user/thunk";
import { styled } from "@mui/material/styles";
import { useState } from "react";
import { useAppDispatch } from "src/redux/store";
import { Navigate, useNavigate } from "react-router";
import { toast } from "react-toastify";

const MainContent = styled(Box)(
  ({ theme }) => `
      height: 100%;
      display: flex;
      flex: 1;
      overflow: auto;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      background-color: #666666;
  `
);

function Login() {
  const dispatch = useAppDispatch();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const handleLogin = async () => {
    const { error } = await dispatch(
      loginThunk({ password, userName: username })
    ).unwrap();

    if (error) {
      toast.error("Không thể đăng nhập!");
    } else {
      toast.success("Đăng nhập thành công!");
      navigate("/");
    }
  };
  return (
    <>
      <Helmet>
        <title>Đăng nhập</title>
      </Helmet>
      <MainContent>
        <Container maxWidth="md">
          <Container maxWidth="xs">
            <Card sx={{ textAlign: "center", mt: 3, p: 4 }}>
              <FormControl variant="outlined" fullWidth>
                <Typography variant="h2" sx={{ mt: 2, mb: 5 }}>
                  Đăng nhập
                </Typography>
                <Stack spacing={2}>
                  <TextField
                    id="username"
                    label="Tên đăng nhập"
                    variant="outlined"
                    value={username}
                    onChange={(event) => {
                      setUsername(event.target.value);
                    }}
                  />
                  <TextField
                    id="password"
                    label="Mật khẩu"
                    variant="outlined"
                    value={password}
                    onChange={(event) => {
                      setPassword(event.target.value);
                    }}
                  />
                </Stack>
              </FormControl>
              <Divider sx={{ my: 4 }} />
              <Button variant="contained" onClick={handleLogin}>
                Đăng nhập
              </Button>
            </Card>
          </Container>
        </Container>
      </MainContent>
    </>
  );
}

export default Login;
