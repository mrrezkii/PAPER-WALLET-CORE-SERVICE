package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update the details of an existing user based on the provided user ID and data
// @Tags user-controller
// @Accept json
// @Produce application/json
// @Param user body user.UpdateUserRequestDto true "User Data for Update"
// @Success 200 {object} domain.User
// @Router /paper-wallet-core-service/users [put]
func (u *UserController) UpdateUser(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req user.UpdateUserRequestDto
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	toUser := userDtoToUser(req.User)
	err := u.userUsecase.Update(ctx, &toUser)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, req.User)

}
