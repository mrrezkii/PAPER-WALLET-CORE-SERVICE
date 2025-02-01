package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	user2 "PAPER-WALLET-SERVICE-CORE/shared/dto/user"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

func formatCurrency(amount decimal.Decimal, scale uint8) string {
	amountFormatted := amount.Round(int32(scale))
	amountStr := amountFormatted.String()
	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]

	var formattedIntPart string
	for i, digit := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			formattedIntPart += ","
		}
		formattedIntPart += string(digit)
	}

	if len(parts) > 1 {
		return formattedIntPart + "." + parts[1]
	}
	return formattedIntPart
}

func generateWording(lang, currency string, scale uint8, amount, balance decimal.Decimal) string {
	amountFormatted := formatCurrency(amount, scale)
	balanceFormatted := formatCurrency(balance, scale)
	newBalance := balance.Sub(amount).Round(int32(scale))
	newBalanceFormatted := formatCurrency(newBalance, scale)

	messages := map[string]string{
		"en": "Success! You have successfully requested to disburse `%s %s`. Your previous balance was `%s %s`, and after the disbursement, your new balance is `%s %s`",
		"id": "Sukses! Anda berhasil meminta untuk mencairkan `%s %s`. Saldo Anda sebelumnya adalah `%s %s`, dan setelah pencairan, saldo baru Anda menjadi `%s %s`",
	}

	messageTemplate, exists := messages[lang]
	if !exists {
		messageTemplate = messages["en"]
	}

	return fmt.Sprintf(
		messageTemplate,
		currency, amountFormatted,
		currency, balanceFormatted,
		currency, newBalanceFormatted,
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
