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

type BrandHandler struct {
    BrandRepo repository.BrandRepository
}

func NewBrandHandler(brandRepo repository.BrandRepository) *BrandHandler {
    return &BrandHandler{BrandRepo: brandRepo}
}

// Gin handlers
func (h *BrandHandler) CreateBrand(c *gin.Context) {
    var brand models.Brand
    if err := c.ShouldBindJSON(&brand); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.BrandRepo.Create(&brand); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create brand"})
        return
    }

    c.JSON(http.StatusCreated, brand)
}

func (h *BrandHandler) GetBrands(c *gin.Context) {
    brands, err := h.BrandRepo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve brands"})
        return
    }

    c.JSON(http.StatusOK, brands)
}

func (h *BrandHandler) GetBrand(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
        return
    }
    
    brand, err := h.BrandRepo.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
        return
    }

    c.JSON(http.StatusOK, brand)
}

// Gorilla Mux handlers
func (h *BrandHandler) CreateBrandMux(w http.ResponseWriter, r *http.Request) {
    var brand models.Brand
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&brand); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := h.BrandRepo.Create(&brand); err != nil {
        http.Error(w, "Could not create brand", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(brand)
}

func (h *BrandHandler) GetBrandMux(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid brand ID", http.StatusBadRequest)
        return
    }
    
    brand, err := h.BrandRepo.GetByID(id)
    if err != nil {
        http.Error(w, "Brand not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(brand)
}

func (h *BrandHandler) GetBrandsMux(w http.ResponseWriter, r *http.Request) {
    brands, err := h.BrandRepo.GetAll()
    if err != nil {
        http.Error(w, "Failed to retrieve brands", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(brands)
}