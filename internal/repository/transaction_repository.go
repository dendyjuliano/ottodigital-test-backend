package repository

import (
	"database/sql"
	"otto-test-go/internal/models"
)

type TransactionRepositoryImpl struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
    return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) CreateTransaction(transaction *models.Transaction) error {
    query := `INSERT INTO transactions 
              (customer_id, total_amount, total_points, status) 
              VALUES (?, ?, ?, ?)`
    
    result, err := r.db.Exec(
        query,
        transaction.CustomerID,
        transaction.TotalAmount,
        transaction.TotalPoints,
        transaction.Status,
    )
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    transaction.ID = int(id)
    return nil
}

func (r *TransactionRepositoryImpl) GetTransactionByID(id int) (*models.Transaction, error) {
    query := `SELECT id, customer_id, total_amount, total_points, status, created_at 
              FROM transactions WHERE id = ?`
    
    transaction := &models.Transaction{}
    err := r.db.QueryRow(query, id).Scan(
        &transaction.ID,
        &transaction.CustomerID,
        &transaction.TotalAmount,
        &transaction.TotalPoints,
        &transaction.Status,
        &transaction.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    
    // Get transaction items
    items, err := r.GetTransactionItemsByTransactionID(id)
    if err != nil {
        return nil, err
    }
    transaction.Items = items
    
    return transaction, nil
}

func (r *TransactionRepositoryImpl) CreateTransactionItem(item *models.TransactionItem) error {
    query := `INSERT INTO transaction_items 
              (transaction_id, voucher_id, quantity, points_used, amount) 
              VALUES (?, ?, ?, ?, ?)`
    
    result, err := r.db.Exec(
        query,
        item.TransactionID,
        item.VoucherID,
        item.Quantity,
        item.PointsUsed,
        item.Amount,
    )
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    item.ID = int(id)
    return nil
}

func (r *TransactionRepositoryImpl) GetTransactionItemsByTransactionID(transactionID int) ([]models.TransactionItem, error) {
    query := `SELECT ti.id, ti.transaction_id, ti.voucher_id, ti.quantity, ti.points_used, ti.amount,
              v.code, v.brand_id, v.discount, v.valid_until
              FROM transaction_items ti
              JOIN vouchers v ON ti.voucher_id = v.id
              WHERE ti.transaction_id = ?`
    
    rows, err := r.db.Query(query, transactionID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var items []models.TransactionItem
    for rows.Next() {
        var item models.TransactionItem
        var voucher models.Voucher
        var validUntil sql.NullTime
        
        err := rows.Scan(
            &item.ID,
            &item.TransactionID,
            &item.VoucherID,
            &item.Quantity,
            &item.PointsUsed,
            &item.Amount,
            &voucher.Code,
            &voucher.BrandID,
            &voucher.Discount,
            &validUntil,
        )
        if err != nil {
            return nil, err
        }
        
        voucher.ID = item.VoucherID
        if validUntil.Valid {
            voucher.ValidUntil = validUntil.Time
        }
        item.Voucher = &voucher
        
        items = append(items, item)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return items, nil
}