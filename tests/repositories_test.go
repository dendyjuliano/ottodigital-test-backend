package tests

import (
	"database/sql"
	"otto-test-go/internal/db"
	"otto-test-go/internal/models"
	"otto-test-go/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var tx *sql.Tx

// Update any tests that rely on database-specific features

func setupTest() error {
    err := db.Connect()
    if err != nil {
        return err
    }

    // Start a transaction
    tx, err = db.GetDB().Begin()
    if err != nil {
        return err
    }

    // Disable foreign key checks
    _, err = tx.Exec("SET FOREIGN_KEY_CHECKS = 0")
    if err != nil {
        return err
    }
    
    // Reset tables for testing
    _, err = tx.Exec("TRUNCATE TABLE vouchers")
    if err != nil {
        return err
    }
    
    _, err = tx.Exec("TRUNCATE TABLE brands")
    if err != nil {
        return err
    }
    
    // Re-enable foreign key checks
    _, err = tx.Exec("SET FOREIGN_KEY_CHECKS = 1")
    if err != nil {
        return err
    }
    
    return nil
}

// Add a teardown function to be called after each test
func teardownTest(t *testing.T) {
    if tx != nil {
        tx.Rollback()
    }
}

func TestCreateBrand(t *testing.T) {
    err := setupTest()
    if err != nil {
        t.Fatalf("Failed to setup test: %v", err)
    }
    defer teardownTest(t)

    repo := repository.NewBrandRepository(db.GetDB())

    brand := models.Brand{Name: "Test Brand"}
    err = repo.Create(&brand)

    assert.NoError(t, err)
    assert.NotEmpty(t, brand.ID)
}

func TestGetBrand(t *testing.T) {
    err := setupTest()
    if err != nil {
        t.Fatalf("Failed to setup test: %v", err)
    }
    defer teardownTest(t)

    repo := repository.NewBrandRepository(db.GetDB())

    brand := models.Brand{Name: "Test Brand for Get"}
    err = repo.Create(&brand)
    assert.NoError(t, err)

    fetchedBrand, err := repo.GetByID(brand.ID)

    assert.NoError(t, err)
    assert.Equal(t, brand.Name, fetchedBrand.Name)
}

func TestCreateVoucher(t *testing.T) {
    err := setupTest()
    if err != nil {
        t.Fatalf("Failed to setup test: %v", err)
    }
    defer teardownTest(t)

    brandRepo := repository.NewBrandRepository(db.GetDB())
    voucherRepo := repository.NewVoucherRepository(db.GetDB())

    // Create a brand first
    brand := models.Brand{Name: "Test Brand for Voucher"}
    err = brandRepo.Create(&brand)
    assert.NoError(t, err)

    voucher := models.Voucher{
        Code: "TEST123", 
        BrandID: brand.ID,
        Discount: 10.0,
        ValidUntil: time.Now().AddDate(1, 0, 0), // 1 year from now
    }
    err = voucherRepo.Create(&voucher)

    assert.NoError(t, err)
    assert.NotEmpty(t, voucher.ID)
}

func TestGetVouchersByBrand(t *testing.T) {
    err := setupTest()
    if err != nil {
        t.Fatalf("Failed to setup test: %v", err)
    }
    defer teardownTest(t)

    brandRepo := repository.NewBrandRepository(db.GetDB())
    voucherRepo := repository.NewVoucherRepository(db.GetDB())

    // Create a brand first
    brand := models.Brand{Name: "Test Brand for Voucher List"}
    err = brandRepo.Create(&brand)
    assert.NoError(t, err)

    voucher1 := models.Voucher{
        Code: "TEST123", 
        BrandID: brand.ID,
        Discount: 10.0,
        ValidUntil: time.Now().AddDate(1, 0, 0), // 1 year from now
    }
    voucher2 := models.Voucher{
        Code: "TEST456", 
        BrandID: brand.ID,
        Discount: 20.0,
        ValidUntil: time.Now().AddDate(1, 0, 0), // 1 year from now
    }
    
    err = voucherRepo.Create(&voucher1)
    assert.NoError(t, err)
    
    err = voucherRepo.Create(&voucher2)
    assert.NoError(t, err)

    vouchers, err := voucherRepo.GetByBrandID(brand.ID)

    assert.NoError(t, err)
    assert.Len(t, vouchers, 2)
}