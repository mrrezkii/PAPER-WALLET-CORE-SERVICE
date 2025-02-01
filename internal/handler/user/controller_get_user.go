package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetUser godoc
// @Summary Get user by ID
// @Description Get details of a user by their ID
// @Tags user-controller
// @Accept json
// @Produce application/json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Router /paper-wallet-core-service/users/{id} [get]
func (u *UserController) GetUser(c echo.Context) error {
	var ctx = c.Request().Context()
	id := c.Param("id")

	if id == "" {
		return echo.ErrNotFound
	}

	user, err := u.userUsecase.FindOne(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
