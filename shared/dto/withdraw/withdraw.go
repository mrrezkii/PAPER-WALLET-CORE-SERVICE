package withdraw

import (
	"PAPER-WALLET-SERVICE-CORE/shared/dto"
	user2 "PAPER-WALLET-SERVICE-CORE/shared/dto/user"
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
		Message          string               `json:"message"`
		Detail           user2.UserDto        `json:"detail"`
	}
)
