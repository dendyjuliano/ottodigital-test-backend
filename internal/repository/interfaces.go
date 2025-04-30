package repository

import (
	"otto-test-go/internal/models"
)

type BrandRepository interface {
    Create(brand *models.Brand) error
    GetByID(id int) (*models.Brand, error)
    GetAll() ([]models.Brand, error) // Add this method
}

type VoucherRepository interface {
    Create(voucher *models.Voucher) error
    GetByBrandID(brandID int) ([]models.Voucher, error)
}