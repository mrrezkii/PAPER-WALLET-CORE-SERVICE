package withdraw

import "PAPER-WALLET-SERVICE-CORE/internal/dto"

type (
	WithdrawRequestDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		UserID           string               `json:"userId" validate:"required"`
		Amount           string               `json:"amount" validate:"required"`
	}

	WithdrawResponseDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		UserID           string               `json:"userId"`
		CurrentBalance   string               `json:"currentBalance"`
	}
)
