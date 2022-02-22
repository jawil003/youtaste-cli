import React from "react";
import { BrowserRouter, Routes } from "react-router-dom";

export interface Props {}

/**
 * An Router React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Router: React.FC<Props> = () => {
  return (
    <BrowserRouter>
      <Routes></Routes>
    </BrowserRouter>
  );
};
