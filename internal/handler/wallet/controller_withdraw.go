package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto/withdraw"
	"github.com/labstack/echo/v4"
	"net/http"
)

// All godoc
// @Summary Withdraw funds from wallet
// @Description Withdraw funds from the user's wallet
// @Tags wallet-controller
// @Accept json
// @Produce application/json
// @Param request body withdraw.WithdrawRequestDto true "Withdraw Request"
// @Success 200 {object} map[string]interface{} "Withdrawal Successful"
// @Router /paper-wallet-core-service/wallet/withdraw [post]
func (u *WalletController) Withdraw(c echo.Context) error {
	var (
		req withdraw.WithdrawRequestDto
	)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Withdrawal successful",
		"userId":  req.UserID,
		"amount":  req.Amount,
	})

}
