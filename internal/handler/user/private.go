package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	user2 "PAPER-WALLET-SERVICE-CORE/shared/dto/user"
)

func userDtoToUser(user user2.UserDto) domain.User {
	return domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Currency: user.Currency,
		Scale:    user.Scale,
		Balance:  user.Balance,
	}
}
