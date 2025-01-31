package main

import (
	"PAPER-WALLET-SERVICE-CORE/config"
	_ "PAPER-WALLET-SERVICE-CORE/docs"
	"PAPER-WALLET-SERVICE-CORE/internal/handler/user"
	"PAPER-WALLET-SERVICE-CORE/internal/handler/wallet"
	"PAPER-WALLET-SERVICE-CORE/internal/repository"
	"PAPER-WALLET-SERVICE-CORE/internal/usecase"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config := config.NewConfig()
	userRepository := repository.NewUserRepository(config)
	userUsecase := usecase.NewUserUsecase(userRepository)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	user.NewController(e, userUsecase)
	wallet.NewController(e, userUsecase)

	e.Logger.Fatal(e.Start(":" + config.ServerPort))

}
