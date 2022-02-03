import React from "react";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { CreateNewUserView } from "./views/create-new-user";
import { CreateOrderView } from "./views/create-order";
import { ErrorView } from "./views/error";
import { Routes as NRoutes } from "./enums/routes.enum";
import { Background } from "./components/background/background";
function App() {
  return (
    <Background>
      <BrowserRouter>
        <Routes>
          <Route path={NRoutes.NEW} element={<CreateNewUserView />} />
          <Route path={NRoutes.NEW_ORDER} element={<CreateOrderView />} />
          <Route path={NRoutes.ERROR} element={<ErrorView />} />
          <Route index element={<Navigate to={NRoutes.NEW} />} />
        </Routes>
      </BrowserRouter>
    </Background>
  );
}

export default App;
