package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto/withdraw"
	"PAPER-WALLET-SERVICE-CORE/internal/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Withdraw godoc
// @Summary Withdraw funds from wallet
// @Description Withdraw funds from the user's wallet
// @Param X-Channel-Id header string true "Channel identifier" default(iOS)
// @Param X-Request-Id header string true "Unique request identifier" default(RequestId)
// @Param X-Service-Id header string true "Service identifier" default(gateway)
// @Param X-Username header string true "Username associated with the request"
// @Param Accept-Language header string true "Language preference for the response" default(id)
// @Param User-Agent header string true "User agent identifier" default(User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/537.36)
// @Param Authorization header string true "Authorization token" default(Basic dXNlcm5hbWU6cGFzc3dvcmQ=)
// @Param X-App-Version header string false "Application version" default(1.2.3)
// @Tags wallet-controller
// @Accept json
// @Produce application/json
// @Param request body withdraw.WithdrawRequestDto true "Withdraw Request"
// @Success 200 {object} map[string]interface{} "Withdrawal Successful"
// @Router /paper-wallet-core-service/wallet/withdraw [post]
func (u *WalletController) Withdraw(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req withdraw.WithdrawRequestDto
	)

	req.MandatoryRequest = handler.MandatoryRequest(ctx)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	responseDto, err := u.walletUsecase.Withdraw(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responseDto)

}
