package withdraw

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
	"github.com/shopspring/decimal"
)

type (
	WithdrawRequestDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		UserID           string               `json:"userId" validate:"required"`
		Amount           decimal.Decimal      `json:"amount" validate:"required"`
	}

	WithdrawResponseDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		UserID           string               `json:"userId"`
		PreviousBalance  decimal.Decimal      `json:"previousBalance"`
		CurrentBalance   decimal.Decimal      `json:"currentBalance"`
	}
)
