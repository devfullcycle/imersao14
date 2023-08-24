import { createTheme } from "@mui/material/styles";
import { Roboto } from "next/font/google";

const roboto = Roboto({
  weight: ["300", "400", "500", "700"],
  subsets: ["latin"],
  display: "swap",
});

const defaultTheme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#FFCD00",
      contrastText: "#242526",
    },
    background: {
      default: "#242526",
    },
  },
  typography: {
    fontFamily: roboto.style.fontFamily,
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        "html, body": {
          minHeight: "100vh",
          display: "flex",
          flex: 1,
          flexDirection: "column",
        },
      },
    },
    MuiAppBar: {
      styleOverrides: {
        colorPrimary: {
          backgroundColor: "#FFCD00",
          color: "#242526",
        },
      },
    },
  },
});

export default defaultTheme;
