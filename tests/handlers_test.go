package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"otto-test-go/internal/api/handlers"
	"otto-test-go/internal/db"
	"otto-test-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    // Connect to DB
    if err := db.Connect(); err != nil {
        panic(err)
    }
    
    // Set up repositories
    brandRepo := repository.NewBrandRepository(db.GetDB())
    voucherRepo := repository.NewVoucherRepository(db.GetDB())
    
    // Set up handlers
    brandHandler := handlers.NewBrandHandler(brandRepo)
    voucherHandler := handlers.NewVoucherHandler(voucherRepo)
    
    // Define routes for testing
    r.POST("/brands", brandHandler.CreateBrand)
    r.GET("/vouchers/:brand_id", voucherHandler.GetVouchersByBrand)
    
    return r
}

// Rename these functions to avoid conflicts with repositories_test.go
func TestCreateBrandHandler(t *testing.T) {  // Changed from TestCreateBrand
    router := setupRouter()

    // Mock request payload
    payload := `{"name": "Test Brand"}`
    req, _ := http.NewRequest("POST", "/brands", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    // Record the response
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert the response
    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetVouchersByBrandHandler(t *testing.T) {  // Changed from TestGetVouchersByBrand
    router := setupRouter()

    // Mock request
    req, _ := http.NewRequest("GET", "/vouchers/1", nil)

    // Record the response
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert the response
    assert.Equal(t, http.StatusOK, w.Code)
}