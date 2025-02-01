package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	user2 "PAPER-WALLET-SERVICE-CORE/shared/dto/user"
	"fmt"
	"github.com/shopspring/decimal"
)

func generateWording(lang string, amount, balance decimal.Decimal) string {
	messages := map[string]string{
		"en": "Success! You have successfully requested to disburse %s. Your previous balance was %s, and after the disbursement, your new balance is %s",
		"id": "Sukses! Anda berhasil meminta untuk mencairkan %s. Saldo Anda sebelumnya adalah %s, dan setelah pencairan, saldo baru Anda menjadi %s.",
	}

	newBalance := balance.Sub(amount)
	messageTemplate, exists := messages[lang]
	if !exists {
		messageTemplate = messages["en"]
	}
	return fmt.Sprintf(messageTemplate, amount.String(), balance.String(), newBalance.String())
}

func userToUserDto(user *domain.User) user2.UserDto {
	return user2.UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Currency: user.Currency,
		Balance:  user.Balance,
		Scale:    user.Scale,
	}
}
