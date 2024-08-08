# Backend Assignment (Attis)

## Description

This project consists of two backend services that manage user accounts and transactions (send/withdraw).

### Account Manager Service
- **User:** Login with ID/Password
- **Payment Account:** One user can have multiple accounts (e.g., credit, debit, loan)
- **Payment History:** Records of transactions

### Payment Manager Service
- **Transaction:** Basic information like amount, timestamp, toAddress, status
- **Core transaction process function:** This function will be executed by `/send` or `/withdraw` API

## Features

1. Users need to register/log in and then be able to call APIs.
2. APIs for two operations: send/withdraw. Account statements will be updated after the transaction is successful.
3. APIs to retrieve all accounts and transactions per account of the user.
4. Write Swagger docs for implemented APIs (Optional)
5. Auto Debit/Recurring Payments: Users should be able to set up recurring payments. These payments will automatically be processed at specified intervals. (Optional)

## Tech-stack

- **Authentication:** Recommended using authentication 3rd party (Supertokens, Supabase, etc.)
- **API Server:** Golang (Gin framework)
- **Database:** PostgreSQL
- **Containerization:** Docker (docker-compose)

## Preparation

### Prerequisites

- Docker
- Docker Compose
- Golang (1.20 or later)

### Steps

1. **Clone the Repository:**

    ```sh
    git clone <repository-url>
    cd backend-assignment
    ```

2. **Set Up Environment Variables:**

    Create a `.env` file in the root directory of the project with the following content:

    ```env
    DB_HOST=db
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=postgres
    DB_PORT=5432
    ```

3. **Install Dependencies:**

    Ensure all dependencies are resolved and `go.sum` is generated:

    ```sh
    go mod tidy
    ```

## Running the Project

1. **Build and Run the Containers:**

    Use Docker Compose to build and run the containers:

    ```sh
    docker-compose up --build
    ```

    This command will:
    - Build the Go application
    - Start the PostgreSQL database
    - Run the Go application on `http://localhost:8080`

## Project Structure

    backend-assignment/
    ├── Dockerfile
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── .env
    └── models/
        ├── account.go
        ├── transaction.go
        └── user.go
    └── routes/
        ├── account.go
        ├── auth.go
        └── transaction.go
    

## API Endpoints

### Auth
- **POST /register:** Register a new user
- **POST /login:** Login a user

### Accounts
- **POST /accounts:** Create a new account
- **GET /accounts/:userId:** Get all accounts for a user

### Transactions
- **POST /transactions:** Create a new transaction (send/withdraw)
- **GET /transactions/:accountId:** Get all transactions for an account
