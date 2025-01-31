package domain

import "PAPER-WALLET-SERVICE-CORE/internal/dto"

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`

	dto.BaseTableFields `json:"-,inline"`
}
