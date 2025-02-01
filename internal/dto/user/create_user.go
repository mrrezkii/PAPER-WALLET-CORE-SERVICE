package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
)

type (
	CreateUserRequestDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             domain.User          `json:"user"`
	}
	CreateUserResponseDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             domain.User          `json:"user"`
	}
)
