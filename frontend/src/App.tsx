import React from "react";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { CreateNewUserView } from "./views/create-new-user";
import { ErrorView } from "./views/error";
import { Routes as NRoutes } from "./enums/routes.enum";
import { Background } from "./components/background/background";
import { Helmet } from "react-helmet";
import { OrderConfirmation } from "./views/order-confirmation";
import { QueryClient, QueryClientProvider } from "react-query";
import { logger } from "./config/logger";
import { CookiesProvider } from "react-cookie";
import { Auth } from "./components/auth/auth";
import { Poll } from "./views/poll";
import { AdminNewView } from "./views/admin/new";
import { PositionTrackerDefault } from "./components/position-tracker/position-tracker";
import { ActiveOnRoutes } from "./components/activeOnRoutes";

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
    <BrowserRouter basename="/app">
      <QueryClientProvider client={queryClient}>
        <CookiesProvider>
          <Background>
            <Helmet>
              <link rel="icon" type="image/x-icon" href="/favicon.ico" />
            </Helmet>
            <div className="flex h-full w-full">
              <ActiveOnRoutes
                routes={[NRoutes.NEW, NRoutes.POLLS, NRoutes.ORDER_CONFIRM]}
              >
                {PositionTrackerDefault}
              </ActiveOnRoutes>
              <div className="flex-1">
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
                    path={NRoutes["ADMIN_NEW"]}
                    element={
                      <Auth mode="ADMIN">
                        <AdminNewView />
                      </Auth>
                    }
                  />
                  <Route
                    path={NRoutes["ADMIN_OVERVIEW"]}
                    element={
                      <Auth mode="ADMIN">
                        <AdminNewView />
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
                  <Route
                    path={NRoutes.POLLS}
                    element={
                      <Auth>
                        <Poll />
                      </Auth>
                    }
                  />
                  <Route path={NRoutes.ERROR} element={<ErrorView />} />
                  <Route index element={<Navigate to={NRoutes.NEW} />} />
                </Routes>
              </div>
              <ActiveOnRoutes
                routes={[NRoutes.NEW, NRoutes.POLLS, NRoutes.ORDER_CONFIRM]}
              >
                <div className="w-40" />
              </ActiveOnRoutes>
            </div>
          </Background>
        </CookiesProvider>
      </QueryClientProvider>
      <div id="modal" className="absolute top-0 left-0 w-0 h-0 z-50" />
    </BrowserRouter>
  );
}

export default App;
