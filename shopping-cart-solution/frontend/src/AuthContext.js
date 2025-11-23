import React, { createContext, useState } from "react";

export const AuthContext = createContext(null);

export function AuthProvider({ children }) {
  const [token, setToken] = useState(null);
  const [user, setUser] = useState(null);
  const [cartId, setCartId] = useState(null);

  const login = (token, user) => {
    setToken(token);
    setUser(user);
  };

  const value = {
    token,
    user,
    cartId,
    setCartId,
    login,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
