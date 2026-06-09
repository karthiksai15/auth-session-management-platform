# API Gateway, Authentication & Order Management Platform

## Project Overview

This project is a microservices-based platform developed using Spring Boot. It provides secure user authentication, centralized API routing, and order management functionality.

The main goal of this project was to understand how real-world backend systems are structured using API Gateways, authentication services, databases, caching systems, and Dockerized deployment.

The application consists of separate services for authentication and order management, which communicate through an API Gateway.

## Features

* User Registration and Login
* JWT-Based Authentication
* Refresh Token Management using Redis
* Role-Based Access Control
* API Gateway for Request Routing
* Load Balancing for Authentication Service
* Order Creation and Cancellation
* PostgreSQL Database Integration
* Dockerized Deployment
* Swagger API Documentation
* Unit Testing with JUnit and Mockito

## Technologies Used

### Backend

* Java 17
* Spring Boot
* Spring Security
* Spring Data JPA
* Spring Cloud Gateway

### Database & Cache

* PostgreSQL
* Redis

### Frontend

* HTML
* CSS
* JavaScript

### Tools

* Docker
* Docker Compose
* Postman
* Swagger

## Architecture

```text
                        Frontend
                   (HTML, CSS, JS)
                              |
                              |
                              v
                    API Gateway Service
                              |
          --------------------------------------
          |                                    |
          |                                    |
          v                                    v

 Authentication Service                Order Service
   (JWT, Redis, Users)                (Orders Module)

          |                                    |
          |                                    |
          v                                    v

      PostgreSQL                         PostgreSQL
          |
          |
          v

        Redis
```

## Authentication Flow

1. User registers with name, email, and password.
2. User logs in using email and password.
3. Authentication Service verifies credentials.
4. JWT Access Token and Refresh Token are generated.
5. Refresh Token is stored in Redis.
6. Access Token is used for protected API requests.
7. API Gateway validates requests before forwarding them to backend services.

## Order Management Flow

1. User logs into the system.
2. User creates an order.
3. Order details are stored in PostgreSQL.
4. User can view their orders.
5. User can cancel their own orders.
6. Admin users can view all orders.

## Design Patterns Used

### Factory Pattern

Used to return the appropriate authentication strategy.

### Strategy Pattern

Used to implement authentication logic.

### Observer Pattern

Used to handle actions that occur after user registration.

### Builder Pattern

Used while creating API response objects.

## Project Structure

```text
api-gateway-order-platform
│
├── api-gateway
│
├── auth-service
│
├── order-service
│
├── frontend
│
├── postman
│
├── docs
│
└── docker-compose.yml
```

## API Endpoints

### Authentication

POST /api/auth/register

POST /api/auth/login

POST /api/auth/refresh

POST /api/auth/logout

### Orders

POST /api/orders

GET /api/orders

GET /api/orders/{id}

DELETE /api/orders/{id}

### Admin

GET /api/admin/orders

## Running the Project

Build the services:

```bash
mvn clean package
```

Start all containers:

```bash
docker compose up --build
```

Stop all containers:

```bash
docker compose down
```

## Testing

The project includes unit tests using:

* JUnit 5
* Mockito

API testing was performed using Postman.

## Screenshots

Add screenshots here after running the project:

* Login Page
* Registration Page
* Orders Dashboard
* Swagger Documentation
* Docker Containers
* Postman Responses

## Learning Outcomes

Through this project I learned:

* Microservice architecture fundamentals
* API Gateway implementation
* JWT authentication and authorization
* Redis integration
* Docker containerization
* PostgreSQL database integration
* Spring Security
* Design patterns in real applications
* Writing unit tests using JUnit and Mockito
