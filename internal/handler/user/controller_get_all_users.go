package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Param X-Channel-Id header string true "Channel identifier" default(iOS)
// @Param X-Request-Id header string true "Unique request identifier" default(RequestId)
// @Param X-Service-Id header string true "Service identifier" default(gateway)
// @Param X-Username header string true "Username associated with the request"
// @Param Accept-Language header string true "Language preference for the response" default(id)
// @Param User-Agent header string true "User agent identifier" default(User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/537.36)
// @Param Authorization header string true "Authorization token" default(Basic dXNlcm5hbWU6cGFzc3dvcmQ=)
// @Param X-App-Version header string false "Application version" default(1.2.3)
// @Tags user-controller
// @Accept json
// @Produce application/json
// @Success 200 {array} domain.User
// @Router /paper-wallet-core-service/users [get]
func (u *UserController) GetAllUsers(c echo.Context) error {
	var ctx = c.Request().Context()

	users, err := u.userUsecase.Find(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
