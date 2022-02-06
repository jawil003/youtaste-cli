import React from "react";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { CreateNewUserView } from "./views/create-new-user";
import { CreateOrderView } from "./views/create-order";
import { ErrorView } from "./views/error";
import { Routes as NRoutes } from "./enums/routes.enum";
import { Background } from "./components/background/background";
import { Helmet } from "react-helmet";
import { OrderConfirmation } from "./views/order-confirmation";
import { QueryClient, QueryClientProvider } from "react-query";
import { logger } from "./config/logger";
import { CookiesProvider } from "react-cookie";
import { Auth } from "./components/auth/auth";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnMount: false,
      refetchOnReconnect: false,
      refetchOnWindowFocus: false,
      retry: false,
      retryDelay: 0,
    },
  },
});

logger.info(process.env, "Environment loaded");

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <CookiesProvider>
        <Background>
          <Helmet>
            <link rel="icon" type="image/x-icon" href="/favicon.ico" />
          </Helmet>
          <BrowserRouter basename="/app">
            <Routes>
              <Route
                path={NRoutes.NEW}
                element={
                  <Auth mode="NO_USER">
                    <CreateNewUserView />
                  </Auth>
                }
              />
              <Route
                path={NRoutes.NEW_ORDER}
                element={
                  <Auth>
                    <CreateOrderView />
                  </Auth>
                }
              />
              <Route
                path={NRoutes.ORDER_CONFIRM}
                element={
                  <Auth>
                    <OrderConfirmation />
                  </Auth>
                }
              />
              <Route path={NRoutes.ERROR} element={<ErrorView />} />
              <Route index element={<Navigate to={NRoutes.NEW} />} />
            </Routes>
          </BrowserRouter>
        </Background>
      </CookiesProvider>
    </QueryClientProvider>
  );
}

export default App;
