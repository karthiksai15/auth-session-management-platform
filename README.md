# Authentication & Session Management Platform

## Project Overview

This project is a secure authentication and session management platform developed using Go, Gin, PostgreSQL, Redis, Docker, and Nginx.

The main goal of this project was to understand how modern backend applications handle authentication, authorization, session management, and secure API access using JWT tokens and Redis-backed sessions.

The application supports user registration, login, refresh token management, role-based access control, and protected API access through a layered backend architecture.

## Features

* User Registration and Login
* JWT-Based Authentication
* Refresh Token Management using Redis
* Role-Based Access Control (RBAC)
* Protected REST APIs
* User Profile Management
* Admin User Management APIs
* PostgreSQL Database Integration
* Session Invalidation on Logout
* Dockerized Deployment
* Nginx Frontend Hosting

## Technologies Used

### Backend

* Go
* Gin
* JWT
* PostgreSQL
* Redis

### Frontend

* HTML
* CSS
* JavaScript
* Nginx

### Tools

* Docker
* Docker Compose
* Postman
* Git
* GitHub

## Architecture

```text
                     Frontend
                (HTML, CSS, JS)
                        |
                        |
                        v

                  Gin Backend API
                        |
         --------------------------------
         |              |              |
         |              |              |
         v              v              v

      Routes       Middleware      Handlers
                                        |
                                        |
                                        v

                                   Services
                                        |
                                        |
                                        v

                                  Repository
                                        |
                                        |
                                        v

                                  PostgreSQL

                                        ^
                                        |
                                        |
                                     Redis
                           (Refresh Token Store)
```

## Authentication Flow

1. User registers using name, email, and password.
2. Password is hashed before being stored in PostgreSQL.
3. User logs in using email and password.
4. Credentials are validated by the authentication service.
5. JWT Access Token and Refresh Token are generated.
6. Refresh Token is stored in Redis.
7. Access Token is used for accessing protected APIs.
8. Middleware validates JWT tokens before allowing access.
9. On logout, the Redis session is removed and future refresh requests are rejected.

## Session Management Flow

1. User logs in successfully.
2. Refresh token is stored in Redis.
3. Access token expires after a short duration.
4. Client requests a new access token using the refresh token.
5. Redis validates the active session.
6. A new access token is generated.
7. Logout removes the stored refresh token from Redis.

## Role-Based Access Control

### USER

* View Profile
* Update Profile

### ADMIN

* View Profile
* Update Profile
* View All Users

## Project Structure

```text
secure-auth-platform
│
├── backend
│   ├── cmd
│   ├── config
│   ├── handlers
│   ├── middleware
│   ├── models
│   ├── repository
│   ├── routes
│   ├── services
│   └── utils
│
├── frontend
│
├── postgres
│
└── docker-compose.yml
```

## API Endpoints

### Authentication

POST /auth/register

POST /auth/login

POST /auth/refresh

POST /auth/logout

### User

GET /users/profile

PUT /users/profile

### Admin

GET /admin/users

## Running the Project

Start all services:

```bash
docker-compose up --build
```

Stop all services:

```bash
docker-compose down
```

Health Check:

```bash
curl http://localhost:8080/health
```

## Testing

The application was tested using:

* Postman
* Docker Containers
* PostgreSQL
* Redis

The following workflows were verified:

* User Registration
* Login Authentication
* Protected Route Access
* Profile Update
* Refresh Token Generation
* Logout and Session Invalidation
* Role-Based Access Control
* Admin API Access

## Screenshots

Add screenshots after testing:

* Login API Response
* Profile API Response
* Admin User API
* Logout Session Invalidation
* Frontend Login Page
* Frontend Profile Page
* Docker Containers Running

## Learning Outcomes

Through this project I learned:

* JWT authentication and authorization
* Session management using Redis
* Role-based access control
* Layered backend architecture
* PostgreSQL integration
* Docker containerization
* REST API development using Gin
* Secure password storage using bcrypt
* Middleware implementation in Go
