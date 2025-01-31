package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewController(e *echo.Echo, userUsecase usecase.UserUsecase) {
	controller := &UserController{userUsecase}

	e.GET("/paper-wallet-core-service/users", controller.GetAllUsers)
	e.GET("/paper-wallet-core-service/users:id", controller.GetUser)
	e.POST("/paper-wallet-core-service/users", controller.UpsertUser)
	e.DELETE("/paper-wallet-core-service/users", controller.DeleteUser)
}
