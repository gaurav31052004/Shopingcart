import React, { useContext, useEffect, useState } from "react";
import { AuthContext } from "./AuthContext";
import { listItems, addToCart, getMyCart, createOrder, listOrders } from "./api";

export default function ItemsList() {
  const { token, cartId, setCartId } = useContext(AuthContext);
  const [items, setItems] = useState([]);
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    async function loadItems() {
      try {
        const data = await listItems(token);
        setItems(data);
      } catch (e) {
        console.error(e);
      }
    }
    loadItems();
  }, [token]);

  const handleAddToCart = async (item) => {
    try {
      const cart = await addToCart(token, item.id);
      setCartId(cart.id);
      window.alert(`Added ${item.name} to cart`);
    } catch (e) {
      window.alert("Failed to add to cart");
    }
  };

  const handleShowCart = async () => {
    try {
      const cart = await getMyCart(token);
      if (!cart || !cart.items || cart.items.length === 0) {
        window.alert("Cart is empty");
        return;
      }
      const text = cart.items
        .map((ci) => `cart_id: ${ci.cart_id}, item_id: ${ci.item_id}`)
        .join("\n");
      window.alert(text);
    } catch (e) {
      window.alert("No open cart");
    }
  };

  const handleCheckout = async () => {
    try {
      const currentCart = await getMyCart(token);
      if (!currentCart || !currentCart.id) {
        window.alert("No open cart");
        return;
      }
      await createOrder(token, currentCart.id);
      window.alert("Order successful");
    } catch (e) {
      window.alert("Checkout failed");
    }
  };

  const handleOrderHistory = async () => {
    try {
      const data = await listOrders(token);
      setOrders(data);
      if (!data || data.length === 0) {
        window.alert("No orders yet");
        return;
      }
      const text = data.map((o) => `Order id: ${o.id}`).join("\n");
      window.alert(text);
    } catch (e) {
      window.alert("Failed to fetch orders");
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <div style={{ marginBottom: 20 }}>
        <button onClick={handleCheckout}>Checkout</button>
        <button style={{ marginLeft: 10 }} onClick={handleShowCart}>
          Cart
        </button>
        <button style={{ marginLeft: 10 }} onClick={handleOrderHistory}>
          Order History
        </button>
      </div>

      <h2>Items</h2>
      <ul>
        {items.map((item) => (
          <li
            key={item.id}
            style={{ cursor: "pointer" }}
            onClick={() => handleAddToCart(item)}
          >
            {item.name} ({item.status})
          </li>
        ))}
      </ul>
    </div>
  );
}
