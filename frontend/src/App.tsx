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

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnMount: false,
      refetchOnReconnect: false,
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Background>
        <Helmet>
          <link rel="icon" type="image/x-icon" href="/favicon.ico" />
        </Helmet>
        <BrowserRouter basename="/app">
          <Routes>
            <Route path={NRoutes.NEW} element={<CreateNewUserView />} />
            <Route path={NRoutes.NEW_ORDER} element={<CreateOrderView />} />
            <Route
              path={NRoutes.ORDER_CONFIRM}
              element={<OrderConfirmation />}
            />
            <Route path={NRoutes.ERROR} element={<ErrorView />} />
            <Route index element={<Navigate to={NRoutes.NEW} />} />
          </Routes>
        </BrowserRouter>
      </Background>
    </QueryClientProvider>
  );
}

export default App;
