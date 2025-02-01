package handler

import (
	"PAPER-WALLET-SERVICE-CORE/helper"
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"strings"
)

func RegisterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		mandatoryRequest := helper.GetMandatoryRequest(c)
		err := c.Validate(&mandatoryRequest)
		if err != nil && isApplySetMandatoryRequest(c.Request().RequestURI) {
			return err
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
