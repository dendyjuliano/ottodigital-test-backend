package repository

import (
	"database/sql"
	"otto-test-go/internal/models"
)

type BrandRepositoryImpl struct {
    db *sql.DB
}

func NewBrandRepository(db *sql.DB) BrandRepository {
    return &BrandRepositoryImpl{db: db}
}

func (r *BrandRepositoryImpl) Create(brand *models.Brand) error {
    // MySQL syntax for INSERT and getting last insert ID
    query := `INSERT INTO brands (name) VALUES (?)`
    result, err := r.db.Exec(query, brand.Name)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    brand.ID = int(id)
    return nil
}

func (r *BrandRepositoryImpl) GetByID(id int) (*models.Brand, error) {
    // MySQL uses ? instead of $1 for parameterized queries
    query := `SELECT id, name FROM brands WHERE id = ?`
    brand := &models.Brand{}
    err := r.db.QueryRow(query, id).Scan(&brand.ID, &brand.Name)
    if err != nil {
        return nil, err
    }
    return brand, nil
}

func (r *BrandRepositoryImpl) GetAll() ([]models.Brand, error) {
    query := `SELECT id, name FROM brands`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var brands []models.Brand
    for rows.Next() {
        var b models.Brand
        if err := rows.Scan(&b.ID, &b.Name); err != nil {
            return nil, err
        }
        brands = append(brands, b)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return brands, nil
}