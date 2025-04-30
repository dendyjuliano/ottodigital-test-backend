package tests

import (
	"otto-test-go/internal/db"
	"otto-test-go/internal/models"
	"otto-test-go/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Update any tests that rely on database-specific features

func setupTest() error {
    err := db.Connect()
    if err != nil {
        return err
    }

    // Reset tables for testing
    _, err = db.GetDB().Exec("TRUNCATE TABLE vouchers")
    if err != nil {
        return err
    }
    
    _, err = db.GetDB().Exec("TRUNCATE TABLE brands")
    if err != nil {
        return err
    }
    
    return nil
}

func TestCreateBrand(t *testing.T) {
    err := setupTest()
    if err != nil {
        t.Fatalf("Failed to setup test: %v", err)
    }

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

    brandRepo := repository.NewBrandRepository(db.GetDB())
    voucherRepo := repository.NewVoucherRepository(db.GetDB())

    // Create a brand first
    brand := models.Brand{Name: "Test Brand for Voucher"}
    err = brandRepo.Create(&brand)
    assert.NoError(t, err)

    voucher := models.Voucher{
        Code: "TEST123", 
        BrandID: brand.ID,
        Discount: 10.0,  // Add missing field
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

    brandRepo := repository.NewBrandRepository(db.GetDB())
    voucherRepo := repository.NewVoucherRepository(db.GetDB())

    // Create a brand first
    brand := models.Brand{Name: "Test Brand for Voucher List"}
    err = brandRepo.Create(&brand)
    assert.NoError(t, err)

    voucher1 := models.Voucher{
        Code: "TEST123", 
        BrandID: brand.ID,
        Discount: 10.0,  // Add missing field
    }
    voucher2 := models.Voucher{
        Code: "TEST456", 
        BrandID: brand.ID,
        Discount: 20.0,  // Add missing field
    }
    
    err = voucherRepo.Create(&voucher1)
    assert.NoError(t, err)
    
    err = voucherRepo.Create(&voucher2)
    assert.NoError(t, err)

    vouchers, err := voucherRepo.GetByBrandID(brand.ID)

    assert.NoError(t, err)
    assert.Len(t, vouchers, 2)
}