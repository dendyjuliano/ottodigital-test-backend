package repository

import (
	"otto-test-go/internal/models"
)

type BrandRepository interface {
    Create(brand *models.Brand) error
    GetByID(id int) (*models.Brand, error)
    GetAll() ([]models.Brand, error)
}

type VoucherRepository interface {
    Create(voucher *models.Voucher) error
    GetByBrandID(brandID int) ([]models.Voucher, error)
    GetByID(id int) (*models.Voucher, error)
}

type TransactionRepository interface {
    CreateTransaction(transaction *models.Transaction) error
    GetTransactionByID(id int) (*models.Transaction, error)
    CreateTransactionItem(item *models.TransactionItem) error
    GetTransactionItemsByTransactionID(transactionID int) ([]models.TransactionItem, error)
}