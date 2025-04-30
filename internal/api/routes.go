package api

import (
	"otto-test-go/internal/api/handlers"
	"otto-test-go/internal/db"
	"otto-test-go/internal/repository"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the API routes using Gin
func SetupRouter() *gin.Engine {
    router := gin.Default()
    
    // Connect to database
    if err := db.Connect(); err != nil {
        panic(err)
    }
    
    // Set up repositories
    brandRepo := repository.NewBrandRepository(db.GetDB())
    voucherRepo := repository.NewVoucherRepository(db.GetDB())
    transactionRepo := repository.NewTransactionRepository(db.GetDB())  // Add this
    
    // Set up handlers
    brandHandler := handlers.NewBrandHandler(brandRepo)
    voucherHandler := handlers.NewVoucherHandler(voucherRepo)
    transactionHandler := handlers.NewTransactionHandler(transactionRepo, voucherRepo)  // Add this
    
    // Brand routes
    router.POST("/brands", brandHandler.CreateBrand)
    router.GET("/brands/:id", brandHandler.GetBrand)
    router.GET("/brands", brandHandler.GetBrands)
    
    // Voucher routes
    router.POST("/vouchers", voucherHandler.CreateVoucher)
    router.GET("/vouchers/brand/:brand_id", voucherHandler.GetVouchersByBrand)
    
    // Transaction routes - add these new routes
    router.POST("/transaction/redemption", transactionHandler.CreateRedemption)
    router.GET("/transaction/redemption", transactionHandler.GetTransactionDetail)
    
    return router
}