package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/repository"
	"PAPER-WALLET-SERVICE-CORE/shared/dto"
	"PAPER-WALLET-SERVICE-CORE/shared/dto/withdraw"
	"context"
	"errors"
	"github.com/shopspring/decimal"
)

type (
	WalletUsecase interface {
		Withdraw(ctx context.Context, request withdraw.WithdrawRequestDto) (withdraw.WithdrawResponseDto, error)
	}
	userUsecase struct {
		repository repository.UserRepository
	}
)

func NewWalletUsecase(repository repository.UserRepository) WalletUsecase {
	return &userUsecase{repository}
}

func (u userUsecase) Withdraw(ctx context.Context, request withdraw.WithdrawRequestDto) (withdraw.WithdrawResponseDto, error) {
	if request.Amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return withdraw.WithdrawResponseDto{}, errors.New("invalid amount")
	}

	user, err := u.repository.FindOne(ctx, map[string]interface{}{
		"ID":        request.UserID,
		"IsDeleted": 0,
	})

	if err != nil {
		return withdraw.WithdrawResponseDto{}, err
	}

	if user == nil {
		return withdraw.WithdrawResponseDto{}, errors.New("user not found")
	}

	if user.Balance.Cmp(request.Amount) <= 0 {
		return withdraw.WithdrawResponseDto{}, errors.New("insufficient balance")
	}

	message := generateWording(request.MandatoryRequest.Language,
		user.Currency,
		user.Scale,
		request.Amount,
		user.Balance)

	user.Balance = user.Balance.Sub(request.Amount)
	err = u.repository.Update(ctx, user)
	if err != nil {
		return withdraw.WithdrawResponseDto{}, err
	}

	return withdraw.WithdrawResponseDto{
		MandatoryRequest: dto.MandatoryRequest{},
		Message:          message,
		Detail:           userToUserDto(user),
	}, nil

}
