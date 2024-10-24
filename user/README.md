# User Microservice - TruckGo Project

## Overview

The `user` microservice manages user information and profiles within the TruckGo platform. It provides functionality for creating, updating, and retrieving user data, along with handling roles and statuses for different types of users such as drivers and customers.

## Features

- **Create User**: Allows new users to be created in the system.
- **Get User Information**: Retrieves detailed user information based on their unique ID.
- **Update User Profile**: Enables users to update their personal information such as name, email, etc.
- **Role and Status Management**: Manages user roles (e.g., driver, customer) and statuses to control access and permissions across the platform.

## API Endpoints

This service exposes several gRPC endpoints, including:

- **NewDriver(NewDriverRequest)**:   Registers a new user in the system.
- **NewCustomer(NewCustomerRequest)**:   Registers a new user in the system.
- **ListDrivers(ListDriverRequest)**:   Retrieves a list of users based on specific criteria.
- **ListCustomers(ListCustomerRequest)**:   Retrieves a list of users based on specific criteria.
- **UpdateUser(UpdateUserRequest)**: Updates user information.
- **GetType(TypeRequest)**: Retrieves user details by ID.
- **GetUser(UserRequest)**: Retrieves user details by ID.

## Request Format

Requests and responses use **gRPC** protocol, with message definitions provided in the `proto` files. Below is an example of a user creation request in JSON format:

```json
{
  "username": "driver123",
  "email": "driver123@truckgo.io",
  "password": "securepassword",
  "role": "driver",
  "status": "active"
}
```

## Technology Stack

	•	Go: The service is built using Go programming language for performance and scalability.
	•	gRPC: For efficient communication between services.
	•	PostgreSQL: Database for storing user data and session information.
	•	Docker: Used for containerization and ease of deployment.
