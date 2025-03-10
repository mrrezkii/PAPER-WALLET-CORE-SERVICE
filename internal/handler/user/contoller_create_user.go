package user

import (
	"PAPER-WALLET-SERVICE-CORE/internal/handler"
	"PAPER-WALLET-SERVICE-CORE/shared/dto/user"
	"PAPER-WALLET-SERVICE-CORE/shared/response"
	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
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
// @Param user body user.CreateUserRequestDto true "User Data"
// @Success 201 {object} user.CreateUserResponseDto
// @Router /paper-wallet-core-service/users [post]
func (u *UserController) CreateUser(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req user.CreateUserRequestDto
	)

	req.MandatoryRequest = handler.MandatoryRequest(ctx)
	if err := c.Bind(&req); err != nil {
		return handler.Response(c, nil, response.NewResponseStandard(response.BAD_REQUEST, err))
	}

	if err := c.Validate(req); err != nil {
		return handler.Response(c, nil, response.NewResponseStandard(response.BAD_REQUEST, err))
	}

	toUser := userDtoToUser(req.User)
	err := u.userUsecase.Create(ctx, &toUser)
	if err != nil {
		return handler.Response(c, nil, response.NewResponseStandard(response.SYSTEM_ERROR, err))
	}

	return handler.Response(c, nil, nil)

}
