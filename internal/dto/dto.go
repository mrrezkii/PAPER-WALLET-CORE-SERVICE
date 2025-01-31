package dto

type (
	MandatoryRequest struct {
		ChannelID     string `json:"channelId" validate:"required"`
		RequestID     string `json:"requestId" validate:"required"`
		ServiceID     string `json:"serviceId" validate:"required"`
		Username      string `json:"username" validate:"required"`
		Language      string `json:"language" validate:"required"`
		UserAgent     string `json:"userAgent" validate:"required"`
		Authorization string `json:"authorization" validate:"required"`
		AppVersion    string `json:"appVersion,omitempty"`
	}

	BaseResponse struct {
		Code        string      `json:"code"`
		Message     string      `json:"message"`
		Data        interface{} `json:"data"`
		Errors      []string    `json:"errors"`
		ServiceTime int64       `json:"serviceTime"`
	}
)
