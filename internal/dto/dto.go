package dto

type (
	MandatoryRequest struct {
		ChannelID     string `json:"X-Channel-Id" validate:"required"`
		RequestID     string `json:"X-Request-Id" validate:"required"`
		ServiceID     string `json:"X-Service-Id" validate:"required"`
		Username      string `json:"X-Username" validate:"required"`
		Language      string `json:"Accept-Language" validate:"required"`
		UserAgent     string `json:"User-Agent" validate:"required"`
		Authorization string `json:"Authorization" validate:"required"`
		AppVersion    string `json:"X-App-Version,omitempty"`
	}

	BaseResponse struct {
		Code        string      `json:"code"`
		Message     string      `json:"message"`
		Data        interface{} `json:"data"`
		Errors      []string    `json:"errors"`
		ServiceTime int64       `json:"serviceTime"`
	}
)
