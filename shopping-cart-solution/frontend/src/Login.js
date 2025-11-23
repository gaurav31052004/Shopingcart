import React, { useState, useContext } from "react";
import { loginApi } from "./api";
import { AuthContext } from "./AuthContext";

export default function Login() {
  const { login } = useContext(AuthContext);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = await loginApi(username, password);
      login(data.token, data.user);
    } catch (err) {
      window.alert("Invalid username/password");
    }
  };

  return (
    <div style={{ maxWidth: 400, margin: "40px auto" }}>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: 10 }}>
          <label>
            Username:
            <input
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              style={{ width: "100%" }}
            />
          </label>
        </div>
        <div style={{ marginBottom: 10 }}>
          <label>
            Password:
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              style={{ width: "100%" }}
            />
          </label>
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}
