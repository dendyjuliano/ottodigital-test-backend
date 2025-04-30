package models

import (
	"fmt"
	"time"
)

type Voucher struct {
    ID         int       `json:"id"`
    Code       string    `json:"code"`
    BrandID    int       `json:"brand_id"`
    Discount   float64   `json:"discount"`
    Points     int       `json:"points"`          // Points required to redeem this voucher
    ValidUntil time.Time `json:"valid_until,omitempty"`
    Redeemed   bool      `json:"redeemed"`        // Indicates if voucher has been redeemed
}

func (v *Voucher) Validate() error {
    if v.Code == "" {
        return fmt.Errorf("voucher code cannot be empty")
    }
    if v.BrandID <= 0 {
        return fmt.Errorf("invalid brand ID")
    }
    if v.Discount <= 0 {
        return fmt.Errorf("discount must be greater than zero")
    }
    if v.Points < 0 {
        return fmt.Errorf("points cannot be negative")
    }
    return nil
}