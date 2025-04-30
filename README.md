# Otto Test Go API

## Overview

This project is a Go-based API service for managing brands and vouchers. It provides endpoints for creating and retrieving brands and vouchers, along with necessary middleware for authentication and authorization.

## Project Structure

```
otto-test-go
├── app                      # Compiled binary
├── cmd
│   ├── api                  # Main application
│   │   └── main.go
│   └── migrate              # Migration tool
│       └── main.go
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   ├── brand_handler.go
│   │   │   ├── transaction_handler.go
│   │   │   └── voucher_handler.go
│   │   ├── middleware
│   │   │   └── auth_middleware.go
│   │   └── routes.go
│   ├── config
│   │   └── config.go
│   ├── db
│   │   ├── db.go
│   │   └── migration.go
│   ├── models
│   │   ├── brand.go
│   │   ├── transaction.go
│   │   └── voucher.go
│   └── repository
│       ├── brand_repository.go
│       ├── interfaces.go
│       ├── transaction_repository.go
│       └── voucher_repository.go
├── migrations               # SQL migration files
│   ├── 000001_create_brands_table.up.sql
│   ├── 000001_create_brands_table.down.sql
│   ├── 000002_create_vouchers_table.up.sql
│   ├── 000002_create_vouchers_table.down.sql
│   ├── 000003_create_transactions_table.up.sql
│   ├── 000003_create_transactions_table.down.sql
│   ├── 000004_create_transaction_items_table.up.sql
│   └── 000004_create_transaction_items_table.down.sql
├── tests
│   ├── handlers_test.go
│   └── repositories_test.go
├── scripts
│   └── schema.sql          # Database schema
├── test_api.sh             # API testing script
├── go.mod
├── go.sum
└── README.md
```

## Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd otto-test-go
   ```

2. **Install Dependencies**
   Ensure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

3. **Database Setup**

   - Configure your database connection in `internal/config/config.go`.

4. **Run the API**
   # Build the application
   ```bash
   go build -o app ./cmd/api
   ```
   # Run the application
   ```bash
   ./app
   ```

## API Endpoints

- **Brands**

  - `POST /brands` - Create a new brand
  - `GET /brands` - Retrieve all brands

- **Vouchers**

  - `POST /vouchers` - Create a new voucher
  - `GET /vouchers?brand_id=<id>` - Retrieve vouchers by brand

- **Transactions**
  - `POST /transaction/redemption` - Create a redemption transaction with multiple vouchers
  - `GET /transaction/redemption?transactionId={id}` - Get transaction details

## Testing

Run the unit tests to ensure everything is functioning correctly:

```bash
go test ./tests/...
```

## Documentation

- API documentation can be found in `docs/api.md`.
- Database schema documentation is available in `docs/database_schema.md`.
