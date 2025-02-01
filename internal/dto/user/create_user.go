package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
)

type (
	CreateUserRequestDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             UserDto              `json:"user"`
	}
	CreateUserResponseDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             UserDto              `json:"user"`
	}
)
