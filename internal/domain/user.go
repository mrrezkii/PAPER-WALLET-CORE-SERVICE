package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID string `json:"id"`

	Name     string          `json:"name" validate:"required"`
	Currency string          `json:"currency" validate:"required"`
	Scale    uint8           `json:"scale" validate:"required"`
	Balance  decimal.Decimal `json:"balance" validate:"required"`

	CreatedDate time.Time `json:"createdDate"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedDate time.Time `json:"updatedDate"`
	UpdatedBy   string    `json:"updatedBy"`
	Version     int       `json:"version"`
	IsDeleted   int       `json:"isDeleted"`
}
