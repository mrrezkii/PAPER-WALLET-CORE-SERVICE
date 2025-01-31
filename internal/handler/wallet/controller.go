package wallet

import (
	"PAPER-WALLET-SERVICE-CORE/internal/usecase/wallet"
	"github.com/labstack/echo/v4"
)

type WalletController struct {
	walletUsecase wallet.WalletUsecase
}

func NewController(e *echo.Echo, walletUsecase wallet.WalletUsecase) {
	controller := &WalletController{walletUsecase}

	e.POST("/paper-wallet-core-service/wallet/withdraw", controller.Withdraw)
}
