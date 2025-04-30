package handlers

import (
	"encoding/json"
	"net/http"
	"otto-test-go/internal/models"
	"otto-test-go/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type VoucherHandler struct {
    VoucherRepo repository.VoucherRepository
}

func NewVoucherHandler(voucherRepo repository.VoucherRepository) *VoucherHandler {
    return &VoucherHandler{VoucherRepo: voucherRepo}
}

// Gin handlers
func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
    var voucher models.Voucher
    if err := c.ShouldBindJSON(&voucher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.VoucherRepo.Create(&voucher); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create voucher"})
        return
    }

    c.JSON(http.StatusCreated, voucher)
}

func (h *VoucherHandler) GetVouchersByBrand(c *gin.Context) {
    brandIDStr := c.Param("brand_id")
    brandID, err := strconv.Atoi(brandIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
        return
    }
    
    vouchers, err := h.VoucherRepo.GetByBrandID(brandID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve vouchers"})
        return
    }

    c.JSON(http.StatusOK, vouchers)
}

// Gorilla Mux handlers
func (h *VoucherHandler) CreateVoucherMux(w http.ResponseWriter, r *http.Request) {
    var voucher models.Voucher
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&voucher); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := h.VoucherRepo.Create(&voucher); err != nil {
        http.Error(w, "Could not create voucher", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(voucher)
}

func (h *VoucherHandler) GetVouchersByBrandMux(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    brandIDStr := vars["brandId"]
    brandID, err := strconv.Atoi(brandIDStr)
    if err != nil {
        http.Error(w, "Invalid brand ID", http.StatusBadRequest)
        return
    }
    
    vouchers, err := h.VoucherRepo.GetByBrandID(brandID)
    if err != nil {
        http.Error(w, "Could not retrieve vouchers", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(vouchers)
}