package domain

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
	"github.com/shopspring/decimal"
)

type User struct {
	Name     string          `json:"name"`
	Currency string          `json:"currency"`
	Scale    uint32          `json:"scale"`
	Balance  decimal.Decimal `json:"balance"`

	dto.BaseTableFields `json:"-,inline"`
}
