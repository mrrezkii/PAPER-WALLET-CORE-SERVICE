package handler

import (
	"PAPER-WALLET-SERVICE-CORE/shared/dto"
	"PAPER-WALLET-SERVICE-CORE/shared/response"
	"github.com/labstack/echo/v4"
	"time"
)

func Response(ec echo.Context, data interface{}, err error) error {
	if err == nil {
		successResponse := response.NewResponseStandard(response.SUCCESS, nil)
		return ec.JSON(successResponse.GetHTTPStatusCode(), &dto.BaseResponse{
			Code:       successResponse.GetCode(),
			Message:    successResponse.GetMessage(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
	if s, ok := err.(response.ResponseStandard); ok {
		return ec.JSON(s.GetHTTPStatusCode(), &dto.BaseResponse{
			Code:       s.GetCode(),
			Message:    s.GetMessage(),
			Errors:     s.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	} else {
		errResponse := response.NewResponseStandard(response.SYSTEM_ERROR, err)
		return ec.JSON(errResponse.GetHTTPStatusCode(), &dto.BaseResponse{
			Code:       errResponse.GetCode(),
			Message:    errResponse.GetMessage(),
			Errors:     errResponse.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
}
