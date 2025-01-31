package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/usecase"
	"github.com/labstack/echo/v4"
)

type WalletController struct {
	userUsecase usecase.UserUsecase
}

func NewController(e *echo.Echo, userUsecase usecase.UserUsecase) {
	controller := &WalletController{userUsecase}

	e.POST("/paper-wallet-core-service/wallet/withdraw", controller.Withdraw)
}
