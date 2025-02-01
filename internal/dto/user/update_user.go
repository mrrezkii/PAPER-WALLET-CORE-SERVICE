package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
)

type (
	UpdateUserRequestDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             domain.User          `json:"user"`
	}
	UpdateUserResponseDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             domain.User          `json:"user"`
	}
)
