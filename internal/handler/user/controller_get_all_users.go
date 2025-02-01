package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags user-controller
// @Accept json
// @Produce application/json
// @Success 200 {array} domain.User
// @Router /users [get]
func (u *UserController) GetAllUsers(c echo.Context) error {
	var ctx = c.Request().Context()

	users, err := u.userUsecase.Find(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
