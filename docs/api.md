# API Documentation

## Overview

This document provides an overview of the API endpoints available for managing brands, vouchers, and transactions in the application. Each endpoint includes details about the request and response formats, as well as usage examples.

## Base URL

```
http://localhost:8080/api
```

## Endpoints

### Brands

#### Create Brand

- **Endpoint:** `POST /brands`
- **Description:** Creates a new brand.
- **Request Body:**
  ```json
  {
    "name": "Brand Name",
    "description": "Brand Description"
  }
  ```
- **Response:**
  - **201 Created**
    ```json
    {
      "id": 1,
      "name": "Brand Name",
      "description": "Brand Description"
    }
    ```
  - **400 Bad Request**
    ```json
    {
      "error": "Invalid input"
    }
    ```

#### Get All Brands

- **Endpoint:** `GET /brands`
- **Description:** Retrieves a list of all brands.
- **Response:**
  - **200 OK**
    ```json
    [
      {
        "id": 1,
        "name": "Brand Name",
        "description": "Brand Description"
      },
      {
        "id": 2,
        "name": "Another Brand",
        "description": "Another Description"
      }
    ]
    ```

### Vouchers

#### Create Voucher

- **Endpoint:** `POST /vouchers`
- **Description:** Creates a new voucher.
- **Request Body:**
  ```json
  {
    "code": "VOUCHER123",
    "discount": 20,
    "brand_id": 1
  }
  ```
- **Response:**
  - **201 Created**
    ```json
    {
      "id": 1,
      "code": "VOUCHER123",
      "discount": 20,
      "brand_id": 1
    }
    ```
  - **400 Bad Request**
    ```json
    {
      "error": "Invalid input"
    }
    ```

#### Get Vouchers by Brand

- **Endpoint:** `GET /vouchers/brand/{brand_id}`
- **Description:** Retrieves all vouchers associated with a specific brand.
- **Response:**
  - **200 OK**
    ```json
    [
      {
        "id": 1,
        "code": "VOUCHER123",
        "discount": 20,
        "brand_id": 1
      }
    ]
    ```

### Transactions

#### Create Transaction

- **Endpoint:** `POST /transactions`
- **Description:** Creates a new transaction.
- **Request Body:**
  ```json
  {
    "voucher_id": 1,
    "amount": 100.0,
    "date": "2023-01-01T00:00:00Z"
  }
  ```
- **Response:**
  - **201 Created**
    ```json
    {
      "id": 1,
      "voucher_id": 1,
      "amount": 100.0,
      "date": "2023-01-01T00:00:00Z"
    }
    ```
  - **400 Bad Request**
    ```json
    {
      "error": "Invalid input"
    }
    ```

#### Get Transactions by Voucher

- **Endpoint:** `GET /transactions/voucher/{voucher_id}`
- **Description:** Retrieves all transactions associated with a specific voucher.
- **Response:**
  - **200 OK**
    ```json
    [
      {
        "id": 1,
        "voucher_id": 1,
        "amount": 100.0,
        "date": "2023-01-01T00:00:00Z"
      }
    ]
    ```

## Error Handling

All error responses will follow the structure:

```json
{
  "error": "Error message"
}
```

## Usage Example

To create a new brand, send a `POST` request to `/brands` with the required JSON body.

## Conclusion

This API provides a simple interface for managing brands, vouchers, and transactions. For further details on the database schema, refer to the `database_schema.md` document.
