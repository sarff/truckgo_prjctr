# Auth Microservice - TruckGo Project

## Overview

The `auth` microservice handles authentication and authorization for the TruckGo platform. This microservice is responsible for user registration, login, and token management using **gRPC** and JWT (JSON Web Tokens) for secure access control. It supports different user types such as drivers and customers, and ensures secure communication between clients and other services in the system.

## Features

- **User Registration**: Allows new users to register by providing necessary credentials and details.
- **User Login**: Authenticates users by verifying credentials and generating a JWT token for future requests.
- **Token Validation**: Ensures that users accessing protected routes or services provide a valid JWT.
- **Password Update**: Supports secure password updates for authenticated users.
- **Role-based Access**: Manages different user roles (e.g., driver, customer) to handle access levels for various operations.

## API Endpoints

This service exposes several gRPC endpoints, including but not limited to:

- **RegisterUser**: Registers a new user in the system.
- **LoginUser**: Authenticates a user and returns a JWT token.
- **ValidateToken**: Validates the provided JWT for secure access to other services.
- **UpdatePassword**: Allows users to securely update their password.

## Request Format

Requests and responses use **gRPC** protocol, with message definitions provided in the `proto` files. Below is a sample request and response in JSON format:

```json
{
  "login": "testj1wt132@local.io",
  "password": "1234567890",
  "TypeUser": "customer"
}
```

## Technology Stack

	•	Go: The service is built using Go programming language for performance and scalability.
	•	gRPC: For efficient communication between services.
	•	JWT: For secure authentication and authorization.
	•	PostgreSQL: Database for storing user data and session information.
	•	Docker: Used for containerization and ease of deployment.
