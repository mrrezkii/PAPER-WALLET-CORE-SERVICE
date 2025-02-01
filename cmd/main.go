package main

import (
	"PAPER-WALLET-SERVICE-CORE/config"
	_ "PAPER-WALLET-SERVICE-CORE/docs"
	"PAPER-WALLET-SERVICE-CORE/internal/handler"
	"PAPER-WALLET-SERVICE-CORE/internal/handler/user"
	"PAPER-WALLET-SERVICE-CORE/internal/handler/wallet"
	"PAPER-WALLET-SERVICE-CORE/internal/repository"
	user2 "PAPER-WALLET-SERVICE-CORE/internal/usecase/user"
	wallet2 "PAPER-WALLET-SERVICE-CORE/internal/usecase/wallet"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config := config.NewConfig()
	userRepository := repository.NewUserRepository(config)

	userUsecase := user2.NewUserUsecase(userRepository)
	walletUsecase := wallet2.NewWalletUsecase(userRepository)

	e := echo.New()
	e.Use(handler.RegisterMiddleware)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	user.NewController(e, userUsecase)
	wallet.NewController(e, walletUsecase)

	e.Logger.Fatal(e.Start(":" + config.ServerPort))

}
