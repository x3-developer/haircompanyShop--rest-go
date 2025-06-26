package docsResponse

import "haircompany-shop-rest/internal/modules/v1/image/dto"

type imageField struct {
	Field     string `json:"field" example:"images"`
	ErrorCode string `json:"errorCode" example:"INVALID_FORMAT"`
}

type ImageUpload200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ImageUpload400 struct {
	IsSuccess bool         `json:"isSuccess" example:"false"`
	Message   string       `json:"message" example:"validation errors occurred"`
	ErrorCode string       `json:"errorCode" enums:"BAD_REQUEST"`
	Fields    []imageField `json:"fields,omitempty"`
}

type ImageUpload413 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"file too large"`
	ErrorCode string `json:"errorCode" enums:"REQUEST_TOO_LARGE"`
}

type ImageUpload500 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"failed to upload image"`
	ErrorCode string `json:"errorCode" enums:"SERVER_ERROR"`
}
