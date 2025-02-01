package response

import (
	"net/http"
)

type (
	ResponseStandard interface {
		Error() string
		GetCode() string
		GetMessage() string
		GetErrors() []string
		GetHTTPStatusCode() int
	}

	errorStandard struct {
		Errors     []string
		Code       string
		Message    string
		HTTPStatus int
	}

	ResponseCode struct {
		Code           string `json:"code"`
		Message        string `json:"message"`
		httpStatusCode int
	}
)

const (
	SUCCESS        = "SUCCESS"
	DATA_NOT_EXIST = "DATA_NOT_EXIST"
	BAD_REQUEST    = "BAD_REQUEST"
	SYSTEM_ERROR   = "SYSTEM_ERROR"
)

var (
	responseCodes = map[string]ResponseCode{
		SUCCESS: {
			Code:           SUCCESS,
			Message:        "SUCCESS",
			httpStatusCode: http.StatusOK,
		},
		DATA_NOT_EXIST: {
			Code:           DATA_NOT_EXIST,
			Message:        "No Data Exist",
			httpStatusCode: http.StatusNotFound,
		},
		BAD_REQUEST: {
			Code:           BAD_REQUEST,
			Message:        "Bad Request",
			httpStatusCode: http.StatusBadRequest,
		},
		SYSTEM_ERROR: {
			Code:           SYSTEM_ERROR,
			Message:        "System Error",
			httpStatusCode: http.StatusInternalServerError,
		},
	}
)

func NewResponseStandard(code string, err error) ResponseStandard {
	if code == SUCCESS {
		resCode := responseCodes[SUCCESS].Code
		resMessage := responseCodes[SUCCESS].Message
		resHttpStatusCode := responseCodes[SUCCESS].httpStatusCode

		return &errorStandard{
			Errors:     []string{},
			Code:       resCode,
			Message:    resMessage,
			HTTPStatus: resHttpStatusCode,
		}
	}

	errCode := responseCodes[SYSTEM_ERROR].Code
	errMessage := responseCodes[SYSTEM_ERROR].Message
	errHttpStatusCode := responseCodes[SYSTEM_ERROR].httpStatusCode
	errors := make([]string, 0)

	if data, ok := responseCodes[code]; ok {
		errCode = data.Code
		errMessage = data.Message
		errHttpStatusCode = data.httpStatusCode

		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	return &errorStandard{
		Errors:     errors,
		Code:       errCode,
		Message:    errMessage,
		HTTPStatus: errHttpStatusCode,
	}
}

func (e errorStandard) Error() string {
	err := e.Errors
	if len(err) > 0 {
		return err[0]
	}
	return ""
}

func (e errorStandard) GetCode() string {
	return e.Code
}

func (e errorStandard) GetMessage() string {
	return e.Message
}

func (e errorStandard) GetErrors() []string {
	return e.Errors
}

func (e errorStandard) GetHTTPStatusCode() int {
	return e.HTTPStatus
}
