# Otto Test Go API

## Overview
This project is a Go-based API service for managing brands and vouchers. It provides endpoints for creating and retrieving brands and vouchers, along with necessary middleware for authentication and authorization.

## Project Structure
```
otto-test-go
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   ├── brand_handler.go
│   │   │   └── voucher_handler.go
│   │   ├── middleware
│   │   │   └── auth_middleware.go
│   │   ├── routes.go
│   │   └── server.go
│   ├── config
│   │   └── config.go
│   ├── db
│   │   ├── migrations
│   │   │   ├── 000001_create_brands_table.up.sql
│   │   │   ├── 000001_create_brands_table.down.sql
│   │   │   ├── 000002_create_vouchers_table.up.sql
│   │   │   └── 000002_create_vouchers_table.down.sql
│   │   └── db.go
│   ├── models
│   │   ├── brand.go
│   │   └── voucher.go
│   └── repository
│       ├── brand_repository.go
│       └── voucher_repository.go
├── pkg
│   ├── logger
│   │   └── logger.go
│   └── utils
│       └── response.go
├── tests
│   ├── handlers_test.go
│   └── repositories_test.go
├── docs
│   ├── api.md
│   └── database_schema.md
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
   - Run the migrations to set up the database schema:
     ```bash
     # Assuming you have a migration tool like goose or migrate
     migrate -path internal/db/migrations -database <your-database-url> up
     ```

4. **Run the API**
   Start the API server:
   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints
- **Brands**
  - `POST /brands` - Create a new brand
  - `GET /brands` - Retrieve all brands

- **Vouchers**
  - `POST /vouchers` - Create a new voucher
  - `GET /vouchers?brand_id=<id>` - Retrieve vouchers by brand

## Testing
Run the unit tests to ensure everything is functioning correctly:
```bash
go test ./tests/...
```

## Documentation
- API documentation can be found in `docs/api.md`.
- Database schema documentation is available in `docs/database_schema.md`.

## Contribution
Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.