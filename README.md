# Authentication, Authorization & Session Management System

A production-grade, full-stack authentication system built with Go, PostgreSQL, Redis, and vanilla HTML/JS/CSS. This project demonstrates secure user registration, login, role-based authorization, and robust session management using short-lived Access Tokens and long-lived Refresh Tokens.

## 🚀 Features

- **User Registration:** Secure password hashing using `bcrypt`.
- **Login & Session Management:**
  - Short-lived **JWT Access Tokens** (15 minutes).
  - Long-lived **Refresh Tokens** (7 days) securely stored and managed in Redis.
- **Role-Based Access Control (RBAC):** `USER` and `ADMIN` roles enforced via Gin middleware.
- **Profile Management:** Users can view and update their profile details.
- **Logout:** Invalidation of active sessions via Redis key deletion.
- **Frontend Integration:** A clean, responsive UI built with pure HTML, CSS, and vanilla JS leveraging the `fetch` API.
- **Dockerized:** Fully orchestrated with Docker Compose for seamless deployment.

## 🛠 Tech Stack

- **Backend:** Go (Golang), Gin Framework
- **Database:** PostgreSQL (raw SQL queries, no ORM)
- **Caching/Sessions:** Redis
- **Security:** `golang-jwt`, `bcrypt`
- **Frontend:** Vanilla HTML, CSS, JavaScript
- **Deployment:** Docker, Docker Compose, Nginx (Reverse Proxy)

## 📦 How to Run

Ensure you have Docker and Docker Compose installed on your machine.

1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   cd auth-system
   ```

2. **Start the application:**
   ```bash
   docker-compose up -d --build
   ```

3. **Access the Application:**
   - **Frontend UI:** Open your browser and navigate to `http://localhost`
   - **Backend API:** `http://localhost/api`
   
   *Note: On the first startup, a default Admin account is automatically seeded into the database.*
   - **Admin Email:** `admin@example.com`
   - **Admin Password:** `secret123`

4. **Stop the application:**
   ```bash
   docker-compose down
   ```

## 🔌 API Endpoints

| Method | Endpoint | Description | Auth Required | Role Required |
|--------|----------|-------------|---------------|---------------|
| `GET`  | `/api/health` | Health check | ❌ | None |
| `POST` | `/api/auth/register` | Register a new user | ❌ | None |
| `POST` | `/api/auth/login` | Login and receive tokens | ❌ | None |
| `POST` | `/api/auth/refresh` | Issue new access token using refresh token | ❌ | None |
| `POST` | `/api/auth/logout` | Logout (invalidates refresh token in Redis) | ✅ | None |
| `GET`  | `/api/users/profile` | Get logged-in user's profile | ✅ | None |
| `PUT`  | `/api/users/profile` | Update user's name | ✅ | None |
| `GET`  | `/api/admin/users` | Retrieve all users | ✅ | `ADMIN` |

## 🏗 Architecture

1. **Nginx (Frontend container):** Serves static assets on port 80 and proxies any requests starting with `/api/` to the backend service.
2. **Go Backend:** Listens internally on port 8080.
3. **PostgreSQL:** Persists user data, securely storing hashed passwords and role definitions.
4. **Redis:** Manages stateful refresh tokens with a 7-day TTL. When a user logs out, their key is deleted from Redis, invalidating their session instantly.
