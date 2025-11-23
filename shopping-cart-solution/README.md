# Shopping Cart Fullstack Solution

This folder contains a reference implementation for the Shopping Cart assignment.

## Tech stack

- Backend: Go, Gin, GORM, SQLite
- Frontend: React (Create React App)
- Postman: collection with sample requests

## Backend setup

```bash
cd backend
go mod tidy
go run main.go
```

The API will run on `http://localhost:8080`.

### Available endpoints

- `POST /users` – create user
- `GET /users` – list users
- `POST /users/login` – login, returns `{ token, user }`
- `POST /items` – create item
- `GET /items` – list items
- `POST /carts` – **(auth)** add single item to logged-in user's cart
- `GET /carts` – **(auth)** get current open cart for logged-in user
- `POST /orders` – **(auth)** convert the given cart to an order
- `GET /orders` – **(auth)** list orders for logged-in user

Auth is via `Authorization: Bearer <token>` header.

Database is an `shopping-cart.db` SQLite file in the backend folder.

## Frontend setup

```bash
cd frontend
npm install
npm start
```

The React app assumes the backend is running on `http://localhost:8080`.

### Flow

1. Seed some items using the `/items` POST endpoint (use Postman).
2. Click an item in the UI to add to cart.
3. Use **Cart** button to view cart items (shown in `window.alert`).
4. Use **Checkout** to create an order.
5. Use **Order History** to see previous orders (order ids in `window.alert`).

## Postman

Import `postman/ShoppingCart.postman_collection.json` into Postman to quickly test the APIs.
