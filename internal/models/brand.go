package models

import (
	"fmt"
)

type Brand struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func (b *Brand) Validate() error {
    if b.Name == "" {
        return fmt.Errorf("brand name cannot be empty")
    }
    return nil
}