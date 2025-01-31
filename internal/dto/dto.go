package dto

import "time"

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

	BaseTableFields struct {
		ID          string    `json:"_id"`
		CreatedDate time.Time `json:"createdDate"`
		CreatedBy   string    `json:"createdBy"`
		UpdatedDate time.Time `json:"updatedDate"`
		UpdatedBy   string    `json:"updatedBy"`
		Version     int       `json:"version"`
		IsDeleted   int       `json:"isDeleted"`
	}
)
