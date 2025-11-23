const BASE_URL = "http://localhost:8080";

async function request(path, { method = "GET", token, body } = {}) {
  const headers = { "Content-Type": "application/json" };
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const res = await fetch(`${BASE_URL}${path}`, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
  });

  if (!res.ok) {
    const errorBody = await res.json().catch(() => ({}));
    throw new Error(errorBody.error || res.statusText);
  }
  return res.json();
}

export function loginApi(username, password) {
  return request("/users/login", {
    method: "POST",
    body: { username, password },
  });
}

export function listItems(token) {
  return request("/items", { token });
}

export function addToCart(token, itemId) {
  return request("/carts", {
    method: "POST",
    token,
    body: { item_id: itemId },
  });
}

export function getMyCart(token) {
  return request("/carts", { token });
}

export function createOrder(token, cartId) {
  return request("/orders", {
    method: "POST",
    token,
    body: { cart_id: cartId },
  });
}

export function listOrders(token) {
  return request("/orders", { token });
}
