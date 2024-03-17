import React from "react";
import { Route, Routes } from "react-router-dom";
import Auth from "./pages/auth/Auth";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/auth" element={<Auth />} />
    </Routes>
  );
};

export default AppRoutes;
