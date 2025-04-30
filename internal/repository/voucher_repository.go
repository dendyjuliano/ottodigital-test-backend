package repository

import (
	"database/sql"
	"otto-test-go/internal/models"
)

type VoucherRepositoryImpl struct {
    db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
    return &VoucherRepositoryImpl{db: db}
}

func (r *VoucherRepositoryImpl) Create(voucher *models.Voucher) error {
    // MySQL syntax for INSERT and getting last insert ID
    query := `INSERT INTO vouchers (code, brand_id, discount, valid_until) 
              VALUES (?, ?, ?, ?)`
    
    result, err := r.db.Exec(
        query, 
        voucher.Code, 
        voucher.BrandID, 
        voucher.Discount, 
        voucher.ValidUntil,
    )
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    voucher.ID = int(id)
    return nil
}

func (r *VoucherRepositoryImpl) GetByBrandID(brandID int) ([]models.Voucher, error) {
    // MySQL uses ? instead of $1 for parameterized queries
    query := `SELECT id, code, brand_id, discount, valid_until FROM vouchers WHERE brand_id = ?`
    rows, err := r.db.Query(query, brandID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var vouchers []models.Voucher
    for rows.Next() {
        var v models.Voucher
        var validUntil sql.NullTime
        if err := rows.Scan(&v.ID, &v.Code, &v.BrandID, &v.Discount, &validUntil); err != nil {
            return nil, err
        }
        if validUntil.Valid {
            v.ValidUntil = validUntil.Time
        }
        vouchers = append(vouchers, v)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return vouchers, nil
}

func (r *VoucherRepositoryImpl) GetByID(id int) (*models.Voucher, error) {
    query := `SELECT id, code, brand_id, discount, valid_until FROM vouchers WHERE id = ?`
    
    voucher := &models.Voucher{}
    var validUntil sql.NullTime
    err := r.db.QueryRow(query, id).Scan(
        &voucher.ID, 
        &voucher.Code, 
        &voucher.BrandID, 
        &voucher.Discount, 
        &validUntil,
    )
    if err != nil {
        return nil, err
    }
    
    if validUntil.Valid {
        voucher.ValidUntil = validUntil.Time
    }
    
    return voucher, nil
}