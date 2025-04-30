package models

import (
	"time"
)

// Transaction represents a voucher redemption transaction
type Transaction struct {
    ID            int       `json:"id"`
    CustomerID    int       `json:"customer_id"` // ID of customer making the redemption
    TotalPoints   int       `json:"total_points"` // Total points used in this transaction
    TotalAmount   float64   `json:"total_amount"` // Total monetary value of the vouchers
    Status        string    `json:"status"` // "completed", "pending", "failed"
    CreatedAt     time.Time `json:"created_at"`
    Items         []TransactionItem `json:"items,omitempty"`
}

// TransactionItem represents an individual voucher in a redemption transaction
type TransactionItem struct {
    ID            int       `json:"id"`
    TransactionID int       `json:"transaction_id"`
    VoucherID     int       `json:"voucher_id"`
    Quantity      int       `json:"quantity"`
    PointsUsed    int       `json:"points_used"`
    Amount        float64   `json:"amount"`
    Voucher       *Voucher  `json:"voucher,omitempty"`
}

// RedemptionRequest represents the request payload for voucher redemption
type RedemptionRequest struct {
    CustomerID  int                    `json:"customer_id" binding:"required"`
    VoucherItems []RedemptionVoucherItem `json:"voucher_items" binding:"required,dive"`
}

// RedemptionVoucherItem represents an individual voucher item in the redemption request
type RedemptionVoucherItem struct {
    VoucherID int `json:"voucher_id" binding:"required"`
    Quantity  int `json:"quantity" binding:"required,min=1"`
}

// TransactionResponse represents the response for a completed transaction
type TransactionResponse struct {
    Transaction Transaction       `json:"transaction"`
    Items       []TransactionItem `json:"items"`
}