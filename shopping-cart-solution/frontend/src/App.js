import React, { useContext } from "react";
import Login from "./Login";
import ItemsList from "./ItemsList";
import { AuthContext } from "./AuthContext";

export default function App() {
  const { token } = useContext(AuthContext);

  if (!token) {
    return <Login />;
  }

  return <ItemsList />;
}
