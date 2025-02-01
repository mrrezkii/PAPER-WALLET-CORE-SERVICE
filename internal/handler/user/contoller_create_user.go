package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags user-controller
// @Accept json
// @Produce application/json
// @Param user body user.CreateUserRequestDto true "User Data"
// @Success 201 {object} user.CreateUserResponseDto
// @Router /paper-wallet-core-service/users [post]
func (u *UserController) CreateUser(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req user.CreateUserRequestDto
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := u.userUsecase.Create(ctx, &req.User)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, req.User)

}
