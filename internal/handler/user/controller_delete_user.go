package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/dto/user"
	"PAPER-WALLET-SERVICE-CORE/internal/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID, either hard or soft delete based on the provided flag
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
// @Param user body user.DeleteUserRequestDto true "Delete User Request"
// @Success 204 {object} nil "User successfully deleted"
// @Router /paper-wallet-core-service/users [delete]
func (u *UserController) DeleteUser(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req user.DeleteUserRequestDto
	)

	req.MandatoryRequest = handler.MandatoryRequest(ctx)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
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
