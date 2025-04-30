# Database Schema Documentation

## Brands Table

### Table Name: `brands`

| Column Name | Data Type    | Constraints                 | Description                       |
| ----------- | ------------ | --------------------------- | --------------------------------- |
| id          | INT          | PRIMARY KEY, AUTO_INCREMENT | Unique identifier for each brand. |
| name        | VARCHAR(255) | NOT NULL                    | Name of the brand.                |

## Vouchers Table

### Table Name: `vouchers`

| Column Name | Data Type     | Constraints                                 | Description                                 |
| ----------- | ------------- | ------------------------------------------- | ------------------------------------------- |
| id          | INT           | PRIMARY KEY, AUTO_INCREMENT                 | Unique identifier for each voucher.         |
| code        | VARCHAR(50)   | NOT NULL                                    | Unique code for the voucher.                |
| discount    | DECIMAL(10,2) | NOT NULL                                    | Discount value associated with the voucher. |
| brand_id    | INT           | NOT NULL, FOREIGN KEY REFERENCES brands(id) | Identifier for the associated brand.        |
| valid_until | DATETIME      |                                             | Expiration date of the voucher.             |

## Transactions Table

### Table Name: `transactions`

| Column Name  | Data Type     | Constraints                         | Description                                             |
| ------------ | ------------- | ----------------------------------- | ------------------------------------------------------- |
| id           | INT           | PRIMARY KEY, AUTO_INCREMENT         | Unique identifier for each transaction.                 |
| customer_id  | INT           | NOT NULL                            | Identifier for the customer.                            |
| total_amount | DECIMAL(10,2) | NOT NULL, DEFAULT 0.00              | Total amount of the transaction.                        |
| total_points | INT           | NOT NULL, DEFAULT 0                 | Total points used in the transaction.                   |
| status       | VARCHAR(20)   | NOT NULL, DEFAULT 'pending'         | Status of the transaction (pending, completed, failed). |
| created_at   | TIMESTAMP     | NOT NULL, DEFAULT CURRENT_TIMESTAMP | Timestamp when the transaction was created.             |

## Transaction Items Table

### Table Name: `transaction_items`

| Column Name    | Data Type     | Constraints                                       | Description                                |
| -------------- | ------------- | ------------------------------------------------- | ------------------------------------------ |
| id             | INT           | PRIMARY KEY, AUTO_INCREMENT                       | Unique identifier for each item.           |
| transaction_id | INT           | NOT NULL, FOREIGN KEY REFERENCES transactions(id) | Identifier for the associated transaction. |
| voucher_id     | INT           | NOT NULL, FOREIGN KEY REFERENCES vouchers(id)     | Identifier for the voucher.                |
| quantity       | INT           | NOT NULL, DEFAULT 1                               | Quantity of vouchers redeemed.             |
| points_used    | INT           | NOT NULL, DEFAULT 0                               | Points used for this item.                 |
| amount         | DECIMAL(10,2) | NOT NULL, DEFAULT 0.00                            | Amount for this item.                      |

## Relationships

- Each voucher is associated with one brand, indicated by the `brand_id` foreign key in the `vouchers` table.
- A brand can have multiple vouchers, establishing a one-to-many relationship between brands and vouchers.
- Each transaction item is associated with one transaction and one voucher.
- A transaction can have multiple transaction items, establishing a one-to-many relationship between transactions and transaction items.
- A voucher can be associated with multiple transaction items, establishing a one-to-many relationship between vouchers and transaction items.

## SQL Migrations

The database schema is managed using migration files in the `/migrations` directory. Key migrations include:

1. `000001_create_brands_table` - Creates the brands table
2. `000002_create_vouchers_table` - Creates the vouchers table
3. `000003_create_transactions_table` - Creates the transactions table
4. `000004_create_transaction_items_table` - Creates the transaction_items table

Each migration includes `up.sql` (for applying the migration) and `down.sql` (for reverting the migration).
