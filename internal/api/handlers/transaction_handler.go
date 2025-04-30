package handlers

import (
	"net/http"
	"otto-test-go/internal/models"
	"otto-test-go/internal/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionRepo repository.TransactionRepository
	VoucherRepo     repository.VoucherRepository
}

func NewTransactionHandler(transactionRepo repository.TransactionRepository, voucherRepo repository.VoucherRepository) *TransactionHandler {
	return &TransactionHandler{
		TransactionRepo: transactionRepo,
		VoucherRepo:     voucherRepo,
	}
}

// CreateRedemption handles POST /transaction/redemption
func (h *TransactionHandler) CreateRedemption(c *gin.Context) {
	// Parse request
	var req models.RedemptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if len(req.VoucherItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No voucher items provided"})
		return
	}

	// Create transaction
	transaction := &models.Transaction{
		CustomerID: req.CustomerID,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	// Calculate totals and validate vouchers
	var totalAmount float64
	var totalPoints int

	// Begin transaction to ensure data consistency
	for _, item := range req.VoucherItems {
		// Get voucher details
		voucher, err := h.VoucherRepo.GetByID(item.VoucherID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid voucher ID: " + strconv.Itoa(item.VoucherID),
			})
			return
		}

		// Check if voucher is valid
		if !voucher.ValidUntil.IsZero() && voucher.ValidUntil.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Voucher has expired: " + voucher.Code,
			})
			return
		}

		// Calculate amount and points
		amount := voucher.Discount * float64(item.Quantity)
		points := int(amount) * 50 // Assuming 1 point = 50 cents

		totalAmount += amount
		totalPoints += points
	}

	// Update transaction totals
	transaction.TotalAmount = totalAmount
	transaction.TotalPoints = totalPoints

	// Save the transaction
	if err := h.TransactionRepo.CreateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Create transaction items
	for _, item := range req.VoucherItems {
		voucher, _ := h.VoucherRepo.GetByID(item.VoucherID)

		amount := voucher.Discount * float64(item.Quantity)
		points := int(amount) * 50 // Same calculation as above

		transactionItem := &models.TransactionItem{
			TransactionID: transaction.ID,
			VoucherID:     item.VoucherID,
			Quantity:      item.Quantity,
			PointsUsed:    points,
			Amount:        amount,
		}

		if err := h.TransactionRepo.CreateTransactionItem(transactionItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction item"})
			return
		}
	}

	// Update transaction status to completed
	transaction.Status = "completed"
	// In a real app, you'd update the transaction status in the database here

	// Get full transaction with items
	fullTransaction, err := h.TransactionRepo.GetTransactionByID(transaction.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction details"})
		return
	}

	c.JSON(http.StatusCreated, fullTransaction)
}

// GetTransactionDetail handles GET /transaction/redemption?transactionId={id}
func (h *TransactionHandler) GetTransactionDetail(c *gin.Context) {
	// Get transaction ID from query params
	transactionIDStr := c.Query("transactionId")
	if transactionIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction ID is required"})
		return
	}

	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	// Get transaction with items
	transaction, err := h.TransactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}