package user

import "github.com/shopspring/decimal"

type UserDto struct {
	ID string `json:"id"`

	Name     string          `json:"name" validate:"required"`
	Currency string          `json:"currency" validate:"required"`
	Scale    uint32          `json:"scale" validate:"required"`
	Balance  decimal.Decimal `json:"balance" validate:"required"`
}
