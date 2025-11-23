# Shopping Cart Solution

A full‑stack shopping cart application built with **Go (Gin)**, **GORM**, and a **React** frontend, including a **Postman collection** for API testing.

## 📂 Project Structure
```
shopping-cart-solution/
│── backend/        # Go + Gin backend API
│── frontend/       # React frontend application
│── postman/        # Postman collection for API testing
│── README.md       # Project documentation
```

---

## 🚀 Features

### ✅ Frontend (React)
- User signup & login  
- JWT-based authentication  
- Product listing UI  
- Add to cart / remove from cart  
- Checkout flow  
- Order summary screen  

### ✅ Backend (Go + Gin + GORM)
- User registration & secure login  
- Access & refresh token authentication  
- Product CRUD APIs  
- Shopping cart APIs  
- Order creation & retrieval  
- PostgreSQL/MySQL support via GORM  

### ✅ Postman Collection
- Includes all API endpoints:
  - Auth APIs
  - Product APIs
  - Cart APIs
  - Order APIs

---

## 📦 Installation & Setup

### 1️⃣ Backend Setup
```bash
cd backend
go mod tidy
go run main.go
```

### 2️⃣ Frontend Setup
```bash
cd frontend
npm install
npm start
```

### 3️⃣ Database Setup
- Create DB in PostgreSQL/MySQL
- Update the connection string in backend `.env`

Example:
```
DB_USER=root
DB_PASS=password
DB_NAME=shopping
DB_HOST=localhost
DB_PORT=3306
JWT_SECRET=your_secret
```

---

## 🧪 Running with Postman
Import the Postman collection from:
```
/postman/shopping_cart.postman_collection.json
```

You can test:
- Register / Login
- Products
- Cart Operations
- Order APIs

---

## 📘 Tech Stack

### Backend
- Go (Gin)
- GORM (ORM)
- JWT Authentication
- PostgreSQL / MySQL

### Frontend
- React
- Axios
- Context API / LocalStorage

---

## 🔐 Authentication Flow
1. User logs in → receives **access token + refresh token**  
2. Access token used for protected routes  
3. Refresh token used to generate new access tokens  
4. Tokens stored securely in HttpOnly cookies / localStorage  

---

## 🧱 Folder Breakdown

### Backend
```
backend/
│── controllers/
│── middlewares/
│── models/
│── routes/
│── services/
│── utils/
│── main.go
```

### Frontend
```
frontend/
│── src/
    │── components/
    │── pages/
    │── context/
    │── api/
```

---

## 📄 License
This project is for learning and demonstration purposes.

---

## 🤝 Contributing
Feel free to submit issues or pull requests!
