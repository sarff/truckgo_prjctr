
# TruckGo - Microservices Architecture

**TruckGo** is a microservice-based system for managing authorization, orders, payments, and shipping. The project consists of several microservices that interact through gRPC to perform operations with customers and manage business processes.

## Project Structure

### Core Microservices:

- **auth**: A microservice for user authorization and authentication using JWT tokens.
- **Order-Service**: Responsible for managing orders.
- **Payment-Service**: Handles payment processing and transactions.
- **Shipping-Service**: Manages shipping operations.
- **rest**: For interaction between front-end and back-end.
- **user**: Manages user profiles, roles, and statuses.
- **Shared**: Modules for logging, configurations, and general logic used by all services.

### Project Tree
```
truckgo/
├── auth/
├── order-service/
├── payment-service/
├── shipping/
├── rest/
├── user/
├── shared/
└── README.md
```

# Running the TruckGo Project with Docker Compose

To run the **TruckGo** project using Docker Compose, follow these steps:

## Prerequisites

- Ensure you have Docker and Docker Compose installed on your machine.

## Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/alexandear/truckgo.git
    cd truckgo
    ```

2. Create a `.env` file if needed and adjust environment variables according to your setup.

3. Build and run the containers:

    ```bash
    docker-compose up --build
    ```

4. The services should now be running. You can access the services as follows:

    - **Auth Database**: `localhost:8880`
    - **User Database**: `localhost:8881`
    - **Order-Service Database**: `localhost:8882`
    - **Payment-Service Database**: `localhost:8883`
    - **Shipping-Service Database**: `localhost:8884`
    - **gRPC Auth**: `localhost:50047`
    - **gRPC Shipping-Service**: `localhost:50051`
    - **gRPC Order-Service**: `localhost:50049`
    - **gRPC User**: `localhost:50048`

## Docker Compose File

Here is an example `docker-compose.yml` file used to run the **TruckGo** project: