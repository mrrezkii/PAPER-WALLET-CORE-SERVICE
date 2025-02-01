package user

import (
	"PAPER-WALLET-SERVICE-CORE/shared/dto"
)

type DeleteUserRequestDto struct {
	MandatoryRequest dto.MandatoryRequest `json:"-"`
	UserID           string               `json:"userId" validate:"required"`
	IsHardDelete     bool                 `json:"isHardDelete"`
}

type DeleteUserResponseDto struct {
	MandatoryRequest dto.MandatoryRequest `json:"-"`
}
