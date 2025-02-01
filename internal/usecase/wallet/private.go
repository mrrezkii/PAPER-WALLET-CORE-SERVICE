package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	user2 "PAPER-WALLET-SERVICE-CORE/shared/dto/user"
	"fmt"
	"github.com/shopspring/decimal"
)

func generateWording(lang, currency string, scale int32, amount, balance decimal.Decimal) string {
	amountFormatted := amount.Round(scale)
	balanceFormatted := balance.Round(scale)
	newBalance := balanceFormatted.Sub(amountFormatted)

	messages := map[string]string{
		"en": "Success! You have successfully requested to disburse %s%s. Your previous balance was %s%s, and after the disbursement, your new balance is %s%s",
		"id": "Sukses! Anda berhasil meminta untuk mencairkan %s%s. Saldo Anda sebelumnya adalah %s%s, dan setelah pencairan, saldo baru Anda menjadi %s%s",
	}

	messageTemplate, exists := messages[lang]
	if !exists {
		messageTemplate = messages["en"]
	}

	return fmt.Sprintf(
		messageTemplate,
		currency, amountFormatted.String(),
		currency, balanceFormatted.String(),
		currency, newBalance.String(),
	)
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
