# Authentication & Session Management Platform

## Overview

Authentication and Session Management Platform built using Go, Gin, PostgreSQL, Redis, Docker, and Nginx.

The application provides secure user authentication, JWT-based authorization, Redis-backed session management, role-based access control, and protected REST APIs. It follows a layered architecture that separates routing, business logic, data access, and authentication concerns.

## Features

* User Registration and Login
* JWT Access and Refresh Tokens
* Redis-Based Session Management
* Role-Based Access Control (RBAC)
* Protected API Endpoints
* User Profile Management
* Admin User Management APIs
* PostgreSQL Persistence
* Session Invalidation on Logout
* Dockerized Deployment

## Technology Stack

### Backend

* Go
* Gin
* PostgreSQL
* Redis
* JWT

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

## System Architecture

```text
+----------------------+
|      Frontend        |
|   HTML, CSS, JS      |
+----------+-----------+
           |
           v
+----------------------+
|      Nginx Server    |
+----------+-----------+
           |
           v
+----------------------+
|      Gin Backend     |
+----------+-----------+
           |
           v
+----------------------+
|      Middleware      |
| JWT & Role Validation|
+----------+-----------+
           |
           v
+----------------------+
|       Handlers       |
+----------+-----------+
           |
           v
+----------------------+
|       Services       |
|   Business Logic     |
+----------+-----------+
           |
           v
+----------------------+
|      Repository      |
|   Database Access    |
+----------+-----------+
           |
           v
+----------------------+
|      PostgreSQL      |
|      User Data       |
+----------------------+

           ^
           |
           |
+----------------------+
|        Redis         |
| Refresh Token Store  |
+----------------------+
```

## Authentication & Authorization Workflow

```text
+----------------------+
|      User Login      |
+----------+-----------+
           |
           v
+----------------------+
| Validate Credentials |
+----------+-----------+
           |
           v
+----------------------+
| Generate JWT Tokens  |
+----------+-----------+
           |
           v
+----------------------+
| Store Refresh Token  |
|      in Redis        |
+----------+-----------+
           |
           v
+----------------------+
| Return Tokens        |
|      to Client       |
+----------+-----------+
           |
           v
+----------------------+
| Protected API Call   |
+----------+-----------+
           |
           v
+----------------------+
| Validate JWT Token   |
+----------+-----------+
           |
           v
+----------------------+
| Check User Role      |
+----------+-----------+
           |
     +-----+-----+
     |           |
     v           v

+-----------+ +-----------+
| Authorized| | Rejected  |
|  Request  | |  Request  |
+-----------+ +-----------+
```

## Role-Based Access Control

| Role  | Permissions                                  |
| ----- | -------------------------------------------- |
| USER  | View Profile, Update Profile                 |
| ADMIN | View Profile, Update Profile, View All Users |

## API Endpoints

### Authentication

```text
POST /auth/register
POST /auth/login
POST /auth/refresh
POST /auth/logout
```

### User

```text
GET /users/profile
PUT /users/profile
```

### Admin

```text
GET /admin/users
```

## Project Structure

```text
auth-session-management-platform
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
├── docker-compose.yml
│
└── README.md
```

## Running the Application

Start all services:

```bash
docker-compose up -d
```

Stop all services:

```bash
docker-compose down
```

View running containers:

```bash
docker ps
```

Verify backend health:

```bash
curl http://localhost:8080/health
```

Build backend locally:

```bash
cd backend
go build ./...
```

## Testing

The following workflows were verified using Postman and Docker:

* User Registration
* User Login
* JWT Authentication
* Protected Route Access
* Profile Retrieval
* Profile Update
* Refresh Token Generation
* Logout and Session Invalidation
* Role-Based Access Control
* Admin User Access

## Screenshots

Add screenshots after testing:

* Login API Response
* User Profile API Response
* Admin Users API Response
* Session Invalidation After Logout
* Login Page
* Profile Page
* Docker Containers Running

## Future Improvements

* Swagger API Documentation
* Unit and Integration Testing
* Rate Limiting
* Email Verification
* Password Reset Workflow
* CI/CD Pipeline Integration

## License

This project was developed for learning and portfolio purposes.
