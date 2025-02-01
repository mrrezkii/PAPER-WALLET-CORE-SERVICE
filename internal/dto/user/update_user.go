package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
)

type (
	UpdateUserRequestDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             UserDto              `json:"user"`
	}
	UpdateUserResponseDto struct {
		MandatoryRequest dto.MandatoryRequest `json:"-"`
		User             UserDto              `json:"user"`
	}
)
