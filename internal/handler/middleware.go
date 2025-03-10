package handler

import (
	"PAPER-WALLET-SERVICE-CORE/shared"
	"PAPER-WALLET-SERVICE-CORE/shared/dto"
	"PAPER-WALLET-SERVICE-CORE/shared/response"
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"strings"
)

func RegisterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		mandatoryRequest := shared.GetMandatoryRequest(c)
		err := c.Validate(&mandatoryRequest)
		if err != nil && isApplySetMandatoryRequest(c.Request().RequestURI) {
			return Response(c, nil, response.NewResponseStandard(response.BAD_REQUEST, err))
		}

		if mandatoryRequest.RequestID == "" || mandatoryRequest.RequestID == "RequestId" {
			mandatoryRequest.RequestID = uuid.New().String()
		}
		ctx := setMandatoryRequest(c.Request().Context(), mandatoryRequest)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func setMandatoryRequest(c context.Context, mandatoryRequest dto.MandatoryRequest) context.Context {
	return context.WithValue(c, "MANDATORY_REQUEST", mandatoryRequest)
}

func isApplySetMandatoryRequest(uri string) bool {
	if strings.HasPrefix(uri, "/swagger") {
		return false
	}
	return true
}

func MandatoryRequest(c context.Context) dto.MandatoryRequest {
	mandatoryRequest := c.Value("MANDATORY_REQUEST")
	if mandatoryRequest == nil {
		return dto.MandatoryRequest{}
	}
	return mandatoryRequest.(dto.MandatoryRequest)
}
