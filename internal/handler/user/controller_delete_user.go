package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID, either hard or soft delete based on the provided flag
// @Tags user-controller
// @Accept json
// @Produce application/json
// @Param user body user.DeleteUserRequestDto true "Delete User Request"
// @Success 204 {object} nil "User successfully deleted"
// @Router /users [delete]

func (u *UserController) DeleteUser(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req user.DeleteUserRequestDto
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var err error
	if req.IsHardDelete {
		err = u.userUsecase.HardDelete(ctx, req.UserID)
	} else {
		err = u.userUsecase.SoftDelete(ctx, req.UserID)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
