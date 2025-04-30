# Database Schema Documentation

## Brands Table

### Table Name: `brands`

| Column Name | Data Type | Constraints         | Description                          |
|-------------|-----------|---------------------|--------------------------------------|
| id          | SERIAL    | PRIMARY KEY         | Unique identifier for each brand.   |
| name        | VARCHAR   | NOT NULL, UNIQUE    | Name of the brand.                   |
| created_at  | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Timestamp when the brand was created. |
| updated_at  | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | Timestamp when the brand was last updated. |

## Vouchers Table

### Table Name: `vouchers`

| Column Name | Data Type | Constraints         | Description                          |
|-------------|-----------|---------------------|--------------------------------------|
| id          | SERIAL    | PRIMARY KEY         | Unique identifier for each voucher. |
| code        | VARCHAR   | NOT NULL, UNIQUE    | Unique code for the voucher.        |
| discount    | DECIMAL   | NOT NULL            | Discount value associated with the voucher. |
| brand_id    | INT       | FOREIGN KEY REFERENCES brands(id) | Identifier for the associated brand. |
| created_at  | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Timestamp when the voucher was created. |
| updated_at  | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | Timestamp when the voucher was last updated. |

## Relationships

- Each voucher is associated with one brand, indicated by the `brand_id` foreign key in the `vouchers` table.
- A brand can have multiple vouchers, establishing a one-to-many relationship between brands and vouchers.