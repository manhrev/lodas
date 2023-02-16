import {
  Box,
  Card,
  Divider,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { subDays } from "date-fns";
import RecentOrdersTable from "src/content/applications/Transactions/RecentOrdersTable";
import { CryptoOrder } from "src/models/crypto_order";
import DraftsIcon from "@mui/icons-material/Drafts";
import InboxIcon from "@mui/icons-material/Inbox";
import { Helmet } from "react-helmet-async";

function createData(
  name: string,
  calories: number,
  fat: number,
  carbs: number,
  protein: number
) {
  return { name, calories, fat, carbs, protein };
}

const rows = [
  createData("Frozen yoghurt", 159, 6.0, 24, 4.0),
  createData("Ice cream sandwich", 237, 9.0, 37, 4.3),
  createData("Eclair", 262, 16.0, 24, 6.0),
  createData("Cupcake", 305, 3.7, 67, 4.3),
  createData("Gingerbread", 356, 16.0, 49, 3.9),
];

const Home = () => {
  const cryptoOrders: CryptoOrder[] = [
    {
      id: "1",
      orderDetails: "Fiat Deposit",
      orderDate: new Date().getTime(),
      status: "completed",
      orderID: "VUVX709ET7BY",
      sourceName: "Bank Account",
      sourceDesc: "*** 1111",
      amountCrypto: 34.4565,
      amount: 56787,
      cryptoCurrency: "ETH",
      currency: "$",
    },
    {
      id: "2",
      orderDetails: "Fiat Deposit",
      orderDate: subDays(new Date(), 1).getTime(),
      status: "completed",
      orderID: "23M3UOG65G8K",
      sourceName: "Bank Account",
      sourceDesc: "*** 1111",
      amountCrypto: 6.58454334,
      amount: 8734587,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "3",
      orderDetails: "Fiat Deposit",
      orderDate: subDays(new Date(), 5).getTime(),
      status: "failed",
      orderID: "F6JHK65MS818",
      sourceName: "Bank Account",
      sourceDesc: "*** 1111",
      amountCrypto: 6.58454334,
      amount: 8734587,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "4",
      orderDetails: "Fiat Deposit",
      orderDate: subDays(new Date(), 55).getTime(),
      status: "completed",
      orderID: "QJFAI7N84LGM",
      sourceName: "Bank Account",
      sourceDesc: "*** 1111",
      amountCrypto: 6.58454334,
      amount: 8734587,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "5",
      orderDetails: "Fiat Deposit",
      orderDate: subDays(new Date(), 56).getTime(),
      status: "pending",
      orderID: "BO5KFSYGC0YW",
      sourceName: "Bank Account",
      sourceDesc: "*** 1111",
      amountCrypto: 6.58454334,
      amount: 8734587,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "6",
      orderDetails: "Fiat Deposit",
      orderDate: subDays(new Date(), 33).getTime(),
      status: "completed",
      orderID: "6RS606CBMKVQ",
      sourceName: "Bank Account",
      sourceDesc: "*** 1111",
      amountCrypto: 6.58454334,
      amount: 8734587,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "7",
      orderDetails: "Fiat Deposit",
      orderDate: new Date().getTime(),
      status: "pending",
      orderID: "479KUYHOBMJS",
      sourceName: "Bank Account",
      sourceDesc: "*** 1212",
      amountCrypto: 2.346546,
      amount: 234234,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "8",
      orderDetails: "Paypal Withdraw",
      orderDate: subDays(new Date(), 22).getTime(),
      status: "completed",
      orderID: "W67CFZNT71KR",
      sourceName: "Paypal Account",
      sourceDesc: "*** 1111",
      amountCrypto: 3.345456,
      amount: 34544,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "9",
      orderDetails: "Fiat Deposit",
      orderDate: subDays(new Date(), 11).getTime(),
      status: "completed",
      orderID: "63GJ5DJFKS4H",
      sourceName: "Bank Account",
      sourceDesc: "*** 2222",
      amountCrypto: 1.4389567945,
      amount: 123843,
      cryptoCurrency: "BTC",
      currency: "$",
    },
    {
      id: "10",
      orderDetails: "Wallet Transfer",
      orderDate: subDays(new Date(), 123).getTime(),
      status: "failed",
      orderID: "17KRZHY8T05M",
      sourceName: "Wallet Transfer",
      sourceDesc: "John's Cardano Wallet",
      amountCrypto: 765.5695,
      amount: 7567,
      cryptoCurrency: "ADA",
      currency: "$",
    },
  ];

  return (
    <>
      <Helmet>
        <title>Trang chủ</title>
      </Helmet>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Dessert (100g serving)</TableCell>
              <TableCell align="right">Calories</TableCell>
              <TableCell align="right">Fat&nbsp;(g)</TableCell>
              <TableCell align="right">Carbs&nbsp;(g)</TableCell>
              <TableCell align="right">Protein&nbsp;(g)</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.name}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="right">{row.calories}</TableCell>
                <TableCell align="right">{row.fat}</TableCell>
                <TableCell align="right">{row.carbs}</TableCell>
                <TableCell align="right">{row.protein}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Card>
        <RecentOrdersTable cryptoOrders={cryptoOrders} />
      </Card>
    </>
  );
};

export default Home;
