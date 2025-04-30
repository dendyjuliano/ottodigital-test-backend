CREATE TABLE IF NOT EXISTS transaction_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT NOT NULL,
    voucher_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    points_used INT NOT NULL DEFAULT 0,
    amount DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (voucher_id) REFERENCES vouchers(id)
);