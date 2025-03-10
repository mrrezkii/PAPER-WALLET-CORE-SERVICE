package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/handler"
	"PAPER-WALLET-SERVICE-CORE/shared/dto/user"
	"PAPER-WALLET-SERVICE-CORE/shared/response"
	"github.com/labstack/echo/v4"
)

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update the details of an existing user based on the provided user ID and data
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
// @Param user body user.UpdateUserRequestDto true "User Data for Update"
// @Success 200 {object} domain.User
// @Router /paper-wallet-core-service/users [put]
func (u *UserController) UpdateUser(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req user.UpdateUserRequestDto
	)

	req.MandatoryRequest = handler.MandatoryRequest(ctx)
	if err := c.Bind(&req); err != nil {
		return handler.Response(c, nil, response.NewResponseStandard(response.BAD_REQUEST, err))
	}

	if err := c.Validate(req); err != nil {
		return handler.Response(c, nil, response.NewResponseStandard(response.BAD_REQUEST, err))
	}

	toUser := userDtoToUser(req.User)
	err := u.userUsecase.Update(ctx, &toUser)
	if err != nil {
		return handler.Response(c, nil, response.NewResponseStandard(response.SYSTEM_ERROR, err))
	}

	return handler.Response(c, nil, nil)

}
