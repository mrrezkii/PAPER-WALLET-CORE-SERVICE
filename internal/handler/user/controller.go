package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/usecase/user"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase user.UserUsecase
}

func NewController(e *echo.Echo, userUsecase user.UserUsecase) {
	controller := &UserController{userUsecase}

	e.GET("/paper-wallet-core-service/users", controller.GetAllUsers)
	e.GET("/paper-wallet-core-service/users:id", controller.GetUser)
	e.POST("/paper-wallet-core-service/users:id", controller.CreateUser)
	e.PUT("/paper-wallet-core-service/users", controller.UpdateUser)
	e.DELETE("/paper-wallet-core-service/users", controller.DeleteUser)
}
