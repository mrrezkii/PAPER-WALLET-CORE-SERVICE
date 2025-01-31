package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID string `json:"id"`

	Name     string          `json:"name"`
	Currency string          `json:"currency"`
	Scale    uint32          `json:"scale"`
	Balance  decimal.Decimal `json:"balance"`

	CreatedDate time.Time `json:"createdDate"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedDate time.Time `json:"updatedDate"`
	UpdatedBy   string    `json:"updatedBy"`
	Version     int       `json:"version"`
	IsDeleted   int       `json:"isDeleted"`
}
